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
		},
		DisableCache: true,
	})
	return viewSetup
}
