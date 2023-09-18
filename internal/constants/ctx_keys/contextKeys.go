package ctxkeys

import "github.com/gaganchawara/loans/internal/constants"

const (
	RpcMethodKey  = "method"
	UriKey        = "uri"
	GitCommitHash = "git_commit_hash"
	TraceId       = "trace_id"
	UserID        = "user_id"
	AdminID       = "admin_id"
	AccountType   = "account_type"
)

func AllKeys() []string {
	return []string{
		RpcMethodKey,
		UriKey,
		GitCommitHash,
		TraceId,
		UserID,
		AdminID,
		AccountType,
	}
}

func HeaderKeyMap() map[string]string {
	return map[string]string{
		RpcMethodKey:                   RpcMethodKey,
		UriKey:                         UriKey,
		TraceId:                        TraceId,
		GitCommitHash:                  GitCommitHash,
		constants.UserIDHeaderKey:      UserID,
		constants.AdminIDHeaderKey:     AdminID,
		constants.AccountTypeHeaderKey: AccountType,
	}
}
