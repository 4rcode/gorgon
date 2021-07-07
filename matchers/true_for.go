/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

import (
	"fmt"
)

// TrueFor TODO
func TrueFor(
	provided, expected interface{},
	args ...interface{},
) Matcher {
	if len(args) < 1 {
		args = append(args, "<=>")
	}

	return True(
		fmt.Sprintf(
			"%#v %v %#v", provided,
			fmt.Sprint(args...), expected))
}
