/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core_test

import (
	"fmt"

	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/core"
	"github.com/4rcode/gadget/dummy"
	"github.com/4rcode/gadget/matchers"
)

var t dummy.Context

func ExampleWithFallback() {
	is := gadget.With(
		t, core.WithFallback(
			matchers.True("xyz")))

	is(0 >= 1)
	// Output:
	// provided: false
	// expected: xyz
}

func ExampleWithFormat() {
	is := gadget.With(
		t, core.WithFormat("got %v want %v"))

	is(0 >= 1)
	// Output:
	// got false want true
}

func ExampleWithPrinter() {
	is := gadget.With(
		t, core.WithPrinter(
			func(v interface{}) string {
				return fmt.Sprintf(">>>  %#v  <<<", v)
			}))

	is(0 >= 1)
	// Output:
	// provided: >>>  false  <<<
	// expected: true
}
