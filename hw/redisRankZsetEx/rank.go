package redisrankzsetex

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const gamePowerRankKey = "game:king:power_rank"

type GamePowerRank struct {
	rdb     *redis.Client
	ctx     context.Context
	rankKey string
}

func NewGamePowerRank(redisArr string, redispwd string, db int) (*GamePowerRank, error) {
	// init redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisArr,
		Password: redispwd,
		DB:       db,
	})

	//verify redis connect
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("redis connect faild : %v", err)
	}
	return &GamePowerRank{
		rdb:     rdb,
		ctx:     ctx,
		rankKey: gamePowerRankKey,
	}, nil
}

func (g *GamePowerRank) InitPlayers(playerPowers map[string]int) (int64, error) {
	var (
		zSet     []*redis.Z
		addCount int64
		err      error
	)
	for uid, power := range playerPowers {
		zSet = append(zSet, &redis.Z{
			Member: uid,
			Score:  float64(power),
		})
	}

	//
	addCount, err = g.rdb.ZAdd(g.ctx, g.rankKey, zSet...).Result()
	if err != nil {
		return 0, fmt.Errorf("init players faild: %w", err)
	}
	return addCount, nil
}

// parm : n(player count)
func (g *GamePowerRank) GetTopN(n int64) ([]redis.Z, error) {
	if n <= 0 {
		return nil, fmt.Errorf("find count must greater than 0")
	}

	topN, err := g.rdb.ZRevRangeWithScores(g.ctx, g.rankKey, 0, n-1).Result()
	if err != nil {
		return nil, fmt.Errorf("get top %d faild ", n)
	}
	return topN, nil
}

func (g *GamePowerRank) GetPlayerInfo(playerID string) (int, int, error) {
	//get power
	power, err := g.rdb.ZScore(g.ctx, g.rankKey, playerID).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("get player %v faild", playerID)
	}
	powerRes := int(power)

	//get rank  (zrevrank 返回排名从0 开始)
	rank, err := g.rdb.ZRevRank(g.ctx, g.rankKey, playerID).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("get player %v rank faild", playerID)
	}
	rankRes := int(rank + 1)

	return powerRes, rankRes, nil
}

func (g *GamePowerRank) UpdatePlayerRank(playerID string, newPower int) error {
	_, err := g.rdb.ZAdd(g.ctx, g.rankKey, &redis.Z{
		Member: playerID,
		Score:  float64(newPower),
	}).Result()
	if err != nil {
		return fmt.Errorf("update player %v rank %d faild", playerID, newPower)
	}
	return nil
}

func (g *GamePowerRank) UpdatePowerIncr(playerID string, incr int) (int, error) {
	newPowerFloat, err := g.rdb.ZIncrBy(g.ctx, g.rankKey, float64(incr), playerID).Result()
	if err != nil {
		return 0, fmt.Errorf("增量更新%s战力失败: %w", playerID, err)
	}
	newPower := int(newPowerFloat)

	return newPower, nil
}

func (g *GamePowerRank) DelPlayer(playerIDs ...string) (int64, error) {
	var members []interface{}
	for _, pid := range playerIDs {
		members = append(members, pid)
	}

	delCount, err := g.rdb.ZRem(g.ctx, g.rankKey, members...).Result()
	if err != nil {
		return 0, fmt.Errorf("del player faild: %v", err)
	}

	return delCount, nil
}

func (g *GamePowerRank) GetPlayersByPowerRange(minPower, maxPower string) ([]redis.Z, error) {
	players, err := g.rdb.ZRevRangeByScoreWithScores(g.ctx, g.rankKey, &redis.ZRangeBy{
		Min: minPower,
		Max: maxPower,
	}).Result()
	if err != nil {
		return nil, fmt.Errorf("查询战力区间[%s, %s]玩家失败: %w", minPower, maxPower, err)
	}

	return players, nil
}

// GetRankStats 获取排行榜统计信息
// 返回：总玩家数 / 战力≥指定值的玩家数 / 错误
func (g *GamePowerRank) GetRankStats(minPower string) (int64, int64, error) {
	// 总玩家数
	total, err := g.rdb.ZCard(g.ctx, g.rankKey).Result()
	if err != nil {
		return 0, 0, fmt.Errorf("统计总玩家数失败: %w", err)
	}

	// 战力≥minPower的玩家数
	count, err := g.rdb.ZCount(g.ctx, g.rankKey, minPower, "+inf").Result()
	if err != nil {
		return 0, 0, fmt.Errorf("统计战力≥%s玩家数失败: %w", minPower, err)
	}

	return total, count, nil
}

// PrintTopN 格式化打印TOP N玩家信息（辅助方法，提升可读性）
func PrintTopN(topN []redis.Z) {
	fmt.Println("\n🏆 战力TOP排名：")
	for idx, z := range topN {
		playerID := z.Member.(string)
		power := int(z.Score)
		fmt.Printf("第 %d 名：%s，战力：%d\n", idx+1, playerID, power)
	}
}

// PrintPlayerPowerRange 格式化打印战力区间玩家信息（辅助方法）
func PrintPlayerPowerRange(players []redis.Z, minPower string) {
	fmt.Printf("\n🎯 战力 ≥ %s 的玩家：\n", minPower)
	for _, z := range players {
		playerID := z.Member.(string)
		power := int(z.Score)
		fmt.Printf("玩家：%s，战力：%d\n", playerID, power)
	}
}
