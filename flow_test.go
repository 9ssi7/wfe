package wfe

import "testing"

func TestNewWithCron(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		flowName string
		trigger  string
		wantName string
		wantKind Kind
	}{
		{
			name:     "Basic Cron Flow",
			flowName: "testCronFlow",
			trigger:  "@every 1s",
			wantName: "testCronFlow",
			wantKind: KindCron,
		},
		{
			name:     "Basic Cron Flow",
			flowName: "testCronFlow",
			trigger:  "@every 1s",
			wantName: "testCronFlow",
			wantKind: KindCron,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewWithCron[string](tt.flowName, tt.trigger)
			if got := f.Name(); got != tt.wantName {
				t.Errorf("Name() = %v, want %v", got, tt.wantName)
			}
			if got := f.Kind(); got != tt.wantKind {
				t.Errorf("Kind() = %v, want %v", got, tt.wantKind)
			}
		})
	}
}
