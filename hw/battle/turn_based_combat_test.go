package battle

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBattleManager_StartBattle(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// 创建角色
	warrior := &Character{
		Name:    "战士",
		MaxHP:   100,
		HP:      100,
		Attack:  15,
		Defense: 10,
		Speed:   5,
		Skills: []Skill{
			DamageSkill{Power: 20, name: "重劈"},
			StatusSkill{
				name:        "火焰斩",
				Effect:      Burning,
				Rounds:      3,
				SuccessRate: 0.7,
			},
		},
	}

	mage := &Character{
		Name:    "法师",
		MaxHP:   80,
		HP:      80,
		Attack:  20,
		Defense: 5,
		Speed:   8,
		Skills: []Skill{
			DamageSkill{Power: 25, name: "火球术"},
			StatusSkill{
				name:        "剧毒云雾",
				Effect:      Poisoned,
				Rounds:      3,
				SuccessRate: 0.6,
			},
		},
	}

	// 创建战斗管理器
	bm := NewBattleManager()
	bm.Characters = []*Character{warrior, mage}

	// 启动战斗日志打印
	go func() {
		for event := range bm.EventChan {
			fmt.Println(event)
		}
	}()

	// 开始战斗
	bm.StartBattle()
	close(bm.EventChan)
}
