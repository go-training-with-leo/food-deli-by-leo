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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
