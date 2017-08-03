// Code generated by go-bindata.
// sources:
// frontend/application.coffee
// frontend/application.js
// frontend/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _frontendApplicationCoffee = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xdf\x4f\xdb\x30\x10\x7e\xf7\x5f\x71\x6b\x91\x9c\x48\x25\x14\xc6\xf6\x50\x29\x93\x36\xf6\x53\x02\x34\xd1\xa1\x3d\x4c\x68\x72\xe2\x4b\x63\x48\xed\xc8\x76\x28\x13\xf4\x7f\x9f\x1c\x27\x69\x52\x0a\x02\x5e\xa8\xfd\xdd\x7d\xb9\xdc\xf7\xdd\xc5\x60\x5a\x69\xfc\xc9\x8c\x59\x29\xcd\x21\x06\x59\x15\x05\x21\xa9\x46\x66\x71\x8e\xa9\x46\x0b\x31\x04\x21\xec\x7f\x20\x00\xa6\xbd\xd8\x0b\xe8\x38\x53\x7a\x79\xd2\x8b\xa3\x61\x94\x09\xc9\x03\x6a\xf1\xce\x32\x8d\x8c\x86\xd1\x2d\x2b\x82\x90\x10\x00\x91\xd5\x39\x78\x67\xb5\xbb\x2f\xb5\x2a\x81\xa6\x39\xa6\x37\xc8\x29\x01\xa8\xb9\x87\x95\x9c\x31\x9b\x47\x9a\x49\xae\x96\x41\x18\x59\x35\xb7\x5a\xc8\x45\xf0\xf6\x7d\x18\x99\x2a\x31\xfe\x74\x14\xb6\xc9\xbe\xb0\x6f\x22\x49\x50\x0b\x93\x7f\xfc\x32\x8f\x50\xa6\x81\x47\x26\x5b\xf4\x75\x4d\x7b\x11\xbb\x66\x77\x40\x59\x29\x0e\xfc\x0b\xd3\x49\xcd\xc6\x99\x65\xb3\xfa\x57\xcb\x3c\x6b\xfe\x77\xf0\xaf\x7f\x25\xce\x60\x74\x6d\x94\x1c\xf9\x0a\x2c\xb3\x95\x39\x51\x1c\xdb\xcc\xa3\xe9\x61\x9b\xe6\xdb\xc4\x1b\xe0\x78\x3a\x9d\x81\x51\x4b\xb4\xb9\x90\x8b\xdf\x5a\xc9\x45\x83\xbc\x7b\x12\x39\x9e\x1e\xcf\x3a\x19\xfc\xdf\x18\xce\x54\x7a\x03\x99\xd2\x20\xa4\x45\x9d\xb1\x14\xc1\xa2\xb1\xa2\xcb\x82\x9d\xcf\xdf\xdc\xff\x15\x7c\x06\x34\x53\x2a\x61\x9a\xba\x96\x64\xac\x30\x48\x88\x7b\xc3\x73\x65\xbf\xaa\x4a\xf2\x9e\xfc\x4e\x41\xa9\x6c\xe6\xae\x69\x18\x99\x5c\xad\x9c\xba\x39\x33\xf9\xa9\x62\xfd\x48\x77\x05\x31\xac\x84\xe4\x6a\x15\x15\x2a\x65\x56\x28\x19\xb9\x6b\x6f\x06\xf7\x2b\x2a\x50\x2e\x6c\x0e\x71\x0c\xd3\xba\x34\x8d\xb6\xd2\x92\x6c\xf2\x39\xa6\x8a\xe3\xe5\xc5\x8f\x13\xb5\x2c\x95\x44\x69\x03\x87\xd4\xea\x95\x4c\x5b\x03\xb1\x67\x32\x65\x21\x2c\xd0\x07\xea\xd9\x6b\xac\x47\x7f\x54\xd3\x37\xa4\x35\xf8\x67\x7a\xb5\xdb\x75\x1e\x3d\xbc\xaa\x4d\xcb\x3b\xfe\xce\x70\x87\xe1\xc6\x39\x23\xe7\x9c\x05\xda\x83\xf1\xbd\xe0\xeb\xd1\xe4\x35\xee\xa8\x05\xed\xf7\xb9\x73\x8d\xb3\x40\xae\x56\x9f\x99\x65\x84\x08\x29\xec\x27\x21\xb9\xd9\x92\x61\xc7\xf0\x25\x42\x72\xa0\xa6\x4a\x96\xc2\xd2\x09\xf4\x47\xb8\xd5\x0e\x57\xf3\x66\x1c\x22\xc9\x6e\x13\xa6\xf7\x13\x37\x60\x5d\x72\x5a\x88\xf4\x86\x4e\xa0\x0b\xac\x13\xbd\x8a\x6d\x8c\xeb\x47\x9a\x33\xb9\x40\x3a\x81\x56\x7a\x42\xba\x94\x5e\xa1\x1b\xdd\x35\x66\x10\x0f\xcf\x5e\xb3\x80\x8e\x69\xe8\xc5\x68\xbc\x37\x70\xac\x23\x73\x4d\x1a\xec\x9f\xef\x8d\x39\x98\x65\x51\xe7\x63\x2f\xfc\x96\x9c\x6f\x9a\x7d\xb6\x71\x7c\x93\x3b\x1a\xdf\x6f\xce\xeb\x87\xfa\xd4\x4b\x5c\x3b\xd5\x2a\x5d\xf8\xc8\x67\xca\x5e\x8f\x07\x44\x23\xd2\x74\xba\x64\x12\x8b\xf3\xb6\x25\x34\x8c\x72\xc1\x31\x08\xfb\xa8\x87\x2e\x2f\x4e\x37\x93\xf4\x04\xea\xb7\xaa\x90\x65\x65\xfd\x4a\x75\xa5\xbd\x2c\x38\x53\x69\x65\x5e\xca\x6c\xb0\xc0\xd4\xfa\x7d\xfd\xc4\x77\xa1\x35\xe6\x4e\x5d\x20\x86\xbe\x28\x2f\x54\x64\x7b\x69\x73\x7c\x7e\x69\xbf\xbe\xbb\x03\xb4\xb7\xbf\x06\xf7\xdd\xd6\x5d\xb9\xad\xbb\x93\xf3\x02\x19\xef\x1e\xf9\x58\xb2\x01\xbc\xeb\x4b\xd8\x7e\x43\xc8\x70\xc3\x6f\x4d\xf6\xa3\x42\xda\x35\x4b\xf6\x7c\x50\xb7\x12\xea\xc7\xb7\x13\x18\x84\xe4\x7f\x00\x00\x00\xff\xff\x2c\x12\x62\xf6\xcc\x07\x00\x00")

