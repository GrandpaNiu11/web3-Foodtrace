package main

import (
	"FoodtraceGo/bindings"
	"fmt"
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

	// 3. 创建合约实例

	contractAddress := common.HexToAddress("0x2C86fBFD95F55550E7516bD917BbEa0f74304D21")

	contractInstance, err := bindings.NewFoodtrace(contractAddress, client)
	watchLogisticsRecordedEvents(contractInstance)

}
func watchLogisticsRecordedEvents(instance *bindings.Foodtrace) {
	eventChan := make(chan *bindings.FoodtraceLogisticsRecorded)

	sub, err := instance.WatchLogisticsRecorded(nil, eventChan)
	if err != nil {
		log.Fatal("创建物流事件订阅失败:", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("👂 开始监听物流记录事件...")

	for {
		select {
		case event := <-eventChan:
			fmt.Printf("\n🚚 新物流记录事件!\n")
			fmt.Printf("食品ID: %d\n", event.Id)
			fmt.Printf("物流信息: %s\n", event.LogisticsRecord)
			fmt.Printf("区块: %d\n", event.Raw.BlockNumber)

		case err := <-sub.Err():
			log.Printf("物流事件监听错误: %v", err)
			return
		}
	}
}
