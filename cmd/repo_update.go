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
	"standardized/internal"
)

var updateCmd = &cobra.Command{
	Use:   "update [REPO NAME]",
	Short: "Update objects definitions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println("Updating repo: " + args[0])
			r := tools.ObjRepo{Name: args[0]}
			r.Load()
			r.Update()
		} else {
			config_dir := tools.GetConfigDir()

			files, err := ioutil.ReadDir(config_dir)
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {
				mode := f.Mode()
				if mode.IsDir() && f.Name()[:1] != "." {
					fmt.Println("Updating repo: " + f.Name())
					r := tools.ObjRepo{Name: f.Name()}
					r.Load()
					r.Update()
				}
			}
		}
	},
}

func init() {
	repoCmd.AddCommand(updateCmd)
}
