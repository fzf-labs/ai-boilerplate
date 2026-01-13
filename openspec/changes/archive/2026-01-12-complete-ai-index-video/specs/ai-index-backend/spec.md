## ADDED Requirements

### Requirement: AI index video records are user-scoped
The system SHALL create, list, update, update status, fetch, and delete AI video records via `ai_index_video` endpoints scoped to the current tenant/admin.

#### Scenario: Create a video record
- **WHEN** an authenticated admin calls `/admin/v1/ai_index_video/create`
- **THEN** the record is stored in `ai_video_record` with tenant/admin from request metadata

#### Scenario: Update a video record
- **WHEN** an authenticated admin updates a video record they own
- **THEN** the record is updated and the response indicates success

#### Scenario: Update video status
- **WHEN** an authenticated admin updates the status of a video record they own
- **THEN** the status field is updated

#### Scenario: Fetch a video record
- **WHEN** an authenticated admin calls `/admin/v1/ai_index_video/info` for a record they own
- **THEN** the record details are returned

#### Scenario: List video records
- **WHEN** an authenticated admin calls `/admin/v1/ai_index_video/list`
- **THEN** the response only includes records owned by the current tenant/admin

#### Scenario: Delete a video record
- **WHEN** an authenticated admin deletes a video record they own
- **THEN** the record is removed; otherwise a not-found error is returned
