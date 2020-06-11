package domain

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/keikohi/gorchitect/domain/err"
)

func LoadGoFiles(dirpath string) ([]string, error) {
	var filepaths []string
	er := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			filepaths = append(filepaths, path)
		}
		return err
	})
	if er != nil || len(filepaths) == 0 {
		return nil, err.NewExitError(err.Error, errors.New("cannot find go files"))
	}
	return toSrcPath(dirpath, filepaths), nil
}

func toSrcPath(projectRoot string, filepaths []string) []string {
	filepath.Dir(projectRoot)
	for i := range filepaths {
		filepaths[i] = strings.Replace(filepaths[i], filepath.Dir(projectRoot)+filepath.FromSlash("/"), "", -1)
	}
	return filepaths
}
