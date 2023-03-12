package pkg

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	md "github.com/LeilaBeken/golang_ass_3/models"
)

func GetDB() (*gorm.DB, error) {
    dsn := "user=postgres password=belelik04 dbname=golang3 host=localhost port=5432 sslmode=disable TimeZone=UTC"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {panic("failed to connect database")}

	db.AutoMigrate(&md.Book{})

    return db, nil
}

type book struct{
	*md.Book
}

func (b *book) GetByID(id uint) error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.First(b, id)
    return result.Error
}

func (b *book) Create() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Create(b)
    return result.Error
}

func (b *book) Update() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Save(b)
    return result.Error
}

func (b *book) Delete() error {
    db, err := GetDB()
	if(err != nil){panic(err)}
    result := db.Delete(b)
    return result.Error
}