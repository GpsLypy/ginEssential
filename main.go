package main

import (
	"fmt"

	"github.com/GpsLypy/ginEssentail/common"
	_ "github.com/GpsLypy/ginEssentail/controller"
	"github.com/GpsLypy/ginEssentail/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	db = common.GetDB()
	fmt.Println(db)
	r := gin.Default()
	r = router.CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
