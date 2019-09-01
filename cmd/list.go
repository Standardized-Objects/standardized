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
  "standardized/common"
  "github.com/spf13/cobra"
  "os"
  "io"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
  Use:   "list",
  Short: "List available repositories",
  Run: func(cmd *cobra.Command, args []string) {
    configFile := common.GetConfigDir() + "/repos.yaml"

    _, err := os.Stat(configFile)
    if os.IsNotExist(err) {
      fmt.Println("Add a repository first.")
      os.Exit(0)
    }

    f, _ := os.Open(configFile)
    io.Copy(os.Stdout, f)
    f.Close()
  },
}

func init() {
  repoCmd.AddCommand(listCmd)
}
