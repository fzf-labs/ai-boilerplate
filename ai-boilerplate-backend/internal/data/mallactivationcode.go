package data

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewMallActivationCodeRepo(
	logger log.Logger,
	data *Data,
	mallActivationCodeRepo *ai_boilerplate_repo.MallActivationCodeRepo,
) *MallActivationCodeRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mallActivationCode"))
	return &MallActivationCodeRepo{
		log:                    l,
		data:                   data,
		MallActivationCodeRepo: mallActivationCodeRepo,
	}
}

type MallActivationCodeRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MallActivationCodeRepo
}

// GetBatchNo 获取批次号
func (m *MallActivationCodeRepo) GetBatchNo(ctx context.Context) (string, error) {
	date := carbon.Now().Format("Ymd")
	batchNo, err := m.data.rueidis.Do(ctx, m.data.rueidis.B().Incrby().Key(constant.ActivationCodeBatchNo.Key(date)).Increment(1).Build()).ToInt64()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%04d", date, batchNo), nil
}

// GenerateCode 生成激活码
func (m *MallActivationCodeRepo) GenerateCode(ctx context.Context, num int32) ([]string, error) {
	_ = ctx

	// 使用当前时间作为随机种子
	seed := uint64(time.Now().UnixNano())
	rng := rand.New(rand.NewPCG(seed, seed+1))
	// 定义字符集，排除容易混淆的字符（0、O、I、1）
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	codes := make([]string, 0, int(num))
	codeMap := make(map[string]struct{}, int(num)) // 用于去重
	for len(codes) < int(num) {
		var code string
		// 生成16位激活码，每4位用"-"分隔
		for i := 0; i < 16; i++ {
			code += string(chars[rng.IntN(len(chars))])
			if (i+1)%4 == 0 && i < 15 {
				code += "-"
			}
		}
		// 确保激活码唯一
		if _, exists := codeMap[code]; !exists {
			codeMap[code] = struct{}{}
			codes = append(codes, code)
		}
	}
	return codes, nil
}
