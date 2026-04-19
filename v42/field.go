// Copyright 2025 Henghua Zhang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fields

//go:generate go run ./internal/gen -in tag.go -out field_names.go -pkg fields

import (
	"fmt"
	"strconv"
)

// Tag is the numeric identifier of a FIX 4.2 field. The untyped TagXxx
// integer constants in tag.go assign implicitly to this type, so callers may
// write either `fields.Tag(38)` or `fields.TagOrderQty` interchangeably.
//
// Tag intentionally satisfies fmt.Stringer so it can be used directly in
// logging and error paths: an unknown tag falls back to its numeric form so
// parsers can report custom or out-of-spec fields without losing information.
type Tag int

// Int returns the tag's numeric value. It exists purely for readability at
// call sites where a plain `int(t)` conversion would be ambiguous (e.g.
// formatting inside a map literal).
func (t Tag) Int() int { return int(t) }

// Name returns the canonical FIX field name for t (for example, "OrderQty" for
// TagOrderQty). It returns the empty string for tags that are not part of the
// FIX 4.2 dictionary shipped in this package.
func (t Tag) Name() string { return tagNames[t] }

// IsValid reports whether t corresponds to a known FIX 4.2 field defined in
// this package. User-defined and out-of-spec tags therefore return false even
// if they are numerically valid integers.
func (t Tag) IsValid() bool {
	_, ok := tagNames[t]
	return ok
}

// String implements fmt.Stringer. Known tags format as `Name(N)`; unknown tags
// format as their decimal number so output remains informative for custom or
// future-version fields without allocating via fmt.Sprintf for the common
// known case.
func (t Tag) String() string {
	if name, ok := tagNames[t]; ok {
		return name + "(" + strconv.Itoa(int(t)) + ")"
	}
	return strconv.Itoa(int(t))
}

// TagByName returns the Tag whose canonical FIX name is exactly name. The
// lookup is case-sensitive — FIX field names are specified with specific
// capitalization (ClOrdID, OrderQty, MsgSeqNum, ...) and callers should not
// normalize away that casing. The second return value is false when name is
// not part of the dictionary.
func TagByName(name string) (Tag, bool) {
	t, ok := tagsByName[name]
	return t, ok
}

// TagByNumber returns the Tag for a numeric identifier. It is a thin wrapper
// that also validates membership in the FIX 4.2 dictionary, so callers can
// distinguish known from unknown tag numbers in a single call.
func TagByNumber(n int) (Tag, bool) {
	t := Tag(n)
	_, ok := tagNames[t]
	return t, ok
}

// AllTags returns the set of tags defined by this package keyed by their
// numeric value. The returned map is a fresh copy; callers may mutate it
// freely. It is intended for reflective uses (CLI dumps, test fixtures) and
// not for hot paths — use TagByNumber / TagByName instead.
func AllTags() map[Tag]string {
	out := make(map[Tag]string, len(tagNames))
	for t, name := range tagNames {
		out[t] = name
	}
	return out
}

// MustTagByName is the panicking form of TagByName for use in package-level
// variable initializers where a missing tag indicates a programming error
// rather than a runtime input problem.
func MustTagByName(name string) Tag {
	t, ok := tagsByName[name]
	if !ok {
		panic(fmt.Sprintf("fields: unknown FIX tag name %q", name))
	}
	return t
}
