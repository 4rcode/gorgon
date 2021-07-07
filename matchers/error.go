/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

import (
	"errors"
)

// Error TODO
func Error(expected error) Matcher {
	return NewFormatted(func(
		provided interface{},
	) bool {
		if provided == expected {
			return true
		}

		err, ok := provided.(error)

		if !ok {
			return false
		}

		return errors.Is(err, expected)
	}, "%#v", expected)
}
