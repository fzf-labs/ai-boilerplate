# Service Per-Method File Examples (Based on Existing Code)

> When to read: Only read this file during **Step 6 (Implement business logic)** when you need a copyable per-file skeleton/pattern for a service method.

Note: The snippets below are **adaptable skeletons**. They may omit `import`s, full field mapping, or full `resp.List` construction. Always follow the referenced implementation file as the source of truth.

## Contents

- Single record (ID from context metadata)
- List (pagination + filters + ordering)
- Related lookup (batch query + map backfill)
- Create (NewData → assign → CreateOneCache → return ID)
- Status update (load → nil-check → DeepCopy → UpdateOneCacheWithZero)
- Delete (delete by ID)

---

## Single record (ID from context metadata)

Reference: `ai-boilerplate-backend/internal/service/app_v1_user_getuserinfo.go`

Use for: App APIs that fetch “current user” data bound to auth metadata.

```go
func (a *AppV1UserService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoReply, error) {
	resp := &pb.GetUserInfoReply{}

	userId := meta.GetMetadataFromClient(ctx, constant.XMdUserId)
	data, err := a.userRepo.FindOneCacheByID(ctx, userId)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Info = &pb.UserInfo{
		Id:        data.ID,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
```

---

## List (pagination + filters + ordering)

Reference: `ai-boilerplate-backend/internal/service/admin_v1_syspost_getsyspostlist.go`

Use for: Admin list pages with pagination, filters, ordering, and time ranges.

```go
resp := &pb.GetXxxListReply{Total: 0, List: []*pb.XxxInfo{}}
param := &condition.Req{
	Page:     req.GetPage(),
	PageSize: req.GetPageSize(),
	Query:    []*condition.QueryParam{},
	Order:    []*condition.OrderParam{{Field: "created_at", Order: condition.DESC}},
}
if req.GetName() != "" {
	param.Query = append(param.Query, &condition.QueryParam{
		Field: "name", Value: "%" + req.GetName() + "%", Exp: condition.LIKE, Logic: condition.AND,
	})
}
list, p, err := a.xxxRepo.FindMultiCacheByCondition(ctx, param)
if err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
resp.Total = int32(p.Total)
```

---

## Related lookup (batch query + map backfill)

Reference: `ai-boilerplate-backend/internal/service/admin_v1_user_getuserlist.go`

Use for: lists that need extra related info (e.g., membership/role) per row.

```go
list, p, err := a.userRepo.FindMultiCacheByCondition(ctx, param)
if err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
resp.Total = int32(p.Total)

userIds := make([]string, 0, len(list))
for _, v := range list {
	userIds = append(userIds, v.ID)
}
userMembershipList, err := a.userMembershipRepo.FindMultiCacheByUserIDS(ctx, userIds)
if err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
userMembershipMap := make(map[string]*pb.UserMembershipInfo, len(userMembershipList))
for _, v := range userMembershipList {
	userMembershipMap[v.UserID] = &pb.UserMembershipInfo{Id: v.ID, UserId: v.UserID}
}
// backfill: build resp.List items and attach userMembershipMap[v.ID]
```

---

## Create (NewData → assign → CreateOneCache → return ID)

Reference: `ai-boilerplate-backend/internal/service/admin_v1_aiaudiorecord_createaiaudiorecord.go`

Use for: standard create endpoints.

```go
resp := &pb.CreateXxxReply{}
data := a.xxxRepo.NewData()
data.Title = req.GetTitle()
// ... assign other fields
if err := a.xxxRepo.CreateOneCache(ctx, data); err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
resp.Id = data.ID
return resp, nil
```

---

## Status update (load → nil-check → DeepCopy → UpdateOneCacheWithZero)

Reference: `ai-boilerplate-backend/internal/service/admin_v1_aiaudiorecord_updateaiaudiorecordstatus.go`

Use for: updating a single field like status while keeping an old-data snapshot.

```go
resp := &pb.UpdateXxxStatusReply{}
data, err := a.xxxRepo.FindOneCacheByID(ctx, req.GetId())
if err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
if data == nil || data.ID == "" {
	return nil, pb.ErrorReasonDataRecordNotFound()
}
oldData := a.xxxRepo.DeepCopy(data)
data.Status = req.GetStatus()
if err := a.xxxRepo.UpdateOneCacheWithZero(ctx, data, oldData); err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
return resp, nil
```

---

## Delete (delete by ID)

Reference: `ai-boilerplate-backend/internal/service/admin_v1_aiaudiorecord_deleteaiaudiorecord.go`

Use for: standard delete endpoints.

```go
resp := &pb.DeleteXxxReply{}
if err := a.xxxRepo.DeleteOneCacheByID(ctx, req.GetId()); err != nil {
	return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
}
return resp, nil
```

