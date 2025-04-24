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

func SinglePathArg(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	files, err := filepath.Glob(args[0])
	if err != nil {
		return fmt.Errorf("invalid path specified: %s", args[0])
	}
	if len(files) == 0 {
		return fmt.Errorf("no files found")
	}

	return nil
}
