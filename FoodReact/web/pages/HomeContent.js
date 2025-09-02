import React, {useEffect, useState} from "react";

import styles from '../index.module.scss';
import {Table} from "@douyinfe/semi-ui";
import axios from "axios";

const Component = () => {
    const [foodItemsData, setFoodItemsData] = useState([]);
    useEffect(() => {
        console.log('Component mounted');
        getFoodItems()

    }, []);
  async  function getFoodItems(){
      const res  =  await axios.get('http://127.0.0.1:8080/selectFoodItems')
      setFoodItemsData(res.data.data)
      console.log(res.data.data)

    }

    const columns = [
        {
            title: '食品名称',
            dataIndex: 'FoodName',
        },
        {
            title: '出仓日期',
            dataIndex: 'OutboundDate',
        },
        {
            title: '管理员地址',
            dataIndex: 'BlockchainAddress',
        },
        {
            title: '交易哈希',
            dataIndex: 'TransactionHash',
        },
    ];


    return (
        <div className={styles.frame}>
            <div className={styles.customers}>
                <p className={styles.item}>基于FISCO BCOS的食品溯源平台-信息看板;</p>
                 <Table columns={columns} dataSource={foodItemsData} pagination={false} />
            </div>
        </div>
    );
}

export default Component;
