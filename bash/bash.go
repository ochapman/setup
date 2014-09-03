/*
* File Name:	bash.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-09-03
 */
package bash

import (
	"errors"
	"github.com/ochapman/setup"
	"io/ioutil"
	"os"
	"path"
)

type Bash struct {
	template string
	config   string
}

// Append to user current bash config
func (b *Bash) Config() error {
	home := os.Getenv("HOME")
	if home == "" {
		return errors.New("$HOME not found")
	}
	srcbuf, err := ioutil.ReadFile(b.template)
	if err != nil {
		return err
	}
	conf := path.Join(home, b.config)
	dst, err := os.OpenFile(conf, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	dst.Write(srcbuf)
	defer dst.Close()
	return nil
}

func NewConf(template, config string) setup.Configer {
	return &Bash{
		template: template,
		config:   config,
	}
}
