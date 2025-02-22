package conf

type Yaml_config struct {
	General General
}

type General struct {
	Openai_key   string
	Openai_url   string
	Openai_model string
}

var Prompt = `你需要把这张图片的内容完全翻译成中文。其中页脚不需要翻译。不要加分隔符。如果翻译到"定理"和"定义"等名词需要加粗，如**定理** ：xxx。公式需要另起一行居中。并且使用latex.如果有小标题用###`
