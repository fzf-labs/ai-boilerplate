package service

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/uuidutil"
	"github.com/fzf-labs/kratos-contrib/meta"
	"gorm.io/datatypes"
)

const userMessageBatchSize = 500

// SendUserMessage App-用户消息-发送
func (a *AdminV1UserMessageService) SendUserMessage(ctx context.Context, req *pb.SendUserMessageReq) (*pb.SendUserMessageReply, error) {
	resp := &pb.SendUserMessageReply{}
	audienceType := req.GetAudienceType()
	if audienceType == "" {
		return nil, pb.ErrorReasonParamError()
	}

	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)

	userIDs, err := a.resolveAudienceUserIDs(ctx, audienceType, req)
	if err != nil {
		return nil, err
	}
	if len(userIDs) == 0 {
		return nil, pb.ErrorReasonParamError()
	}

	messageID := uuidutil.GenUUID()
	sentAt := time.Now()
	audienceValue := buildAudienceValue(req)

	records := make([]*ai_boilerplate_model.UserMessage, 0, len(userIDs))
	for _, userID := range userIDs {
		records = append(records, &ai_boilerplate_model.UserMessage{
			MessageID:     messageID,
			UserID:        userID,
			Category:      req.GetCategory(),
			Title:         req.GetTitle(),
			Summary:       req.GetSummary(),
			CoverURL:      req.GetCoverURL(),
			Content:       req.GetContent(),
			LinkURL:       req.GetLinkURL(),
			AudienceType:  audienceType,
			AudienceValue: audienceValue,
			SentAt:        sentAt,
			AdminID:       adminID,
		})
	}

	if err := a.userMessageRepo.CreateBatchCache(ctx, records, userMessageBatchSize); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.MessageId = messageID
	resp.Total = int32(len(records))
	return resp, nil
}

func (a *AdminV1UserMessageService) resolveAudienceUserIDs(ctx context.Context, audienceType string, req *pb.SendUserMessageReq) ([]string, error) {
	switch audienceType {
	case "all":
		return a.listUserIDsByCondition(ctx, []*condition.QueryParam{
			{
				Field: "status",
				Value: 1,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		})
	case "users":
		userIDs := uniqueStrings(req.GetUserIds())
		if len(userIDs) == 0 {
			return nil, pb.ErrorReasonParamError()
		}
		return userIDs, nil
	case "segment":
		return a.listSegmentUserIDs(ctx, req)
	default:
		return nil, pb.ErrorReasonParamError()
	}
}

func (a *AdminV1UserMessageService) listSegmentUserIDs(ctx context.Context, req *pb.SendUserMessageReq) ([]string, error) {
	membershipType := req.GetMembershipType()
	activeWithinDays := req.GetActiveWithinDays()
	if membershipType == "" && activeWithinDays == 0 {
		return nil, pb.ErrorReasonParamError()
	}

	var idsByMembership []string
	var err error
	if membershipType != "" {
		idsByMembership, err = a.listUserIDsByMembership(ctx, membershipType)
		if err != nil {
			return nil, err
		}
	}

	if activeWithinDays > 0 {
		return a.listUserIDsByActivity(ctx, activeWithinDays, idsByMembership)
	}

	return uniqueStrings(idsByMembership), nil
}

func (a *AdminV1UserMessageService) listUserIDsByMembership(ctx context.Context, membershipType string) ([]string, error) {
	pageSize := int32(userMessageBatchSize)
	page := int32(1)
	userIDs := make([]string, 0)
	for {
		param := &condition.Req{
			Page:     page,
			PageSize: pageSize,
			Query: []*condition.QueryParam{
				{
					Field: "membership_type",
					Value: membershipType,
					Exp:   condition.EQ,
					Logic: condition.AND,
				},
				{
					Field: "status",
					Value: 1,
					Exp:   condition.EQ,
					Logic: condition.AND,
				},
			},
		}
		list, p, err := a.userMembershipRepo.FindMultiByCondition(ctx, param)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, item := range list {
			userIDs = append(userIDs, item.UserID)
		}
		if p.Total == 0 || page*pageSize >= p.Total {
			break
		}
		page++
	}
	return uniqueStrings(userIDs), nil
}

func (a *AdminV1UserMessageService) listUserIDsByActivity(ctx context.Context, activeWithinDays int32, filterIDs []string) ([]string, error) {
	threshold := time.Now().AddDate(0, 0, -int(activeWithinDays))
	if len(filterIDs) == 0 {
		return a.listUserIDsByCondition(ctx, []*condition.QueryParam{
			{
				Field: "status",
				Value: 1,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "updated_at",
				Value: threshold,
				Exp:   condition.GTE,
				Logic: condition.AND,
			},
		})
	}

	userIDs := make([]string, 0, len(filterIDs))
	for _, chunk := range chunkStrings(uniqueStrings(filterIDs), userMessageBatchSize) {
		param := &condition.Req{
			Query: []*condition.QueryParam{
				{
					Field: "id",
					Value: chunk,
					Exp:   condition.IN,
					Logic: condition.AND,
				},
				{
					Field: "status",
					Value: 1,
					Exp:   condition.EQ,
					Logic: condition.AND,
				},
				{
					Field: "updated_at",
					Value: threshold,
					Exp:   condition.GTE,
					Logic: condition.AND,
				},
			},
		}
		list, _, err := a.userRepo.FindMultiByCondition(ctx, param)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, item := range list {
			userIDs = append(userIDs, item.ID)
		}
	}
	return uniqueStrings(userIDs), nil
}

func (a *AdminV1UserMessageService) listUserIDsByCondition(ctx context.Context, query []*condition.QueryParam) ([]string, error) {
	pageSize := int32(userMessageBatchSize)
	page := int32(1)
	userIDs := make([]string, 0)
	for {
		param := &condition.Req{
			Page:     page,
			PageSize: pageSize,
			Query:    query,
		}
		list, p, err := a.userRepo.FindMultiByCondition(ctx, param)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, item := range list {
			userIDs = append(userIDs, item.ID)
		}
		if p.Total == 0 || page*pageSize >= p.Total {
			break
		}
		page++
	}
	return uniqueStrings(userIDs), nil
}

func buildAudienceValue(req *pb.SendUserMessageReq) datatypes.JSON {
	payload := map[string]any{}
	if len(req.GetUserIds()) > 0 {
		payload["userIds"] = req.GetUserIds()
	}
	if req.GetMembershipType() != "" {
		payload["membershipType"] = req.GetMembershipType()
	}
	if req.GetActiveWithinDays() > 0 {
		payload["activeWithinDays"] = req.GetActiveWithinDays()
	}
	if len(payload) == 0 {
		return datatypes.JSON([]byte("{}"))
	}
	value, err := json.Marshal(payload)
	if err != nil {
		return datatypes.JSON([]byte("{}"))
	}
	return datatypes.JSON(value)
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func chunkStrings(values []string, size int32) [][]string {
	if size <= 0 {
		return [][]string{values}
	}
	chunkSize := int(size)
	result := make([][]string, 0)
	for start := 0; start < len(values); start += chunkSize {
		end := start + chunkSize
		if end > len(values) {
			end = len(values)
		}
		result = append(result, values[start:end])
	}
	return result
}
