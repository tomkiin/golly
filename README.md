## 关于golly

golly是一款基于go+vue开发的前后端分离可平台化配置的爬虫框架，灵活、轻量

## 运行此项目

#### 数据库

配置api/model/db.go的配置信息，新建mysql数据库，导入golly.sql文件

#### 后端

```bash
cd golly
go run main.go
```

#### 前端

```bash
npm run serve
```

访问 http://localhost:8080 进入系统

## 使用教程

- 运行任务：点击任务管理，系统默认配置了一个任务，点击运行按钮即可启动，运行时可点击日志按钮查看获取的信息。
- 配置任务：输入需要爬取的任务名称、种子地址、源地址添加任务，点击配置按钮添加收集器规则，收集器解析规则需遵守[goquery](https://github.com/PuerkitoBio/goquery)的写法，收集器支持两种数据获取方式（attr和text），比如你想获取a标签的href信息，你需要选择attr选项，并在参数名称一栏中填写“href”即可；text选项则直接获取text文本。收集器获取的数据会作为下一个收集器的入口URL，请注意URL的拼接（通常是你填写的源地址+a标签的href信息）

## 设计思路

golly项目大致分为6大模块：

- api - 提供api控制
- app - 爬虫核心代码
- cache - 缓存模块
- cron - 异步任务
- errors - 自定义错误包
- web - 前端

爬虫架构设计参考了golang并发模式的流水线模式，收集器的设计参考了[colly](https://github.com/gocolly/colly)框架的设计，页面解析使用了[goquery](https://github.com/PuerkitoBio/goquery)

## 未来更新

- [ ] 任务报错处理
- [ ] cookie支持
- [ ] 反爬虫支持
- [ ] 存储器接口开发，支持mysql、mongo、文件等存储方式
- [ ] 配置线上化
- [ ] 分布式
- [ ] 未完待续