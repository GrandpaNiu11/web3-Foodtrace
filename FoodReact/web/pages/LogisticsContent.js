import React, { useEffect, useRef, useState } from 'react'; // 只添加 useState
import "@douyinfe/semi-ui/dist/css/semi.css";
import styles from "../index.module.scss";
import {Card, Form, Row, Toast} from "@douyinfe/semi-ui";
import Section from "@douyinfe/semi-ui/lib/es/form/section";
import axios from "axios";

const LogisticsContent = () => {
     const [wlfrom, setwlFrom] = useState({});
    const [wlzcfrom, setwlzcFrom] = useState({});
   async  function selectFoodId(){
         console.log("1213123",wlfrom)
         const response = await axios.get(`http://127.0.0.1:8080/getFoodInfo?foodid=${wlfrom.foodID}`, );
       const jsonData = response.data.data;

       // 格式化键值对
       const message = Object.entries(jsonData)
           .map(([key, value]) => `${key}: ${value}`)
           .join('\n');
       Toast.success(`该食品id为:`+message)
     }

     async function zcwlzcfrom(){
         console.log(document.cookie);
         console.log("caewcawcwa",wlzcfrom)
         const response = await axios.post('http://127.0.0.1:8080/updateLogistice', {
          "foodId":wlzcfrom.foodID,
          "logisticsInfo": wlzcfrom.logisticsInfo
         }, {
             withCredentials: true,
         });
         Toast.success(`注册成功交易hssh为:`+response.data.hash)

   }



    return (
        <>
            <div className={styles.frame}>
                <div className={styles.customers}>
                    <p className={styles.item}>基于FISCO BCOS的食品溯源平台-物流界面</p>

                    <div className={styles.content_all}>
                        <div className={styles.content_left}>
                            <Card theme="solid"
                                  style={{ backgroundColor: '#eff1f7', color: 'white', minWidth: 500, minHeight: 300 }}>
                                <Form

                                    style={{ padding: 10, width: '100%' }}
                                    onValueChange={values=>setwlFrom(values)}
                                >
                                    <Section text={'物流查询'}>
                                        <Row style={{ marginTop: 20 }}>

                                            <Form.Input
                                                field="foodID"
                                                label="食品id"
                                                placeholder={'请输入食品id'}
                                                style={{ width: '80%'}}
                                                trigger='blur'
                                            />
                                        </Row>
                                    </Section>
                                </Form>
                                <span  className={styles.queryLink} onClick={() => { selectFoodId()}} >查询此id对应的链上信息</span>
                            </Card>

                            <Card theme="solid"
                                  style={{ backgroundColor: '#eff1f7', color: 'white', minWidth: 500, minHeight: 300 ,marginTop: 20}}>
                                <Form

                                    style={{ padding: 10, width: '100%' }}
                                    onValueChange={values=>setwlzcFrom(values)}
                                >
                                    <Section text={'物流注册'}>
                                        <Row style={{ marginTop: 20 }}>

                                            <Form.Input
                                                field="foodID"
                                                label="食品id"
                                                placeholder={'请输入食品id'}
                                                style={{ width: '80%'}}
                                                trigger='blur'
                                            />
                                        </Row>

                                        <Row style={{ marginTop: 20 }}>

                                            <Form.Input
                                                field="logisticsInfo"
                                                label="物流信息"
                                                placeholder={'物流追踪'}
                                                style={{ width: '80%'}}
                                                trigger='blur'
                                            />
                                        </Row>
                                    </Section>
                                </Form>
                                <span  className={styles.queryLink} onClick={() => { zcwlzcfrom()}} >注册此id对应的物流信息</span>

                            </Card>
                        </div>

                        <div className={styles.content_right}    >


                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default LogisticsContent;