/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

// Not TODO
func Not(
	value interface{},
) Matcher {
	matcher, ok := value.(Matcher)

	if !ok {
		matcher = EqualTo(value)
	}

	return NewFormatted(func(
		provided interface{},
	) bool {
		return !matcher.Matches(provided)
	}, "NOT %v", matcher)
}
