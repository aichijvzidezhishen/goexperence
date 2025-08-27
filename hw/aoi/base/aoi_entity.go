package base

import (
	"sync"
)

type Position struct {
	X, Y float32
}

// 实体接口
type Entity interface {
	GetID() int
	GetPosition() Position
	OnEnterAOI(other Entity) // 进入AOI回调
	OnLeaveAOI(other Entity) // 离开AOI回调
}

// 网格定义
type Grid struct {
	GridID   int
	Entities map[int]Entity // 当前网格内的实体
	mu       sync.RWMutex
}

func NewGrid(id int) *Grid {
	return &Grid{
		GridID:   id,
		Entities: make(map[int]Entity),
	}
}

// AOI 管理器
type AOIManager struct {
	GridWidth  float32        // 网格宽度
	GridHeight float32        // 网格高度
	Grids      map[int]*Grid  // 所有网格
	Entities   map[int]Entity // 所有实体
	mu         sync.RWMutex   // 读写锁
}

// 创建一个新的AOIManager实例
func NewAOIManager(gridWidth, gridHeight float32) *AOIManager {
	// 返回一个指向AOIManager实例的指针
	return &AOIManager{
		// 设置网格宽度
		GridWidth: gridWidth,
		// 设置网格高度
		GridHeight: gridHeight,
		// 创建一个空的网格映射
		Grids: make(map[int]*Grid),
		// 创建一个空的实体映射
		Entities: make(map[int]Entity),
	}
}
