package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

const TableNameActivity = "activity"

// Activity mapped from table <activity>
type Activity struct {
	ID       string `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid();comment:id" json:"id"`
	TenantID string `gorm:"column:tenant_id;type:character varying(64);not null;comment:租户id" json:"tenantId"`

	Title    string `gorm:"column:title;type:character varying(200);not null;comment:标题" json:"title"`
	ImageURL string `gorm:"column:image_url;type:character varying(500);not null;comment:图片URL" json:"imageUrl"`
	LinkURL  string `gorm:"column:link_url;type:character varying(500);not null;comment:跳转链接" json:"linkUrl"`
	LinkType string `gorm:"column:link_type;type:character varying(32);not null;comment:跳转类型" json:"linkType"`

	Sort   int32 `gorm:"column:sort;type:integer;not null;comment:排序" json:"sort"`
	Status int16 `gorm:"column:status;type:smallint;not null;comment:状态(-1禁用,1开启)" json:"status"`

	StartTime sql.NullTime   `gorm:"column:start_time;type:timestamp with time zone;comment:开始时间" json:"startTime"`
	EndTime   sql.NullTime   `gorm:"column:end_time;type:timestamp with time zone;comment:结束时间" json:"endTime"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;comment:删除时间" json:"deletedAt"`
}

func (*Activity) TableName() string { return TableNameActivity }

type ActivityListFilter struct {
	TenantID string
	Keyword  string
	Status   int16 // 0 means "all"
	Page     int32
	PageSize int32
}

type AppActivityListFilter struct {
	TenantID string
	Page     int32
	PageSize int32
	Now      time.Time
}

func NewActivityRepo(
	logger log.Logger,
	data *Data,
) *ActivityRepo {
	l := log.NewHelper(log.With(logger, "module", "data/activity"))
	return &ActivityRepo{
		log:  l,
		data: data,
	}
}

type ActivityRepo struct {
	log  *log.Helper
	data *Data
}

func (r *ActivityRepo) NewData() *Activity {
	return &Activity{}
}

func (r *ActivityRepo) Create(ctx context.Context, activity *Activity) error {
	return r.data.gorm.WithContext(ctx).Create(activity).Error
}

func (r *ActivityRepo) GetByID(ctx context.Context, id string) (*Activity, error) {
	var activity Activity
	err := r.data.gorm.WithContext(ctx).First(&activity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *ActivityRepo) Update(ctx context.Context, activity *Activity) error {
	return r.data.gorm.WithContext(ctx).Save(activity).Error
}

func (r *ActivityRepo) Delete(ctx context.Context, id string) error {
	return r.data.gorm.WithContext(ctx).Delete(&Activity{}, "id = ?", id).Error
}

func (r *ActivityRepo) ListForAdmin(ctx context.Context, filter ActivityListFilter) (list []*Activity, total int64, err error) {
	query := r.data.gorm.WithContext(ctx).Model(&Activity{})
	if filter.TenantID != "" {
		query = query.Where("tenant_id = ?", filter.TenantID)
	}
	if filter.Keyword != "" {
		query = query.Where("title ILIKE ?", "%"+filter.Keyword+"%")
	}
	if filter.Status != 0 {
		query = query.Where("status = ?", filter.Status)
	}

	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := int((filter.Page - 1) * filter.PageSize)
	if offset < 0 {
		offset = 0
	}
	var rows []Activity
	err = query.
		Order("updated_at DESC").
		Offset(offset).
		Limit(int(filter.PageSize)).
		Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	for i := range rows {
		list = append(list, &rows[i])
	}
	return list, total, nil
}

func (r *ActivityRepo) ListForApp(ctx context.Context, filter AppActivityListFilter) (list []*Activity, total int64, err error) {
	query := r.data.gorm.WithContext(ctx).Model(&Activity{})
	if filter.TenantID != "" {
		query = query.Where("tenant_id = ?", filter.TenantID)
	}

	now := filter.Now
	query = query.
		Where("status = ?", int16(1)).
		Where("(start_time IS NULL OR start_time <= ?)", now).
		Where("(end_time IS NULL OR end_time >= ?)", now)

	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := int((filter.Page - 1) * filter.PageSize)
	if offset < 0 {
		offset = 0
	}
	var rows []Activity
	err = query.
		Order("sort ASC").
		Order("created_at DESC").
		Offset(offset).
		Limit(int(filter.PageSize)).
		Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	for i := range rows {
		list = append(list, &rows[i])
	}
	return list, total, nil
}
