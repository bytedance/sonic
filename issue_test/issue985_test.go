package issue_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
)

func TestIssue985ArrayClear(t *testing.T) {
	for _, input := range []string{`[]`, `[null]`, `[1]`, `[1,2]`, `[1,2,3]`, `null`} {
		t.Run(input, func(t *testing.T) {
			for _, makeArray := range []func() interface{}{
				func() interface{} { return &[2]interface{}{"old", "tail"} },
				func() interface{} { return &[2]*int{new(int), new(int)} },
				func() interface{} { return &[2]int{7, 8} },
				func() interface{} { return &[0]interface{}{} },
			} {
				got, want := makeArray(), makeArray()
				if err := json.Unmarshal([]byte(input), want); err != nil {
					t.Fatal(err)
				}
				if err := sonic.Unmarshal([]byte(input), got); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%T: got %v, want %v", got, got, want)
				}
			}
		})
	}
}

type issue985Payload struct {
	Status string `json:"status"`
	Data   struct {
		Result []struct {
			Metric map[string]string `json:"metric"`
			Value  [2]interface{}    `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func TestIssue985ConcurrentFirstDecode(t *testing.T) {
	const rows = 3400
	var b bytes.Buffer
	b.WriteString(`{"status":"success","data":{"resultType":"vector","result":[`)
	for i := 0; i < rows; i++ {
		if i != 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, `{"metric":{"__name__":"up","instance":"172.%d.%d.%d:1888","job":"agent_env_all","env":"pro"},"value":[1788943336,"1"]}`, 16+i%40, i/250, i%250+1)
	}
	b.WriteString(`]}}`)

	// Distinct types keep each round's top-level decoder cold.
	for round := 0; round < 10; round++ {
		typ := reflect.StructOf([]reflect.StructField{{
			Name: "Payload", Type: reflect.TypeOf(issue985Payload{}),
			Anonymous: true, Tag: reflect.StructTag(fmt.Sprintf(`issue985:"%d"`, round)),
		}})
		start := make(chan struct{})
		var wg sync.WaitGroup
		for worker := 0; worker < 8; worker++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dst := reflect.New(typ)
				<-start
				if err := sonic.Unmarshal(b.Bytes(), dst.Interface()); err != nil {
					t.Error(err)
					return
				}
				got := dst.Elem().Field(0).Interface().(issue985Payload)
				if got.Status != "success" || len(got.Data.Result) != rows {
					t.Errorf("unexpected payload: status=%q rows=%d", got.Status, len(got.Data.Result))
					return
				}
				for _, row := range got.Data.Result {
					if row.Metric["__name__"] != "up" || row.Value != [2]interface{}{float64(1788943336), "1"} {
						t.Errorf("unexpected row: %+v", row)
						return
					}
				}
			}()
		}
		close(start)
		wg.Wait()
	}
}
