## Go-Gin

Instalación de dependencias:

    $ go get -u github.com/gin-gonic/gin
    $ go get -u github.com/jinzhu/gorm

Las variables globales deben empezar con mayúscula

Instalar dependencias:

    $ govendor install

Autoreload ante cambios:

    $ go get github.com/codegangsta/gin
    $ gin run main.go

Migraciones con DBMATE - ubicaciones:

    $ dbmate -d "db/migrations" -e "ACCESS" new <<nombre_de_migracion>>
    $ dbmate -d "db/migrations" -e "ACCESS" up

### BDD con cucumber

Instalación de dependencias:

	$ gem install cucumber
	$ bundler install

Ejecutar pruebas:

	$ cucumber 
    $ cucumber features/<file_name>.feature

---

Fuentes:

+ https://github.com/gin-gonic/gin
+ https://www.youtube.com/watch?v=8s5YvgS5TuI
+ http://motion-express.com/blog/gorm:-a-simple-guide-on-crud
+ http://gorm.io/
+ https://github.com/jinzhu/gorm
+ https://stackoverflow.com/questions/38541598/why-is-it-possible-to-export-variable-of-private-type
+ https://github.com/codegangsta/gin
+ https://github.com/gin-gonic/gin
+ https://github.com/foolin/gin-template
+ https://markhneedham.com/blog/2016/12/23/go-templating-with-the-gin-web-framework/
+ https://github.com/foolin/gin-template/blob/master/examples/advance/views/index.tpl
+ https://github.com/foolin/gin-template/blob/master/examples/advance/main.go
+ https://golang.org/pkg/text/template/
+ https://github.com/revel/revel/blob/master/template_functions.go
+ https://gobyexample.com/json
+ https://stackoverflow.com/questions/32443738/setting-up-route-not-found-in-gin
+ https://stackoverflow.com/questions/43153181/json-and-struct-composition-in-go
+ http://gorm.io/docs/delete.html
+ http://gorm.io/docs/update.html
+ http://gorm.io/docs/create.html
+ https://forum.golangbridge.org/t/gorm-save-is-not-returning-primary-key/6289/11
+ https://github.com/badoux/checkmail
+ https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/
+ https://codereview.stackexchange.com/questions/125846/encrypting-strings-in-golang