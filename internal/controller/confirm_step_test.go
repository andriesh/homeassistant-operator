package controller

import (
	"testing"

	"github.com/przemekhys/homeassistant-operator/internal/haclient"
)

func TestConfirmStepData(t *testing.T) {
	tests := []struct {
		name string
		form *haclient.FlowResponse
		want string
	}{
		{"no schema falls back to confirmed", &haclient.FlowResponse{}, "confirmed"},
		{
			"newer HA advertises confirmed_ok",
			&haclient.FlowResponse{DataSchema: []haclient.FlowField{{Name: "confirmed_ok"}}},
			"confirmed_ok",
		},
		{
			"older HA advertises confirmed",
			&haclient.FlowResponse{DataSchema: []haclient.FlowField{{Name: "confirmed"}}},
			"confirmed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := confirmStepData(tt.form)
			if v, ok := got[tt.want].(bool); !ok || !v || len(got) != 1 {
				t.Errorf("expected {%s: true}, got %v", tt.want, got)
			}
		})
	}
}
