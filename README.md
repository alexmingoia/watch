# watch

Watch files and folders for changes, and run commands when they change.

## Installation

### Download binaries

Download a pre-built binary release for Linux, OS X, or Windows

* [Linux 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.3.0-linux-amd64.tbz) SHA1 checksum: e8ddec7113b149a2ed5da585a597b990338ad182
* [Linux 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.3.0-linux-386.tbz) SHA1 checksum: dd5df6da001ee9a6b73ea0dfcbecf6619b104098
* [OS X 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.3.0-darwin-amd64.tbz) SHA1 checksum: 1e3c38545d8a84d95230401cd74219eb077950bb
* [OS X 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.3.0-darwin-386.tbz) SHA1 checksum: 8bc05f061f711ef7e1e9427eef3e93a547a02c96
* [Windows 64-bit](https://s3.amazonaws.com/acm-binaries/watch-0.3.0-windows-amd64.zip) SHA1 checksum: 0a69d19a79ba92209a87f4088aa83d568f536139
* [Windows 32-bit](https://s3.amazonaws.com/acm-binaries/watch-0.3.0-windows-386.zip) SHA1 checksum: 63f3f85f990d8371daf3fcae9765ac6347131bf4

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
`-n, --no-recurse`       Skip subfolders (Default: false)  
`-V, --version`          Output the version number
`-q, --quiet`            Suppress all output (Default: false)

Intervals can be milliseconds(ms), seconds(s), minutes(m), or hours(h).
The format is the integer followed by the abbreviation.

## MIT Licensed
