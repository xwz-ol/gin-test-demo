package model

type User struct {
	Id	 string  `xorm:"id PK notnull" json:"id"`
	Name string	`xorm:"name" json:"name"`
}