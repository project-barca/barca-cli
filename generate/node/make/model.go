package make

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pterm/pterm"

	"golang.org/x/text/language"
)

func Model(lang string, directory string, collection string, database string) {
	// Configurações para internacionalização do software
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("translate/en.json")
	bundle.LoadMessageFile("translate/fr.json")
	bundle.LoadMessageFile("translate/pt-BR.json")

	switch lang {
	case "portugues-brasileiro":
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	case "english":
		localizer = i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String(), language.French.String())
	case "francais":
		localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String(), language.BrazilianPortuguese.String())
	default:
		localizer = i18n.NewLocalizer(bundle, language.BrazilianPortuguese.String(), language.English.String(), language.French.String())
	}

	switch database {
	case "mongo":
		_, errorPath := os.Stat("./" + directory)

		localizeConfigErrorModelMongoExpress := i18n.LocalizeConfig{
			MessageID: "error_model_mongo_express",
		}
		localizeConfigSuccessModelMongoExpress := i18n.LocalizeConfig{
			MessageID: "success_model_mongodb_express",
		}
		resultErrorModelMongoExpress, _ := localizer.Localize(&localizeConfigErrorModelMongoExpress)
		resultSuccessMongoExpress, _ := localizer.Localize(&localizeConfigSuccessModelMongoExpress)

		if os.IsNotExist(errorPath) {
			pterm.Error.Println(resultErrorModelMongoExpress)
		} else {
			path := directory + "/model"

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(path, 0755)
			}
			f, err := os.Create(path + "/" + collection + ".js")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			firstLine := "module.exports = mongoose => {\n  var schema = mongoose.Schema(\n    {\n      \n\n    },\n    { timestamps: true }\n);\n\nschema.method(`toJson`, function() {\n  const { __v, _id, ...object } = this.toObject()\n  object.id = _id;\n  return object;\n})\n\nconst " + strings.Title(collection) + " = mongoose.model('" + collection + "', schema)\n\nreturn " + strings.Title(collection) + ";\n\n"
			data := []byte(firstLine)

			_, err2 := f.Write(data)

			if err2 != nil {
				log.Fatal(err2)
			}

			line2 := "};"
			data2 := []byte(line2)

			var idx int64 = int64(len(data))

			_, err3 := f.WriteAt(data2, idx)

			if err3 != nil {
				log.Fatal(err3)
			}

			pterm.Success.Println(resultSuccessMongoExpress)
		}
	case "mysql":
		// mysql
	case "dynamo":
		// dynamo
	case "oracle":
		// oracle
	case "postegresql":
		// postgresql
	case "sqlserver":
		// sqlserver
	case "db2":
		// db2
	default:
		// barcaSQL
	}

}
