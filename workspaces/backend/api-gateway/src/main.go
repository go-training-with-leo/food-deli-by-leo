package main

import (
	"api-gateway/src/component"
	"api-gateway/src/database"
	"api-gateway/src/modules/restaurant/restauranttransport/ginrestaurant"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func main() {
	db := database.CreateInstance()

	if err := runService(db); err != nil {
		log.Fatal("can not start the server.\n", err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	// CRUD

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var restaurant Restaurant

			if err := db.First(&restaurant, "id = ?", id).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "cannot get restaurant",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "restaurant retrieved",
				"data":    restaurant,
			})
		})

		restaurants.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			var data RestaurantUpdate

			if err := ctx.ShouldBindJSON(&data); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid request body",
				})
				return
			}

			fmt.Println(&data)

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "cannot update restaurant",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "restaurant updated",
			})
		})

		restaurants.DELETE("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			if err := db.Where("id = ?", id).Delete(&Restaurant{}).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "cannot delete restaurant",
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "restaurant deleted",
			})
		})
	}

	return r.Run()
}
