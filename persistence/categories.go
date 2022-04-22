package persistence

import (
	"context"

	"github.com/leonardonatali/graphql/graph/model"
)

type CategoryRepository interface {
	GetAllCategories() ([]*model.Category, error)
	CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error)
}
