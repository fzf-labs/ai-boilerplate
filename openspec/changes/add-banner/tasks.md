## 1. Implementation
- [x] 1.1 Add `banner` SQL schema in `ai-boilerplate-backend/doc/sql/ai_boilerplate/banner.sql`
- [x] 1.2 Generate GORM model for `banner`
- [x] 1.3 Generate admin and app proto files from `banner` (sqltopb)
- [x] 1.4 Edit proto to add filters/fields for admin list and app active list
- [x] 1.5 Generate API code (pb.go/grpc/http) for admin and app
- [x] 1.6 Implement admin CRUD/status services for banners
- [x] 1.7 Implement app banner list service with active filtering
- [x] 1.8 Update `AppV1HomeService.GetBannerList` to read from banner data (position "home")
- [ ] 1.9 Run backend quality checks
