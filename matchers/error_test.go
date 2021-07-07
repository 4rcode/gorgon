/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/4rcode/gadget/matchers"
)

func TestError(t *testing.T) {
	type err struct{ error }

	custom := fmt.Errorf("error %w", io.EOF)

	for tcId, tc := range []struct {
		provided interface{}
		expected error
		result   bool
	}{
		{nil, nil, true},
		{nil, io.EOF, false},
		{nil, custom, false},
		{nil, (*err)(nil), false},
		{error(nil), nil, true},
		{error(nil), io.EOF, false},
		{error(nil), custom, false},
		{error(nil), (*err)(nil), false},
		{(*err)(nil), nil, false},
		{(*err)(nil), io.EOF, false},
		{(*err)(nil), custom, false},
		{(*err)(nil), (*err)(nil), true},
		{io.EOF, nil, false},
		{io.EOF, io.EOF, true},
		{io.EOF, custom, false},
		{io.EOF, (*err)(nil), false},
		{io.EOF, io.ErrClosedPipe, false},
		{custom, nil, false},
		{custom, io.EOF, true},
		{custom, custom, true},
		{custom, (*err)(nil), false},
		{custom, io.ErrClosedPipe, false},
	} {
		t.Logf("TC %d", tcId+1)
		matcher := matchers.Error(tc.expected)

		if matcher.Matches(tc.provided) != tc.result {
			t.Errorf("unexpected %#v", tc.provided)
		}
	}
}
