package files

import (
	"os"
	"path/filepath"
	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) *Storage {
	return &Storage{basePath: basePath}
}

func (s *Storage) Save(page *storage.Page) (err error) {
	const (
		op     = "files.Save"
		msgErr = "can not save"
	)
	defer func() { err = e.WrapIfErr(op, msgErr, err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

}
