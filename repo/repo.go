package repo

import (
	"github.com/guneyin/sbda-product-category-service/config"
	"github.com/guneyin/sbda-product-category-service/data"
	"log"
)

type Repo struct {
	cfg      *config.Config
	db       *data.DB
	Category *CategoryRepo
}

func NewRepo(cfg *config.Config) (*Repo, error) {
	if cfg == nil {
		log.Fatal("invalid config 3")
	}

	db, err := data.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Repo{
		cfg:      cfg,
		db:       db,
		Category: newCategoryRepo(db),
	}, nil
}
