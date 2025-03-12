package issue_test

import (
	"log"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
)

type MyJs struct {
	Arr [][2]string
}

func TestIssue762(t *testing.T) {
	type parsedStruct struct {
		parsed MyJs
		err    error
	}
	conv := func(buf []byte) parsedStruct {
		var x MyJs
		err := sonic.Unmarshal(buf, &x)
		return parsedStruct{x, err}
	}

	const msg = `{"Arr":[["foo","bar"]]}`
	log.Printf("%#v\n", conv([]byte(msg)))

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j :=0; j < 5000; j++ {
				conv([]byte(msg))
			}
		}()
	}
	wg.Wait()
	log.Println("wg done")
}