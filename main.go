package main

import (
	"fmt"
	"os"

	"github.com/GpsLypy/ginEssentail/common"
	_ "github.com/GpsLypy/ginEssentail/controller"
	"github.com/GpsLypy/ginEssentail/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	db = common.GetDB()
	fmt.Println(db)
	r := gin.Default()
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
