// lib/withAuth.js
import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { Spin } from '@douyinfe/semi-ui';

const withAuth = (WrappedComponent) => {
    return (props) => {
        const router = useRouter();
        const [isLoading, setIsLoading] = useState(true);

        useEffect(() => {
            const checkAuth = () => {
                // 获取 token
                const token = getCookie('address');
                // 定义公开路径（不需要登录就能访问的页面）
                const publicPaths = ['/login', '/register', '/forgot-password'];
                const currentPath = router.pathname;

                if (!token && !publicPaths.includes(currentPath)) {
                    // 未登录且访问的是需要认证的页面，跳转到登录页
                    router.push('/login');
                } else if (token && publicPaths.includes(currentPath)) {
                    // 已登录但访问的是登录页，跳转到首页
                    router.push('/');
                } else {
                    setIsLoading(false);
                }
            };

            checkAuth();
        }, [router]);

        if (isLoading) {
            return (
                <div style={{
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                    height: '100vh'
                }}>
                    <Spin size="large" tip="权限验证中..." />
                </div>
            );
        }

        return <WrappedComponent {...props} />;
    };
};

// 获取 cookie 的辅助函数
export const getCookie = (name) => {
    if (typeof document === 'undefined') return null;
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
    return null;
};

export default withAuth;