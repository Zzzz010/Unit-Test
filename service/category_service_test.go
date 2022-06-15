package service

import (
	"belajar-go-unit-test/entity"
	"belajar-go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "010",
		Name: "Raung",
	}
	categoryRepository.Mock.On("FindById", "010").Return(category)

	hasil, err := categoryService.Get("010")
	assert.Nil(t, err)
	assert.NotNil(t, hasil)
	assert.Equal(t, category.Id, hasil.Id)
	assert.Equal(t, category.Name, hasil.Name)
}
