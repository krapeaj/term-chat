package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func GetConfig(data []byte) map[string]interface{} {
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		panic(err)
	}
	c, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		panic(fmt.Errorf("yaml file is not a map[interface{}]interface{}"))
	}
	config, err := convertKeysToStrings(c);
	if err != nil {
		panic(err)
	}
	return config
}

func convertKeysToStrings(m map[interface{}]interface{}) (map[string]interface{}, error) {
	newMap := make(map[string]interface{})
	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}
		vMap, ok := v.(map[interface{}]interface{})
		if ok {
			var err error
			if v, err = convertKeysToStrings(vMap); err != nil {
				return nil, err
			}
		}
		newMap[str] = v
	}

	return newMap, nil
}
