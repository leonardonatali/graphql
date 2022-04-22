package persistence

import (
	"context"

	"github.com/leonardonatali/graphql/graph/model"
)

type CourseRepository interface {
	GetAllCourses() ([]*model.Course, error)
	CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error)
	GetCoursesByCategoryID(ctx context.Context, categoryID string) ([]*model.Course, error)
}
