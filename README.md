# appletMessagesServer
项目说明

    本项目小程序客服消息处理服务，用户在微信公众平台开启消息推送，并将cfg.json的推送参数配置到微信公众平台开启消息推送对应的设置后即可实现小程序推送消息的处理。
    
使用说明

    1.将本代码下载到本地，然后在cfg.json填入自己的参数。参数填入请参考cfg.example.json文件
    2.使用go build 编译改项目，然后使用项目中自带的shell脚本control来启动该服务，启动命令  ./control start ，更多control请知晓 ./control help
