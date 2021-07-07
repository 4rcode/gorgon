/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core_test

import (
	"testing"

	"github.com/4rcode/gadget/core"
	"github.com/4rcode/gadget/mock"
)

func TestInspector(t *testing.T) {
	var context mock.Context
	var fallback mock.Matcher

	inspector := core.NewInspector(
		core.WithLogger(context.Errorf),
		core.WithHelper(context.Helper),
		core.WithFallback(&fallback),
		core.WithFormat("got %v want %v"))

	if !inspector("") {
		t.Errorf("unexpected false")
	}

	if context != "" {
		t.Errorf("unexpected %#v", context)
	}

	context = ""
	fallback = "not <nil>"

	if inspector(nil) {
		t.Errorf("unexpected true")
	}

	if context != "#got <nil> want not <nil>" {
		t.Errorf("unexpected %#v", context)
	}
}

func TestInspectorWithMatchers(t *testing.T) {
	var context mock.Context
	var fallback mock.Matcher = "right"
	var matcher mock.Matcher = "a good value"

	inspector := core.NewInspector(
		core.WithLogger(context.Errorf),
		core.WithHelper(context.Helper),
		core.WithFallback(fallback),
		core.WithFormat("got %v want %v"))

	for _, td := range []struct {
		provided interface{}
		expected string
	}{
		{"err", "#got \"err\" want a good value"},
		{struct{}{}, "#got struct {}{} want a good value"},
	} {
		for id, is := range []core.Inspector{
			inspector, inspector, inspector,
		} {
			context = ""

			if is(td.provided, matcher) {
				t.Errorf("the return value should be false")
			}

			if string(context) != td.expected {
				t.Errorf("id: %d\n\ngot  %#v\nwant %#v\n\n", id, context, td.expected)
			}
		}
	}
}
