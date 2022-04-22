package persistence

import (
	"context"

	"github.com/leonardonatali/graphql/graph/model"
)

type ChapterRepository interface {
	GetAllChapters() ([]*model.Chapter, error)
	CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error)
	GetChaptersByCourseID(ctx context.Context, courseID string) ([]*model.Chapter, error)
}
