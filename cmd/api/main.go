package main

import (
	"context"

	"github.com/gaganchawara/loans/internal/boot"
)

func main() {
	ctx := context.Background()

	boot.Initialize(ctx)
}
