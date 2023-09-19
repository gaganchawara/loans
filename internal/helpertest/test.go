package helpertest

import (
	"database/sql"
	"time"

	"github.com/gaganchawara/loans/internal/enums/loanstatus"

	"github.com/gaganchawara/loans/internal/enums/repaymentstatus"
	"github.com/gaganchawara/loans/internal/loan/entity"
)

func GetTestLoanEntity() *entity.Loan {
	now := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	return &entity.Loan{
		Id:         "loanId1",
		Amount:     10000,
		UserId:     "randomUserId",
		Term:       3,
		Status:     loanstatus.Pending,
		ApprovedBy: "randomAdminId",
		Entity: entity.Entity{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

func GetTestRepaymentEntity() *entity.Repayment {
	now := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	return &entity.Repayment{
		Id:         "repaymentId1",
		LoanId:     "loanId1",
		Amount:     10000,
		PaidAmount: 0,
		Status:     repaymentstatus.Pending,
		DueDate: sql.NullTime{
			Time:  time.Now().AddDate(0, 0, 7),
			Valid: true,
		},
		Entity: entity.Entity{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}
