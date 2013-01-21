# Gokogiri 

by Zhigang Chen and Hampton Catlin

## This Fork

The [main gokogiri repository](http://github.com/moovweb/gokogiri) is under under active development by the folks at [Moovweb](http://www.moovweb.com).  This fork aims to:

* Make gokogiri easier to install with `go get`
* Make API changes (where possible) to be more like Go's built-in XML support
* Add support for parts of libxml2 (like sax) not present in gokogiri

Towards this end, so far:

* [rubex](http://github.com/moovweb/rubex), a cgo wrapper around oniguruma, has been replaced with go's built-in regex
* `master` is the main development branch, based on the `go1` branch, but with memory fixes from master patched in

## About

A libxml2 wrapper for Go.

To install

    - go get github.com/moovweb/gokogiri
	
To run test

    - go test github.com/moovweb/gokogiri/html
    - go test github.com/moovweb/gokogiri/xml
