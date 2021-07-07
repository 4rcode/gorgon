/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core_test

import (
	"testing"

	"github.com/4rcode/gadget/core"
)

func TestDefaultPrinter(t *testing.T) {
	for _, td := range []struct {
		in  interface{}
		out string
	}{
		{nil, "<nil>"},
		{true, "true"},
		{12.3, "12.3"},
		{"abc", "\"abc\""},
		{func() {}, "github.com/4rcode/gadget/core_test.TestDefaultPrinter.func1"},
	} {
		text := core.DefaultPrinter(td.in)

		if text != td.out {
			t.Errorf("\ngot  %#v\nwant %#v", text, td.out)
		}
	}
}
