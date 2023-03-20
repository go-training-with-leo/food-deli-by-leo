package restaurantmodel

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Restaurant struct {
	Id      uuid.UUID `json:"id" gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid();"`
	Name    string    `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Address string    `json:"address" gorm:"column:address;type:varchar(100);not null"`
	// Status    int       `json:"status" gorm:"columnd:status;type:integer;not null"`
	// CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	// UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:address;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id      uuid.UUID `json:"id" gorm:"column:id;default:gen_random_uuid();"`
	Name    string    `json:"name" gorm:"column:name;"`
	Address string    `json:"address" gorm:"column:address;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	return nil
}
