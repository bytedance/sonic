//go:build (amd64 && go1.17 && !go1.24) || (arm64 && go1.20 && !go1.24)
// +build amd64,go1.17,!go1.24 arm64,go1.20,!go1.24

package sonic_test

import (
	"encoding/json"
	"testing"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
)

func TestUnescapedCharInString(t *testing.T) {
	var data =  "\"" + "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0A\x0B\x0C\x0D\x0E\x0F\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1A\x1B\x1C\x1D\x1E\x1F" +  "\""
	
	t.Run("Default", func(t *testing.T) {
		var sonicv, jsonv string
		sonice := sonic.Unmarshal([]byte(data), &sonicv)
		assert.NoError(t, sonice)
		assert.Equal(t, sonicv, data[1:len(data) - 1])
		
		jsone := json.Unmarshal([]byte(data), &jsonv)
		assert.True(t, strings.Contains(jsone.Error(), ("invalid char")))
		assert.Equal(t, jsonv, "")
	})

	t.Run("ValidateString", func(t *testing.T) {
		var sonicv, jsonv string
		api := sonic.Config {
			ValidateString: true,
		}.Froze()
		sonice := api.Unmarshal([]byte(data), &sonicv)
		assert.Error(t, sonice)
		assert.Equal(t, sonicv, "")
		
		jsone := json.Unmarshal([]byte(data), &jsonv)
		assert.True(t, strings.Contains(jsone.Error(), ("invalid char")))
		assert.Equal(t, jsonv, "")
	})
}
