package providers

import (
	"galaveg/config"
	"galaveg/utils/path"
	"html/template"
	"strings"
)

func MustHtmlEngine(cfg *config.Config) *template.Template {
	funcMap := template.FuncMap{}

	funcMap["eq"] = func(a, b string) bool { return a == b }
	funcMap["eqInt"] = func(a, b int) bool { return a == b }
	funcMap["ne"] = func(a, b string) bool { return a != b }
	funcMap["neInt"] = func(a, b int) bool { return a != b }
	funcMap["replace"] = strings.ReplaceAll
	funcMap["unless"] = func(v any) bool {
		return v == nil || v == "" || v == false
	}
	funcMap["dict"] = func(values ...interface{}) map[string]interface{} {
		dict := make(map[string]interface{})
		for i := 0; i < len(values); i += 2 {
			key := values[i].(string)
			value := values[i+1]
			dict[key] = value
		}
		return dict
	}
	funcMap["sub1"] = func(x int) int { return x - 1 }
	funcMap["startsWith"] = func(s, prefix string) bool {
		return strings.HasPrefix(s, prefix)
	}

	files := path.MustCollectFilepathBySuffix(cfg.GetFolder(cfg.Templates.Html.Path), ".gohtml")

	return template.Must(template.New("").Delims("{{", "}}").Funcs(funcMap).ParseFiles(files...))
}
