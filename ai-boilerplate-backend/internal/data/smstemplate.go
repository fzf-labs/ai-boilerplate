package data

import (
	"context"
	"regexp"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewSmsTemplateRepo(
	logger log.Logger,
	data *Data,
	smsTemplateRepo *ai_boilerplate_repo.SmsTemplateRepo,
) *SmsTemplateRepo {
	l := log.NewHelper(log.With(logger, "module", "data/smsTemplate"))
	return &SmsTemplateRepo{
		log:             l,
		data:            data,
		SmsTemplateRepo: smsTemplateRepo,
	}
}

type SmsTemplateRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SmsTemplateRepo
}

func (r *SmsTemplateRepo) IDToName(ctx context.Context, ids []string) (map[string]string, error) {
	resp := make(map[string]string)
	list, err := r.FindMultiCacheByIDS(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp[v.ID] = v.TemplateName
	}
	return resp, nil
}

// 解析TemplateContent中的变量`{}`
func (r *SmsTemplateRepo) ParseTemplateParams(templateContent string) map[string]string {
	matches := regexp.MustCompile(`\{([^}]+)\}`).FindAllStringSubmatch(templateContent, -1)
	templateParams := make(map[string]string, 0)
	for _, match := range matches {
		templateParams[match[1]] = ""
	}
	return templateParams
}
