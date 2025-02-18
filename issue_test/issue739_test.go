package issue_test

import (
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
)

type myfoo struct{}

func TestIssue739(t *testing.T) {
	var bar myfoo
	s := `{"a":"b
c"}`
	fmt.Println(sonic.ConfigDefault.UnmarshalFromString(s, &bar))
}