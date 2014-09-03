/*
* File Name:	git.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */
package git

import (
	"bytes"
	"errors"
	"github.com/ochapman/setup"
	"io"
	"os"
	"path"
)

type Git struct {
	user     setup.User
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
	defer src.Close()
	conf := path.Join(home, g.config)
	dst, err := os.OpenFile(conf, os.O_CREATE|os.O_WRONLY, 0644)
	defer dst.Close()
	err = g.parseUser(src, dst)
	return err
}

func (g *Git) parseUser(r io.Reader, w io.Writer) error {
	buf := make([]byte, 1000)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	buf = bytes.Replace(buf[:n], []byte("%NAME%"), []byte(g.user.Name), -1)
	buf = bytes.Replace(buf, []byte("%EMAIL%"), []byte(g.user.Email), -1)
	_, err = w.Write(buf)
	return err
}

func NewConf(user setup.User, template, config string) setup.Configer {
	return &Git{
		user:     user,
		template: template,
		config:   config,
	}
}
