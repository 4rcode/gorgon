/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gadget

import (
	"github.com/4rcode/gadget/core"
)

// LenientContext TODO
type LenientContext interface {
	Context
	Logf(string, ...interface{})
}

// WithLenient TODO
func WithLenient(
	context LenientContext,
	options ...core.Option,
) Inspector {
	return factory(options).
		new(context, func() core.Logger {
			return context.Logf
		})
}
