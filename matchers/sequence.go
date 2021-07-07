/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

import (
	"fmt"
)

// Matcher TODO
type Matcher interface {
	Matches(interface{}) bool
}

// Sequence TODO
type Sequence struct {
	values   []interface{}
	expected bool
	inverted bool
	prefix   string
	failed   Matcher
}

func (s *Sequence) Matches(value interface{}) bool {
	if len(s.values) < 1 {
		s.values = append(s.values, Func(nil))
	}

	for _, arg := range s.values {
		matcher, ok := arg.(Matcher)

		if !ok {
			matcher = EqualTo(arg)
		}

		s.failed = matcher

		if matcher.Matches(value) == s.expected {
			return s.result()
		}
	}

	return !s.result()
}

func (s *Sequence) result() bool {
	return (s.expected || s.inverted) &&
		!(s.expected && s.inverted)
}

func (s *Sequence) String() string {
	return fmt.Sprintf("%s%v", s.prefix, s.failed)
}
