package issue_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	info := &ShowInfo{
		StartTime: JSONTime(time.Now()),
	}
	infos := &[]*ShowInfo{info}
	bytes, err1 := json.Marshal(infos)
	fmt.Printf("%+v\n", string(bytes))
	spew.Dump(bytes, err1)

	jbytes, err2 := sonic.Marshal(infos)
	fmt.Printf("%+v\n", string(jbytes))
	spew.Dump(jbytes, err2)
	require.Equal(t, bytes, jbytes)
}

type ShowInfo struct {
	StartTime JSONTime
}

type JSONTime time.Time

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(*t).Unix(), 10)), nil
}