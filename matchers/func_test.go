/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers_test

import (
	"testing"

	"github.com/4rcode/gadget/matchers"
)

func TestFunc(t *testing.T) {
	var matcher matchers.Func = func(interface{}) bool {
		return true
	}

	if !matcher.Matches(nil) {
		t.Errorf("unexpected false")
	}

	if matcher.String() != "a matching value" {
		t.Errorf("unexpected %#v", matcher)
	}
}
