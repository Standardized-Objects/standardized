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
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"log"
	"os"
	"path/filepath"
)

type ObjRepo struct {
	Path      string
	AuthType  string
	AuthValue string
	Url       string
	Name      string
}

func (r *ObjRepo) Start() {

	r.Path = filepath.Join(GetConfigDir(), r.Name)

	if _, err := os.Stat(r.Path); os.IsNotExist(err) {
		os.Mkdir(r.Path, os.ModePerm)
	}

	viper.Set("type", r.AuthType)
	viper.Set("value", r.AuthValue)
	viper.Set("url", r.Url)
	viper.WriteConfigAs(filepath.Join(r.Path, "auth.yaml"))
}

func (r *ObjRepo) Load() {
	r.Path = filepath.Join(GetConfigDir(), r.Name)
	viper.SetConfigFile(filepath.Join(r.Path, "auth.yaml"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	r.AuthType = viper.GetString("type")
	r.AuthValue = viper.GetString("value")
	r.Url = viper.GetString("url")
}

func (r *ObjRepo) Clone() {
	opts := git.CloneOptions{URL: r.Url, Progress: os.Stdout}

	switch r.AuthType {
	case "ssh":
		sshAuth, kr := ssh.NewPublicKeysFromFile("git", r.AuthValue, "")
		if kr != nil {
			log.Fatal(kr)
		}
		opts.Auth = sshAuth
	case "github":
		opts.Auth = &http.BasicAuth{
			Username: "standardized", // yes, this can be anything except an empty string
			Password: r.AuthValue,
		}
	}

	_, err := git.PlainClone(filepath.Join(r.Path, "src"), false, &opts)

	if err != nil {
		log.Fatal(err)
	}
}

func (r *ObjRepo) Update() {
	g, _ := git.PlainOpen(filepath.Join(r.Path, "src"))
	w, _ := g.Worktree()
	opts := git.PullOptions{RemoteName: "origin", Progress: os.Stdout}

	switch r.AuthType {
	case "ssh":
		sshAuth, kr := ssh.NewPublicKeysFromFile("git", r.AuthValue, "")
		if kr != nil {
			log.Fatal(kr)
		}
		opts.Auth = sshAuth
	case "github":
		opts.Auth = &http.BasicAuth{
			Username: "standardized", // yes, this can be anything except an empty string
			Password: r.AuthValue,
		}
	}

	w.Pull(&opts)
}
