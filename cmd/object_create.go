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
	"github.com/spf13/cobra"
	"os"
	"github.com/Standardized-Objects/standardized/tools"
)

var outputPath string

var createCmd = &cobra.Command{
	Use:   "create [OBJECT]",
	Short: "Cenerate an object from Object Definition",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid arguments")
			os.Exit(0)
		}

		// Load object definition
		o := tools.ObjDef{
			Name:   args[0],
			OutDir: outputPath,
		}
		o.Load()
		// Run pre create hooks
		o.RunHooks("precreate")
		// Apply values
		o.Apply()
		// Run post create hooks
		o.RunHooks("postcreate")
	},
}

func init() {
	objectCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&outputPath, "output", "o", "_current_dir_", "Output path, will be created if not exists")
}
