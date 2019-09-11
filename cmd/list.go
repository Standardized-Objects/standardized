/*
Copyright © 2019 Juan Ezquerro LLanes <arrase@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
  "fmt"
  "standardized/internal"
  "github.com/spf13/cobra"
  "io/ioutil"
  "log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List available repositories",
  Run: func(cmd *cobra.Command, args []string) {
    files, err := ioutil.ReadDir(common.GetConfigDir())
    if err != nil {
      log.Fatal(err)
    }

    for _, f := range files {
      fmt.Println(f.Name())
    }
  },
}

func init() {
  repoCmd.AddCommand(listCmd)
}
