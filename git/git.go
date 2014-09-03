/*
* File Name:	git.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */
package git

import (
	"errors"
	"github.com/ochapman/setup"
	"io"
	"os"
	"path"
)

type Git struct {
	template string
	config   string
}

func (g *Git) Config() error {
	home := os.Getenv("HOME")
	if home == "" {
		return errors.New("$HOME not found")
	}
	src, err := os.Open(g.template)
	if err != nil {
		return err
	}
	conf := path.Join(home, g.config)
	dst, err := os.OpenFile(conf, os.O_CREATE|os.O_WRONLY, 0644)
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func NewConf(template, config string) setup.Configer {
	return &Git{
		template: template,
		config:   config,
	}
}
