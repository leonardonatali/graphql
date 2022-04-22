package graph

import (
	"github.com/leonardonatali/graphql/persistence"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Categories persistence.CategoryRepository
	Chapters   persistence.ChapterRepository
	Courses    persistence.CourseRepository
}

func NewResolver(categories persistence.CategoryRepository, chapters persistence.ChapterRepository, courses persistence.CourseRepository) Resolver {
	return Resolver{
		Categories: categories,
		Chapters:   chapters,
		Courses:    courses,
	}
}
