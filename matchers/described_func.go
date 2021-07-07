/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

// DescribedFunc TODO
type DescribedFunc struct {
	function    Func
	description string
}

// Matches TODO
func (f DescribedFunc) Matches(value interface{}) bool {
	return f.function.Matches(value)
}

// String TODO
func (f DescribedFunc) String() string {
	if len(f.description) > 0 {
		return f.description
	}

	return f.function.String()
}
