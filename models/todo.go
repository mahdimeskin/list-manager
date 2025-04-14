package models

type ToDo struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Note string `json:"note"`
	Completed bool `json:"completed"`
}