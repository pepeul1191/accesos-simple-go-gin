package configs

import (
	"html/template"
	"time"

	gintemplate "github.com/foolin/gin-template"
)

func GetViewSetup() *gintemplate.TemplateEngine {
	var viewSetup = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		//Master:    "layouts/master",
		Partials: []string{
			"partials/blank_footer",
			"partials/blank_header",
		},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
			"loadCSS": func(arrayCss []string) string {
				var rpta = ""
				for i := 0; i < len(arrayCss); i++ {
					var temp = "<link rel=\"stylesheet\" type=\"text/css\" href=\"" + Constants["STATIC_URL"] + arrayCss[i] + ".css\"/>"
					rpta = rpta + temp
				}
				return rpta
			},
			"loadJS": func(arrayJs []string) string {
				var rpta = ""
				for i := 0; i < len(arrayJs); i++ {
					var temp = "<script src=\"" + Constants["STATIC_URL"] + arrayJs[i] + ".js\" type=\"text/javascript\"></script>"
					rpta = rpta + temp
				}
				return rpta
			},
			"raw": func(text string) template.HTML { // de revel
				return template.HTML(text)
			},
		},
		DisableCache: true,
	})
	return viewSetup
}
