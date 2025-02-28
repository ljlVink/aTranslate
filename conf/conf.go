package conf

type Yaml_config struct {
	General General
}

type General struct {
	Openai_key   string
	Openai_url   string
	Openai_model string
}

var Prompt = `你需要把这张图片的内容完全翻译成中文。
1.需要使用latex的地方要用$包裹，另起一行居中用$$。
2.不要加分隔线。
3.如果翻译到"定理"和"定义"等名词需要加粗，如**定理** ：xxx。
4.如果有小标题用###`
