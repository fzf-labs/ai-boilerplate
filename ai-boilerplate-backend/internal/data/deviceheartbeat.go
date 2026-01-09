package data

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dromara/carbon/v2"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/rueidis"
)

// NewDeviceHeartbeatRepo 创建设备心跳仓库
func NewDeviceHeartbeatRepo(logger log.Logger, data *Data) *DeviceHeartbeatRepo {
	return &DeviceHeartbeatRepo{
		logger: log.NewHelper(log.With(logger, "module", "data/deviceheartbeat")),
		data:   data,
	}
}

// DeviceHeartbeatRepo 设备心跳仓库实现
type DeviceHeartbeatRepo struct {
	logger *log.Helper
	data   *Data
}

// RecordHeartbeat 记录设备心跳
func (r *DeviceHeartbeatRepo) RecordHeartbeat(ctx context.Context, sn string) error {
	now := carbon.Now()
	results := r.data.rueidis.DoMulti(ctx,
		r.data.rueidis.B().Set().Key(constant.DeviceHeartbeat.Key(sn)).Value(now.ToDateString()).Ex(constant.DeviceHeartbeat.TTL()).Build(),
		r.data.rueidis.B().Zadd().Key(constant.DeviceHeartbeatSorted.Key()).ScoreMember().ScoreMember(float64(now.Timestamp()), sn).Build(),
	)
	for _, result := range results {
		if result.Error() != nil {
			return fmt.Errorf("failed to record heartbeat: %w", result.Error())
		}
	}
	return nil
}

// IsDeviceOnline 检查设备是否在线（通过有序集合检查，更高效）
func (r *DeviceHeartbeatRepo) IsDeviceOnline(ctx context.Context, sn string) (bool, error) {
	// 检查设备是否在有序集合中（比检查单独的key更高效）
	result := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zscore().Key(constant.DeviceHeartbeatSorted.Key()).Member(sn).Build())
	if result.Error() != nil {
		// 如果是not found错误，说明设备不在线
		if result.Error() == rueidis.Nil {
			return false, nil
		}
		return false, fmt.Errorf("failed to check device online status: %w", result.Error())
	}
	// 能获取到score说明设备在线
	return true, nil
}

// IsDeviceOnlineBatch 批量检查设备是否在线
func (r *DeviceHeartbeatRepo) IsDeviceOnlineBatch(ctx context.Context, sns []string) (map[string]bool, error) {
	if len(sns) == 0 {
		return make(map[string]bool), nil
	}
	commands := make([]rueidis.Completed, 0, len(sns))
	for _, sn := range sns {
		commands = append(commands, r.data.rueidis.B().Zscore().Key(constant.DeviceHeartbeatSorted.Key()).Member(sn).Build())
	}
	results := r.data.rueidis.DoMulti(ctx, commands...)
	statusMap := make(map[string]bool, len(sns))
	for i, result := range results {
		sn := sns[i]
		if result.Error() != nil {
			// 如果是not found错误，说明设备不在线
			statusMap[sn] = false
		} else {
			// 能获取到score说明设备在线
			statusMap[sn] = true
		}
	}
	return statusMap, nil
}

// GetAllOnlineDevices 获取所有在线设备
func (r *DeviceHeartbeatRepo) GetAllOnlineDevices(ctx context.Context) ([]string, error) {
	onlineSortedKey := constant.DeviceHeartbeatSorted.Key()
	// 使用 ZRANGE 获取所有在线设备
	result := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zrange().Key(onlineSortedKey).Min("-inf").Max("+inf").Build())
	if result.Error() != nil {
		return nil, fmt.Errorf("failed to get all online devices: %w", result.Error())
	}
	devices, err := result.AsStrSlice()
	if err != nil {
		return nil, fmt.Errorf("failed to parse devices result: %w", err)
	}
	return devices, nil
}

