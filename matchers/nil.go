/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

import (
	"reflect"
)

// Nil TODO
func Nil() Matcher {
	return NewFormatted(func(
		provided interface{},
	) bool {
		if provided == nil {
			return true
		}

		value := reflect.ValueOf(provided)
		kind := value.Kind()

		if reflect.Chan <= kind &&
			kind <= reflect.Slice &&
			value.IsNil() {
			return true
		}

		return false
	}, "%#v", nil)
}
