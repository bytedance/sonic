package utf8

import (
	"github.com/bytedance/sonic/internal/rt"
	"github.com/bytedance/sonic/internal/native/types"
	"github.com/bytedance/sonic/internal/native"
    "os"
    "fmt"
)

var (
    allCorrect = os.Getenv("SONIC_UTF8_ALL_RIGHT")  != ""
)

type InvalidUTF8Error struct {
    Pos  int
    Src  string
}

func (self InvalidUTF8Error) Error() string {
    return fmt.Sprintf("invalid UTF-8 postion is %d in string: %s", self.Pos, self.Src)
} 

// CorrectWith corrects the invalid utf8 byte with repl string.
func CorrectWith(dst []byte, src []byte, repl string) ([]byte, error) {
    if allCorrect {
        panic("should be right utf8, not panic here")
    }
    var err error
    sstr := rt.Mem2Str(src)
    sidx := 0

    /* state machine records the invalid postions */
    m := types.NewStateMachine()
    m.Sp = 0 // invalid utf8 numbers

    for sidx < len(sstr) {
        scur  := sidx
        ecode := native.ValidateUTF8(&sstr, &sidx, m)

        if m.Sp != 0 {
            if m.Sp > len(sstr) {
                panic("numbers of invalid utf8 exceed the string len!")
            }
            /* error record the first invalid utf8 position */
            err = &InvalidUTF8Error {
                Pos: m.Vt[0],
                Src: string(sstr),
            }
        }
        
        for i := 0; i < m.Sp; i++ {
            ipos := m.Vt[i] // invalid utf8 position
            dst  = append(dst, sstr[scur:ipos]...)
            dst  = append(dst, repl...)
            scur = m.Vt[i] + 1
        }
        /* append the remained valid utf8 bytes */
        dst = append(dst, sstr[scur:sidx]...)

        /* not enough space, reset and continue */
        if ecode != 0 {
            m.Sp = 0
        }
    }
    types.FreeStateMachine(m)
    return dst, err
}

// Validate is a simd-accelereated drop-in replacement for the standard library's utf8.Valid.
func Validate(src []byte) bool {
    return ValidateString(rt.Mem2Str(src))
}

func ValidateString(src string) bool {
    return native.ValidateUTF8Fast(&src) == 0
}