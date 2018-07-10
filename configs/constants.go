package configs

var Constants = make(map[string]string)

func SetConstants() {
	Constants["BASE_URL"] = "http://localhost:8080/"
	Constants["STATIC_URL"] = "http://localhost:8080/public/"
}
