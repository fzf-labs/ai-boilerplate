package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSmsChannelRepo(
	logger log.Logger,
	data *Data,
	smsChannelRepo *ai_boilerplate_repo.SmsChannelRepo,
) *SmsChannelRepo {
	l := log.NewHelper(log.With(logger, "module", "data/smsChannel"))
	return &SmsChannelRepo{
		log:            l,
		data:           data,
		SmsChannelRepo: smsChannelRepo,
	}
}

type SmsChannelRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SmsChannelRepo
}

func (r *SmsChannelRepo) IDToName(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	list, err := r.FindMultiCacheByIDS(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp[v.ID] = v.Name
	}
	return resp, nil
}
