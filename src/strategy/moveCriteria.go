package strategy

type MoveCriteria interface {
	Move() error
	Spin(instruction rune) MoveCriteria
	GetDirection() string
}
