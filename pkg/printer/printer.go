package printer

import (
	"bytes"
	"gopkg.in/yaml.v3"
)

func PrintYaml[T interface{}](data T) error {
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)

	if err := yamlEncoder.Encode(&data); err != nil {
		return err
	}
	println(string(b.Bytes()))
	return nil
}
