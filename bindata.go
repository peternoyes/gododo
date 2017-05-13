// Code generated by go-bindata.
// sources:
// data/firmware
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

var _dataFirmware = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x6f\x6c\x1b\x47\x76\xc0\xdf\xfe\x11\x39\x4b\xd2\x22\x2d\xd3\x27\x5a\x61\xec\x0d\x4f\x56\x2c\xf5\xe0\xca\xd7\x34\x60\x1d\x85\x5e\xea\x68\x87\xd4\x6d\x12\x29\x47\x19\x49\x7a\x4d\xd6\x89\xd5\xd2\x39\x27\x51\x12\xa7\x91\x75\x8c\x96\xbc\x78\x75\x8a\x22\x05\xa6\xa4\x43\x9c\xf6\x0b\xc3\x6a\x05\xaf\x0b\x5f\xc3\x0f\x07\xa4\x28\x1c\xe8\x1a\x38\x50\x8c\x2b\xb0\xb8\x36\x4d\xae\x40\x00\xb7\x8e\xdd\x24\x77\x3d\xec\x01\x39\x5b\x6e\x64\xb3\x78\xb3\xfa\x67\x5b\x6e\x72\x45\xdb\x4f\xf3\x93\xf0\x76\xe6\xcd\x9b\xf7\xde\xcc\xec\xcc\xee\x52\x49\x8f\xa6\x8f\xa5\x23\x1f\xbe\x25\x6b\x13\xb9\xe3\xb9\x13\x39\x65\xe0\xc3\xb7\x2a\xb5\xbf\x90\x03\x25\xd9\x3c\x22\xaf\x9b\x92\x3d\x25\xf9\xfd\x92\xfc\x45\xe9\x2c\xe7\xdc\x2b\xdf\x35\x99\x6e\x0d\x39\x24\xd7\x2a\x3b\x1e\xf5\xf6\x82\xfa\xab\x82\xfc\xfa\x98\xa5\x1b\x9c\x75\xc4\xe0\x2d\x30\xa2\xc6\xad\x72\xcb\x94\xf5\xea\xa2\x82\x37\xa2\x72\xcb\x94\x9c\x1e\x93\xef\x9a\x6c\x0d\xd9\x97\xe9\xc5\xb9\x2c\xef\x33\xd4\xdb\x0b\xef\x7d\xc7\x09\xbe\xb7\xd7\xa9\x7f\xaf\xdb\x69\x7a\xef\x1e\xa7\x05\xfd\xd5\x17\x65\x28\x59\x7f\x59\x39\x22\x6f\x98\x90\x3d\x25\xf5\xf6\x02\xd6\xdf\x58\x5d\xf7\x94\x30\xe4\x77\xaf\x0f\xf9\xdc\xa2\x42\x58\x0e\xb9\x6e\x42\xfe\xd6\x84\x0c\x25\xf9\x97\x13\x58\x76\xfb\xe5\x56\xcc\xac\x3a\xda\x13\x4d\xff\x7a\x42\x85\x87\x54\x7f\x71\x4d\xcf\x03\xd7\x7b\xa6\x26\x45\x6a\x22\x58\x79\x43\x2c\x43\x45\xa8\x72\x25\x61\xce\xbe\x7c\xc6\xf1\x5e\xe0\x2f\x88\x6a\xba\xa8\x42\x5e\x7e\x2c\x7f\x70\xa1\x56\x99\x7d\x95\xd3\xb5\x0a\xa0\xb4\xfc\xe3\x0d\xba\xb5\x65\x7c\x83\xae\xa5\x4f\xae\xd7\x71\x2e\x72\xe3\x21\x5d\xc3\x32\x71\x2e\x9f\x0c\xe9\x9a\x05\x86\xcf\xea\x30\xfc\x56\x52\xbe\x35\x5f\x06\xb9\x2d\x5f\xf2\xcd\xcd\x2a\xf6\x45\xad\x0c\x55\x5f\x89\x60\xf9\xd2\x48\xd5\xd7\x43\xda\xff\x13\x8d\x89\x55\x30\x24\xd9\x93\x97\xff\x30\x2f\x07\xf2\xf2\xfd\xf9\x88\x49\x0e\x28\x06\x39\xfa\xcb\x0b\x92\xfd\xef\xd4\x8b\xa6\xbe\x9c\x07\x06\xe3\xff\x99\x3d\x07\xf7\x3d\x97\x3b\xf0\xd4\x9f\x6d\xdf\xbe\x1d\xd4\xa7\xf7\xed\x5f\x2c\xde\xb3\xef\xc9\x3e\xf8\xce\xe1\xe7\x0e\xf5\x3d\x09\x7b\xfb\x9e\x7d\xee\xc0\xd3\x4f\xed\x94\x77\x6c\x6f\xdf\xbe\x03\xba\x9f\x7d\x7a\xff\xf3\x8f\x1f\xda\x29\x7f\xfb\xc0\x21\x70\x55\xdf\x3b\x70\x08\xac\x1f\x8f\xcb\xba\x25\x1a\x9c\x05\x7d\xdc\x52\x31\xe4\x16\xc1\xe0\x2c\xc5\x2d\x6e\x47\x51\xa4\x4a\x14\x3f\x58\x2e\x71\x06\x67\x95\x5d\x9b\x19\x14\xc7\x51\xec\x42\x71\x12\xc5\xab\x28\xe6\x50\xfc\x23\x0a\x05\xc5\xfb\x28\x1c\x14\x1f\xa2\xa8\x43\xf1\x0b\x14\xbf\x41\xf1\x2f\x28\x1a\x50\x94\x96\x5d\x5d\x2b\xe4\x7d\x86\xf5\x37\xe3\x32\xdd\xcf\x9c\x35\x60\xf0\xd6\x47\x86\x60\x8d\xd1\xc3\x42\xac\x0a\x25\x6e\xd5\x61\xf1\xbc\x21\xa7\xc7\x34\xb9\x77\xd2\x90\x50\xf8\xac\xd3\xae\xa9\x29\xa5\x5b\xbd\xc7\xab\x82\x11\xc8\x75\x75\x75\x9d\x38\xf5\x93\x31\x83\x3b\x75\x6a\xcc\xe0\xa7\x7d\x55\x6e\x6b\xc0\xe1\x2d\x8e\xf6\x6a\xfc\x1d\xba\x9a\x8d\x8e\x64\x06\xa6\x7d\xeb\xb9\x12\xa7\x99\x81\x4c\x6d\xda\xb7\x03\x8b\xd8\x3b\xb2\xe2\x07\x85\x1f\x45\x40\xee\x9b\x34\x84\x61\xd1\xf4\x77\x75\x75\x19\x8d\xa6\xbf\xd5\xeb\x78\x23\xa6\x74\x80\x18\x92\xd9\xb8\xda\x71\x20\xd2\xc7\xd1\x15\xe2\x31\xc1\x11\x33\x62\xd7\xe3\xa1\x38\x62\x5f\xc6\x8b\xfa\xe7\xc3\x38\x6a\xac\x09\x25\xce\x94\xe2\x9f\x12\xa7\xc5\x90\x22\x56\x01\x7b\x1d\xf5\xb8\xfd\x22\xa6\xaf\x4f\x30\x04\x5a\x15\x0d\x71\xda\x37\xa2\x3e\x34\xac\xad\x95\xd6\xef\x9c\x8f\x05\x18\xfd\x8b\xa5\xd8\x8d\x37\xc4\xc6\x60\x9f\xba\xc1\x6e\x41\xb1\xe9\xbf\x9d\x8d\x56\x2f\x06\xbf\x6e\x6e\x57\x87\x34\x1b\x1d\x51\xd5\x47\x22\xd3\xbe\x91\x0a\x98\x9b\xec\x6e\xf3\x16\xbb\x7e\xd5\x14\xa8\xf7\x8f\xb4\x0a\xce\xb6\x56\xce\x69\xa8\x72\xeb\xa9\xfe\x92\x5b\x50\xef\x1f\xa9\x72\x3b\x96\x34\x3b\x96\x34\x3d\x4b\x9a\x1e\x57\x63\xde\x62\xdf\x7a\x2c\x3d\x7a\xfc\x93\xaa\x60\x34\xe5\x8e\x9b\x4d\xd8\xfc\xeb\xd1\xe3\x55\xa1\x0c\x4b\xee\xef\x6b\xe5\x9c\xad\xab\x8c\xaa\x5c\x1d\x35\xfb\x15\x9a\x19\x4d\x65\x70\x15\xea\xfd\x23\xd7\x18\x6d\xbd\xde\x68\xeb\x1a\x46\xbb\xaf\x37\x42\xc5\x0d\x0b\xeb\xbd\xe9\xc2\xee\xff\xa1\xda\xf2\x4a\xdc\x22\xe7\x1b\x8d\xa6\x69\xdf\x88\xdb\xa1\x8e\x9a\x7b\x70\xc6\x1c\x41\xfd\xab\x97\xcd\x5b\xec\x07\x66\x1a\x4f\xfd\xc3\x18\xde\xa4\x55\xc1\x77\xc6\x5e\xc0\x3b\x77\xa6\xe9\xd4\xfb\x63\x3b\xea\x4a\x75\x55\xa1\x0b\x55\x75\xa5\xba\x11\xfb\x83\xaf\x66\xa8\x5e\x19\x6d\x15\x9c\x43\xad\x9c\x73\x67\x55\x98\x69\x5c\xb2\xac\x0a\x33\x4d\x2b\xbe\xce\xdf\xac\x49\xbd\x32\x8a\x4d\xf1\xb6\x33\xf6\x17\x3b\x16\xdb\xe2\x4f\x60\x85\xf6\x3b\x77\xf3\xc6\xc5\x9e\xe8\xb4\x67\x95\xd3\x9e\x6b\xe3\x5d\xdf\xa4\x5e\x19\x35\x6f\xb1\x0f\x2e\xcd\xbc\x3f\x77\x7c\x34\xbd\x3c\x4e\xd3\x7f\xc3\x38\x4d\xff\x52\xa6\xb9\x13\x23\xf6\x1c\x5d\x1d\x7f\x19\xbe\x5a\x17\xf5\xca\xa8\xfa\x13\x9c\x9d\xcb\xad\x9c\xf3\xed\x6b\x62\x9a\xfe\x95\xf9\x30\xfd\x2b\xf3\x81\x51\xfe\x79\x25\xca\x5a\x26\xea\x95\xd1\x1b\x5c\x2d\x4f\x11\x1a\x2e\x4f\x11\x3a\xfb\xa7\x15\x67\x6b\x1b\xad\xe5\x6e\x69\xe6\x96\xc2\xf6\xac\x9d\xd9\xf5\x26\xee\xdd\x1a\x5b\xbe\x5b\xaf\xb9\x57\x57\xee\xd4\x97\x47\xb4\x69\xa9\x02\xc3\x92\xaf\x65\x8d\xff\x19\xe9\x88\x14\xd1\xca\xb0\x7c\x12\x07\x96\x8f\x0b\x7a\x70\x48\xcb\x67\x72\x20\x7e\xde\x27\x27\xc6\x8c\x88\xe9\x8f\x9f\x97\xb0\xb4\xe9\x74\xe4\xcd\x86\x32\x67\x4a\x33\x3e\xc3\x37\x2c\x99\xfe\x99\x80\x11\x18\xf6\x9b\xfe\xd3\xd2\x9b\x21\x53\x9a\xf1\x1b\xfe\x61\xc9\xf4\x51\xad\x8f\xf6\x32\x22\xcb\x6e\x36\x99\x91\x2e\xa3\xc9\xf4\x9d\x0e\xbc\xe9\xb5\x38\x23\xa0\xfe\xfc\x15\xab\x66\x04\x8e\xd9\xf1\x15\x7f\xa7\x7d\x47\x1f\x4e\x9b\x81\xb4\xfc\x81\x91\x33\x02\x39\xc3\x6f\x36\xc5\xcf\x6f\x6a\xf7\x18\x4d\x17\x7c\xf6\xd9\x48\x5f\xc4\x68\x8a\x98\x52\x5f\xc0\x90\x2e\xf8\x6c\x1b\xe3\x1e\x6d\xbb\x89\xbd\xb4\x6c\xef\xeb\x0b\x18\xbe\x0b\x92\x6d\x6b\xad\xb5\x50\x5d\xa6\x16\x39\xc0\x69\x91\xb7\x96\x1e\x74\x15\x28\x83\xf5\xe6\xb0\xd0\x27\xe0\xb3\x90\x47\x11\x1a\x97\xf5\x2a\x37\x7e\x9b\x3e\x17\x9a\xff\xe4\xa6\x6d\xf6\xfc\x05\xfe\x93\x61\xc1\x22\xcd\x82\xf3\xf6\x2a\x87\x16\x54\x44\x3c\xd2\xe6\xec\x2f\x22\x17\xf8\x33\xf6\x6f\x35\x28\xb8\x7f\x03\x03\x87\x0f\x0f\x0e\x7e\xff\xfb\x1c\x2f\x92\x90\xac\x14\x80\x13\xbc\xc1\x2d\xbb\x74\x28\xcc\x9e\x75\xe6\x17\xae\xae\xfd\x5a\x52\x98\xd5\xda\xdb\x43\xa1\x50\x7b\xbb\xa6\xcd\x16\xbe\xf4\x2d\xe6\xec\x59\x4d\xd3\xb4\xc2\xa2\x61\xa1\xb0\x52\xbb\x79\xcb\x75\xe1\x34\xb9\xbd\xbd\xbd\xdf\x6e\x6f\x97\xe5\x81\xa7\x3e\xa8\x16\x7f\xfa\x4e\x2d\x7e\x8c\x10\x12\x77\xa0\x58\xeb\xf8\x92\xf8\xb5\x1a\xca\x62\x71\x68\x08\x60\x68\xa8\x58\x5c\xa9\xdd\xbc\x65\x99\xfe\xab\x5e\x8e\xe6\x36\xab\x28\xe7\x6a\x9a\x2c\xcb\x72\x7b\x7b\x84\x08\xde\x75\x9b\x5e\x1f\xfb\xd9\xa5\x5f\xa8\xf7\x24\x15\x45\xa3\x28\x8a\x32\x3b\x7b\x63\xfa\xe0\xf5\x7a\x3c\x1e\x0f\xc7\xb9\x35\x8e\x5b\xa9\xdd\xbc\x65\x05\x21\x18\x8d\xd7\x0b\x02\xef\x11\x3d\x5e\x51\x14\x3f\x1e\xaa\xf7\x08\x8b\xe6\xdc\x5a\xa3\xe5\x04\xbe\x3e\x7a\xf1\xef\x79\x41\x78\xe9\x9d\x7f\xab\xb4\x74\xc4\xb5\x55\x8d\x0e\x11\x3f\xce\x7e\x2c\x12\x07\xc0\xe9\x96\xc1\xc1\x35\xe8\x3e\x0b\xfd\xb3\xfd\x5f\x32\x87\x2b\x93\x79\x76\xb1\x30\xbf\xe0\x15\xae\x0b\x3c\xfc\xb3\xf4\x40\x7b\xc8\x35\xb3\x67\xb5\x1d\x51\x8f\x20\x78\xae\xb1\x29\x76\x36\x9b\xd3\x5b\x9b\x3b\x8b\x00\xc5\x42\x01\x8a\x1c\x57\xe0\x0a\xc5\x02\x14\xbf\x6a\x78\x00\xef\xe6\xb8\x7b\xd3\x09\x9e\x00\x89\x44\x42\xa1\x50\x28\x12\x21\x81\x80\x87\x17\xe8\x7c\x14\x66\xcf\xf6\x47\xea\x85\xb5\x3b\x07\x43\xf2\xef\xb7\x6d\x95\x43\x41\x00\x12\x24\x10\xe4\xf8\x20\x04\x79\xf0\x12\xef\x57\x4f\x80\x8e\x95\x13\xae\x5e\x15\x03\x98\x00\x32\x3f\x1f\x21\x24\x10\x08\x78\x96\x92\xf8\xbf\x26\xe0\xe1\xb7\x08\x9e\x40\xc4\xad\x45\x02\xc1\x84\x28\x8a\x81\x95\xf6\x78\x2a\xb4\x0d\x52\x71\x80\x7c\x98\x40\x3e\x9b\x02\x31\x2f\xd2\x6b\x3e\x9c\x03\x80\x3c\x09\xe5\x21\x9e\x8a\x43\xb4\x3f\x4a\xd5\xdf\xcd\x1e\x02\x05\x60\x7f\x36\x4d\xd5\x24\x0f\xf1\x6c\x56\x06\x7a\x24\x4b\xa3\xe9\x63\xe9\x19\xa9\xfc\xf6\x88\x7d\xe5\x8c\x23\x7d\x56\x7e\x77\xc4\xbe\xa2\xfe\xc7\xd1\xdc\xf1\xdc\x09\x53\xd2\xae\xf9\x3a\x9f\x1d\x17\x74\xaa\xb0\xf4\x71\x5e\xaf\x0c\xbd\x0a\xf8\x7d\xb1\xee\x24\xa7\xb7\x72\xc6\x66\xeb\xbe\x71\x51\xb7\xde\x19\xaf\xd3\x4f\xfa\xf5\x56\x5d\x52\xa4\xd0\xb8\x5f\xb7\x66\xc7\xeb\xd1\xaa\xde\x08\x1a\x21\xa3\xc1\xd8\x60\x84\x2d\xce\x58\x6f\x6c\xb4\x82\xe3\x3e\xfc\xfe\x20\xba\x86\x6f\x8d\xf5\xc3\x41\xfa\x6e\xf9\x35\xac\x34\x0c\x6f\xb0\x80\x1a\x6e\xd4\xcc\x7a\x47\x7e\x77\xbd\x1d\x9d\x0e\xcd\x34\xd8\x75\xd5\xfa\x71\xa2\xcf\x55\xeb\x6d\xaf\x05\x46\x48\x1d\x2d\xcd\x1d\x09\x19\xeb\xcd\x06\x47\x7e\x77\xa3\x1d\x9d\x0e\x57\x1b\x68\x7b\x83\xed\x47\x07\xe6\xd7\x1c\x49\xfd\x71\x69\xee\x48\xd8\xd8\x88\x5f\x1a\x0d\xc6\x06\x8c\xe7\xfe\xcc\xb0\xf9\xfa\x9f\x19\xe4\xed\x13\xab\x7f\x69\x70\x8b\x9e\x93\x21\xbd\xc2\x69\x16\x54\x40\xb3\x92\xf2\xd9\x92\x66\x70\xc3\x3c\xbe\xcf\x39\x7e\xf9\x6c\x69\xce\xbe\x78\x81\x57\x37\x4d\x58\x80\x4d\xe9\x63\xe9\xd1\x74\x99\xaf\xe4\xce\xd8\x57\x46\xec\xf9\xdc\x89\xdc\xf1\x9c\xfb\x85\xf3\xa0\xc1\x97\xa1\x22\xcb\x4e\x09\x0f\xea\xf9\x33\x4e\xdd\x05\x5e\xbd\x7f\x42\xfb\x86\x88\x9f\x5f\xeb\xd4\x1f\x94\xd0\x6e\x9d\xb9\xce\x59\xd0\xac\x8e\x71\xd0\x35\x2b\x81\x92\x2e\x51\x19\x8e\xf8\x2a\xc4\xea\xf0\x48\x47\x79\x49\x19\x07\xfd\xb3\x5f\x83\x1e\xf9\x06\xe8\x21\x2e\xde\xe2\x3b\x63\x9f\x37\x7d\x98\xdd\x62\x1c\xf9\xd0\x84\x25\xc8\x2f\x61\x4a\x8b\x02\x23\x63\x66\xb4\x62\x48\x98\x96\x29\x61\x1e\x9f\xb9\x79\xbc\x37\x21\x0f\x4e\x5c\xd3\xdf\x23\xbf\x84\x3a\x2c\xf2\x6b\xba\xaa\x72\xf2\x4b\x13\xe8\x68\xce\x76\x5c\x27\xe1\xc9\xc5\x0e\xa2\xdb\x57\x3b\xc9\xe9\x33\x9b\x6d\x3e\x53\xab\x80\x76\xe2\xd4\xbf\x4e\x1a\xdc\xa9\x4b\x93\x06\x7f\x90\x03\x13\xe2\x9f\xc2\x89\x37\xe0\x22\x38\x9c\x76\x91\xd3\x96\xea\xe9\x8b\x60\xf3\x17\x39\x6c\xc0\xeb\x89\x1c\xbd\x33\xe8\x74\x5b\x5c\x89\x9b\xc3\x6f\x94\x39\x6b\xe9\x73\x6c\xf5\x2b\x04\x77\x5a\x72\x04\xf5\x6f\x27\x2d\x38\xed\x7b\x13\x0b\x9a\xfc\xfa\x98\xf5\xbe\xc1\x59\x93\x8b\x3f\x66\x89\x4b\x3f\x66\xa9\x3f\x9f\xec\x7d\xea\xb9\xe7\xfb\xfb\x9f\x7e\xf6\x50\xdf\x7e\x58\x9f\x9e\x9e\x4d\x7f\x04\x9e\xef\xf5\x3d\x79\x68\xb0\xf0\x3a\xd9\x12\xbe\x6b\x54\x7f\xed\x87\x63\xc3\xc6\x2b\x47\x4b\xa5\x89\x52\x69\x62\x62\x62\x6c\x78\x6a\x6a\xd2\x98\x9c\x74\x7f\xf3\xa2\xc1\x6f\x45\x11\x5d\xca\xcf\xfa\x3b\x43\xb0\xa6\x0c\x71\x26\x7a\xea\xf5\x29\xc3\x7b\xea\x8d\x29\x83\xcc\xdc\x7a\xca\x9a\x8a\xf4\x79\x0d\x7c\xc5\x27\x06\xa1\xb7\x4c\x7a\xce\x90\x2c\xaf\xe1\xb3\xa0\x42\xf6\x48\x47\x85\x48\x9f\xef\x89\x3f\x95\xce\xd8\xbf\x35\x7c\xa6\x10\xe9\x93\x8c\x3a\x53\xec\xf3\x19\x9e\x63\xe9\xb2\xa7\x5a\x57\xf2\xe2\xc7\x50\x5d\xc9\x9b\x3b\x6e\x79\x57\xb9\x52\xd5\xa9\x1b\x9e\xf1\xe0\xad\xdf\x18\xfd\x7a\xdb\x8e\xf8\xae\x3d\xf7\x66\xff\xf8\xb1\x03\xfd\x2f\x2c\x1e\x16\x43\xc5\x1f\x55\x7f\x54\x1c\x2a\x0e\x3d\x71\xdf\x13\x43\xc5\xc0\xe6\x44\x3e\xb1\x39\x40\xa2\x09\x3d\x11\x25\xf1\xf8\x6b\xb5\xd7\xe2\xf1\x48\xc7\xdb\xb5\xb7\x3b\x22\x00\xd1\x44\x14\x1f\x9b\xe7\x7e\x7a\xae\x56\x83\x68\x2c\x16\x8b\x42\xed\xdc\x47\x1f\x7d\x74\xae\xd6\x3f\x32\x32\xf2\xbc\xe0\x85\xfa\x9e\xdf\xf4\xd4\x03\xc8\xfd\x09\x9e\xa7\x32\xf6\x6c\x02\x62\xd1\x70\x34\x06\x7a\x22\x11\x8d\x12\x42\x48\x34\x9a\x48\xe8\x10\x8e\xe9\xb1\x30\x00\x3c\x0a\x8f\x02\x80\x47\xd2\x25\x1d\xa0\xab\xb7\xb7\xb7\x15\x20\x80\x00\x4c\x55\x6a\x95\x29\x20\x22\x5f\x73\x5f\x50\x6a\x8a\x1c\x22\x84\x64\xda\xa2\x84\x44\xdb\x32\x84\x90\x5a\x81\x42\xa2\x6d\xa4\x2d\x4a\x66\x9d\x85\xda\x82\x33\x2b\x04\x77\xd5\x76\x05\x57\x3d\x03\x1e\xa5\xd2\x0b\x78\xcc\x87\xf5\xb0\x1e\x06\x68\x6e\xd3\xdb\x1a\x00\xbe\xbe\x81\xec\x7f\x0c\xe0\xce\xcc\x5e\xb9\xdb\xb5\x01\x80\x68\x2c\xe9\xf6\x4b\xc6\xa2\x00\xb1\xb0\x1e\x8e\x01\x10\x92\x20\x04\x00\x94\x76\x6c\xc1\x61\x60\x0d\x8f\x4d\x00\x39\x44\x44\x1e\x20\xd1\x93\xd9\x9d\x00\x80\x4e\x1d\xb5\x8f\xf5\x64\x32\x7b\x00\x6e\x4b\x66\xee\xfd\x03\x80\x48\xb8\x41\x0f\x01\xdc\xbe\x7b\xf7\xee\x3f\x02\xe8\xe8\xca\x64\x76\x00\x70\xcf\x48\x75\x02\x46\xcf\x64\xee\x04\xd8\x93\xc9\xb4\x6e\x46\x77\x61\x1a\x5c\xb9\x83\x06\x0a\xd3\x64\xc2\x08\xcd\x28\x4c\x00\x78\xee\xa1\x3a\x8c\x97\xfc\x93\x87\xee\x03\xc8\x37\xac\x6f\xc8\x03\xe8\xae\x97\x44\x32\x99\x8c\x01\xe8\xc9\x64\x32\xe1\x2a\x93\x00\xba\x24\x49\x1c\x6d\xcb\x1c\x06\xd0\x09\x21\x3a\x3a\xd3\xd1\xb5\xac\x28\xca\x2e\x54\xd2\x48\xba\xa2\x28\x0a\x80\xce\x07\x78\x1d\x40\x17\x49\x48\x77\x7d\x26\x5c\x2f\x1e\xac\xf5\xdc\xf6\x08\xd6\x36\xb5\xee\x01\x68\xc9\x64\x32\xdf\xc4\x47\xa2\x8e\x2f\x24\xbb\x5c\x67\x5b\x64\x45\xde\x02\xa0\xcb\x11\x59\x07\x78\x3c\x4c\xc2\x8f\x03\x08\xe2\x80\x28\x00\xec\xeb\xc9\xec\xfe\x16\x80\x8e\x19\x02\xd0\x85\x05\x48\x26\x93\xba\x0e\x10\x22\x22\x09\x81\xbb\xa6\x05\x00\x8e\x17\x31\xc1\x6c\x36\x3b\x00\xa0\xa7\x52\xf8\x48\x8b\xa7\x52\xa9\x94\x7b\xd1\x01\x9f\x53\xd9\x07\x01\xc8\x90\xc4\xf1\x00\x91\xe9\xe9\x69\xb4\x14\x45\x71\x00\x00\x52\x2f\xe2\x3a\x28\x85\x23\x2f\x02\x80\x1e\x0a\x6d\x4b\xd1\x41\xa3\x32\x2f\xc6\xc5\x3c\x5e\xa8\x65\xdc\x75\xbd\xd0\xdc\xdc\x1c\x01\x88\x34\x37\x37\x2f\xe0\x53\x52\x14\x45\x80\x74\x36\x9b\x6d\x06\x10\x77\xa5\x52\x32\x40\x87\xa2\xc8\x79\x80\xa8\xac\xc8\x51\x80\xbc\xd2\xae\xe4\x01\x52\xdb\xa8\xeb\x68\xb9\x5c\xc6\xda\xfe\xac\x9a\x02\x20\x89\x17\x92\xf4\x36\xa2\x2f\x99\xc9\xe4\x0b\x09\xba\x70\x02\xcf\x01\xf4\xa7\x53\xe9\x7e\x80\xcd\xb7\xbd\xb1\xaf\x01\x60\xa7\xa2\xc8\x83\x74\x28\x7b\x1f\xc2\xd1\xee\xed\x1d\x04\x88\x65\xb3\xd9\x41\x3a\x76\xac\xc9\x7b\x7b\xf7\x0e\x60\x66\xd3\xfb\x9b\xd1\x72\x6f\xef\xc3\x00\x3b\xb3\xd9\xec\xc3\xb4\x1f\xd6\x78\x65\x40\xe1\xf1\xe6\x3b\xdc\x49\x03\x0e\xe2\x30\x9f\xd9\xd6\xbc\xed\x19\x80\xc1\x8d\x1b\x37\x0e\x02\xe4\xb3\xd9\xbd\xbb\x01\x9a\xb3\x03\x38\x67\x43\x92\x9e\xc9\x00\xb4\xa7\xd5\xae\x3b\x00\x76\xa6\x52\xa9\x9d\x58\xeb\x52\xdb\x01\x3a\x3a\x93\xb1\x3c\x40\x47\xb2\x13\x47\xdb\x54\x2e\x97\x07\x01\xee\xee\xec\xec\xbc\x1b\xe0\x6e\x45\x51\xee\xc6\x4d\xa2\xc7\x62\x00\xe9\x44\x26\x99\xec\x84\x0d\xe1\x81\xf0\x06\xd0\x25\x4f\xbc\xbb\x7b\x1b\x28\x23\x43\x12\x4f\x93\xdf\x7b\x18\x57\x60\x10\x27\xa2\x3d\x9d\x56\xbf\x89\x33\xd8\x79\x5b\x1e\x60\x50\xf2\xfb\x30\x33\x69\x43\x0c\x4d\xb6\xfe\xde\x76\xcc\xba\xa5\xb5\x05\x17\xbc\xfb\x5e\x45\x06\x88\x8a\x74\x01\x50\x46\x01\x1a\x89\xfa\x44\x0f\x5e\xee\xa0\x1d\xe0\x45\x77\x73\x64\xc2\xb1\x64\x32\x16\xce\xc4\xc2\x64\x7d\x8a\xfe\xf5\x9e\xa0\x7f\xb5\xde\xda\x89\x5a\x6f\x6d\x71\x01\xf0\xa8\x70\xb7\x91\x5b\xab\xa1\x96\x90\x79\x32\x4f\x95\x0b\xb4\xed\x92\xbb\x56\xb4\x2d\x1c\xfe\x5c\x5c\xc0\x4b\x63\x68\x0b\x5a\x06\x49\x90\x5a\x6e\x71\x9d\xcd\xbb\x07\x44\x90\x9e\x03\x64\xf1\x32\xef\x9e\x09\x35\xb2\x9a\xda\xa2\x92\xee\xdf\x1a\x36\x02\x6c\x09\x35\x62\x6d\x41\xfc\x1c\xb7\x75\x23\xd6\x30\xde\xe7\xae\xc9\xa5\xf0\x0a\x97\x16\x6b\x68\x41\x93\x20\x04\x4d\x69\x6d\x1e\x03\x62\x0e\x41\x1a\x61\x8b\x7b\x42\x2c\xd0\x0b\x6d\xc3\xe0\x35\xda\xa1\x46\x3b\xe0\x10\x10\x6c\xaa\xb9\x38\x2e\x58\x74\xdb\xb0\x14\x74\xc1\x6d\x11\x3f\x08\x50\x93\xe8\xc1\xa2\x73\xf4\x93\x43\xcc\xe3\xcb\x24\x24\x1f\xef\xc5\x83\x25\x9e\x4a\xa9\x77\x00\x2c\xc8\x74\x5b\x88\x62\x1e\x57\x8c\xf4\xea\xbd\x04\x20\x91\xc9\x64\x12\x00\x8f\xec\xe3\xf6\x3d\x02\xb0\x73\xf7\xee\xdd\x71\xdc\x5b\xf8\x0f\x71\x65\x01\x77\xdc\xb6\x6c\x36\xb5\x0d\x60\x88\xe3\xb8\x21\x80\x36\x04\x20\x95\x7a\x14\x77\x78\xcf\x03\xd9\x07\xbb\x01\xba\x1f\xcc\x3e\xd0\x03\x00\x57\x71\x83\xe3\x6e\xd6\xe9\x0a\xb4\xe1\x98\xd3\xcd\x07\xd3\xcd\x00\xc0\xd7\xd1\xb6\x68\x34\x4a\xc7\x80\xb3\x41\x42\x9a\x7b\x28\x07\xb9\x7a\xd4\x49\xeb\x7c\x00\x90\x40\xfe\x37\xdf\xc4\x19\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\x06\x83\xc1\x60\x30\x18\x0c\xc6\xff\x80\x7b\x26\xa1\xd0\x50\xe0\x0a\xff\x15\x00\x00\xff\xff\x49\x31\x91\xda\x00\x80\x00\x00")

func dataFirmwareBytes() ([]byte, error) {
	return bindataRead(
		_dataFirmware,
		"data/firmware",
	)
}

func dataFirmware() (*asset, error) {
	bytes, err := dataFirmwareBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/firmware", size: 32768, mode: os.FileMode(420), modTime: time.Unix(1494444115, 0)}
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
	"data/firmware": dataFirmware,
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
	"data": &bintree{nil, map[string]*bintree{
		"firmware": &bintree{dataFirmware, map[string]*bintree{}},
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

