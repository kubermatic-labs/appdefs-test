package appdefs

import (
	"fmt"
	"testing"
)

func TestGetAppDefFiles(t *testing.T) {
	files, err := GetAppDefFiles()
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		finfo, err := file.Stat()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(finfo.Name())
	}
}
