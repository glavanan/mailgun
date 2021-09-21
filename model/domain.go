package model

//Domain ...
type Domain struct {
	Name    string `json:"name" gorm:"primaryKey"`
	Bounce  int    `json:"bounce"`
	Deliver int    `json:"deliver"`
}
