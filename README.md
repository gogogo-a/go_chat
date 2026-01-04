# KamaChat

【代码随想录知识星球】项目分享-基于go+vue实现的聊天室+仿微信项目

---

## 📑 目录导航

- [项目概述](#项目概述)
- [功能特性](#功能特性)
- [安装与运行](#安装与运行)
- [项目结构详解](#项目结构详解)
  - [后端项目结构](#后端项目结构)
  - [前端项目结构](#前端项目结构)
- [核心技术架构](#核心技术架构)
  - [后端架构设计](#后端架构设计)
  - [前端架构设计](#前端架构设计)
- [项目学习路线](#项目学习路线)
  - [适合人群](#适合人群)
  - [学习顺序](#学习顺序)
  - [学习建议](#学习建议)

---

# 项目概述

## 简介
KamaChat 是一个前后端分离的即时通讯项目，具备后台管理、单聊群聊、联系人管理、多种消息（文本 / 文件 / 视频）处理、离线消息处理以及音视频通话等功能，旨在打造类似微信的聊天体验。

## 技术栈

### 后端技术
- **编程语言**：Go 1.20+
- **Web 框架**：Gin（HTTP 路由、中间件）
- **ORM 框架**：GORM（MySQL 数据库操作）
- **数据库**：MySQL 8.0+
- **缓存**：Redis 6.0+（GoRedis 客户端）
- **消息队列**：Kafka（可选，支持 channel 和 kafka 两种模式）
- **实时通信**：WebSocket（Gorilla WebSocket）
- **日志库**：Zap（高性能结构化日志）
- **加密**：AES（密码加密）、SSL/TLS（HTTPS）
- **短信服务**：阿里云短信服务

### 前端技术
- **框架**：Vue3（组合式 API）
- **路由**：Vue Router 4
- **状态管理**：Vuex 4
- **UI 组件**：Element Plus
- **实时通信**：WebSocket
- **音视频**：WebRTC
- **构建工具**：Webpack / Vue CLI
- **包管理**：Yarn / npm

## 核心亮点
1. ✅ **分层架构清晰**：MVC 模式，职责分明，易于维护
2. ✅ **WebSocket 长连接**：实时消息推送，用户体验佳
3. ✅ **离线消息处理**：消息持久化，上线自动拉取
4. ✅ **Redis 缓存优化**：减少数据库查询，提升性能
5. ✅ **Kafka 消息队列**：支持高并发场景（可选）
6. ✅ **WebRTC 音视频**：P2P 通话，支持 TURN 中继
7. ✅ **后台管理系统**：用户/群组管理、权限控制
8. ✅ **完整的用户系统**：注册、登录、短信验证


# 功能特性
1. 即时通讯功能
   + 单聊与群聊：支持一对一私密聊天和群组聊天，消息实时推送。
   + 联系人管理：可添加、删除、拉黑联系人，处理好友申请等。
   + 消息类型：支持文本、文件、音视频等多种类型消息的发送与接收。
   + 离线消息处理：确保用户离线时消息不丢失，上线后可正常接收。
2. 音视频通话：基于 WebRTC 实现 1 对 1 音视频通话，包括发起、拒绝、接收、挂断通话等功能。
3. 后台管理：具备后台管理界面，靓号用户可进行人员管控等维护操作。
4. 安全与验证：登录注册采用 SMS 短信验证方式，并支持 SSL 加密，保障用户信息安全。
5. 后台mysql数据库：使用 GORM 进行数据库操作，确保数据持久化存储。
6. 日志记录：使用 Zap 日志库记录系统运行日志，便于问题排查与性能监控。
7. 消息队列：使用 Kafka 处理消息队列，确保消息的高效传输与处理。
8. redis缓存：使用 GoRedis 进行缓存操作，提高系统性能。
9. WebSocket：使用 WebSocket 实现实时消息推送，保证消息的实时性。

# 快速开始（本地开发）

## 环境要求
- Go 1.20+
- Node.js 16+
- MySQL 8.0+
- Redis 6.0+

## 本地运行步骤

### 1. 克隆项目
```bash
git clone <repository-url>
cd KamaChat
```

### 2. 配置后端

#### 2.1 安装依赖
```bash
go mod tidy
```

#### 2.2 配置文件
```bash
cp configs/config.toml configs/config_local.toml
```

修改 `configs/config_local.toml`：
```toml
[mainConfig]
appName = "KamaChat"
host = "0.0.0.0"
port = 8000

[mysqlConfig]
host = "127.0.0.1"
port = 3306
user = "root"
password = "your_password"
databaseName = "kama_chat_server"

[redisConfig]
host = "127.0.0.1"
port = 6379
password = ""
db = 0

[kafkaConfig]
messageMode = "channel"  # 本地开发使用 channel 模式
hostPort = "127.0.0.1:9092"

[logConfig]
logPath = "./cmd/kama_chat_server/logs"

[staticSrcConfig]
staticAvatarPath = "./static/avatars"
staticFilePath = "./static/files"
```

#### 2.3 创建数据库
```bash
mysql -u root -p
CREATE DATABASE kama_chat_server CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXIT;
```

#### 2.4 启动 Redis
```bash
redis-server
```

#### 2.5 运行后端
```bash
cd cmd/kama_chat_server
go run main.go
```

后端将在 `https://127.0.0.1:8000` 启动（需配置 SSL 证书）。

---

### 3. 配置前端

#### 3.1 安装依赖
```bash
cd web/chat-server
yarn install
# 或
npm install
```

#### 3.2 修改后端 API 地址
编辑 `web/chat-server/vue.config.js`，确保代理配置正确：
```javascript
module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'https://127.0.0.1:8000',
        changeOrigin: true,
        secure: false
      }
    }
  }
}
```

#### 3.3 运行前端
```bash
yarn serve
# 或
npm run serve
```

前端将在 `http://localhost:8080` 启动。

---

### 4. 测试项目

1. 打开浏览器访问 `http://localhost:8080`
2. 注册两个测试账号
3. 使用两个浏览器窗口登录不同账号
4. 互相添加好友
5. 发送消息测试

---

## 常见问题

### Q1: SSL 证书问题
**问题：** 后端启动报错 "can't find certificate"

**解决：** 本地开发可以使用 mkcert 生成自签名证书
```bash
# 安装 mkcert
brew install mkcert  # macOS
# 或在 Windows 下载 mkcert.exe

# 生成证书
mkcert -install
mkcert 127.0.0.1 localhost

# 将生成的证书放到 pkg/ssl/ 目录
```

或者修改 `cmd/kama_chat_server/main.go` 使用 HTTP：
```go
// 改为
if err := https_server.GE.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
    zlog.Fatal("server running fault")
}
```

### Q2: 数据库连接失败
检查 MySQL 是否启动：
```bash
# macOS/Linux
sudo service mysql status

# Windows
net start mysql
```

### Q3: Redis 连接失败
检查 Redis 是否启动：
```bash
redis-cli ping
# 返回 PONG 表示正常
```

### Q4: 前端请求跨域
确保后端已开启 CORS（代码中已配置）。

---

# 安装与运行（云服务器部署）
此次安装运行为一键部署，即可在Ubuntu22.04的云服务器上部署上线，公网都可以访问。
在执行脚本代码之前，需要做一些前置准备。
![](docs/image/3.png)
![](docs/image/4.png)
把端口3306（mysql）, 6379（redis）, 443（前端访问）, 80（云服务器http访问）, 22（ssh）, 3478（turn服务器，用于音视频公网转发）, 8000（后端访问）等端口开放。
![](docs/image/5.png)
打开前端src/views/chat/contact/ContactChat.vue，找到ICE_CFG配置，更新对应的turn服务器的相关配置。turn服务器就是你的云服务器。
如果需要本地通信的话，就需要把iceServers删掉，让ICE_CFG置空。
```toml
[mainConfig]
appName = "your app name"
host = "0.0.0.0"
port = 8000

[mysqlConfig]
host = "127.0.0.1"
port = 3306
user = "root"
password = "123456"
databaseName = "your database name"

[redisConfig]
host = "127.0.0.1"
port = 6379
password = ""
db = 0

[authCodeConfig]
accessKeyID = "your accessKeyID in alibaba cloud"
accessKeySecret = "your accessKeySecret in alibaba cloud"
signName = "阿里云短信测试"
templateCode = "SMS_154950909"

[logConfig]
logPath = "your log path"

[kafkaConfig]
messageMode = "channel"# 消息模式 channel or kafka
hostPort = "127.0.0.1:9092" # "127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094" 多个kafka服务器
loginTopic = "login"
chatTopic = "chat_message"
logoutTopic = "logout"
partition = 0 # kafka partition
timeout = 1 # 单位秒

[staticSrcConfig]
staticAvatarPath = "./static/avatars"
staticFilePath = "./static/files"
```

你需要修改相应的后端配置文件中的内容。还需要先完成手机验证的功能，这篇需要看“后端开发”里的“手机验证”功能。

在这些都完成之后，就可以开始执行脚本代码了。

```bash
#!/bin/bash

# 更新系统软件包
sudo apt update && sudo apt upgrade -y

# 安装 MySQL
echo "Installing MySQL..."
sudo apt install mysql-server -y

# 配置 MySQL 安全
sudo mysql_secure_installation

# 启动并启用 MySQL 服务
sudo systemctl start mysql
sudo systemctl enable mysql

# 自动创建数据库
echo "Creating database 'kama_chat_server'..."
sudo mysql -u root -p <<EOF
CREATE DATABASE kama_chat_server;
EOF

# 安装 Redis
echo "Installing Redis..."
sudo apt install redis-server -y

# 配置 Redis
# sudo nano /etc/redis/redis.conf  # 修改 bind 127.0.0.1 改为 bind 0.0.0.0（如果需要外部访问）

# 启动并启用 Redis 服务
sudo systemctl restart redis
sudo systemctl enable redis

# 卸载旧版本 Node.js 和 npm，如果不是纯净版的Ubuntu的话
echo "Uninstalling previous versions of Node.js and npm..."
sudo apt remove --purge -y nodejs npm

# # 安装 Node.js 版本管理工具（nvm）
echo "Installing Node Version Manager (nvm)..."
rm -rf ~/.nvm
export NVM_NODE_MIRROR=https://npmmirror.com/mirrors/node/
export NVM_NPM_MIRROR=https://npmmirror.com/mirrors/npm/
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash



# 设置 NVM_DIR 环境变量（避免重复写入 ~/.bashrc）
if ! grep -q "export NVM_DIR=~/.nvm" ~/.bashrc; then
    echo 'export NVM_DIR="$HOME/.nvm"' >> ~/.bashrc
    echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm' >> ~/.bashrc
    echo '[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion' >> ~/.bashrc
fi


# 手动设置 NVM_DIR 并加载 nvm
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # 加载 nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # 加载 bash_completion

source ~/.bashrc

# 检查 nvm 是否存在
if ! command -v nvm &> /dev/null; then
    echo "nvm could not be found. Please ensure it is installed and added to your PATH."
    exit 1
fi

# 安装指定版本的 Node.js（例如 v16.x）
echo "Installing Node.js v16.x..."
nvm install 16
nvm use 16


# 加载环境变量
source ~/.bashrc

# 安装 Go
echo "Installing Go..."
wget https://mirrors.aliyun.com/golang/go1.20.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz

cp -r /usr/local/go/bin/* /usr/bin
# 设置 Go 环境变量
echo "Configuring Go environment..."

export PATH=$PATH:/usr/local/go/bin

 # 设置 Go 环境变量（避免重复写入 ~/.bashrc）
 if ! grep -q "export GOPATH=$HOME/go" ~/.bashrc; then
     echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
 fi
 source ~/.bashrc


# # 配置 Go 代理
echo "Configuring Go proxy..."
go env -w GOPROXY=https://goproxy.cn,direct

# 安装 Vue.js 开发环境
echo "Installing Vue.js development environment..."
sudo apt install npm -y



# 方案1：使用 npm 安装 Yarn
# sudo npm install -g yarn

# 方案2：使用cnpm 安装 Yarn
sudo npm install -g cnpm --registry=https://registry.npmjs.org
sudo cnpm install -g yarn

# 安装 Vue CLI
sudo cnpm install -g @vue/cli

# 重新安装项目依赖
cd ~/project/KamaChat/web/chat-server

yarn cache clean
rm -rf node_modules

yarn install # 会把package.json中所有依赖配置好的

#打包项目成dist，放到/var/www/html/，此时就可以通过云服务器的公网ip看到前端页面了
rm -rf /var/www/html/* 
rm -rf /root/project/KamaChat/web/chat-server/dist
yarn build
sudo cp -r /root/project/KamaChat/web/chat-server/dist/* /var/www/html # 改成自己的项目路径
sudo chmod -R 755 /var/www/html
sudo chown -R www-data:www-data /var/www/html

cd ~/project/KamaChat

# 安装 ssl 模块
echo "Installing ssl..."
sudo apt-get install openssl
sudo apt-get install libssl-dev

# # 创建根密钥，生成证书签名请求 (CSR)，创建根证书
openssl genrsa -out /etc/ssl/private/root.key 2048
openssl req -new -key /etc/ssl/private/root.key -out /etc/ssl/certs/root.csr
openssl x509 -req -in /etc/ssl/certs/root.csr -out /etc/ssl/certs/root.crt -signkey /etc/ssl/private/root.key -CAcreateserial -days 3650

# 生成服务器密钥，生成服务器证书签名请求 (CSR)，创建服务器证书扩展文件
openssl genrsa -out /etc/ssl/private/server.key 2048
openssl req -new -key /etc/ssl/private/server.key -out /etc/ssl/certs/server.csr
sudo nano v3.ext
# 内容如下
# authorityKeyIdentifier=keyid,issuer
# basicConstraints=CA:FALSE
# keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
# subjectAltName = @alt_names
# [alt_names]
# IP.1 = xxxxxxxxx # 你的云服务器地址

# # 使用根证书为服务器证书签名
openssl x509 -req -in /etc/ssl/certs/server.csr -CA /etc/ssl/certs/root.crt -CAkey /etc/ssl/private/root.key -CAcreateserial -out /etc/ssl/certs/server.crt -days 500 -sha256 -extfile v3.ext


# 打开Apache2配置文件
sudo nano /etc/apache2/sites-enabled/000-default.conf
# 添加如下内容
# <VirtualHost *:443>
#     ServerAdmin webmaster@localhost
#     DocumentRoot /var/www/html

#     SSLEngine on
#     SSLProxyEngine on

#     # 替换为您的自签名证书路径
#     SSLCertificateFile /etc/ssl/certs/server.crt
#     SSLCertificateKeyFile /etc/ssl/private/server.key

#     # 如果有中间证书，添加以下行
#     # SSLCertificateChainFile /path/to/your_intermediate.crt

#     ErrorLog ${APACHE_LOG_DIR}/error.log
#     CustomLog ${APACHE_LOG_DIR}/access.log combined

#     # 以下配置可选，用于启用 HTTP 到 HTTPS 重定向，也可以把这段添加到80端口那儿
#     <IfModule mod_rewrite.c>
#         RewriteEngine On
#         RewriteCond %{HTTPS} off
#         RewriteRule ^/?(.*) https://%{SERVER_NAME}%{REQUEST_URI} [R=301,L]
#     </IfModule>
# </VirtualHost>

# 启用ssl模块，启用ssl站点，重启服务
sudo a2enmod ssl
sudo a2ensite 000-default.conf
sudo systemctl restart apache2


# 配置turn服务器
echo "Installing coturn..."
sudo apt install coturn
sudo nano /etc/coturn/coturn.conf
# 配置以下参数
# listening-ip=0.0.0.0

# external-ip=xxxxx # 外部 IP 地址（替换为你的服务器公网 IP）

# listening-port=3478 # 监听端口

# user=username:password # 用户名和密码（替换为你的用户名和密码）

# tls-certificate=/etc/ssl/certs/server.crt # SSL 证书路径（如果需要加密通信）
# tls-private-key=/etc/ssl/private/server.key

sudo systemctl start coturn
sudo systemctl enable coturn


# 将后端打包部署
cd ~/project/KamaChat/cmd/kama_chat_server # 里面是main.go
go build -o kama_chat_backend main.go
sudo cp kama_chat_backend /usr/local/bin/

sudo nano /etc/systemd/system/kama_chat_backend.service
# 配置以下内容
# [Unit]
# Description=kama chat service
# After=network.target

# [Service]
# User=kama_chat  # 替换为你的用户名
# Group=kama_chat  # 替换为你的用户名
# WorkingDirectory=/root/project/KamaChat/cmd/kama_chat_server  # 替换为你的项目路径
# ExecStart=/usr/local/bin/kama_chat_backend  # 替换为你的可执行文件路径
# Restart=on-failure
# RestartSec=5

# [Install]
# WantedBy=multi-user.target

# 把后端服务起起来
sudo systemctl daemon-reload
sudo systemctl start kama_chat_backend
sudo systemctl enable kama_chat_backend

# 输出完成信息
echo "Deployment complete!"
```


在Ubuntu22.04云服务器上执行该脚本，它就会自动部署相关的依赖，并把go后端和vue前端部署到对应的位置，之后的访问可以通过https://xxxxx:443去访问。如果在前端访问后端的时候报错“NetWork error”时，可能后端还没部署好，可以重启一下。

# 项目结构详解

## 后端项目结构

KamaChat 后端采用标准的 Go 项目分层架构，遵循 MVC 设计模式和依赖注入原则。

```
KamaChat/
├── api/                                    # API 接口层（Controller 层）
│   └── v1/                                 # API v1 版本
│       ├── controller.go                   # 公共响应处理函数
│       ├── chatroom_controller.go          # 聊天室相关接口（在线用户查询）
│       ├── group_info_controller.go        # 群组管理接口（创建、解散、更新群组等）
│       ├── message_controller.go           # 消息相关接口（获取消息列表、文件上传等）
│       ├── session_controller.go           # 会话管理接口（打开、删除会话等）
│       ├── user_contact_controller.go      # 联系人管理接口（添加、删除、拉黑联系人等）
│       ├── user_info_controller.go         # 用户信息接口（登录、注册、更新用户信息等）
│       └── ws_controller.go                # WebSocket 连接处理
│
├── cmd/                                    # 应用程序入口
│   └── kama_chat_server/
│       ├── main.go                         # 主程序入口，初始化服务并启动
│       ├── kama_chat_backend               # 编译后的可执行文件
│       └── logs/                           # 日志文件目录
│
├── internal/                               # 内部私有代码（不可被外部导入）
│   ├── config/                             # 配置管理
│   │   └── config.go                       # 配置文件解析（MySQL、Redis、Kafka等）
│   │
│   ├── dao/                                # Data Access Object（数据访问层）
│   │   └── gorm.go                         # GORM 数据库连接初始化
│   │
│   ├── dto/                                # Data Transfer Object（数据传输对象）
│   │   ├── request/                        # 请求参数结构体（30+ 个）
│   │   │   ├── login_request.go            # 登录请求
│   │   │   ├── register_request.go         # 注册请求
│   │   │   ├── create_group_request.go     # 创建群组请求
│   │   │   ├── chat_message_request.go     # 聊天消息请求
│   │   │   └── ......                      # 其他业务请求结构体
│   │   │
│   │   └── respond/                        # 响应数据结构体（20+ 个）
│   │       ├── login_respond.go            # 登录响应
│   │       ├── get_userinfo_respond.go     # 获取用户信息响应
│   │       ├── get_message_list_respond.go # 获取消息列表响应
│   │       └── ......                      # 其他业务响应结构体
│   │
│   ├── https_server/                       # HTTPS 服务器配置
│   │   └── https_server.go                 # Gin 路由初始化、CORS 配置、路由注册
│   │
│   ├── model/                              # 数据模型层（对应数据库表）
│   │   ├── user_info.go                    # 用户信息表模型
│   │   ├── group_info.go                   # 群组信息表模型
│   │   ├── message.go                      # 消息表模型
│   │   ├── session.go                      # 会话表模型
│   │   ├── user_contact.go                 # 用户联系人表模型
│   │   └── contact_apply.go                # 好友申请表模型
│   │
│   └── service/                            # 业务逻辑层
│       ├── chat/                           # WebSocket 聊天服务
│       │   ├── server.go                   # 聊天服务器（管理在线客户端、消息转发）
│       │   ├── client.go                   # WebSocket 客户端连接封装
│       │   └── kafka_server.go             # Kafka 模式的聊天服务器
│       │
│       ├── gorm/                           # 数据库业务逻辑
│       │   ├── user_info_service.go        # 用户相关业务（登录、注册、查询等）
│       │   ├── group_info_service.go       # 群组相关业务
│       │   ├── message_service.go          # 消息相关业务
│       │   ├── session_service.go          # 会话相关业务
│       │   ├── user_contact_service.go     # 联系人相关业务
│       │   └── chatroom_service.go         # 聊天室相关业务
│       │
│       ├── kafka/                          # Kafka 消息队列服务
│       │   └── kafka_service.go            # Kafka 生产者/消费者实现
│       │
│       ├── redis/                          # Redis 缓存服务
│       │   └── redis_service.go            # Redis 连接池、缓存操作
│       │
│       ├── sms/                            # 短信验证服务
│       │   └── auth_code_service.go        # 阿里云短信验证码发送
│       │
│       └── aes/                            # 加密服务
│           └── aes.go                      # AES 加密/解密工具
│
├── pkg/                                    # 公共可复用代码（可被外部导入）
│   ├── constants/                          # 常量定义
│   │   └── constants.go                    # 全局常量（通道大小、超时时间等）
│   │
│   ├── enum/                               # 枚举类型定义
│   │   ├── contact/                        # 联系人枚举
│   │   │   ├── contact_status_enum/        # 联系人状态（正常、拉黑）
│   │   │   └── contact_type_enum/          # 联系人类型（用户、群组）
│   │   ├── contact_apply/                  # 好友申请枚举
│   │   │   └── contact_apply_status_enum/  # 申请状态（待处理、已同意、已拒绝）
│   │   ├── group_info/                     # 群组枚举
│   │   │   ├── add_mode_enum/              # 加群方式（需审核、直接加入）
│   │   │   └── group_status_enum/          # 群组状态（正常、禁用）
│   │   ├── message/                        # 消息枚举
│   │   │   ├── message_status_enum.go      # 消息状态（已发送、未发送）
│   │   │   └── message_type_enum.go        # 消息类型（文本、文件、音视频）
│   │   └── user_info/                      # 用户枚举
│   │       └── user_status_enum.go         # 用户状态（正常、禁用）
│   │
│   ├── ssl/                                # SSL/TLS 证书管理
│   │   └── tls_handler.go                  # HTTPS 重定向中间件
│   │
│   ├── util/                               # 工具函数
│   │   └── random/                         # 随机数生成工具
│   │       └── random_int.go               # UUID 生成等
│   │
│   └── zlog/                               # 日志管理
│       └── logger.go                       # Zap 日志封装
│
├── configs/                                # 配置文件目录
│   ├── config.toml                         # 生产环境配置示例
│   └── config_local.toml                   # 本地开发配置（需自行创建）
│
├── static/                                 # 静态资源目录
│   ├── avatars/                            # 用户头像存储
│   └── files/                              # 聊天文件存储
│
├── test/                                   # 测试代码
│   ├── config/                             # 配置测试
│   ├── dao/                                # 数据库测试
│   └── zlog/                               # 日志测试
│
├── docs/                                   # 项目文档
│   ├── 业务逻辑.md                          # 业务逻辑详细说明
│   └── image/                              # 文档图片
│
├── go.mod                                  # Go 模块依赖管理
├── go.sum                                  # 依赖校验文件
├── LICENSE                                 # 开源协议
└── README.md                               # 项目说明文档
```

## 前端项目结构

前端采用 Vue3 + Vue Router + Vuex 架构，实现 SPA（单页应用）。

```
web/chat-server/
├── src/                                    # 源代码目录
│   ├── main.js                             # Vue 应用入口文件
│   ├── App.vue                             # 根组件
│   │
│   ├── assets/                             # 静态资源
│   │   ├── css/                            # 样式文件
│   │   │   └── chat.css                    # 聊天界面样式
│   │   ├── img/                            # 图片资源
│   │   │   └── chat_server_background.jpg  # 背景图
│   │   └── js/                             # 工具函数
│   │       ├── random.js                   # 随机数生成
│   │       └── valid.js                    # 表单验证
│   │
│   ├── components/                         # 公共组件
│   │   ├── Modal.vue                       # 通用模态框组件
│   │   ├── SmallModal.vue                  # 小型模态框组件
│   │   ├── VideoModal.vue                  # 音视频通话模态框（WebRTC）
│   │   ├── NavigationModal.vue             # 导航模态框
│   │   ├── ContactListModal.vue            # 联系人列表选择框
│   │   ├── DeleteGroupModal.vue            # 删除群组确认框（管理员）
│   │   ├── DeleteUserModal.vue             # 删除用户确认框（管理员）
│   │   ├── DisableGroupModal.vue           # 禁用群组确认框（管理员）
│   │   ├── DisableUserModal.vue            # 禁用用户确认框（管理员）
│   │   └── SetAdminModal.vue               # 设置管理员确认框
│   │
│   ├── router/                             # 路由配置
│   │   └── index.js                        # Vue Router 路由表配置
│   │
│   ├── store/                              # 状态管理
│   │   └── index.js                        # Vuex 全局状态管理（用户信息、消息等）
│   │
│   └── views/                              # 页面视图
│       ├── access/                         # 登录注册模块
│       │   ├── Login.vue                   # 账号密码登录页面
│       │   ├── Register.vue                # 注册页面
│       │   └── SmsLogin.vue                # 短信验证码登录页面
│       │
│       ├── chat/                           # 聊天模块
│       │   ├── contact/                    # 联系人相关
│       │   │   ├── ContactList.vue         # 联系人列表（显示好友、群组、申请）
│       │   │   └── ContactChat.vue         # 聊天界面（消息列表、输入框、WebRTC）
│       │   ├── session/                    # 会话相关
│       │   │   └── SessionList.vue         # 会话列表（最近聊天）
│       │   └── user/                       # 用户相关
│       │       └── OwnInfo.vue             # 个人信息设置页面
│       │
│       └── manager/                        # 管理员模块
│           └── Manager.vue                 # 后台管理页面（用户/群组管理）
│
├── public/                                 # 公共静态资源
│   └── index.html                          # HTML 模板
│
├── babel.config.js                         # Babel 配置
├── jsconfig.json                           # JavaScript 项目配置
├── vue.config.js                           # Vue CLI 配置（代理、打包等）
├── package.json                            # npm 依赖管理
├── yarn.lock                               # Yarn 锁定文件
├── LICENSE                                 # 开源协议
└── README.md                               # 前端说明文档
```

---

## 数据库表结构设计

项目共有 **6 张核心数据表**，采用 GORM 自动迁移生成。

### 1. user_info（用户信息表）
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | BIGINT | 自增主键 |
| uuid | VARCHAR(20) | 用户唯一标识（U + 时间戳 + 随机数）|
| nickname | VARCHAR(20) | 用户昵称 |
| telephone | CHAR(11) | 手机号（索引）|
| email | CHAR(30) | 邮箱 |
| avatar | VARCHAR(255) | 头像 URL |
| gender | TINYINT | 性别（0:男, 1:女）|
| signature | VARCHAR(100) | 个性签名 |
| password | CHAR(18) | 密码（AES 加密）|
| birthday | CHAR(8) | 生日 |
| created_at | DATETIME | 创建时间 |
| deleted_at | DATETIME | 软删除时间 |
| last_online_at | DATETIME | 最后在线时间 |
| last_offline_at | DATETIME | 最后离线时间 |
| is_admin | TINYINT | 是否管理员（0:否, 1:是）|
| status | TINYINT | 状态（0:正常, 1:禁用）|

**索引：**
- 唯一索引：`uuid`
- 普通索引：`telephone`、`created_at`、`status`

---

### 2. group_info（群组信息表）
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | BIGINT | 自增主键 |
| uuid | VARCHAR(20) | 群组唯一标识（G + 时间戳 + 随机数）|
| group_name | VARCHAR(50) | 群组名称 |
| notice | VARCHAR(255) | 群公告 |
| owner_id | VARCHAR(20) | 群主 UUID |
| avatar | VARCHAR(255) | 群头像 URL |
| members | JSON | 成员 UUID 列表（JSON 数组）|
| add_mode | TINYINT | 加群方式（0:直接加入, 1:需审核）|
| created_at | DATETIME | 创建时间 |
| status | TINYINT | 状态（0:正常, 1:禁用）|

**JSON 字段示例：**
```json
members: ["U1234567890", "U0987654321", "U1111111111"]
```

---

### 3. user_contact（用户联系人关系表）
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | BIGINT | 自增主键 |
| user_id | VARCHAR(20) | 用户 UUID |
| contact_id | VARCHAR(20) | 联系人 UUID（可以是用户或群组）|
| contact_type | TINYINT | 联系人类型（0:用户, 1:群组）|
| status | TINYINT | 状态（0:正常, 1:已拉黑）|
| created_at | DATETIME | 创建时间 |

**联合索引：** (`user_id`, `contact_id`)

---

### 4. contact_apply（好友申请表）
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | BIGINT | 自增主键 |
| apply_user_id | VARCHAR(20) | 申请人 UUID |
| receive_user_id | VARCHAR(20) | 接收人 UUID |
| contact_type | TINYINT | 联系人类型（0:用户, 1:群组）|
| contact_id | VARCHAR(20) | 目标联系人 UUID |
| apply_info | VARCHAR(255) | 申请说明 |
| status | TINYINT | 状态（0:待处理, 1:已同意, 2:已拒绝, 3:已拉黑）|
| last_apply_time | DATETIME | 最后申请时间 |

---

### 5. session（会话表）
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | BIGINT | 自增主键 |
| uuid | VARCHAR(20) | 会话唯一标识（S + 时间戳 + 随机数）|
| user_id | VARCHAR(20) | 用户 UUID |
| contact_id | VARCHAR(20) | 联系人 UUID |
| contact_type | TINYINT | 联系人类型（0:用户, 1:群组）|
| last_message | TEXT | 最后一条消息内容 |
| last_message_type | TINYINT | 最后消息类型（0:文本, 1:文件, 2:音视频）|
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

**作用：** 类似微信首页的会话列表，记录用户的聊天会话。

---

### 6. message（消息表）
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | BIGINT | 自增主键 |
| uuid | VARCHAR(20) | 消息唯一标识（M + 时间戳 + 随机数）|
| session_id | VARCHAR(20) | 所属会话 UUID |
| type | TINYINT | 消息类型（0:文本, 1:文件, 2:音视频）|
| content | TEXT | 文本消息内容 |
| url | VARCHAR(255) | 文件/图片 URL |
| send_id | VARCHAR(20) | 发送者 UUID |
| send_name | VARCHAR(50) | 发送者昵称 |
| send_avatar | VARCHAR(255) | 发送者头像 |
| receive_id | VARCHAR(20) | 接收者 UUID（用户或群组）|
| file_size | VARCHAR(20) | 文件大小 |
| file_type | VARCHAR(50) | 文件类型 |
| file_name | VARCHAR(255) | 文件名 |
| status | TINYINT | 状态（0:未发送, 1:已发送）|
| av_data | TEXT | 音视频通话数据（WebRTC SDP/ICE）|
| created_at | DATETIME | 创建时间 |

**索引：** `session_id`、`send_id`、`receive_id`

---

### 表关系图
```
user_info (用户)
    ├── 1:N → user_contact (联系人关系)
    ├── 1:N → contact_apply (好友申请)
    ├── 1:N → session (会话)
    ├── 1:N → message (发送的消息)
    └── 1:N → group_info (创建的群组)

group_info (群组)
    ├── 1:N → user_contact (群成员关系)
    ├── 1:N → session (群会话)
    └── 1:N → message (群消息)

session (会话)
    └── 1:N → message (会话中的消息)
```

---

# 核心技术架构

## 后端架构设计

### 1. 分层架构
- **API 层（Controller）**：处理 HTTP 请求，参数验证，调用 Service 层
- **Service 层**：业务逻辑处理，协调多个 DAO 操作
- **DAO 层**：数据访问层，封装数据库操作
- **Model 层**：数据模型定义，对应数据库表结构

### 2. 核心技术选型
- **Web 框架**：Gin（高性能 HTTP 框架）
- **ORM**：GORM（对象关系映射，简化数据库操作）
- **缓存**：Redis（存储会话、消息缓存，提升性能）
- **消息队列**：Kafka（可选，处理高并发消息）
- **实时通信**：WebSocket（双向通信，消息实时推送）
- **音视频**：WebRTC（P2P 音视频通话）
- **日志**：Zap（高性能结构化日志）
- **加密**：AES（密码加密）、SSL/TLS（HTTPS）

### 3. WebSocket 聊天服务架构
```
客户端 WebSocket 连接
        ↓
   WS Controller
        ↓
   Chat Server (管理所有在线客户端)
   ├── Login Channel    (登录通道)
   ├── Logout Channel   (登出通道)
   └── Transmit Channel (消息转发通道)
        ↓
   消息分发逻辑
   ├── 单聊：发送给指定用户
   ├── 群聊：发送给群组所有成员
   └── 离线消息存储到 MySQL
```

### 4. 消息流程
1. 用户 A 发送消息 → WebSocket 发送到服务器
2. 服务器接收消息 → 存储到 MySQL（持久化）
3. 判断接收者是否在线：
   - **在线**：通过 WebSocket 推送给用户 B
   - **离线**：消息已存储，用户 B 登录时自动拉取
4. 使用 Redis 缓存最近消息列表，减少数据库查询

## 前端架构设计

### 1. Vue3 架构
- **组件化开发**：复用性强，易于维护
- **Vue Router**：单页应用路由管理
- **Vuex**：全局状态管理（用户信息、消息列表等）
- **WebSocket**：与后端建立长连接，接收实时消息
- **WebRTC**：音视频通话（P2P + TURN 服务器中继）

### 2. 核心页面
- **登录注册**：账号登录、短信登录、注册
- **会话列表**：显示最近聊天，类似微信首页
- **联系人管理**：好友列表、群组列表、好友申请
- **聊天界面**：消息收发、文件上传、音视频通话
- **个人中心**：修改头像、昵称、签名等
- **后台管理**：管理员管理用户和群组（禁用、删除等）

# 项目学习路线

## 适合人群
- Go 后端开发学习者
- Vue 前端开发学习者
- 想了解即时通讯系统实现原理的开发者
- 需要毕设/课设项目参考的学生

## 学习顺序

### 阶段一：环境准备与基础知识（1-2天）

#### 1. 技术栈预习
**后端必备知识：**
- Go 基础语法（变量、函数、结构体、接口、协程、通道）
- Gin 框架基础（路由、中间件、参数绑定）
- GORM 使用（模型定义、CRUD、关联查询）
- MySQL 数据库基础
- Redis 基础操作
- WebSocket 原理

**前端必备知识：**
- Vue3 基础（组合式 API、响应式原理）
- Vue Router 路由管理
- Vuex 状态管理
- WebSocket 客户端使用
- WebRTC 基础（可后期学习）

#### 2. 环境搭建
```bash
# 1. 安装必要软件
- Go 1.20+
- Node.js 16+
- MySQL 8.0+
- Redis 6.0+
- Git

# 2. 克隆项目
git clone <repository-url>
cd KamaChat

# 3. 后端配置
cp configs/config.toml configs/config_local.toml
# 修改 config_local.toml 中的数据库、Redis 配置

# 4. 安装后端依赖
go mod tidy

# 5. 创建数据库
mysql -u root -p
CREATE DATABASE kama_chat_server;

# 6. 前端配置
cd web/chat-server
yarn install
```

---

### 阶段二：数据库设计与模型层（2-3天）

#### 学习目标
理解数据库表结构设计，掌握 GORM 模型定义和关系映射。

#### 学习路径
1. **阅读数据库设计**
   - 查看 `internal/model/` 下的所有模型文件
   - 理解 6 张核心表的作用：
     - `user_info`：用户信息表
     - `group_info`：群组信息表
     - `user_contact`：用户联系人关系表
     - `contact_apply`：好友申请表
     - `message`：消息表
     - `session`：会话表

2. **重点文件**
   ```
   internal/model/user_info.go      # 用户表（理解软删除、索引）
   internal/model/message.go        # 消息表（理解消息类型）
   internal/model/group_info.go     # 群组表（理解 JSON 字段存储成员）
   ```

3. **实践任务**
   - 启动项目，观察 GORM 自动建表
   - 尝试手动插入测试数据
   - 理解 GORM 标签含义（`gorm:"column:xxx;index;not null"`）

---

### 阶段三：配置与工具层（1天）

#### 学习目标
理解项目配置管理、日志系统、常量定义。

#### 学习路径
1. **配置管理**
   ```
   internal/config/config.go        # 配置文件解析
   configs/config_local.toml        # 配置文件
   ```

2. **日志系统**
   ```
   pkg/zlog/logger.go               # Zap 日志封装
   ```

3. **常量与枚举**
   ```
   pkg/constants/constants.go       # 全局常量
   pkg/enum/                        # 各种枚举定义
   ```

4. **工具函数**
   ```
   pkg/util/random/                 # UUID 生成
   internal/service/aes/aes.go      # AES 加密
   ```

---

### 阶段四：用户模块（3-4天）

#### 学习目标
掌握完整的用户注册、登录、信息管理流程。

#### 学习路径

**1. 注册流程**
```
前端：views/access/Register.vue
  ↓ 发送注册请求
后端：api/v1/user_info_controller.go → Register()
  ↓ 调用
后端：internal/service/gorm/user_info_service.go → RegisterUser()
  ↓ 存储
数据库：user_info 表
```

**重点代码：**
- `api/v1/user_info_controller.go` → `Register()` 函数
- `internal/service/gorm/user_info_service.go` → `RegisterUser()` 函数
- 理解密码加密、UUID 生成、参数校验

**2. 登录流程**
```
前端：views/access/Login.vue
  ↓ 发送登录请求
后端：api/v1/user_info_controller.go → Login()
  ↓ 验证密码
后端：internal/service/gorm/user_info_service.go → Login()
  ↓ 返回用户信息
前端：存储到 Vuex，建立 WebSocket 连接
```

**重点代码：**
- `api/v1/user_info_controller.go` → `Login()` 函数
- 理解密码验证、返回 Token（如果有）

**3. 短信登录**
```
后端：internal/service/sms/auth_code_service.go  # 阿里云短信
api/v1/user_info_controller.go → SendSmsCode()、SmsLogin()
```

**4. 用户信息管理**
- 修改个人信息：`UpdateUserInfo()`
- 上传头像：`api/v1/message_controller.go → UploadAvatar()`

**实践任务：**
- 注册一个新用户
- 使用账号密码登录
- 修改用户昵称、头像

---

### 阶段五：WebSocket 与实时通信（4-5天，核心）

#### 学习目标
理解 WebSocket 连接管理、消息转发机制、在线状态管理。

#### 学习路径

**1. WebSocket 连接建立**
```
前端：src/views/chat/contact/ContactChat.vue
  ↓ 建立 WebSocket 连接
后端：api/v1/ws_controller.go → WsLogin()
  ↓ 升级 HTTP 为 WebSocket
后端：internal/service/chat/client.go → Client 结构体
  ↓ 注册到聊天服务器
后端：internal/service/chat/server.go → Login Channel
```

**重点代码：**
```go
// api/v1/ws_controller.go
func WsLogin(c *gin.Context) {
    // 1. 升级为 WebSocket 连接
    conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
    
    // 2. 创建 Client 对象
    client := &chat.Client{
        Uuid: uuid,
        Conn: conn,
        SendBack: make(chan *chat.MessageBack),
    }
    
    // 3. 注册到 ChatServer
    chat.ChatServer.SendClientToLogin(client)
    
    // 4. 启动接收和发送协程
    go client.Read()
    go client.Write()
}
```

**2. 聊天服务器架构**
```
internal/service/chat/server.go      # 核心：Server 结构体
internal/service/chat/client.go      # 客户端连接封装
```

**Server 结构体解析：**
```go
type Server struct {
    Clients  map[string]*Client  // 所有在线客户端（key: uuid）
    mutex    *sync.Mutex         // 并发安全锁
    Transmit chan []byte         // 消息转发通道
    Login    chan *Client        // 登录通道
    Logout   chan *Client        // 登出通道
}
```

**3. 消息收发流程**

**发送消息：**
```
前端：用户输入消息，点击发送
  ↓ WebSocket.send(JSON.stringify(message))
后端：Client.Read() 接收消息
  ↓ 发送到 Transmit Channel
后端：Server.Start() 处理 Transmit Channel
  ↓ 判断接收者类型（单聊/群聊）
  ↓ 存储到 MySQL
  ↓ 查找接收者是否在线
  ↓ 如果在线，发送到接收者的 SendBack Channel
后端：Client.Write() 发送消息
  ↓ WebSocket 推送给前端
前端：接收消息，显示在聊天界面
```

**4. 消息类型处理**
- 文本消息：`message_type_enum.Text`
- 文件消息：`message_type_enum.File`
- 音视频消息：`message_type_enum.AudioOrVideo`

**实践任务：**
- 注册两个用户，互相发送消息
- 观察后端日志，理解消息流转过程
- 尝试断开 WebSocket，观察离线消息处理

---

### 阶段六：联系人与会话管理（2-3天）

#### 学习目标
掌握好友添加、群组创建、会话管理逻辑。

#### 学习路径

**1. 好友申请流程**
```
前端：ContactList.vue → 添加好友
  ↓
后端：api/v1/user_contact_controller.go → ApplyContact()
  ↓ 存储到 contact_apply 表
前端：被申请人查看好友申请
  ↓
后端：GetNewContactList() 查询待处理申请
  ↓
前端：同意/拒绝申请
  ↓
后端：PassContactApply() / RefuseContactApply()
  ↓ 同意后插入 user_contact 表
```

**2. 群组管理**
```
api/v1/group_info_controller.go      # 群组接口
internal/service/gorm/group_info_service.go  # 群组业务逻辑
```

**核心功能：**
- 创建群组：`CreateGroup()`
- 加入群组：`EnterGroupDirectly()` / `CheckGroupAddMode()`
- 退出群组：`LeaveGroup()`
- 解散群组：`DismissGroup()`

**3. 会话管理**
```
api/v1/session_controller.go
internal/service/gorm/session_service.go
```

**会话列表：**
- 用户会话：`GetUserSessionList()`（类似微信首页）
- 群组会话：`GetGroupSessionList()`

**实践任务：**
- 创建一个群组
- 邀请多个用户加入
- 在群组中发送消息
- 观察会话列表更新

---

### 阶段七：消息管理与文件上传（2天）

#### 学习目标
理解消息持久化、消息列表查询、文件上传处理。

#### 学习路径

**1. 消息列表查询**
```
api/v1/message_controller.go
  ├── GetMessageList()        # 单聊消息列表
  └── GetGroupMessageList()   # 群聊消息列表
```

**Redis 缓存策略：**
```go
// 第一次查询：从 MySQL 读取，存入 Redis
key: "message_list_{sendId}_{receiveId}"
value: JSON 格式的消息列表
TTL: 30 分钟

// 后续查询：直接从 Redis 读取
```

**2. 文件上传**
```
api/v1/message_controller.go
  ├── UploadAvatar()  # 上传头像
  └── UploadFile()    # 上传聊天文件
```

**实践任务：**
- 发送文本、图片、文件消息
- 查看数据库 message 表
- 查看 Redis 缓存

---

### 阶段八：音视频通话（3-4天，难点）

#### 学习目标
理解 WebRTC P2P 通信原理，掌握 ICE、SDP、TURN 服务器配置。

#### 学习路径

**1. WebRTC 基础**
- 学习 WebRTC 三大核心：
  - **ICE**：网络穿透（NAT 打洞）
  - **SDP**：会话描述协议
  - **STUN/TURN**：中继服务器

**2. 前端 WebRTC 实现**
```
src/views/chat/contact/ContactChat.vue  # 核心代码
src/components/VideoModal.vue          # 音视频界面
```

**通话流程：**
```
用户 A：点击"视频通话"
  ↓ 创建 RTCPeerConnection
  ↓ 发送 offer（通过 WebSocket）
  ↓
后端：转发 offer 给用户 B
  ↓
用户 B：收到通话请求
  ↓ 创建 RTCPeerConnection
  ↓ 发送 answer（通过 WebSocket）
  ↓
用户 A：收到 answer
  ↓
双方交换 ICE Candidate
  ↓
建立 P2P 连接，开始通话
```

**3. TURN 服务器配置**
```javascript
const ICE_CFG = {
  iceServers: [
    {
      urls: 'turn:your-server-ip:3478',
      username: 'your-username',
      credential: 'your-password'
    }
  ]
};
```

**实践任务：**
- 本地测试（无需 TURN 服务器）
- 公网测试（需配置 TURN 服务器）

---

### 阶段九：后台管理（1-2天）

#### 学习目标
理解管理员权限控制、用户管理、群组管理。

#### 学习路径

**1. 管理员判断**
```go
// 在 Controller 中判断
if user.IsAdmin != 1 {
    JsonBack(c, "权限不足", -2, nil)
    return
}
```

**2. 管理功能**
```
api/v1/user_info_controller.go
  ├── DisableUsers()  # 禁用用户
  ├── DeleteUsers()   # 删除用户
  └── SetAdmin()      # 设置管理员

api/v1/group_info_controller.go
  ├── SetGroupsStatus()  # 禁用群组
  └── DeleteGroups()     # 删除群组
```

**3. 前端管理页面**
```
src/views/manager/Manager.vue
```

---

### 阶段十：项目部署（2-3天）

#### 学习目标
掌握 Ubuntu 服务器部署、HTTPS 配置、MySQL/Redis 部署。

#### 学习路径
1. 按照 README.md 中的部署脚本操作
2. 配置域名和 SSL 证书
3. 配置 TURN 服务器（coturn）
4. 使用 systemd 管理后端服务

---

## 学习建议

### 1. 循序渐进
- 不要一次性理解所有代码
- 从单个功能模块入手，跟踪完整流程
- 边学边实践，修改代码观察效果

### 2. 调试技巧
- 后端：使用 `zlog.Info()` / `zlog.Debug()` 打印日志
- 前端：使用 `console.log()` 查看数据流
- 数据库：使用 MySQL Workbench 查看表数据
- Redis：使用 Redis Desktop Manager 查看缓存

### 3. 扩展练习
- 添加表情包功能
- 实现消息撤回
- 添加群组管理员功能
- 实现消息已读/未读状态
- 优化消息加载（分页、虚拟滚动）

### 4. 参考文档
- Gin 官方文档：https://gin-gonic.com/
- GORM 官方文档：https://gorm.io/
- Vue3 官方文档：https://cn.vuejs.org/
- WebRTC 教程：https://webrtc.org/

---

# docs

在 `/docs/业务逻辑.md` 中，介绍了具体的业务设计，对业务有问题的同学可以查看了解一下。

# todoList

-  多对多群聊

-  nginx分布式部署

  
