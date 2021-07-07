/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

// NoneOf TODO
func NoneOf(
	args ...interface{},
) *Sequence {
	return &Sequence{
		values: args, expected: true, inverted: true, prefix: "NOT "}
}
