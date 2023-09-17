package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

const (
	GitCommitHash = "git_commit_hash"
)

// UnaryServerGitCommitHashInterceptor and sets git_commit_id in the gRPC context.
func UnaryServerGitCommitHashInterceptor(gitCommit string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, GitCommitHash, gitCommit)
		return handler(ctx, req)
	}
}
