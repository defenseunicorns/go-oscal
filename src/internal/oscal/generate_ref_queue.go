package oscal

type RefQueue struct {
	refs   []string
	refMap map[string]bool
}

func NewRefQueue() *RefQueue {
	return &RefQueue{
		refs:   []string{},
		refMap: map[string]bool{},
	}
}

func (r *RefQueue) Add(ref string) {
	if _, ok := r.refMap[ref]; !ok {
		r.refMap[ref] = true
		r.refs = append(r.refs, ref)
	}
}

func (r *RefQueue) Pop() string {
	if len(r.refs) > 0 {
		ref := r.refs[0]
		if len(r.refs) != 1 {
			r.refs = r.refs[1:]
		} else {
			r.refs = []string{}
		}
		return ref
	}
	return ""
}

func (r *RefQueue) Len() int {
	return len(r.refs)
}
