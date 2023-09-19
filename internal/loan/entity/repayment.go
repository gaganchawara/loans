package entity

import (
	"time"

	"github.com/gaganchawara/loans/pkg/utils"
	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"

	"github.com/gaganchawara/loans/internal/enums/repaymentstatus"
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

func (e *Repayment) TableName() string {
	return TableRepayment
}

func NewRepaymentEntity() *Repayment {
	e := &Repayment{
		Id:     utils.GenerateUniqueID(),
		Status: repaymentstatus.Pending,
	}
	e.RefreshTimestamps()

	return e
}

func (e *Repayment) Proto() *loansv1.Repayment {
	proto := loansv1.Repayment{
		Id:         e.Id,
		LoanId:     e.LoanId,
		Amount:     e.Amount,
		PaidAmount: e.PaidAmount,
		Status:     e.Status.String(),
		DueDate:    e.DueDate.Unix(),
		CreatedAt:  e.CreatedAt.Unix(),
		UpdatedAt:  e.UpdatedAt.Unix(),
	}

	if e.DeletedAt != nil {
		proto.DeletedAt = e.DeletedAt.Unix()
	}

	return &proto
}
