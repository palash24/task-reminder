package internal

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/palash24/task-reminder/config"
)

func NewDb(cfg config.Config) *pg.DB {
	DB := pg.Connect(&pg.Options{
		User:     cfg.DbUser,
		Password: cfg.DBPassword,
		Database: cfg.DBName,
		Addr:     fmt.Sprintf("%s:%s", cfg.DBAddr, cfg.DBPort),
	})
	if DB == nil {
		log.Fatal("Failed to connect to the database")
	}
	return DB
}
