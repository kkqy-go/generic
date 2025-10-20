# generic

简洁的 Go 1.18+ 泛型工具库，提供对标准库常用类型的类型安全封装。

包含以下组件：
- Optional[T]：可选值容器，支持 JSON 序列化/反序列化（null ↔ 未呈现）。
- AtomicValue[T]：对 sync/atomic.Value 的类型安全封装（Load/Store/Swap/CompareAndSwap）。
- SyncMap[K, V]：对 sync.Map 的类型安全封装，附加 Len、Clear 等便捷方法。

## 安装

```bash
go get -u github.com/kkqy-go/generic
```

要求：Go 1.18 及以上。

## 快速上手

### Optional[T]
```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/kkqy-go/generic"
)

func main() {
    // 直接构造已呈现的可选值
    o := generic.NewOptional(42)
    fmt.Println(o.Presented()) // true
    fmt.Println(o.Value())     // 42

    // 默认值
    fmt.Println(o.ValueOr(0)) // 42

    // JSON: 已呈现 -> 值
    b, _ := json.Marshal(o)
    fmt.Println(string(b)) // "42"

    // JSON: null -> 未呈现
    var o2 generic.Optional[int]
    _ = json.Unmarshal([]byte("null"), &o2)
    fmt.Println(o2.Presented()) // false

    // 指针访问
    if p := o.Ptr(); p != nil {
        fmt.Println(*p) // 42
    }

    // 清空
    o.Clear()
    fmt.Println(o.Presented()) // false
}
```

### AtomicValue[T]
```go
package main

import (
    "fmt"
    "github.com/kkqy-go/generic"
)

func main() {
    var v generic.AtomicValue[int]

    // 尚未 Store 前，Load 返回类型零值
    fmt.Println(v.Load()) // 0

    v.Store(10)
    fmt.Println(v.Load()) // 10

    old := v.Swap(20)
    fmt.Println(old)      // 10
    fmt.Println(v.Load()) // 20

    // CAS
    if v.CompareAndSwap(20, 30) {
        fmt.Println(v.Load()) // 30
    }
}
```

### SyncMap[K, V]
```go
package main

import (
    "fmt"
    "github.com/kkqy-go/generic"
)

func main() {
    var m generic.SyncMap[string, int]

    // 基本存取
    m.Store("a", 1)
    if v, ok := m.Load("a"); ok {
        fmt.Println(v) // 1
    }

    // LoadOrStore
    actual, loaded := m.LoadOrStore("b", 2)
    fmt.Println(actual, loaded) // 2 false（第一次）

    // 遍历
    m.Range(func(k string, v int) bool {
        fmt.Println(k, v)
        return true
    })

    // 长度与清空
    fmt.Println(m.Len()) // 2
    m.Clear()
    fmt.Println(m.Len()) // 0
}
```

## 包文档
- https://pkg.go.dev/github.com/kkqy-go/generic

## 设计动机
- 利用 Go 泛型，让常用并发/容器类型更安全、更易用。
- 贴近标准库语义，零依赖，小而简单。

## 开发
```bash
go test ./...
```

欢迎提交 Issue 与 PR 改进本库。
