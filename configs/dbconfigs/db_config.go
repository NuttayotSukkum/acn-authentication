package dbconfigs

import (
	"fmt"
	"github.com/NuttayotSukkum/acn/acn-authentication/configs"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func InitDB(cfg *configs.Configs) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		cfg.Database.DatabaseUser,
		cfg.Database.DatabasePass,
		cfg.Database.DatabaseHost,
		cfg.Database.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	} else {
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatalf("database migration failed: %v", err)
		}
	}
	log.Printf("database connection success")
	return db
}
