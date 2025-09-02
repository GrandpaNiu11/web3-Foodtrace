'use client'
import styles from './login.module.scss';
import { Card, Form, Input, Button, Typography, Toast } from '@douyinfe/semi-ui';
import { useState } from "react";
import axios from "axios";

export default function Login() {
    const [from, setFrom] = useState({});
    const [loading, setLoading] = useState(false);

    async function login() {
        if (!from.username || !from.password) {
            Toast.warning('请输入用户名和密码');
            return;
        }

        setLoading(true);
        try {
            const response = await axios.post('http://127.0.0.1:8080/userLogin', {
                username: from.username,
                password: from.password
            });

            if (response.data.message == "该用户不存在") {
                Toast.error('账号或密码错误');
            } else {
                document.cookie = `address=${response.data.address}; path=/; max-age=86400; SameSite=Lax`;
                Toast.success('登录成功 🎉');
                setTimeout(() => {
                    window.location.href = '/';
                }, 1000);
            }
        } catch (error) {
            Toast.error('登录失败，请重试');
            console.error('Login error:', error);
        } finally {
            setLoading(false);
        }
    }

    return (
        <div style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            minHeight: '94vh',
            background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
            padding: '20px'
        }}>
            <Card
                style={{
                    maxWidth: 450,
                    width: '100%',
                    padding: '40px 30px',
                    borderRadius: '16px',
                    boxShadow: '0 20px 40px rgba(0, 0, 0, 0.1)',
                    border: 'none',
                    backdropFilter: 'blur(10px)',
                    backgroundColor: 'rgba(255, 255, 255, 0.95)'
                }}
                bodyStyle={{
                    padding: 0
                }}
            >
                {/* Logo 或标题区域 */}
                <div style={{ textAlign: 'center', marginBottom: '40px' }}>
                    <Typography.Title
                        heading={2}
                        style={{
                            marginBottom: '10px',
                            color: '#333',
                            fontWeight: 600
                        }}
                    >
                        欢迎回来
                    </Typography.Title>
                    <Typography.Text type="tertiary" style={{ fontSize: '16px' }}>
                        请输入您的账号密码登录系统
                    </Typography.Text>
                </div>

                <Form layout="vertical" onValueChange={values => setFrom(values)}>
                    <Form.Input
                        field="username"
                        label="用户名"
                        placeholder="请输入用户名"
                        style={{
                            marginBottom: '24px',
                            borderRadius: '8px'
                        }}
                        size="large"
                        rules={[{ required: true, message: '请输入用户名' }]}
                    />

                    <Form.Input
                        field="password"
                        label="密码"
                        type="password"
                        placeholder="请输入密码"
                        style={{
                            marginBottom: '30px',
                            borderRadius: '8px'
                        }}
                        size="large"
                        rules={[{ required: true, message: '请输入密码' }]}
                    />

                    <Button
                        theme="solid"
                        type="primary"
                        htmlType="button"
                        block
                        onClick={login}
                        loading={loading}
                        style={{
                            height: '46px',
                            borderRadius: '8px',
                            fontSize: '16px',
                            fontWeight: 500,
                            background: 'linear-gradient(135deg, #667eea, #764ba2)',
                            border: 'none',
                            marginBottom: '20px'
                        }}
                    >
                        {loading ? '登录中...' : '登录账户'}
                    </Button>
                </Form>

                {/* 底部提示 */}
                <div style={{
                    textAlign: 'center',
                    marginTop: '20px',
                    paddingTop: '20px',
                    borderTop: '1px solid #f0f0f0'
                }}>
                    <Typography.Text type="tertiary" style={{ fontSize: '14px' }}>
                        食品溯源系统 © 2024
                    </Typography.Text>
                </div>
            </Card>
        </div>
    )
}