package database

import (
	"apitest/domain/entity"
	"apitest/domain/repository"
	"apitest/infrastructure/persistence"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	User repository.UserRepository
	db   *gorm.DB
}

func InitConfiguration() (*Repositories, error) {
	DbHost := os.Getenv("DBHOST")
	DbUser := os.Getenv("DBUSER")
	DbPassword := os.Getenv("DBPASSWORD")
	DbPort := os.Getenv("DBPORT")
	DbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Repositories{
		User: persistence.NewUserRepository(db),
		db:   db,
	}, nil
}

func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.User{})
}
