package ctxkeys

const (
	RpcMethodKey  = "method"
	UriKey        = "uri"
	GitCommitHash = "git_commit_hash"
	TraceId       = "trace_id"
)

func AllKeys() []string {
	return []string{
		RpcMethodKey,
		UriKey,
		TraceId,
		GitCommitHash,
	}
}
