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

func TestNewInspector(t *testing.T) {
	if core.NewInspector() == nil {
		t.Errorf("cannot be nil")
	}

	if core.NewInspector(nil) == nil {
		t.Errorf("cannot be nil")
	}
}
