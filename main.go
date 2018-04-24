package main

import (
	"fmt"
	"github.com/akrylysov/algnhsa"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"om/handlers"
)

func init() {
	setUpViper()
	registerDatabase()

}

func main() {

	http.HandleFunc("/product", handlers.ProductHandler)

	// Insert the middleware
	fmt.Println(cast.ToString(viper.Get("base_url")))
	algnhsa.ListenAndServe(http.DefaultServeMux, nil)

}

//function to register the database to beego orm
func registerDatabase() {
	runmode := cast.ToString(viper.Get("runmode"))

	mysql := viper.Get(runmode + ".mysql").(map[string]interface{})
	mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string)
	log.Println("conf", mysqlConf)
	orm.RegisterDataBase("default", "mysql", mysqlConf)
	orm.Debug = true
}

//set up config file from conf folder
func setUpViper() {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetEnvPrefix("global")
}
