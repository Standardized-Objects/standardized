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
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"standardized/internal"
)

var repoDeleteCmd = &cobra.Command{
	Use:   "delete [REPO NAME]",
	Short: "Delete repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid arguments")
			os.Exit(0)
		}

		repo_dir := filepath.Join(tools.GetConfigDir(), args[0])

		if tools.Exists(repo_dir) {

			prompt := promptui.Prompt{
				Label:     "Delete repository [" + args[0] + "]",
				IsConfirm: true,
			}
			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Canceled %v\n", err)
				return
			}

			if result == "y" {
				err := os.RemoveAll(repo_dir)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Repo deleted")
				}
			}
		}
	},
}

func init() {
	repoCmd.AddCommand(repoDeleteCmd)
}
