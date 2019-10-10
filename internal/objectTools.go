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
package tools

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"text/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ObjDef struct {
	Name    string
	OutDir  string
	Path    string
	OutPath string
}

func (o *ObjDef) Load() {
	// Split object name into repo and folder
	obj := strings.Split(o.Name, "/")
	if len(obj) != 2 {
		fmt.Println("Invalid Object Definition : " + o.Name)
		os.Exit(0)
	}

	curr_dir, _ := os.Getwd()
	if obj[0] == "_local" {
		o.Path = filepath.Join(curr_dir, ".stdized", obj[1])
	} else {
		o.Path = filepath.Join(GetConfigDir(), obj[0], "src", obj[1])
	}

	if o.OutDir != "_current_dir_" {
		o.OutPath = filepath.Join(curr_dir, o.OutDir)
		CreateIfNotExists(o.OutPath, os.ModePerm)
	} else {
		o.OutPath = curr_dir
	}

	if _, err := os.Stat(o.Path); os.IsNotExist(err) {
		fmt.Println("Invalid path: " + o.Path)
		os.Exit(0)
	}
}

func (o *ObjDef) Apply() {
	templates_dir := filepath.Join(o.Path, "templates")
	CopyDirectory(templates_dir, o.OutPath)

	// Read object configuration
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(o.Path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Apply values
	values := viper.Get("values")
	config := make(map[string]string, len(values.([]interface{})))

	for _, data := range values.([]interface{}) {

		prompt := promptui.Prompt{
			Label: data.(map[interface{}]interface{})["description"].(string),
		}

		if _default, ok := data.(map[interface{}]interface{})["default"].(string); ok {
			prompt.Default = _default
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Fail %v\n", err)
			return
		}

		config[data.(map[interface{}]interface{})["tag"].(string)] = result
	}

	wlk_err := filepath.Walk(o.OutPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				t, err := template.ParseFiles(path)
				if err != nil {
					log.Fatal(err)
				}

				f, err := os.Create(path)
				if err != nil {
					log.Fatal("create file: ", err)
				}

				err = t.Execute(f, config)
				if err != nil {
					log.Fatal("execute: ", err)
				}
				f.Close()

			}
			return nil
		})

	if wlk_err != nil {
		log.Println(wlk_err)
		os.Exit(0)
	}
}

func (o *ObjDef) RunHooks(hooks_dir string) {
	scripts_path := filepath.Join(o.Path, hooks_dir)
	if Exists(scripts_path) {
		scripts, _ := ioutil.ReadDir(scripts_path)
		for _, scpt := range scripts {
			mode := scpt.Mode()
			if !mode.IsDir() && scpt.Name()[:1] != "." {
				log.Printf("Running hook: " + scpt.Name())
				cmd := exec.Command(filepath.Join(scripts_path, scpt.Name()))
				cmd.Dir = o.OutPath
				cmd.Run()
				log.Printf("Done: " + scpt.Name())
			}
		}
	}
}
