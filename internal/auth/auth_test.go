package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"no auth header",
			args{
				headers: http.Header{},
			},
			"",
			true,
		},
		{
			"malformed auth header",
			args{
				headers: http.Header{
					"Authorization": []string{"Bearer"},
				},
			},
			"",
			true,
		},
		{
			"correct auth header",
			args{
				headers: http.Header{
					"Authorization": []string{"ApiKey test"},
				},
			},
			"test",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
