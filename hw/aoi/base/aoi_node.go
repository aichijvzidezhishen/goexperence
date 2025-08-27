package base

// 获取实体所在网格ID
func (m *AOIManager) GetGridID(pos Position) int {
	x := int(pos.X / m.GridWidth)
	y := int(pos.Y / m.GridHeight)
	return y*1000 + x // 生成唯一ID（假设地图横向不超过1000格）
}

// 添加实体
func (m *AOIManager) AddEntity(entity Entity) {
	m.mu.Lock()
	defer m.mu.Unlock()

	gridID := m.GetGridID(entity.GetPosition())
	grid, exists := m.Grids[gridID]
	if !exists {
		grid = NewGrid(gridID)
		m.Grids[gridID] = grid
	}

	grid.mu.Lock()
	grid.Entities[entity.GetID()] = entity
	grid.mu.Unlock()

	m.Entities[entity.GetID()] = entity
}

// 移除实体
func (m *AOIManager) RemoveEntity(entityID int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	entity, exists := m.Entities[entityID]
	if !exists {
		return
	}

	gridID := m.GetGridID(entity.GetPosition())
	if grid, ok := m.Grids[gridID]; ok {
		grid.mu.Lock()
		delete(grid.Entities, entityID)
		grid.mu.Unlock()
	}

	delete(m.Entities, entityID)
}

// 实体移动
func (m *AOIManager) MoveEntity(entityID int, newPos Position) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entity, exists := m.Entities[entityID]
	if !exists {
		return
	}

	oldGridID := m.GetGridID(entity.GetPosition())
	newGridID := m.GetGridID(newPos)

	if oldGridID != newGridID {
		// 跨网格移动，触发离开旧网格和进入新网格
		oldGrid := m.Grids[oldGridID]
		newGrid, exists := m.Grids[newGridID]
		if !exists {
			newGrid = NewGrid(newGridID)
			m.Grids[newGridID] = newGrid
		}

		// 从旧网格移除
		oldGrid.mu.Lock()
		delete(oldGrid.Entities, entityID)
		oldGrid.mu.Unlock()

		// 加入新网格
		newGrid.mu.Lock()
		newGrid.Entities[entityID] = entity
		newGrid.mu.Unlock()

		// 触发事件
		m.triggerCrossGridEvents(entity, oldGridID, newGridID)
	}
}

// 触发跨网格事件
func (m *AOIManager) triggerCrossGridEvents(entity Entity, oldGridID, newGridID int) {
	// 获取旧网格周围的实体并通知离开
	oldNeighbors := m.GetNeighborGrids(oldGridID)
	for _, grid := range oldNeighbors {
		grid.mu.RLock()
		for _, e := range grid.Entities {
			if e.GetID() != entity.GetID() {
				e.OnLeaveAOI(entity)
				entity.OnLeaveAOI(e)
			}
		}
		grid.mu.RUnlock()
	}

	// 获取新网格周围的实体并通知进入
	newNeighbors := m.GetNeighborGrids(newGridID)
	for _, grid := range newNeighbors {
		grid.mu.RLock()
		for _, e := range grid.Entities {
			if e.GetID() != entity.GetID() {
				e.OnEnterAOI(entity)
				entity.OnEnterAOI(e)
			}
		}
		grid.mu.RUnlock()
	}
}

// 获取周围九宫格网格
func (m *AOIManager) GetNeighborGrids(gridID int) []*Grid {
	// 实现九宫格逻辑（此处简化）
	// TODO: 根据 gridID 计算周围网格
	return []*Grid{m.Grids[gridID]} // 示例仅返回当前网格
	// return []
}
