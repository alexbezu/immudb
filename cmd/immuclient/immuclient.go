/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"strings"

	c "github.com/codenotary/immudb/cmd/helper"
	immuclient "github.com/codenotary/immudb/cmd/immuclient/command"
)

func main() {
	cmd := immuclient.NewCommand()
	err := immuclient.Execute(cmd)
	if err != nil {
		fmt.Println(cmd.Aliases)
		if strings.HasPrefix(err.Error(), "unknown command") {
			os.Exit(0)
		}
		c.QuitWithUserError(err)
	}
}
