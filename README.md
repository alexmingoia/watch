# watch

Watch files and folders for changes, and run commands when they change.

## Installation

### Download binaries

Download a pre-build binary release for Linux, OS X, or Windows

* [watch-1.0-darwin-amd64.tar.gz]() SHA1 checksum: 
* [watch-1.0-linux-amd64.tar.gz]() SHA1 checksum: 
* [watch-1.0-windows-amd64.tar.gz]() SHA1 checksum: 

### Compile from source

```sh
go get -u github.com/alexmingoia/watch
```

## Usage

```sh
Usage:
  watch paths... [options]

Example:
  watch src --on-change 'make build'

Options:
      --on-change <arg>  Run command on any change
  -h, --halt             Exits on error (Default: false)
  -i, --interval <arg>   Run command once within this interval (Default: 1s)
  -r, --recursive        Watch subfolders (Default: true)
  -q, --quiet            Suppress all output (Default: false)

Intervals can be milliseconds(ms), seconds(s), minutes(m), or hours(h).
The format is the integer followed by the abbreviation.
```

## MIT Licensed
