package anyto

import "testing"

func TestAnyPointerE_Int_Success(t *testing.T) {
	got, err := Anyto("42").To().PointerE().Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || *got != 42 {
		t.Errorf("unexpected result: %v", got)
	}
}

func TestAnyPointerE_Int_Error(t *testing.T) {
	got, err := Anyto("not-int").To().PointerE().Int()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != nil {
		t.Errorf("expected nil pointer, got %v", *got)
	}
}

func TestAnyPointerE_Bool_Success(t *testing.T) {
	got, err := Anyto(true).To().PointerE().Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || !*got {
		t.Error("expected non-nil pointer to true")
	}
}

func TestAnyPointerE_String_Success(t *testing.T) {
	got, err := Anyto(123).To().PointerE().String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || *got != "123" {
		t.Errorf("unexpected result: %v", got)
	}
}

func TestAnyPointerE_Float64_Error(t *testing.T) {
	got, err := Anyto("not-float").To().PointerE().Float64()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != nil {
		t.Errorf("expected nil, got %v", *got)
	}
}
