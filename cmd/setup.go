/*
* File Name:	setup.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */

package main

import (
	"github.com/ochapman/setup/git"
)

func main() {
	g := git.NewConf("template/git.config", ".gitconfig")
	g.Config()
}
