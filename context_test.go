/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gadget_test

import (
	"testing"

	"github.com/4rcode/gadget"
)

func TestWithFunctions(t *testing.T) {
	for tdId, td := range []struct {
		context testing.TB
	}{
		{nil},
		{t},
	} {
		t.Logf("tdId=%v", tdId+1)

		for tcId, tc := range []struct {
			inspector gadget.Inspector
		}{
			{gadget.With(td.context)},
			{gadget.WithExclusive(td.context)},
			{gadget.WithLenient(td.context)},
			{gadget.WithStrict(td.context)},
		} {
			t.Logf("tcId=%v", tcId+1)

			if tc.inspector == nil {
				t.Errorf("value %v cannot be nil", tcId)
			}
		}
	}
}
