package db

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/alvarotor/user-go/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB_PG(cfg *model.Config, l *slog.Logger) *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.POSTGRES_USER, cfg.POSTGRES_PASSWORD, cfg.POSTGRES_DB)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		l.Error("failed to connect database PostGres")
		log.Fatal("failed to connect database PostGres")
	}

	l.Info("DB PostGres Connection established...")

	if err = db.AutoMigrate(&model.User{}); err == nil {
		log.Fatal("failed to AutoMigrate database PostGres")
		// &&
		// db.Migrator().HasTable(&model.User{}) {
		// if err := db.First(&model.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// 	// db.Create(&model.User{
		// 	// 	Email: "goodbytes23@gmail.com",
		// 	// 	Name:  "Alvaro",
		// 	// })
		// }
	}

	return db
}
