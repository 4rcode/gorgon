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
	"github.com/4rcode/gadget/mock"
)

func TestNoneOf(t *testing.T) {
	for i, tc := range []struct {
		matchers.Matcher
		bool
		string
	}{
		{matchers.NoneOf(), true, "NOT true"},
		{matchers.NoneOf(""), true, "NOT \"\""},
		{matchers.NoneOf("value"), false, "NOT \"value\""},
		{matchers.NoneOf("different", "value"), false, "NOT \"value\""},
		{matchers.NoneOf(mock.Matcher("different"), mock.Matcher("value")), false, "NOT value"},
		{matchers.NoneOf(mock.Matcher("different"), mock.Matcher("again")), true, "NOT again"},
	} {
		t.Logf("tc %v", i+1)

		if tc.Matcher.Matches("value") != tc.bool {
			t.Errorf("unexpected %v", !tc.bool)
		}

		if text := fmt.Sprintf("%v", tc.Matcher); text != tc.string {
			t.Errorf("unexpected %v", text)
		}
	}
}
