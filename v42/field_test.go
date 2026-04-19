// Copyright 2025 Henghua Zhang. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fields

import (
	"testing"
)

// TestTagNameRoundTrip asserts that every tag declared in tag.go resolves to a
// name via Tag.Name and that the same name resolves back to the same tag
// through TagByName. This is the core invariant of the lookup tables: if it
// breaks, every caller that reflects on FIX fields (parsers, loggers,
// validators) also breaks.
func TestTagNameRoundTrip(t *testing.T) {
	if len(tagNames) == 0 {
		t.Fatal("tagNames is empty; the generator must run before tests")
	}
	if len(tagNames) != len(tagsByName) {
		t.Fatalf("tagNames (%d) and tagsByName (%d) have different cardinality", len(tagNames), len(tagsByName))
	}
	for tag, name := range tagNames {
		got, ok := TagByName(name)
		if !ok {
			t.Errorf("TagByName(%q) missing; should resolve to tag %d", name, tag)
			continue
		}
		if got != tag {
			t.Errorf("TagByName(%q) = %d, want %d", name, got, tag)
		}
		if tag.Name() != name {
			t.Errorf("Tag(%d).Name() = %q, want %q", tag, tag.Name(), name)
		}
	}
}

// TestTagStringKnownAndUnknown covers both the common case (a known FIX tag
// formats as Name(N)) and the fallback for custom / out-of-spec tags, which
// must still produce a useful decimal representation so parsers can log them
// without branching.
func TestTagStringKnownAndUnknown(t *testing.T) {
	known := []struct {
		tag  Tag
		want string
	}{
		{TagAccount, "Account(1)"},
		{TagBeginString, "BeginString(8)"},
		{TagMsgType, "MsgType(35)"},
		{TagEncodedListStatusText, "EncodedListStatusText(446)"},
	}
	for _, tc := range known {
		if got := tc.tag.String(); got != tc.want {
			t.Errorf("Tag(%d).String() = %q, want %q", tc.tag, got, tc.want)
		}
	}

	// A tag number well outside the FIX 4.2 dictionary — 5001 is commonly used
	// for user-defined / vendor-extension tags — must stringify to its bare
	// integer so downstream tools can still surface it.
	unknown := Tag(5001)
	if got := unknown.String(); got != "5001" {
		t.Errorf("Tag(5001).String() = %q, want %q", got, "5001")
	}
	if unknown.IsValid() {
		t.Errorf("Tag(5001).IsValid() = true, want false")
	}
	if unknown.Name() != "" {
		t.Errorf("Tag(5001).Name() = %q, want empty string", unknown.Name())
	}
}

// TestTagByNumber verifies the membership check composed into the lookup. A
// caller that only needs "is this a real tag?" should not have to probe both
// the number and the name map separately.
func TestTagByNumber(t *testing.T) {
	if got, ok := TagByNumber(1); !ok || got != TagAccount {
		t.Errorf("TagByNumber(1) = %d, %v; want TagAccount, true", got, ok)
	}
	if got, ok := TagByNumber(446); !ok || got != TagEncodedListStatusText {
		t.Errorf("TagByNumber(446) = %d, %v; want TagEncodedListStatusText, true", got, ok)
	}
	// Known gap: tag 101 is reserved and deliberately left as a blank in tag.go.
	if _, ok := TagByNumber(101); ok {
		t.Errorf("TagByNumber(101) = _, true; FIX 4.2 leaves tag 101 reserved")
	}
	if _, ok := TagByNumber(0); ok {
		t.Errorf("TagByNumber(0) = _, true; FIX tags start at 1")
	}
}

// TestTagByName_CaseSensitive guards against well-meaning callers that may
// lower-case inputs. FIX field names use specific capitalization (ClOrdID,
// OnBehalfOfCompID, ...) and normalizing loses information that downstream
// consumers need.
func TestTagByName_CaseSensitive(t *testing.T) {
	if _, ok := TagByName("account"); ok {
		t.Error("TagByName lookups must be case-sensitive (got match for \"account\")")
	}
	if got, ok := TagByName("Account"); !ok || got != TagAccount {
		t.Errorf("TagByName(%q) = %d, %v; want TagAccount, true", "Account", got, ok)
	}
}

// TestAllTagsIsACopy protects against accidental mutation of the internal
// lookup table. A caller that receives AllTags() and later clears it must not
// affect subsequent callers.
func TestAllTagsIsACopy(t *testing.T) {
	snap := AllTags()
	if len(snap) != len(tagNames) {
		t.Fatalf("AllTags() returned %d entries; want %d", len(snap), len(tagNames))
	}
	for k := range snap {
		delete(snap, k)
	}
	if len(tagNames) == 0 {
		t.Fatal("mutating AllTags()' result corrupted internal tagNames map")
	}
}

// TestMustTagByNamePanics ensures the must-variant fails loudly rather than
// returning a zero Tag, which is the whole reason for its existence.
func TestMustTagByNamePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustTagByName(bogus) did not panic")
		}
	}()
	_ = MustTagByName("DefinitelyNotAField")
}

// TestTagConstantsMatchGeneratedTable asserts that the hand-written TagXxx
// constants, which are value-verified in tag_test.go, line up with the
// generated name table. If tag.go and field_names.go drift, go generate was
// forgotten — surface that before a release rather than at integration time.
func TestTagConstantsMatchGeneratedTable(t *testing.T) {
	cases := map[string]Tag{
		"Account":               TagAccount,
		"BeginString":           TagBeginString,
		"ClOrdID":               TagClOrdID,
		"MsgType":               TagMsgType,
		"Side":                  TagSide,
		"OrdType":               TagOrdType,
		"TransactTime":          TagTransactTime,
		"ExDestination":         TagExDestination,
		"CxlRejReason":          TagCxlRejReason,
		"CouponRate":            TagCouponRate,
		"ContractMultiplier":    TagContractMultiplier,
		"MDReqID":               TagMDReqID,
		"EncodedListStatusText": TagEncodedListStatusText,
	}
	for name, tag := range cases {
		got, ok := TagByName(name)
		if !ok {
			t.Errorf("generated table missing %q", name)
			continue
		}
		if got != tag {
			t.Errorf("generated table: %q = %d, constant TagXxx = %d", name, got, tag)
		}
	}
}

func BenchmarkTagByNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = TagByNumber(35)
	}
}

func BenchmarkTagByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = TagByName("MsgType")
	}
}

func BenchmarkTagString(b *testing.B) {
	t := Tag(TagMsgType)
	for i := 0; i < b.N; i++ {
		_ = t.String()
	}
}
