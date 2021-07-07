/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core

type inspector struct {
	fallback Matcher
	format   string
	helper   Helper
	logger   Logger
	printer  Printer
	supplier Supplier
}

func (i *inspector) inspect(
	value interface{}, args ...interface{},
) bool {
	if len(args) < 1 {
		args = append(args, i.fallback)
	}

	matcher := i.supplier(args...)

	if matcher.Matches(value) {
		return true
	}

	i.helper()
	i.logger(
		i.format, i.printer(value), matcher)

	return false
}
