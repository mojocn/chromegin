# 基于golang chrome-dp 网页截图微服务

## 1. 4Dockerfile功能说明
1. 安装中文字体
2. 安装linux chrome浏览器
3. 安装golang 环境
3. chromedp + gin RESTful API 业务代码编译和运行

## 2. Chromedp RESTful API 接口说明
|  表头   | 表头  |
|  ----  | ----  |
| Method | GET |
| URL  | 127.0.0.1:6666/python/ss |
| u  | url-encode 之后的截图网址 |
| c  | url-encode 时间 格式 `2018-09-09 12:12:12` |
