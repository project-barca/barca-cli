package file

import (
	"encoding/xml"
	"io/ioutil"
)

func ReadXML(fileName string, key interface{}) interface{} {
	data, _ := ioutil.ReadFile(fileName)

	interfaceKey := key
	_ = xml.Unmarshal([]byte(data), &interfaceKey)

	return interfaceKey
}
