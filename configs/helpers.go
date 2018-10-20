package configs

import (
	"fmt"
	"math/rand"
	"time"
)

func HelperLoadCSS(css_array []string) (out string) {
	var rpta string = ""
	for i := 0; i < len(css_array); i++ {
		rpta = rpta + "<link href=\"" + Constants["STATIC_URL"] + css_array[i] + ".css\" rel=\"stylesheet\" type=\"text/css\"/>"
	}
	out = rpta
	return
}

func HelperLoadJS(js_array []string) (out string) {
	var rpta string = ""
	for i := 0; i < len(js_array); i++ {
		rpta = rpta + "<script type=\"text/javascript\" src=\"" + Constants["STATIC_URL"] + js_array[i] + ".js\"></script>"
	}
	out = rpta
	return
}

func EmailFormatCheck(email string) (out bool) {
	err := ValidateFormat("ç$€§/az@gmail.com")
	rpta := true
	if err != nil {
		fmt.Println(err)
		rpta = false
	}
	return rpta
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
