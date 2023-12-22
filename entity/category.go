package entity

import (
	pb "github.com/guneyin/sbda-sdk/pb"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"name"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Tags        string `gorm:"tags"`
	Image       string `gorm:"image"`
}

func ToCategoryEntity(in *pb.CreateCategoryRequest) *Category {
	return &Category{
		Name:        in.Name,
		Title:       in.Title,
		Description: in.Description,
		Tags:        in.Tags,
		Image:       in.Image,
	}
}

func FromCategoryEntity(in *Category) *pb.CreateCategoryResponse {
	return &pb.CreateCategoryResponse{
		ID:          int32(in.ID),
		Name:        in.Name,
		Title:       in.Title,
		Description: in.Description,
		Tags:        in.Tags,
		Image:       in.Image,
	}
}
