/** @type {import('next').NextConfig} */
const nextConfig = {
    transpilePackages: [
        '@douyinfe/semi-ui',
        '@douyinfe/semi-icons',
        '@douyinfe/semi-illustrations',
        '@douyinfe/semi-foundation',
        'date-fns',
        'date-fns-tz'
    ],
    experimental: {
        esmExternals: 'loose'
    },
    webpack: (config) => {
        config.module.rules.push({
            test: /date-fns-tz/,
            type: 'javascript/auto'
        });
        return config;
    }
};

module.exports = nextConfig;
