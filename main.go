package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varcahr(11);not null;unique"`
	PassWord  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	//defer db.Clo
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据验证

		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须11位",
			})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能少于6位",
			})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)

		}
		//判断手机号是否存在
		log.Println(name, telephone, password)
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "该手机号已经注册",
			})
			return
		}
		//创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
		}
		//返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func RandomString(n int) string {
	var letters = []byte("fdsfdfsfsdd")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	dsn := "root:Mysql123..@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faild to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db

}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("telephone= ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
