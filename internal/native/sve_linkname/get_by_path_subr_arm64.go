//go:build !noasm || !appengine
// +build !noasm !appengine

// DO NOT EDIT.

package sve_linkname

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __get_by_path_entry__() uintptr

var (
	_subr__get_by_path uintptr = __get_by_path_entry__() +  48
)

const (
	_stack__get_by_path = 224
)

var (
	_ = _subr__get_by_path
)

const (
	_ = _stack__get_by_path
)
