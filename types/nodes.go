package types

type (
	nodes []*Tree
)

func (ns nodes) Len() int {
	return len(ns)
}

func (ns nodes) Less(i, j int) bool {
	in := int(ns[i].Name[0])
	if ns[i].IsFile {
		in += 1000
	}
	jn := int(ns[j].Name[0])
	if ns[j].IsFile {
		jn += 1000
	}
	return in < jn
}

func (ns nodes) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}
