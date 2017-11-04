# bing

将 Bing 每日壁纸上传到七牛云。

## 动机

* [Pipe](https://github.com/b3log/pipe) 博客平台有个特性是随机为文章配图，所以需要一些漂亮的图片
* Bing 壁纸原始地址会过期，为了尽量保证图片的可访问性，所以下载原图后上传到七牛云

## 使用方式

1. 命令行参数 -bucket, -ak, -sk
2. 编译出可执行二进制后配置定时任务每天执行。定时任务最好每天执行 1 次以上，主要是为了降低出错概率

## 访问路径

上传七牛云的文件路径是 `bing/yyyyMMdd.jpg`，客户端可以按此规则生成随机路径进行访问。

## 一些细节

* 壁纸分辨率是 1920*1080，按每张大小 500K 计算一年大概需要 180M 存储空间
* 可以直接访问我上传好的壁纸（带 CDN）：`https://img.hacpai.com/bing/yyyyMMdd.jpg`，起始日期是 20171104，即 https://img.hacpai.com/bing/20171104.jpg

## 鸣谢

* 微软 Bing 搜索以及各图片作者
* [GoRequest](https://github.com/parnurzeal/gorequest)：简单好用的 golang HTTP 客户端
