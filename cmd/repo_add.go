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
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"standardized/internal"
)

var sshAuth bool
var sshKey string
var githubToken string

var addCmd = &cobra.Command{
	Use:   "add [NAME] [GIT URL]",
	Short: "Add Standardized Objects Definitions repositories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid arguments")
			os.Exit(0)
		}

		if args[0] == "_local" {
			fmt.Println("Reserved string: _local")
			os.Exit(0)
		}

		r := tools.ObjRepo{
			Name: args[0],
			Url:  args[1],
		}

		home, _ := homedir.Dir()
		if sshAuth {
			if sshKey == "" {
				sshKey = home + "/.ssh/id_rsa"
			}
			r.AuthType = "ssh"
			r.AuthValue = sshKey
		} else if githubToken != "" {
			r.AuthType = "github"
			r.AuthValue = githubToken
		} else {
			r.AuthType = "public"
			r.AuthValue = ""
		}

		r.Start()
		r.Clone()
	},
}

func init() {
	repoCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&sshAuth, "ssh", "s", false, "Use SSH for repo auth")
	addCmd.Flags().StringVarP(&sshKey, "key", "k", "", "SSH private key")
	addCmd.Flags().StringVarP(&githubToken, "token", "t", "", "Use GitHub Personal Access Token for repo auth")
}
