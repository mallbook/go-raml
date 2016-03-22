// Code generated by go-bindata.
// sources:
// ../templates/client_helper_resource.tmpl
// ../templates/client_resource.tmpl
// ../templates/generic_main.tmpl
// ../templates/oauth2_middleware.tmpl
// ../templates/oauth2_scopes_match.tmpl
// ../templates/python_client.tmpl
// ../templates/python_client_utils.tmpl
// ../templates/python_server_resource.tmpl
// ../templates/server_main.tmpl
// ../templates/server_python_main.tmpl
// ../templates/server_resources_api.tmpl
// ../templates/server_resources_interface.tmpl
// ../templates/struct.tmpl
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

var _TemplatesClient_helper_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\xd1\x6f\xda\x3e\x10\x7e\xc6\x7f\xc5\xfd\x22\xfd\x26\x1b\xd2\xb0\xa9\x4f\xad\x84\x26\xad\x55\x35\x4d\x5a\xb7\x41\xa7\x3d\x54\x55\x6b\xc8\x51\xb2\x12\x3b\x38\x17\x36\x84\xf8\xdf\xe7\x73\x42\x0a\xac\xab\xfa\x32\x69\x2f\x21\x3e\xdf\xe7\xfb\xbe\xcf\x77\x64\xbd\x3e\x82\x14\xa7\x99\x41\x88\x26\xf3\x0c\x0d\xdd\xce\x70\x5e\xa0\xbb\x75\x58\xda\xca\x4d\xb0\x8c\xe0\x68\xb3\x11\x85\x9e\x3c\xe8\x7b\x84\x3a\x49\x88\x2c\x2f\xac\x23\x90\xa2\x13\x8d\x57\xe4\xb3\xfc\x0b\x9a\x89\x4d\x33\x73\xdf\xff\x5e\x5a\xc3\x81\xcc\xf2\xd3\x20\xf5\x67\x44\x05\xbf\x4f\x73\xe2\x1f\xca\x72\x8c\x84\x12\x62\x5a\x99\x09\x04\x1c\xbe\xb3\xe9\x4a\xa6\x9a\x34\x64\x86\xd0\x4d\xf5\x04\xd7\x1b\x05\x32\xb3\xc9\x10\x75\x8a\x2e\x06\x74\xce\x3a\x05\x6b\xd1\x19\x87\x05\x9c\x0e\x80\x6b\x25\x1f\xb5\x2b\x67\x7a\x1e\xe0\x4a\x74\xb2\x69\xd8\xfd\x6f\x00\x26\x9b\x73\x7a\xc7\x21\x55\xce\xf0\x32\x00\x45\x67\x23\xb6\xb1\x40\x3f\xb9\xc4\x1f\x75\x15\x39\x56\x31\xe7\x89\x4d\xc3\x2e\xb5\x43\x5c\x7c\xcb\x68\x16\x08\xe6\x48\x33\x9b\xc6\x50\xb9\xf9\x88\x1c\x94\xe4\xbc\xe0\x18\x0e\x79\xc7\x8d\x51\xc0\xc2\x93\xb3\xf0\x1e\xc3\x2c\x54\x28\x21\xd7\xc5\x75\x8d\xbc\xd9\xc3\x2c\xca\xcf\xda\xe9\xbc\x39\xd5\x6b\xef\x06\xf8\x10\xcb\xc2\x9a\x12\xf7\x0c\xf0\x64\x5a\x0f\x0e\x0c\x7c\xa9\x03\xa2\xd3\xef\xc3\xc4\xa1\x26\x04\x9a\x21\x38\x5c\x54\x58\x12\x3b\xb3\x68\xcf\x0e\x0c\x82\x3b\x61\xf3\xd0\x80\xde\x96\x74\x0c\x4c\xe9\xe5\xa5\xa7\xd6\xc1\x43\x0c\x4b\xae\xe1\xb4\xf1\xad\xb5\x75\x87\x21\x1d\x26\x93\xbc\x0f\x91\x64\x84\x24\x7d\xaa\xef\x9d\x64\x54\x78\x67\x68\x2a\xa3\xff\x97\x51\xbc\x54\xaa\x3e\xab\x29\x50\x5b\x9e\x9c\x5b\xe9\xc1\xaa\xbd\xc0\x71\x95\xcd\xd3\x2f\x15\xba\xd5\x28\xf8\x5a\x37\xd9\xd3\x77\xa0\x1a\xef\xd7\x41\xc7\x1c\x4d\x6d\x28\x0c\x06\xf0\x7a\x57\x4b\x14\xd5\x95\xc7\xba\xc4\x70\x34\xcb\x88\xde\x46\x4f\xe9\x0a\xe5\x18\xfb\x98\xdc\x1b\xc0\x83\xb7\x2e\x1a\x44\xfe\xf9\xa8\x4b\x2e\x15\x47\x5f\x45\x7b\xb2\x5a\xd8\xf5\x29\x13\x6a\x97\xea\xe8\xcd\x0d\x8b\xec\xf7\xcf\xf9\x06\x1d\x16\x7e\x66\xb9\xe7\x86\x17\x67\xc7\xc7\x27\x27\x5c\x18\x05\xad\x0a\x84\x90\xc0\x43\x97\x5c\xf9\x07\x43\x9a\x89\xf9\x30\xfa\x74\x09\x76\xe9\x6f\x25\x4b\xd1\x7b\xd2\x06\x6b\xeb\x24\x41\x97\xb1\x0a\x76\xf2\xa5\x6f\xcc\xeb\x1b\x1e\x9b\xdd\x86\x6c\xc8\xd6\x1b\xb2\xad\x25\xbb\xa4\x92\x0b\xeb\x72\x4d\xf2\x2e\xba\xf3\xf2\xc2\x56\xa0\x78\x7c\xe2\x97\x3e\xa8\x1e\x27\xae\x25\x76\x85\x3f\xe9\x37\x62\x1c\xfc\x03\x31\xde\xfa\xbb\xc4\xbe\x9a\xfc\x29\xcf\x2a\xf3\x8c\x6b\x7b\x18\x39\x6e\x48\xa8\x9a\x1d\x93\xa3\xb2\x1d\xb4\x50\xde\x8f\x52\x89\xcc\xa7\xb7\xcb\xa6\xe7\x03\x71\xd3\x99\xfe\xff\xe9\xb9\x21\x6b\xe7\xab\x4b\x30\x08\xf7\x2e\xa9\x54\x62\x67\x06\x0f\xd4\xec\x1b\x5d\x99\x67\xac\xde\xc3\xfc\x43\x6a\x0e\x68\x36\x53\xbe\x1d\xe5\x9d\x1e\xd8\xbf\xfc\x6d\x1e\x1f\xc1\xdf\x41\x34\x69\xf8\xd4\xfd\x0a\x00\x00\xff\xff\xf3\x16\xb8\x6d\x14\x07\x00\x00")

func TemplatesClient_helper_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesClient_helper_resourceTmpl,
		"../templates/client_helper_resource.tmpl",
	)
}

