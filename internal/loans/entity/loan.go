package entity

import (
	"github.com/gaganchawara/loans/internal/enums/loanstatus"
	"time"
)

type Loan struct {
	Id         string          `gorm:"column:id" json:"id"`
	UserId     string          `gorm:"column:user_id" json:"user_id"`
	Amount     int64           `gorm:"column:amount" json:"amount"`
	Term       int32           `gorm:"column:term" json:"term"`
	Status     loanstatus.Type `gorm:"column:status" json:"status"`
	ApprovedBy string          `gorm:"column:approved_by" json:"approved_by"`
	DisbursedAt time.Time      `gorm:"column:disbursed_at" json:"disbursed_at"`
	Entity
}

const TableLoan = "loan"

func (model *Loan) TableName() string {
	return TableLoan
}