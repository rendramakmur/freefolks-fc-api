package domains

import "time"

type GameRegistration struct {
	Id                  int               `gorm:"column:gr_id;primaryKey"`
	PlayerInformationId int               `gorm:"column:gr_pi_id"`
	PlayerInformation   PlayerInformation `gorm:"foreginKey:PlayerInformationId"`
	GameDataId          int               `gorm:"column:gr_gd_id"`
	GameData            GameData          `gorm:"foreignKey:GameDataId"`
	IsOutfield          string            `gorm:"column:gr_is_outfield"`
	Amount              float64           `gorm:"column:gr_amount"`
	TransactionNumber   string            `gorm:"column:gr_transaction_number"`
	CreatedAt           time.Time         `gorm:"column:gr_created_at"`
	CreatedBy           int               `gorm:"column:gr_created_by"`
	UpdatedAt           time.Time         `gorm:"column:gr_updated_at"`
	UpdatedBy           int               `gorm:"column:gr_updated_by"`
}
