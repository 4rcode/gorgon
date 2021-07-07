/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package dummy_test

import (
	"github.com/4rcode/gadget/dummy"
)

func ExampleContext() {
	var logger dummy.Context

	logger.Helper()
	logger.Errorf("%s", "a")
	logger.Fatalf("%d", 1)
	logger.Logf("%.2f", 2.3)
	logger.Skipf("%#v", struct{}{})
	// Output:
	// a12.30struct {}{}
}
