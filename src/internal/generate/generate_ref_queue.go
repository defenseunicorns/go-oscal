package generate

type RefQueue struct {
	refs       []string
	refMap     map[string]bool
	refHistory []string
}

func NewRefQueue() RefQueue {
	return RefQueue{
		refs:       []string{},
		refMap:     map[string]bool{},
		refHistory: []string{},
	}
}

func (r *RefQueue) Add(ref string) {
	if has := r.refMap[ref]; !has {
		r.refMap[ref] = true
		r.refs = append(r.refs, ref)
		r.refHistory = append(r.refHistory, ref)
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

func (r *RefQueue) History() []string {
	return r.refHistory
}

func (r *RefQueue) Len() int {
	return len(r.refs)
}
