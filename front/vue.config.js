module.exports = {
    devServer: {
        port: 8080, // 端口号
        https: false, // https:{type:Boolean}
        open: true, // 配置自动启动浏览器
        hotOnly: true, // 热更新
        proxy: {
            // 配置多个跨域
            '/api/': {
                // target: 'http://dqynb.ml:44632/notes',//跨域接口的地址
                target: "http://127.0.0.1:44632/notes",
                changeOrigin: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        },
    }
}