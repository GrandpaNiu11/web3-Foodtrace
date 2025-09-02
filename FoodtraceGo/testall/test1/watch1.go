package main

import (
	"FoodtraceGo/bindings"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 1. 连接到以太坊节点（本地 Ganache / Hardhat / 测试网）
	client, err := ethclient.Dial("ws://127.0.0.1:7565")
	if err != nil {
		log.Fatal("连接节点失败:", err)
	}
	defer client.Close()

	fmt.Println("成功连接到以太坊节点")

	// 2. 查看区块高度
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("获取区块高度失败:", err)
	}
	fmt.Printf("当前区块高度: %d\n", blockNumber)

	// 获取链 ID
	chainID, err := client.ChainID(context.Background()) // 使用 context.Background() 或自定义 context
	if err != nil {
		log.Fatal(err)
	}

	// 打印链 ID
	fmt.Printf("Chain ID: %d\n", chainID)

	// 3. 创建合约实例

	contractAddress := common.HexToAddress("0x2C86fBFD95F55550E7516bD917BbEa0f74304D21")

	contractInstance, err := bindings.NewFoodtrace(contractAddress, client)

	getAllHistoricalEvents(contractInstance, client)
}
func getAllHistoricalEvents(instance *bindings.Foodtrace, client *ethclient.Client) {
	fmt.Println("📜 查询历史事件...")

	// 获取当前区块
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 查询从0到当前区块的所有事件
	startBlock := uint64(0)
	if currentBlock > 1000 {
		startBlock = currentBlock - 10
	}

	// 食品注册事件
	foodEvents, err := instance.FilterFoodRegistered(&bind.FilterOpts{
		Start: startBlock,
		End:   &currentBlock,
	})
	if err != nil {
		log.Printf("查询历史食品事件失败: %v", err)
	} else {
		fmt.Printf("\n🍎 历史食品注册事件 (%d 个):\n", countEvents(foodEvents))
		for foodEvents.Next() {
			event := foodEvents.Event
			fmt.Printf("- ID: %d, 名称: %s, 生产商: %s\n",
				event.Id, event.Name, event.Producer.Hex())
		}
	}

	// 物流记录事件
	logisticsEvents, err := instance.FilterLogisticsRecorded(&bind.FilterOpts{
		Start: startBlock,
		End:   &currentBlock,
	})
	if err != nil {
		log.Printf("查询历史物流事件失败: %v", err)
	} else {
		fmt.Printf("\n🚚 历史物流记录事件 (%d 个):\n", countEvents(logisticsEvents))
		for logisticsEvents.Next() {
			event := logisticsEvents.Event
			fmt.Printf("- 食品ID: %d, 物流信息: %s\n", event.Id, event.LogisticsRecord)
		}
	}
}

func countEvents(iterator interface{ Next() bool }) int {
	count := 0
	for iterator.Next() {
		count++
	}
	return count
}
