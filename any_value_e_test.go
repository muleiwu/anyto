package anyto

import "testing"

func TestAnyValueE_Int_Success(t *testing.T) {
	got, err := Anyto("42").To().ValueE().Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 42 {
		t.Errorf("Int() = %d, want 42", got)
	}
}

func TestAnyValueE_Int_Error(t *testing.T) {
	got, err := Anyto("not-int").To().ValueE().Int()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != 0 {
		t.Errorf("Int() = %d, want 0 on error", got)
	}
}

func TestAnyValueE_Bool_Success(t *testing.T) {
	got, err := Anyto(true).To().ValueE().Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got {
		t.Error("Bool() = false, want true")
	}
}

func TestAnyValueE_String(t *testing.T) {
	got, err := Anyto(123).To().ValueE().String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "123" {
		t.Errorf("String() = %q, want %q", got, "123")
	}
}

func TestAnyValueE_Float64_Error(t *testing.T) {
	_, err := Anyto("not-float").To().ValueE().Float64()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
