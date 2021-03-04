# 基于golang chrome-dp 网页截图微服务

## 1. Dockerfile功能说明
1. 安装中文字体
2. 安装linux chrome浏览器
3. 安装golang 环境
3. chromedp + gin RESTful API 业务代码编译和运行

```dockerfile
FROM centos:7.6.1810

# 安装中文字体和chrome
RUN yum install -y wget && \
    yum install -y wqy-microhei-fonts wqy-zenhei-fonts && \
    wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
    yum install -y ./google-chrome-stable_current_*.rpm && \
    google-chrome --version && \
    rm -rf *.rpm

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.io GOPATH=/gopath PATH="${PATH}:/usr/local/go/bin"
# 定义使用的Golang 版本
ARG GO_VERSION=1.13.3

# 安装 golang 1.13.3
RUN wget "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf *.tar.gz && \
    go version && go env;


WORKDIR $GOPATH
COPY . chromegin

RUN cd chromegin && go build -o app;

EXPOSE 6666

# 保存图片网页图片截图
VOLUME /data

CMD ["chromegin/app"]
```

## 2. Docker 编译和运行
```bash
git clone https://github.com/mojocn/chromegin.git && cd chromegin
# 编译build image 名称位chromegin  docker run 挂在host主机/data/chrome_screen_shot 目录保存图片
docker build -t chromegin . && docker run -p 6666:6666 -v /data/chrome_screen_shot:/data --name chromegin chromegin 
```

从dockerhub上pull
`docker pull mojotvcn/chromegin`

## 3. Chromedp RESTful API 接口说明
|  表头   | 表头  |
|  ----  | ----  |
| Method | GET |
| URL  | 127.0.0.1:6666/python/ss |
| URL  | 127.0.0.1:6666/open/chromedp/screen/shot |
| u  | url-encode 之后的截图网址 |
| c  | url-encode 时间 格式 `2018-09-09 12:12:12` |


## 5 使用
```bash
curl --location --request POST 'http://localhost:6666/api' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sel":"",
    "url":"https://localhost:9527/#/jms/report/2021-02-01/2021-03-03",
    "wait":3
}'
```