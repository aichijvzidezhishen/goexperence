package battle

import (
	"fmt"
	"math/rand"
	"time"
)

// 角色状态类型
type StatusEffect int

const (
// Normal StatusEffect = iota
// Poisoned
// Burning
)

// // 角色结构体
// type Character struct {
// 	Name         string
// 	HP           int
// 	MaxHP        int
// 	Attack       int
// 	Defense      int
// 	Speed        int
// 	Status       StatusEffect
// 	StatusRounds int
// 	Skills       []Skill
// 	Mux          sync.Mutex
// }

// 技能接口
type Skill interface {
	Execute(caster, target *Character) string
	Name() string
}

// 直接伤害技能
type DamageSkill struct {
	Power    int
	Cost     int
	Cooldown int
	name     string
}

// 执行伤害技能
func (s DamageSkill) Execute(caster, target *Character) string {
	// 加锁，防止并发访问
	caster.Mux.Lock()
	defer caster.Mux.Unlock()

	// 计算伤害值
	damage := s.Power + caster.Attack - target.Defense
	// 如果伤害值小于0，则设为0
	if damage < 0 {
		damage = 0
	}
	// 对目标造成伤害
	target.HP -= damage
	// 返回攻击者的名字、技能名称和伤害值
	return fmt.Sprintf("%s 使用 %s 造成 %d 点伤害", caster.Name, s.name, damage)
}

func (s DamageSkill) Name() string {
	return s.name
}

// 状态效果技能
type StatusSkill struct {
	Effect      StatusEffect
	Rounds      int
	SuccessRate float32
	name        string
}

func (s StatusSkill) Execute(caster, target *Character) string {
	if rand.Float32() < s.SuccessRate {
		target.Status = s.Effect
		target.StatusRounds = s.Rounds
		return fmt.Sprintf("%s 被施加了 %s", target.Name, s.name)
	}
	return fmt.Sprintf("%s 抵抗了 %s", target.Name, s.name)
}

func (s StatusSkill) Name() string {
	return s.name
}

// 战斗管理器
type BattleManager struct {
	Characters []*Character
	EventChan  chan string
}

func NewBattleManager() *BattleManager {
	return &BattleManager{
		EventChan: make(chan string, 100),
	}
}

// // 处理状态效果
// // ProcessStatusEffects 函数用于处理战斗管理器中的状态效果
// func (bm *BattleManager) ProcessStatusEffects() {
// 	// 遍历战斗管理器中的所有角色
// 	for _, c := range bm.Characters {
// 		// 加锁，防止并发访问
// 		c.Mux.Lock()
// 		// 如果角色的状态不是正常状态，并且状态轮数大于0
// 		if c.Status != Normal && c.StatusRounds > 0 {
// 			// 根据角色的状态，进行不同的处理
// 			switch c.Status {
// 			case Poisoned:
// 				// 如果角色中毒，则造成最大生命值的10%的伤害
// 				damage := c.MaxHP / 10
// 				c.HP -= damage
// 				// 将事件发送到事件通道
// 				bm.EventChan <- fmt.Sprintf("%s 受到中毒伤害 %d", c.Name, damage)
// 			case Burning:
// 				// 如果角色灼烧，则造成最大生命值的12.5%的伤害
// 				damage := c.MaxHP / 8
// 				c.HP -= damage
// 				// 将事件发送到事件通道
// 				bm.EventChan <- fmt.Sprintf("%s 受到灼烧伤害 %d", c.Name, damage)
// 			}
// 			// 状态轮数减1
// 			c.StatusRounds--
// 		}
// 		// 解锁
// 		c.Mux.Unlock()
// 	}
// }

// // 开始战斗
// func (bm *BattleManager) StartBattle() {
// 	// 按速度排序
// 	sortedChars := make([]*Character, len(bm.Characters))
// 	copy(sortedChars, bm.Characters)
// 	for i := 0; i < len(sortedChars)-1; i++ {
// 		for j := 0; j < len(sortedChars)-i-1; j++ {
// 			if sortedChars[j].Speed < sortedChars[j+1].Speed {
// 				sortedChars[j], sortedChars[j+1] = sortedChars[j+1], sortedChars[j]
// 			}
// 		}
// 	}

// 	round := 1
// 	for {
// 		bm.EventChan <- fmt.Sprintf("=== 第 %d 回合 ===", round)
// 		bm.ProcessStatusEffects()

// 		// 执行角色行动
// 		for _, c := range sortedChars {
// 			if c.HP <= 0 {
// 				continue
// 			}

// 			// 简单AI：随机选择技能和目标
// 			if len(c.Skills) > 0 {
// 				target := bm.FindAliveTarget(c)
// 				if target == nil {
// 					bm.EventChan <- "战斗结束!"
// 					return
// 				}

// 				skill := c.Skills[rand.Intn(len(c.Skills))]
// 				result := skill.Execute(c, target)
// 				bm.EventChan <- result

// 				if target.HP <= 0 {
// 					bm.EventChan <- fmt.Sprintf("%s 被击败了!", target.Name)
// 				}
// 			}
// 		}

// 		// 检查存活状态
// 		if aliveCount := bm.CountAlive(); aliveCount < 2 {
// 			bm.EventChan <- "战斗结束!"
// 			return
// 		}

// 		round++
// 		time.Sleep(1 * time.Second) // 回合间隔
// 	}
// }

// 辅助方法
// 在BattleManager中查找一个存活的敌人
func (bm *BattleManager) FindAliveTarget(self *Character) *Character {
	// 遍历所有角色
	for _, c := range bm.Characters {
		// 如果角色不是自己且HP大于0
		if c != self && c.HP > 0 {
			// 返回该角色
			return c
		}
	}
	// 如果没有找到存活的敌人，返回nil
	return nil
}

func (bm *BattleManager) CountAlive() int {
	count := 0
	for _, c := range bm.Characters {
		if c.HP > 0 {
			count++
		}
	}
	return count
}

func main() {
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
