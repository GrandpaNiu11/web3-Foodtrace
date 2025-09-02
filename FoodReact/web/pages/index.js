import React, {useEffect, useState} from "react";
import {Nav, Avatar, Descriptions, Toast, Layout, Dropdown, Notification} from '@douyinfe/semi-ui';
import {IconHelpCircle, IconBell, IconUserCircle,} from '@douyinfe/semi-icons';
import styles from '../index.module.scss';
import HomeContent from "./HomeContent"; // 主页组件
import RegisterContent from "./RegisterContent"; // 食品登记组件
import LogisticsContent from "./LogisticsContent"; // 物流溯源组件
import SettingsContent from "./SettingContent";
import Footer from "@douyinfe/semi-ui/lib/es/datePicker/footer";
import {useRouter} from "next/router";
import withAuth from "../lib/withAuth";
import axios from "axios"; // 设置组件


const Component = () => {

    const [selectedKey, setSelectedKey] = useState('Home'); // 用于跟踪当前选中的导航项
    const [foodItemCount, setFoodItemCount] = useState();
    const [logisticsRecordedCount, setLogisticsRecordedCount] = useState('Home');
    const {Footer} = Layout;
    const handleMenuClick = (itemKey) => {
        setSelectedKey(itemKey); // 切换选中项
    };
    useEffect(() => {
        getFoodItemCount();
        getLogisticsRecordedCount()

    })
    function getFoodItems(){
        // 清空 address cookie
        document.cookie = 'address=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/';

        Toast.success('退出登录成功');

        // 跳转到首页
        window.location.href = '/';


    }





    // 获取注册食品数
    async function getFoodItemCount(){
        const response = await axios.get('http://127.0.0.1:8080/foodItemCount', {
        });
        console.log(response)
        setFoodItemCount(response.data.data.count)
    }
    //获取链上交易数
   //获取溯源交易数

    async function getLogisticsRecordedCount(){
        const response = await axios.get('http://127.0.0.1:8080/logisticsRecordedCount', {
        });
        console.log(response)
        setLogisticsRecordedCount(response.data.data.count)
    }

    // 根据选中的导航项条件渲染内容
    const renderContent = () => {
        switch (selectedKey) {
            case 'Home':
                return <HomeContent/>;
            case 'Register':
                return <RegisterContent/>;
            case 'Logistics':
                return <LogisticsContent/>;
            case 'Settings':
                return <SettingsContent/>;
            default:
                return <HomeContent/>;
        }
    };


    return (
        <div className={styles.frame}>
            <Nav
                mode="horizontal"
                header={{
                    logo: <img src="../static/logo_transparent.png" style={{height: "70px", width: "70px"}}
                               alt="Logo"/>,
                    text: "Food Trace",
                }}
                footer={
                    <div className={styles.dIv}>
                        <IconHelpCircle size="large" className={styles.semiIconsFeishuLogo}/>
                        <IconBell size="large" className={styles.semiIconsFeishuLogo}/>
                        <Dropdown
                            render={
                                <Dropdown.Menu>
                                    <Dropdown.Item
                                        style={{color: "rgba(var(--semi-grey-3), 1)"}}>区块链地址</Dropdown.Item>
                                    <Dropdown.Item
                                        onClick={() => {getFoodItems()}}
                                        style={{color: "rgba(var(--semi-grey-3), 1)"}} >退出登录</Dropdown.Item>
                                </Dropdown.Menu>
                            }
                        >
                            <Avatar
                                size="small"
                                src="https://sf6-cdn-tos.douyinstatic.com/obj/eden-cn/ptlz_zlp/ljhwZthlaukjlkulzlp/root-web-sites/avatarDemo.jpeg"
                                color="blue"
                                className={styles.avatar}
                            >
                                示例
                            </Avatar>
                        </Dropdown>
                    </div>
                }
                className={styles.nav}
            >
                <Nav.Item itemKey="Home" text="信息看板" onClick={() => handleMenuClick('Home')}/>
                <Nav.Item itemKey="Register" text="食品登记" onClick={() => handleMenuClick('Register')}/>
                <Nav.Item itemKey="Logistics" text="物流溯源" onClick={() => handleMenuClick('Logistics')}/>
                <Nav.Item itemKey="Settings" text="资料设置" onClick={() => handleMenuClick('Settings')}/>
            </Nav>

            <div className={styles.content}>
                <div className={styles.frame18637}>
                    <div className={styles.frame1321314182}>
                        <div className={styles.workitemIcon}>
                            <IconUserCircle className={styles.semiIconsUserCircle}/>
                        </div>
                        <Descriptions
                            data={[{key: "注册食品数", value: foodItemCount}]}
                            row={true}
                            className={styles.descriptions}
                        />
                    </div>

                    <div className={styles.frame1321314182}>
                        <div className={styles.buttonOnlyIconSecond}>
                            <IconUserCircle className={styles.semiIconsUserCircle}/>
                        </div>
                        <Descriptions
                            data={[{key: "链上交易数", value: 99}]}
                            row={true}
                            className={styles.descriptions}
                        />
                    </div>

                    <div className={styles.frame1321314182}>
                        <div className={styles.buttonOnlyIconSecond2}>
                            <IconUserCircle className={styles.semiIconsUserCircle}/>
                        </div>
                        <Descriptions
                            data={[{key: "溯源登记数", value: logisticsRecordedCount}]}
                            row={true}
                            className={styles.descriptions}
                        />
                    </div>
                </div>

                <div>
                    {renderContent()} {/* 根据选中的项渲染相应的组件内容 */}
                </div>
            </div>
         
            <Footer>
                &copy; 2025 Food Trace @FISCO BCOS-BLCOKCHAIN. Version 1.0.0
            </Footer>
        </div>
    );
}

export default withAuth(Component);

