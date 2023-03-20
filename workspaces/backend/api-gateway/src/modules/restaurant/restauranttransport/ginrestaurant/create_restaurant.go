package ginrestaurant

import (
	"api-gateway/src/common"
	"api-gateway/src/component"
	"api-gateway/src/modules/restaurant/restaurantbiz"
	"api-gateway/src/modules/restaurant/restaurantmodel"
	"api-gateway/src/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
