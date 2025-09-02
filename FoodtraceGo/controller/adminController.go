package controller

import (
	"FoodtraceGo/dbs"
	"FoodtraceGo/models"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
}

func (a AdminController) UserRegister(c *gin.Context) {
	var req struct {
		Username string `json:"username"  bindings: "required" `
		Password string `json:"password"  bindings: "required" `
	}
	// 自动绑定和验证
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "参数验证失败",
		})
		return
	}
	//检查用户是否存在
	user := models.Users{}
	result := dbs.DB.Where("username = ?", req.Username).First(&user)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error":   "USERNAME_EXISTS",
			"message": "用户名已被注册",
		})
		return
	}
	privateKey, address, err := GeneratePrivateKeyWithAddress()
	if err != nil {
		c.Error(err)
	}
	fmt.Printf("私钥: %s\n", privateKey)
	fmt.Printf("地址: %s\n", address)

	//生成
	creat := models.Users{
		Username: req.Username,
		Password: req.Password,
		Address:  address,
		Private:  privateKey,
	}
	dbs.DB.Create(&creat)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    "创建成功",
		"status":  true,
	})
}

func (a AdminController) UserLogin(c *gin.Context) {
	var req struct {
		Username string `json:"username"  bindings: "required" `
		Password string `json:"password"  bindings: "required" `
	}

	// 自动绑定和验证
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"message": "参数验证失败",
		})
		return
	}
	user := models.Users{}
	result := dbs.DB.Where("username = ? And password = ?", req.Username, req.Password).First(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "该用户不存在",
			"status":  true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"address": user.Address,
		"message": "Login Success",
	})
}

// GeneratePrivateKeyWithAddress 生成私钥和对应的地址
func GeneratePrivateKeyWithAddress() (privateKeyHex, address string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	// 获取私钥十六进制
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex = hex.EncodeToString(privateKeyBytes)

	// 获取地址
	address = crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	return privateKeyHex, address, nil
}
