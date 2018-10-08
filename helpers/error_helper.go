package helpers

import "github.com/swp/access/configs"

func ErrorAccessCSS() []string {
	rpta := []string{}
	if configs.Constants["ambiente_static"] == "desarrollo" {
		rpta = []string{
			"bower_components/bootstrap/dist/css/bootstrap.min",
			"bower_components/font-awesome/css/font-awesome.min",
			"assets/css/constants",
			"assets/css/error",
		}
	} else if configs.Constants["ambiente_static"] == "produccion" {
		rpta = []string{
			"dist/error.min",
		}
	}
	return rpta
}

func ErrorAccessJS() []string {
	rpta := []string{}
	if configs.Constants["ambiente_static"] == "desarrollo" {
		rpta = []string{}
	} else if configs.Constants["ambiente_static"] == "produccion" {
		rpta = []string{}
	}
	return rpta
}
