/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gadget

import (
	"github.com/4rcode/gadget/core"
)

// ExclusiveContext TODO
type ExclusiveContext interface {
	Context
	Skipf(string, ...interface{})
}

// WithExclusive TODO
func WithExclusive(
	context ExclusiveContext,
	options ...core.Option,
) Inspector {
	return factory(options).
		new(context, func() core.Logger {
			return context.Skipf
		})
}
