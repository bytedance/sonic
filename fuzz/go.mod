module github.com/bytedance/sonic/fuzz

go 1.18

require (
	github.com/bytedance/gopkg v0.0.0-20221122125632-68358b8ecec6
	github.com/bytedance/sonic v1.11.5-alpha3
	github.com/davecgh/go-spew v1.1.1
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/bytedance/sonic/loader v0.1.0-rc // indirect
	github.com/cloudwego/base64x v0.1.2 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/bytedance/sonic => ../.
