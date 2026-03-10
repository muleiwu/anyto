# anyto

[![Go Reference](https://pkg.go.dev/badge/github.com/muleiwu/anyto.svg)](https://pkg.go.dev/github.com/muleiwu/anyto)
[![Go Report Card](https://goreportcard.com/badge/github.com/muleiwu/anyto)](https://goreportcard.com/report/github.com/muleiwu/anyto)

[English](README.md) | [中文](README.zh-CN.md)

基于 [spf13/cast](https://github.com/spf13/cast) 封装的 Go 类型转换流式 API。

将 `any` 通过链式调用转换为目标类型 — 支持**值**、**指针**和**错误感知**模式，覆盖 **38 种类型**（共 152 个方法）。

## 安装

```bash
go get github.com/muleiwu/anyto
```

## 快速开始

```go
import "github.com/muleiwu/anyto"

// 值转换（失败时返回零值）
s := anyto.Anyto("42").Int()          // 42
b := anyto.Anyto("true").Bool()       // true
f := anyto.Anyto("3.14").Float64()    // 3.14

// 值转换 + 错误
n, err := anyto.Anyto("42").To().ValueE().Int()

// 指针转换（失败时返回 nil）
p := anyto.Anyto("42").Pointer().Int() // *int → 42

// 指针转换 + 错误
p, err := anyto.Anyto("42").To().PointerE().Int()
```

## API 概览

### 入口函数

```go
anyto.Anyto(v any) AnyValue
```

### 四种转换模式

| 模式 | 调用方式 | 返回值 | 失败时 |
|------|----------|--------|--------|
| **值** | `Anyto(v).Xxx()` | `T` | 零值 |
| **值 + 错误** | `Anyto(v).To().ValueE().Xxx()` | `(T, error)` | `(零值, error)` |
| **指针** | `Anyto(v).Pointer().Xxx()` | `*T` | `nil` |
| **指针 + 错误** | `Anyto(v).To().PointerE().Xxx()` | `(*T, error)` | `(nil, error)` |

### 导航结构

```
Anyto(v) → AnyValue
             ├── .Xxx()           → T           （直接快捷方式）
             ├── .Pointer().Xxx() → *T          （指针快捷方式）
             └── .To() → AnyTo
                          ├── .Value()   → AnyValue   → .Xxx() → T
                          ├── .ValueE()  → AnyValueE  → .Xxx() → (T, error)
                          ├── .Pointer() → AnyPointer  → .Xxx() → *T
                          └── .PointerE()→ AnyPointerE → .Xxx() → (*T, error)
```

### 支持的类型

| 分类 | 方法 |
|------|------|
| **基础类型** | `Bool`, `String` |
| **整数** | `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64` |
| **浮点数** | `Float32`, `Float64` |
| **时间** | `Time`, `Duration` |
| **切片** | `Slice`, `BoolSlice`, `StringSlice`, `IntSlice`, `Int8Slice`, `Int16Slice`, `Int32Slice`, `Int64Slice`, `UintSlice`, `Uint8Slice`, `Uint16Slice`, `Uint32Slice`, `Uint64Slice`, `Float32Slice`, `Float64Slice`, `DurationSlice` |
| **映射** | `StringMap`, `StringMapString`, `StringMapStringSlice`, `StringMapBool`, `StringMapInt`, `StringMapInt64` |

## 使用示例

### 处理 JSON 解码数据

```go
data := map[string]any{"port": "8080", "debug": "true", "rate": "0.75"}

port := anyto.Anyto(data["port"]).Int()          // 8080
debug := anyto.Anyto(data["debug"]).Bool()        // true
rate := anyto.Anyto(data["rate"]).Float64()       // 0.75
```

### 可选字段的安全指针转换

```go
func getConfig(raw map[string]any) *int {
    return anyto.Anyto(raw["timeout"]).Pointer().Int()
    // 如果 "timeout" 键不存在或无法转换，返回 nil
}
```

### 带错误处理的转换

```go
val, err := anyto.Anyto(input).To().ValueE().Int()
if err != nil {
    log.Printf("无效输入: %v", err)
    return
}
```

## 代码生成

转换方法通过类型表自动生成。如需重新生成：

```bash
go generate ./...
```

## 许可证

[MIT](LICENSE)
