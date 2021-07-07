/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core_test

import (
	"fmt"
	"testing"

	"github.com/4rcode/gadget/core"
	"github.com/4rcode/gadget/mock"
)

func TestOptions(t *testing.T) {
	var context mock.Context
	var fallback mock.Matcher = "fallback"

	inspector := core.NewInspector(
		core.WithLogger(context.Errorf),
		core.WithHelper(context.Helper),
		core.WithFallback(fallback),
		core.WithFormat("got %v want %v"),
		core.WithPrinter(func(v interface{}) string {
			return fmt.Sprintf("[%#v]", v)
		}),
		core.WithSupplier(func(args ...interface{}) core.Matcher {
			return mock.Matcher(fmt.Sprint(args...) + " matcher")
		}), nil)

	inspector(123)

	if context != "#got [123] want fallback matcher" {
		t.Errorf("unexpected %#v", context)
	}
}
