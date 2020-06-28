package app

import (
	"archive/tar"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"upm/types"
	"upm/utils"
)

type (
	asset struct {
		Id     string
		Path   string
		IsFile bool
		Image  string
	}
)

func OpenPackage(filename string) *types.PackageInfo {
	defer func() {
		if err := recover(); err != nil {
			utils.Toast(fmt.Sprint(err), "error")
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	tr := tar.NewReader(g)

	info := &types.PackageInfo{
		Name: filepath.Base(filename),
		Path: filename,
		Tree: &types.Tree{
			Name:   "Assets",
			Path:   "Assets",
			IsFile: false,
		},
	}

	type chanInfo struct {
		Id    string
		Key   string
		Value interface{}
	}
	assetChan := make(chan chanInfo)
	end := make(chan struct{})
	go func() {
		for {
			th, err := tr.Next()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}

			id := strings.Split(th.Name, "/")[0]

			if strings.HasSuffix(th.Name, "/asset") {
				assetChan <- chanInfo{id, "isFile", true}
			}

			if strings.HasSuffix(th.Name, "/pathname") {
				data, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				assetChan <- chanInfo{id, "path", string(data)}
			}

			if strings.HasSuffix(th.Name, "/preview.png") {
				data, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				img := "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
				assetChan <- chanInfo{id, "image", img}
			}
		}
		close(end)
	}()

	func() {
		assets := make(map[string]*asset)
	F:
		for {
			select {
			case info := <-assetChan:
				_, ok := assets[info.Id]
				if !ok {
					assets[info.Id] = &asset{}
				}
				switch info.Key {
				case "isFile":
					assets[info.Id].IsFile = info.Value.(bool)
				case "path":
					assets[info.Id].Path = info.Value.(string)
				case "image":
					assets[info.Id].Image = info.Value.(string)
				}
			case <-end:
				break F
			}
		}
		for _, v := range assets {
			mt(info.Tree, *v)
		}
	}()
	return info
}

func mt(t *types.Tree, info asset) {
	ps := strings.Split(info.Path, "/")
	now := t
	for i := 1; i < len(ps); i++ {
		var nt *types.Tree
		if fn := now.NodeFind(ps[i]); fn == nil {
			nt = &types.Tree{
				Name:   ps[i],
				Path:   now.Path + "/" + ps[i],
				IsFile: false,
			}
			if i == len(ps)-1 && info.IsFile {
				nt.IsFile = true
				nt.AssetID = info.Id
				nt.Image = info.Image
				nt.Ext = filepath.Ext(ps[i])
				now.Nodes = append(now.Nodes, nt)
			} else {
				now.Nodes = append([]*types.Tree{nt}, now.Nodes...)
			}
			now.SortNodes()
			now = nt
		} else {
			now = fn
		}
	}
}
