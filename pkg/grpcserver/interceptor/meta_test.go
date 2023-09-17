package interceptors

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestHeaderInterceptor(t *testing.T) {
	headerTags := []string{"tag1", "tag2"}
	interceptor := HeaderInterceptor(headerTags)

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		// Verify that the expected values have been set in the context.
		value := ctx.Value("tag1")
		if value == nil || value.(string) != "value1" {
			t.Errorf("Expected 'tag1' to be set in context with value 'value1', got: %v", value)
		}

		return nil, nil
	}

	// Create a context with incoming metadata.
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("tag1", "value1"))

	_, err := interceptor(ctx, nil, &grpc.UnaryServerInfo{}, handler)
	if err != nil {
		t.Errorf("Interceptor error: %v", err)
	}
}

func Test_extractAndSetMetadataValues(t *testing.T) {
	headerTags := []string{"tag1", "tag2"}

	// Create a context with incoming metadata.
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("tag1", "value1", "tag2", "value2", "tag3", "value3"))

	// Call the function to extract and set metadata values in the context.
	ctx = extractAndSetMetadataValues(ctx, headerTags)

	// Verify that the expected values have been set in the context.
	value1 := ctx.Value("tag1")
	if value1 == nil || value1.(string) != "value1" {
		t.Errorf("Expected 'tag1' to be set in context with value 'value1', got: %v", value1)
	}

	value2 := ctx.Value("tag2")
	if value2 == nil || value2.(string) != "value2" {
		t.Errorf("Expected 'tag2' to be set in context with value 'value2', got: %v", value2)
	}

	// Verify that 'tag3' is not set in the context.
	value3 := ctx.Value("tag3")
	if value3 != nil {
		t.Errorf("Expected 'tag3' not to be set in context, but it has a value: %v", value3)
	}
}

func Test_in(t *testing.T) {
	// Test when the needle exists in the haystack.
	haystack := []string{"abc", "def", "ghi"}
	needle := "def"
	if !in(needle, haystack) {
		t.Errorf("Expected 'Def' to be found in 'haystack', but it wasn't.")
	}

	// Test when the needle does not exist in the haystack.
	needle = "xyz"
	if in(needle, haystack) {
		t.Errorf("Expected 'xyz' not to be found in 'haystack', but it was.")
	}
}
