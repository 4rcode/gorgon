/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package dummy

import (
	"fmt"
)

type Context struct{}

func (Context) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (Context) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (Context) Helper() {}

func (Context) Logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (Context) Skipf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
