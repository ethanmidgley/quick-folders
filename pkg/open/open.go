package open

import (
	"errors"
	"os/exec"
	"path/filepath"
	"runtime"
)

func OpenFolder(dir string) error {

	stc := map[string]string{
		"darwin":  "open",
		"linux":   "xdg-open",
		"windows": "explorer",
	}

	sys := runtime.GOOS

	if c, has := stc[sys]; has {
		dir = filepath.FromSlash(dir)
		e := exec.Command(c, dir)
		// there needs to be a better way to debug this as we assume there is no error
		// by default the command returns an error as it opens new window
		e.Run()
		return nil
	} else {
		return errors.New("unsupported platform")
	}

}
