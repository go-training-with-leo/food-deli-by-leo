package restaurantmodel

type Filter struct {
	Name string `json:"name,omitempty" form:"name"`
}
