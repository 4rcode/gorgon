/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core

import (
	"github.com/4rcode/gadget/matchers"
)

// Inspector TODO
type Inspector func(interface{}, ...interface{}) bool

// NewInspector TODO
func NewInspector(
	options ...Option,
) Inspector {
	i := inspector{format: DefaultFormat}

	for _, option := range options {
		if option != nil {
			option(&i)
		}
	}

	if i.fallback == nil {
		i.fallback = matchers.Func(nil)
	}

	if i.helper == nil {
		i.helper = DefaultHelper
	}

	if i.logger == nil {
		i.logger = DefaultLogger
	}

	if i.printer == nil {
		i.printer = DefaultPrinter
	}

	if i.supplier == nil {
		i.supplier = DefaultSupplier
	}

	return i.inspect
}
