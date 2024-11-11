package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Car struct {
	Pemilik  string `gorm:"varchar(255)" json:"pemilik"`
	Merk     string `gorm:"varchar(255)" json:"merk"`
	Harga    int    `gorm:"int" json:"harga"`
	TypeCars string `gorm:"varchar(255)" json:"typecars" valid:"required~Type is required"`
	ID       uint   `gorm:"primaryKey"`
	UserID   int    `gorm:"int"`
	User     *User
}

func (c *Car) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	err = nil
	return
}
