// Code generated by go-bindata.
// sources:
// admin_http_assets/app.css
// admin_http_assets/app.js
// admin_http_assets/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _admin_http_assetsAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\xce\x4b\x8f\xb1\x4a\xce\xc9\x4f\xcc\x8e\xd5\x51\x00\x72\x74\xe1\xec\x94\xc4\x92\x44\x5d\x64\x81\x0a\x64\x9e\x1e\x8c\x0d\x64\x22\x24\x20\xe2\x19\x99\x29\xa9\x0a\xd5\x5c\x0a\x40\x90\x92\x59\x5c\x90\x93\x58\x69\xa5\x90\x97\x9f\x97\xaa\xa0\x98\x99\x5b\x90\x5f\x54\x92\x98\x57\x62\xcd\x55\x0b\x08\x00\x00\xff\xff\x9c\x9e\x21\xb0\x7a\x00\x00\x00")

func admin_http_assetsAppCssBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppCss,
		"admin_http_assets/app.css",
	)
}

func admin_http_assetsAppCss() (*asset, error) {
	bytes, err := admin_http_assetsAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.css", size: 122, mode: os.FileMode(420), modTime: time.Unix(1434641335, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _admin_http_assetsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x4b\x6f\xe3\x36\x10\xbe\xe7\x57\x0c\x04\x03\x92\xb0\x8a\x9c\xf6\x50\xa0\x32\x82\x36\x7d\x2c\xb0\x87\x00\xc5\x76\xdb\x1e\x5c\x2f\x40\x4b\x8c\x22\x84\x26\x05\x92\xf2\xc6\x70\xf5\xdf\x3b\xa4\x5e\x14\xed\xcd\xa6\x49\x7b\xab\x01\x43\x12\x39\x8f\x6f\xbe\x19\x8e\x46\x7b\x22\x81\xd4\x35\x5c\x03\xa7\x9f\x80\xf0\xb2\x61\x44\xa6\x3b\x51\x34\x8c\x46\x41\x4e\xe4\x56\xf0\x4b\x49\x19\x39\x5c\xf2\x32\x48\x60\x1d\xf0\xf2\x3d\x55\xa2\x91\x39\xc5\xc7\xa0\xa9\xd2\xad\x10\x5a\x69\x49\xea\x60\x13\xaf\x2e\x2e\xd0\x5a\x9a\x0b\xae\xa5\x60\x8c\xca\x28\xb8\x25\x15\xff\x51\x33\xab\xbb\x50\xb9\xa8\xad\xde\x42\x3a\x46\x16\xe8\x8f\x18\x89\xbb\x86\xe7\xba\x12\x3c\xea\x04\x13\x18\xc5\xf0\xd6\x0a\xc5\xc7\x0b\x80\x6e\x37\x25\x68\x5f\x2b\x84\xbe\xde\xac\x70\xd5\x84\xf2\x81\x6c\x19\xc5\x95\x51\x2f\x0a\x96\xda\xac\x2d\x83\x78\x90\xf9\x81\x91\xfc\x81\x55\x4a\x7b\x72\xdb\x61\x5d\x2d\xb3\x8a\x17\xf4\x71\x52\xb9\x29\x4b\x49\x4b\xa2\x85\xf4\x74\xc8\xb8\x71\xaa\xf4\x5e\x34\xda\xc7\x22\xcd\x1a\x8a\x3e\xd0\x03\x86\x7b\xc4\x4b\x06\xe1\xf7\x78\x09\x5b\x7c\x6c\x47\xdd\x9f\xa8\xd2\x15\x27\x86\x8b\xcf\x5b\x58\x16\x93\xd4\xdc\xfd\xc5\x44\xd2\x9e\xb0\xaa\xb8\x29\x0a\x34\x61\xa8\x5a\x7e\x5c\x7f\xcc\x36\x6f\xfe\xcc\xd6\x57\x97\xdf\x6e\xde\x2c\x96\x2b\x4f\xd4\xa2\xfe\x70\xa8\xa9\x15\x56\x94\x17\xd1\x0d\x63\x7f\xbd\xad\xa4\xd2\xf1\x2d\xd1\xf9\xfd\xd2\xd7\xa0\x25\x7d\x44\xe9\x68\xcc\x5e\x0c\x26\x4b\xe6\x27\xa9\x6e\x24\x1f\x1f\xcd\x0f\xd1\xeb\x6c\xca\x34\x9a\x68\x68\x3c\x93\x30\x3f\x43\x42\xa5\x7e\x37\xf6\xd1\xb4\x96\x0d\x5d\x79\x12\x5a\x1e\xe0\x08\xde\x22\xd8\x32\x46\x44\x3f\x3f\xd6\xbd\x69\x5f\xaf\x85\xdc\x44\x11\x9d\x71\x0a\x8e\xcb\x3b\xc2\xd4\x89\xcf\xd6\x7b\xee\xc3\xeb\xb5\x5c\xe9\x41\xb2\x35\x8b\x6d\x1c\x99\x93\xe1\xe5\xa4\x2c\xdf\x22\x09\xff\x36\x71\xd5\x1d\x74\x5b\x70\x7d\x0d\x81\x6a\x76\xc1\xb9\x38\x47\x07\xe7\x98\xf5\xa3\x9c\x9b\x24\xfb\xf2\xf5\x26\x7b\xd1\x13\x92\x9f\xa4\xad\x3f\xb3\x23\x03\x55\xf1\x18\x77\x38\xec\xb1\x4f\x4b\xaa\x27\x2a\x0b\xa2\x49\x3c\xa0\xec\x0d\xe8\xbe\x3b\x98\xbd\xce\x6b\x77\xe2\x5a\xdf\xcb\xdc\x2f\x96\x14\x26\xab\x6f\x91\x53\x2b\x88\x8e\xe6\x98\x64\xc8\x48\x69\x3a\xe3\x3b\xae\xa9\x44\x9a\xb2\x6f\xae\x12\xf8\x83\x54\x98\xab\xaf\xbe\xbe\xea\x1c\x0c\x0d\xab\x28\x66\x9d\xe4\x24\xef\xe7\x1b\x9b\x07\x24\x5d\x28\xb2\xa7\x51\x9c\xea\x7b\xca\xa7\x80\xf1\x80\xd7\x53\x5e\x1c\x8d\xa2\x78\x0d\x74\xc7\xd8\xc0\x8c\x25\x2e\x19\x1d\x53\x29\x4f\xfc\x4e\x11\x1c\x77\xaa\xcc\x00\x65\x52\xc3\x7a\x8a\x37\x42\xb6\x1b\x8f\xfd\x51\x4f\xd2\x9d\xd8\xd3\xf3\x24\x4d\xe9\x36\x05\x89\xaf\x98\xbb\x4a\xee\xa2\xf0\x46\x52\x38\x88\x06\x54\xd3\xdf\x7c\x22\x5c\x83\x16\x50\x50\x46\xb1\x01\x4f\x2d\x1a\x28\x37\x4d\x83\x8b\x14\x42\x78\x03\xc6\xde\x04\x7c\xf2\x99\x76\x8a\xd1\x31\xb4\x1d\x35\xcc\x50\xf0\x69\x2e\xfc\x12\x42\xb2\x87\xe6\x6f\xdb\x91\xb9\x8f\x8e\xbf\x54\xf8\x7a\x41\xd2\x6d\xd1\x27\xf0\x6b\x2d\x04\xcb\xec\x69\x49\xc0\xdb\xeb\xb3\x63\xba\x2f\x36\x5f\xdb\x75\x03\xbf\x92\x06\x0f\x2f\x28\x22\xab\xfa\x0f\xcb\xe8\xbf\x0a\xe8\xff\xfa\x7a\x6e\x7d\x8d\x7b\x26\x6f\x2e\x70\x3b\x12\x7c\x29\xfb\x5d\xce\x6d\xca\x8f\x38\x6b\x58\x9d\x64\x56\x3c\x73\xcf\x98\x81\x0e\x91\x97\x87\x67\x67\xe0\x5c\x6f\xed\xe8\x77\x07\xb0\xd7\xb2\x3f\x0e\x6d\x4f\x92\x3f\x7a\x7c\xfd\xd9\xee\x42\x38\x39\x7c\x38\x8e\x7d\x09\xfe\x77\xa1\x83\xa8\x4b\xc7\x88\xc6\x4c\x80\x66\xa4\x7b\x09\x96\xf9\xac\xe8\x22\x4a\xe0\x19\xac\xce\x60\xb9\x83\xe7\x29\xba\x04\x5e\x4c\x1b\xfe\xb9\x4f\x1a\xd8\xa4\xf7\xbe\xcd\xc0\x67\x27\xfc\x77\x5c\x69\xc2\x73\x3b\x39\xdb\x05\xab\x1b\x0d\x08\x35\xdd\xd5\x8c\x68\xfa\x9b\xc4\x5e\x13\x36\x35\xd6\x5c\x97\x8d\x5b\x2b\x7b\xaf\x77\x2c\xec\x2b\x17\x10\xf2\x56\x10\x59\x0c\x6d\xa8\x5f\x9e\xbe\x4e\x32\x07\xcb\xf8\xc9\x31\x43\xd1\x9f\x14\x77\xe0\x19\xc8\xef\xa3\xb1\xd7\x95\xbf\x2b\x1e\x66\x81\xce\x07\xa6\xb9\x8b\x34\x67\x42\xd1\xc8\x35\xeb\x0c\xad\xed\x89\xe9\xdc\x28\xb1\xe7\x9b\x2f\x2a\xb5\xab\x94\x8a\xc2\x4e\x31\x3c\x67\x7c\x38\xec\x60\x5d\x64\xbd\xab\x64\x1c\x46\x95\x60\x7b\x5c\x9e\xbc\x58\x98\xd9\x67\x21\xf4\xf3\xdd\xf0\x3d\x89\xc6\x0e\xb3\x00\xd5\x1a\x53\xbf\x71\x91\x5c\xb8\xd7\xb6\x9b\xbf\x60\x5e\x11\x58\xef\xaa\x61\x7a\xfe\xc2\x82\xc8\x4b\x90\xd3\x24\x23\x97\xcc\xa9\x1d\xb5\xc6\xf1\xdf\x01\x00\x00\xff\xff\x9f\xe6\x84\xf7\xf9\x0e\x00\x00")

