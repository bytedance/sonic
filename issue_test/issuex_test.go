/**
 * Copyright 2023 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package issue_test

import (
	// `encoding/json`
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/require"
)

var _ = ast.ErrNotExist
func TestUnmarshalErrorIXXXnMapSlice(t *testing.T) {
	subscribeContent := `
{
	"data": {
	"streamId": "",
	"audio": true,
	"ipType": "",
	"connId": "",
	"signalParams": {
		"params": {},
		"userId": "vendor168678041123_ins0",
		"deviceType": ""
	},
	"roomId": "",
	"screen": false,
	"data": true,
	"elapse": 7,
	"sessionId": "",
	"config": {},
	"peerConnectionMode": 0,
	"sdpInfo": {},
	"peerConnectionId": "",
	"eventSessionId": "",
	"video": true,
	"mediaConfig": {"downLink_VNM":{"qos_fb_pub_enable":1}}
	}
}`
	root, _ := sonic.GetFromString(subscribeContent)
	// root.LoadAll();
	sp := root.GetByPath("data", "signalParams")
	node := sp.Get("params")
	require.True(t, node.Exists())
	root.GetByPath("data", "mediaConfig")
	vs, _ := sp.Get("userId").String()
	require.Equal(t, "vendor168678041123_ins0", vs)
	vi, e := sp.Interface()
	require.NoError(t, e)
	fmt.Printf("%v", vi)
	// get2, err2 := node2.String()
	// require.Nil(t, err2)
	// require.Equal(t, get2, "vendor168678041123_ins0")
	vv, e := root.Interface()
	require.NoError(t, e)
	fmt.Printf("%v", vv)
	b, e := root.MarshalJSON()
	require.NoError(t, e)
	print(string(b))
}
