package battle

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// 新增AI接口
type AIController interface {
	ChooseSkill(c *Character) Skill
	ChooseTarget(c *Character, bm *BattleManager) *Character
}

// 基础AI结构体
type BaseAI struct {
	Personality string // 性格类型：aggressive/defensive/supportive
}

// 进攻型AI
type AggressiveAI struct {
	BaseAI
}

func (ai AggressiveAI) ChooseSkill(c *Character) Skill {
	// 优先使用高伤害技能
	var maxPower int
	var selected Skill
	for _, skill := range c.Skills {
		if dmgSkill, ok := skill.(DamageSkill); ok {
			if dmgSkill.Power == 0 {

			}
			if dmgSkill.Power > maxPower {
				maxPower = dmgSkill.Power
				selected = skill
			}
		}
	}
	if selected != nil {
		return selected
	}
	// 没有伤害技能则随机选择
	return c.Skills[rand.Intn(len(c.Skills))]
}

func (ai AggressiveAI) ChooseTarget(c *Character, bm *BattleManager) *Character {
	// 优先攻击血量最低的敌人
	var minHP = math.MaxInt32
	var target *Character
	for _, enemy := range bm.Characters {
		if enemy != c && enemy.HP > 0 {
			if enemy.HP < minHP {
				minHP = enemy.HP
				target = enemy
			}
		}
	}
	return target
}

// 防御型AI
type DefensiveAI struct {
	BaseAI
}

func (ai DefensiveAI) ChooseSkill(c *Character) Skill {
	// 血量低于30%时尝试使用治疗技能
	if float32(c.HP)/float32(c.MaxHP) < 0.3 {
		for _, skill := range c.Skills {
			if _, ok := skill.(HealSkill); ok {
				return skill
			}
		}
	}

	// 否则使用防御技能或普通攻击
	for _, skill := range c.Skills {
		if _, ok := skill.(ShieldSkill); ok {
			return skill
		}
	}
	return c.Skills[rand.Intn(len(c.Skills))]
}

func (ai DefensiveAI) ChooseTarget(c *Character, bm *BattleManager) *Character {
	// 优先攻击对自己威胁最大的敌人（攻击力最高）
	var maxAttack int
	var target *Character
	for _, enemy := range bm.Characters {
		if enemy != c && enemy.HP > 0 {
			if enemy.Attack > maxAttack {
				maxAttack = enemy.Attack
				target = enemy
			}
		}
	}
	return target
}

// 新增治疗技能
type HealSkill struct {
	Power int
	name  string
}

func (s HealSkill) Execute(caster, target *Character) string {
	healAmount := s.Power
	if target.HP+healAmount > target.MaxHP {
		healAmount = target.MaxHP - target.HP
	}
	target.HP += healAmount
	return fmt.Sprintf("%s 使用 %s 恢复了 %d 生命值", caster.Name, s.name, healAmount)
}

func (s HealSkill) Name() string {
	return s.name
}

// 新增护盾技能
type ShieldSkill struct {
	ShieldValue int
	Duration    int
	name        string
}

func (s ShieldSkill) Execute(caster, target *Character) string {
	// 实现护盾逻辑（需要扩展角色结构体）
	return fmt.Sprintf("%s 获得 %d 点护盾", target.Name, s.ShieldValue)
}

func (s ShieldSkill) Name() string {
	return s.name
}

// 修改角色结构体
type Character struct {
	Name         string
	HP           int
	MaxHP        int
	Attack       int
	Defense      int
	Speed        int
	Status       StatusEffect
	StatusRounds int
	Skills       []Skill
	AI           AIController // 新增AI控制器
	Mux          sync.Mutex
}

// 修改战斗管理器中的行动逻辑
func (bm *BattleManager) StartBattle() {
	// ...（同之前排序逻辑）...
	stedChars := make([]*Character, len(bm.Characters))
	copy(stedChars, bm.Characters)

	for i := 0; i < len(stedChars)-1; i++ {
		for j := 0; j < len(stedChars)-i-1; j++ {
			if stedChars[j].Speed < stedChars[j+1].Speed {
				stedChars[j], stedChars[j+1] = stedChars[j+1], stedChars[j]
			}
		}

		round := 1
		for {
			// ...（同之前状态处理）...

			bm.EventChan <- fmt.Sprintf("第 %d 回合", round)
			bm.ProcessStatusEffects()
			// 修改后的行动逻辑
			for _, c := range stedChars {
				if c.HP <= 0 {
					continue
				}

				target := c.AI.ChooseTarget(c, bm)
				if target == nil {
					bm.EventChan <- "战斗结束!"
					return
				}

				skill := c.AI.ChooseSkill(c)
				result := skill.Execute(c, target)
				bm.EventChan <- result

				// ...（同之前逻辑）...
			}

			// ...（同之前检查逻辑）...
		}
	}
}

// 创建更智能的角色
func createSmartWarrior() *Character {
	return &Character{
		Name:    "精英战士",
		MaxHP:   120,
		HP:      120,
		Attack:  18,
		Defense: 12,
		Speed:   6,
		AI:      AggressiveAI{},
		Skills: []Skill{
			DamageSkill{Power: 25, name: "狂暴斩击"},
			StatusSkill{
				name:        "破甲攻击",
				Effect:      DefenseDown,
				Rounds:      2,
				SuccessRate: 0.8,
			},
			ShieldSkill{
				name:        "钢铁防御",
				ShieldValue: 30,
				Duration:    2,
			},
		},
	}
}

func createSupportMage() *Character {
	return &Character{
		Name:    "支援法师",
		MaxHP:   90,
		HP:      90,
		Attack:  15,
		Defense: 8,
		Speed:   7,
		AI:      DefensiveAI{},
		Skills: []Skill{
			HealSkill{Power: 40, name: "治疗术"},
			DamageSkill{Power: 20, name: "奥术飞弹"},
			StatusSkill{
				name:        "群体护盾",
				Effect:      Shielded,
				Rounds:      2,
				SuccessRate: 1.0,
			},
		},
	}
}

// 扩展状态效果
const (
	// ...原有状态...
	Normal StatusEffect = iota
	Poisoned
	Burning
	DefenseDown
	Shielded
)

// 在状态处理中添加新效果
func (bm *BattleManager) ProcessStatusEffects() {
	for _, c := range bm.Characters {
		c.Mux.Lock()
		if c.Status != Normal && c.StatusRounds > 0 {
			switch c.Status {
			case DefenseDown:
				originalDefense := c.Defense
				c.Defense = int(float32(c.Defense) * 0.7)
				bm.EventChan <- fmt.Sprintf("%s 防御力下降至 %d", c.Name, c.Defense)
				defer func() {
					c.Defense = originalDefense
				}()
			case Shielded:
				// 护盾实现逻辑
			}
			c.StatusRounds--
		}
		c.Mux.Unlock()
	}
}
