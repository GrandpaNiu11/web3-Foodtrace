package controller

import (
	"FoodtraceGo/bindings"
	"FoodtraceGo/dbs"
	"FoodtraceGo/models"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	FoodClient           *ethclient.Client
	FoodContractInstance *bindings.Foodtrace
	initOnce             sync.Once
	initializationError  error
	privateKey           string
)

type FoodController struct {
}

func init() {
	InitFoodtrace()
}

func (f FoodController) AddLogisticsProvider(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}
	address, _ := data["address"].(string)
	addr := common.HexToAddress(address)
	hase := addLogisticsProviderff(FoodContractInstance, FoodClient, privateKey, addr)
	c.JSON(200, gin.H{
		"message": "success",
		"hash":    hase,
		"status":  true,
	})

}

// 食品生产方注册食品
func (f FoodController) AddRegisterFood(c *gin.Context) {
	var req struct {
		Name  string `json:"name"  bindings: "required" `
		Price string ` json:"price" bindings: "required" `
		Store string ` json:"store" bindings: "required" `
		Date  string ` json:"date" bindings: "required" `
		Code  string ` json:"code" bindings: "required" `
	}
	// 自动绑定和验证
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "参数验证失败",
		})
		return
	}
	details := fmt.Sprintf("价格:%.2f,仓库:%s,出仓日期:%s,编码:%s", req.Price, req.Store, req.Date, req.Code)

	hase := registerFood(FoodContractInstance, FoodClient, privateKey, req.Name, details)

	key, err := GetAddressFromPrivateKey(privateKey)
	if err != nil {
		c.Error(err)
		return
	}
	price, err := strconv.ParseFloat(req.Price, 64)
	//插入数据库
	var data = models.FoodItems{
		FoodName:          req.Name,
		OutboundPrice:     price,
		ShippingWarehouse: req.Store,
		OutboundDate:      req.Date,
		Code:              req.Code,
		Description:       details,
		BlockchainAddress: key,
		TransactionHash:   hase,
	}
	tx := dbs.DB.Create(&data)
	fmt.Println(tx.RowsAffected)
	c.JSON(200, gin.H{
		"message": "success",
		"hash":    hase,
		"status":  true,
	})
}

// 物流管理方记录食品物流信息
func (f FoodController) AddUpdateLogistice(c *gin.Context) {
	// 获取特定的 Cookie，例如 "isAuthenticated"
	cookie, err := c.Request.Cookie("address")
	if err != nil {
		if err == http.ErrNoCookie {
			// Cookie 不存在
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未找到认证 Cookie"})
			return
		}
		// 其他错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取 Cookie 失败: " + err.Error()})
		return
	}

	// 获取 Cookie 值
	authToken := cookie.Value
	//根据地址查密钥
	user := models.Users{}
	dbs.DB.Where("address = ?", authToken).First(&user)
	//
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}
	foodIdStr, _ := data["foodId"].(string)
	// 创建 big.Int 对象
	foodId := new(big.Int)

	// 将字符串转换为 big.Int，10 表示十进制
	_, success := foodId.SetString(foodIdStr, 10)
	logisticsInfo, _ := data["logisticsInfo"].(string)
	if !success {
		fmt.Errorf("failed to parse foodId as big.Int: %s", foodIdStr)
		return
	}
	log.Println(logisticsInfo)
	hase := updateLogistice(FoodContractInstance, FoodClient, user.Private, foodId, logisticsInfo)
	c.JSON(200, gin.H{
		"message": "success",
		"hash":    hase,
		"status":  true,
	})
}

func (f FoodController) LogisticsRecordedCount(c *gin.Context) {
	count, err := FoodContractInstance.LogisticsRecordedCount(nil)
	if err != nil {
		log.Fatal("获取食品数量失败:", err)
	}
	fmt.Printf("当前食品数量: %d\n", count)
	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"count": count.String(), // 使用 String() 方法将 big.Int 转换为字符串
		},
		"status": true,
	})
}

func (f FoodController) FoodItemCount(c *gin.Context) {
	count, err := FoodContractInstance.FoodItemCount(nil)
	if err != nil {
		log.Fatal("获取食品数量失败:", err)
	}
	fmt.Printf("当前食品数量: %d\n", count)
	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"count": count.String(), // 使用 String() 方法将 big.Int 转换为字符串
		},
		"status": true,
	})
}

func (f FoodController) FoodItems(c *gin.Context) {
	foodIdStr := c.Query("foodid")
	if foodIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodid 参数缺失"})
		return
	}
	bigInt, _ := StringToBigInt(foodIdStr)

	fmt.Printf("foodIdStr: %s\n", bigInt)
	data, err := FoodContractInstance.FoodItems(nil, bigInt)
	if err != nil {
		log.Fatal("获取食品数量失败:", err)
	}
	fmt.Printf("当前食品数量: %d\n", data)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
		"status":  true,
	})
}

func (f FoodController) LogisticsProviders(c *gin.Context) {
	addressStr := c.Query("address")
	if addressStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodid 参数缺失"})
		return
	}
	addr := common.HexToAddress(addressStr)
	flag, err := FoodContractInstance.LogisticsProviders(nil, addr)
	if err != nil {
		log.Fatal("获取食品数量失败:", err)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    flag,
		"status":  true,
	})
}

