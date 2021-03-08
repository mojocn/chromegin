## 1. 编译golang代码
FROM centos:7.6.1810

# 安装中文字体和chrome
RUN yum install -y wget git&& \
    yum install -y wqy-microhei-fonts wqy-zenhei-fonts && \
    wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
    yum install -y ./google-chrome-stable_current_*.rpm && \
    google-chrome --version && \
    rm -rf *.rpm


ENV GOPROXY=https://goproxy.io  PATH="${PATH}:/usr/local/go/bin"
# 定义使用的Golang 版本
ARG GO_VERSION=1.16

# 安装 golang 1.6
RUN wget "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf *.tar.gz && \
    go version && go env;


ARG GO_DIR=.

ARG BUILD_DIR=/gobin

WORKDIR $BUILD_DIR
# git describe --tags --always

#COPY ksbastion-sshark .
COPY $GO_DIR/go.mod .
COPY $GO_DIR/go.sum .
RUN go mod download

COPY $GO_DIR .

RUN export GITHASH=$(git rev-parse --short HEAD) && \
    export BUILDAT=$(date) && \
    go build -ldflags "-w -s -X 'main.BuildAt=$BUILDAT' -X 'main.GitHash=$GITHASH'"


VOLUME /pic
EXPOSE 6666

#remove golang files
RUN rm -rf /usr/local/go /root/go /root/.cache /root/.config

CMD ["/gobin/chromegin"]