module github.com/bytedance/sonic/fuzz

go 1.18

require (
	github.com/bytedance/gopkg v0.1.3
	github.com/bytedance/sonic v1.11.5-alpha3
	github.com/davecgh/go-spew v1.1.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/sys v0.22.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/bytedance/sonic => ../.
