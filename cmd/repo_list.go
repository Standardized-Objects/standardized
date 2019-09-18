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
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"path/filepath"
	"standardized/internal"
)

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available repositories",
	Run: func(cmd *cobra.Command, args []string) {
		config_dir := tools.GetConfigDir()

		files, err := ioutil.ReadDir(config_dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			viper.SetConfigFile(filepath.Join(config_dir, f.Name()) + "/auth.yaml")
			err := viper.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
			fmt.Println(f.Name() + " : " + viper.GetString("url"))
		}
	},
}

func init() {
	repoCmd.AddCommand(repoListCmd)
}
