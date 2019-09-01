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
package common

import (
  // "gopkg.in/src-d/go-git.v4"
  "fmt"
  "os"
  homedir "github.com/mitchellh/go-homedir"
)

func GetConfigDir() string {
  home, err := homedir.Dir()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return home + "/.standardize"
}

