package strategy

type MoveCriteria interface {
	Move() error
	Spin(command rune) MoveCriteria
	GetDirection() string
}
