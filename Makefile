# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

file = .coverprofile.test

.PHONY: clean cover all

all: clean cover

clean:
	rm -f $(file)

cover: $(file)
	go tool cover -func=$(file)

$(file):
	go test -cover -coverprofile=$(file) ./...