func TemplatesClient_helper_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesClient_helper_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/client_helper_resource.tmpl", size: 1812, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesClient_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x4d\x4f\xdb\x40\x10\x3d\x7b\x7f\xc5\xd4\x4a\x91\x4d\x8d\x73\xaf\x94\x0b\x90\x7e\x48\x14\xd1\x00\xed\xb1\x32\xf6\x98\x18\x12\xaf\xb3\xbb\x0e\x42\xae\xff\x7b\x67\x76\x6d\x27\x21\x2d\x52\x91\xda\x5e\x9a\x53\x76\x3e\xdf\xbc\x79\xbb\x6e\x9a\x23\xc8\x30\x2f\x4a\x04\x3f\x5d\x14\x58\x9a\x6f\x0a\xb5\xac\x55\x8a\x3e\x1c\xb5\xad\xa8\x92\xf4\x3e\xb9\x45\x70\x4e\x21\x8a\x65\x25\x95\x81\x40\x78\x7e\x89\x66\x3c\x37\xa6\xf2\x85\xe7\xf9\x58\xa6\x32\x2b\xca\xdb\xf1\x9d\x96\xa5\xb5\xe4\x4b\xe3\x8b\x50\x88\x54\x96\xda\x26\x28\x29\xcd\xf5\xec\x0c\x26\xe0\x37\x4d\x7c\x9c\x68\xbc\x9e\x7d\x6c\x5b\x1b\xd4\x34\xa3\xa4\x2a\xce\x93\x25\xc2\xdb\x09\xc4\xfc\x87\x9a\x9b\xc7\x0a\x81\x62\xdd\x11\xb4\x51\x75\x6a\xa0\x11\x9e\x43\x03\xdc\x3d\x3e\x71\xc8\x5a\x21\xf2\xba\x4c\xe1\x1c\x1f\x86\x8c\x20\x84\xc3\x4d\x3a\xe7\x71\xf5\x12\x1f\x82\xc1\x1a\x92\x31\xee\xca\x4d\xb6\x0b\x36\x2d\x21\x46\x53\xab\x12\x52\x2e\xde\x34\xa0\x92\x92\x98\x18\xdd\x47\x30\x5a\x5b\x98\x9f\xd0\xcc\x65\xa6\x81\xa0\x6e\xb9\x73\xf6\xe7\x1c\x30\x5a\xc7\xef\x08\xd3\x89\x5c\x2e\xa9\xa2\x8d\x1b\x8f\x69\x20\x72\x33\x9e\x06\xcb\x8c\x4c\x0c\x3b\x48\x19\x69\xcf\x01\xc1\xb2\x51\x5d\x83\x6e\x18\x6b\xb9\x48\x54\xb2\xd4\x14\xe0\x8e\x33\xd4\xd5\xb1\xcc\x1e\x6d\xb9\x22\xa7\xd9\x60\xcb\x0a\xbe\x4f\x3d\xa3\xbe\xd3\xa1\x1d\x8f\x9d\xb4\x12\x8c\x50\x29\xa9\x42\xa6\x65\xa5\x6d\x59\x86\x7c\x53\x17\x8b\xec\x73\x8d\xea\xf1\xd2\x28\xda\x67\xb0\xe2\xff\xae\x2b\x71\xc5\x7a\xa1\x36\xb8\xe2\x36\x5f\x50\xdd\x80\xff\x7e\x7a\xc5\x5d\x68\xe5\xbf\x44\x00\xeb\x44\x41\x0d\xfb\x88\x1d\x01\x94\x4a\xb4\xa4\x0a\x13\x83\xa0\x90\x3a\x92\x60\xe4\xcd\x1d\xa6\x86\x5c\x64\x88\x80\xb0\x32\x3a\x3b\x00\x6d\x78\xe6\x82\x02\xdb\x3c\x82\x5e\x59\xbb\x00\xac\x8a\x2f\x12\x33\xef\x40\xbc\xe9\x1b\x0e\x40\x86\x08\xeb\xed\x58\x88\xa0\x2c\x16\x34\xaa\xc7\x73\x52\xdb\x57\x13\x36\x30\x4d\xcf\x4d\xd8\x49\xa5\xb6\xd9\x16\xaf\x8b\xc7\x85\xc6\x8d\x7b\xd7\x77\x04\x04\xc7\xde\x32\xcf\xb3\x2c\xe4\x52\x01\xa9\xcb\x8a\xcb\xa9\x69\x8e\x49\x86\x4a\x53\x77\xe8\x7f\xc4\x47\xfc\xc1\x9a\xe3\x4b\x34\x01\xc5\xd3\x4d\x8b\x2f\x2b\xda\x96\xc9\x03\xff\xf5\x9a\x08\x59\x87\xa1\x4d\xe8\xb8\xcd\x24\x98\xf9\x40\xad\xe5\x54\x57\x03\xa9\xbd\xfe\xe3\x53\x19\x50\xc8\x3f\x18\xdd\xf3\xe8\x0d\x42\x05\x0c\x2b\xe6\xca\x74\x07\xa5\xc6\x20\x14\xcf\xa9\x8a\x2b\x6d\x7a\xbb\x91\xf8\xf5\x61\x81\x9c\x22\x3d\x47\xa8\x82\xa1\x62\x18\x3b\x53\x70\x50\x87\x62\x83\x6e\xab\x86\x2b\x40\x30\xc5\x13\x80\x2e\xf6\xa9\xec\x4f\xa7\x67\xd3\xab\xa9\x6f\x4b\xbc\x54\xbd\x5d\x8d\x41\xc0\x7f\x5c\xbf\xfb\xbb\xf8\xdb\xba\x73\x0f\xea\x9e\xe2\xb6\xf6\xf1\xf2\x57\xc4\xa1\xdc\x52\x76\x26\x89\xea\xaf\x85\x99\x73\x68\xe0\xdb\x4c\xde\x1e\x7d\x74\x7e\x97\xf3\x9f\x50\x1e\xed\xbe\x37\xab\x0e\x69\xdb\x1e\xec\xe6\x38\xc7\x77\xb8\x92\x67\xf2\x01\x15\x27\xf6\x04\x44\x3d\xd1\x51\xbf\xbd\xff\xb7\x6f\xef\xf6\x6d\x0e\xfc\x99\xdd\x9c\x76\x0e\x3f\x02\x00\x00\xff\xff\xec\xe0\x3a\x34\xc9\x08\x00\x00")

func TemplatesClient_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesClient_resourceTmpl,
		"../templates/client_resource.tmpl",
	)
}

func TemplatesClient_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesClient_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/client_resource.tmpl", size: 2249, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesGeneric_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4f\xcd\x4b\x2d\xca\x4c\x8e\xcf\x4d\xcc\xcc\x8b\x2f\x49\xcd\x2d\xc8\x49\x2c\x49\x55\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x00\x49\x70\x71\xa5\x95\xe6\x25\x83\x99\x1a\x9a\xd5\x5c\x5c\xb5\x5c\xd5\xd5\xa9\x79\x29\x40\x55\x80\x00\x00\x00\xff\xff\xdc\x57\x73\x81\x49\x00\x00\x00")

func TemplatesGeneric_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesGeneric_mainTmpl,
		"../templates/generic_main.tmpl",
	)
}

