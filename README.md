# anyto

[![Go Reference](https://pkg.go.dev/badge/github.com/muleiwu/anyto.svg)](https://pkg.go.dev/github.com/muleiwu/anyto)
[![Go Report Card](https://goreportcard.com/badge/github.com/muleiwu/anyto)](https://goreportcard.com/report/github.com/muleiwu/anyto)

[English](README.md) | [中文](README.zh-CN.md)

A fluent API for Go type conversion, wrapping [spf13/cast](https://github.com/spf13/cast).

Convert `any` to target types with a single chainable call — supports **value**, **pointer**, and **error-aware** modes across **38 types** (152 methods total).

## Installation

```bash
go get github.com/muleiwu/anyto
```

## Quick Start

```go
import "github.com/muleiwu/anyto"

// Value (returns zero value on error)
s := anyto.Anyto("42").Int()          // 42
b := anyto.Anyto("true").Bool()       // true
f := anyto.Anyto("3.14").Float64()    // 3.14

// Value with error
n, err := anyto.Anyto("42").To().ValueE().Int()

// Pointer (returns nil on error)
p := anyto.Anyto("42").Pointer().Int() // *int → 42

// Pointer with error
p, err := anyto.Anyto("42").To().PointerE().Int()
```

## API Overview

### Entry Point

```go
anyto.Anyto(v any) AnyValue
```

### Four Conversion Modes

| Mode | Access | Returns | On Error |
|------|--------|---------|----------|
| **Value** | `Anyto(v).Xxx()` | `T` | zero value |
| **ValueE** | `Anyto(v).To().ValueE().Xxx()` | `(T, error)` | `(zero, error)` |
| **Pointer** | `Anyto(v).Pointer().Xxx()` | `*T` | `nil` |
| **PointerE** | `Anyto(v).To().PointerE().Xxx()` | `(*T, error)` | `(nil, error)` |

### Navigation

```
Anyto(v) → AnyValue
             ├── .Xxx()           → T           (direct shortcut)
             ├── .Pointer().Xxx() → *T          (pointer shortcut)
             └── .To() → AnyTo
                          ├── .Value()   → AnyValue   → .Xxx() → T
                          ├── .ValueE()  → AnyValueE  → .Xxx() → (T, error)
                          ├── .Pointer() → AnyPointer  → .Xxx() → *T
                          └── .PointerE()→ AnyPointerE → .Xxx() → (*T, error)
```

### Supported Types

| Category | Methods |
|----------|---------|
| **Basic** | `Bool`, `String` |
| **Integer** | `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64` |
| **Float** | `Float32`, `Float64` |
| **Time** | `Time`, `Duration` |
| **Slices** | `Slice`, `BoolSlice`, `StringSlice`, `IntSlice`, `Int8Slice`, `Int16Slice`, `Int32Slice`, `Int64Slice`, `UintSlice`, `Uint8Slice`, `Uint16Slice`, `Uint32Slice`, `Uint64Slice`, `Float32Slice`, `Float64Slice`, `DurationSlice` |
| **Maps** | `StringMap`, `StringMapString`, `StringMapStringSlice`, `StringMapBool`, `StringMapInt`, `StringMapInt64` |

## Examples

### Handling JSON-decoded data

```go
data := map[string]any{"port": "8080", "debug": "true", "rate": "0.75"}

port := anyto.Anyto(data["port"]).Int()          // 8080
debug := anyto.Anyto(data["debug"]).Bool()        // true
rate := anyto.Anyto(data["rate"]).Float64()       // 0.75
```

### Safe pointer conversion for optional fields

```go
func getConfig(raw map[string]any) *int {
    return anyto.Anyto(raw["timeout"]).Pointer().Int()
    // Returns nil if "timeout" key is missing or not convertible
}
```

### Error-aware conversion

```go
val, err := anyto.Anyto(input).To().ValueE().Int()
if err != nil {
    log.Printf("invalid input: %v", err)
    return
}
```

## Code Generation

The conversion methods are generated from a type table. To regenerate:

```bash
go generate ./...
```

## License

[MIT](LICENSE)