func frontendApplicationCoffeeBytes() ([]byte, error) {
	return bindataRead(
		_frontendApplicationCoffee,
		"frontend/application.coffee",
	)
}

func frontendApplicationCoffee() (*asset, error) {
	bytes, err := frontendApplicationCoffeeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "frontend/application.coffee", size: 1996, mode: os.FileMode(436), modTime: time.Unix(1501789334, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _frontendApplicationJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xe0\x92\x00\xa4\x50\x8d\xb1\xb3\x6c\x0f\x36\xf4\xb0\x79\x5b\x3b\xa0\x2b\x86\x64\xc5\x1e\x86\x62\xa0\xc4\x93\xc5\x56\x26\x05\x92\x8a\x53\xb4\xfe\xee\x03\x29\x4a\xa2\x14\xbb\xff\xf2\x10\xe1\x78\xc7\xbb\xe3\xdd\xef\x7e\xe7\xeb\x6b\xf4\x1c\x24\x68\x66\x81\xa3\xfc\x3d\xda\xaa\xb2\x04\xb8\x2f\xb4\x68\x2c\x5a\xd1\xd5\x0d\xbd\x5d\x90\xb2\x95\x85\x15\x4a\x92\x04\x7d\x58\x20\xf4\xc0\x34\x2a\x34\x30\x0b\xf7\x50\x68\xb0\x29\xe2\xcc\xb2\x57\xca\xfe\xae\x5a\xc9\x53\x54\x31\x53\xbd\x54\x8c\xa7\x48\x48\x61\x7f\x11\x92\x9b\x14\x49\x38\xf4\xd6\xc6\x7f\xb7\xde\x03\xf7\x62\xab\xe1\x2f\x66\xcc\x41\x69\x27\x57\xea\xf0\x2b\xb3\x2c\x45\x46\xed\xc1\x56\x42\xee\xfe\xd1\x4a\xee\x36\x8b\x05\x9a\x19\xa3\x0c\xc9\xb6\xae\xbd\x26\xce\x08\x65\x68\x96\x73\x97\x75\x17\x79\xe3\x65\xd3\x5b\x5e\x11\x7c\x59\x2a\xbd\xdf\x46\x0e\x70\x42\x4b\x21\x39\xc1\x16\x1e\x2d\xd3\xc0\x70\x42\x1f\x58\x4d\x92\xee\xae\x28\x11\x71\xd7\xe0\xd1\x6a\xa7\x6a\xb4\x6a\x08\x2e\x2a\x28\xde\x01\xc7\x49\x1f\xf2\x44\xba\x7f\x32\x5b\x51\xcd\x24\x57\x7b\x92\x50\xab\xee\xad\x16\x72\x47\x7e\xf8\x29\xa1\xa6\xcd\x4d\x27\xdd\x84\x30\x51\x92\xcf\x45\x9e\x83\x16\xa6\xfa\xf9\xb7\x7b\x0a\xb2\x20\x66\x2c\x66\x14\x21\x5c\x3c\xfa\xff\x57\x94\xbd\x65\x8f\x04\xb3\x46\x5c\x77\xd5\xc1\xe9\x90\x99\x6b\xd9\x7a\x90\xfa\x48\xeb\xf0\x0d\xc7\xc7\x34\xb2\xfe\xfb\x7d\x03\x6b\x74\xf1\xd6\x28\x79\xd1\x9f\x1b\xcb\x6c\x6b\xb6\x8a\x43\xec\xeb\x66\xb9\x5a\xcf\xba\x3c\xe8\x6e\x97\xcb\xf5\xac\xb1\xa3\xf2\xc7\x4f\x29\x6f\x97\xb7\xeb\xa7\x6d\xed\xfe\x34\xd8\x56\xcb\x69\x4c\x12\x1b\xf4\x0f\xfc\x4f\xf0\x35\xc2\xa5\x52\x39\xd3\x38\xd2\x1f\x87\x92\xf7\xd5\xeb\xbf\xbd\x26\x84\x28\x59\x6d\xc0\x9d\x1c\x3d\xea\x62\xe4\x9f\x42\x5d\xb8\xe5\xc0\x22\x95\x2d\x9d\x19\x4e\xa8\x83\x78\x87\xa5\xce\x4b\x3f\x31\xe7\x70\xeb\xf4\x29\x12\x3c\x45\x0d\xd3\xd6\x74\x09\xb9\x43\x94\xa1\x83\x90\x5c\x1d\x68\xad\x0a\xe6\xee\x51\x77\x3c\xa2\xd4\x49\xb4\x06\xb9\xb3\x15\xca\xb2\x0c\x2d\xc7\xba\x75\xa9\xc5\x80\x09\x1e\x39\x14\x8a\xc3\xeb\xbb\x3f\xb6\x6a\xdf\x28\x09\xd2\x7a\x37\xa1\x0e\x3e\x03\x94\x79\x63\x6a\x9a\x5a\x58\x82\x3f\xe2\x68\x30\xbc\x41\x1c\xf3\x66\x8c\x19\x22\x78\x93\x7f\x97\x6f\x36\xe7\xc6\xa4\x33\x58\xbd\x89\xb3\x13\x7c\x88\x3a\x4c\xca\x6a\xda\x9c\x80\xf8\x0b\x87\xf8\x1d\xd8\xeb\x0b\xf4\xcc\x57\xed\xc3\x37\xc1\xd8\x03\x6e\x42\x6d\x11\xc2\x1d\x50\x03\x51\x9d\x80\x4b\xd7\xd6\x81\xfe\x4e\xf5\xf5\x0c\xed\xe4\x9e\x76\x4c\x9b\xef\x85\xc5\xe9\x84\xd5\xc2\x5b\x3d\x96\x46\x36\xa5\x92\x3d\xe4\x4c\x7f\x9f\x3b\x56\x19\x1c\x14\xb5\x28\xde\xe1\x88\x76\x67\x85\x22\x1d\x6c\x7a\x73\x57\xd6\xa2\x62\x72\xe7\x28\xa2\x47\x63\xf4\x92\xc1\xcd\xa9\x97\x8c\xd0\xd3\x50\xa2\x6c\x2a\xf7\x10\xb9\xc4\xc9\xd0\xf0\xd3\xa3\x34\x19\xde\x38\x8e\x6b\xc1\x53\x16\x7f\xe1\x67\xa2\xd5\x75\x4c\xe7\x2f\x02\x82\x99\x65\x74\x18\xf8\x11\x9a\x33\x9c\x7d\x97\x75\xfb\x63\xc2\xd7\xa3\x97\x48\x78\x86\x2e\x3e\x3a\x30\x4d\x1d\xc4\xe8\x6c\x75\x8d\x32\x44\x3e\xf1\xf6\xc4\x79\xb9\x0c\x5e\x82\xe3\xb1\xa3\x0d\x93\x50\xbf\xea\xcb\x8c\x13\x5a\x09\x0e\x24\x99\x19\x74\xda\xd7\x77\x2f\x27\x24\x72\xc6\xa0\xdb\x60\x42\x36\xad\x0d\xeb\xab\xd5\xf5\x97\x5f\x28\x55\xd1\x9a\xaf\x08\x60\xa0\x86\xc2\x92\x29\xd0\xce\xec\xeb\xa1\xe7\x61\x84\x3e\xdf\xee\xd9\xd2\x8e\x3a\xfc\xd5\xed\x9d\xaf\x53\x0e\x5f\xb6\x4e\xbf\xa9\x4b\x73\x83\x68\x0b\xcc\x55\xc3\xe2\x3b\xb8\xc5\x77\xce\xf9\x1d\x30\x3e\x84\x8f\x31\x10\xad\x9a\xa7\x86\x27\x7f\xcd\x98\x91\x19\x42\x3b\x26\xab\xf7\x33\xfb\xec\x49\xba\xf3\xad\x76\x35\xff\xe1\x18\x31\xe2\x2c\xe9\x9e\x72\xc2\xf5\x64\xb3\x58\x1c\x13\x5a\xb0\xba\x26\xb6\x12\x26\xd9\x2c\xfe\x0f\x00\x00\xff\xff\xde\xe1\x35\x2f\xa9\x0a\x00\x00")

func frontendApplicationJsBytes() ([]byte, error) {
	return bindataRead(
		_frontendApplicationJs,
		"frontend/application.js",
	)
}

func frontendApplicationJs() (*asset, error) {
	bytes, err := frontendApplicationJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "frontend/application.js", size: 2729, mode: os.FileMode(436), modTime: time.Unix(1501789349, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _frontendIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x7b\x53\xdb\xb8\x16\xff\xbb\xfd\x14\xa7\xee\xcc\xed\x6e\xb7\xb6\xf3\x28\xd0\x52\x27\x33\x14\x58\x0a\x0b\xa4\x3c\x0a\x85\x9d\xfd\x43\xb6\x8f\x6d\x05\x59\x32\x92\x9c\x10\xee\xdc\xef\x7e\x47\xb2\x1d\xf2\x62\xcb\xb6\xdd\x76\x9a\x44\xf2\xf1\x91\x7e\xe7\xfc\xce\x43\x6a\xf0\x62\x67\xb0\x7d\x7e\xf5\x79\x17\x32\x9d\xb3\xfe\xf3\xc0\x7c\x01\x23\x3c\xed\x39\xc8\x9d\xfe\x73\x80\x20\x43\x12\x9b\x1f\x00\x41\x8e\x9a\x40\x94\x11\xa9\x50\xf7\x9c\x52\x27\xee\x3b\x67\xf6\x51\xa6\x75\xe1\xe2\x6d\x49\x47\x3d\xe7\xab\xfb\x65\xcb\xdd\x16\x79\x41\x34\x0d\x19\x3a\x10\x09\xae\x91\xeb\x9e\xb3\xbf\xdb\xc3\x38\xc5\xb9\x37\x39\xc9\xb1\xe7\x8c\x28\x8e\x0b\x21\xf5\x8c\xf0\x98\xc6\x3a\xeb\xc5\x38\xa2\x11\xba\x76\xf0\x06\x28\xa7\x9a\x12\xe6\xaa\x88\x30\xec\xb5\x1b\x45\x2f\x5c\x17\xce\x33\x04\x12\x8a\x11\x42\x17\xac\x62\x4d\x52\x05\xaf\xf3\x52\xe9\xd7\x10\x89\x1c\x21\xa1\x52\x69\xa0\x1c\x74\x86\x60\xb0\x7d\x00\xc2\x27\x20\x74\x86\xd2\x8e\x9b\xb5\xc1\xbc\x54\xbd\xf3\x9a\x24\x1a\xe5\x6b\xf3\x8a\xc2\x4a\xa5\xeb\xd6\xab\x6a\xaa\x19\xf6\x07\xe7\x67\xe0\xc2\x80\x23\x9c\xd3\x1c\xe1\x0c\x23\x89\x5a\x05\x7e\xf5\xf4\xf9\xc3\x06\x3f\x0a\xa1\x95\x96\xa4\xb0\x1a\x9e\x3d\x0b\x18\xe5\x37\x20\x91\xf5\x1c\xa5\x27\x0c\x55\x86\xa8\x1d\xc8\x24\x26\x3d\xc7\xd8\x53\x6d\xfa\x7e\x14\xf3\xa1\xf2\x22\x26\xca\x38\x61\x44\xa2\x17\x89\xdc\x27\x43\x72\xe7\x33\x1a\x2a\x3f\x34\x3a\xc7\x44\x47\x99\xdf\xf5\xba\xde\x86\x9f\x30\xa2\xd9\xa4\x9a\x37\x6b\x79\x39\xe5\x5e\xa4\x94\x63\xf7\x51\xfd\xa1\x5c\x63\x2a\xa9\x9e\xf4\x1c\x95\x91\xce\xda\xba\x2b\xdb\x97\x74\x78\xe9\x9f\x1d\x1f\xa5\x83\xf1\xcd\xda\xe1\xce\xcd\xc6\xc9\xe9\x27\xb9\x2e\x2e\xae\xae\xc2\xab\xcb\xa3\xb1\xdf\xbe\x19\x7c\x4d\xae\x0e\x92\xb4\xe7\x40\x24\x85\x52\x42\xd2\x94\xf2\x9e\x43\xb8\xe0\x93\x5c\x94\xca\x01\xbf\xb6\xcb\x4f\xc0\x95\x08\xae\x5d\x32\x46\x25\x72\xf4\xdf\x7a\x1b\x5e\xcb\x8f\xd4\xfc\xf4\xd3\x90\xe1\xb5\x94\x07\xd1\x78\x27\xf2\xbb\xe5\x4e\xa6\x62\xbd\xde\x56\x87\x1d\x31\xf8\x78\xd5\x5d\xef\xdc\x1e\x75\x99\xe0\xed\x74\xb2\x7b\x77\x73\xd8\xfa\x7b\x64\x0f\x7e\xfc\x74\x7e\x74\xb8\x06\x2a\xa3\x39\x10\x1e\xc3\x29\xaa\x42\xf0\xd8\x1b\x2a\x48\x84\x84\xfd\xdd\x77\xa0\xca\xc2\x70\x19\x44\x52\x0b\x23\xc3\x1c\xb9\x56\xf6\x85\x1c\x63\x4a\xe0\xb6\x44\x49\x71\x86\x4d\x46\xf5\xe5\xd6\xe9\xf1\xfe\xf1\xde\xe6\xac\xd2\x58\xa0\xe2\xaf\x34\x8c\x85\xbc\x01\x9a\xc0\x44\x94\x60\xa2\xc5\xb2\xb8\x20\x29\xc2\x88\x12\x48\x28\xc3\x4d\xdf\x9f\x53\xf7\x27\x4d\x80\x69\xd8\xdf\x85\xf7\x7f\xf5\x6b\x2b\x05\x2a\x92\xb4\xd0\xa0\x64\xf4\x64\x67\x98\xcc\xb0\xa6\x32\x3a\xf2\xbb\xde\x86\xd7\x7d\x18\x5b\x17\x0c\xe7\x3c\xb0\xda\x0b\xdd\x83\x89\xbf\x17\x9e\x1d\xca\xb4\x25\xde\x4f\xd6\xae\xd7\x78\xbb\x1c\xb7\x6e\xef\xae\x77\xb7\x3f\x6d\x6c\xaf\x0f\x4e\x8a\x8b\x8f\xe9\xf1\xef\x57\xa4\xf5\x37\xfc\xea\x07\x7e\xb5\xf9\x1f\xc1\x22\xa7\x86\xf5\xdb\xde\x5b\xaf\x33\x9d\x78\x2a\x94\x74\x9d\x6e\x25\xa3\xeb\xe2\x37\xbe\x73\xd2\x39\x8f\xcf\x4f\xfd\x8b\x8b\x3f\x0e\x92\x6e\xb8\x27\xc5\xdb\x32\x5c\x4b\x46\x97\x67\x97\x17\xa7\xb4\x73\xbc\xfb\x64\x28\xc1\x8b\x3f\x91\xc7\x34\xf9\xcb\x78\xaf\x9a\xb1\x91\xd3\x00\x7d\x69\xf8\xae\x33\xca\xd3\xb1\x14\x3c\x7d\x03\x2f\xb9\xd0\x89\x28\x79\xfc\x06\x5e\x16\x84\x23\x3b\x45\x12\x57\x99\xa7\x99\xa9\x46\x5f\x4e\x0f\xe1\xbf\x53\x4c\x31\x55\x05\x23\x93\x4d\xe0\x82\xe3\x87\x7a\xfa\x7f\xf5\xb7\x97\x08\xa1\x51\xce\x88\xdb\x68\x53\xf4\x1e\x37\xa1\xe5\xbd\xc7\xfc\xc3\xf4\x89\xc6\x3b\xed\x12\x46\x53\xbe\x09\x11\x72\x8d\xf2\xe1\x59\x24\x98\x90\x9b\xf0\xb2\x93\x98\xbf\xf3\xab\x04\xfe\x14\x57\xe0\x37\xd5\x25\x08\x45\x3c\xb1\x09\x91\x93\x11\x44\x8c\x28\xd5\x73\x38\x19\x85\x44\x42\xf5\xe5\xc6\x98\x90\x92\x69\xc7\x48\x3d\x0b\x62\x3a\x15\x33\x09\x9b\x50\x8e\xd2\x4d\x58\x49\xe3\x4a\xe0\x59\x95\x6e\xa5\x09\x36\xf3\x4f\x8b\x34\x65\x08\x29\x6a\x48\xa5\x28\x0b\x8c\x6d\xa4\x86\xa8\x0d\xde\x5c\x84\x94\x61\x63\x9b\x3a\x37\xcf\xaf\x52\xef\xc2\x6c\x18\x65\xbd\xc6\xb3\x20\x2c\xb5\x16\x1c\xf4\xa4\xc0\x9e\x53\x0d\x9c\x85\x37\xea\x95\x23\xc1\x18\x29\x14\xc6\x0e\xc4\x44\x93\x7a\xda\xec\xbe\x9a\x6f\xa6\x89\x4c\x4d\x75\x7d\x19\x2a\x17\xef\x48\x5e\x30\x74\x6b\x45\x8d\xa4\xdb\x76\x80\x48\x4a\x5c\xbc\x2b\x08\x8f\x31\xee\x39\x09\x61\x0a\x9b\x4d\x3d\x0b\x54\x41\x78\xb3\x0b\x25\x5d\xc1\xd9\xc4\xe9\x9f\x57\xfb\xe0\x64\x44\x53\xa2\xa9\xe0\x81\x6f\xe4\x56\xbe\x44\x23\xc1\xdd\x90\x48\x4b\xd2\x7f\x43\x28\xf0\x2b\x63\x35\x43\xb2\x60\xb4\xd0\x38\xae\xa9\x17\x2f\x9d\x47\x0b\x2c\xa9\x1d\xe5\xc7\x74\xd4\x7f\xfe\xe0\xf8\x6d\xc1\x18\x46\xda\xe6\x47\xc3\x28\x53\x8e\xd4\x1b\xe3\xf2\x5c\xbd\xb1\x84\xa8\x2a\x7e\x53\xec\x0d\x17\xac\x47\x28\x4f\x57\xba\xbf\x31\x3e\x2c\x38\xc3\x01\x1a\xf7\x9c\xbf\x75\x56\x83\xb1\x64\x33\x20\x1b\x3d\x33\x3f\x25\x4d\x33\xfd\xe0\x44\x46\xfb\x01\x99\x5a\xc0\x2e\xc3\x71\x5c\x21\x77\xfa\x01\x6d\x94\x25\x04\x12\xe2\x16\xac\x4a\x29\xb4\x0f\xc7\x38\xae\x0d\x64\xec\x13\xf8\x8c\x4e\x8d\x5e\xb2\x59\x7b\x59\x4b\xf9\xde\xc2\x8e\x1b\xf8\x73\x32\x0b\x31\xd6\x74\x2f\x3e\x27\xa3\x3e\xd4\xe9\x6a\x55\x48\x3a\xd3\x3c\x3d\xf3\x54\x8a\xf1\x74\x7e\xf1\x3d\xe6\xe6\xb1\xfb\x0e\xea\x1f\x45\xa9\x32\xb7\x33\x23\x3c\x2f\x4e\x18\x4a\x0d\xf6\xd3\x8d\x09\x4f\x51\x3a\x20\x85\x89\x2a\x3b\x57\x1b\xad\x4e\x93\x73\x5a\x00\x16\x0d\x78\x5b\xa2\x32\x51\xe1\x46\x54\x46\xa6\x61\xb5\x21\x96\xd1\x38\x46\xde\x73\xb4\x2c\xb1\x32\xef\x9c\x92\xf3\x8c\x2a\xa0\x0a\xb8\xa8\xa8\xa6\xac\xd9\x6d\x79\x26\x12\x81\x09\x71\x63\x08\x95\x08\xf9\x9f\x0c\x19\xa3\xc5\x87\x59\x24\x15\x69\x7f\x00\xda\x7c\x49\xf8\x37\x00\x9e\x35\x2b\xc0\xd8\x84\x89\x5d\xc7\x83\xfd\x57\x39\x8c\x50\x4e\x40\x09\x29\x27\xa6\xe3\x2e\x0d\x7e\xaa\xbe\x8d\x72\x6e\xd8\x84\xed\x77\x50\xa4\x6d\x58\xf1\x88\xed\x6c\xed\x03\xfb\xe9\x16\x92\xe6\x44\x4e\x2a\x7b\xd9\xa9\xe3\x87\x20\x9a\xb7\xd7\xa2\x06\x9b\xeb\xe9\x92\x5d\xcd\x99\xa8\x3b\x2f\x68\xdb\x7c\xa7\xbf\x2d\x91\x68\x04\x02\x1c\xc7\x35\x15\x02\x3f\xeb\x2e\x2c\xb3\xe8\xf5\x55\x2b\x9b\x6a\xb8\xbc\xac\xc9\x5e\x16\x87\xf9\x51\x2d\xb6\x12\xc9\xa2\x4e\x23\xee\xda\xaa\xb7\x42\xd0\xf4\xea\x24\x44\x66\x48\xda\x73\x54\xad\xaf\xd2\x6b\x4b\xd2\x66\xe0\x5b\x81\x95\xaf\x9a\x26\x80\x48\x24\x73\x4b\x99\xe8\x97\x82\x19\xc2\x8e\x55\xcf\x59\xab\xc9\xda\x64\x2e\xbf\x79\x69\xc5\xae\x97\x8d\xf3\x4f\xc1\x50\x5e\x18\x32\xda\x82\x1c\x65\x18\xdd\x84\xe2\xce\xa9\x0f\x98\x78\xa7\x25\x71\x15\x46\xa5\xac\x13\xb7\x9d\x71\xc0\x0a\x62\xfc\x2d\xeb\x54\xd2\xfd\x23\x72\x83\x40\x35\xd8\x21\x54\xea\xe0\x17\xe4\x91\x9c\x14\xda\x3c\x08\x31\x11\xd2\x64\x03\x6e\xf8\x03\x5a\xd4\xc9\x41\x8e\x50\xfe\xfa\xa8\x39\x1f\x03\x5f\x21\xaa\xe1\x87\x9a\x43\xa8\xb9\xab\xca\x28\x42\xa5\x9c\x1a\xa9\x2a\xc3\x9c\x6a\x07\x46\x84\x95\xd8\x73\x6a\x26\x3e\xa4\xa4\x17\xcb\x6c\xf2\x8d\x25\xbf\xc1\xcd\xf9\x00\x5d\x74\xc5\x6c\x9c\x4d\xf7\x33\x8d\xb3\x69\xf3\xf9\xd3\xe3\xac\x26\x67\x64\x41\xc6\x2f\x7e\x6a\x8c\x15\xcb\x0e\xb8\x12\xa5\x6c\x32\xfb\x98\xa8\x66\x5d\xdb\x45\x28\x2d\x24\xc6\x50\x2a\xeb\x67\x53\x0d\xbe\x9c\x1e\x6e\x2e\x99\x7a\x49\xeb\x93\x19\x3d\xc7\x67\x13\x38\xce\x23\xa1\x86\x24\x36\x7d\xde\xb2\x9b\x57\x90\x6a\x15\xca\xcf\x0c\x89\x42\x90\x98\x63\x1e\xa2\xac\x4a\x9a\x80\x54\x54\xf4\xad\x90\x99\xca\x26\x15\xb2\x04\x88\x02\x9d\x11\x73\x22\x2d\x59\x0c\x31\x2a\x2d\xc5\x64\x86\x70\x1e\x1c\x94\x4a\x43\x41\x94\x32\x11\xa1\x05\x98\x62\x25\x38\x02\x32\x85\x2f\xbe\x65\xa1\x1f\x60\xe2\x72\xc6\x7f\x38\x18\xfd\x74\x2a\x9e\x56\xf2\xd6\x2e\x0d\x72\xcf\xfb\xa9\x8c\x7c\x32\x53\x9e\x98\x8b\xa7\x4c\x79\x3c\x11\x2f\x1b\x1b\x1e\x61\x4d\x60\xfc\xce\xd3\xfe\x96\x36\x6d\x34\x15\x7c\xd3\x1c\xef\xec\x94\x89\x9b\x57\x12\xc1\xac\x05\x0a\x71\x1a\x21\x82\x47\xe8\xc1\x96\x02\x25\x04\x37\x44\x32\xed\x92\x44\x26\x48\xfc\x70\x9f\x31\xd3\x4d\x8d\x29\x63\x10\x22\xa4\x86\x3d\x4a\x40\x4e\x26\xa1\x39\x4f\x15\x13\xc3\x2c\x2e\xc6\x2b\xba\x0e\xf8\x2e\x52\x55\x63\x98\xf6\xbc\x75\xa7\x31\xbd\x50\x59\x94\x90\x62\x0c\xd3\xe3\xfa\x3f\x6e\x60\xa0\x3a\x6c\xcf\xb9\xf2\xb3\x18\xa3\x49\x28\xe1\x04\xa6\xcd\x7f\x73\xab\x91\x52\x9d\x95\xa1\xbd\xcb\x38\x2c\xef\x69\x82\xd2\x17\x5a\x39\xfd\x66\x30\x38\x3f\xb3\x07\xa2\x95\x70\x67\x06\xb3\xa8\x2d\x90\xe1\x49\x69\x9a\xb9\x5f\x38\x9a\x14\x4e\xe4\xc4\x1e\x87\xa6\x57\x95\xaf\x14\x1c\x90\x11\x39\xab\x2e\x5a\x0a\x56\xa6\x94\xab\x5f\x1f\x6e\x99\xbe\xe3\x06\x66\x78\x6b\x56\xf4\xdb\x5e\xbb\xe3\xbd\xad\x47\x2b\x6f\x5f\x96\xef\x5e\xae\x85\xda\x0d\x4f\x0f\xc3\xe3\x93\xfb\xc3\x82\xff\xb1\x7f\xb3\x1b\xcb\xcf\xa3\x0d\x36\x98\xbc\xdf\xee\x6c\x64\x9f\x4e\x7e\xfb\x5a\xbc\x23\x6f\x8f\xee\xb6\x4e\xfe\xc1\xdd\x8b\xeb\xc2\x3e\x8f\x58\x19\x23\x10\xc6\x20\x12\x79\x41\x19\xc6\x0d\x5a\xf8\x25\x44\x26\xc6\xbf\xbe\x01\x21\x81\xd6\x82\x94\xc7\x74\x44\xe3\x92\x30\x7b\xf3\xa6\x0c\x95\x39\x62\x8c\xf1\x0f\xd9\x46\x8f\xa9\xd6\x28\xdd\xe9\xf5\x6d\x7d\xab\x3b\x54\x0b\x37\xba\xdf\xb6\xd5\x97\xb5\xeb\x5d\xfc\x23\xd9\x3b\x1e\x0c\x49\xab\xb5\x71\x74\xb4\xd3\xbd\xfa\xb8\xdf\xda\xea\x0e\xce\xae\x07\x27\x21\xee\xad\xdf\x77\x92\xce\x55\x2b\x2b\xdf\x3d\xdd\x56\xdf\x01\x29\xa5\x61\x88\x92\xaa\xcc\x25\xa8\xfc\xb6\xd7\xf2\x5a\xf3\x73\x4f\xc4\xf3\xe9\xfd\x47\xff\xeb\xe9\x0e\x1d\x1e\xd3\xdf\xbb\x24\xbd\x18\x6a\x7a\xdf\xde\x3f\xbc\xa0\xdd\x48\xed\xc6\x7b\xf7\xef\xaf\x3e\xa6\x6b\xe7\xc3\x4b\x51\x1e\x3d\x09\xcf\x32\x20\x52\x14\x8c\x46\xf6\x7a\xc4\xec\x67\x0e\x7b\xe0\x57\x97\x54\x81\x5f\xfd\x6f\xc9\xf3\xff\x07\x00\x00\xff\xff\x89\x9c\xe8\x11\x3f\x19\x00\x00")

func frontendIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_frontendIndexHtml,
		"frontend/index.html",
	)
}

func frontendIndexHtml() (*asset, error) {
	bytes, err := frontendIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "frontend/index.html", size: 6463, mode: os.FileMode(436), modTime: time.Unix(1501786186, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"frontend/application.coffee": frontendApplicationCoffee,
	"frontend/application.js": frontendApplicationJs,
	"frontend/index.html": frontendIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"frontend": &bintree{nil, map[string]*bintree{
		"application.coffee": &bintree{frontendApplicationCoffee, map[string]*bintree{}},
		"application.js": &bintree{frontendApplicationJs, map[string]*bintree{}},
		"index.html": &bintree{frontendIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

