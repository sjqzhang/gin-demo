package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjqzhang/gmock"
	"github.com/sjqzhang/gmock/mockdb"
	"github.com/sjqzhang/gmock/util"
	"github.com/spf13/viper"
	"jkdev.cn/api/common"
	"jkdev.cn/api/middleware"
	"jkdev.cn/api/route"
	"os"
	"path/filepath"
)

var MockDB *mockdb.MockGORM

func init() {

	MockDB = gmock.NewGORMFromDSN("./db", "mysql", "root:mock@tcp(127.0.0.1:63307)/mock?charset=utf8&parseTime=True&loc=Local")

	InitConfig()
	common.InitDB()

	r := gin.Default()

	r.Use(middleware.AuthMiddlewareMock(), util.DumpWithOptions(true, true, true, false, false, func(dumpStr string) {
		fmt.Println(dumpStr)
	}))

	r = route.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		go func() {
			panic(r.Run(":" + port))
		}()
	}
}

func Reset() {
	MockDB.ResetAndInit()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	workDir = filepath.Dir(workDir)
	os.Chdir(workDir)

	viper.SetConfigName("application")

	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
