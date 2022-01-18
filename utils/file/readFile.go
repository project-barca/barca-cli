package file

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

func ReadXML(fileName string, key interface{}) interface{} {
	data, _ := ioutil.ReadFile(fileName)

	interfaceKey := key
	_ = xml.Unmarshal([]byte(data), &interfaceKey)

	return interfaceKey
}

func ReadJSON(fileName string, key interface{}) interface{} {
	data, _ := ioutil.ReadFile(fileName)

	interfaceKey := key
	_ = json.Unmarshal([]byte(data), &interfaceKey)

	return interfaceKey
}
