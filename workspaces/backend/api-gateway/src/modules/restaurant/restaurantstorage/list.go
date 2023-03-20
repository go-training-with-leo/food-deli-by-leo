package restaurantstorage

import (
	"api-gateway/src/common"
	"api-gateway/src/modules/restaurant/restaurantmodel"
	"context"
	"strings"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if len(strings.TrimSpace(v.Name)) > 0 {
			db = db.Where("name = ?", v.Name)
		}
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
