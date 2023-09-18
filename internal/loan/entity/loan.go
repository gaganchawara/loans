package entity

import (
	"time"

	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"

	"github.com/gaganchawara/loans/internal/enums/loanstatus"
)

type Loan struct {
	Id          string          `gorm:"column:id" json:"id"`
	UserId      string          `gorm:"column:user_id" json:"user_id"`
	Amount      int64           `gorm:"column:amount" json:"amount"`
	Term        int32           `gorm:"column:term" json:"term"`
	Status      loanstatus.Type `gorm:"column:status" json:"status"`
	ApprovedBy  string          `gorm:"column:approved_by" json:"approved_by"`
	DisbursedAt time.Time       `gorm:"column:disbursed_at" json:"disbursed_at"`
	Entity
}

const TableLoan = "loan"

func (e *Loan) TableName() string {
	return TableLoan
}

func (e *Loan) Proto() *loansv1.Loan {
	var disbursedAt = e.DisbursedAt.Unix()

	return &loansv1.Loan{
		Id:          e.Id,
		UserId:      e.UserId,
		Amount:      e.Amount,
		Term:        e.Term,
		Status:      e.Status.String(),
		ApprovedBy:  &e.ApprovedBy,
		DisbursedAt: &disbursedAt,
		CreatedAt:   e.CreatedAt.Unix(),
		UpdatedAt:   e.UpdatedAt.Unix(),
		DeletedAt:   e.DeletedAt.Unix(),
	}
}
