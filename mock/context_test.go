/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package mock_test

import (
	"fmt"

	"github.com/4rcode/gadget/mock"
)

func ExampleContext() {
	var t mock.Context

	t.Helper()
	t.Errorf("%s", "a")
	t.Fatalf("%d", 1)
	t.Logf("%.2f", 2.3)
	t.Skipf("%#v", struct{}{})

	fmt.Println(t)
	// Output:
	// #a12.30struct {}{}
}
