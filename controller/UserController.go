package controller

import (
	_ "fmt"
	"github.com/GpsLypy/ginEssentail/common"
	"github.com/GpsLypy/ginEssentail/model"
	"github.com/GpsLypy/ginEssentail/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	db := common.GetDB()
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
		name = utils.RandomString(10)

	}
	//判断手机号是否存在
	//log.Println(name, telephone, password)
	if isTelephoneExist(db, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "该手机号已经注册",
		})
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "该手机号已经注册",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		PassWord:  string(hasedPassword),
	}

	db.Create(&newUser)

	//返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()
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
		name = utils.RandomString(10)

	}
	//判断手机号是否存在
	var user model.User
	db.Where("telephone= ?", telephone).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	//log.Println(name, telephone, password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token := "111"

	//返回结果
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功",
	})
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone= ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
