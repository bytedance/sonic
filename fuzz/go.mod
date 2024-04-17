module github.com/bytedance/sonic/fuzz

go 1.18

require (
	github.com/bytedance/gopkg v0.0.0-20221122125632-68358b8ecec6
	github.com/bytedance/sonic v1.11.5-alpha3
	github.com/davecgh/go-spew v1.1.1
	github.com/stretchr/testify v1.8.1
)

require github.com/cloudwego/iasm v0.2.0 // indirect

replace github.com/bytedance/sonic => ../.
