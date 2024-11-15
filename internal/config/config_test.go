package config

import (
	"testing"
	"time"
)

func TestCoalesce(t *testing.T) {
	tests := []struct {
		name     string
		first    interface{}
		second   interface{}
		expected interface{}
	}{
		{"IntegerFirstDefault", 0, 42, 42},
		{"IntegerFirstNonDefault", 10, 42, 10},
		{"StringFirstDefault", "", "default", "default"},
		{"StringFirstNonDefault", "value", "default", "value"},
		{"BoolFirstDefault", false, true, true},
		{"BoolFirstNonDefault", true, false, true},
		{"DurationFirstDefault", time.Duration(0), time.Second, time.Second},
		{"DurationFirstNonDefault", 5 * time.Second, time.Second, 5 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.expected.(type) {
			case int:
				result := coalesce(tt.first.(int), tt.second.(int))
				if result != tt.expected.(int) {
					t.Errorf("Coalesce(%v, %v) = %v; want %v", tt.first, tt.second, result, tt.expected)
				}
			case string:
				result := coalesce(tt.first.(string), tt.second.(string))
				if result != tt.expected.(string) {
					t.Errorf("Coalesce(%v, %v) = %v; want %v", tt.first, tt.second, result, tt.expected)
				}
			case bool:
				result := coalesce(tt.first.(bool), tt.second.(bool))
				if result != tt.expected.(bool) {
					t.Errorf("Coalesce(%v, %v) = %v; want %v", tt.first, tt.second, result, tt.expected)
				}
			case time.Duration:
				result := coalesce(tt.first.(time.Duration), tt.second.(time.Duration))
				if result != tt.expected.(time.Duration) {
					t.Errorf("Coalesce(%v, %v) = %v; want %v", tt.first, tt.second, result, tt.expected)
				}
			default:
				t.Errorf("Unsupported type for test case: %v", tt.expected)
			}
		})
	}
}
