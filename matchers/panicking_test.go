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

func TestPanicking(t *testing.T) {
	var receivers [3]interface{}

	pointers := []*interface{}{
		nil, &receivers[0],
		&receivers[1], &receivers[2]}

	matcher := matchers.Panicking(pointers...)
	text := fmt.Sprintf("%v", matcher)

	if text != "PANIC" {
		t.Errorf("unexpected %#v", text)
	}

	if matcher.Matches(nil) {
		t.Errorf("Panicking should be false")
	}

	for _, expected := range []interface{}{
		(*int)(nil), 0, "err", nil,
	} {
		if matcher.Matches(func() {
			if expected != nil {
				panic(expected)
			}
		}) != (expected != nil) {
			t.Errorf("Panicking should be %#v", expected != nil)
		}

		if receivers != [3]interface{}{
			expected, expected, expected,
		} {
			t.Errorf("unexpected values")
		}
	}
}
