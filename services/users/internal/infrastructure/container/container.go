package container

import (
	"fmt"
	"gardener/services/users/internal/infrastructure/config"
	"gardener/services/users/internal/models/user"
	"gardener/services/users/internal/models/user/profile"
	"gardener/services/users/internal/services"
	sUser "gardener/services/users/internal/services/user"
	"log"
	"os"
	"time"

	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() (*dig.Container, error) {
	var err error
	container := dig.New()

	err = container.Provide(func() (*config.Config, error) {
		cfg, err := config.New()
		fmt.Println(cfg, err)
		if err != nil {
			return nil, err
		}

		return cfg, nil
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(cfg *config.Config) (*gorm.DB, error) {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  true,          // Disable color
			},
		)
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow", cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger.LogMode(logger.Info)})
		if err != nil {
			return nil, err
		}

		if err := db.AutoMigrate(&user.User{}, &profile.Profile{}); err != nil {
			return nil, err
		}

		return db, nil
	})
	if err != nil {
		return nil, err
	}

	err = container.Provide(func(repo *gorm.DB) services.UserService {
		return sUser.New(repo)
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}
