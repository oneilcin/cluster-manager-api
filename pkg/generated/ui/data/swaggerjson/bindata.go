// Code generated by go-bindata.
// sources:
// assets/generated/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x6d\x6f\xe3\x36\x12\xfe\x9e\x5f\x31\xd0\x1d\x70\x59\x20\xb1\xb7\x39\x1c\x50\xe4\xd3\xa5\x09\xae\x35\xd2\xb4\x41\xbd\x68\x50\x5c\x17\xc6\x98\x1a\xcb\xac\x25\x52\x4b\x52\xf6\x7a\x0f\xf9\xef\x07\x52\x92\x2d\xc9\x92\xa3\x48\x49\xd6\x2d\x5c\xa0\x88\xd7\x7c\x99\xf7\xe1\x43\x0e\xe9\xff\x9d\x00\x78\x7a\x85\x41\x40\xca\xbb\x04\xef\x62\xf0\xde\x3b\xb3\xdf\x71\x31\x93\xde\x25\xd8\x76\x00\xcf\x70\x13\x92\x6d\xbf\x0e\x13\x6d\x48\xc1\x1d\x0a\x0c\x48\xc1\xd5\xfd\xc8\xf5\x07\xf0\x96\xa4\x34\x97\xc2\xf6\x5a\xbe\x1f\xe4\x13\x01\x78\x4c\x0a\x83\xcc\x6c\x66\x03\xf0\x04\x46\x6e\xba\x31\x46\x3a\x11\x01\x5c\xff\x74\xfd\x21\xeb\x0e\xe0\x25\x2a\xb4\x8d\x73\x63\x62\x7d\x39\x1c\x06\xdc\xcc\x93\xe9\x80\xc9\x68\xa8\xd3\xfe\xe7\x4c\x30\x33\x64\x29\x2f\xe7\x51\xca\xcb\x39\xc6\x7c\x3b\x07\x45\xc8\xdd\x2c\xe8\x47\x5c\xfc\xbb\x38\x70\xc0\xa5\xe7\xba\x3d\x9e\x00\x3c\x3a\x69\x35\x9b\x53\x44\xda\xbb\x84\xff\xa6\x3c\x3b\xda\xb9\x00\xf6\x1f\x76\xc4\x47\xd7\x97\x49\xa1\x93\x52\x67\x8c\xe3\x90\x33\x34\x5c\x8a\xe1\x1f\x5a\x8a\x6d\xdf\x58\x49\x3f\x61\x2d\xfb\xa2\x99\xeb\xad\xca\x87\x18\xf3\xe1\xf2\x9b\x5c\xca\xa2\xf6\x02\x2a\x2a\xd3\xb2\x9f\x44\x11\xaa\xb5\x15\xf7\x81\x87\x21\x28\x32\x8a\xd3\x92\xc0\xcc\x09\xb4\x41\x93\x68\x90\x33\x40\xc8\x26\x03\x14\x3e\x70\xa3\x61\x91\x4c\x89\x49\x31\xe3\x01\xcc\xa4\x02\x26\x85\x20\x66\xf8\x92\x9b\xf5\x46\x95\x00\x9e\x8c\x49\x39\x96\x47\xbe\xa5\xf1\x3d\x99\xcc\x0f\x8a\x9d\x14\xe9\x58\x0a\x4d\xba\xc4\x1b\x80\x77\xf1\xfe\x7d\xe5\x2b\x00\xcf\x27\xcd\x14\x8f\x4d\xe6\x31\x85\x89\x52\x89\xac\x41\x70\x67\x18\x80\xf7\x77\x45\x33\x3b\xe2\x6f\x43\x9f\x66\x5c\x70\x3b\x83\xce\xb5\x34\xc9\x7c\x61\x82\x31\xdf\x72\xf9\x0b\xc5\xe1\xda\x2b\x4d\xf4\x78\x52\xf7\xf9\xb1\x20\x4e\x8c\x0a\x23\x32\xa4\xb6\xc6\x4b\xff\xab\x08\x92\xbb\xb2\xfb\x7b\xb6\x57\xc8\x9f\x30\x22\x6b\x07\x6b\x95\xdc\x12\x46\xc2\x94\x20\x94\x72\x41\x3e\x24\xf1\xa0\x3a\x05\x77\x23\x3f\x25\xa4\xd6\xd5\x26\x45\x9f\x12\xae\xc8\x9a\x64\x86\xa1\xa6\x4a\xb3\x59\xc7\x8e\x31\x6d\x14\x17\x41\x51\xfc\xc7\xb3\xa7\xc5\xc1\x95\x1e\x68\x62\x8a\xcc\x64\x41\xeb\x09\xf7\x9f\x90\xed\xc3\x9c\x60\xec\xfa\xdf\xd2\x7a\xe4\x3b\x77\xba\xba\x1f\xc1\x15\x63\xa4\xf5\x21\x8a\x85\x8e\x33\x2b\x5d\x6b\xd1\x52\x61\x6e\x69\xbd\x11\x0f\x0f\x4f\x3c\x45\x81\x65\xfc\x69\x99\x7e\x71\x1d\x0f\x56\x94\x2f\x89\xa2\x01\xc6\x71\x3b\xdf\xbb\x8a\xe3\x03\xf6\x3a\x27\x8b\x21\x81\xc2\xb4\x90\xe5\x83\xeb\x78\xd8\x86\x89\x51\xeb\x95\x54\x6d\x4c\x73\x9f\x75\x3d\x6c\x81\x74\x32\xdd\x70\xde\x32\xdd\x15\x46\x1c\xaa\x6c\xb1\x92\x4b\xee\x97\x16\xea\x3a\x71\x8a\x2b\x53\x3e\x44\xc3\x29\xae\xf4\x10\x17\x7a\xb8\x8c\x56\xa8\x68\x48\x86\xbd\x7b\x1b\xc9\x36\x9f\x3f\x16\x56\x64\x83\x41\x75\x2d\xce\xf1\xe8\x76\xf0\xc7\x93\x8a\x72\x3c\x9f\x42\x32\xb4\x1f\x30\xa5\x7d\xb6\x00\x69\x0f\xf8\xb9\x71\x5d\xff\x04\xf8\xa7\xc4\xe8\xa1\x40\xa0\x87\x39\x1a\xe0\xba\x08\x81\xfe\xa1\xc1\x0e\xb4\x48\xc8\x27\x6d\x94\x5c\x1f\x4c\xf4\x1c\x41\xd0\x11\x04\x7d\x5d\x51\x8e\x20\xe8\x90\x0d\x73\x04\x41\x47\x10\xf4\x7a\x92\xbd\x24\x08\x8a\x93\x27\x8e\x8c\x9c\xc0\xda\x1a\xb1\x0d\x08\xba\x56\x84\x7f\x0a\x10\x54\x62\xf4\x4d\x40\xd0\x54\xfa\x3b\x3e\x90\xba\x47\x5d\x4b\xc1\x3b\x8c\x4a\xaa\xce\xf1\xd2\x0a\xb8\xd3\x41\x1b\xf1\xbb\xfb\xdb\x49\x41\x7b\xd5\x13\xcc\x61\xc8\xb5\xe9\x76\x8c\x89\x60\xc7\xda\xd8\xcc\xe6\xd2\xad\x4e\x27\x7f\xb4\x04\x0f\xd8\x39\xcb\x9c\xbe\x89\x77\x1e\x11\xed\x11\xd1\x7e\x5d\x51\x8e\x88\xf6\x90\x0d\x73\x44\xb4\x47\x44\xfb\x7a\x92\xbd\x16\xc2\xd8\x96\x9e\x9f\x05\x2e\x12\x25\x20\x1b\x0a\x5c\xcc\xa4\x8a\x1c\x8c\x00\x9c\xca\xc4\x00\xc6\x1c\x34\xa9\xe5\x5e\x20\xfc\x3d\x99\x5f\xd3\x19\x46\xdb\x09\x0e\x1c\x73\x64\x0c\x77\xc2\x1b\x5d\x8c\xb5\xa9\xb2\x17\x58\xdb\xd6\xb9\x4b\x00\xf5\xea\x76\x3c\x8e\x89\x5d\xdd\x8e\x47\x42\x1b\x14\x8c\xbe\x57\x32\x89\x8b\x86\xcd\xdd\x4b\x4e\xff\x20\xb6\x4d\xbb\x36\x24\x62\x52\x86\x57\x34\x9d\x87\x4c\x49\xf7\x15\x17\x3d\x2b\xb5\xe5\x57\x1d\x6c\x2e\x10\x85\x00\x0a\x1c\x27\xf5\x5a\x49\xe7\x7b\x3e\x8d\x5c\x4a\xb0\xbd\xe1\x74\x6c\x50\xf8\xa8\xfc\xc9\xcd\xc5\x64\x79\x71\x06\x64\xd8\xe0\x5d\x3d\xc9\x88\x8b\xc9\xa7\x04\x85\xe1\x66\xdd\x44\x9a\x0b\x43\x41\x25\x49\x78\xa9\x97\x66\xcd\xff\xbc\x68\x60\xec\x8e\x0b\x1e\x25\x11\x88\x24\x9a\x92\xb2\x2a\xe0\x19\xab\x1a\x4e\x7d\x9a\x61\x12\x1a\x0d\x46\xc2\x17\x52\xb2\x89\x45\xfc\xfc\xaa\x2c\xe2\xe7\x6e\x2c\x9e\x54\x58\xad\x31\x87\xb3\xb5\x06\xe7\xb0\x76\xfb\xe1\xcc\x83\xc2\xaf\x23\xe6\x95\x52\x52\xd9\x9d\x1f\x52\x77\x7e\x18\xdf\xa0\xc1\x6b\x12\x95\x8b\x1c\xcf\xf5\xe5\x0c\x03\x76\xf1\xb4\x87\x39\x67\x73\x48\x27\x80\xd3\x44\x9f\x13\x6a\x73\xfe\xcd\x5e\x1f\xc3\x25\xf2\x10\xa7\x3c\xe4\x66\x3d\xf9\x22\xc5\x6e\x12\xcb\x49\xa3\x52\x58\x5e\x27\x3c\x6e\x28\xaa\xf6\x6f\xbf\xee\x55\x19\x2f\xb2\x02\x8e\x95\x82\x10\xd3\x33\xc8\x3f\x5f\x30\xf7\x79\x45\xf6\xb3\xbf\x2b\x5d\xb3\xed\x53\x3a\xd6\x4e\x90\x19\xaa\x8d\x5d\x5f\x2c\x4d\xbd\x54\x0a\x89\xfe\x35\x08\x51\x05\x74\x4c\x1e\x7f\x9d\xe4\x71\xaf\xb2\x4b\x5a\x89\x22\x7f\x54\x89\xab\x67\x7b\xda\x32\x66\x16\x08\x77\x5d\x12\x7f\xbd\xbf\x06\xee\x9f\xc1\x34\x44\xb1\x70\xf8\xd8\xfe\xff\xbb\xc7\x1c\xef\x20\x05\xb9\x2f\xd6\x32\xf9\xdd\x3b\x83\x19\x0f\x43\xf2\x81\xcf\xec\x17\x80\x8a\xe0\xbb\xdf\x7e\xb6\x73\xd4\x5b\x5d\x13\x4b\x94\xcd\x36\x4e\x87\x5d\xd9\x1c\x67\xb3\xec\x5b\xb2\x39\x46\x13\x25\x43\x9a\xa0\xea\x96\x50\xad\x2e\x46\x57\x77\x60\x27\x71\x12\x17\x6f\x75\x9d\xa2\x12\xef\x72\x3b\x6a\x2d\x19\x77\xc0\xd0\xf7\x5b\x79\xd2\x7f\xa4\x82\xd5\x9c\x04\x68\x19\x11\x98\x39\x17\x81\x76\xba\xc3\x50\x11\xfa\x6b\x48\x75\xed\xef\x71\x9e\x5f\xef\x1e\x50\x91\xf5\x9f\xf4\xd3\x1d\xb2\x39\x17\xee\x8b\x3e\xbe\x93\x68\x52\xbd\x00\x55\x3e\x81\xd3\xd8\x78\xfc\x43\xb6\xb1\xaa\xb7\xd1\x5c\x6a\xd3\x99\x92\x1d\xdc\x8a\x4a\x2c\x55\x23\x95\x1e\x19\xc6\xf2\x60\xa7\x6e\xc7\x43\xbe\xed\x7e\x86\xb4\x35\x5b\xd7\xc5\xb7\x7a\xb3\xa1\xd9\xb8\xa4\x14\x46\xc9\x10\xe2\x10\x05\x0d\xe0\xc3\x9c\x6b\x10\xd2\x27\xe0\x1a\xa4\x08\xd7\x80\x10\xa1\x73\x5a\x6e\x71\x2e\xd7\x30\xe3\x14\xfa\xb6\x39\xcd\x60\xfe\xa0\x95\xd3\x5a\xfa\x4c\x91\x4f\xc2\x70\x0c\x5d\xee\x4c\x74\x6a\x67\xe7\xae\x5c\x04\xc5\x10\x29\xfb\x6e\x65\x57\x52\xb3\xa5\xea\xe1\xb4\x01\x37\x93\xdd\x1d\xe2\xf3\xbc\xc9\x60\x00\x52\xa4\xfb\x00\x6e\x40\x51\x2c\x35\x37\x52\xad\xeb\xed\x69\x49\x32\x19\x45\xbc\x87\xff\xa2\x9e\x6f\xb6\x1e\xdc\x40\x36\x5d\x23\x39\xa3\x88\x26\xda\xa0\xe9\x16\x9c\x0f\x73\x32\x73\xbb\x38\x29\x10\xd2\x38\xaa\x76\x46\x58\xa1\x06\x16\x12\x8a\x34\x23\x4d\x13\x1e\x36\x30\x61\x9b\xfc\x89\xdf\x95\x81\x1b\xb7\x7a\xcc\x1c\x05\xbf\x41\x4c\xd9\xcb\x8e\x99\x57\x59\x22\x81\x73\x4e\xdf\x3a\x29\x93\x51\xcc\x43\xaa\xa7\x98\x35\xaa\x4e\xf4\xae\xb3\xc1\x8e\x54\x43\xdc\x87\x68\xac\x8f\x77\x9a\xff\x3e\x1b\x0c\xdc\xa4\x66\x4a\xe9\xa5\x07\x72\x43\x50\x89\x10\x36\xe8\xd2\x3b\xe6\x19\xed\x93\xe2\xdf\x3c\xfa\x6a\x8e\x09\xae\x1e\xc6\xd7\xdb\x60\xee\x13\x7d\xe5\x6a\x43\xd7\x68\x68\xae\x41\x34\x02\x89\x4a\x31\xa0\x1f\xe5\xa6\x12\x41\x3d\xf5\x1e\xfb\xb4\xc6\x13\xfc\xd7\x4f\xc2\x75\x6e\xf0\x25\x51\x39\xa4\x18\x93\x5a\x72\x46\x57\x8c\xc9\x44\x98\x3e\x2e\xc1\x42\x4e\xc2\xf4\x71\x87\x6b\x37\xc3\xc8\x87\x53\x5c\xe0\xa5\xab\x12\xdc\x34\x6c\x24\x32\x62\xa9\x43\xf4\x24\x98\xfa\x42\x46\x34\x5f\xb3\xdb\xed\x0e\xac\xd6\x31\x55\x5d\x9a\x7a\xa6\xeb\xea\x3b\x80\x0c\x41\xbb\x73\x6b\x50\xa4\x65\xa2\xdc\x76\x84\x3e\x5f\x42\x28\xd1\x87\x29\x86\x76\xd3\xa0\xde\xb5\xb4\xdb\xcb\x04\x70\x56\xab\xe9\xaa\xb9\xba\x0a\x4e\xc3\x29\x5a\x5a\x48\xe9\x4a\xa8\xb6\xbc\xf2\x72\x70\xab\x44\xab\xa1\xf6\xd1\x90\x8b\x2a\x25\x88\xce\x99\xa8\xb9\x30\xf1\x55\x92\x43\x96\x17\x6e\xc8\x20\x0f\xed\xb6\xb4\x8f\x97\x75\xd4\xcb\xe8\xa6\xf2\xa0\xa6\xde\x04\x9d\xf7\x2c\x35\x4f\x76\x1a\x8c\xec\x5e\x59\x75\x84\x5e\xdb\x3b\xb1\xdb\xc7\x5a\x4f\x52\xdc\xbe\xdd\xea\x4d\xb5\xf0\x0c\xcc\x21\x22\xf7\x0a\xcc\x7e\xac\x67\xa2\x2d\x88\xc8\x1c\xe4\xe8\x1a\x6f\xe2\x1a\xad\xad\x52\x53\x70\xe9\x63\x9e\x50\xb2\xea\x16\xad\xbd\x80\x6e\x85\x70\x0b\xde\xce\xc9\x6b\x55\x8d\xac\x76\x39\x73\x4d\xad\xab\x5f\x3b\xeb\x62\x03\x72\x2f\x67\x49\xb7\x25\x79\xda\xce\x59\xe3\x2e\x40\xea\xc2\x63\x2d\xe6\x6a\xe0\x36\x83\x10\x3a\xed\x5c\x86\x19\x36\x88\x31\x5c\xd8\xbf\x29\xb0\x70\x70\x62\x93\xef\xeb\x70\x45\x55\xae\xfc\xb0\x32\x3d\x94\x7b\x81\x1a\x40\xbd\x2a\x5a\x15\x02\x0b\x13\x35\xd4\x0b\x2a\xc7\xb0\xbd\x03\xe4\xa1\x77\x80\xf8\x68\x70\xc2\xaa\xc5\x9f\xd6\x9a\xa8\xab\x21\xed\x09\xa6\x87\xf1\x1b\x85\x52\x79\x87\xf8\x92\x81\xb4\xc1\xbe\xbd\xd4\x55\x73\x6a\x5e\xcf\xe4\x77\xbf\xfd\x0c\xa9\xa7\x1e\x8e\xfb\x37\x14\x98\xde\xdc\xfd\xef\x74\xf0\x55\x8a\xef\xcd\x4f\xa5\x37\x17\x95\x1b\xcf\x54\xf2\xbb\x31\x5d\x3d\xbb\xa4\x80\xfb\x6c\x36\x97\x04\xf6\x9c\xf4\x66\xdd\x40\xc7\xc4\xf8\x2c\x7b\xe1\xdf\x57\xfb\x25\xe2\x5f\xc3\x0c\x45\xfc\xb1\x11\xf1\x1c\x58\xa2\x14\x09\x13\xae\xd3\x03\x63\xae\x01\x57\x1a\xa4\x82\x08\xb1\x21\x8a\x16\xdf\xea\xde\x87\xb0\xcb\xed\x01\xde\x6d\x32\x25\x25\xc8\x50\x03\x39\x5c\x75\xcf\x6c\xb5\xf9\x7f\x7f\xc6\x6d\xb2\x7a\x91\x23\xbb\xfa\xbe\x10\x4f\x19\x68\xab\xa5\x93\x5e\xea\x7a\x19\x42\xdb\x2a\x52\x43\x81\x86\x07\xf3\x49\xb1\x34\xdf\x64\xda\xa9\x94\x21\xa1\x68\xaa\xa3\xd4\x36\xef\x3b\x9a\xce\x53\x02\xd7\xf0\xc3\x55\x03\xae\x27\xb3\x92\x6a\x31\x99\xe1\x54\x71\xd6\xd9\xe7\xd2\xe1\x59\xee\xa9\x1c\xe4\x76\x8a\xe9\xf4\xda\x55\x8f\x60\x96\x8b\xb7\xd6\xf2\x0a\x75\x31\xed\xa6\x47\x57\x5c\x83\xa2\x4f\x09\xe9\x86\x72\xc0\xee\xef\x96\x3c\xd3\x11\x0b\x9b\xc7\x66\xf3\xf8\xee\xf0\xa1\xba\x2b\xca\x39\x83\xfc\xf6\x5d\x5f\xa3\x15\x02\xa1\x87\xe5\xa2\xb4\x08\xfb\x36\x10\x62\x5f\x05\xf8\x69\x18\x91\xf5\xd6\xb0\x72\xd7\x62\x98\x8c\x62\xc5\x35\xf5\xdc\x7a\xd6\x3c\x49\x3e\xbc\x40\xb8\x96\x49\x19\xa4\xda\xc8\xcf\x5e\x26\x37\x61\x8e\x1e\x5b\xfc\x71\x69\x5b\xbf\x1b\x52\x6d\x75\x5b\xf7\x98\xe4\xf0\x94\x3b\xd2\xe5\xfc\x9d\x96\x54\xf5\x5a\xdb\x30\xdf\x97\x47\x5e\x2d\x68\x9e\x48\x3d\x4f\x87\xca\x8f\xd5\xa7\x49\x7d\x6c\xf7\x97\xb3\x5b\xdf\xfc\x5f\x38\x5d\xee\xa4\xd7\xd2\x65\xe7\x03\xd4\xeb\x0c\x36\x37\xca\xdd\x3a\xfb\xf3\x6d\x03\xac\x4b\xe5\x98\xf0\xda\x3b\x11\x7b\x54\xfb\xf4\xd5\x8a\xfd\xd5\xf2\x62\xcf\x5d\x0b\xec\x56\x19\x1c\x15\x57\x54\xc8\xe1\x7a\x29\xa7\x6d\x6f\x83\xd3\x67\x43\x4a\x60\x78\x23\x59\xe1\x3a\x78\xe5\x56\xcb\x9d\x54\x94\xdd\xc7\xdf\xf3\xd3\x73\x3d\x7e\x30\xce\xb2\x73\xf2\x78\xf2\xff\x00\x00\x00\xff\xff\x7a\x30\x0d\xa9\x08\x4f\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 20232, mode: os.FileMode(420), modTime: time.Unix(1537147136, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": {apiSwaggerJson, map[string]*bintree{}},
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
