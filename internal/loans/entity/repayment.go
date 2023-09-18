package entity

import (
	"github.com/gaganchawara/loans/internal/enums/repaymentstatus"
	"time"
)

type Repayment struct {
	Id         string               `gorm:"column:id" json:"id"`
	LoanId     string               `gorm:"column:loan_id" json:"loan_id"`
	Amount     int64                `gorm:"column:amount" json:"amount"`
	PaidAmount int64                `gorm:"column:paid_amount" json:"paid_amount"`
	Status     repaymentstatus.Type `gorm:"column:status" json:"status"`
	DueDate    time.Time            `gorm:"column:due_date" json:"due_date"`
	Entity
}

const TableRepayment = "repayment"

func (model *Repayment) TableName() string {
	return TableRepayment
}