package data

import (
	"fmt"
	"github.com/guneyin/sbda-product-category-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const dbName = "product"

func newPGClient(cfg *config.Config) (*gorm.DB, error) {
	if cfg == nil {
		log.Fatal("invalid config 1")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPwd)

	err := createDBIfNotExist(dsn, dbName)
	if err != nil {
		return nil, err
	}

	dsn = fmt.Sprintf("dbname=%s %s", dbName, dsn)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
}

func createDBIfNotExist(dsn, dbName string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	var cnt int32
	row := db.Debug().Raw("SELECT count(datname) cnt FROM pg_database WHERE datname = ?", dbName).Row()
	row.Scan(&cnt)

	if cnt == 0 {
		err = db.Debug().Exec(fmt.Sprintf("CREATE database %s", dbName)).Error
	}

	return nil
}
