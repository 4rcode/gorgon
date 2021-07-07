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

func TestNot(t *testing.T) {
	m := matchers.Not(mock.Matcher(""))

	if m.Matches("") {
		t.Errorf("unexpected true")
	}

	if text := fmt.Sprintf("%v", m); text != "NOT " {
		t.Errorf("unexpected %#v", text)
	}

	m = matchers.Not(mock.Matcher("error"))

	if !m.Matches("false") {
		t.Errorf("unexpected false")
	}

	if text := fmt.Sprintf("%v", m); text != "NOT error" {
		t.Errorf("unexpected %#v", text)
	}

	matchers.Not(nil)

	// t.Errorf("panic expected") TODO test
}
