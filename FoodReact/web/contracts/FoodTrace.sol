// SPDX-License-Identifier: MIT
pragma solidity ^0.6.10;
// 开启 ABIEncoderV2 支持
pragma experimental ABIEncoderV2;

contract FoodTrace {
    struct FoodItem {
        uint id;  //食品id
        string name;  //食品名
        address producer;  //生产者
        string details;  //细节
        string[] logisticsRecords; // 物流记录
        bool logisticsRecorded; // 记录物流信息的状态
    }

    address public owner; // 合约的拥有者

    uint public logisticsRecordedCount; // 记录已注册物流信息的食品数
    uint public foodItemCount; // 食品项计数

    mapping(uint => FoodItem) public foodItems; // 食品项存储
    mapping(address => bool) public logisticsProviders; // 存储注册的物流方地址

    // 事件定义
    event FoodRegistered(uint id, string name, address producer);
    event LogisticsUpdated(uint id, string logisticsInfo);
    event LogisticsProviderRegistered(address provider);

    // 构造函数，设置合约的拥有者
    constructor() public {
        owner = msg.sender; // 设定合约创建者为拥有者
    }

    // 仅允许拥有者执行某些操作
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // 注册物流方
    function registerLogisticsProvider(address _provider) public onlyOwner {
        logisticsProviders[_provider] = true; // 注册物流方
        emit LogisticsProviderRegistered(_provider); // 触发事件
    }

    // 食品生产方注册食品
    function registerFood(string memory _name, string memory _details) public onlyOwner{
        foodItemCount++;
        foodItems[foodItemCount] = FoodItem({
            id: foodItemCount,
            name: _name,
            producer: msg.sender,
            details: _details,
            logisticsRecords: new string[](0), // 初始化物流记录为空数组
            logisticsRecorded: false // 默认未记录物流信息
        });

        emit FoodRegistered(foodItemCount, _name, msg.sender);
    }

    // 物流管理方记录食品物流信息
    function updateLogistics(uint _foodId, string memory _logisticsInfo) public {
        require(foodItems[_foodId].id != 0, "Food item does not exist"); // 确保食品存在
        require(logisticsProviders[msg.sender], "Not a registered logistics provider"); // 验证调用者是否是注册的物流方
        require(_foodId <= foodItemCount, "Food ID exceeds the current food item count");//确保输入的ID不超过当前的foodItemCount

        logisticsRecordedCount++;
        foodItems[_foodId].logisticsRecords.push(_logisticsInfo); // 添加物流信息
        foodItems[_foodId].logisticsRecorded = true; // 设置为已记录物流信息

        emit LogisticsUpdated(_foodId, _logisticsInfo);
    }

    // 用户查看食品记录
    function getFoodInfo(uint _foodId) public view returns (string memory, address, string memory, string[] memory, bool) {
        require(foodItems[_foodId].id != 0, "Food item does not exist"); // 确保食品存在
        require(_foodId <= foodItemCount, "Food ID exceeds the current food item count");//确保输入的ID不超过当前的foodItemCount
        FoodItem memory foodItem = foodItems[_foodId];
        return (foodItem.name, foodItem.producer, foodItem.details, foodItem.logisticsRecords, foodItem.logisticsRecorded);//返回食品信息和物流状态
    }
}
