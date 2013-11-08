# watch

Watch files and folders for changes, and run commands when they change.

## Installation

### Download binaries

Download a pre-built binary release for Linux, OS X, or Windows

* [Linux 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.2.0-linux-amd64.tbz) SHA1 checksum: 6cf6cd78d6971c5996eb2d95847c7460ce790409
* [Linux 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.2.0-linux-386.tbz) SHA1 checksum: a68a858f067c96cb43819c808a5b5b248c7e8e2c
* [OS X 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.2.0-darwin-amd64.tbz) SHA1 checksum: b1b769c42901c700248dd5a25270a9d962227ad8
* [OS X 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.2.0-darwin-386.tbz) SHA1 checksum: 02a59dfb8c72231b7ce238ae862064c26fa58d60
* [Windows 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.2.0-windows-amd64.zip) SHA1 checksum: 09a9079a92929fddf90a860b32932be047907720
* [Windows 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.2.0-windows-386.zip) SHA1 checksum: 8cc96257cb207ebe171039fcbd2051bcea16016f

### Compile from source

    go get -u github.com/alexmingoia/watch

## Usage

    watch paths... [options]

### Example

    watch src --on-change 'make build'

### Options

`    --on-change <arg>`  Run command on any change  
`-h, --halt`             Exits on error (Default: false)  
`-i, --interval <arg>`   Run command once within this interval (Default: 1s)  
`-r, --recursive`        Watch subfolders (Default: true)  
`-q, --quiet`            Suppress all output (Default: false)

Intervals can be milliseconds(ms), seconds(s), minutes(m), or hours(h).
The format is the integer followed by the abbreviation.

## MIT Licensed
