# FIX Protocol

A Go package for mapping FIX protocol field tags to semantic constants. This package defines FIX protocol fields (such as Account, BeginString, CheckSum, ClOrdID, etc.) as Go constants using iota, avoiding hardcoded numeric values in FIX protocol development. Each constant is accompanied by detailed comments explaining the corresponding FIX field's meaning, usage scenarios, and protocol constraints, which helps to improve code readability, reduce errors, and simplify the maintenance of FIX-related projects (such as trading systems, financial message parsing systems).

```go
package main

import (
	"fmt"

	"github.com/zhh2001/fix-protocol/v42"
)

func main() {
	fmt.Println(fields.TagBeginString) // 8
}

```
