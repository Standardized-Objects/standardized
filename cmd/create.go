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
  "standardized/internal"
  "os"
  "strings"
  "path/filepath"
  //"bufio"
)

var outputPath string

var createCmd = &cobra.Command{
  Use:   "create [OBJECT]",
  Short: "Cenerate an object from Object Definition",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
      fmt.Println("Invalid arguments")
      os.Exit(1)
    }

    obj := strings.Split(args[0],"/")

    output_dir , _ := os.Getwd()
    if outputPath != "" {
      output_dir = filepath.Join(output_dir, outputPath)
    }

    obj_dir := filepath.Join(filepath.Join(filepath.Join(tools.GetConfigDir(), obj[0]), "src"), obj[1])
    templates_dir := filepath.Join(obj_dir, "templates")

    tools.CopyDirectory(templates_dir, output_dir)

    viper.SetConfigType("yaml")
    viper.SetConfigName("config")
    viper.AddConfigPath(obj_dir)
    err := viper.ReadInConfig()
    if err != nil {
      panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    values := viper.Get("values")
    // var config map[string]string

    for _, data := range values.([]interface{}) {
      for _, v := range data.(map[interface{}]interface{}) {
        switch t := v.(type) {
        case string, []int:
          fmt.Println(t)
        default:
          fmt.Println("wrong type")
        }
      }
    }
  },
}

func init() {
  objectCmd.AddCommand(createCmd)
  createCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path, will be created if not exists")
}
