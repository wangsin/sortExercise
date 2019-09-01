package basicsort

type Element struct {
	Slice []interface{}
	By    func(a, b interface{}) bool
}

func (ele Element) Len() int           { return len(ele.Slice) }
func (ele Element) Swap(i, j int)      { ele.Slice[i], ele.Slice[j] = ele.Slice[j], ele.Slice[i] }
func (ele Element) Less(i, j int) bool { return ele.By(ele.Slice[i], ele.Slice[j]) }
