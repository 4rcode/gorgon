/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package mock

import (
	"fmt"
)

type Context string

func (c *Context) Errorf(f string, a ...interface{}) {
	c.Logf(f, a...)
}

func (c *Context) Fatalf(f string, a ...interface{}) {
	c.Logf(f, a...)
}

func (c *Context) Helper() {
	c.Logf("#")
}

func (c *Context) Logf(f string, a ...interface{}) {
	*c = *c + Context(fmt.Sprintf(f, a...))
}

func (c *Context) Skipf(f string, a ...interface{}) {
	c.Logf(f, a...)
}

func (c Context) String() string {
	return string(c)
}
