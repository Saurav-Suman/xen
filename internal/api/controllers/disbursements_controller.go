package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"xend/internal/pkg/persistence"

	http_err "xend/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type InitiateDisbursement struct {
	DisbursementCode  string `json:"disbursement_code" binding:"required"`
	BeneficiaryIDName string `json:"beneficiary_id_name" binding:"required"`
	Location          struct {
		BranchID     string `json:"branch_id" binding:"required"`
		BranchName   string `json:"branch_name" binding:"required"`
		Address      string `json:"address" binding:"required"`
		City         string `json:"city" binding:"required"`
		PhoneNumber  string `json:"phone_number" binding:"required"`
		OperatorName string `json:"operator_name" binding:"required"`
		Metadata     struct {
			ExtraField1 float64   `json:"extra_field_1"`
			ExtraField2 time.Time `json:"extra_field_2"`
		} `json:"metadata"`
	} `json:"location" binding:"required"`
}

type CommitDisbursement struct {
	RefNumber string  `json:"ref_number" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

//  Status mapping
/*
1 - CREATED
2 - LOCKED
3 - EXPIRED
4 - COMMITED
*/

//var mappiing map[int]string{1:"CREATED",2:"LOCKED",3:"EXPIRED",4:"COMMITED"}

func Initiate(c *gin.Context) {
	s := persistence.DisbursementsRepository{}
	var initiateInput InitiateDisbursement
	err := c.BindJSON(&initiateInput)
	if err != nil {

		http_err.NewError(c, 400, errors.New("INVALID_PARAMS"))
		c.Abort()
		return
	}
	if disbursement, err := s.Get(initiateInput.DisbursementCode); err != nil {
		http_err.NewError(c, 400, errors.New("INVALID_CODE"))
		log.Println(err)
	} else {
		fmt.Println(disbursement)
		if disbursement.Status == 1 {
			disbursementdata, err := s.UpdateStatus(disbursement.RefNumber, 2)
			if err != nil {
				http_err.NewError(c, 400, errors.New("INTERNAL_ERROR"))
				log.Println(err)
			}
			c.JSON(http.StatusOK, disbursementdata)
		} else if disbursement.Status == 2 {
			http_err.NewError(c, 400, errors.New("LOCKED"))
			log.Println(err)
			c.Abort()
			return
		} else if disbursement.Status == 3 {
			http_err.NewError(c, 400, errors.New("EXPIRED"))
			log.Println(err)
			c.Abort()
			return
		}
		//c.JSON(http.StatusOK, disbursement)
	}

}

func Commit(c *gin.Context) {
	s := persistence.DisbursementsRepository{}
	var commitInput CommitDisbursement
	err := c.BindJSON(&commitInput)
	if err != nil {

		http_err.NewError(c, 400, errors.New("INVALID_PARAMS"))
		c.Abort()
		return
	}
	if disbursement, err := s.GetRef(commitInput.RefNumber); err != nil {
		http_err.NewError(c, 400, errors.New("CASH_DISBURSEMENT_NOT_FOUND"))
		log.Println(err)
	} else {
		fmt.Println(disbursement.Status)
		fmt.Println(disbursement)
		if disbursement.Status == 2 {
			disbursementdata, err := s.UpdateStatus(commitInput.RefNumber, 4)
			if err != nil {
				http_err.NewError(c, 400, errors.New("INTERNAL_ERROR"))
				log.Println(err)
			}
			c.JSON(http.StatusOK, disbursementdata)
		} else if disbursement.Status == 1 {
			http_err.NewError(c, 400, errors.New("CASH_DISBURSEMENT_NOT_FOUND"))
			log.Println(err)
			c.Abort()
			return
		} else {
			http_err.NewError(c, 400, errors.New("INVALID_CASH_DISBURSEMENT"))
			log.Println(err)
			c.Abort()
			return
		}

		//c.JSON(http.StatusCreated, disbursement)
	}

}
