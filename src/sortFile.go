package src

import "os"

type ByLen []os.FileInfo

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i].Name()) < len(a[j].Name())
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
