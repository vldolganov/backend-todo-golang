package models

type Todo struct {
	ID         uint   `json:"id" gorm:"autoIncrement" gorm:"primaryKey: type:uuid"`
	TaskText   string `json:"taskText"`
	IsComplete bool   `json:"isComplete"`
}
