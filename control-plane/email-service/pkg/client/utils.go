package client

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"

	"github.com/run-ai/runai-common-packages/logging"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func GetStructName(templateData interface{}) string {
	if t := reflect.TypeOf(templateData); t.Kind() == reflect.Ptr {
		return "*" + toSnakeCase(t.Elem().Name())
	} else {
		return toSnakeCase(t.Name())
	}
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func StructToMap(obj interface{}) map[string]interface{} {
	b, err := json.Marshal(obj)
	if err != nil {
		logging.Errorf("marshal error: %s", err)
		return nil
	}
	data := new(map[string]interface{})
	err = json.Unmarshal(b, data)
	if err != nil {
		logging.Errorf("unmarshal error: %s", err)
		return nil
	}
	return *data
}
