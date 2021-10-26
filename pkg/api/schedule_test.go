package api

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestErgast_Schedule(t *testing.T) {
	type fields struct {
		BaseURL    string
		HTTPClient *http.Client
	}
	type args struct {
		ctx  context.Context
		year int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Race
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ergast := &Ergast{
				BaseURL:    tt.fields.BaseURL,
				HTTPClient: tt.fields.HTTPClient,
			}
			got, err := ergast.Schedule(tt.args.ctx, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ergast.Schedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ergast.Schedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
