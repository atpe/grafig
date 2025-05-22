package file

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	Path   = "path"
	Output = "output"
)

func SingleGlobArg(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	_, err := IsGlob(args[0])

	return err
}

func IsGlob(str string) (bool, error) {
	files, err := filepath.Glob(str)
	if err != nil {
		return false, fmt.Errorf("invalid path specified: %s", str)
	}
	if len(files) == 0 {
		return false, fmt.Errorf("no files found")
	}

	return true, nil
}
