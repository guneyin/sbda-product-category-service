package data

import (
	"github.com/guneyin/sbda-product-category-service/config"
	"github.com/guneyin/sbda-product-category-service/entity"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	db *gorm.DB
}

func NewDB(cfg *config.Config) (*DB, error) {
	if cfg == nil {
		log.Fatal("invalid config 2")
	}

	db, err := newPGClient(cfg)
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(&entity.Category{})

	return &DB{db: db}, nil
}

func (d *DB) Save(data any) (any, error) {
	res := d.db.Save(&data)
	if res.Error != nil {
		return nil, res.Error
	}

	//m := new(gorm.Model)
	//_ = res.Row().Scan(&m)

	return &data, nil
}
