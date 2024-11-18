package repository

import (
	"github.com/bndrmrtn/my-cloud/database/repository/paginator"
	"gorm.io/gorm"
)

func PaginateList[T any](db *gorm.DB, cursor string) (*paginator.Pagination[T], error) {
	var t T

	data, err := paginator.Paginate[T](db.Model(&t), &paginator.Config{
		Cursor:     cursor,
		Order:      "desc",
		PointsNext: true,
		Limit:      20,
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
