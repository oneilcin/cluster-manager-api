// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x73\xdb\x36\x12\x7f\xe7\xa7\xd8\xd1\xcb\x39\x37\xb6\x98\x38\xcd\x5d\x47\xaa\x6f\xaa\xca\x69\xa2\x49\x62\x7b\x42\xd7\x9e\x3c\x69\x20\x70\x45\xe1\x0c\x02\x2c\x00\x4a\x66\x32\xfe\xee\x37\xf8\x43\x8a\xa4\xa8\xa4\xd7\xf4\xa1\x9a\x49\x4c\x02\xbb\x3f\xee\xfe\xf6\x0f\x96\x8c\x63\x98\xcb\xa2\x52\x2c\xdb\x18\x38\x7f\xfe\xe2\x47\x48\x48\xae\x4b\x91\x41\x72\x99\xc0\x9c\xcb\x32\x85\x2b\x62\xd8\x16\x61\x2e\xf3\xa2\x34\x4c\x64\x70\x8b\x24\x07\x52\x9a\x8d\x54\x7a\x1c\xc5\x71\x14\xc7\xf0\x9e\x51\x14\x1a\x53\x28\x45\x8a\x0a\xcc\x06\x61\x56\x10\xba\xc1\x7a\xe7\x14\xee\x50\x69\x26\x05\x9c\x8f\x9f\xc3\x89\x15\x18\x85\xad\xd1\xb3\xa9\x85\xa8\x64\x09\x39\xa9\x40\x48\x03\xa5\x46\x30\x1b\xa6\x61\xcd\x38\x02\x3e\x52\x2c\x0c\x30\x01\x54\xe6\x05\x67\x44\x50\x84\x1d\x33\x1b\xf7\x9c\x80\x62\x2d\x81\x4f\x01\x43\xae\x0c\x61\x02\x08\x50\x59\x54\x20\xd7\x6d\x41\x20\x26\x18\x6d\x7f\x1b\x63\x8a\x49\x1c\xef\x76\xbb\x31\x71\x06\x8f\xa5\xca\x62\xee\x45\x75\xfc\x7e\x31\x7f\x7d\x95\xbc\x3e\x3b\x1f\x3f\x0f\x4a\xbf\x09\x8e\x5a\x83\xc2\xdf\x4b\xa6\x30\x85\x55\x05\xa4\x28\x38\xa3\x64\xc5\x11\x38\xd9\x81\x54\x40\x32\x85\x98\x82\x91\xd6\xe8\x9d\x62\x96\xb7\x53\xd0\x72\x6d\x76\x44\xa1\x85\x49\x99\x36\x8a\xad\x4a\xd3\xe1\xac\x36\x91\xe9\x8e\x80\x14\x40\x04\x8c\x66\x09\x2c\x92\x11\xfc\x32\x4b\x16\xc9\xa9\x05\xb9\x5f\xdc\xbe\xbd\xfe\xed\x16\xee\x67\x1f\x3f\xce\xae\x6e\x17\xaf\x13\xb8\xfe\x08\xf3\xeb\xab\xcb\xc5\xed\xe2\xfa\x2a\x81\xeb\x5f\x61\x76\xf5\x09\xde\x2d\xae\x2e\x4f\x01\x99\xd9\xa0\x02\x7c\x2c\x94\xf5\x40\x2a\x60\x96\x4d\x4c\x1d\x75\x09\x62\xc7\x84\xb5\xf4\x26\xe9\x02\x29\x5b\x33\x0a\x9c\x88\xac\x24\x19\x42\x26\xb7\xa8\x84\xcd\x84\x02\x55\xce\xb4\x8d\xaa\x06\x22\x52\x0b\xc3\x59\xce\x0c\x31\x6e\xe9\xc0\xaf\x71\x64\x45\x3e\x30\xba\x21\xc8\xe1\x0e\x05\x7e\x66\x04\x7e\xca\xb7\xfe\xea\xe7\x2c\x27\x8c\x8f\xa9\xcc\xff\x63\xe5\x66\x9c\x3d\x10\x78\x4f\x94\x46\x01\x3f\x11\x7b\x37\xe6\xee\xae\x2d\x18\xe9\x4a\x18\xf2\x08\x17\x30\x2a\x94\x34\xf2\xe5\x68\x1a\x45\x05\xa1\x0f\xd6\x54\xca\x4b\x6d\x50\x2d\x73\x22\x48\x86\x6a\x49\x0a\x36\x8d\x22\x96\x17\x52\x19\x18\x65\x52\x66\x1c\x63\x52\xb0\x98\x08\x21\x83\xd9\x63\x07\x33\x9a\x36\x62\xee\x9e\x9e\x65\x28\xce\xf4\x8e\x64\x19\xaa\x58\x16\x4e\x74\x50\x2d\x8a\xfc\x2e\x9c\x64\xaa\xa0\xe3\x8c\x18\xdc\x91\xca\x6f\xd3\x65\x86\x62\x19\x50\xc6\x01\x65\x2c\x0b\x14\xa4\x60\xdb\xf3\x7a\xe7\x19\x5c\xc0\x97\x08\x80\x89\xb5\x9c\xb8\x2b\x00\xc3\x0c\xc7\x09\x8c\xe6\xde\x25\xf8\xe0\x5d\x82\xd9\xcd\x62\x34\x75\x12\x5b\x5f\x60\x13\x18\x6d\x9f\x8f\xcf\xc7\xcf\xc3\x32\x95\xc2\x10\x6a\x6a\x1c\xfb\x13\x24\xb7\x50\x75\xad\xcf\xaf\xe6\xb7\x41\xd8\xfe\x4a\xc5\x27\x30\xb2\x85\xa1\x27\x71\x9c\x31\xb3\x29\x57\x96\xeb\x58\x7b\xf9\x33\x2a\xa8\x89\x03\xb5\x67\x81\xda\x33\x52\xb0\x16\x06\xda\x00\x4d\x60\x44\xd2\x9c\x89\x9f\xdb\x8a\x63\x26\x83\xdc\x93\xfd\xe3\xfe\xc3\x47\x83\x4a\x10\xbe\x4c\x25\xd5\xb5\xa1\xdf\x6b\x46\x8a\x9a\x2a\xe6\x28\x9e\xc0\xe8\x83\x54\x08\x64\x25\x4b\x03\xc7\x18\x7c\x8a\x00\x34\xdd\x60\x8e\x7a\x02\x6f\x6f\x6f\x6f\x92\x69\x7f\xc5\x2e\x50\x29\x74\xe9\x56\x46\xa1\xf0\xed\x23\xe2\xff\x6a\x29\x1c\x4c\xa1\x64\x5a\xd2\x63\xfb\x4f\xd3\x28\xd2\xa8\xb6\x8c\x62\x63\x88\xf7\xd7\xd6\x33\xe3\xdc\xea\x6f\x99\xeb\x94\xa4\xce\x5f\xb7\xaf\x0a\x0a\x73\x85\xc4\x60\xad\x77\xd2\xb9\xfd\xa0\xb3\x67\xa0\xd0\x94\x4a\xe8\xde\xd6\x47\x2c\x78\xf5\xac\x95\x00\x4d\x86\xba\x0a\x18\x93\x82\x8d\x2d\xd1\x75\xde\xed\x7f\x45\x69\x60\x02\x23\x57\x23\xdb\x17\x35\xdb\xa3\x8e\xcc\x4a\xa6\x95\x15\xfa\xe7\x7e\xf9\x29\x44\xb8\xe3\x98\x42\xa3\x18\x6e\x7d\x9b\xd1\x86\x98\x52\xdb\xd6\xdc\x78\x69\x5b\x08\x30\xa3\xe1\xa1\x5c\x21\x95\x62\xcd\x32\xd7\x85\xa8\x14\x02\xa9\x61\x5b\x66\xaa\x86\x89\x37\x68\x1a\x1a\xf6\xd7\x5d\x0e\xf6\xeb\x7f\x9e\x80\x0c\xbf\x4e\xc0\xa0\xa7\x29\x72\x34\x38\x10\xbf\x4b\xb7\xd1\x18\xde\xb9\xed\xda\xde\xd9\xfa\xf3\xe6\x07\x4b\xfe\x6f\x0f\x9a\x58\x11\xe0\x4c\x1b\x1b\xa7\xa0\xa8\x07\x42\xf0\xde\x8a\x9c\x74\xef\x8f\x85\xc2\xee\xfd\xd5\xe1\x88\xad\x8d\xdf\xf6\xa8\x54\xa2\x6e\x92\xae\xb5\xaa\xdc\x95\x66\x68\x0b\xa4\x60\x60\x2b\xb3\x15\xae\x37\x68\xc2\xd4\xb2\x68\x89\x9f\xec\x97\x0f\x9c\x0c\xeb\x7f\x99\x83\xc1\xdc\x01\xdf\x9e\xa2\x28\x47\xad\xed\x29\xd7\x6f\x03\xfb\x86\x72\x45\x72\xac\xc7\x9f\xba\xca\x8c\x84\x15\xee\xbb\x0c\xa6\x4e\xd8\x0e\x1b\x22\x73\x27\x03\x5c\xc0\x8b\x69\x8d\x70\xbb\x09\xb2\xf6\x28\xaf\x67\x01\xc7\x83\x93\xe8\x3c\xfa\x26\xc8\x25\x05\xd2\xbd\xd2\x05\x9c\x4f\x8f\x5a\xeb\x88\x6a\x35\xc0\x0d\xba\x19\x45\x2a\x37\x06\xb6\xcd\xde\x11\xdd\x36\xda\xce\x5d\x6e\x42\xb4\x83\x18\x6a\x13\xf9\x4e\x24\x39\xc8\x87\x03\x07\x52\x34\x84\x71\xdd\x67\x22\xa8\x82\x42\x5d\x48\xa1\xd1\x7b\xe4\x37\x17\x06\xf3\x46\xb0\xef\x42\xa7\xe1\xfc\x11\xb6\xb9\x94\x0f\x76\xd0\x2b\x86\xb9\x1e\x84\xee\x51\xb3\xd0\x1d\x5c\x26\x7c\x1b\xad\xb4\xc1\xfc\xd0\xf9\xb6\x2b\x97\xce\xfb\xaf\x3a\xd4\x6f\x44\xed\x88\x10\x63\xc7\xd1\xd6\xb3\xff\xa1\xbd\xe9\x46\xda\x33\xd6\x28\x59\x7d\xd3\xab\xc3\x6e\xb6\x7f\xc2\x5c\x96\x3c\xed\xf8\xb6\xc2\x1a\x38\x24\xe7\x50\x5c\x93\xe6\x00\xb1\xaa\xed\x2c\x08\x86\x84\x13\xe6\x78\xec\x42\x97\x82\x2f\xc7\xb7\xbf\x2b\x06\x41\xe9\xfd\x60\xff\xc4\xc2\x56\x41\x3a\x94\x6e\x87\x36\xb7\x85\xf6\xc6\x5c\xf6\x72\xad\xed\x3c\x4b\x3b\x36\x0c\x64\xe6\x40\xcc\xce\xa7\x43\x51\xd7\x1d\xa2\x07\xb4\x1b\xa2\x5f\x0e\x19\xdd\xca\xbe\xbf\xb7\xe9\x03\xfa\xad\x41\xc4\xc8\x7a\x0e\xb1\x97\x47\xe0\x5a\xf2\x17\xf0\xc3\xf1\xae\xd7\x69\x94\x83\xa5\xd6\x74\xcf\x33\xa0\xa5\x52\x28\x0c\x0f\xfd\x8e\x69\x20\x3b\xf7\xf6\x96\x13\xa2\xbf\xd9\xbb\xeb\xf3\x4e\xae\xe1\x5d\xb9\x42\x25\xd0\x60\x47\xeb\xe1\x47\xbd\xac\x85\x1c\x8f\x6e\x53\x0a\x94\xeb\xc6\x8a\x65\xfb\xb4\xdc\x9f\x57\xe1\x11\xb3\xfb\x64\xe0\x64\x38\x38\x1d\x66\xf7\x89\xf3\xd7\x5a\xdf\x10\xfe\x14\xfd\x81\xd6\xcf\x34\xbc\x9d\xed\xeb\x6b\xc3\xb2\x0d\xaf\x96\x64\x4b\x18\x77\x2f\xdd\x8e\xeb\x96\x41\x6b\xb2\x52\x8c\x86\xd6\x5b\xea\xde\x09\x87\x66\x27\xd5\xc3\x32\x08\x5d\xc0\xab\x69\x74\x34\x50\xb5\xcd\x5f\xa2\x9e\xbf\x97\xc4\x10\x98\xa3\xa8\x13\x60\x76\x9f\xd8\x25\xbf\x02\x29\x31\x64\x49\xfd\x75\x3b\x20\x73\x85\x29\x0a\xc3\x08\xd7\xce\xba\x92\x75\x3b\x5f\x0d\xd5\x96\xa3\xad\xeb\x76\x9e\xff\xf2\xe9\x1a\x98\xc1\x5c\xd7\x4a\x37\x2a\xe4\x5e\xa9\x30\xb5\xf5\x66\xcf\x46\x2d\x4b\x45\xb1\x9b\xe1\x0b\xa1\x8d\xfb\x84\x92\x29\x59\x16\xbd\x7e\x34\xbb\x4f\xea\xfd\x37\x76\x1b\x58\xb8\x5b\x7a\x69\x4f\xf6\x3e\x68\x8c\x6e\x0e\xc8\xa8\xa9\xec\x92\xd2\xc9\x1b\xaf\xa8\x30\x73\x53\x51\xa9\xcf\x90\x68\x73\xf6\xe2\x14\xd0\xd0\xf1\xb3\x46\x32\xc4\x2c\xc8\x35\x54\x76\x40\x42\x1e\x30\xce\x4c\x05\x9f\xa5\x40\xdd\x02\x5c\x9d\x42\x7d\x7d\x4e\xdd\xf5\x0e\xed\x75\xda\x7f\x52\x43\x40\x78\x64\x1b\x75\xe9\x51\x1b\xf6\xf7\x49\x6b\xd3\x81\x76\x83\x5a\x86\x6f\x27\xd4\xe6\x91\x45\xea\xc7\xb7\xc5\x4e\x3b\xce\x07\x65\x95\x20\x55\x68\xde\x61\xb5\x48\x1d\xe0\xec\x66\x01\x33\x4a\x51\xeb\x3e\x3d\xda\x49\x2e\x1f\xb0\x5a\xb6\x9b\xe7\x01\x96\xd7\x7e\x87\x55\x83\x47\xbe\x86\xe7\x37\x2d\xec\x90\xeb\xbf\x4a\x05\xbb\x0d\x0a\xd0\x32\x77\x1f\xeb\x44\xa6\x81\xd8\x97\x6c\xae\x90\xa4\x95\x27\x20\xd4\x5e\xcb\xe7\x81\x34\x3d\x70\xfd\xee\x66\x0e\x2c\x3d\x85\x15\x27\xe2\xc1\x19\x6b\xff\x8d\x3c\xa2\xed\x4d\xee\xbe\x92\xe5\xe8\x14\xd6\x8c\x73\x4c\x81\xad\xdd\x07\x44\x6b\x80\xad\x8c\xbb\x9b\x79\xdf\xab\x6d\x41\x87\xe8\x49\x90\x96\xca\xa6\x8e\xcb\xef\x01\x2a\xdc\xae\xcf\x7e\xaf\x7f\x40\x45\xaf\xa0\x20\xc5\x35\x13\xf6\xd5\xc9\x54\x05\xba\xb7\x5a\x51\xe6\x2b\xdb\xdb\xd6\x4d\x39\xe9\x3e\x2f\xdd\xaa\xeb\x50\xd2\xe0\x3b\xbc\x93\xfc\xd5\x98\x13\x95\xe1\x91\x62\x71\x42\x7d\x2f\x3f\x30\xc1\xf2\x32\x1f\x32\x04\x4e\x52\x5c\x93\x92\x1b\x97\xbe\x9f\x51\xc9\x3d\x24\x13\xe6\xe5\x39\xe4\x4c\x2c\x7f\x2f\x89\x30\x96\xa6\xc6\xff\x1a\x99\x3c\x7e\x07\x32\x79\x6c\x23\xbf\x6c\xbd\xdc\xc4\xb1\x9d\xc5\xda\xe7\x97\x4d\xd8\xc4\xbf\xa0\xb5\xa6\xb5\xfd\x9b\x98\x1f\xe4\xe2\x18\xfc\xd4\x66\x53\xa4\xd6\xae\xc7\xc3\x43\xbd\xfe\x84\xb7\x06\x59\xa0\xf2\x27\x9d\x7d\xe5\xb8\x7e\x77\x64\xb8\xae\xa1\x06\x5e\x10\x0f\xf2\xd9\x90\x0c\xa4\x1f\x16\x33\x66\xdf\x37\x0a\xa9\x99\x91\xaa\xea\xc7\x2e\x63\xa6\x75\x18\x1f\xd6\xf1\x86\xe8\x4d\x3d\xce\x58\x24\x2a\xf3\x9c\x99\x21\x14\xbf\x73\x10\xad\x81\x53\xd6\x28\x44\xe7\x2a\xe5\x48\x84\xaf\x69\x7b\x32\x0d\xc2\x5a\xe1\xa5\x9d\x9a\x70\x1f\xae\x00\x7d\xe9\x6a\x73\xed\x4f\xb5\xbe\xae\x5b\x5c\xa6\x5e\xef\x87\x8e\xde\xdd\x3e\xc2\x99\x6b\x9f\xa9\x1f\xb6\xf2\x82\x71\x3c\xb0\x41\xb6\xf8\x79\xd5\xc1\x99\x7b\x0d\xb5\x3f\xf1\x5b\x7a\xb4\xde\xbc\x80\x7f\x75\xb4\x6e\x38\x31\x36\x72\xc0\x8c\x27\xc1\x0b\xfa\x96\x1b\x83\x2a\x85\xfb\xd2\xdd\x1a\x6a\x02\x62\x51\x2b\x5e\xc0\xbf\xfb\x0d\xa1\x76\xa9\x95\x14\x6e\x6b\x20\x57\x82\x37\x9d\x09\xab\x1e\xff\xa3\xff\x05\x00\x00\xff\xff\xa7\x89\x23\xc6\x9f\x19\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 6559, mode: os.FileMode(420), modTime: time.Unix(1534874472, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
