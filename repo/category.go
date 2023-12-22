package repo

import (
	"context"
	"github.com/guneyin/sbda-product-category-service/data"
	"github.com/guneyin/sbda-product-category-service/entity"
	pb "github.com/guneyin/sbda-sdk/pb"
)

type CategoryRepo struct {
	db *data.DB
}

func newCategoryRepo(db *data.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (cr *CategoryRepo) Create(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	e := entity.ToCategoryEntity(in)

	d, err := cr.db.Save(e)
	if err != nil {
		return nil, err
	}

	val, ok := d.(*entity.Category)
	if !ok {
		return nil, ErrCreateCategoryFailed
	}

	return entity.FromCategoryEntity(val), nil
}
