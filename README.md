# Welcome to Revel First item
revel build myapp
revel run myapp


The directory structure of a generated Revel application:

    conf/             配置文件目录
        app.conf      app的主要配置
        routes        路由配置

    app/              应用资源
        init.go       初始化注册
        controllers/  控制器
        models/       站点模型
        jobs/         任务吧？
        services/     服务
        views/        模板目录

    messages/         Message files

    public/           静态文件资源
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help:
    可以在app/controller 加一个文件，另外写哪里的业务
    c.Response.Status=201 设置方法状态码
    引入了xorm go get github.com/go-xorm/xorm

# goRevel
    http://ace.jeka.by/
# 问题
    模板不好用，无法识别 url action的匹配