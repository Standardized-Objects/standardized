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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"standardized/internal"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available objects",
	Run: func(cmd *cobra.Command, args []string) {
		// From saved git repos
		config_dir := tools.GetConfigDir()

		files, err := ioutil.ReadDir(config_dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			objs, _ := ioutil.ReadDir(filepath.Join(config_dir, f.Name(), "src"))
			for _, o := range objs {
				mode := o.Mode()
				if mode.IsDir() && o.Name()[:1] != "." {
					fmt.Println(filepath.Join(f.Name(), o.Name()))
				}
			}
		}

		// From current working dir
		curr_dir, _ := os.Getwd()
		local_objs := filepath.Join(curr_dir, ".stdized")
		if tools.Exists(local_objs) {
			lobjs, _ := ioutil.ReadDir(local_objs)
			for _, lo := range lobjs {
				lmode := lo.Mode()
				if lmode.IsDir() && lo.Name()[:1] != "." {
					fmt.Println(filepath.Join("_local", lo.Name()))
				}
			}
		}
	},
}

func init() {
	objectCmd.AddCommand(listCmd)
}
