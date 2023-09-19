package entity

import (
	"context"
	"time"

	"github.com/gaganchawara/loans/internal/iam"

	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/gaganchawara/loans/pkg/utils"

	loansv1 "github.com/gaganchawara/loans/rpc/loans/v1"

	"github.com/gaganchawara/loans/internal/enums/loanstatus"
)

type Loan struct {
	Id          string          `gorm:"column:id;primary_key" json:"id"`
	UserId      string          `gorm:"column:user_id" json:"user_id"`
	Amount      int64           `gorm:"column:amount" json:"amount"`
	Term        int32           `gorm:"column:term" json:"term"`
	Status      loanstatus.Type `gorm:"column:status" json:"status"`
	ApprovedBy  string          `gorm:"column:approved_by" json:"approved_by"`
	DisbursedAt *time.Time      `gorm:"column:disbursed_at" json:"disbursed_at"`
	Entity
}

const TableLoan = "loan"

func NewLoanEntity(_ context.Context) (*Loan, errors.Error) {
	loan := &Loan{
		Id:     utils.GenerateUniqueID(),
		Status: loanstatus.Pending,
	}
	loan.RefreshTimestamps()

	return loan, nil
}

func (e *Loan) TableName() string {
	return TableLoan
}

func (e *Loan) MarkApproved(ctx context.Context) errors.Error {
	adminId, ierr := iam.GetAdminId(ctx)
	if ierr != nil {
		return ierr
	}

	e.ApprovedBy = adminId
	now := time.Now()
	e.DisbursedAt = &now

	return nil
}

func (e *Loan) MarkRejected(_ context.Context) errors.Error {
	e.Status = loanstatus.Rejected

	return nil
}

func (e *Loan) Proto() *loansv1.Loan {
	proto := loansv1.Loan{
		Id:         e.Id,
		UserId:     e.UserId,
		Amount:     e.Amount,
		Term:       e.Term,
		Status:     e.Status.String(),
		ApprovedBy: &e.ApprovedBy,
		CreatedAt:  e.CreatedAt.Unix(),
		UpdatedAt:  e.UpdatedAt.Unix(),
	}

	if e.DisbursedAt != nil {
		disbursedAt := e.DisbursedAt.Unix()
		proto.DisbursedAt = &disbursedAt
	}

	if e.DeletedAt != nil {
		proto.DeletedAt = e.DeletedAt.Unix()
	}

	return &proto
}
