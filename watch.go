package main

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"github.com/jessevdk/go-flags"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const usage = `
Usage:
  watch paths... [options]

Example:
  watch src --on-change 'make build'

Options:
      --on-change <arg>  Run command on any change
  -h, --halt             Exits on error (Default: false)
  -i, --interval <arg>   Run command once within this interval (Default: 1s)
  -n, --no-recurse       Skip subfolders (Default: false)
  -q, --quiet            Suppress standard output (Default: false)

Intervals can be milliseconds(ms), seconds(s), minutes(m), or hours(h).
The format is the integer followed by the abbreviation.
`

var (
	last     time.Time
	interval time.Duration
	paths    []string
	err      error
)

var opts struct {
	Help        bool   `short:"h" long:"help"      description:"Show this help message" default:false`
	Halt        bool   `short:"h" long:"halt"      description:"Exits on error (Default: false)" default:false`
	Quiet       bool   `short:"q" long:"quiet"     description:"Suppress standard output (Default: false)" default:false`
	Interval    string `short:"i" long:"interval"  description:"Run command once within this interval (Default: 1s)" default:"1s"`
	NoRecursive bool   `short:"n" long:"no-recursive" description:"Skip subfolders (Default: false)" default:false`
	OnChange    string `long:"on-change"           description:"Run command on change."`
}

func init() {
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	paths, err = ResolvePaths(args[1:])

	if len(paths) <= 0 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2) // 2 for --help exit code
	}

	interval, err = time.ParseDuration(opts.Interval)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	last = time.Now().Add(-interval)
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	done := make(chan bool)

	// clean-up watcher on interrupt (^C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		if !opts.Quiet {
			fmt.Fprintln(os.Stdout, "Interrupted. Cleaning up before exiting...")
		}
		watcher.Close()
		os.Exit(0)
	}()

	// process watcher events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if !opts.Quiet {
					fmt.Fprintln(os.Stdout, ev)
				}
				if time.Since(last).Nanoseconds() > interval.Nanoseconds() {
					last = time.Now()
					err = ExecCommand()
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						if opts.Halt {
							os.Exit(1)
						}
					}
				}
			case err := <-watcher.Error:
				fmt.Fprintln(os.Stderr, err)
				if opts.Halt {
					os.Exit(1)
				}
			}
		}
	}()

	// add paths to be watched
	for _, p := range paths {
		err = watcher.Watch(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// wait and watch
	<-done
}

func ExecCommand() error {
	if opts.OnChange == "" {
		return nil
	} else {
		args := strings.Split(opts.OnChange, " ")
		cmd := exec.Command(args[0], args[1:]...)

		if !opts.Quiet {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
		cmd.Stdin = os.Stdin

		return cmd.Run()
	}
}

// Resolve path arguments by walking directories and adding subfolders.
func ResolvePaths(args []string) ([]string, error) {
	var stat os.FileInfo
	resolved := make([]string, 0)

	var once sync.Once
	var recurse error = nil

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if recurse == nil && opts.NoRecursive && info.IsDir() {
			once.Do(func() {
				recurse = filepath.SkipDir
			})
		}

		resolved = append(resolved, path)

		return recurse
	}

	for _, path := range args {
		if path == "" {
			continue
		}

		stat, err = os.Stat(path)
		if err != nil {
			return nil, err
		}

		if !stat.IsDir() {
			resolved = append(resolved, path)
			continue
		}

		err = filepath.Walk(path, walker)
	}

	return resolved, nil
}
