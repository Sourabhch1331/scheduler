package internal

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"scheduler-server/pkg/utils"
	"strings"
)

func GetConfigString(key string) string {
	val := getValueForKey(strings.Split(key, "."), 0, getConfigData())
	if val == nil {
		log.Panicf("config for key not found: %s", key)
	}
	return utils.GetStringFromInterface(val)
}

func GetConfigInt64(key string) int64 {
	val := getValueForKey(strings.Split(key, "."), 0, getConfigData())
	if val == nil {
		log.Panicf("config for key not found: %s", key)
	}
	return utils.GetInt64FromInterface(val)
}

func getConfigData() interface{} {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Panicf("unable to open file")
	}
	defer func() {
		_ = file.Close()
	}()

	var data interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		log.Panicf("Invalid formate in config file")
	}

	return data
}

func getValueForKey(keys []string, idx int, data interface{}) interface{} {
	if data == nil || idx >= len(keys) {
		return nil
	}
	mp, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}

	if idx == len(keys)-1 {
		return mp[keys[idx]]
	}

	for k, v := range mp {
		if k == keys[idx] && v != nil {
			return getValueForKey(keys, idx+1, v)
		}
	}
	return nil
}
