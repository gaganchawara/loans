package interceptors

import (
	"context"
	"encoding/base64"
	"strings"

	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/v2/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type BasicAuthCreds struct {
	Username string
	Password string
}

const (
	AuthorizationHeaderKey = "authorization"
	OriginServiceKey       = "origin_service"
)

// UnaryServerBasicAuthInterceptor creates an authenticator interceptor with the given BasicAuth
func UnaryServerBasicAuthInterceptor(creds []BasicAuthCreds) grpc.UnaryServerInterceptor {
	return grpcauth.UnaryServerInterceptor(BasicAuthFunc(creds))
}

func BasicAuthFunc(creds []BasicAuthCreds) grpcauth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		return BasicAuth(ctx, creds)
	}
}

// BasicAuth implements basic authentication for services
func BasicAuth(ctx context.Context, authCreds []BasicAuthCreds) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.New(codes.Unauthenticated, "couldn't parse incoming context").Err()
	}

	headers := md.Get(AuthorizationHeaderKey)
	if len(headers) < 1 {
		return ctx, status.New(codes.Unauthenticated, "invalid authorization headers").Err()
	}

	username, password, ok := parseBasicAuth(headers[0])
	if !ok {
		return ctx, status.New(codes.Unauthenticated, "parsing auth headers failed").Err()
	}

	if username == "" || password == "" {
		return ctx, status.New(codes.Unauthenticated, "empty username/password").Err()
	}

	for _, cred := range authCreds {
		if cred.Username == username {
			if cred.Password == password {
				ctx = context.WithValue(ctx, OriginServiceKey, username)
				return ctx, nil
			} else {
				return ctx, status.New(codes.Unauthenticated, "invalid password").Err()
			}
		}
	}

	return ctx, status.New(codes.Unauthenticated, "invalid username/password").Err()
}

// parseBasicAuth parses authorization header value to return username, password strings
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}

	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}

	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}

	return cs[:s], cs[s+1:], true
}
