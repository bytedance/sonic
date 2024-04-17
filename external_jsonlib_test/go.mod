module github.com/bytedance/sonic/external_jsonlib_test

go 1.18

require (
	github.com/buger/jsonparser v1.1.1
	github.com/bytedance/sonic v1.11.5-alpha3
	github.com/goccy/go-json v0.9.11
	github.com/json-iterator/go v1.1.12
	github.com/stretchr/testify v1.8.1
	github.com/tidwall/gjson v1.14.3
	github.com/tidwall/sjson v1.2.5
)

require github.com/cloudwego/iasm v0.2.0 // indirect

replace github.com/bytedance/sonic => ../.
