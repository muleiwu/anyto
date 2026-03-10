package anyto

import "testing"

func TestAnyPointer_Int_Success(t *testing.T) {
	got := Anyto("42").Pointer().Int()
	if got == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *got != 42 {
		t.Errorf("*Int() = %d, want 42", *got)
	}
}

func TestAnyPointer_Int_Error(t *testing.T) {
	got := Anyto("not-int").Pointer().Int()
	if got != nil {
		t.Errorf("expected nil pointer, got %v", *got)
	}
}

func TestAnyPointer_Bool_Success(t *testing.T) {
	got := Anyto(true).Pointer().Bool()
	if got == nil || !*got {
		t.Error("expected non-nil pointer to true")
	}
}

func TestAnyPointer_String_Success(t *testing.T) {
	got := Anyto(123).Pointer().String()
	if got == nil || *got != "123" {
		t.Errorf("unexpected result: %v", got)
	}
}

func TestAnyPointer_StringSlice_Success(t *testing.T) {
	got := Anyto([]any{"a", "b"}).Pointer().StringSlice()
	if got == nil || len(*got) != 2 {
		t.Errorf("unexpected result: %v", got)
	}
}
