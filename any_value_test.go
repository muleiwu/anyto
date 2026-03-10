package anyto

import (
	"testing"
	"time"
)

func TestAnyValue_Bool(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want bool
	}{
		{"true", true, true},
		{"string true", "true", true},
		{"false", false, false},
		{"invalid returns zero", "not-a-bool", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anyto(tt.in).Bool(); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyValue_String(t *testing.T) {
	if got := Anyto(123).String(); got != "123" {
		t.Errorf("String() = %q, want %q", got, "123")
	}
}

func TestAnyValue_Int(t *testing.T) {
	if got := Anyto("42").Int(); got != 42 {
		t.Errorf("Int() = %d, want 42", got)
	}
	if got := Anyto("not-int").Int(); got != 0 {
		t.Errorf("Int() on invalid = %d, want 0", got)
	}
}

func TestAnyValue_Float64(t *testing.T) {
	if got := Anyto("3.14").Float64(); got != 3.14 {
		t.Errorf("Float64() = %f, want 3.14", got)
	}
}

func TestAnyValue_Duration(t *testing.T) {
	if got := Anyto("1s").Duration(); got != time.Second {
		t.Errorf("Duration() = %v, want %v", got, time.Second)
	}
}

func TestAnyValue_StringSlice(t *testing.T) {
	in := []any{"a", "b", "c"}
	got := Anyto(in).StringSlice()
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("StringSlice() = %v, want [a b c]", got)
	}
}

func TestAnyValue_StringMap(t *testing.T) {
	in := map[string]any{"key": "value"}
	got := Anyto(in).StringMap()
	if got["key"] != "value" {
		t.Errorf("StringMap() = %v, want map[key:value]", got)
	}
}

func TestAnyValue_Time(t *testing.T) {
	ts := "2023-01-15T10:30:00Z"
	got := Anyto(ts).Time()
	if got.Year() != 2023 || got.Month() != time.January || got.Day() != 15 {
		t.Errorf("Time() = %v, unexpected", got)
	}
}
