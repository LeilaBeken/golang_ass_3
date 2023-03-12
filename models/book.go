package models

type Book struct{
    ID       uint   `json:"id" gorm:"primaryKey"`
	Title        string `json:"title"`
	Description string `json:"description"`
    Price       int    `json:"price"`
}