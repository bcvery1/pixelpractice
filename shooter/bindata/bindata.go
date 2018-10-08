// Code generated by go-bindata.
// sources:
// assets/rock.png
// assets/ship.png
// DO NOT EDIT!

package bindata

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

var _assetsRockPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x82\x08\x7d\xf7\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x35\x00\x00\x00\x34\x08\x06\x00\x00\x00\x2a\xba\x70\xd5\x00\x00\x00\x06\x62\x4b\x47\x44\x00\xff\x00\xff\x00\xff\xa0\xbd\xa7\x93\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x0b\x13\x00\x00\x0b\x13\x01\x00\x9a\x9c\x18\x00\x00\x00\x07\x74\x49\x4d\x45\x07\xe2\x0a\x05\x0d\x0d\x1e\xad\x4f\x28\xc7\x00\x00\x08\x0f\x49\x44\x41\x54\x68\xde\xed\x9a\x6f\x4c\x53\xfb\x19\xc7\x3f\xbf\xd3\x9e\x52\xa0\x05\x81\xda\x02\xa2\x68\x74\x64\x64\x78\x33\x03\xdc\xeb\x46\x87\x2e\xd1\x68\xae\xe6\xee\x85\xcb\xe2\xcc\xa2\xf7\xd5\xdd\x5e\x2c\xba\x17\x6e\x8b\xd9\x66\x32\x97\xc5\x97\x46\x47\x16\xc9\xc5\x17\x73\x59\xb6\x25\x5b\x72\x97\xdc\x17\xdb\x2b\x83\x11\x45\xa3\x28\x4c\x25\x32\xa0\xca\x9f\x52\x40\x28\x15\xae\x94\xb6\xe7\xfc\xf6\x42\x4e\x53\xca\xa1\xb4\xa5\x75\x6e\xd9\x93\x34\x6d\x4f\xcf\xe9\xef\x7c\x7f\xdf\xe7\xcf\xf7\x79\x5a\xf8\x1f\x34\xf1\x96\xd7\x90\xff\x95\xbb\xd4\xd1\xd1\x61\x39\x76\xec\xd8\x63\x21\xc4\x17\x80\x06\x48\x21\x84\x34\x9e\x2d\x16\x4b\x1f\xb0\xfb\x5d\x66\xea\x1b\xc0\x4f\x80\x4d\x36\x9b\x6d\xbb\xaa\xaa\x5b\x5e\xbf\x7e\x2d\xb6\x6f\xdf\xce\xb9\x73\xe7\x50\x55\x15\x21\xde\x2c\x11\x89\x44\xb8\x76\xed\x1a\x3d\x3d\x3d\x44\xa3\x51\x09\x0c\x02\x75\xef\x0a\xa8\xf7\x9d\x4e\xe7\xf7\x55\x55\xfd\x58\xd7\x75\xc5\xeb\xf5\xe2\xf5\x7a\x71\xbb\xdd\x08\x21\xe2\x20\x52\xd9\xad\x5b\xb7\xe8\xe8\xe8\x90\xcb\xee\xf8\x29\xf0\x83\xff\x24\xa8\x49\x21\xc4\xa6\x83\x07\x0f\xda\x8e\x1f\x3f\x8e\xa2\x28\x1b\x75\x55\xba\xba\xba\x70\x38\x1c\xd1\xf9\xf9\xf9\x46\xe0\x9f\x6f\x13\x94\x1d\x98\x53\x55\xb5\xe0\xfc\xf9\xf3\x6c\xdd\xba\x35\x67\xbb\x1a\x08\x04\xb8\x70\xe1\x02\xe1\x70\x18\x5d\xd7\x9f\x03\xfb\x81\x17\x6f\x03\xd4\x2b\x87\xc3\xe1\x6c\x6b\x6b\xcb\xb9\xff\xeb\xba\x8e\xae\xeb\x0c\x0f\x0f\x73\xe9\xd2\x25\xc2\xe1\x30\x05\x05\x05\x54\x57\x57\xfb\x2a\x2b\x2b\x7f\x7c\xfb\xf6\xed\xbf\xe6\x03\xd4\xaf\x80\x9f\x5d\xbe\x7c\x59\x94\x96\x96\xe6\x05\x54\x62\x1c\x46\xa3\x51\xda\xdb\xdb\xe9\xeb\xeb\x23\x16\x8b\x21\x84\xc0\x62\xb1\x50\x5a\x5a\xfa\xaf\x97\x2f\x5f\xd6\xe5\x0a\xd4\x42\x43\x43\x43\xf1\xd9\xb3\x67\x73\x0e\x48\xca\x37\x65\x4b\x08\x11\x7f\x1d\x5f\x74\x61\x81\x07\x0f\x1e\x10\x0c\x06\x19\x18\x18\xe0\xd9\xb3\x67\xc6\x06\x7c\x57\x4a\xf9\xa7\x54\xdf\xab\xa4\x11\x4b\xf6\xd3\xa7\x4f\xe7\x85\x21\x03\x50\xe2\xb3\x61\xc5\xc5\xc5\x00\x94\x95\x95\xb1\x6d\xdb\x36\x5a\x5b\x5b\x71\x3a\x9d\x00\x7f\x54\x14\xe5\x97\x1b\x01\xf5\x17\x8b\xc5\x12\xb5\xd9\x6c\xb9\xaf\x25\x09\xe9\x5f\x4a\x19\x67\x4a\x4a\x19\xff\xcc\x6a\xb5\x02\xe0\xf1\x78\x10\x42\xd0\xd8\xd8\x88\xdb\xed\x46\x4a\x79\x5e\x55\xd5\xef\x64\x0b\xaa\xa9\xaa\xaa\x2a\x90\x97\x02\x99\xa2\x9e\x19\x00\x77\xef\x7e\x23\x3c\x14\x45\xa1\xb6\xb6\x96\xea\xea\x6a\xea\xeb\xeb\x71\xbb\xdd\x68\x9a\xf6\x67\x55\x55\xb7\x67\x03\x6a\xd3\x81\x03\x07\x86\xf3\x15\x4b\xc9\xc7\x12\x59\x02\x70\x38\x1c\x2b\xce\xb1\xd9\x6c\x08\x21\xa8\xaf\xaf\xc7\x62\xb1\xe0\x74\x3a\x07\x32\x05\xf5\x75\x00\xaf\xd7\xfb\xcd\x5c\x02\xd2\x34\x6d\x15\x4b\x52\x4a\x14\x45\x89\x27\x8c\x44\x57\x4c\xb6\x8a\x8a\x0a\x00\x9a\x9b\x9b\x99\x9d\x9d\x55\x15\x45\xf9\x28\x13\x50\xbf\x71\xb9\x5c\x16\xab\xd5\x9a\x53\x25\x6f\xb1\x58\xd6\xcc\x7e\xc9\x32\xcb\xcc\x45\x1d\x0e\x07\x42\x08\x6c\x36\x1b\x76\xbb\x1d\x21\xc4\xb5\x74\x41\xd9\x2d\x16\xcb\x97\x8f\x1e\x3d\x6a\xcd\x47\x3c\x99\xa5\xf2\xc4\xd7\x66\x0c\x99\xd9\x8e\x1d\x3b\xd0\x75\xdd\x95\x2e\xa8\xcf\xec\x76\x7b\xd1\xfe\xfd\xfb\xc9\x57\x92\x90\x52\xc6\x0b\x6f\x32\x23\xc6\xfb\xc9\xc9\xc9\x94\x49\xc6\xe3\xf1\x18\x1b\x60\x5f\x17\x94\xd5\x6a\x3d\xd0\xd2\xd2\xb2\x21\x26\x92\x1f\xc9\x0c\x18\x80\x92\xcf\xd1\x34\x2d\xfe\x7a\x78\xd8\x3c\x47\x59\xad\xd6\xe4\xfa\xf6\xc3\xf5\x40\x9d\x90\x52\x5a\x4e\x9c\x38\x91\xb3\xcc\x96\x78\xdc\x00\x60\xc4\x56\x22\x4b\xc6\x71\x03\x6c\x2c\x16\x5b\x37\x2e\x55\x55\x05\xf8\x24\x25\x28\xbb\xdd\x7e\xb1\xbc\xbc\x7c\x43\x0c\x25\xbb\x89\x99\x8b\x99\xb1\x27\x84\x40\xd3\x34\x00\xa6\xa7\xa7\xd7\x5c\xc7\x50\x1b\x46\x5c\x01\xb5\x29\x41\x69\x9a\xb6\xad\xa9\xa9\x29\x6b\x86\x8c\x1d\xd6\x34\x2d\x2e\x85\x92\x55\xc4\x5a\xcd\xa4\x21\x5e\xc3\xe1\x30\xfd\xfd\xfd\x6b\xae\x95\xe8\x7e\x55\x55\x55\x08\x21\x54\x20\x7e\xd3\xab\xb2\x5b\x2c\x16\xe3\xf0\xe1\xc3\x19\x83\x1a\x1a\x1a\x62\x6c\x6c\xcc\xf4\x33\x55\x55\xa9\xa9\xa9\xa1\xaa\xaa\xca\x70\x97\x38\xb0\xe4\x38\x8b\x46\xa3\xdc\xbd\x7b\x37\xe5\x5a\xc1\x60\x70\xc5\x75\x8a\xa2\x08\x4d\xd3\x3e\x06\xee\xaf\x02\xd5\xdc\xdc\xfc\xdb\x87\x0f\x1f\x92\x4d\x8b\xe1\xf7\xfb\xd7\xfc\x2c\x1a\x8d\xe2\xf3\xf9\xf0\xf9\x7c\x2b\x35\x58\x53\x13\x45\x45\x45\x71\x76\xef\xdc\xb9\x93\x96\x47\x44\x22\x11\x33\x37\xdf\x64\xca\x54\x75\x75\xf5\x89\xde\xde\xde\xac\xe2\xa9\xb5\xb5\x95\xce\xce\xce\xb4\x6b\x0c\xc0\xfd\xfb\xf7\x33\x5e\x27\x16\x8b\xad\x5a\x43\x55\x55\x62\xb1\xd8\x57\x4d\x63\xea\xd1\xa3\x47\x25\x89\x41\x98\xa9\xed\xd9\xb3\x07\x97\xcb\x45\x41\x41\x41\x5e\xea\x9b\x94\x92\xf1\xf1\xf1\x55\xc7\xcb\xca\xca\x00\x76\x98\x32\x15\x0e\x87\xc9\x36\xf3\x01\x94\x94\x94\xd0\xd0\xd0\x10\x7f\xdf\xdd\xdd\x4d\x38\x1c\xce\x19\xa8\x68\x34\x6a\x7a\xbc\xae\xae\x0e\xbf\xdf\x6f\x33\x05\xf5\xea\xd5\x2b\x7d\xcb\x96\x2d\x96\x5c\xdd\xc4\xde\xbd\x7b\x91\x52\xb2\xb8\xb8\x48\x4f\x4f\xcf\x9a\x75\x27\x5d\x96\x66\x66\x66\xd6\xeb\x0d\xeb\x80\x01\x25\x29\x9d\x47\x8c\xec\x94\x4b\x2b\x2c\x2c\xc4\xeb\xf5\xb2\x6f\xdf\xbe\x78\xe3\x97\xa9\x22\x09\x85\x42\x2c\x2d\x2d\xad\x07\xea\x83\x55\x31\xa5\x69\xda\x5c\x2e\xdd\x25\xb1\x00\x1b\xc1\xdd\xd2\xd2\xb2\xaa\x4f\x4a\xc7\x96\x5b\xf9\x35\x8b\xf9\xf2\x0c\xf2\x43\x33\x50\xfd\x66\x81\x98\x4b\x11\x0b\xd0\xd8\xd8\x48\x65\x65\xe5\xba\xad\x7e\xe2\x8d\x1b\xdd\xef\xe6\xcd\x9b\x57\x29\x97\x84\x99\xc7\x67\x66\x8a\xe2\x52\x30\x18\x24\x5f\x96\x78\x93\x75\x75\x75\x19\xc7\x14\x40\x51\x51\x11\xc5\xc5\xc5\x66\x7d\x97\x06\x4c\x9a\x81\xfa\x1c\x48\x29\x51\x36\x02\x26\xd1\x15\x85\x10\x14\x15\x15\xa5\x7d\xad\x31\xe2\x16\x42\x90\x1c\xf7\x09\xac\x3d\x32\xd5\x7e\xd5\xd5\xd5\xfa\xd5\xab\x57\xc9\xa7\x25\xea\xb6\x74\xcf\x4f\xcc\x9c\x73\x73\x73\x2b\xdc\x2f\x10\x08\x00\x84\x81\x39\x53\x50\xbb\x76\xed\x7a\x7f\x61\x61\x81\xce\xce\xce\xbc\x75\xbc\x86\xa5\x8a\x2b\x33\x11\x6b\xf6\x1d\x42\x08\x26\x26\x26\xc2\xc0\xcc\x9a\x2a\xfd\xc6\x8d\x1b\x0f\x34\x4d\xfb\xf4\xfa\xf5\xeb\xb1\x50\x28\xa4\xe5\x73\xd6\x97\xd8\x17\xa5\x7b\xbd\xa1\xfc\x13\x7f\x71\x09\x85\x42\x0a\xf0\x8b\xf5\xda\xf9\x4f\x34\x4d\x9b\x3c\x73\xe6\xcc\xb4\x94\x52\xcf\xe7\x20\x33\x53\x9b\x9f\x9f\x5f\xf1\xfe\xf1\xe3\xc7\x3a\x10\x01\xae\xa7\x33\x4d\xda\x05\x94\x5d\xbc\x78\x71\x30\x5f\x6e\x98\x69\x4d\x94\x52\x32\x37\x37\x17\x67\xec\xf9\xf3\xe7\x4c\x4f\x4f\x47\x8d\x71\x5e\x3a\xa0\xc2\xc0\xcf\x07\x06\x06\x6a\x83\xc1\x60\x34\x1f\x09\x63\xad\x19\xc4\x5a\x96\xa8\x28\x0a\x0b\x0b\x19\x1d\x1d\xd5\x80\xdf\x91\xf4\x63\x5d\x3a\x3e\xe0\x77\x38\x1c\xb6\xb6\xb6\xb6\x8a\x5c\xb2\x24\xa5\xe4\xe6\xcd\x9b\x19\xb5\x1c\x63\x63\x63\x28\x8a\x82\x94\x92\xee\xee\x6e\x16\x17\x17\xbb\x81\xaf\x65\x32\xcc\x34\xec\xdb\x0b\x0b\x0b\xce\xb5\xba\xda\x6c\xb3\x5f\xa6\xbd\x94\xd1\x84\x4a\x29\x19\x1e\x1e\x66\x71\x71\x31\x68\x06\x28\x5d\x50\xb7\x81\x81\x2b\x57\xae\xe8\xb9\x04\x95\x3c\xbf\x48\x65\x13\x13\x13\xf1\x91\x5a\x34\x1a\xe5\xc5\x8b\x17\x31\x60\x5f\x2a\x65\x9b\x8e\x35\x4f\x4d\x4d\x45\xdb\xdb\xdb\xb3\x06\x94\x38\xe7\xf3\xfb\xfd\x69\x27\x89\x40\x20\xc0\xd2\xd2\x12\x42\x08\xa6\xa7\xa7\xe9\xea\xea\xd2\x81\x1f\x91\xe2\x47\xef\x74\x0b\x45\x0c\x08\xf8\xfd\xfe\x8f\xc2\xe1\x30\x56\xab\x15\x97\xcb\x95\x95\x8a\xe8\xed\xed\x4d\x39\xcf\x48\xb4\xf1\xf1\x71\x22\x91\x08\x42\x08\x9e\x3c\x79\x82\xcf\xe7\x8b\x01\x3f\x05\x2e\xa7\xba\x2e\x93\xea\xd7\x03\x7c\x30\x34\x34\xf4\x25\x97\xcb\xc5\xe4\xe4\x24\x53\x53\x53\x78\x3c\x9e\x75\xff\x7a\xa0\xeb\x3a\x73\x73\x73\xdc\xbb\x77\x2f\x6d\x86\x46\x46\x46\xd0\x34\x0d\x45\x51\xe8\xee\xee\x26\x14\x0a\x85\x80\x7a\xe0\xef\xeb\x6e\x60\x16\x9e\xd4\xa7\x28\xca\x57\x0e\x1d\x3a\xa4\x78\x3c\x9e\x38\x0b\x3b\x77\xee\xa4\xa6\xa6\x66\x45\xeb\x3d\x3a\x3a\xca\xc8\xc8\x48\xca\x36\xc4\x2c\xee\x46\x46\x46\x90\x52\x12\x08\x04\x18\x1c\x1c\x94\xb1\x58\x6c\x7c\x79\xae\x37\x99\x96\x57\x64\x19\x22\xdd\x40\xa3\xc3\xe1\x10\x47\x8e\x1c\xb1\x14\x16\x16\xe6\xac\x76\xf9\x7c\x3e\x9e\x3e\x7d\xca\xec\xec\x2c\xba\xae\xeb\x52\xca\xcf\x81\x6f\x65\xe4\xea\x1b\x58\xff\x10\xf0\x7b\x21\x44\x45\x49\x49\x89\x72\xf0\xe0\xc1\xac\x3a\xda\x38\xfd\x7d\x7d\xf4\xf7\xf7\xc7\xc7\xce\x91\x48\xe4\x0f\xc0\xf7\xb2\x92\x61\x39\xd8\xdc\xf7\x84\x10\x7f\x13\x42\xd4\x96\x97\x97\x53\x5f\x5f\x4f\x5d\x5d\x9d\x58\x4e\xd9\x52\x4a\x29\x12\xf5\x9e\x11\x7f\x33\x33\x33\x0c\x0e\x0e\x32\x31\x31\x41\x30\x18\x44\x51\x14\x4a\x4a\x4a\x02\x33\x33\x33\x7b\xd9\xe0\x3f\x5e\x72\x69\xef\x2d\xbb\xe5\x17\xcb\x02\xf3\x15\x30\x0b\xf4\x3a\x9d\xce\x60\x45\x45\x45\xc4\xed\x76\x47\x0b\x0a\x0a\x96\x84\x10\x11\x45\x51\xa4\xcd\x66\x93\xe5\xe5\xe5\xb2\xb9\xb9\x79\xec\xe4\xc9\x93\x2a\xef\xb8\xd9\x79\xf3\xcf\xb0\x5f\x03\xff\x00\xba\x96\x1f\x37\x80\xab\xcb\x41\xcf\xa9\x53\xa7\x04\xff\xb7\xf4\xec\xdf\xb2\xd1\xbc\x14\xf1\x01\xf0\x6c\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xc6\x9d\x11\x65\x82\x08\x00\x00")

