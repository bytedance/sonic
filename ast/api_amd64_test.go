// +build amd64,go1.15,!go1.21

package ast

import (
    `testing`

    `github.com/bytedance/sonic/encoder`
    `github.com/stretchr/testify/assert`
)

func TestSortNodeTwitter(t *testing.T) {
    root, err := NewSearcher(_TwitterJson).GetByPath()
    if err != nil {
        t.Fatal(err)
    }
    obj, err := root.MapUseNumber()
    if err != nil {
        t.Fatal(err)
    }
    exp, err := encoder.Encode(obj, encoder.SortMapKeys)
    if err != nil {
        t.Fatal(err)
    }
    if err := root.SortKeys(true); err != nil {
        t.Fatal(err)
    }
    act, err := root.MarshalJSON()
    if err != nil {
        t.Fatal(err)
    }
    assert.Equal(t, len(exp), len(act))
    assert.Equal(t, string(exp), string(act))
}