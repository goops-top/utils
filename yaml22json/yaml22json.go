package yaml22json

import (
	jyaml "github.com/ghodss/yaml"
)

func JSON2Yaml(jsonData string) (yamlData string, err error) {
	jObj := []byte(jsonData)
	yObj, err := jyaml.JSONToYAML(jObj)
	if err != nil {
		yamlData = ""
		return
	}
	yamlData = string(yObj)
	return
}

func Yaml2JSON(yamlData string) (jsonData string, err error) {
	yObj := []byte(yamlData)
	jObj, err := jyaml.YAMLToJSON(yObj)

	if err != nil {
		yamlData = ""
		return
	}
	jsonData = string(jObj)
	return

}
