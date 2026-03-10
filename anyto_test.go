package anyto

import "testing"

func TestAnyto_ReturnsAnyValue(t *testing.T) {
	av := Anyto("hello")
	if got := av.String(); got != "hello" {
		t.Errorf("Anyto().String() = %q, want %q", got, "hello")
	}
}

func TestAnyValue_To_Value(t *testing.T) {
	got := Anyto(42).To().Value().Int()
	if got != 42 {
		t.Errorf("To().Value().Int() = %d, want 42", got)
	}
}

func TestAnyValue_To_ValueE(t *testing.T) {
	got, err := Anyto(42).To().ValueE().Int()
	if err != nil {
		t.Errorf("To().ValueE().Int() error = %v", err)
	}
	if got != 42 {
		t.Errorf("To().ValueE().Int() = %d, want 42", got)
	}
}

func TestAnyValue_To_Pointer(t *testing.T) {
	got := Anyto(42).To().Pointer().Int()
	if got == nil || *got != 42 {
		t.Errorf("To().Pointer().Int() unexpected result")
	}
}

func TestAnyValue_To_PointerE(t *testing.T) {
	got, err := Anyto(42).To().PointerE().Int()
	if err != nil {
		t.Errorf("To().PointerE().Int() error = %v", err)
	}
	if got == nil || *got != 42 {
		t.Errorf("To().PointerE().Int() unexpected result")
	}
}

func TestAnyValue_Pointer_Shortcut(t *testing.T) {
	got := Anyto("true").Pointer().Bool()
	if got == nil || *got != true {
		t.Errorf("Pointer().Bool() unexpected result")
	}
}
