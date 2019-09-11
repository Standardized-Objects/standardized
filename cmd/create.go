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
  "os"
  "github.com/spf13/cobra"
  "standardized/internal"
)

var createCmd = &cobra.Command{
  Use:   "create <NAME>",
  Short: "Create a new Object Definition with the given name",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
      fmt.Println("Invalid arguments")
      os.Exit(1)
    }

    os.MkdirAll(args[0] + "/templates", os.ModePerm)

    s, _ := os.Create(args[0] + "/spec.yaml")
    s.Write([]byte("name: {{.objName}}\ndescription: Custom Object Definition\ncontent: []\n"))
    s.Close()

    values := map[string]string{
      "objName": args[0],
    }

    common.ParseTemplate(args[0] + "/spec.yaml", values)

    r, _ := os.Create(args[0] + "/templates/README.txt")
    r.Write([]byte("Add your template files here\n"))
    r.Close()

    fmt.Println("New Object Definition: " + args[0])
  },
}

func init() {
  objectCmd.AddCommand(createCmd)
}
