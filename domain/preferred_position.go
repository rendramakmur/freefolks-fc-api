package domain

import "time"

type PreferredPosition struct {
	Id                  int               `gorm:"column:pp_id;primaryKey"`
	PlayerInformationId int               `gorm:"column:pp_pi_id"`
	PlayerInformation   PlayerInformation `gorm:"foreignKey:PlayerInformationId"`
	Position            int               `gorm:"column:pp_position"`
	CreatedAt           time.Time         `gorm:"column:pp_created_at"`
	CreatedBy           int               `gorm:"column:pp_created_by"`
	UpdatedAt           time.Time         `gorm:"column:pp_updated_at"`
	UpdatedBy           int               `gorm:"column:pp_updated_by"`
}
