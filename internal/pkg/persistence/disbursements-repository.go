package persistence

import (
	"fmt"
	models "xend/internal/pkg/models/disbursements"
)

type DisbursementsRepository struct{}

var disbursementsRepository *DisbursementsRepository

func GetDisbursementsRepository() *DisbursementsRepository {
	if disbursementsRepository == nil {
		disbursementsRepository = &DisbursementsRepository{}
	}
	return disbursementsRepository
}

func (r *DisbursementsRepository) Get(id string) (*models.Disbursements, error) {
	var disbursement models.Disbursements
	where := models.Disbursements{}
	where.DisbursementCode = id
	err := Find(&where, &disbursement, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//disbursement.Location=json.Unmarshal(disbursement.Location, &a)
	return &disbursement, err
}

func (r *DisbursementsRepository) GetRef(id string) (*models.Disbursements, error) {
	var disbursement models.Disbursements
	where := models.Disbursements{}
	where.RefNumber = id
	err := Find(&where, &disbursement, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//disbursement.Location=json.Unmarshal(disbursement.Location, &a)
	return &disbursement, err
}

func (r *DisbursementsRepository) UpdateStatus(id string, val string) (*models.Disbursements, error) {
	var disbursement models.Disbursements
	where := models.Disbursements{}
	where.RefNumber = id
	err := Find(&where, &disbursement, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	disbursement.Status = val
	err = Updates(&where, &disbursement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//disbursement.Location=json.Unmarshal(disbursement.Location, &a)
	return &disbursement, err
}
