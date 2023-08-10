module github.com/bytedance/sonic/fuzz

go 1.18

require (
	github.com/bytedance/gopkg v0.0.0-20221122125632-68358b8ecec6
	github.com/bytedance/sonic v1.10.0-rc
	github.com/davecgh/go-spew v1.1.1
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/bytedance/sonic => ../.
