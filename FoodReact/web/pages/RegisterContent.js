import React from 'react';

import "@douyinfe/semi-ui/dist/css/semi.css";
import styles from "../index.module.scss";
import {Button, Card, Col, Form, Input, Row, Spin, Toast} from "@douyinfe/semi-ui";
import Section from "@douyinfe/semi-ui/lib/es/form/section";
import axios from "axios";

const RegisterContent = () => {
     const [from, setFrom] = React.useState({});
    const [loading, setLoading] = React.useState(false);

    async function djsl() {
        setLoading(true);
        const response = await axios.post('http://127.0.0.1:8080/registerFood', {
            name: from.name,
            price: from.price,
            store: from.store,
            date: from.date,
            code: from.code
        });
        console.log(response)
        if (response.status === 200){
            Toast.success(`信息链上记录成功交易,hash为:`+response.data.hash)
            setFrom({});
        }else {
            alert("添加失败")
        }
        setLoading(false);
    }
    return (
        <>
            <Spin spinning={loading} tip="数据提交中...">
            <div className={styles.frame}>
                <div className={styles.customers}>
                    <p className={styles.item}>基于FISCO BCOS的食品溯源平台-登记界面</p>
                    <Card theme="solid"
                          style={{ backgroundColor: '#bab9b9', color: 'white', minWidth: 900 ,minHeight: 600   }}>

                        <Form

                            style={{ padding: 10, width: '100%' }}
                            onValueChange={values=>setFrom(values)}
                        >
                            <Section text={'基本信息'}>
                                <Row>

                                <Form.Input
                                    field="name"
                                    label="名称（Input）"
                                    placeholder={'请输入食品名称'}
                                    style={{ width: '80%' }}
                                    trigger='blur'
                                />
                                    <Form.Input
                                        field="price"
                                        label="价格（Input）"
                                        placeholder={'请输入价格'}
                                        style={{ width: '80%' }}
                                        trigger='blur'
                                    />
                                </Row>
                            </Section>

                            <Section text={'出仓信息'}>
                                <Form.Input
                                    field="store"
                                    label="发货仓库"
                                    placeholder={'请输入发货仓库'}
                                    style={{ width: '80%' }}
                                    trigger='blur'
                                />
                                <Form.Input
                                    field="date"
                                    label="出仓日期"
                                    placeholder={'出仓日期'}
                                    style={{ width: '80%' }}
                                    trigger='blur'
                                />
                                <Form.Input
                                    field="code"
                                    label="出仓编码"
                                    placeholder={'请输入出仓编码'}
                                    style={{ width: '80%' }}
                                    trigger='blur'
                                />

                            </Section>
                            <Button theme="solid" type="primary" htmlType="submit" block onClick={() => djsl()}>
                                登记上链
                            </Button>
                        </Form>
                         </Card>


                </div>
            </div>
            </Spin>
        </>
    );
};

export default RegisterContent;