func assetsRockPngBytes() ([]byte, error) {
	return bindataRead(
		_assetsRockPng,
		"assets/rock.png",
	)
}

func assetsRockPng() (*asset, error) {
	bytes, err := assetsRockPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/rock.png", size: 2178, mode: os.FileMode(438), modTime: time.Unix(1538745210, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsShipPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x87\x01\x78\xfe\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x05\x00\x00\x00\x0f\x08\x06\x00\x00\x00\x2c\xf4\x05\x83\x00\x00\x00\x06\x62\x4b\x47\x44\x00\xff\x00\xff\x00\xff\xa0\xbd\xa7\x93\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x2c\x83\x00\x00\x2c\x83\x01\x5b\xa5\x57\x36\x00\x00\x00\x07\x74\x49\x4d\x45\x07\xe2\x0a\x08\x10\x04\x1d\x03\x10\x2a\xca\x00\x00\x01\x14\x49\x44\x41\x54\x18\xd3\x45\xc8\xbd\x4b\x42\x51\x00\x40\xf1\x73\x3f\x9c\xa2\xa4\x45\x5b\x82\x27\x42\x0f\x69\x50\x70\x76\x4e\x84\xa2\xa9\xb5\x25\x50\x9a\xdc\xda\xa3\x7f\xa1\x41\x68\x28\x08\x5c\x12\x22\x82\x10\x91\x06\xc1\xa5\xad\x21\xa4\x48\xa7\x4a\x8a\x1e\x98\xde\x8a\xde\xf3\xde\x86\x8a\x7e\xdb\x39\x00\xec\xe6\xf3\xb3\x81\xef\x47\xfc\x92\xdb\x95\x0a\xc9\x54\xaa\xa0\x40\xed\xff\xdd\x9b\x74\x9a\xfb\x6c\xd6\x39\x70\xcf\x9e\x37\x06\x10\x7b\xe5\x32\x33\xa3\x91\x8d\x29\x25\x3e\x8c\xe1\x34\x0c\x85\x6e\xf5\x7a\x34\x1a\x0d\x01\xe0\x9c\xb3\xcb\x57\x57\xc8\x52\xa9\x84\xd6\x9a\x93\xb3\x73\x84\x10\xd2\x18\x83\x4c\x24\x12\x44\x51\x84\x5c\x58\x04\x20\xa6\x35\x52\x4a\xf9\x13\xc1\x10\x6b\x2d\x73\xf1\x38\xf2\x35\x08\x50\x4a\xb1\x5a\x5c\xc1\x39\x37\x1d\x0c\x06\xc8\xf9\x5a\x8d\x28\x8a\x68\xb6\x2f\x51\x4a\xa9\x97\x4e\x07\xb9\x6e\xad\x15\x42\x30\x36\x06\x80\xad\x6e\x77\xaa\xa5\x10\xe2\xa8\x5e\xbf\xfd\xec\xf7\x97\x8e\x9f\x1e\x82\x4d\x88\xeb\xb7\x30\x1c\xc6\xda\xed\x8b\x6b\x63\xec\x57\x32\xd9\xdc\xb0\x76\x4d\x1e\x78\xde\xdd\x3b\x54\x1f\x5b\xad\x4c\x76\x32\xa9\x1e\xfa\x7e\x48\x31\x93\x61\x27\x97\xa3\xc0\xbf\x6f\x26\xcf\x6f\x42\xde\x0e\x9f\x1f\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x62\x65\x99\xcf\x87\x01\x00\x00")

func assetsShipPngBytes() ([]byte, error) {
	return bindataRead(
		_assetsShipPng,
		"assets/ship.png",
	)
}

func assetsShipPng() (*asset, error) {
	bytes, err := assetsShipPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/ship.png", size: 391, mode: os.FileMode(438), modTime: time.Unix(1539014669, 0)}
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
	"assets/rock.png": assetsRockPng,
	"assets/ship.png": assetsShipPng,
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
	"assets": &bintree{nil, map[string]*bintree{
		"rock.png": &bintree{assetsRockPng, map[string]*bintree{}},
		"ship.png": &bintree{assetsShipPng, map[string]*bintree{}},
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
