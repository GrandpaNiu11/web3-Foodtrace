import React, {useEffect, useState} from 'react';
import {
    Layout,
    Empty,
} from '@douyinfe/semi-ui';

import "@douyinfe/semi-ui/dist/css/semi.css";
import {
    IllustrationNoContent,
    IllustrationNoContentDark,
} from "@douyinfe/semi-illustrations";
const LogisticsContent = () => {
    const { Header, Footer, Sider, Content } = Layout;

    return (
        <>
        <Empty
            image={<IllustrationNoContent style={{ width: 450, height: 450 }} />}
            darkModeImage={<IllustrationNoContentDark style={{ width: 450, height: 450 }} />}
            description={'暂未开放此功能'}
        />

        </>
    );
};

export default LogisticsContent;