// GetOnlineDevices 获取在线设备列表（按心跳时间倒序）
func (r *DeviceHeartbeatRepo) GetOnlineDevicesByPage(ctx context.Context, page, pageSize int32) ([]string, int64, error) {
	onlineSortedKey := constant.DeviceHeartbeatSorted.Key()
	// 先获取总数
	totalResult := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zcard().Key(onlineSortedKey).Build())
	if totalResult.Error() != nil {
		return nil, 0, fmt.Errorf("failed to get total device count: %w", totalResult.Error())
	}
	total, err := totalResult.AsInt64()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse total count: %w", err)
	}

	if total == 0 {
		return []string{}, 0, nil
	}
	// 计算分页偏移量
	start := int64(page * pageSize)
	stop := int64((page+1)*pageSize) - 1
	if start >= total {
		return []string{}, total, nil
	}
	if stop >= total {
		stop = total - 1
	}
	// 使用 ZREVRANGE 按时间戳倒序获取设备（最新心跳的设备在前）
	results := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zrevrange().Key(onlineSortedKey).Start(start).Stop(stop).Build())
	if results.Error() != nil {
		return nil, 0, fmt.Errorf("failed to get online devices: %w", results.Error())
	}
	devices, err := results.AsStrSlice()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to parse devices result: %w", err)
	}

	return devices, total, nil
}

// GetOnlineDeviceCount 获取在线设备数量
func (r *DeviceHeartbeatRepo) GetOnlineDeviceCount(ctx context.Context) (int64, error) {
	onlineSortedKey := constant.DeviceHeartbeatSorted.Key()
	// 使用 ZCARD 获取有序集合中元素数量
	result := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zcard().Key(onlineSortedKey).Build())
	if result.Error() != nil {
		return 0, fmt.Errorf("failed to get online device count: %w", result.Error())
	}
	count, err := result.AsInt64()
	if err != nil {
		return 0, fmt.Errorf("failed to parse count result: %w", err)
	}
	return count, nil
}

// CleanExpiredHeartbeats 清理过期的心跳记录
func (r *DeviceHeartbeatRepo) CleanExpiredHeartbeats(ctx context.Context) error {
	// 获取2分钟前的时间戳
	expiredTime := carbon.Now().SubMinutes(2).Timestamp()
	onlineSortedKey := constant.DeviceHeartbeatSorted.Key()
	// 直接删除有序集合中的过期设备
	// 单独的心跳key会通过TTL自动过期（1分钟），无需手动删除
	result := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zremrangebyscore().Key(onlineSortedKey).Min("-inf").Max(strconv.FormatInt(expiredTime, 10)).Build())
	if result.Error() != nil {
		return fmt.Errorf("failed to clean expired heartbeats: %w", result.Error())
	}
	return nil
}

// GetDeviceLastHeartbeatTime 获取设备最后心跳时间
func (r *DeviceHeartbeatRepo) GetDeviceLastHeartbeatTime(ctx context.Context, sn string) (int64, error) {
	result := r.data.rueidis.Do(ctx, r.data.rueidis.B().Zscore().Key(constant.DeviceHeartbeatSorted.Key()).Member(sn).Build())
	if result.Error() != nil {
		if result.Error() == rueidis.Nil {
			return 0, nil // 设备不在线，返回0
		}
		return 0, fmt.Errorf("failed to get device last heartbeat time: %w", result.Error())
	}

	timestamp, err := result.AsFloat64()
	if err != nil {
		return 0, fmt.Errorf("failed to parse timestamp: %w", err)
	}

	return int64(timestamp), nil
}

// GetOnlineDevicesByTimeRange 获取指定时间范围内有心跳的设备
func (r *DeviceHeartbeatRepo) GetOnlineDevicesByTimeRange(ctx context.Context, startTime, endTime int64) ([]string, error) {
	onlineSortedKey := constant.DeviceHeartbeatSorted.Key()
	result := r.data.rueidis.Do(ctx,
		r.data.rueidis.B().Zrangebyscore().Key(onlineSortedKey).Min(strconv.FormatInt(startTime, 10)).Max(strconv.FormatInt(endTime, 10)).Build())
	if result.Error() != nil {
		return nil, fmt.Errorf("failed to get devices by time range: %w", result.Error())
	}
	devices, err := result.AsStrSlice()
	if err != nil {
		return nil, fmt.Errorf("failed to parse devices result: %w", err)
	}
	return devices, nil
}
