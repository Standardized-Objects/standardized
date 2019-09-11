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
  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
  "path/filepath"
  "os"
  "log"
)

var addCmd = &cobra.Command{
  Use:   "add [NAME] [GIT URL] [GITHUB TOKEN]",
  Short: "Add Standardized Objects Definitions repositories",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 3 {
      fmt.Println("Invalid arguments")
      os.Exit(1)
    }

    _, err := git.PlainClone(filepath.Join(common.GetConfigDir(), args[0]), false, &git.CloneOptions{
      // The intended use of a GitHub personal access token is in replace of your password
      // because access tokens can easily be revoked.
      // https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
      Auth: &http.BasicAuth{
        Username: "standardized", // yes, this can be anything except an empty string
        Password: args[2],
      },
      URL:      args[1],
      Progress: os.Stdout,
    })

    if err != nil {
      log.Fatal(err)
    }
  },
}

func init() {
  repoCmd.AddCommand(addCmd)
}
