/*
* File Name:	vim.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */
package vim

import (
	"bytes"
	"errors"
	"github.com/ochapman/setup"
	"io"
	"os"
	"path"
)

type Vim struct {
	user     setup.User
	template string
	config   string
}

func (v *Vim) Config() error {
	home := os.Getenv("HOME")
	if home == "" {
		return errors.New("$HOME is not found")
	}
	src, err := os.Open(v.template)
	if err != nil {
		return err
	}
	defer src.Close()
	conf := path.Join(home, v.config)
	dst, err := os.OpenFile(conf, os.O_CREATE|os.O_WRONLY, 0644)
	defer dst.Close()
	err = v.parseUser(src, dst)
	return err

}

func (v *Vim) parseUser(r io.Reader, w io.Writer) error {
	buf := make([]byte, 2000)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	buf = bytes.Replace(buf[:n], []byte("%NAME%"), []byte(v.user.Name), -1)
	buf = bytes.Replace(buf, []byte("%EMAIL%"), []byte(v.user.Email), -1)
	_, err = w.Write(buf)
	return err
}

func NewConf(user setup.User, template, config string) setup.Configer {
	return &Vim{
		user:     user,
		template: template,
		config:   config,
	}
}
