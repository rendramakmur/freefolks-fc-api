package postgresql

import (
	"freefolks-fc/domain"

	"gorm.io/gorm"
)

type PlayerInformationRepository interface {
	FindAll() ([]domain.PlayerInformation, error)
}

type playerInformationRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *playerInformationRepository {
	return &playerInformationRepository{db: db}
}

func (pi *playerInformationRepository) FindAll() ([]domain.PlayerInformation, error) {
	playerInformation := []domain.PlayerInformation{}
	err := pi.db.Find(&playerInformation).Error

	return playerInformation, err
}