func admin_http_assetsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsAppJs,
		"admin_http_assets/app.js",
	)
}

func admin_http_assetsAppJs() (*asset, error) {
	bytes, err := admin_http_assetsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/app.js", size: 3833, mode: os.FileMode(420), modTime: time.Unix(1434641335, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _admin_http_assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x1b\xed\x6e\xdc\x36\xf2\x7f\x9e\x82\x11\x0a\x38\xc1\x55\x92\xed\xb4\x68\x6b\xec\x2e\x90\xba\xc9\x35\x77\xcd\xc5\x70\x52\xdc\x1d\x82\x1c\xc0\x15\xb9\x2b\xda\x94\xa8\x90\xd4\xda\xbe\x20\xef\x7e\x33\x94\xb4\xd2\x6a\x25\xad\xb6\xfe\xea\xb5\x69\x2c\x92\xc3\xe1\x7c\xcf\x70\xcc\x4e\x9e\xfe\xf2\xee\xf4\xc3\xbf\xcf\x5e\x91\xd8\x26\x72\xf6\x64\x82\x3f\x88\xa4\xe9\x72\xea\xf1\xd4\x23\xe9\xd2\xa7\x59\x36\xf5\x22\xaa\xe7\x2a\xf5\x35\x97\xf4\xc6\x4f\x97\x1e\x42\x72\xca\x66\x4f\x08\x99\x24\xdc\x52\x12\xc5\x54\x1b\x6e\xa7\x5e\x6e\x17\xfe\x8f\x5e\xbd\x10\x5b\x9b\xf9\xfc\x73\x2e\x56\x53\xef\x5f\xfe\xef\x2f\xfd\x53\x95\x64\xd4\x8a\xb9\xe4\x1e\x89\x54\x6a\x79\x0a\xbb\xde\xbc\x9a\x72\xb6\xe4\x8d\x7d\x29\x4d\xf8\xd4\x5b\x09\x7e\x95\x29\x6d\x1b\xa0\x57\x82\xd9\x78\xca\xf8\x4a\x44\xdc\x77\x83\x6f\x89\x48\x85\x15\x54\xfa\x26\xa2\x92\x4f\x8f\xb6\xd0\x30\x6e\x22\x2d\x32\x2b\x54\xda\xc0\xb4\x05\x46\x73\x1b\x2b\xbd\x05\x21\x45\x7a\x49\x80\xf5\xa9\x27\x22\x44\x10\x6b\xbe\x98\x7a\x41\x10\xc2\x9f\x05\x5d\xe1\x64\x00\x7f\x01\x30\x42\x5b\x61\x25\x9f\x9d\x16\x02\x3b\x77\x02\xfb\xc7\x5f\x09\x65\x89\x48\x27\x61\xb1\xe8\xe0\x9e\xfa\x3e\xf9\x8d\x5a\x6e\x2c\x9c\x97\x64\x42\x72\x46\x68\xca\x08\xc0\x89\x85\x80\xc1\xe9\xfb\xf7\xc4\xf7\x5b\x14\x18\x7b\x23\xb9\x89\x39\xb7\x15\x1d\x61\x98\xd0\xeb\x88\xa5\xc1\x5c\x29\x6b\xac\xa6\x19\x0e\x00\x65\xb8\x9e\x08\x5f\x04\xc7\xc1\x61\x18\x19\x53\xcf\x05\x70\x4e\x00\x33\x5e\x4d\xcd\x3b\x27\x20\x2a\x89\x8d\x79\xc2\xef\xf1\x6c\xdf\x1d\xd0\xa2\x60\xf0\x1c\x30\xc2\x16\xb1\xbf\x7e\x78\xfb\xdb\xf7\xc4\xc4\x22\x71\x52\x3b\xe7\x26\x53\x29\x0b\x2e\x0c\x79\xf3\xea\x47\x62\xf2\x0c\xcd\x86\xa8\x45\x09\xc8\x25\x9c\x98\x5a\x53\x88\x98\x33\x41\xc9\xe7\x9c\x6b\xc1\x4d\xc5\x27\x20\xfd\x28\x16\x44\x5a\x40\x40\x7e\xfa\xe4\xe6\x0a\xab\x21\x46\x47\x53\x0f\x0d\xd9\x9c\x84\xa1\x32\x26\x28\xb9\x46\x46\xd1\x61\xbe\x07\x32\x56\xc0\xe8\x0f\xc1\x71\x3d\x76\xec\x5d\x00\xc9\x93\xb0\x40\x33\x16\xa3\x2e\x58\x09\x8f\x82\xef\x00\x5f\x39\xea\xc6\xf6\xf4\x23\x4f\x99\x58\x7c\x42\x16\x26\x61\xe1\x91\x4f\x26\x73\xc5\x6e\x88\x56\x12\x0d\x5f\x45\x39\xf2\xed\x91\x5a\x72\xaf\xc5\x35\x58\x57\x4a\x57\x73\xaa\x2b\xe6\x99\x58\x91\x48\x52\x63\xa6\x5e\xb9\x50\xfc\xf0\x45\xba\xe2\xe0\xd8\x5e\x89\x0f\x66\xc5\x92\x3a\x3f\xc2\x7d\x9b\x3b\xd1\x6d\xa8\x48\xb9\x2e\xd7\xba\xf0\xfa\x48\x64\x03\x02\x60\x68\x0b\x62\xae\x41\x47\x6b\xcd\x7b\x6d\x57\x9a\x84\x74\x8d\x3e\x04\xfc\x03\x67\x45\x4a\x4a\x9a\x19\x4e\xaa\x8f\xe6\xb1\xb9\x6c\x40\x57\xec\xc2\x8f\x06\x8c\xb3\xca\x0a\x8a\x46\x56\xac\xf8\xc6\xaa\x23\x7e\x4d\xe7\xaf\x2a\xe1\x0d\xe2\x0a\x02\xa5\x68\xa1\xeb\xdb\x3f\xa7\xec\x2d\xb7\x5a\x44\x26\x3c\xfe\x2e\x06\x55\xa3\x88\x7f\xa6\x68\xac\x6e\x76\x10\xf3\x24\xcc\x65\xb7\x50\x40\xe1\x61\x00\x5c\xd5\xb2\x00\x8d\xd7\x30\xe5\xc7\x93\x3e\x45\x96\x6a\x4f\x60\xe4\x92\x01\xae\xc0\x94\xe4\x7a\xea\xbd\x85\xc9\x53\x2b\x1b\x86\xd0\xa5\x0a\xad\xae\x8a\x9d\x52\xd1\xcb\xa6\xd6\x01\x87\xc5\x05\xcd\x33\x4e\x21\xd6\x16\x13\x22\x25\xee\x03\x2c\xfd\xcb\x17\xf7\x15\x24\x66\xf9\xf5\x2b\xb0\x8f\x83\x06\x82\x0d\x7a\xa5\x9f\x30\xff\xe8\xb8\xad\x9d\xf8\x68\xf6\xb3\xa4\xd1\xa5\x14\xc6\x82\x7f\x1c\xb5\x96\x2d\x85\x2c\x54\x21\x29\x06\xee\x6f\x64\x93\xf1\xd4\x70\xd6\x42\x88\x7b\xaa\xbc\xd7\x9e\xd7\xdb\x93\x0e\x7c\xf6\x96\xda\x28\x26\x19\xa8\x59\x5c\x43\x06\x88\x87\xe1\x4c\x3e\x87\x18\x29\xd2\xe5\x6e\x50\xcd\x97\x7c\x10\xe3\xcb\x08\x3d\xd5\x74\x83\xc0\xec\x16\xc9\x08\xd9\xc1\xde\xc4\xba\x90\xd2\x50\xd6\x1c\x15\xe5\x64\x15\xcc\x2b\x01\x6f\xc9\xaa\x57\x2a\xb8\xc0\x40\xbf\xf3\xa0\x90\x0a\xaa\xd7\x76\x08\xb5\x09\xb9\x96\xcb\x18\x60\x27\x99\x1d\x80\x6b\xbd\xf3\x6b\xeb\x47\x10\x23\x31\x2a\x81\x3f\x3a\x5b\x15\xd1\x25\x98\x2e\x4f\xd4\x8a\xaf\x0d\xe8\xd9\x37\x02\xac\xe2\xfa\x39\x40\xad\x83\xc2\x52\xde\x64\x31\x16\x00\x64\xfd\xe5\x17\xdb\xfc\x48\xe8\x08\x6a\x9c\x70\x86\x9e\xdb\x4d\x49\x8f\x0a\x50\xd8\x2d\x4b\x0d\x9d\xac\xb7\xad\xfb\xe5\x72\x09\xbc\x52\xab\xb4\xe9\xb0\xef\x85\xd2\x49\x55\xda\x2c\x97\xaf\x61\xe4\x55\x84\xe3\x92\x4f\x19\x23\xee\x63\xa9\x55\x9e\x91\x98\x1a\x7f\xc1\x39\x9b\x03\xc3\x95\xe7\x2f\xdc\x26\x90\x09\xc8\x3f\x11\xe8\xa6\x8c\xd5\x87\x3e\x7b\x0e\x6b\x6a\x45\xa5\x60\x50\xc9\x3c\x8a\x77\xbd\xce\x53\x67\xe5\x43\x7e\x70\xbe\xcb\x51\xde\xe5\x36\xcb\xad\x93\x05\xb5\x43\x80\x6f\xd0\x4c\x80\x5f\xf2\xcc\xf0\xe8\xf9\x10\xe4\x3f\xa9\xb0\xbb\xa1\xee\xcf\x43\x69\xed\xa1\xb4\x36\x92\xfd\x7d\x94\x06\x8b\x3c\xdd\xed\x73\x74\x94\xcf\x39\x40\x10\xf6\xeb\xc4\x8e\x81\xac\xa4\x3d\x06\x16\xe5\x7d\x47\x1e\xdf\xb0\xef\x5b\xbb\xfc\x28\xa5\xf6\x78\x7d\x41\x76\xdf\x7c\xbf\xd6\x36\x7c\xbc\xc3\xb5\xbb\x37\xba\xcd\x22\x45\x3f\x00\x81\x24\x8a\x61\x29\x9e\xf2\x2b\x90\x46\x00\x4e\xe6\x95\x81\x64\x81\x9f\xcd\x03\xca\x92\xc0\x23\x19\x04\x4a\x1e\x2b\xc9\xb0\x38\xa8\xdc\x12\xe2\x08\x5e\xfe\x34\x56\x9c\x4b\x1f\x2e\x7e\x20\xf4\x14\xee\x75\x18\x32\x00\x33\x82\x0d\xd1\x83\x49\x1e\x83\x4f\xac\xae\xd6\x31\x0c\x0d\x32\x00\xc5\x38\x1c\x03\x9b\x1d\x02\x93\xd1\xb4\x07\x03\xd7\x5a\xe9\xa0\x24\xc9\x9b\xbd\xba\xce\x78\x64\xf1\x12\x56\xaa\x1f\xa8\x27\x8b\x92\x8d\x13\xc8\xcb\x09\x51\x9a\xd0\x15\x24\x66\x44\x3a\x40\x74\xb3\x00\x6b\x2f\x8d\x30\xcf\xbb\x53\x9b\x8b\x7b\x95\xe2\x74\x31\xd8\xad\xba\x12\x70\x4b\x5d\x05\xb6\x7d\x95\xe5\xb0\xdd\x52\x5d\x25\x8e\x3e\x85\x39\xcc\x58\x0d\xe5\x12\x2e\x2e\xfc\x1a\x8a\x0a\x63\x5c\x4a\xf8\x7f\x51\x54\x11\x12\x2b\x4d\xa9\xdc\x2e\x70\xb4\x5b\x55\x00\x59\x3b\xd8\x63\x70\x62\x6f\x32\xbc\x15\xe6\xc9\x1c\x2f\x0b\x5b\x7c\x55\x01\xbc\xe2\x4c\xac\xc7\x9d\xbc\x35\x4d\x2e\xfc\xcf\xc7\x43\xff\xa7\x4f\x7f\xf9\x26\xdc\xdb\xe4\xaa\x53\x6e\x69\x75\x35\x9a\xde\x48\x41\x0a\xce\xc9\x33\xc8\xb5\x90\xea\xa1\xae\x31\xcf\x89\x55\x04\xca\x1b\x65\x39\xf9\x9c\xd3\xd4\x8a\xff\x42\xd1\x4a\x2a\x64\x7f\x5a\x9b\xdc\xa1\x49\x4c\xaf\x95\x16\xaf\xdc\xf7\x7d\x6a\x10\x4f\xb8\xa5\xf6\x0a\x14\x7f\x54\x73\x09\xbd\x26\x88\xe1\xbe\xd4\xd5\xb3\x30\xcf\xad\x85\xac\x53\x8a\x76\x6e\x53\x02\xff\xf9\x26\x71\x3f\x32\x2d\x12\xaa\x6f\xdc\xf7\x5c\x2a\xac\xd1\x0b\xa5\x15\xa5\xb9\x13\x3e\x13\x06\x4b\x3e\x56\xcb\xa1\x96\xe2\x4b\xc6\x26\x61\x71\xc2\x3e\x14\xef\x51\xb0\x74\x5f\x53\x42\x34\x91\xed\xab\xcb\x39\x04\x2f\x77\xcd\x1d\xba\xb6\x80\x21\x5b\x7e\x67\x17\x17\x3c\x92\x3f\xf2\x9d\xa5\xa8\xfb\x5d\xc1\x5f\xb4\xe3\xae\xf8\x81\x94\x04\x5d\x10\x4b\x49\x43\x62\xae\xab\xde\x6b\xd7\x4e\xc7\x03\xb9\xe4\x37\x83\x17\x1e\x07\x84\xd6\xb1\xbb\x8b\xf0\x28\x8d\x09\xc6\x30\x47\x0f\x81\xbc\xcf\x94\x92\x43\x00\x67\x50\xb1\xcb\x21\xfe\xb0\xeb\x87\xde\x3b\x3d\x26\xf7\x78\xcf\xd2\xf5\x3d\xcb\x19\xeb\x1e\x57\x2c\x8c\x03\x45\xf4\x1a\xb8\x5a\x40\xbe\xbf\xf1\x08\xd5\x82\xfa\xb1\x60\x60\x83\x60\x97\x3a\xe7\xae\x11\x8c\xa1\xa9\x2f\xcc\x34\x72\x82\x48\x17\x0a\xdb\x69\x3a\x00\xb3\xe9\xbf\x23\x75\xee\x40\x1b\xda\x73\x4b\x82\x36\xc0\xf5\xce\xde\xce\xe0\xe6\x11\xed\x9e\xc1\xfd\x3b\x6e\xa3\xed\xbd\x6b\x5b\xf1\x5e\x78\xa3\x45\x5a\x6f\x3a\xee\xbc\x47\x16\xe1\xc6\x49\xfd\x9e\x9b\x46\xce\xc4\x9a\x56\xc9\xd0\x2a\x75\xc0\xb8\x81\x18\xdb\xec\xd7\x77\xf0\x34\xc3\x30\xb4\xd3\x0e\x41\xaa\x2b\x8d\x84\x8a\x65\x6c\x87\x0c\x12\x42\xd7\x70\xea\xbb\xc5\x6a\x21\x62\x47\xe3\x97\x03\x46\xd3\x25\xd7\x07\x27\xe4\x29\x0b\x54\x2a\x45\xca\xbf\x25\x07\xa8\x18\x98\xaa\x66\xbe\xa2\x55\xb0\xd1\x26\x79\x27\x87\x8c\xed\x54\xde\xf6\x9c\x51\x4d\xce\x3f\x78\x06\x2d\x02\xf4\xdd\x63\xef\xec\xc1\x74\xe3\x77\x06\xb8\xae\xeb\x58\x60\x30\x23\x78\x43\x36\x1a\x33\x06\xae\xd3\xe5\xb9\x03\x1e\xdd\xe0\x81\x54\x4c\x90\xfb\xe3\x22\x73\x69\x6b\x90\x0d\x75\xe9\x1b\xb1\x4c\x1f\x8a\x95\xe1\x20\xf6\x4b\x1d\x42\x8a\x50\xf6\xed\xe3\xb5\xc1\xf7\x6e\x87\x3d\x54\x86\xdd\xef\xd6\xd5\xd5\x05\x70\xb9\x22\xf8\x3b\xbf\xa9\x2e\x59\x1f\x20\xf1\xf6\x5c\xb2\xd6\x1d\xb5\xcd\xb6\x4d\x55\x18\x76\x9e\xfa\x80\x4c\x14\x94\xef\xe4\x62\x83\x78\xc3\x53\xf6\x52\x4a\x57\x3e\x0e\xb5\x0c\xdd\x09\x0e\x69\x0f\x55\x1b\xd7\xc9\xf5\xed\xc1\xd1\xb4\xfb\x42\xd9\xba\x4a\xb6\xb7\xf7\x5d\x26\x9b\xc4\x63\x8f\x10\xc7\xaf\x85\x36\xd6\xcd\x0c\x5d\x21\x7b\xaf\x8f\x0f\xa8\xae\x33\x97\x13\x2b\x85\x55\xa3\xdd\x2a\x2b\x21\x1f\x99\xfa\xf7\x55\xb2\xad\x18\x68\x4c\xec\xe6\xa1\x06\x7e\x64\x36\x36\x3a\xb5\xe7\x63\x3b\xb5\xe7\x7b\x77\x6a\x1f\x90\xa5\xf2\x8e\x57\x31\xb5\x1e\xee\x11\xd2\x6a\x14\x5b\xbf\x3a\x28\x57\x1e\x92\xc7\xa2\xd9\x02\xa5\x57\x74\x39\x57\xd7\x5e\xa7\x35\x16\x25\x4a\x69\x89\x1b\xf5\xca\x06\xbf\x7f\x32\xb2\xcf\xca\x9a\xa4\x0c\x01\x9b\x15\xca\x48\xc2\x7b\xb2\xef\xbd\xf7\xb9\xea\x20\x3d\xba\xd3\x75\xf7\xbf\x8c\x6f\xf7\xb8\x36\xdf\xe0\xd4\x83\xc6\x73\x9b\xf5\x0b\x1c\xec\x01\x85\xeb\x17\x37\xae\xf3\x53\x4d\xff\x5c\x3d\x98\x83\x22\x4d\x73\xf2\x37\xba\xa2\xef\xdd\xe3\x2f\x87\x6c\xba\xf7\x3f\x8d\xa7\x6e\xe4\x0c\x1d\x8d\x11\x6a\xf1\xb5\x1f\x81\x8c\x85\x6f\xe5\xf0\xb3\x7a\x32\x46\x8c\x72\xe3\x8c\x2e\xb9\x21\x52\x51\x46\x16\xd4\x58\xbe\x7e\x33\xd6\xf5\x94\x8d\x5e\xd0\xeb\x60\xa9\xd4\x52\x72\x9a\x09\xe3\xde\xb3\xe1\x5c\x28\xc5\xdc\x84\x17\xf8\xe4\xee\x26\x3c\x0a\x8e\xe0\x4f\x39\xda\xfd\x4c\x6e\xfc\x23\xc3\x8b\xf6\xfb\xc6\x4d\xbc\x5d\x44\x03\xcd\x11\x78\x43\x00\xc5\x32\xfe\xf6\xe8\xc2\x04\x4a\x2f\x81\xc2\xe3\xe0\xe8\x30\x2c\x27\xc7\xbd\xe4\x1b\x83\x09\x8a\x62\xa3\x72\x1d\xf1\x31\x5c\x03\x97\x80\x24\x92\x2a\x67\x0b\xd8\xcb\x5b\xc2\xac\x50\xe6\xc2\xaf\xe5\x70\x88\xa2\x3d\x0c\x9b\x73\xbe\xcd\xa4\xd9\x21\x09\xf7\xbe\xb2\x8f\x9c\xc2\xf9\xf0\xaa\x13\x82\xe7\x59\x9e\x40\x8c\xb6\x10\x1f\x04\x78\x5f\x9e\x61\x13\xd5\x45\x91\xb7\x8a\x51\x19\xe0\xf3\xc7\x8e\xa7\x81\x09\x2e\xb6\xdf\xfe\x4d\xe2\x17\x9b\xeb\xee\x75\xac\x37\xfb\xdd\x21\x25\xce\xb3\x4f\xc8\x97\x2f\xee\xa3\xea\x5d\xc5\x2f\x36\x5c\xc9\x7d\x36\x1a\xc6\x8b\x5b\xf4\x8a\xb7\x9a\xc2\xdb\x1c\x60\x34\x68\x3e\x22\x6c\x00\xd4\x67\xb4\x5e\x10\xd2\x39\x97\x48\xc1\xd4\xc3\x1c\xe6\xcd\xce\x8a\x4c\x36\x09\xdd\xca\x06\x6c\x3b\x99\x16\x9c\xe3\x86\x2a\x38\x3b\x14\x63\xaa\xb3\xb2\x4a\x1d\x53\x1c\x6c\x16\xcd\x88\xd4\x15\xb9\x7d\xf5\x72\xab\x4a\x6e\xc0\x0f\xff\x82\x76\xe8\x17\xb3\xad\x42\xb8\x3d\xdc\x4f\xcc\xd8\xbc\xf0\xea\x1e\xf3\x58\x31\xe3\x86\x4a\xcc\x0e\xc5\x08\x31\x1f\x1d\xff\x10\x1c\xc2\xbf\x47\x27\xc7\x87\x87\xdf\x0d\x3e\x76\xe8\xa8\x58\x3a\x04\x8f\x07\xef\x23\xf8\x02\xbe\x47\xf0\x27\xe4\x20\x56\xc6\x9e\xe0\x53\xe8\x83\xfd\x84\xde\xfb\xb4\xb6\xf0\x83\x05\x44\x97\xcd\x57\xbc\x1d\x29\x1d\x13\x38\xe3\x0b\x9a\x4b\xeb\x35\xfa\x0a\x11\x4d\x23\x2e\x9f\x3d\xf7\x66\xa7\x52\x19\xbe\x9d\xa9\x7b\xca\x83\xb2\x2e\x68\xe5\xff\xc5\x46\xea\x6f\x1c\xa3\x2e\xf1\x88\x22\x96\xb4\xcf\xd8\x48\xc9\x55\xea\xae\x63\x1f\x80\xbb\xac\x3f\x09\x8b\xff\x2b\xe2\x7f\x01\x00\x00\xff\xff\x4c\x74\xd8\x75\x26\x31\x00\x00")

func admin_http_assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_admin_http_assetsIndexHtml,
		"admin_http_assets/index.html",
	)
}

func admin_http_assetsIndexHtml() (*asset, error) {
	bytes, err := admin_http_assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "admin_http_assets/index.html", size: 12582, mode: os.FileMode(420), modTime: time.Unix(1434641335, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"admin_http_assets/app.css": admin_http_assetsAppCss,
	"admin_http_assets/app.js": admin_http_assetsAppJs,
	"admin_http_assets/index.html": admin_http_assetsIndexHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"admin_http_assets": &bintree{nil, map[string]*bintree{
		"app.css": &bintree{admin_http_assetsAppCss, map[string]*bintree{
		}},
		"app.js": &bintree{admin_http_assetsAppJs, map[string]*bintree{
		}},
		"index.html": &bintree{admin_http_assetsIndexHtml, map[string]*bintree{
		}},
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
