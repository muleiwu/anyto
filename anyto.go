// Package anyto provides a fluent API for type conversion, wrapping github.com/spf13/cast.
package anyto

import "github.com/muleiwu/gsr"

//go:generate go run ./gen/...

// Anyto creates an AnyValue from any value, serving as the entry point for the fluent API.
func Anyto(v any) gsr.AnyValue {
	return &anyValue{v: v}
}

// anyValue holds a value and provides direct type conversion methods that return T.
// On conversion failure, methods return the zero value of T.
type anyValue struct {
	v any
}

// To returns an AnyTo navigator for choosing between Value, ValueE, Pointer, and PointerE modes.
func (a *anyValue) To() gsr.AnyTo {
	return &anyTo{v: a.v}
}

// Pointer returns an AnyPointer for pointer-returning conversions.
func (a *anyValue) Pointer() gsr.AnyPointer {
	return &anyPointer{v: a.v}
}

// anyTo provides navigation to the four conversion modes.
type anyTo struct {
	v any
}

// Value returns an AnyValue for conversions that return T.
func (a *anyTo) Value() gsr.AnyValue {
	return &anyValue{v: a.v}
}

// ValueE returns an AnyValueE for conversions that return (T, error).
func (a *anyTo) ValueE() gsr.AnyValueE {
	return &anyValueE{v: a.v}
}

// Pointer returns an AnyPointer for conversions that return *T.
func (a *anyTo) Pointer() gsr.AnyPointer {
	return &anyPointer{v: a.v}
}

// PointerE returns an AnyPointerE for conversions that return (*T, error).
func (a *anyTo) PointerE() gsr.AnyPointerE {
	return &anyPointerE{v: a.v}
}

// anyValueE holds a value and provides type conversion methods that return (T, error).
type anyValueE struct {
	v any
}

// anyPointer holds a value and provides type conversion methods that return *T.
// On conversion failure, methods return nil.
type anyPointer struct {
	v any
}

// anyPointerE holds a value and provides type conversion methods that return (*T, error).
// On conversion failure, methods return (nil, error).
type anyPointerE struct {
	v any
}
