package serverMux

import (
	"net/http"
	"reflect"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func Test_headerMatcher(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "test",
			args: args{
				header: "x-merchant-id",
			},
			want:  "x-merchant-id",
			want1: true,
		},
		{
			name: "test",
			args: args{
				header: "x-auth-type",
			},
			want:  "x-auth-type",
			want1: true,
		},
		{
			name: "test",
			args: args{
				header: "authorization",
			},
			want:  "authorization",
			want1: true,
		},
		{
			name: "test",
			args: args{
				header: "x-user-id",
			},
			want:  "x-user-id",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := headerMatcher(tt.args.header)
			if got != tt.want {
				t.Errorf("headerMatcher() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("headerMatcher() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultMarshalOption(t *testing.T) {
	tests := []struct {
		name string
		want runtime.Marshaler
	}{
		{
			name: "test",
			want: &runtime.JSONPb{
				MarshalOptions:   protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true},
				UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := http.NewRequest("GET", "http://example.com", nil)
			if err != nil {
				t.Fatalf(`http.NewRequest("GET", "http://example.com", nil failed with %v; want success`, err)
			}

			got := runtime.NewServeMux(DefaultMarshalOption())
			inb, _ := runtime.MarshalerForRequest(got, request)
			if !reflect.DeepEqual(inb, tt.want) {
				t.Fatalf(`invalid inbound marshaller`)
			}
		})
	}
}
