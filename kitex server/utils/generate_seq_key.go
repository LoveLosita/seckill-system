package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"kitex-server/seckill/kitex_gen/seckill"
	"kitex-server/seckill/seckill_resp"
	"time"
)

type OrderIDGenerator struct {
	rdb        *redis.Client
	seqKey     string
	timeFormat string
}

// NewOrderIDGenerator 创建一个新的 OrderIDGenerator
func NewOrderIDGenerator(rdb *redis.Client, seqKey string) *OrderIDGenerator {
	return &OrderIDGenerator{
		rdb:        rdb,
		seqKey:     seqKey,
		timeFormat: "20060102", // 按天重置可选：年月日格式，或者使用 "20060102150405" 到秒不重置
	}
}

func (g *OrderIDGenerator) GenerateID(ctx context.Context) (string, seckill.Status) {
	// 1. 获取当前时间前缀
	prefix := time.Now().Format(g.timeFormat)
	// 2. Redis 自增
	seq, err := g.rdb.Incr(ctx, g.seqKey+":"+prefix).Result()
	if err != nil {
		return "", seckill_resp.InternalErr(err)
	}
	// 3. 可选：设置过期时间，让序列按天自动归零
	if seq == 1 {
		// 第一次生成就设置 2 天后过期
		_ = g.rdb.Expire(ctx, g.seqKey+":"+prefix, 48*time.Hour).Err()
	}
	// 4. 格式化序列号为 6 位，不足左补 0
	return fmt.Sprintf("%s%06d", prefix, seq), seckill.Status{}
}
