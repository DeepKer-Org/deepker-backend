package models

type Test struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
}

func (t *Test) GetID() uint {
	return t.ID
}

func (t *Test) SetID(id uint) {
	t.ID = id
}
