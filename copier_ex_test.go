package copier_test

import (
	"testing"

	"github.com/admpub/copier"
	"github.com/webx-top/com"
)

// https://github.com/jinzhu/copier/issues/31
func TestNested(t *testing.T) {
	type Nested struct {
		A string
	}
	type ParentA struct {
		*Nested
	}
	type parentB struct {
		*Nested
	}
	type parentC struct {
		*ParentA
	}
	a := ParentA{
		Nested: &Nested{A: "a"},
	}
	b := parentB{}
	copier.Copy(&b, &a)
	if b.A != a.A {
		panic(`no match`)
	}

	com.Dump(b)

	a1 := parentC{
		ParentA: &a,
	}
	b1 := parentC{}

	copier.Copy(&b1, &a1)
	com.Dump(b1)
	if b1.A != a1.A {
		panic(`no match`)
	}
}