func TemplatesGeneric_mainTmpl() (*asset, error) {
	bytes, err := TemplatesGeneric_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/generic_main.tmpl", size: 73, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesOauth2_middlewareTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x92\xcf\x6e\xa3\x30\x10\xc6\xcf\xf6\x53\x8c\x7c\x32\xab\x0d\xfb\x47\x7b\x5a\x29\x87\xbd\x6c\x39\xb4\x15\x4d\x53\xf5\x58\x59\x61\x92\x58\x01\x43\x07\x53\x9a\x22\xde\xbd\x63\x4c\xaa\xa4\x0f\x50\x2e\xd8\x33\xfe\x3e\xff\x3e\x79\x86\x61\x01\x05\x6e\xad\x43\x50\xb5\xe9\xfc\xfe\xf7\x53\x65\x8b\xa2\xc4\xde\x10\x2a\x58\x8c\xa3\x6c\xcc\xe6\x60\x76\x08\xc3\x90\xe6\x71\x79\x6b\x2a\xe4\x86\xb4\x55\x53\x93\x07\x2d\x85\x2a\xeb\x9d\xe2\x9f\x43\xff\x63\xef\x7d\xa3\x64\x22\xe5\xb6\x73\x9b\xa0\x8a\xc7\x6f\x7a\xd2\x0e\x5f\x3d\x84\x7e\x9a\x19\xc7\x97\x50\x72\xb1\x83\x41\x0a\x42\xdf\x91\xbb\x28\xff\x67\x1f\x1d\xcc\x74\x1f\xeb\x2b\x6c\x9b\xda\xb5\xf8\x48\xd6\x23\x7d\x07\x82\x6f\x73\xfd\xb9\xc3\xd6\x27\xec\x03\xf3\x37\x0c\x60\xb7\x90\x66\x68\x0a\xf6\x67\xe8\x53\x23\x64\xcd\x0a\x82\xbf\x4b\xa0\xb9\x9d\x5e\xa1\xd7\x8a\x81\xe7\x6d\xe4\x56\x89\x14\x82\x3d\x4e\x82\xe5\x12\x94\x0a\xa4\x42\x70\xe8\x34\x27\xeb\x7c\xe9\xb4\xfa\xc7\xfd\x9a\xec\x9b\xf1\xb6\x66\xfe\x78\xa1\x6d\x01\xab\xc6\x1f\x27\x13\xd1\xa7\x13\x71\xb4\xd7\x7f\x7e\xfe\x9a\xaa\x31\x31\xaf\xc6\x73\x6a\x2c\x5b\x9c\xd0\xef\x3a\xa4\x63\x6e\xc8\x54\xed\x39\xbf\xaf\x0f\xe8\x22\xfd\xc3\xea\x3a\x9e\xd2\xc9\x47\x84\x33\xd5\xa7\x1c\x51\xf8\x55\x29\x5c\x11\xa0\x85\x08\x2f\x9f\xde\x23\xbd\x60\xb6\x5e\xe7\xba\xe7\x47\x63\xd9\x98\x48\x1e\xa3\x30\x83\xe1\x60\x18\xb6\xf7\x00\x00\x00\xff\xff\xad\x0c\xd8\xa2\x90\x02\x00\x00")

func TemplatesOauth2_middlewareTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesOauth2_middlewareTmpl,
		"../templates/oauth2_middleware.tmpl",
	)
}

func TemplatesOauth2_middlewareTmpl() (*asset, error) {
	bytes, err := TemplatesOauth2_middlewareTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/oauth2_middleware.tmpl", size: 656, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesOauth2_scopes_matchTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\xc1\x6a\xeb\x30\x10\x3c\x4b\x5f\xb1\xf8\x24\x3d\x12\xbf\xb6\xf4\x54\xc8\xb5\xe4\x54\x42\x13\xe8\xa1\x94\x20\xec\x75\x62\xea\x48\xee\x4a\x8e\x0b\x42\xff\x5e\x29\x0a\x69\x4c\xba\x60\x6c\xcd\xcc\x8e\xbd\xb3\xf6\x7e\x0e\x35\x36\xad\x46\x28\x8c\x1a\xdc\xfe\x61\x6b\x2b\xd3\xa3\xdd\x1e\x94\xab\xf6\x05\xcc\x43\xe0\xbd\xaa\x3e\xd5\x0e\xc1\xfb\x72\x95\x1f\x5f\xd4\x01\x23\xc1\xdb\x43\x6f\xc8\x81\xe0\xac\xe8\xcc\xae\x88\x37\x8d\xee\xff\xde\xb9\xbe\xe0\x92\xf3\x66\xd0\x55\xea\xca\x72\xa1\xf1\xdb\x41\x22\xcb\xa5\xd2\x75\x87\x24\x27\x27\xf0\x9c\x11\xba\x81\xf4\x04\x7e\x8e\x26\x22\x39\x89\x31\xe3\xaf\x68\x7b\xa3\x2d\xbe\x51\xeb\x90\x66\x40\xf0\xef\x8c\x7f\x0d\x68\x9d\x8c\x3e\x70\xae\xa3\x22\xc8\xf3\xc0\xfb\x87\x75\xd4\xea\x1d\xbf\x90\xaa\xeb\xcc\x88\xf5\x3a\xf3\x4f\x8b\x8b\xc4\xa7\x6f\xce\x70\x08\x10\x2e\x0d\x8d\x21\xd8\xce\x40\x9d\xc4\xa4\x74\x8c\x64\xea\xf1\xfb\xe2\x2b\xf9\x95\xda\xfe\x25\x4b\xd5\x36\x51\xb6\x58\x24\xeb\x18\x02\xbb\xa6\x52\x68\xe5\x1a\xe9\x88\xcb\xcd\x66\x25\xc6\x38\xaf\xbc\xe9\x4f\x95\xa3\xbb\xa1\x02\x9f\x9e\x18\x4b\xd7\x58\x9e\xd2\x5b\xa2\xaa\x91\xc4\xe3\xdd\xbd\xe4\x2c\x48\x1e\x57\x9a\xfe\x08\xd4\xf5\x69\xf1\x3f\x01\x00\x00\xff\xff\x81\x64\x33\xc0\x1e\x02\x00\x00")

func TemplatesOauth2_scopes_matchTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesOauth2_scopes_matchTmpl,
		"../templates/oauth2_scopes_match.tmpl",
	)
}

