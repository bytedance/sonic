// +build !noasm,amd64 !appengine,amd64

package rt

func StopProf()

// func StopProf() {
//     atomic.AddUint32(&yieldCount, 1)
//     if runtimeProf.hz != 0 {
//         oldHz = runtimeProf.hz
//         runtimeProf.hz = 0
//     }
// }

func StartProf()

// func StartProf() {
//     atomic.AddUint32(&yieldCount, ^uint32(0))
//     if yieldCount == 0 && runtimeProf.hz == 0 {
//         if oldHz == 0 {
//             oldHz = 100
//         }
//         runtimeProf.hz = oldHz
//     }
// }
