/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gadget_test

import (
	"github.com/4rcode/gadget"
	"github.com/4rcode/gadget/dummy"
)

var t dummy.Context

func ExampleWith() {
	is := gadget.With(t)

	is(1 == 1.1)
	// Output:
	// provided: false
	// expected: true
}

func ExampleWithExclusive() {
	is := gadget.WithExclusive(t)

	is(0, 1)
	// Output:
	// provided: 0
	// expected: 1
}

func ExampleWithLenient() {
	is := gadget.WithLenient(t)

	is(1 > 2)
	// Output:
	// provided: false
	// expected: true
}

func ExampleWithStrict() {
	is := gadget.WithStrict(t)

	is(2, 2.1)
	// Output:
	// provided: 2
	// expected: 2.1
}
