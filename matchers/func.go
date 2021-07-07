/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

// Func TODO
type Func func(interface{}) bool

// Matches TODO
func (f Func) Matches(value interface{}) bool {
	if f == nil {
		return value == true
	}

	return f(value)
}

// String TODO
func (f Func) String() string {
	if f == nil {
		return "true"
	}

	return "a matching value"
}
