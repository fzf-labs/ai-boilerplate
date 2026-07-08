package testfixture

import (
	"database/sql"
	"time"

	adminpb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	apppb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
)

const (
	AdminID   = "fixture-admin-001"
	TenantID  = "fixture-tenant-001"
	RoleID    = "fixture-role-001"
	UserID    = "fixture-user-001"
	OrderID   = "fixture-order-001"
	ProductID = "fixture-product-001"
)

func SysAdmin() *ai_boilerplate_model.SysAdmin {
	return &ai_boilerplate_model.SysAdmin{
		ID:       AdminID,
		TenantID: TenantID,
		Username: "fixture_admin",
		Nickname: "Fixture Admin",
		RoleID:   RoleID,
		Status:   int16(constant.StatusEnable),
	}
}

func SysMenus() []*ai_boilerplate_model.SysMenu {
	return []*ai_boilerplate_model.SysMenu{
		{
			ID:     "menu-root",
			Name:   "Root",
			Type:   constant.SysMenuTypeDir.String(),
			Path:   "/root",
			Sort:   1,
			Status: int16(constant.StatusEnable),
		},
		{
			ID:         "menu-child",
			Pid:        "menu-root",
			Name:       "Dashboard",
			Type:       constant.SysMenuTypeMenu.String(),
			Path:       "/dashboard",
			Permission: "dashboard:view",
			Sort:       2,
			Status:     int16(constant.StatusEnable),
		},
		{
			ID:         "menu-button",
			Pid:        "menu-child",
			Name:       "Create",
			Type:       constant.SysMenuTypeButton.String(),
			Permission: "dashboard:create",
			Sort:       3,
			Status:     int16(constant.StatusEnable),
		},
		{
			ID:         "menu-disabled",
			Pid:        "menu-root",
			Name:       "Disabled",
			Type:       constant.SysMenuTypeMenu.String(),
			Path:       "/disabled",
			Permission: "disabled:view",
			Sort:       4,
			Status:     int16(constant.StatusDisable),
		},
		{
			ID:         "menu-duplicate",
			Pid:        "menu-root",
			Name:       "Duplicate Permission",
			Type:       constant.SysMenuTypeButton.String(),
			Permission: "dashboard:view",
			Sort:       5,
			Status:     int16(constant.StatusEnable),
		},
	}
}

func SmsChannel() *ai_boilerplate_model.SmsChannel {
	now := fixedTime()
	return &ai_boilerplate_model.SmsChannel{
		ID:          "fixture-sms-channel-001",
		Name:        "Fixture SMS",
		Operator:    constant.SmsChannelCodeALIYUN.String(),
		Remark:      "fixture only",
		APIKey:      "ak_fixture_1234567890",
		APISecret:   "sk_fixture_0987654321",
		CallbackURL: "https://example.invalid/sms/callback",
		Status:      int16(constant.StatusEnable),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func MallOrder(productType string) *ai_boilerplate_model.MallOrder {
	now := fixedTime()
	return &ai_boilerplate_model.MallOrder{
		ID:             OrderID,
		UserID:         UserID,
		ProductType:    productType,
		ProductID:      ProductID,
		OriginalAmount: 100,
		DiscountAmount: 20,
		ActualAmount:   80,
		RefundAmount:   0,
		Currency:       "CNY",
		PaymentMethod:  "wechat",
		PaymentStatus:  0,
		Status:         "pendingPayment",
		ExpiredTime:    sql.NullTime{Time: now.Add(30 * time.Minute), Valid: true},
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

func PaymentCallbackReq(orderID string, status int32) *apppb.PaymentCallbackReq {
	return &apppb.PaymentCallbackReq{
		OrderId:       orderID,
		PaymentStatus: status,
		TransactionId: "txn_fixture_001",
		PaymentMethod: "wechat",
		CallbackData:  `{"fixture":true}`,
	}
}

func SysAuthLoginReq() *adminpb.SysAuthLoginReq {
	return &adminpb.SysAuthLoginReq{
		Username: "fixture_admin",
		Password: "fixture-password",
	}
}

func fixedTime() time.Time {
	return time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
}
