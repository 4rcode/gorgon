/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package mock_test

import (
	"fmt"

	"github.com/4rcode/gadget/mock"
)

func ExampleMatcher() {
	var matcher mock.Matcher

	fmt.Println(matcher)
	fmt.Println(&matcher)

	fmt.Println(
		matcher.Matches(""))

	matcher = "wrong"

	fmt.Println(matcher)
	fmt.Println(&matcher)

	fmt.Println(
		matcher.Matches(""))
	// Output:
	// true
	// wrong
	// wrong
	// false
}
