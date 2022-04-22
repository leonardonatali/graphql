package persistence

import (
	"context"

	"github.com/leonardonatali/graphql/graph/model"
	"github.com/leonardonatali/graphql/pkg/generators"
	"github.com/pkg/errors"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrCourseNotFound   = errors.New("course not found")
)

type MemoryPersistence struct {
	IDGenerator generators.IDGenerator
	Categories  []*model.Category
	Courses     []*model.Course
	Chapters    []*model.Chapter
}

func NewMemoryPersistence(generator generators.IDGenerator) *MemoryPersistence {
	return &MemoryPersistence{
		IDGenerator: generator,
		Categories:  []*model.Category{},
		Courses:     []*model.Course{},
		Chapters:    []*model.Chapter{},
	}
}

func (m *MemoryPersistence) GetAllCategories() ([]*model.Category, error) {
	return m.Categories, nil
}

func (m *MemoryPersistence) GetCategoryByID(id string) (*model.Category, error) {
	for _, c := range m.Categories {
		if c.ID == id {
			return c, nil
		}
	}

	return nil, ErrCategoryNotFound
}

func (m *MemoryPersistence) GetAllChapters() ([]*model.Chapter, error) {
	return m.Chapters, nil
}

func (m *MemoryPersistence) GetAllCourses() ([]*model.Course, error) {
	return m.Courses, nil
}

func (m *MemoryPersistence) GetCourseByID(id string) (*model.Course, error) {
	for _, c := range m.Courses {
		if c.ID == id {
			return c, nil
		}
	}

	return nil, ErrCourseNotFound
}

func (m *MemoryPersistence) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := &model.Category{
		ID:          m.IDGenerator.NewString(),
		Name:        input.Name,
		Description: input.Description,
	}

	m.Categories = append(m.Categories, category)
	return category, nil
}

func (m *MemoryPersistence) GetCoursesByCategoryID(ctx context.Context, categoryID string) ([]*model.Course, error) {
	courses := []*model.Course{}

	for _, course := range courses {
		if course.Category.ID == categoryID {
			courses = append(courses, course)
		}
	}

	return courses, nil
}

func (m *MemoryPersistence) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	category, err := m.GetCategoryByID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	course := &model.Course{
		ID:          m.IDGenerator.NewString(),
		Name:        input.Name,
		Description: input.Description,
		Category:    category,
	}

	m.Courses = append(m.Courses, course)
	return course, nil
}

func (m *MemoryPersistence) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	course, err := m.GetCourseByID(input.CourseID)
	if err != nil {
		return nil, err
	}

	category, err := m.GetCategoryByID(input.CategoryID)
	if err != nil {
		return nil, err
	}

	chapter := &model.Chapter{
		ID:       m.IDGenerator.NewString(),
		Name:     input.Name,
		Category: category,
		Course:   course,
	}

	m.Chapters = append(m.Chapters, chapter)
	return chapter, nil
}

func (m *MemoryPersistence) GetChaptersByCourseID(ctx context.Context, courseID string) ([]*model.Chapter, error) {
	chapters := []*model.Chapter{}

	for _, chapter := range chapters {
		if chapter.Course.ID == courseID {
			chapters = append(chapters, chapter)
		}
	}

	return chapters, nil
}
