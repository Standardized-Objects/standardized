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
  "gopkg.in/src-d/go-git.v4"
  "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
  "github.com/spf13/viper"
  "path/filepath"
  "os"
  "log"
  "fmt"
)

func RepoInit (name string, auth_type string, auth_value string) string {
  repo_path := filepath.Join(GetConfigDir(), name)

  if _, err := os.Stat(repo_path); os.IsNotExist(err) {
    os.Mkdir(repo_path, os.ModePerm)
  }

  f, _ := os.Create(repo_path + "/auth.yaml")
  f.Write([]byte("type: {{.auth_type}}\nvalue: {{.auth_value}}\n"))
  f.Close()

  config := map[string]string{
    "auth_type":  auth_type,
    "auth_value": auth_value,
  }
  ParseTemplate(repo_path + "/auth.yaml", config)

  return repo_path
}

func ParseRepoAuth(path string) (string, string) {
  viper.SetConfigType("yaml")
  viper.SetConfigName("auth")
  viper.AddConfigPath(path)
  err := viper.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }

  return viper.GetString("type"), viper.GetString("value")
}

func Clone (path string, url string) {
  opts := GetCloneOptions(path, url)
  _, err := git.PlainClone(filepath.Join(path, "src"), false, &opts)

  if err != nil {
    log.Fatal(err)
  }
}

func GetCloneOptions(path string, url string) git.CloneOptions {
  switch rtype, rauth := ParseRepoAuth(path); rtype {
  case "ssh":
    fmt.Println("Not yet.")
  case "github":
    return git.CloneOptions{
      URL: url,
      Auth: &http.BasicAuth{
        Username: "standardized", // yes, this can be anything except an empty string
        Password: rauth,
      },
      Progress: os.Stdout,
    }
  }

  return git.CloneOptions{URL: url, Progress: os.Stdout}
}

func GetPullOptions(path string) git.PullOptions {
  switch rtype, rauth := ParseRepoAuth(path); rtype {
  case "ssh":
    fmt.Println("Not yet.")
  case "github":
    return git.PullOptions{
      RemoteName: "origin",
      Auth: &http.BasicAuth{
        Username: "standardized", // yes, this can be anything except an empty string
        Password: rauth,
      },
      Progress: os.Stdout,
    }
  }
  return git.PullOptions{RemoteName: "origin", Progress: os.Stdout}
}
