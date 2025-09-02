//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract FoodItem {
    struct FoodItem {
        uint id;
        string name;
        address producer;
        string details;
        string[] logisticsRecords;
        bool logisticsRecorded;
    }
    //合约拥有者
    address public owner;

//    记录已注册物流信息的食品数
    uint public logisticsRecordedCount;
    //食品项目数
    uint public foodItemCount;

// 食品项存储
    mapping(uint => FoodItem) public foodItems;
//   存储注册的物流方地址
    mapping(address => bool) public logisticsProviders;

//    事件定义
    event FoodRegistered(uint id, string name, address producer);
    event LogisticsRecorded(uint id, string logisticsRecord);
    event LogisticsProviderAdded(address provider);


    constructor() {
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function.");
        _;
    }
//    注册物流方
    function addLogisticsProvider(address _provider) public onlyOwner {
        logisticsProviders[_provider] = true;
        emit LogisticsProviderAdded(_provider);
    }
    //食品生产方注册食品
    function registerFood(string memory _name, string memory _details) public onlyOwner {
        foodItemCount++;
        foodItems [foodItemCount] = FoodItem({
            id: foodItemCount,
            name: _name,
            producer: msg.sender,
            details: _details,
            logisticsRecords: new string[](0),
            logisticsRecorded: false
        });
        emit FoodRegistered(foodItemCount, _name, msg.sender);
    }

    //物流管理方记录食品物流信息
    function updateLogistice(uint _foodId, string memory _logisticsInfo) public {
        require(foodItems[_foodId].id != 0, "Logistics info has been recorded");
        //查看调用者是否是食品的创建者
        require(logisticsProviders[msg.sender], "Only the producer can update logistics info");
        require(_foodId <= foodItemCount, "Invalid food id");
        logisticsRecordedCount++;
        foodItems[_foodId].logisticsRecords.push(_logisticsInfo);
        foodItems[_foodId].logisticsRecorded = true;
        emit LogisticsRecorded(_foodId, _logisticsInfo);
    }
    //用户查看食品记录
    function getFoodInfo(uint _foodId) public view returns (FoodItem memory) {
        require(foodItems[_foodId].id != 0, "Invalid food id");
        require(_foodId <= foodItemCount, "Invalid food id");
        return foodItems[_foodId];
    }


}