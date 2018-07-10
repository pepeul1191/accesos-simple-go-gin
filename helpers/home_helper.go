package helpers

import "github.com/ginv2/configs"

func HomeIndexCSS() []string {
	rpta := []string{}
	if configs.Constants["ambiente_static"] == "desarrollo" {
		rpta = []string{
			"bower_components/bootstrap/dist/css/bootstrap.min",
			"bower_components/font-awesome/css/font-awesome.min",
			"bower_components/swp-backbone/assets/css/constants",
			"bower_components/swp-backbone/assets/css/dashboard",
			"bower_components/swp-backbone/assets/css/table",
			"assets/css/constants",
			"assets/css/styles",
		}
	} else if configs.Constants["ambiente_static"] == "produccion" {
		rpta = []string{
			"dist/home.min",
		}
	}
	return rpta
}

func HomeIndexJS() []string {
	rpta := []string{}
	if configs.Constants["ambiente_static"] == "desarrollo" {
		rpta = []string{
			"bower_components/jquery/dist/jquery.min",
			"bower_components/bootstrap/dist/js/bootstrap.min",
			"bower_components/underscore/underscore-min",
			"bower_components/backbone/backbone-min",
			"bower_components/handlebars/handlebars.min",
			"bower_components/swp-backbone/layouts/application",
			"bower_components/swp-backbone/views/table",
			"bower_components/swp-backbone/views/modal",
			"models/ubicaciones/departamento",
			"models/ubicaciones/provincia",
			"collections/ubicaciones/departamento_collection",
			"collections/ubicaciones/provincia_collection",
			"data/ubicaciones/tabla_distrito",
			"data/ubicaciones/tabla_provincia",
			"data/ubicaciones/tabla_departamento",
			"views/ubicaciones/tabla_provincia_view",
			"views/ubicaciones/tabla_departamento_view",
			"views/ubicaciones/ubicacion_view",
			"routes/ubicaciones",
		}
	} else if configs.Constants["ambiente_static"] == "produccion" {
		rpta = []string{
			"dist/home.min",
		}
	}
	return rpta
}