func TemplatesOauth2_scopes_matchTmpl() (*asset, error) {
	bytes, err := TemplatesOauth2_scopes_matchTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/oauth2_scopes_match.tmpl", size: 542, mode: os.FileMode(420), modTime: time.Unix(1458573475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_clientTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\xcf\xce\xda\x30\x0c\xc0\xef\x79\x0a\xab\xe2\x00\xfa\xbe\xaf\x0f\x80\xc4\x01\x10\x93\x90\xb6\x09\xb1\x3f\xd7\x28\xb4\x2e\x44\x6b\x92\xce\x49\x98\x50\x96\x77\x5f\xd2\x40\xd5\xc3\x7c\x68\x6d\xcb\xfe\xf9\x5f\x42\xf8\x80\x16\x3b\xa9\x11\xaa\xe1\xe1\x6e\x46\xf3\xa6\x97\xa8\x1d\x77\xa8\x86\x5e\x38\xac\xe0\x23\x46\x26\xd5\x60\xc8\x01\xe1\x6f\x8f\xd6\x59\xd6\x91\x51\xf0\x8c\xf4\x4e\xf6\x16\x9e\x11\x17\x2f\xfb\x96\xa7\x28\x7a\x70\xeb\x48\xea\x2b\x63\xbb\xed\xb7\x03\xff\x71\x3e\xc2\x06\xaa\x10\xea\x9d\xb0\x98\xac\x18\x2b\xc6\x58\xd3\x0b\x6b\x61\x3f\x92\xd6\x0c\x92\xa4\x76\x80\x73\xa9\xa5\xe3\x7c\x69\xb1\xef\x56\xc5\x9f\x25\x9b\xb5\xa7\x3e\x91\x5e\x50\x16\x02\x90\xd0\x57\x84\xc5\xaf\x77\x58\xdc\x61\xbd\x81\xfa\x34\x8e\xf2\x05\xd3\xb7\xb5\x90\xfa\x7f\x81\x43\x58\xdc\xeb\xe2\xff\x2a\x14\xc6\xb8\x1c\x3d\x27\x41\x42\xd9\x18\x67\xa5\xaa\xaa\x9a\x91\xbb\x8c\xee\x32\x3b\x45\x7f\xf2\xba\xd9\x1b\xa5\x52\xcb\x13\x3c\x4b\x46\x75\x31\x86\x80\xba\x9d\xb9\x8f\x0e\xa4\x05\x35\x16\x85\xce\x50\x69\xe2\x27\xd2\x25\xc6\xa2\x1f\x74\x3b\x18\xa9\xdd\x2c\x29\x95\x9f\x74\x4f\x32\x0d\x3c\xcd\xfe\x56\x92\xce\x68\x8d\xa7\x06\x4f\xc2\xdd\x66\x89\x25\x38\x7f\xdf\xfe\x73\x8c\x65\x31\x86\x71\xde\xd5\x94\x44\xe8\x3c\xe9\xe9\xbc\xf5\xd4\x21\xfc\x85\xef\xe6\xb3\xf9\x83\x94\x56\x95\xa0\x65\x5b\xe7\x2d\x5d\xd3\xb6\xde\xe1\x86\xa2\x45\xb2\x9b\xe7\x7f\x95\x8f\x91\x86\xcf\x4b\xc9\x2f\x2b\xab\xf9\xf5\xfc\x0b\x00\x00\xff\xff\x09\x39\x05\x4e\x66\x02\x00\x00")

func TemplatesPython_clientTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_clientTmpl,
		"../templates/python_client.tmpl",
	)
}

func TemplatesPython_clientTmpl() (*asset, error) {
	bytes, err := TemplatesPython_clientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_client.tmpl", size: 614, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_client_utilsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xcd\x6e\xdb\x46\x10\xbe\xeb\x29\x06\x0c\x14\x93\xb5\xa8\xd0\x52\x90\x1f\x21\x6a\x0e\x45\x51\xf4\x92\x93\x73\x0a\x02\x79\x45\x0e\x2d\xd6\xe4\x2e\xbd\xbb\xb4\xc3\x04\x79\xf7\xce\xec\x92\x14\x29\x2b\x69\x0e\x35\x0c\x6a\xb5\xf3\xff\xcd\x7c\x43\x7d\xfb\x16\x43\x86\x79\x21\x11\x82\xba\xb5\x07\x25\x77\x69\x59\xa0\xb4\xbb\xc6\x16\xa5\x09\x20\xfe\xfe\x7d\x56\x54\xb5\xd2\x16\x32\x61\xd1\x16\x15\xf6\xdf\xdd\x79\x36\x23\x73\xd8\x37\x45\x99\xed\xee\x1b\xd4\xed\xce\x58\x5d\xc8\xdb\xd0\x7f\xa9\x85\x16\x95\xd9\x7e\x50\x12\xa3\xcd\x0c\xe8\x2f\x08\x02\xf7\xe9\x4c\xc0\x69\x81\xd3\x42\x8b\xda\x49\x0a\x59\x37\x16\xe0\x66\xec\xe1\x06\x36\x90\x15\xa9\x2d\x94\x14\xba\x75\x6a\xaa\xb1\xac\xe7\xc3\xc1\x63\x61\x0f\x90\x2b\x5d\x09\x0b\x37\xef\xef\xb0\xdd\x3e\x88\xf2\x39\x7d\xae\xf8\xb0\xba\x99\xc4\x2e\x72\x18\x3b\x87\xc2\x00\x67\xe8\x13\xe4\x3f\x8d\xb6\xd1\x92\xf4\x67\xee\xea\xde\xc0\x16\x82\xf7\xde\x98\x82\x00\xf9\x5d\x00\x96\x58\x51\xb2\x13\x57\x47\x17\x64\x73\xb9\x65\x45\xb8\x84\x60\x1b\xd0\x93\x32\x0d\xd9\x26\xe2\x9b\xe7\x9d\xeb\x2e\xd2\xbd\xf9\xb4\x89\xaf\x3e\x77\x70\xde\xa2\x44\x4d\x68\xef\x74\x9e\xae\xd7\xeb\xb7\x61\xb6\x80\x52\xa5\xa2\xdc\xd9\xaf\xdb\x6b\xdd\x9c\x62\xd9\xeb\x43\xa7\xef\x5a\xd3\xa1\x31\x82\xd4\x1b\x65\x54\x0b\xb7\x12\x6c\x5b\xa3\xbb\xe9\x5d\x93\xa0\x31\xe8\xbf\x7a\x17\x5f\x09\x15\x46\xcb\x52\xcc\x85\x47\xdd\x1e\x50\x3f\x16\xa4\x56\x09\x7d\x07\xc2\x40\x63\xd3\xd9\xb8\x21\x3e\x4a\x9f\x49\xd7\x1e\x17\xd0\x27\xb4\x04\xfc\x42\xdd\xbc\x59\x25\xc9\x9b\x38\x79\x19\x27\xab\xeb\x55\xb2\x49\xf8\xff\x32\x79\x4d\xcf\x69\xaf\xac\x6e\x8f\xa0\x52\x2a\x7d\xb2\xc7\xcb\x71\x51\x9c\xf5\x72\x38\xe4\x5a\x55\x7c\x30\x56\x54\x75\x98\x45\x83\x09\x96\x06\x7f\xc5\x01\xd5\x76\xd6\x07\x7e\x49\xb1\xb6\x70\x4d\x08\xfe\xa9\xb5\xd2\x47\x5f\xb5\x30\x66\xd6\x0f\x99\x54\x96\x66\xab\x90\x64\x2b\x53\xe4\x2e\x4e\x42\x44\xa3\x81\x13\x0c\xe9\xe0\x2f\xbc\xf8\xa0\x3c\xc1\x5c\x58\xa0\x91\x73\x08\xaa\xfd\x3f\x98\x12\x82\x7f\x91\x74\xae\x97\x17\x30\x77\x5d\xa4\xb4\xa2\x5f\x0d\xca\x87\x51\xe0\xb3\x75\x87\xbf\x65\x4b\xfe\xb4\x4d\x5d\x62\x18\x7d\xda\xac\x3f\x47\x93\x71\x0d\x2f\xe6\xc9\xcb\x2c\x9e\x27\x2b\xff\xb8\xe6\xc7\x66\x78\xcc\x0d\x65\x36\x81\x37\xcc\x96\x2d\x0a\x4d\xb9\x2c\x2b\x25\xed\x81\x0f\x99\x68\xf9\xe3\xa0\x1a\x7f\x5f\xc8\xc6\x22\x9f\x0c\xa6\x4a\x66\x8b\x89\x03\xd8\x0d\xa4\xe0\xcc\x78\x30\xc7\xac\x88\x18\x00\xc7\x9d\x1d\x5d\xa4\x4d\xc9\x8a\x2a\xcf\x0d\xda\x90\xab\x1a\x69\x4e\xa9\x33\x61\x06\x43\xbc\xf9\x11\x39\x36\x03\x11\x1c\x4b\x4c\x6b\x2c\xf1\xbf\x4f\x66\x31\x62\x46\x07\x52\x32\x81\x8c\xa4\x5d\x0f\x73\xf8\x78\xfd\x07\xf8\xe4\x96\x4e\xe5\xef\xdc\x8b\x32\x85\xc6\xf5\xef\x20\x1e\x10\x84\x6c\x07\xf7\x94\x67\xae\x16\xf0\x88\x27\x0c\x75\xa1\x4f\x88\x39\x84\x9f\xd4\x79\x8e\x39\xcf\x0e\x42\x66\x25\x02\xb7\x06\xf6\x48\xfc\x44\xb8\x7a\xfb\x3a\x81\x4a\x19\xda\xac\xad\x2b\x90\xdc\x22\x6f\x48\xa9\xa6\xd9\x30\x97\x69\x1d\x8f\xed\x96\x63\xa2\x72\x45\xae\xe9\xf0\xce\x09\xa7\x7c\x7b\x06\x1f\x8d\x33\x5a\x91\x83\x54\x34\xa6\x8b\xcc\x10\xc8\x8b\x1e\x01\x28\x51\xd4\xe4\xa9\x9d\xd8\x5a\x9a\x59\x37\xaf\xd5\x9d\x9b\x56\x17\x49\x63\x5d\x0a\x9a\x77\x8e\xb8\x65\xbf\xd1\x71\x84\x7f\x46\xfc\xb3\xce\x46\xc3\xdf\x0d\xbe\xcf\xb9\xc3\x8b\x12\x2a\x8b\xdb\x83\x35\xe2\x81\x96\xdb\xc2\x57\x3b\xbe\x72\x5d\x12\xa5\x3d\x9d\x0c\xbe\xee\x41\x1c\x63\xe5\x12\x70\xfd\x71\x39\x58\xca\xbd\xda\x15\x26\x33\x76\x9a\x6c\xd7\xda\xd8\xe9\x77\xfe\x7f\x52\xdb\x44\x7d\x12\x77\xaa\x7b\x1c\x58\x4f\xa1\x33\x54\xfb\xbf\x38\xb4\x57\xaa\xec\x5e\x17\x8e\x00\xc3\xbb\x2b\x03\xde\xb4\x4f\xe9\x3b\xd6\xa5\x97\xad\xdf\x0d\x66\x7c\xbb\x85\x04\xe2\xdf\xe1\xd2\xbd\x42\xa6\x82\xab\x37\xc9\x20\x5b\x9f\xc8\xe2\xf5\x2b\x2f\x8c\x93\xab\xde\xb0\xaf\x68\x50\xfa\xcf\x6d\xe2\x6b\xe1\x1d\x46\xda\x62\x6f\x42\xaf\x16\xc1\x8b\x17\xc0\x01\x9c\xd8\xaf\xb6\x13\x85\xb9\x93\xb3\xde\xab\x64\x58\xde\x5d\xe0\x77\x90\x3c\xe9\xce\xc5\x3c\x1d\x16\x2c\x2f\xfe\x30\x88\x83\x05\xf8\xed\xe9\x03\x44\x3f\x6e\xed\x53\xe3\xcb\x27\xc6\xfc\x5b\x10\x65\xe6\x7e\xf2\xfd\x1b\x00\x00\xff\xff\x4b\x7e\x1f\x35\x18\x0a\x00\x00")

