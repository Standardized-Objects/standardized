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
  "github.com/manifoldco/promptui"
  "os"
  "strings"
  "path/filepath"
  "log"
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

    // Split object name into repo and folder
    obj := strings.Split(args[0], "/")
    if len(obj) != 2 {
      fmt.Println("Invalid Object Definition : " + args[0])
      os.Exit(0)
    }

    // Copy templates from object definition
    curr_dir , _ := os.Getwd()

    var obj_dir string
    if obj[0] == "_local" {
      obj_dir = filepath.Join(filepath.Join(curr_dir, ".stdized"), obj[1])
    } else {
      obj_dir = filepath.Join(filepath.Join(filepath.Join(tools.GetConfigDir(), obj[0]), "src"), obj[1])
    }

    if !tools.Exists(obj_dir) {
      fmt.Println("Invalid path: " + obj_dir)
      os.Exit(0)
    }

    templates_dir := filepath.Join(obj_dir, "templates")

    var _out string
    if outputPath != "" {
      _out =  filepath.Join(curr_dir, outputPath)
    } else {
      _out = curr_dir
    }

    tools.CreateIfNotExists(_out, os.ModePerm)

    // Run pre create hooks
    tools.RunHooks(filepath.Join(obj_dir, "precreate"), _out)
    tools.CopyDirectory(templates_dir, _out)

    // Read object configuration
    viper.SetConfigType("yaml")
    viper.SetConfigName("config")
    viper.AddConfigPath(obj_dir)
    err := viper.ReadInConfig()
    if err != nil {
      panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    // Apply values
    values := viper.Get("values")
    config :=  make(map[string]string, len(values.([]interface{})))

    for _, data := range values.([]interface{}) {

      prompt := promptui.Prompt{
        Label:    data.(map[interface{}]interface{})["description"].(string),
      }

      if _default, ok := data.(map[interface{}]interface{})["default"].(string); ok {
        prompt.Default =  _default
      }

      result, err := prompt.Run()

      if err != nil {
        fmt.Printf("Fail %v\n", err)
        return
      }

      config[data.(map[interface{}]interface{})["tag"].(string)] = result
    }

    wlk_err := filepath.Walk(_out,
    func(path string, info os.FileInfo, err error) error {
      if err != nil {
        return err
      }

      if !info.IsDir() {
        tools.ParseTemplate(path, config)
      }
      return nil
    })

    if wlk_err != nil {
      log.Println(wlk_err)
      os.Exit(0)
    }

    // Run post create hooks
    tools.RunHooks(filepath.Join(obj_dir, "postreate"), _out)
  },
}

func init() {
  objectCmd.AddCommand(createCmd)
  createCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path, will be created if not exists")
}
