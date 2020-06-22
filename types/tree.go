package types

import (
	"sort"
)

type (
	Tree struct {
		Name    string  `json:"name"`
		Path    string  `json:"id"`
		Ext     string  `json:"ext,omitempty"`
		AssetID string  `json:"assetId,omitempty"`
		IsFile  bool    `json:"isFile,omitempty"`
		Extract bool    `json:"extract,omitempty"`
		Image   string  `json:"image,omitempty"`
		Nodes   []*Tree `json:"children,omitempty"`
	}
)

func (t *Tree) SortNodes() {
	ns := nodes(t.Nodes)
	sort.Sort(ns)
	t.Nodes = ns
}

func (t *Tree) NodeFind(name string) *Tree {
	for _, v := range t.Nodes {
		if v.Name == name {
			return v
		}
	}
	return nil
}
