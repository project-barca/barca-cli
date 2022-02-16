package projeto

import (
	"path/filepath"
	"strings"
	"unicode"

	"github.com/project-barca/barca-cli/generate/node/dependencies"
)

func WhatsIsLanguages(extensions []string) []string {
	var languages []string

	increment_xml := 0
	increment_txt := 0
	increment_dll := 0
	increment_cache := 0
	increment_csproj := 0
	increment_go := 0
	increment_js := 0
	increment_java := 0
	increment_ts := 0
	increment_rb := 0
	increment_erb := 0
	increment_php := 0
	increment_cs := 0
	increment_rs := 0
	increment_c := 0
	increment_cpp := 0
	increment_py := 0
	increment_html := 0
	increment_css := 0
	increment_less := 0
	increment_scss := 0
	increment_sql := 0
	increment_md := 0
	//percent := 100 / len(extensions)
	for i := 0; i < len(extensions); i++ {
		switch extensions[i] {
		case ".go":
			increment_go++
		case ".js":
			increment_js++
		case ".ts":
			increment_ts++
		case ".php":
			increment_php++
		case ".cs":
			increment_cs++
		case ".csproj":
			increment_csproj++
		case ".c":
			increment_c++
		case ".cpp":
			increment_cpp++
		case ".rs":
			increment_rs++
		case ".py":
			increment_py++
		case ".rb":
			increment_rb++
		case ".erb":
			increment_erb++
		case ".html":
			increment_html++
		case ".less":
			increment_less++
		case ".css":
			increment_css++
		case ".scss":
			increment_scss++
		case ".md":
			increment_md++
		case ".sql":
			increment_sql++
		case ".java":
			increment_java++
		case ".xml":
			increment_xml++
		case ".dll":
			increment_dll++
		case ".txt":
			increment_txt++
		case ".cache":
			increment_cache++
		}
	}

	var stacks = map[string]int{
		"rust":       increment_rs,
		"go":         increment_go,
		"c":          increment_c,
		"c++":        increment_cpp,
		"c#":         increment_cs,
		"python":     increment_py,
		"ruby":       increment_rb,
		"javascript": increment_js,
		"typescript": increment_ts,
		"php":        increment_php,
		"java":       increment_java,
		"html":       increment_html,
		"css":        increment_css,
		"less":       increment_less,
		"scss":       increment_scss,
		"sql":        increment_sql,
		"dll":        increment_dll,
		"xml":        increment_xml,
		"markdown":   increment_md,
	}

	var key string
	var value int

	for key, value = range stacks {
		if value == 0 {
			delete(stacks, key)
		}
	}

	for key, value = range stacks {
		languages = append(languages, key)
	}

	// fmt.Print("Rust: " + string(percent*increment_rs) + "%")
	// fmt.Print("Javascript: " + string(percent*increment_js) + "%")
	// fmt.Print("TypeScript: " + string(percent*increment_ts) + "%")
	// fmt.Print("C: " + string(percent*increment_c) + "%")
	// fmt.Print("CSharp: " + string(percent*increment_cs) + "%")
	// fmt.Print("PYthon: " + string(percent*increment_py) + "%")
	// fmt.Print("HTML: " + string(percent*increment_html) + "%")
	// fmt.Print("CSS: " + string(percent*increment_css) + "%")
	// fmt.Print("SQL: " + string(percent*increment_sql) + "%")

	// if increment_go >= 1 {
	// 	languages = append(languages, filepath.Ext("Go"))
	// }
	//	fmt.Print(languages)

	return languages
}

func WhatsIsFrameworks(languages []string, path string) []string {

	var frameworks []string
	var modules []string

	for i := 0; i < len(languages); i++ {
		switch languages[i] {
		case "javascript":
			modules = append(modules, "total", "feather", "loopback", "koa", "sails", "nest", "meteor", "socket", "derby", "express", "@adonisjs/framework", "rapi", "react", "vue", "angular", "bootstrap", "ember")
			for y := 0; y < len(modules); y++ {
				if dependencies.IfExistsModule(path, languages[i], modules[y]) == true {
					frameworks = append(frameworks, strings.ToLowerSpecial(unicode.TurkishCase, modules[y]))
				}
			}
		case "typescript":
			modules = append(modules, "express", "nest", "loopback", "feather")

		case "python":
			modules = append(modules, "Django", "Flask", "pyramid", "falcon", "cherrypy", "bottle", "web2py", "tornado", "cubicweb", "dash")
			for y := 0; y < len(modules); y++ {
				if dependencies.IfExistsModule(path, languages[i], modules[y]) == true {
					frameworks = append(frameworks, strings.ToLowerSpecial(unicode.TurkishCase, modules[y]))
				}
			}
		case "java":
			modules = append(modules, "spring", "play", "gwt", "hibernate", "struts", "vaadin", "grails", "jsf", "grails")

		case "golang":
			modules = append(modules, "gorilla", "echo", "gin", "beego", "fasthttp", "kit", "fiber", "iris")

		case "php":
			modules = append(modules, "laravel", "php", "symfony", "zend", "phalcon", "cake", "yii", "codeigniter")
			for y := 0; y < len(modules); y++ {
				if dependencies.IfExistsModuleComposer(path, modules[y]) == true {
					frameworks = append(frameworks, strings.ToLowerSpecial(unicode.TurkishCase, modules[y]))
				}
			}
		case "ruby":
			modules = append(modules, "rubyonrails", "sinatra-activerecord", "roda", "camping", "cuba", "nancy", "remaze", "ramaze", "goliath", "hanami", "padrinho")

		}
	}

	return frameworks
}

func WhatsIsDatabase(fileName []string) []string {
	var extensions []string
	for i := 0; i < len(fileName); i++ {
		extensions = append(extensions, "da", "de")
	}
	return extensions
}

func WhatsIsWebServer(fileName []string) []string {
	var extensions []string
	for i := 0; i < len(fileName); i++ {
		extensions = append(extensions, filepath.Ext(fileName[i]))
	}
	return extensions
}
