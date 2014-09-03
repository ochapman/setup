/*
* File Name:	common.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */
package setup

type Item interface {
	Installer
	Configer
}

type Installer interface {
	Install() error
}

type Configer interface {
	Config() error
}

type User struct {
	Name  string
	Email string
}
