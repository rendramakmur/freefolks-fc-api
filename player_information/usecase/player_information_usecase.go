package usecase

import (
	"freefolks-fc/domain"
	"freefolks-fc/player_information/repository/postgresql"
)

type PlayerInformationUsecase interface {
	FindAll() ([]domain.PlayerInformation, error)
}

type playerInformationUsecase struct {
	r postgresql.PlayerInformationRepository
}

func NewPlayerInformationUsecase(r postgresql.PlayerInformationRepository) *playerInformationUsecase {
	return &playerInformationUsecase{r: r}
}

func (pu *playerInformationUsecase) FindAll() ([]domain.PlayerInformation, error) {
	playerInformations, err := pu.r.FindAll()

	return playerInformations, err
}
