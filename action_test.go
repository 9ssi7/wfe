package wfe

import (
	"context"
	"errors"
	"testing"
)

func TestAction(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		errRefs []string
		fn      ActionRunner[int]
		wantErr bool
	}{
		{
			name:    "Successful Run",
			ref:     "testAction",
			errRefs: nil,
			fn: func(ctx context.Context, p int) error {
				return nil
			},
			wantErr: false,
		},
		{
			name:    "Error Run",
			ref:     "errorAction",
			errRefs: []string{"someError"},
			fn: func(ctx context.Context, p int) error {
				return errors.New("some error")
			},
			wantErr: true,
		},
		{
			name:    "Multiple Error Refs",
			ref:     "multiErrorAction",
			errRefs: []string{"error1", "error2"},
			fn: func(ctx context.Context, p int) error {
				return errors.New("some error")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			action := NewAction[int](tt.ref, tt.fn, tt.errRefs...)

			if got := action.Reference(); got != tt.ref {
				t.Errorf("Reference() = %v, want %v", got, tt.ref)
			}

			if tt.errRefs == nil {
				if got := action.ErrorRef(); got != nil {
					t.Errorf("ErrorRef() = %v, want nil", got)
				}
			} else {
				if got := action.ErrorRef(); got == nil || *got != tt.errRefs[0] {
					t.Errorf("ErrorRef() = %v, want %v", got, tt.errRefs[0])
				}
			}

			err := action.Run(context.Background(), 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
