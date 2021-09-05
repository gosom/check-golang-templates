package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

func main() {
	var (
		folder  string
		pattern string
	)

	flag.StringVar(&folder, "folder", "", "The folder the templates exist.")
	flag.StringVar(&pattern, "pattern", "", "The pattern to load template definitions (relative to folder). Example *.jet")

	flag.Parse()

	if len(folder) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	if len(pattern) == 0 {
		pattern = "*"
	}

	loader := jet.NewOSFileSystemLoader(folder)
	views := jet.NewSet(
		loader,
	)
	searchpattern := filepath.Join(folder, pattern)

	filenames, err := getMatchedFilenames(folder, pattern)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(3)
	}

	if len(filenames) == 0 {
		fmt.Fprintf(os.Stderr, "pattern matches no files: %#q\n", searchpattern)
		os.Exit(4)
	}

	problematic := map[string]string{}
	for i := range filenames {
		fname := filenames[i][len(folder):]
		fname = strings.TrimLeft(fname, "/")
		_, err := views.GetTemplate("./" + fname)
		if err != nil {
			problematic[filenames[i]] = err.Error()
		}
	}

	fmt.Fprintf(os.Stderr, "checked %d templates...", len(filenames))

	if len(problematic) == 0 {
		fmt.Fprintf(os.Stderr, "OK\n")
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "ERRORS:\n")
	for _, msg := range problematic {
		fmt.Fprintf(os.Stderr, " - "+msg+"\n")
	}
	os.Exit(1)

}

func getMatchedFilenames(folder, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil

		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return matches, nil
}