func TemplatesPython_client_utilsTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_client_utilsTmpl,
		"../templates/python_client_utils.tmpl",
	)
}

func TemplatesPython_client_utilsTmpl() (*asset, error) {
	bytes, err := TemplatesPython_client_utilsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_client_utils.tmpl", size: 2584, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_server_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\x51\x4b\xc3\x30\x14\x85\xdf\xf3\x2b\x2e\xa3\xd0\x0d\xba\xfd\x80\xc1\x40\x14\x05\x41\xc5\x07\xf1\x45\x24\xc4\xf5\xc6\xd5\x35\x49\xbd\x49\x27\xa3\xe6\xbf\x7b\x9b\xb0\x4d\x67\x5f\x1a\xee\x3d\xf9\xce\xe9\xe9\x30\xcc\xa1\x46\xdd\x58\x84\x09\xa1\x77\x3d\xad\x51\x76\xfb\xb0\x71\x56\x06\x34\x5d\xab\x02\x4e\x60\x1e\xa3\x18\x95\x85\xea\x9a\x07\x65\x10\x96\x2b\x58\xa4\xc3\xb8\xd1\xe4\x0c\xe8\x56\xf9\x2d\x34\xa6\x73\x14\xe0\xb2\xed\xb1\xa3\xc6\x86\x0a\x3e\xbc\xb3\x8d\xde\x57\x40\xf8\xd9\xa3\x0f\x82\x41\xf9\xea\x37\x3c\xb9\x3b\xf7\x85\x04\x31\x4a\x06\xc3\xea\x74\x6f\x5a\xfe\x53\x65\x51\x59\x81\x94\x96\x17\x52\xce\x98\x04\xa4\xec\x3b\x42\xb1\xad\xa0\xd8\xa5\x54\xf7\xc8\xd9\x6b\xcf\x4c\x21\x2e\x86\xe1\x98\xf8\x8c\xb3\x20\xd7\x07\x1c\x6d\x8a\xdd\xe2\xda\xd6\x9d\x63\xd7\x18\x19\x6f\x32\x60\xf5\x92\x77\xcf\x48\x6f\x3c\x7f\x9d\x09\xae\x09\xd2\x28\x5b\x8c\xd4\x18\xa7\x69\xf2\xa8\x48\x19\x1f\xe3\x6c\x29\x80\x9f\xb2\x2c\xd3\x7b\x18\x0e\xf1\xf4\x98\x4f\x8f\x01\x59\x7d\xd3\xdb\xf5\x95\x33\x06\x6d\xf0\xa9\xc0\xac\x65\xc1\xf1\x8c\xb6\x3e\x6e\x6e\x03\x34\x1e\x36\xca\xd6\x2d\x77\xa5\x1d\xc1\xaf\x60\x70\xf6\x01\x7f\xfc\x09\x43\x4f\xf6\xf0\x0b\xa6\xa9\xb0\x03\x59\x9c\x4c\x7e\x02\x00\x00\xff\xff\x47\xce\x2f\x6b\x04\x02\x00\x00")

