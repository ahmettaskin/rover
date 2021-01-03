package strategy

type MoveStrategy struct {
	moveCriteria MoveCriteria
}

func NewMoveStrategy(criteria MoveCriteria) *MoveStrategy {
	return &MoveStrategy{
		moveCriteria: criteria,
	}
}

func (ms *MoveStrategy) SetCriteria(criteria MoveCriteria) {
	ms.moveCriteria = criteria
}

func (ms *MoveStrategy) GetCriteria() MoveCriteria {
	return ms.moveCriteria
}
