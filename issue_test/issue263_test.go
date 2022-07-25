package issue_test

import (
    `bytes`
    `strings`
    `testing`

    `github.com/bytedance/sonic/decoder`
    `github.com/stretchr/testify/require`
)

type Response struct {
    Menu Menu `json:"menu"`
}

type Menu struct {
    Items []*Item `json:"items"`
}

type Item struct {
    ID string `json:"id"`
}

func (i *Item) UnmarshalJSON(buf []byte) error {    
	return nil
}

func TestIssue263(t *testing.T) {
    q := `{
		"menu": {
			"items": [
				{`+strings.Repeat(" ", 1024)+`}
			]
		}
	}`

    var response Response
    require.Nil(t, decoder.NewStreamDecoder(bytes.NewReader([]byte(q))).Decode(&response))

    q = `{
        "menu": {
            "items": [
                {"a":"`+strings.Repeat("b", 2048)+`"}
            ]
        }
    }`
    
    require.Nil(t, decoder.NewStreamDecoder(bytes.NewReader([]byte(q))).Decode(&response))
}