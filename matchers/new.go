/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

import (
	"fmt"
)

// New TODO
func New(
	f Func, args ...interface{},
) DescribedFunc {
	return DescribedFunc{
		f, fmt.Sprint(args...)}
}

// NewFormatted TODO
func NewFormatted(
	f Func, format string, args ...interface{},
) DescribedFunc {
	return DescribedFunc{
		f, fmt.Sprintf(format, args...)}
}
