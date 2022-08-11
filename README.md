# Microsoft-OAuth-Test

未调用官方SDK，手动组合请求

回调接收，打印最终的access_token

流程与oauth2.0 code模式调用流程完全一致

client_id、client_secret 配置位于configs.go文件中，需要自行补充

已有编译好的exe，可直接运行

浏览器访问localhost:5001/redirect即可跳转授权界面

程序会在控制台打印最终结果