package insert

import (
	"fmt"

	"github.com/project-barca/barca-cli/utils/file"
)

func ColumnsMySQLModel(path string, table string, field string, types string) {
	var typeData string

	switch types {
	case "string":
		typeData = "String"
	case "int":
		typeData = "Number"
	case "bool":
		typeData = "Boolean"
	case "date":
		typeData = "Date"
	default:
	}

	err := file.InsertStringToFile(path+"/model/"+table+".js", "      "+field+": { type: "+typeData+" }", 4)
	if err != nil {
		fmt.Printf("Deu erro ao tentar escrever no arquivo")
	}

}
