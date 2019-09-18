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
	homedir "github.com/mitchellh/go-homedir"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GetConfigDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf_dir := filepath.Join(home, ".standardize")

	if _, err := os.Stat(conf_dir); os.IsNotExist(err) {
		os.Mkdir(conf_dir, os.ModePerm)
	}

	return conf_dir
}

func ParseTemplate(path string, values map[string]string) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, values)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}

func RunHooks(path string, outdir string) {
	if Exists(path) {
		scripts, _ := ioutil.ReadDir(path)
		for _, scpt := range scripts {
			mode := scpt.Mode()
			if !mode.IsDir() && scpt.Name()[:1] != "." {
				log.Printf("Running hook: " + scpt.Name())
				cmd := exec.Command(filepath.Join(path, scpt.Name()))
				cmd.Dir = outdir
				cmd.Run()
				log.Printf("Done: " + scpt.Name())
			}
		}
	}
}
