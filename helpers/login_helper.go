package helpers

import "github.com/swp/access/configs"

func LoginIndexCSS() []string {
	rpta := []string{}
	if configs.Constants["ambiente_static"] == "desarrollo" {
		rpta = []string{
			"bower_components/bootstrap/dist/css/bootstrap.min",
			"bower_components/font-awesome/css/font-awesome.min",
			"bower_components/swp-backbone/assets/css/constants",
			"bower_components/swp-backbone/assets/css/login",
			"assets/css/login",
		}
	} else if configs.Constants["ambiente_static"] == "produccion" {
		rpta = []string{
			"dist/login.min",
		}
	}
	return rpta
}

func LoginIndexJS() []string {
	rpta := []string{}
	if configs.Constants["ambiente_static"] == "desarrollo" {
		rpta = []string{}
	} else if configs.Constants["ambiente_static"] == "produccion" {
		rpta = []string{}
	}
	return rpta
}
