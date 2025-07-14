package employee

type Employee struct {
	ID       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Position string `gorm:"column:position"`
}