func TemplatesPython_server_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_server_resourceTmpl,
		"../templates/python_server_resource.tmpl",
	)
}

func TemplatesPython_server_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesPython_server_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_server_resource.tmpl", size: 516, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\x41\x6f\xdb\x3c\x0c\x86\xcf\xd6\xaf\xe0\x27\xf4\x20\x25\xae\x53\x20\x97\xa6\x40\x0e\xc5\x57\x14\xe8\xb0\x65\x41\xd2\x9d\x86\xa1\x55\x12\x26\x11\x6a\xcb\x86\x44\x77\x2d\x8c\xfc\xf7\x51\x76\x9b\x25\xeb\xd6\xdb\x80\x5d\x12\x99\x7a\x5f\xea\x21\x6d\xaa\x69\x4e\x61\x85\x6b\xeb\x10\x64\x40\xff\x88\xfe\xae\x30\xd6\xdd\x11\x16\x55\x6e\x08\x25\x9c\xee\x76\xa2\x32\xcb\x07\xb3\x41\x68\x9a\x6c\xda\x2d\x27\xa6\x40\xde\x10\xb6\xa8\x4a\x4f\xa0\x44\x22\xf3\x72\x23\xf9\xcf\x21\x0d\xb6\x44\x55\x5c\x93\x2d\x50\x0a\x5e\x6c\x2c\x6d\xeb\x45\xb6\x2c\x8b\xc1\xa6\xf4\x36\xcf\xcd\xa0\xa8\x9f\xa4\xd0\x42\x0c\x06\x57\x7c\x0e\x78\xac\x3c\x06\x74\x04\xb3\xeb\xff\x87\xc3\xd1\x08\x56\x1c\x16\xf4\x5c\x21\xb4\x82\x98\x2b\xbb\xe5\x9f\x68\xf9\x64\x7c\xd8\x9a\xfc\xc3\xfc\xf3\x04\x4a\x86\xf6\x76\x85\x50\xfc\x0c\x8a\x75\xed\x96\xa0\x08\x7a\xd1\xab\xe1\x40\xaf\x34\xa8\xaf\xdf\x16\xcf\x84\x29\xb0\xb1\xf4\x1a\x1a\x91\x78\xa4\xda\x3b\xe8\x36\xd4\xfe\x2c\xd5\x23\x9d\x5d\x97\xbe\x30\xa4\xee\xe5\x3d\xf4\x3b\x8c\x16\x71\x38\xe2\x47\x0e\x6a\x9d\x82\xb3\xb9\xd8\x1d\x80\xdd\xe2\x13\xbd\x01\x8b\xc1\x3f\x80\xc5\xad\xbf\x0b\xf6\xc5\x15\xbf\xeb\x59\xed\xde\xe9\xda\x91\x47\x2d\x5e\x20\x74\x47\x17\xe1\x28\xb4\xa8\x70\x31\xee\x8e\x9f\xb2\x1c\x23\x4f\xff\x90\xa6\xcf\x81\x14\x02\x79\xeb\x36\x6a\xa1\xb5\x48\xec\xba\xb5\xfd\x37\x8e\x7c\x31\xd1\x6b\x99\x1c\x15\x09\xf3\x26\x3d\x82\x71\xfb\xde\x15\x05\xbd\xef\xc2\xdb\x6a\x8e\x1b\x5d\xbb\x77\x5a\x7d\xe4\xf9\x87\xaa\xf9\x05\x73\xde\x65\xd6\x2f\x47\x1c\x7c\x03\xc7\x2f\xff\x55\xb7\x4f\x11\xe7\x56\x75\xdf\x4c\x2c\x81\xe7\x2b\x9b\xe0\xf7\x59\x59\x13\x7a\x96\x89\xa4\x69\xc0\x1b\xc7\x63\x7c\xf2\x90\xc2\xc9\x63\x14\x65\x33\x0c\x65\xed\x97\x18\xae\x70\x0d\x3c\xd0\x2c\xca\xba\xd9\xbe\x71\xec\x5b\x9b\x25\xb6\x19\x82\xf2\x29\xec\xf7\x2e\xa7\x37\xcd\x4e\xb7\x19\xd1\xad\xa2\x4f\x24\x3c\xfe\xd9\x94\x91\x28\x77\x4a\x06\x32\x9e\x22\x7d\x77\xa7\x48\xd6\xc6\x3b\x21\xfb\x68\x03\xa1\xbb\x74\xab\x79\x8c\x2b\x79\x71\x7e\x76\x7e\x26\x53\xf0\x6d\x19\xf1\x32\x8a\xf9\xe2\x95\xf3\x23\x00\x00\xff\xff\xfc\xa7\xe5\x52\x99\x04\x00\x00")

func TemplatesServer_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_mainTmpl,
		"../templates/server_main.tmpl",
	)
}

func TemplatesServer_mainTmpl() (*asset, error) {
	bytes, err := TemplatesServer_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_main.tmpl", size: 1177, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_python_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xd0\xcd\x6a\xeb\x30\x10\x05\xe0\xbd\x9f\xe2\x60\xb2\x70\xe0\xc6\x0f\x10\xf0\xee\xd2\x55\xe9\xa2\x64\x3f\x28\xf5\xc8\x15\xb1\x7e\x18\x49\x29\x45\xf5\xbb\x57\x8e\x1b\xba\x68\xa1\xda\x09\xe9\x9c\x8f\x99\x52\x0e\x18\x59\x1b\xc7\x68\x23\xcb\x95\x85\xc2\x7b\x7a\xf5\x8e\xac\x32\x8e\x12\xdb\x30\xab\xc4\x2d\x0e\xcb\xd2\x68\xf1\x16\x7a\x56\xf1\x02\x63\x83\x97\x84\x87\xf5\xd2\x34\xa5\x40\x94\x9b\x18\xbb\xcb\x3f\xec\xae\x38\x0e\xe8\x9f\x39\xfa\x2c\x2f\x1c\xff\xb3\xfe\x4e\x97\xd2\x3f\x29\xcb\xf8\xc0\xc9\x3f\xfa\x37\x96\x65\xb9\x77\xfd\xf2\x44\x2a\x98\x5a\xce\x6e\x44\x2d\x68\x54\x08\x18\x36\xb3\x23\x72\xf5\x33\xd1\x7e\xd5\xff\xc6\x6b\xb4\x17\x9e\x4c\x4c\x75\xc2\xf3\x9c\x39\x88\x71\xa9\xfb\x61\x62\x43\xf7\x9b\xba\xa2\x46\xe3\x6e\x61\x18\xd0\xd2\xb6\x19\x6a\x8f\x0d\xea\xb9\x15\x67\xd7\x8d\x7c\xce\xd3\x70\x92\xcc\x5f\xd9\x1b\xfb\x19\x00\x00\xff\xff\xba\x73\xf0\x48\x5f\x01\x00\x00")

func TemplatesServer_python_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_python_mainTmpl,
		"../templates/server_python_main.tmpl",
	)
}

