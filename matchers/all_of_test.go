/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers_test

import (
	"testing"

	"github.com/4rcode/gadget/matchers"
	"github.com/4rcode/gadget/mock"
)

func TestAllOf(t *testing.T) {
	matcher := matchers.AllOf()

	if outcome := matcher.Matches(""); outcome {
		t.Errorf("unexpected %#v", outcome)
	}

	matcher = matchers.AllOf(
		mock.Matcher("that's"), mock.Matcher("wrong"))

	if outcome := matcher.Matches(""); outcome {
		t.Errorf("unexpected %#v", outcome)
	}

	matcher = matchers.AllOf(
		mock.Matcher(""), mock.Matcher("wrong"))

	if outcome := matcher.Matches(""); outcome {
		t.Errorf("unexpected %#v", outcome)
	}

	matcher = matchers.AllOf(
		mock.Matcher(""), mock.Matcher(""))

	if outcome := matcher.Matches(""); !outcome {
		t.Errorf("unexpected %#v", outcome)
	}
}
