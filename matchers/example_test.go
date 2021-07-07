/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers_test

import (
	"fmt"
	"io"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/dummy"
	. "github.com/4rcode/gadget/matchers"
)

var t dummy.Context

func ExampleAllOf() {
	is := gadget.With(t)

	is(0, AllOf(EqualTo(0), EqualTo(1)))
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleTrueFor() {
	is := gadget.With(t)

	is(0 >= 1, TrueFor(0, 1))
	// Output:
	// provided: false
	// expected: 0 <=> 1
}

func ExampleTrueFor_withComparison() {
	is := gadget.With(t)

	is(0 >= 1, TrueFor(0, 1, ">="))
	// Output:
	// provided: false
	// expected: 0 >= 1
}

func ExampleDeepEqualTo() {
	is := gadget.With(t)

	is(0, DeepEqualTo(1))
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleEqualTo_true() {
	is := gadget.With(t)

	is(false, EqualTo(true))
	// Output:
	// provided: false
	// expected: true
}

func ExampleEqualTo_one() {
	is := gadget.With(t)

	is(0, EqualTo(1))
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleEqualTo_struct() {
	is := gadget.With(t)

	a := struct{ int }{0}
	b := struct{ int64 }{0}

	is(a, EqualTo(b))
	// Output:
	// provided: struct { int }{int:0}
	// expected: struct { int64 }{int64:0}
}

func ExampleEqualTo_string() {
	is := gadget.With(t)

	is("something", EqualTo("else"))
	// Output:
	// provided: "something"
	// expected: "else"
}

func ExampleNil() {
	is := gadget.With(t)

	is(123, Nil())

	// Output:
	// provided: 123
	// expected: <nil>
}

func ExampleNot_nil() {
	is := gadget.With(t)

	is((*struct{})(nil), Not(Nil()))
	// Output:
	// provided: (*struct {})(nil)
	// expected: NOT <nil>
}

func ExamplePanicking() {
	is := gadget.With(t)

	is(func() {
		// missing panic
	}, Panicking())

	// Output:
	// provided: github.com/4rcode/gadget/matchers_test.ExamplePanicking.func1
	// expected: PANIC
}

func ExamplePanicking_usePanicValue() {
	is := gadget.With(t)

	var v interface{}

	is(func() {
		panic("xyz")
	}, Panicking(&v))

	fmt.Println(v)

	// Output:
	// xyz
}

func ExampleError() {
	is := gadget.With(t)

	is(io.ErrUnexpectedEOF, Error(io.EOF))

	// Output:
	// provided: &errors.errorString{s:"unexpected EOF"}
	// expected: &errors.errorString{s:"EOF"}
}
