/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

// EqualTo TODO
func EqualTo(
	expected interface{},
) Matcher {
	return NewFormatted(func(
		provided interface{},
	) bool {
		return provided == expected
	}, "%#v", expected)
}
