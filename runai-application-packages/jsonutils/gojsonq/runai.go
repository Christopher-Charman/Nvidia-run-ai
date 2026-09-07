package gojsonq

import (
	"encoding/json"
	"strings"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/logging"
)

// JsonqDelim - the default jsonQ delimiter is "." which is ok for out synthetic resources.
// note that runai job can contain elements like "resources.gpu.limit". if this is ever
// used for json representation of runaijob - choose a different delimiter.
const JsonqDelim = "."

// LoadJson loads any object to a generic map
func LoadJson(anyObject interface{}) (map[string]interface{}, error) {
	//
	//   convert the object to json string
	//
	arJsonBytes, err := json.MarshalIndent(anyObject, "", "\t")
	if err != nil {
		return nil, err
	}

	return LoadJsonBytes(arJsonBytes)
}

func ToBytes(jsonMap map[string]interface{}) []byte {
	bytes, err := json.Marshal(jsonMap)
	if err != nil {
		logging.Errorf("failed to convert json to bytes: %v", err)
	}
	return bytes
}

func LoadJsonBytes(arJsonBytes []byte) (map[string]interface{}, error) {
	//
	//  load the string to a generic map.
	//
	jsonMap := make(map[string]interface{})
	decoder := json.NewDecoder(strings.NewReader(string(arJsonBytes)))

	//
	//  we want to use numbers rather than the mighty default float64, which is not accepted by some
	//  json libraries
	//
	decoder.UseNumber()

	err := decoder.Decode(&jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}

// NewJsonQ loads an object to JsonQ
func NewJsonQ(anObject interface{}) (*JSONQ, error) {

	jsonMap, err := LoadJson(anObject)
	if err != nil {
		return nil, err
	}
	jq := New(SetSeparator(JsonqDelim)).FromInterface(jsonMap)
	if jq.Error() != nil {
		return nil, jq.Error()
	}

	return jq, nil
}

// CleanMap "cleans" entries which are null or having no keys
// for example:
//
//	environment:
//	  image: null
//	  environmentVariables: {}
//	  imagePullPolicy: Always
//
// after clean should look like this:
//
//	environment:
//	  imagePullPolicy: always
//
// because the image and environmentVariables are empty
func CleanMap(aMap map[string]interface{}) {
	for key, value := range aMap {
		//  handle only entries which are maps
		if nestedMap, ok := value.(map[string]interface{}); ok {
			// recursively cleans any nested map
			CleanMap(nestedMap)

			// remove nil entries
			for innerKey, innerValue := range nestedMap {
				if innerValue == nil {
					delete(nestedMap, innerKey)
				}
			}

			// if nested map is empty remove it
			if len(nestedMap) == 0 {
				delete(aMap, key)
			}
		}
	}
}

// RemoveIdenticalKeys remove from destMap all the keys that also exist in sourceMap
// for example
//
//	destMap:
//	  environmentVariables:
//	    AAA: value-of-a
//	    BBB; value-of-b
//	sourceMap:
//	    BBB: another-value-of-b
//
// the result will be //
//
//	destMap:
//	  AAA: value-of-a
//
// and this works recursively on nested keys as well
func RemoveIdenticalKeys(destMap, sourceMap map[string]interface{}) {
	for key := range sourceMap {
		if nestedMap, ok := sourceMap[key].(map[string]interface{}); ok {
			if nestedMap1, ok := destMap[key].(map[string]interface{}); ok {
				RemoveIdenticalKeys(nestedMap1, nestedMap)
			}
		} else {
			delete(destMap, key)
		}
	}
}
