/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package mock

type Matcher string

func (m Matcher) Matches(value interface{}) bool {
	return string(m) == value
}

func (m Matcher) String() string {
	return string(m)
}
