/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gadget

import (
	"github.com/4rcode/gadget/core"
)

// Context TODO
type Context interface {
	Helper()
}

// Inspector TODO
type Inspector func(interface{}, ...interface{}) bool

type factory []core.Option

func (f factory) new(
	helper Context,
	logger func() core.Logger,
) Inspector {
	if helper != nil {
		f = append(f,
			core.WithHelper(helper.Helper),
			core.WithLogger(logger()))
	}

	return Inspector(core.NewInspector(f...))
}
