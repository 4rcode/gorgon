/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package core

// Helper TODO
type Helper func()

// Logger TODO
type Logger func(string, ...interface{})

// Matcher TODO
type Matcher interface {
	Matches(interface{}) bool
}

// Option is a function used to configure an inspector.
type Option func(i *inspector)

// Printer TODO
type Printer func(interface{}) string

// Supplier TODO
type Supplier func(...interface{}) Matcher

// WithFallback configures the message reported by the fallback matcher in case
// of failure.
func WithFallback(matcher Matcher) Option {
	return func(i *inspector) {
		i.fallback = matcher
	}
}

// WithFormat configures the format string used by an Inspector to report a
// message, in case of failure. The string should have a reference to both the
// provided and the expected values, usually using the placeholder "%v".
func WithFormat(format string) Option {
	return func(i *inspector) {
		i.format = format
	}
}

// WithHelper TODO
func WithHelper(helper Helper) Option {
	return func(i *inspector) {
		i.helper = helper
	}
}

// WithLogger TODO
func WithLogger(logger Logger) Option {
	return func(i *inspector) {
		i.logger = logger
	}
}

// WithPrinter configures the printer function for the current value under
// investigation, reported only in case of failure.
func WithPrinter(printer Printer) Option {
	return func(i *inspector) {
		i.printer = printer
	}
}

// WithSupplier TODO
func WithSupplier(supplier Supplier) Option {
	return func(i *inspector) {
		i.supplier = supplier
	}
}