func TemplatesServer_python_mainTmpl() (*asset, error) {
	bytes, err := TemplatesServer_python_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_python_main.tmpl", size: 351, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_resources_apiTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x53\x4f\x6f\xdb\x3e\x0c\x3d\xdb\x9f\x82\x3f\x23\xf8\xc1\x1e\x1a\xa5\x87\x9d\x56\xf4\xb0\x3f\x2d\xd6\x01\xcb\xb2\x15\xe8\x8e\x85\x1a\xd3\x8d\x17\x5b\x72\x68\x39\x46\x20\xf8\xbb\x8f\x94\xdc\xb5\xc9\x89\x22\x1f\x1f\x1f\xf9\x1c\xef\x97\x50\x62\x55\x1b\x84\x8c\xb0\xb7\x03\x6d\xf1\x51\x77\xf5\xa3\xc3\xb6\x6b\xb4\xc3\x0c\x96\xd3\x94\x76\x7a\xbb\xd7\xcf\x08\xde\xab\x4d\x0c\xd7\xba\x45\x2e\xa4\x75\xdb\x59\x72\x90\xa7\x89\xf7\x50\x57\xa0\xd6\x88\xe5\xb7\xfb\x1f\x6b\x98\xa6\x0c\xcd\xd6\x96\xb5\x79\x5e\xfd\xe9\xad\xc9\x18\x80\xa6\xe4\x7c\x9a\x64\x06\xdd\x6a\xe7\x5c\x97\xa5\x45\xea\xfd\x82\x27\x0a\x21\x7c\xb8\x66\x82\xc8\xec\x4e\x5d\x98\x17\x9f\x1f\x37\x77\xd0\x3b\x1a\xb6\x0e\x7c\xca\x63\xb9\x09\x48\x1b\x96\xb4\xd8\x5f\xc0\xe2\x18\x3a\xbf\xa3\xdb\xd9\xb2\x97\x09\x6f\xca\x95\xd4\x2b\x01\x2c\x8e\xea\x76\x30\xdb\xcf\xb6\x6d\xd1\xb8\x9e\x61\xab\x15\x8f\xe0\xea\x34\x79\xcf\xda\x62\xe6\xce\x41\xdd\xc3\x4e\x9b\xb2\x41\x82\xca\x52\xc0\xa8\x07\xa4\xa7\x69\x8a\xf1\x8d\x29\x3b\x5b\x1b\xc7\x0d\x15\x33\xe6\xac\x1f\x5e\xd7\x08\x72\x8b\x88\x8c\x9a\x62\x36\x1f\x41\x76\x56\xbf\xb0\xef\xac\xe9\xf1\x37\xd5\x0e\xe9\x02\x08\xde\xcd\xf9\xc3\x80\xbd\xe3\x4e\xb9\xe6\xf2\xdf\x06\x07\xd9\xe0\x30\x6f\xf0\x73\x40\x3a\x6d\x34\x31\x23\x37\xf7\xc1\x1e\xe0\x5f\x5c\x65\x7f\x60\x89\x0c\x24\x3c\xa8\x5b\x4b\xed\x83\x6e\x06\xcc\xb3\xb9\x92\x15\x91\xf8\xc5\x06\x89\xc5\x33\x1e\xfc\xc9\x96\xa7\xc0\x95\x1c\x35\x49\x7b\x48\xf0\xfd\xe7\x9a\x78\x9d\x30\x16\x89\x84\x5f\x0c\x65\xab\xc7\x2f\xc8\x0e\x23\xe5\xa4\x04\x54\xa8\xf8\xce\xff\x9f\x09\x8a\xab\xd0\xf0\xdf\x35\x98\xba\x91\xb5\x92\x51\x85\xad\xbf\xa2\x96\xb6\xf7\x97\x97\x2c\x29\x21\x74\x03\x99\x34\x99\xce\xe4\xbd\xd5\xd7\x77\x41\xcf\xab\xbe\x39\x11\x04\xc6\x58\x6a\x2f\xb2\x6e\x4c\x94\x35\x16\x2a\x86\xa2\x28\xc2\x8a\xab\xf3\x1b\xf0\xdd\xd8\xc1\xf8\x49\xc0\x13\x36\x76\x84\x46\xfe\x0e\xce\x82\x2e\x4b\xd8\x05\xa1\x01\x36\xaa\x59\x75\xa1\xee\xd1\xe5\xd9\x1e\x4f\xd9\x45\x76\x94\x0b\xf3\x61\xe5\x9b\x8b\xac\x72\xc5\xb3\xc7\xdf\x00\x00\x00\xff\xff\x35\xbb\x9e\xfc\x65\x03\x00\x00")

func TemplatesServer_resources_apiTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_resources_apiTmpl,
		"../templates/server_resources_api.tmpl",
	)
}

func TemplatesServer_resources_apiTmpl() (*asset, error) {
	bytes, err := TemplatesServer_resources_apiTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_resources_api.tmpl", size: 869, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_resources_interfaceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xcf\x6e\x9c\x30\x10\xc6\xcf\xf8\x29\x46\x68\x0f\x10\xed\x9a\x7b\xa5\x9e\xfa\x47\xcd\x21\x51\x15\x45\xcd\x31\xf2\xc2\x00\x6e\xc0\xa6\xf6\xb0\x9b\x08\xf1\xee\x1d\x63\x76\x9b\xb4\x2b\x35\x5c\x18\xcc\x37\x33\xdf\xfc\xc6\xd3\xb4\x83\x0a\x6b\x6d\x10\x52\x87\xde\x8e\xae\xc4\x47\x5d\x3f\x12\xf6\x43\xa7\x08\x53\xd8\xcd\xb3\x18\x54\xf9\xa4\x1a\x84\x69\x92\xdf\x63\x78\xab\x7a\xe4\x1f\xa2\x28\xee\x5b\xed\xa1\xd6\x1d\x02\xbf\xd5\x48\x76\xd7\xa0\x41\xc7\xb9\x15\xec\x5f\xa0\xb1\x3b\xa7\xfa\x8e\x85\x9f\x2d\x18\x4b\x80\x95\x26\xa0\x73\x12\x4b\x5a\x65\x2a\xf0\xda\x94\x5c\x82\xe0\xa8\xbb\x0e\xf6\x08\xf6\x80\xee\xe8\x34\x11\x1a\xa8\x46\xa7\x4d\xc3\x59\x08\x06\x9f\x09\xd6\x0e\xda\x1a\x21\x74\x3f\x58\x47\x90\x89\x24\x6d\x34\xb5\xe3\x5e\x96\xb6\x2f\x1a\xeb\xb8\x8e\x2a\xfa\xf1\x39\xe5\x3f\x06\xa9\x68\x89\x86\x54\x00\x3f\xd3\x04\xba\x06\xf9\xc0\xf2\x1b\x5d\x55\x1d\x1e\x95\x43\x98\x67\x78\x5d\xe1\xe7\xe8\x49\x1b\xe5\x0b\xd5\xe9\x92\x39\x70\x12\x9a\x8a\x67\xce\x85\xa0\x97\x61\x81\x11\x29\x5c\x1b\x42\x57\xab\x60\xff\x14\x4d\x22\x09\x64\x9d\x32\x4c\x6d\xf3\xb4\x85\xcd\x01\x3e\x7c\x04\x79\x83\xd4\xda\xca\x2f\x50\x13\x00\x16\x9d\x35\x75\x10\xd5\x41\xb5\x39\xc8\xaf\xa3\x29\x3f\xd9\xbe\x47\x43\x7e\x95\x16\x05\x77\x64\xc5\x3c\x4f\x53\x34\x92\x84\x59\xf8\xf8\x9a\x02\xfa\x40\xb1\x43\x07\xb5\x75\x8b\x50\xfe\x40\xb7\xe7\x99\x96\xf8\x8b\xa9\x06\xcb\xee\x42\x56\xb2\x9c\x44\x27\x71\x80\x2c\xa0\x91\x77\xe8\x07\x6b\x3c\x3e\x30\x74\x74\x5b\xb8\x5a\x4f\x7f\x8d\xe8\x29\x0f\x03\x05\x00\x8c\x49\xf0\xde\x6b\x36\x78\x81\xc0\x9d\x1d\x09\x7d\xe6\xe0\x8a\xc1\xcb\xe5\x8b\x2b\xe9\x0b\xca\x9c\xd7\xf0\x1f\x48\xdc\x09\xd6\x27\xe8\x78\x65\xc1\xf7\x79\x61\x6f\x04\x4e\x7e\x5b\xe6\xcf\xd2\xbf\xe6\x4d\xb7\xb0\x2c\x50\xde\xe2\x31\x8b\x93\xff\xa9\x30\xcf\xb9\xbc\x6f\xd1\x44\x00\xb1\x82\x0b\xec\x33\x2d\xff\xa1\x94\xe7\xf9\xc9\xda\xda\x25\x12\x4e\xf3\x37\x36\xb1\xf3\xe1\x2e\x31\xe7\x93\xa7\xa5\xe0\x05\x5f\x97\x7a\xbc\xab\x43\xdc\x42\xf2\x2a\x9e\xc5\xe9\x23\xdc\xac\xdf\x01\x00\x00\xff\xff\x7e\x77\x8b\x5e\xd5\x03\x00\x00")

