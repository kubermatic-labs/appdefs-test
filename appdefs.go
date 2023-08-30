package appdefs

import (
	"embed"
	"io/fs"
)

//go:embed application-definitions
var f embed.FS

func GetAppDefFiles() ([]fs.File, error) {
	dirname := "application-definitions"
	files := []fs.File{}
	entries, err := f.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			file, err := f.Open(dirname + "/" + entry.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, file)
		}
	}
	return files, nil
}
