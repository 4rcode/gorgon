/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package matchers

// Panicking TODO
func Panicking(
	receivers ...*interface{},
) Matcher {
	return New(func(
		provided interface{},
	) (ok bool) {
		ok = true

		defer func() {
			value := recover()

			for _, receiver := range receivers {
				if receiver != nil {
					*receiver = value
				}
			}
		}()

		f, ok := provided.(func())

		if !ok || f == nil {
			panic("func() required")
		}

		f()

		return false
	}, "PANIC")
}
