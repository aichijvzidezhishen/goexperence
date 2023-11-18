package base

// type AoiEntity struct {
// 	Key           int64
// 	X             *AoiNode
// 	Y             *AoiNode
// 	ViewEntity    *sync.HashSet //
// 	ViewEntityBak *sync.HashSet
// 	Move          func() []int64
// 	Leave         func() []int64
// }

// func NewAoiEntity(key int64) *AoiEntity {
// 	return &AoiEntity{
// 		Key:           key,
// 		ViewEntity:    sync.NewHashSet(),
// 		ViewEntityBak: sync.NewHashSet(),
// 		Move: func() []int64 {
// 			return (*(this.ViewEntityBak)).Union((*(this.ViewEntity)).Values()).ToArray()
// 		},
// 		Leave: func() []int64 {
// 			return (*(this.ViewEntityBak)).Except((*(this.ViewEntity)).Values()).ToArray()
// 		},
// 	}
// }

// func (this *AoiEntity) Dispose() {
// 	// (*(this.ViewEntity)).Clear()
// 	// Pool<sync.HashSet>.Return(this.ViewEntity)
// 	// (*(this.ViewEntityBak)).Clear()
// 	// Pool<sync.HashSet>.Return(this.ViewEntityBak)
// }

// func (this *AoiNode) Dispose() {
// }

// const (
// 	PoolSize int = 100
// )

// var pool = make([]*AoiEntity, PoolSize)
// var poolMutex = sync.Mutex{}

// func Rent() *AoiEntity {
// 	poolMutex.Lock()
// 	defer poolMutex.Unlock()

// 	if len(pool) == 0 {
// 		return NewAoiEntity(0)
// 	}

// 	var result *AoiEntity
// 	for _, item := range pool {
// 		if item == nil {
// 			result = item
// 			break
// 		}
// 	}

// 	if result == nil {
// 		result = NewAoiEntity(0)
// 	}

// 	pool = append(pool[:len(pool)-1], pool[len(pool):]...)

// 	return result
// }

// func Return(item *AoiEntity) {
// 	poolMutex.Lock()
// 	defer poolMutex.Unlock()

// 	pool = append(pool[:len(pool)-1], item)
// }

// func (this *AoiEntity) GetMove() []int64 {
// 	return this.Move()
// }

// func (this *AoiEntity) GetLeave() []int64 {
// 	return this.Leave()
// }

type AoiEvent interface {
	OnEnterAoi(node []*AoiNode)
	OnUpdateAoi(node []*AoiNode)
	OnLeaveAoi(node []*AoiNode)
}

type AoiAction interface {
	EnterAoi(node *AoiNode)
	LeaveAoi(node *AoiNode)
	MoveAoi(node *AoiNode)
	FindNeighbors(node *AoiNode, radius float32) map[*AoiNode]struct{}
	Print()
}
