package services

import (
	"context"

	"github.com/guisfits/go-grpc/internal/database"
	"github.com/guisfits/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResult := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CreateCategoryResponse{Category: categoryResult}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResult []*pb.Category

	for _, category := range categories {
		categoriesResult = append(categoriesResult, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoriesResult}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	categoryResult := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResult, nil
}
