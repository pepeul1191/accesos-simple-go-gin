package configs

var Constants = make(map[string]string)

func SetConstants() {
	Constants["BASE_URL"] = "http://localhost:3000/"
	Constants["STATIC_URL"] = "http://localhost:3000/public/"
	Constants["ambiente_static"] = "desarrollo" //produccion
}
