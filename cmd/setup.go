/*
* File Name:	setup.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */

package main

import (
	"flag"
	"github.com/ochapman/setup"
	"github.com/ochapman/setup/bash"
	"github.com/ochapman/setup/git"
	"os"
)

func main() {
	var user setup.User
	flag.StringVar(&user.Name, "name", "", "User name")
	flag.StringVar(&user.Email, "email", "", "User email")
	flag.Parse()
	if user.Name == "" || user.Email == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	g := git.NewConf(user, "template/git.config", ".gitconfig")
	g.Config()
	b := bash.NewConf("template/bash", ".bash_profile")
	b.Config()
}
