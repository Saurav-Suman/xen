package disbursements

import (
	"time"
	"xend/internal/pkg/models"
)

type Location struct {
	BranchID     string `json:"branch_id"`
	BranchName   string `json:"branch_name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	PhoneNumber  string `json:"phone_number"`
	OperatorName string `json:"operator_name"`
	Metadata     struct {
		ExtraField1 float64   `json:"extra_field_1"`
		ExtraField2 time.Time `json:"extra_field_2"`
	} `json:"metadata"`
}

type Disbursements struct {
	models.Model
	RefNumber         string    `gorm:"column:ref_number;not null; json:"ref_number" form:"ref_number"`
	DisbursementCode  string    `gorm:"column:disbursement_code;not null;" json:"disbursement_code" form:"disbursement_code"`
	Location          string    `gorm:"column:location;json:"location" form:"location"`
	ChannelCode       string    `gorm:"column:channel_code;not null;" json:"channel_code"`
	Status            int       `gorm:"column:status;not null;" json:"status"`
	BeneficiaryIDName string    `gorm:"column:beneficiary_id_name;not null;" json:"beneficiary_id_name"`
	Amount            float64   `gorm:"column:amount;not null;" json:"amount"`
	Currency          string    `gorm:"column:currency;not null;" json:"currency"`
	Description       string    `gorm:"column:description;not null;" json:"description"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null;" json:"updated_at"`
}

func (m *Disbursements) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Disbursements) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
