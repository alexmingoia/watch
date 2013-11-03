# watch

Watch files and folders for changes, and run commands when they change.

## Installation

### Download binaries

Download a pre-built binary release for Linux, OS X, or Windows

* [Linux 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.1-linux-amd64.tbz) SHA1 checksum: 196be91b6ba3c38a5d9b4992c1f08f0b5d237634
* [Linux 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.1-linux-386.tbz) SHA1 checksum: 89075990c13b05747889536a30018f46d468f852
* [OS X 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.1-darwin-amd64.tbz) SHA1 checksum: 760f13f607a0991031f66f7493b86a0f90ee36ef
* [OS X 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.1-darwin-386.tbz) SHA1 checksum: 40ea20b168c0955f4aa2b1c6fdd2ef3c21302284
* [Windows 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.1-windows-amd64.zip) SHA1 checksum: f41b30e4afdf04e767e379160321b79fbeb7b756
* [Windows 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.1-windows-386.zip) SHA1 checksum: a493ac3f840ed7b7ec1c7ac8f8e9555571082d21

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
