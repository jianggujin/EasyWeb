#  一、EasyWeb

## 1.1 简介

`EasyWeb`是用`GO`语言编写的简易`Web`服务器工具，可以用于打包后的前端项目预览或者其他静态资源的访问。

## 1.2 功能

* 自定义端口与静态资源目录
* 自定义`MineType`扩展
* 访问日志归档

# 二、快速开始

## 2.1 安装

### 2.1.1 二进制包

从发行版下载: [https://github.com/jianggujin/EasyWeb/releases](https://github.com/jianggujin/EasyWeb/releases)

### 2.2.2 源码编译

下载源码之后，运行 `sh build.sh` 命令编译

```shell
git clone https://github.com/jianggujin/EasyWeb.git
cd EasyWeb
sh build.sh
```

## 2.2 使用

1. 编辑`confg.toml`
2. 启动`EasyWeb`

```shell
# easy-web以实际运行环境为准
./bin/easy-web -config config.toml
```

3. 浏览器访问：`http://127.0.0.1:8080`

   > 8080以配置中的端口号为准

## 2.3 配置

```toml
# 服务器配置
[source]
# 端口号，默认值：8080
port = 8080
# 静态资源目录，默认值：static/
static = "static/"
# 自定义MineType类型，当无法识别文件类型时可增加该配置
mine_type = {}

# 高级配置
[advanced]
# 日志文件前缀，默认值：easy-web，不配置则不生成日志文件
log_file = "easy-web"
# 日志级别，默认值：info，可选值：debug, info、warn
log_level = "info"
# 日志归档文件最大保存天数，默认值：7，配置0表示不清理
log_max_history = 7
# 日志文件目录
log_dir = "logs"
# 日志模式，默认值：rolling，可选值为：
# single  单文件模式
# rolling 按天生成日志文件
log_mode = "rolling"
```

# 三、帮忙点个⭐Star

开源不易，如果觉得`EasyWeb`对您有帮助的话，请帮忙在<a target="_blank" href='https://github.com/jianggujin/EasyWeb'><img src="https://img.shields.io/github/stars/jianggujin/EasyWeb.svg?style=flat-square&label=Stars&logo=github" alt="github star"/></a>
的右上角点个⭐Star，您的支持是使`EasyWeb`变得更好最大的动力。如果您愿意的话，可以为作者捐赠一杯咖啡，万分感谢。

<img src="doc\alipay.png" width="45%"><img src="doc\wepay.png" width="45%">

