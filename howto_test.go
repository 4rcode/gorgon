/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gadget_test

import (
	"github.com/4rcode/gadget"
)

func Example_booleanAssertion() {
	is := gadget.With(t)

	result := 1 + 2

	is(result > 10)
	// Output:
	// provided: false
	// expected: true
}

func Example_equalityAssertion() {
	is := gadget.With(t)

	result := 1 + 2

	is(result, 10)
	// Output:
	// provided: 3
	// expected: 10
}
