package ast

import (
	"sync"
)

var bytesPool = sync.Pool{}

func newBuffer() *[]byte {
    if ret := bytesPool.Get(); ret != nil {
        return ret.(*[]byte)
    } else {
        buf := make([]byte, 0, 1024)
        return &buf
    }
}

func freeBuffer(buf *[]byte) {
    *buf = (*buf)[:0]
    bytesPool.Put(buf)
}


func (self *Node) MarshalJSON() ([]byte, error) {
    // if err := self.Check(); err != nil {
    //     return nil, err
    // }
    // buf := newBuffer()
    // err := self.encode(buf)
    // if err != nil {
    //     freeBuffer(buf)
    //     return nil, err
    // }
    // ret := make([]byte, len(*buf))
    // copy(ret, *buf)
    // freeBuffer(buf)
    return []byte(self.node.JSON), nil
}

// func (self *Node) encode(buf *[]byte) (error) {
//     if self.node.Kind == types.Type(V_ANY) {
//         // to pass encoder option
//         err := encoder.EncodeInto(buf, self.any(), 0)
//         if err != nil {
//             return  err
//         }
//         return nil
//     } else 
//     if !self.isMut() { 
//         *buf = append(*buf, self.node.JSON...)
//         return nil
//     } else if self.node.Kind == types.T_ARRAY {
//         *buf = append(*buf, '[')
//         for i, v := range self.node.Kids {
//             if i > 0 {
//                 *buf = append(*buf, ',')
//             }
//             n := self.getKidRaw(v)
//             if err := n.encode(buf); err != nil {
//                 return err
//             }
//         }
//         *buf = append(*buf, ']')
//         return nil
//     } else if self.node.Kind == types.T_OBJECT {
//         *buf = append(*buf, '{')
//         for i:=0; i<len(self.node.Kids); i+=2 {
//             if i > 0 {
//                 *buf = append(*buf, ',')
//             }
//             key := self.key(self.node.Kids[i])
//             *buf = append(*buf, key...)
//             *buf = append(*buf, ':')
//             val := self.getKidRaw(self.node.Kids[i+1])
//             if err := val.encode(buf); err != nil {
//                 return err
//             }
//         }
//         *buf = append(*buf, '}')
//         return nil
//     } else {
//         panic("unreachable")
//     }
// }