func (f FoodController) GetFoodInfo(c *gin.Context) {
	foodIdStr := c.Query("foodid")
	if foodIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "foodid 参数缺失"})
		return
	}
	bigInt, _ := StringToBigInt(foodIdStr)
	data, err := FoodContractInstance.GetFoodInfo(nil, bigInt)
	if err != nil {
		log.Fatal("获取食品数量失败:", err)
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
		"status":  true,
	})

}

func (f FoodController) SelectFoodItems(c *gin.Context) {
	var data []models.FoodItems // 使用切片存储所有记录

	// 查询所有记录
	if err := dbs.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败: " + err.Error()})
		return
	}
	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
		"status":  true,
	})
}

// InitFoodtrace 初始化区块链连接和合约实例
func InitFoodtrace() error {
	cfg, err := ini.Load("./my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	url := cfg.Section("sepolia").Key("url").String()
	address := cfg.Section("sepolia").Key("address").String()
	privateKey = cfg.Section("sepolia").Key("privateKey").String()

	fmt.Println("开始连接以太坊节点")
	initOnce.Do(func() {
		// 1. 连接以太坊节点"http://127.0.0.1:7565
		client, err := ethclient.Dial(fmt.Sprintf("http://%v", url))
		if err != nil {
			initializationError = err
			return
		}
		FoodClient = client

		// 2. 创建合约实例
		contractAddress := common.HexToAddress(address)
		contractInstance, err := bindings.NewFoodtrace(contractAddress, FoodClient)
		if err != nil {
			initializationError = err
			return
		}
		FoodContractInstance = contractInstance

		log.Println("✅ 食品溯源合约初始化成功")
	})
	return initializationError
}

// //食品生产方注册食品
func registerFood(instance *bindings.Foodtrace, client *ethclient.Client, privateKey string, name string, details string) (hase string) {
	// 创建交易选项
	auth, err := createTransactOpts(client, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 调用注册食品方法
	tx, err := instance.RegisterFood(auth, name, details)
	if err != nil {
		log.Fatal("注册食品失败:", err)
	}

	fmt.Printf("✅ 食品注册交易已发送! 交易哈希: %s\n", tx.Hash().Hex())
	waitForTransaction(client, tx, "食品注册")
	return tx.Hash().Hex()
}

// 物流管理方记录食品物流信息
func updateLogistice(instance *bindings.Foodtrace, client *ethclient.Client, privateKey string, foodId *big.Int, logisticsInfo string) (hase string) {
	// 创建交易选项
	auth, err := createTransactOpts(client, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 调用注册食品方法
	tx, err := instance.UpdateLogistice(auth, foodId, logisticsInfo)
	if err != nil {
		log.Fatal("物流管理方记录食品物流信息:", err)
	}

	fmt.Printf("✅ 食品注册交易已发送! 交易哈希: %s\n", tx.Hash().Hex())
	waitForTransaction(client, tx, "物流管理方记录食品物流")
	return tx.Hash().Hex()
}

// 物流管理方记录食品物流信息
func addLogisticsProviderff(instance *bindings.Foodtrace, client *ethclient.Client, privateKey string, address common.Address) (hase string) {
	// 创建交易选项
	auth, err := createTransactOpts(client, privateKey)
	if err != nil {
		log.Fatal(err)

	}

	// 调用注册食品方法
	tx, err := instance.AddLogisticsProvider(auth, address)
	if err != nil {
		log.Fatal("注册物流方:", err)
	}

	fmt.Printf("✅ 物流管理方记录食品物流信息! 交易哈希: %s\n", tx.Hash().Hex())
	waitForTransaction(client, tx, "注册物流方")
	return tx.Hash().Hex()
}

func createTransactOpts(client *ethclient.Client, privateKeyHex string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("私钥格式错误: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("创建交易者失败: %v", err)
	}

	// 设置合理的Gas限制
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(20000000000) // 20 Gwei

	return auth, nil
}
func waitForTransaction(client *ethclient.Client, tx *types.Transaction, operation string) {
	fmt.Printf("⏳ 等待 %s 交易确认...\n", operation)

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Printf("✅ %s 成功! 区块: %d, Gas使用: %d\n",
			operation, receipt.BlockNumber, receipt.GasUsed)
	} else {
		fmt.Printf("❌ %s 失败!\n", operation)
	}
}

// StringToBigInt 将字符串转换为 big.Int，支持多种格式
func StringToBigInt(s string) (*big.Int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, fmt.Errorf("空字符串")
	}

	// 处理十六进制
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
		bigInt := new(big.Int)
		_, success := bigInt.SetString(s, 16)
		if !success {
			return nil, fmt.Errorf("无效的十六进制数字: %s", s)
		}
		return bigInt, nil
	}

	// 先尝试普通整数
	if num, err := strconv.ParseInt(s, 10, 64); err == nil {
		return big.NewInt(num), nil
	}

	// 尝试大整数
	bigInt := new(big.Int)
	_, success := bigInt.SetString(s, 10)
	if !success {
		return nil, fmt.Errorf("无效的数字格式: %s", s)
	}

	return bigInt, nil
}

// GetAddressFromPrivateKey 从私钥十六进制字符串获取以太坊地址
func GetAddressFromPrivateKey(privateKeyHex string) (string, error) {
	// 移除可能的0x前缀
	if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// 解码私钥
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("无效的私钥: %v", err)
	}

	// 转换为ECDSA私钥
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("无效的ECDSA私钥: %v", err)
	}

	// 获取公钥并生成地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("公钥转换失败")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.Hex(), nil
}
