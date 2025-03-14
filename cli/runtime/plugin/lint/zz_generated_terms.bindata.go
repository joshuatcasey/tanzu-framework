// Code generated by go-bindata. DO NOT EDIT.
// sources:
// hack/linter/cli-wordlist.yml

package lint

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataHackLinterCliwordlistYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x54\x51\xb2\x83\x36\x0c\xfc\xcf\x29\xb8\x80\xbe\x3b\xd3\x6b\xb4\x17\x30\x46\x18\x4d\x8c\xe4\x91\x64\xf2\xd2\xd3\x77\xb0\x1d\x92\xbc\x4e\xbf\x58\x2d\x92\xb2\x5e\x2f\x01\x80\xdb\xc4\x52\xd9\xfe\xbc\x4d\x13\x4c\x21\x46\xa9\xec\x1d\x2f\x3b\x31\xdc\xeb\x8c\x51\x78\xa5\xd4\xc9\x52\xfa\xf3\x61\xed\x19\x91\x0a\x94\xa0\x4e\x91\x4a\x70\x12\xee\x74\xae\xe6\xa8\x9f\x38\xa9\xd4\x3e\x1b\x65\x2f\xc2\x38\x7e\x26\x2a\x2e\xc8\x4e\x21\xf7\x8d\x4b\xf0\x00\x45\xc5\x31\x5e\xeb\x88\x8d\xd2\xd6\xfb\x7f\x09\x3a\x4b\x65\x74\x34\x50\xcc\x18\x0c\x1b\x9d\x25\x51\x9f\xdd\x43\xdc\x88\x71\xc3\x90\x7d\x8b\x1b\xc6\xfb\xa0\x39\x24\xdc\x91\x1d\x3e\xc5\x16\xd4\x9d\xcc\x48\xb8\xab\x29\xb9\xbe\x16\x15\xc9\x14\x9f\x1d\xaa\x1c\x74\x36\x8d\x29\xc5\x22\x0d\x3c\x44\xef\x56\x42\xc4\xdb\x74\xa0\xce\x2f\x57\x8f\x40\x39\xcc\x19\xa1\x96\xa4\x61\xc1\xe1\x5d\xc6\xc0\x97\x23\x19\xdf\xee\x29\x06\xef\xe7\x58\x30\xe3\x05\x2d\x2a\xcd\xbd\x48\xd8\xdd\xa0\xbd\x88\x0e\xc8\xab\x0c\x40\x2f\xc6\x3c\xe4\xdc\x0d\x21\xeb\xa4\xd8\x90\x9c\xe8\x3a\xb5\xc5\x90\xfb\x5e\xdb\xe4\xd1\x40\x2d\xcb\x4b\xc3\x10\xdd\xf0\x81\x6a\x4d\x66\xca\x32\x87\x0c\x6b\x0e\xa9\x9d\x72\x82\x69\xc3\x5c\x3a\xca\x92\x60\xa5\xb6\xf2\x9c\x19\x0f\xd4\x59\xce\xeb\x89\xb2\xef\x81\x97\xaf\xd9\x50\xe8\x6f\xb9\x23\xf7\x6a\x26\x5e\x06\x52\x79\x58\x53\xf9\x7f\x61\x3b\x79\xa6\x01\x84\x1d\x7f\xfc\x5d\xa8\xe4\x92\x03\x23\x18\xfd\x33\xd4\x2c\x58\xb2\x3c\xc1\xef\x09\x84\xe1\xf8\xab\x6c\xa8\xf8\xc7\x78\x47\xd6\xae\xa9\x65\x95\x38\x7d\xb3\x2c\x80\x71\x93\x4e\x22\x37\xce\xef\xc9\xfe\xbb\x07\x79\x29\x42\x3c\x84\xac\x18\xbc\x2a\xf6\xf3\x0e\xea\x32\x67\x15\x8d\x9f\x10\xa2\x0d\x13\x53\x2c\x30\xd7\x78\x47\x07\x0e\x3b\xbe\x49\x15\x71\x28\xc1\xb7\x4e\x11\xc7\x5c\x17\x84\x96\xe8\x5f\x81\xee\xb1\xd0\x60\xae\x35\x9e\x22\x3a\xf7\xf5\x0d\x4d\x30\xbd\xf7\x9f\x68\x04\xb8\x85\x65\x5d\x0d\xc7\x31\xc4\x20\x68\xdc\xae\xe2\x3d\x24\x06\x57\x2e\x5a\x5d\xbd\xd4\x31\x75\xda\xdf\x91\xa1\x1e\x2f\x51\x67\xcc\x20\xe4\x0c\x51\x78\x21\x1f\x9f\xdb\xf9\xe2\xba\x27\xf3\x90\xae\x2b\x30\x5f\x50\x15\x84\xf3\xb3\x13\x1e\x34\x0d\x63\x3e\xe4\x3a\xed\x28\xaf\x5f\xf6\x3d\x42\x4f\xb9\xb6\xac\x40\xd5\xdc\xdf\xd4\x11\x97\x6a\x08\xf8\x43\xe6\xc4\x09\x66\x11\x3f\x5b\xcb\xb7\x7d\x5f\x3d\xed\x8b\xad\xbf\x3a\x0e\x6b\x57\x0f\x5f\x79\xfb\x4e\xc0\xab\xe5\xd8\xc1\x71\x2f\x39\x38\x7e\xd8\x77\xfe\x69\xa0\x7e\x44\xf4\x89\x76\xfb\x37\x00\x00\xff\xff\x89\xa5\x48\x5c\x9a\x05\x00\x00")

func bindataHackLinterCliwordlistYmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataHackLinterCliwordlistYml,
		"hack/linter/cli-wordlist.yml",
	)
}

func bindataHackLinterCliwordlistYml() (*asset, error) {
	bytes, err := bindataHackLinterCliwordlistYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "hack/linter/cli-wordlist.yml",
		size:        1434,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

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
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// AssetNames returns the names of the assets.
// nolint: deadcode
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"hack/linter/cli-wordlist.yml": bindataHackLinterCliwordlistYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"hack": {Func: nil, Children: map[string]*bintree{
		"linter": {Func: nil, Children: map[string]*bintree{
			"cli-wordlist.yml": {Func: bindataHackLinterCliwordlistYml, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
