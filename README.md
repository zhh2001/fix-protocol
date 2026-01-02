# FIX Protocol

[![Go Report Card](https://goreportcard.com/badge/github.com/zhh2001/fix-protocol)](https://goreportcard.com/report/github.com/zhh2001/fix-protocol)

A Go package for mapping FIX protocol field tags to semantic constants. This package defines FIX protocol fields (such as Account, BeginString, CheckSum, ClOrdID, etc.) as Go constants using iota, avoiding hardcoded numeric values in FIX protocol development. Each constant is accompanied by detailed comments explaining the corresponding FIX field's meaning, usage scenarios, and protocol constraints, which helps to improve code readability, reduce errors, and simplify the maintenance of FIX-related projects (such as trading systems, financial message parsing systems).

## Core Features

1. **Semantic Field Tags**: Defined in `tag.go`, FIX protocol field tags are mapped to named constants via `iota` (e.g., `TagAccount`, `TagBeginString`, `TagCheckSum`), eliminating magic numbers in business code.
2. **Business Enumeration Constants**: Defined in `enum.go`, common FIX protocol enumeration values are encapsulated as named string constants (e.g., `AdvSideBuy`/`AdvSideSell` for advertisement sides, `ExecInstNotHeld`/`ExecInstAON` for execution instructions), avoiding hardcoded string literals and reducing typos.
3. **Detailed Documentation**: Each constant is annotated with clear explanations of its FIX protocol meaning, usage scenarios, and constraints, improving team collaboration efficiency and reducing the cost of understanding the FIX protocol.

## Usage Examples

### Basic Usage: Field Tag Constants

```go
package main

import (
	"fmt"

	"github.com/zhh2001/fix-protocol/v42/fields"
)

func main() {
	// Print the numeric value of the FIX BeginString field tag
	fmt.Println(fields.TagBeginString) // 8

	// Print the numeric value of the FIX Account field tag
	fmt.Println(fields.TagAccount) // 1
}

```

### Advanced Usage: Business Enumeration Constants

```go
package main

import (
	"fmt"

	"github.com/zhh2001/fix-protocol/v42/fields"
)

func main() {
	// FIX protocol message version string
	fmt.Println("FIX Protocol Version:", fields.BeginStringFIX42) // FIX.4.2

	// Advertisement side enumeration
	fmt.Println("Adv Side - Buy:", fields.AdvSideBuy)   // B
	fmt.Println("Adv Side - Sell:", fields.AdvSideSell) // S

	// Execution instruction enumeration
	fmt.Println("Exec Instruction - Not Held:", fields.ExecInstNotHeld)  // 1
	fmt.Println("Exec Instruction - All or None:", fields.ExecInstAON)   // G
	fmt.Println("Exec Instruction - Do Not Reduce:", fields.ExecInstDNR) // F
}
```

## Core File Explanation

| File      | Responsibility                                                                                                                                 |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `tag.go`  | Defines FIX protocol field tags as `iota` constants (numeric values), with detailed comments for each field.                                   |
| `enum.go` | Defines common FIX protocol enumeration values as string constants, covering transaction sides, execution instructions, commission types, etc. |

## Advantages

- **Improve Code Readability**: Semantic constants replace hardcoded numbers/strings, making the purpose of FIX-related code self-explanatory.
- **Reduce Development Errors**: Avoid typos and incorrect value usage caused by manual hardcoding.
- **Simplify Project Maintenance**: Centralized management of FIX protocol constants, easy to update and iterate with protocol versions.
- **Lower Learning Costs**: Detailed annotations help developers who are new to the FIX protocol quickly grasp the meaning and usage of each field.

## Notes

1. This package currently targets the **FIX.4.2** protocol version, with the import path suffix /v42 corresponding to the protocol version.
2. All constants are defined as unexported (lowercase) or exported (uppercase) according to Go's visibility rules, and exported constants can be directly used in external projects.
3. The package only provides constant mappings and does not include FIX message parsing, serialization, or other business logic, which can be extended based on actual project needs.
