package battle

// import (
// 	"fmt"
// 	"time"
// )

// type Character struct {
// 	ID       int
// 	Name     string
// 	Hp       int
// 	Mp       int
// 	Skills   []Skill
// 	Statuses []StatusEffect
// }

// type Skill struct {
// 	ID       int
// 	Name     string
// 	Damage   int
// 	MpCost   int
// 	Cooldown time.Duration
// 	LastUsed time.Time
// }

// // 状态效果（如眩晕、中毒等）
// type StatusEffect struct {
// 	Name       string
// 	Duration   time.Duration
// 	StartTime  time.Time
// 	EffectFunc func(*Character) // 状态效果的具体逻辑
// }

// // 战斗事件
// type BattleEvent struct {
// 	Attacker *Character
// 	Target   *Character
// 	Skill    *Skill
// 	Time     time.Time
// }

// // 战斗逻辑
// func (c *Character) UseSkill(skill *Skill, target *Character) bool {
// 	// 检查技能是否处于冷却中
// 	if time.Since(skill.LastUsed) >= skill.Cooldown {
// 		fmt.Printf("%s 的技能  %s 还在冷却中! \n", c.Name, skill.Name)
// 		return false
// 	}

// 	// 检查MP 是否足够
// 	if c.Mp < skill.MpCost {
// 		fmt.Printf("%s MP不足，无法释放 %s! \n", c.Name, skill.Name)
// 		return false
// 	}

// 	// 消耗MP
// 	c.Mp -= skill.MpCost

// 	// 更新技能的最后一次使用时间
// 	skill.LastUsed = time.Now()

// 	// 对目标造成伤害
// 	target.Hp -= skill.Damage

// 	// if !skill.Cooldown {
// 	// dewlay := time.Duration(skill.Cooldown)
// 	// time.Sleep(delay)
// 	// time.ANSIC()
// 	// }
// 	return false
// }
