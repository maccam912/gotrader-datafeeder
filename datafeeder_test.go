package datafeeder

import "testing"

func TestRandomDataProducer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go RandomDataProducer()
		})
	}
}
