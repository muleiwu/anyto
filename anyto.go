// Package anyto provides a fluent API for type conversion, wrapping github.com/spf13/cast.
package anyto

//go:generate go run ./gen/...

// Anyto creates an AnyValue from any value, serving as the entry point for the fluent API.
func Anyto(v any) AnyValue {
	return AnyValue{v: v}
}

// AnyValue holds a value and provides direct type conversion methods that return T.
// On conversion failure, methods return the zero value of T.
type AnyValue struct {
	v any
}

// To returns an AnyTo navigator for choosing between Value, ValueE, Pointer, and PointerE modes.
func (a AnyValue) To() AnyTo {
	return AnyTo{v: a.v}
}

// Pointer returns an AnyPointer for pointer-returning conversions.
func (a AnyValue) Pointer() AnyPointer {
	return AnyPointer{v: a.v}
}

// AnyTo provides navigation to the four conversion modes.
type AnyTo struct {
	v any
}

// Value returns an AnyValue for conversions that return T.
func (a AnyTo) Value() AnyValue {
	return AnyValue{v: a.v}
}

// ValueE returns an AnyValueE for conversions that return (T, error).
func (a AnyTo) ValueE() AnyValueE {
	return AnyValueE{v: a.v}
}

// Pointer returns an AnyPointer for conversions that return *T.
func (a AnyTo) Pointer() AnyPointer {
	return AnyPointer{v: a.v}
}

// PointerE returns an AnyPointerE for conversions that return (*T, error).
func (a AnyTo) PointerE() AnyPointerE {
	return AnyPointerE{v: a.v}
}

// AnyValueE holds a value and provides type conversion methods that return (T, error).
type AnyValueE struct {
	v any
}

// AnyPointer holds a value and provides type conversion methods that return *T.
// On conversion failure, methods return nil.
type AnyPointer struct {
	v any
}

// AnyPointerE holds a value and provides type conversion methods that return (*T, error).
// On conversion failure, methods return (nil, error).
type AnyPointerE struct {
	v any
}
