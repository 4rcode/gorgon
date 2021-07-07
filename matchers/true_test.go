/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers_test

import (
	"fmt"
	"testing"

	"github.com/4rcode/gadget/matchers"
)

func TestTrue(t *testing.T) {
	for _, tc := range []struct {
		matchers.Matcher
		string
	}{
		{matchers.True(), "true"},
		{matchers.True("a", "b", "c"), "abc"},
	} {
		if text := fmt.Sprintf("%v", tc.Matcher); text != tc.string {
			t.Errorf("unexpected %#v", text)
		}

		for _, td := range []struct {
			in  interface{}
			out bool
		}{
			{nil, false},
			{123, false},
			{"abc", false},
			{false, false},
			{true, true},
		} {
			if tc.Matcher.Matches(td.in) != td.out {
				t.Errorf("%v should give %#v", td.in, td.out)
			}
		}
	}
}