func TemplatesServer_resources_interfaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_resources_interfaceTmpl,
		"../templates/server_resources_interface.tmpl",
	)
}

func TemplatesServer_resources_interfaceTmpl() (*asset, error) {
	bytes, err := TemplatesServer_resources_interfaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_resources_interface.tmpl", size: 981, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesStructTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8f\x41\x4b\xc4\x30\x14\x84\xef\xfd\x15\x8f\xb2\x47\x37\x7b\x5f\xf0\xe4\x22\x08\xa2\x1e\xbc\xbb\xa1\x9d\x4a\xdc\x36\x8d\x4d\x76\xa1\x84\xfc\x77\x27\x49\x51\x73\xc9\xcb\x4c\x86\xf9\x5e\x8c\x3d\x06\x63\x21\xad\x0f\xcb\xb5\x0b\x1f\x01\x93\x1b\x75\x40\x9b\x52\xe3\x74\x77\xd1\x9f\x90\x18\xd5\x5b\x1d\x5f\xf4\x04\x1a\x4d\x8c\xb2\x68\x4b\x6b\x77\x93\xe3\xbd\xa8\x13\x7c\xb7\x18\x17\xcc\x6c\x85\xfe\xe1\xc0\xcc\xee\x96\x12\x2f\xd8\x9e\x0a\x03\x66\x10\xf5\xe4\x5f\x2d\x9e\xd9\x77\xc2\x20\xfb\xaa\xab\x7f\x52\x51\xf6\x82\xd1\xa3\xd8\x61\x75\xb9\x5e\x54\x2e\xa6\x2b\x95\x52\x62\x23\x3c\x7f\x14\x17\xac\x77\x64\xd1\xe3\x15\x85\xe7\xd1\x60\xec\x3d\x03\x19\x23\xab\xaa\x92\xe7\x0c\x39\xf0\xbd\x7d\x26\xd0\xc3\x3c\xb9\xd9\x9b\x82\x3e\x68\x16\x17\xea\xcd\x7e\x67\x3f\xdf\xe7\x2f\x3f\xdb\x63\x4b\x99\x45\x29\xb5\xe7\xdf\xbd\x2a\x46\x9d\x37\x76\xdb\x6f\x9b\x55\xf5\x27\x00\x00\xff\xff\x5a\x33\x7f\x59\x62\x01\x00\x00")

func TemplatesStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesStructTmpl,
		"../templates/struct.tmpl",
	)
}

func TemplatesStructTmpl() (*asset, error) {
	bytes, err := TemplatesStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/struct.tmpl", size: 354, mode: os.FileMode(420), modTime: time.Unix(1458530859, 0)}
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
	"../templates/client_helper_resource.tmpl": TemplatesClient_helper_resourceTmpl,
	"../templates/client_resource.tmpl": TemplatesClient_resourceTmpl,
	"../templates/generic_main.tmpl": TemplatesGeneric_mainTmpl,
	"../templates/oauth2_middleware.tmpl": TemplatesOauth2_middlewareTmpl,
	"../templates/oauth2_scopes_match.tmpl": TemplatesOauth2_scopes_matchTmpl,
	"../templates/python_client.tmpl": TemplatesPython_clientTmpl,
	"../templates/python_client_utils.tmpl": TemplatesPython_client_utilsTmpl,
	"../templates/python_server_resource.tmpl": TemplatesPython_server_resourceTmpl,
	"../templates/server_main.tmpl": TemplatesServer_mainTmpl,
	"../templates/server_python_main.tmpl": TemplatesServer_python_mainTmpl,
	"../templates/server_resources_api.tmpl": TemplatesServer_resources_apiTmpl,
	"../templates/server_resources_interface.tmpl": TemplatesServer_resources_interfaceTmpl,
	"../templates/struct.tmpl": TemplatesStructTmpl,
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
	"..": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"client_helper_resource.tmpl": &bintree{TemplatesClient_helper_resourceTmpl, map[string]*bintree{}},
			"client_resource.tmpl": &bintree{TemplatesClient_resourceTmpl, map[string]*bintree{}},
			"generic_main.tmpl": &bintree{TemplatesGeneric_mainTmpl, map[string]*bintree{}},
			"oauth2_middleware.tmpl": &bintree{TemplatesOauth2_middlewareTmpl, map[string]*bintree{}},
			"oauth2_scopes_match.tmpl": &bintree{TemplatesOauth2_scopes_matchTmpl, map[string]*bintree{}},
			"python_client.tmpl": &bintree{TemplatesPython_clientTmpl, map[string]*bintree{}},
			"python_client_utils.tmpl": &bintree{TemplatesPython_client_utilsTmpl, map[string]*bintree{}},
			"python_server_resource.tmpl": &bintree{TemplatesPython_server_resourceTmpl, map[string]*bintree{}},
			"server_main.tmpl": &bintree{TemplatesServer_mainTmpl, map[string]*bintree{}},
			"server_python_main.tmpl": &bintree{TemplatesServer_python_mainTmpl, map[string]*bintree{}},
			"server_resources_api.tmpl": &bintree{TemplatesServer_resources_apiTmpl, map[string]*bintree{}},
			"server_resources_interface.tmpl": &bintree{TemplatesServer_resources_interfaceTmpl, map[string]*bintree{}},
			"struct.tmpl": &bintree{TemplatesStructTmpl, map[string]*bintree{}},
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

