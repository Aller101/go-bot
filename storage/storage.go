package storage

import (
	"crypto/sha1"
	"errors"
	"io"
	"read-adviser-bot/lib/e"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("no saved pages")

type Page struct {
	URL      string
	UserName string
}

func (p *Page) Hash() (string, error) {
	const op = "storage.Hash"
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap(op, "can not calc hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap(op, "can not calc hash", err)
	}

	return string(h.Sum(nil)), nil
}
