/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/4rcode/gadget/matchers"
)

// DefaultFormat TODO
const DefaultFormat = "\n\nprovided: %v\nexpected: %v\n\n"

// DefaultHelper TODO
func DefaultHelper() {}

// DefaultLogger TODO
func DefaultLogger(string, ...interface{}) {}

// DefaultPrinter TODO
func DefaultPrinter(value interface{}) string {
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Func {
		return runtime.FuncForPC(
			v.Pointer(),
		).Name()
	}

	return fmt.Sprintf("%#v", value)
}

// DefaultSupplier TODO
func DefaultSupplier(args ...interface{}) Matcher {
	return matchers.AllOf(args...)
}
