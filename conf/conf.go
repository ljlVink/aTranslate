package conf

type Yaml_config struct {
	General General
}

type General struct{
	Openai_key string 
	Openai_url string
}