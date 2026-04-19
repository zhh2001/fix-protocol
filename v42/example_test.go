// Copyright 2025 Henghua Zhang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fields_test

import (
	"fmt"

	fields "github.com/zhh2001/fix-protocol/v42"
)

// ExampleTag_String shows how a Tag formats when printed. Known fields
// include both their canonical FIX name and number; unknown fields fall back
// to bare decimal.
func ExampleTag_String() {
	fmt.Println(fields.Tag(fields.TagMsgType))
	fmt.Println(fields.Tag(5001)) // vendor-extension / unknown tag
	// Output:
	// MsgType(35)
	// 5001
}

// ExampleTagByName demonstrates the reverse lookup a FIX parser performs
// when it encounters a field by name (for example, while consuming a
// human-authored config file).
func ExampleTagByName() {
	tag, ok := fields.TagByName("OrderQty")
	fmt.Println(tag, ok)
	// Output: OrderQty(38) true
}

// ExampleTagByNumber demonstrates the validity check a FIX parser performs
// on inbound wire data, distinguishing dictionary tags from user-defined
// extensions (tag numbers ≥ 5000 are commonly vendor-specific).
func ExampleTagByNumber() {
	if _, ok := fields.TagByNumber(35); ok {
		fmt.Println("35 is a known FIX 4.2 field")
	}
	if _, ok := fields.TagByNumber(5001); !ok {
		fmt.Println("5001 is outside the FIX 4.2 dictionary")
	}
	// Output:
	// 35 is a known FIX 4.2 field
	// 5001 is outside the FIX 4.2 dictionary
}
