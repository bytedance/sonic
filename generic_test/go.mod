module github.com/bytedance/sonic/generic_test

go 1.18

require (
	github.com/bytedance/sonic v1.11.5-alpha3
	github.com/go-json-experiment/json v0.0.0-20220603215908-554802c1e539
	github.com/goccy/go-json v0.9.4
	github.com/json-iterator/go v1.1.12
)

require github.com/cloudwego/iasm v0.2.0 // indirect

replace github.com/bytedance/sonic => ../.
