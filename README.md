# BaiduPCS-Go 百度网盘客户端

[![Build status](https://ci.appveyor.com/api/projects/status/nhx92nqyrfq9su7y?svg=true)](https://ci.appveyor.com/project/iikira/baidupcs-go)
[![GoDoc](https://godoc.org/github.com/iikira/BaiduPCS-Go?status.svg)](https://godoc.org/github.com/iikira/BaiduPCS-Go)

仿 Linux shell 文件处理命令的百度网盘命令行客户端.

This project was largely inspired by [GangZhuo/BaiduPCS](https://github.com/GangZhuo/BaiduPCS)

## 解决错误代码4, No permission to do this operation
```
BaiduPCS-Go config set -appid 266719
```

详见讨论 [#387](https://github.com/iikira/BaiduPCS-Go/issues/387)

## 注意

此文档只针对于最新的commit, 可能不适用于已发布的最新版本.

<!-- toc -->
## 目录

- [特色](#特色)
- [编译/交叉编译 说明](#编译交叉编译-说明)
- [下载/运行 说明](#下载运行-说明)
  * [Windows](#windows)
  * [Linux / macOS](#linux--macos)
  * [Android / iOS](#android--ios)
- [命令列表及说明](#命令列表及说明)
  * [注意 ! ! !](#注意---)
  * [检测程序更新](#检测程序更新)
  * [登录百度帐号](#登录百度帐号)
  * [列出帐号列表](#列出帐号列表)
  * [获取当前帐号](#获取当前帐号)
  * [切换百度帐号](#切换百度帐号)
  * [退出百度帐号](#退出百度帐号)
  * [获取网盘配额](#获取网盘配额)
  * [切换工作目录](#切换工作目录)
  * [输出工作目录](#输出工作目录)
  * [列出目录](#列出目录)
  * [列出目录树形图](#列出目录树形图)
  * [获取文件/目录的元信息](#获取文件目录的元信息)
  * [搜索文件](#搜索文件)
  * [下载文件/目录](#下载文件目录)
  * [上传文件/目录](#上传文件目录)
  * [获取下载直链](#获取下载直链)
  * [手动秒传文件](#手动秒传文件)
  * [修复文件MD5](#修复文件MD5)
  * [获取本地文件的秒传信息](#获取本地文件的秒传信息)
  * [导出文件/目录](#导出文件目录)
  * [创建目录](#创建目录)
  * [删除文件/目录](#删除文件目录)
  * [拷贝文件/目录](#拷贝文件目录)
  * [移动/重命名文件/目录](#移动重命名文件目录)
  * [分享文件/目录](#分享文件目录)
    + [设置分享文件/目录](#设置分享文件目录)
    + [列出已分享文件/目录](#列出已分享文件目录)
    + [取消分享文件/目录](#取消分享文件目录)
  * [离线下载](#离线下载)
    + [添加离线下载任务](#添加离线下载任务)
    + [精确查询离线下载任务](#精确查询离线下载任务)
    + [查询离线下载任务列表](#查询离线下载任务列表)
    + [取消离线下载任务](#取消离线下载任务)
    + [删除离线下载任务](#删除离线下载任务)
  * [回收站](#回收站)
    + [列出回收站文件列表](#列出回收站文件列表)
    + [还原回收站文件或目录](#还原回收站文件或目录)
    + [删除回收站文件或目录/清空回收站](#删除回收站文件或目录清空回收站)
  * [显示和修改程序配置项](#显示和修改程序配置项)
  * [测试通配符](#测试通配符)
  * [工具箱](#工具箱)
- [初级使用教程](#初级使用教程)
  * [1. 查看程序使用说明](#1-查看程序使用说明)
  * [2. 登录百度帐号 (必做)](#2-登录百度帐号-必做)
  * [3. 切换网盘工作目录](#3-切换网盘工作目录)
  * [4. 网盘内列出文件和目录](#4-网盘内列出文件和目录)
  * [5. 下载文件](#5-下载文件)
  * [6. 设置下载最大并发量](#6-设置下载最大并发量)
  * [7. 退出程序](#7-退出程序)
- [常见问题](#常见问题)
- [TODO](#todo)
- [交流反馈](#交流反馈)
- [捐助](#捐助)

<!-- tocstop -->

# 特色

多平台支持, 支持 Windows, macOS, linux, 移动设备等.

百度帐号多用户支持;

通配符匹配网盘路径和 Tab 自动补齐命令和路径, [通配符_百度百科](https://baike.baidu.com/item/通配符);

[下载](#下载文件目录)网盘内文件, 支持多个文件或目录下载, 支持断点续传和单文件并行下载;

[上传](#上传文件目录)本地文件, 支持上传大文件(>2GB), 支持多个文件或目录上传;

[离线下载](#离线下载), 支持http/https/ftp/电驴/磁力链协议.

# 编译/交叉编译 说明
参见 [编译/交叉编译帮助](https://github.com/iikira/BaiduPCS-Go/wiki/编译-交叉编译帮助)

# 下载/运行 说明

Go语言程序, 可直接在[发布页](https://github.com/iikira/BaiduPCS-Go/releases)下载使用.

如果程序运行时输出乱码, 请检查下终端的编码方式是否为 `UTF-8`.

使用本程序之前, 建议学习一些 linux 基础知识 和 基础命令.

如果未带任何参数运行程序, 程序将会进入仿Linux shell系统用户界面的cli交互模式, 可直接运行相关命令.

cli交互模式下, 光标所在行的前缀应为 `BaiduPCS-Go >`, 如果登录了百度帐号则格式为 `BaiduPCS-Go:<工作目录> <百度ID>$ `

程序会提供相关命令的使用说明.

## Windows

程序应在 命令提示符 (Command Prompt) 或 PowerShell 中运行, 在 mintty (例如: GitBash) 可能会有显示问题.

也可直接双击程序运行, 具体使用方法请参见 [命令列表及说明](#命令列表及说明) 和 [初级使用教程](#初级使用教程).

## Linux / macOS

程序应在 终端 (Terminal) 运行.

具体使用方法请参见 [命令列表及说明](#命令列表及说明) 和 [初级使用教程](#初级使用教程).

## Android / iOS

> Android / iOS 移动设备操作比较麻烦, 不建议在移动设备上使用本程序.

安卓, 建议使用 [Termux](https://termux.com) 或 [NeoTerm](https://github.com/NeoTerm/NeoTerm) 或 终端模拟器, 以提供终端环境.

示例: [Android 运行本项目程序参考示例](https://github.com/iikira/BaiduPCS-Go/wiki/Android-运行本项目程序参考示例), 有兴趣的可以参考一下.

苹果iOS, 需要越狱, 在 Cydia 搜索下载并安装 MobileTerminal, 或者其他提供终端环境的软件.

示例: [iOS 运行本项目程序参考示例](https://github.com/iikira/BaiduPCS-Go/wiki/iOS-运行本项目程序参考示例), 有兴趣的可以参考一下.

具体使用方法请参见 [命令列表及说明](#命令列表及说明) 和 [初级使用教程](#初级使用教程).

# 命令列表及说明

## 注意 ! ! !

命令的前缀 `BaiduPCS-Go` 为指向程序运行的全路径名 (ARGv 的第一个参数)

直接运行程序时, 未带任何其他参数, 则程序进入cli交互模式, 运行以下命令时, 要把命令的前缀 `BaiduPCS-Go` 去掉!

cli交互模式已支持按tab键自动补全命令和路径.

## 检测程序更新
```
BaiduPCS-Go update
```

## 登录百度帐号

### 常规登录百度帐号

支持在线验证绑定的手机号或邮箱,
```
BaiduPCS-Go login
```

### 使用百度 BDUSS 来登录百度帐号

[关于 获取百度 BDUSS](https://github.com/iikira/BaiduPCS-Go/wiki/关于-获取百度-BDUSS)

```
BaiduPCS-Go login -bduss=<BDUSS>
```

#### 例子
```
BaiduPCS-Go login -bduss=1234567
```
```
BaiduPCS-Go login
请输入百度用户名(手机号/邮箱/用户名), 回车键提交 > 1234567
```

## 列出帐号列表

```
BaiduPCS-Go loglist
```

列出所有已登录的百度帐号

## 获取当前帐号

```
BaiduPCS-Go who
```

## 切换百度帐号

切换已登录的百度帐号
```
BaiduPCS-Go su <uid>
```
```
BaiduPCS-Go su

请输入要切换帐号的 # 值 >
```

## 退出百度帐号

退出当前登录的百度帐号
```
BaiduPCS-Go logout
```

程序会进一步确认退出帐号, 防止误操作.

## 获取网盘配额

```
BaiduPCS-Go quota
```
获取网盘的总储存空间, 和已使用的储存空间

## 切换工作目录
```
BaiduPCS-Go cd <目录>
```

### 切换工作目录后自动列出工作目录下的文件和目录
```
BaiduPCS-Go cd -l <目录>
```

#### 例子
```
# 切换 /我的资源 工作目录
BaiduPCS-Go cd /我的资源

# 切换 上级目录
BaiduPCS-Go cd ..

# 切换 根目录
BaiduPCS-Go cd /

# 切换 /我的资源 工作目录, 并自动列出 /我的资源 下的文件和目录
BaiduPCS-Go cd -l 我的资源

# 使用通配符
BaiduPCS-Go cd /我的*
```

## 输出工作目录
```
BaiduPCS-Go pwd
```

## 列出目录

列出当前工作目录的文件和目录或指定目录
```
BaiduPCS-Go ls
```
```
BaiduPCS-Go ls <目录>
```

### 可选参数
```
-asc: 升序排序
-desc: 降序排序
-time: 根据时间排序
-name: 根据文件名排序
-size: 根据大小排序
```

#### 例子
```
# 列出 我的资源 内的文件和目录
BaiduPCS-Go ls 我的资源

# 绝对路径
BaiduPCS-Go ls /我的资源

# 降序排序
BaiduPCS-Go ls -desc 我的资源

# 按文件大小降序排序
BaiduPCS-Go ls -size -desc 我的资源

# 使用通配符
BaiduPCS-Go ls /我的*
```

## 列出目录树形图

列出当前工作目录的文件和目录或指定目录的树形图
```
BaiduPCS-Go tree <目录>

# 默认获取工作目录元信息
BaiduPCS-Go tree
```

## 获取文件/目录的元信息
```
BaiduPCS-Go meta <文件/目录1> <文件/目录2> <文件/目录3> ...

# 默认获取工作目录元信息
BaiduPCS-Go meta
```

#### 例子
```
BaiduPCS-Go meta 我的资源
BaiduPCS-Go meta /
```

## 搜索文件

按文件名搜索文件（不支持查找目录）。

默认在当前工作目录搜索.

```
BaiduPCS-Go search [-path=<需要检索的目录>] [-r] <关键字>
```

#### 例子
```
# 搜索根目录的文件
BaiduPCS-Go search -path=/ 关键字

# 搜索当前工作目录的文件
BaiduPCS-Go search 关键字

# 递归搜索当前工作目录的文件
BaiduPCS-Go search -r 关键字
```

## 下载文件/目录
```
BaiduPCS-Go download <网盘文件或目录的路径1> <文件或目录2> <文件或目录3> ...
BaiduPCS-Go d <网盘文件或目录的路径1> <文件或目录2> <文件或目录3> ...
```

### 可选参数
```
  --test          测试下载, 此操作不会保存文件到本地
  --ow            overwrite, 覆盖已存在的文件
  --status        输出所有线程的工作状态
  --save          将下载的文件直接保存到当前工作目录
  --saveto value  将下载的文件直接保存到指定的目录
  -x              为文件加上执行权限, (windows系统无效)
  --share         以分享文件的方式获取下载链接来下载
  --locate        以获取直链的方式来下载
  -p value        指定下载线程数
```

支持多个文件或目录的下载.

下载的文件默认保存到 **程序所在目录** 的 download/ 目录, 支持设置指定目录, 重名的文件会自动跳过!

[关于下载的简单说明](https://github.com/iikira/BaiduPCS-Go/wiki/%E5%85%B3%E4%BA%8E%E4%B8%8B%E8%BD%BD%E7%9A%84%E7%AE%80%E5%8D%95%E8%AF%B4%E6%98%8E)

#### 例子
```
# 设置保存目录, 保存到 D:\Downloads
# 注意区别反斜杠 "\" 和 斜杠 "/" !!!
BaiduPCS-Go config set -savedir D:/Downloads

# 下载 /我的资源/1.mp4
BaiduPCS-Go d /我的资源/1.mp4

# 下载 /我的资源 整个目录!!
BaiduPCS-Go d /我的资源

# 下载网盘内的全部文件!!
BaiduPCS-Go d /
BaiduPCS-Go d *
```

## 上传文件/目录
```
BaiduPCS-Go upload <本地文件/目录的路径1> <文件/目录2> <文件/目录3> ... <目标目录>
BaiduPCS-Go u <本地文件/目录的路径1> <文件/目录2> <文件/目录3> ... <目标目录>
```

* 上传默认采用分片上传的方式, 上传的文件将会保存到, <目标目录>.

* 遇到同名文件将会自动覆盖!!

* 当上传的文件名和网盘的目录名称相同时, 不会覆盖目录, 防止丢失数据.


#### 注意:

* 分片上传之后, 服务器可能会记录到错误的文件md5, 可使用 fixmd5 命令尝试修复文件的MD5值, 修复md5不一定能成功, 但文件的完整性是没问题的.

fixmd5 命令使用方法:
```
BaiduPCS-Go fixmd5 -h
```

* 禁用分片上传可以保证服务器记录到正确的md5.

* 禁用分片上传时只能使用单线程上传, 指定的单个文件上传最大线程数将会无效.

#### 例子:
```
# 将本地的 C:\Users\Administrator\Desktop\1.mp4 上传到网盘 /视频 目录
# 注意区别反斜杠 "\" 和 斜杠 "/" !!!
BaiduPCS-Go upload C:/Users/Administrator/Desktop/1.mp4 /视频

# 将本地的 C:\Users\Administrator\Desktop\1.mp4 和 C:\Users\Administrator\Desktop\2.mp4 上传到网盘 /视频 目录
BaiduPCS-Go upload C:/Users/Administrator/Desktop/1.mp4 C:/Users/Administrator/Desktop/2.mp4 /视频

# 将本地的 C:\Users\Administrator\Desktop 整个目录上传到网盘 /视频 目录
BaiduPCS-Go upload C:/Users/Administrator/Desktop /视频
```

## 获取下载直链
```
BaiduPCS-Go locate <文件1> <文件2> ...
```

#### 注意

若该功能无法正常使用, 提示`user is not authorized, hitcode:101`, 尝试更换 User-Agent 为 `netdisk;8.3.1;android-android`:
```
BaiduPCS-Go config set -user_agent "netdisk;8.3.1;android-android"
```

## 手动秒传文件
```
BaiduPCS-Go rapidupload -length=<文件的大小> -md5=<文件的md5值> -slicemd5=<文件前256KB切片的md5值(可选)> -crc32=<文件的crc32值(可选)> <保存的网盘路径, 需包含文件名>
BaiduPCS-Go ru -length=<文件的大小> -md5=<文件的md5值> -slicemd5=<文件前256KB切片的md5值(可选)> -crc32=<文件的crc32值(可选)> <保存的网盘路径, 需包含文件名>
```

注意: 使用此功能秒传文件, 前提是知道文件的大小, md5, 前256KB切片的 md5 (可选), crc32 (可选), 且百度网盘中存在一模一样的文件.

上传的文件将会保存到网盘的目标目录.

遇到同名文件将会自动覆盖! 

可能无法秒传 20GB 以上的文件!!

#### 例子:
```
# 如果秒传成功, 则保存到网盘路径 /test
BaiduPCS-Go rapidupload -length=56276137 -md5=fbe082d80e90f90f0fb1f94adbbcfa7f -slicemd5=38c6a75b0ec4499271d4ea38a667ab61 -crc32=314332359 /test
```


## 修复文件MD5
```
BaiduPCS-Go fixmd5 <文件1> <文件2> <文件3> ...
```

尝试修复文件的MD5值, 以便于校验文件的完整性和导出文件.

使用分片上传文件, 当文件分片数大于1时, 百度网盘服务端最终计算所得的md5值和本地的不一致, 这可能是百度网盘的bug.

不过把上传的文件下载到本地后，对比md5值是匹配的, 也就是文件在传输中没有发生损坏.

对于MD5值可能有误的文件, 程序会在获取文件的元信息时, 给出MD5值 "可能不正确" 的提示, 表示此文件可以尝试进行MD5值修复.

修复文件MD5不一定能成功, 原因可能是服务器未刷新, 可过几天后再尝试.

修复文件MD5的原理为秒传文件, 即修复文件MD5成功后, 文件的**创建日期, 修改日期, fs_id, 版本历史等信息**将会被覆盖, 修复的MD5值将覆盖原先的MD5值, 但不影响文件的完整性.

注意: 无法修复 **20GB** 以上文件的 md5!!

#### 例子:
```
# 修复 /我的资源/1.mp4 的 MD5 值
BaiduPCS-Go fixmd5 /我的资源/1.mp4
```

## 获取本地文件的秒传信息
```
BaiduPCS-Go sumfile <本地文件的路径>
BaiduPCS-Go sf <本地文件的路径>
```

获取本地文件的大小, md5, 前256KB切片的 md5, crc32, 可用于秒传文件.

#### 例子:
```
# 获取 C:\Users\Administrator\Desktop\1.mp4 的秒传信息
BaiduPCS-Go sumfile C:/Users/Administrator/Desktop/1.mp4
```

## 导出文件/目录
```
BaiduPCS-Go export <文件/目录1> <文件/目录2> ...
BaiduPCS-Go ep <文件/目录1> <文件/目录2> ...
```

导出网盘内的文件或目录, 原理为秒传文件, 此操作会生成导出文件或目录的命令.

#### 注意

**无法导出 20GB 以上的文件!!**

**无法导出文件的版本历史等数据!!**

并不是所有的文件都能导出成功, 程序会列出无法导出的文件列表

#### 例子:
```
# 导出当前工作目录:
BaiduPCS-Go export

# 导出所有文件和目录, 并设置新的根目录为 /root
BaiduPCS-Go export -root=/root /

# 导出 /我的资源
BaiduPCS-Go export /我的资源
```

## 创建目录
```
BaiduPCS-Go mkdir <目录>
```

#### 例子
```
BaiduPCS-Go mkdir 123
```

## 删除文件/目录
```
BaiduPCS-Go rm <网盘文件或目录的路径1> <文件或目录2> <文件或目录3> ...
```

注意: 删除多个文件和目录时, 请确保每一个文件和目录都存在, 否则删除操作会失败.

被删除的文件或目录可在网盘文件回收站找回.

#### 例子
```
# 删除 /我的资源/1.mp4
BaiduPCS-Go rm /我的资源/1.mp4

# 删除 /我的资源/1.mp4 和 /我的资源/2.mp4
BaiduPCS-Go rm /我的资源/1.mp4 /我的资源/2.mp4

# 删除 /我的资源 内的所有文件和目录, 但不删除该目录
BaiduPCS-Go rm /我的资源/*

# 删除 /我的资源 整个目录 !!
BaiduPCS-Go rm /我的资源
```

## 拷贝文件/目录
```
BaiduPCS-Go cp <文件/目录> <目标 文件/目录>
BaiduPCS-Go cp <文件/目录1> <文件/目录2> <文件/目录3> ... <目标目录>
```

注意: 拷贝多个文件和目录时, 请确保每一个文件和目录都存在, 否则拷贝操作会失败.

#### 例子
```
# 将 /我的资源/1.mp4 复制到 根目录 /
BaiduPCS-Go cp /我的资源/1.mp4 /

# 将 /我的资源/1.mp4 和 /我的资源/2.mp4 复制到 根目录 /
BaiduPCS-Go cp /我的资源/1.mp4 /我的资源/2.mp4 /
```

## 移动/重命名文件/目录
```
# 移动:
BaiduPCS-Go mv <文件/目录1> <文件/目录2> <文件/目录3> ... <目标目录>
# 重命名:
BaiduPCS-Go mv <文件/目录> <重命名的文件/目录>
```

注意: 移动多个文件和目录时, 请确保每一个文件和目录都存在, 否则移动操作会失败.

#### 例子
```
# 将 /我的资源/1.mp4 移动到 根目录 /
BaiduPCS-Go mv /我的资源/1.mp4 /

# 将 /我的资源/1.mp4 重命名为 /我的资源/3.mp4
BaiduPCS-Go mv /我的资源/1.mp4 /我的资源/3.mp4
```

## 分享文件/目录
```
BaiduPCS-Go share
```

### 设置分享文件/目录
```
BaiduPCS-Go share set <文件/目录1> <文件/目录2> ...
BaiduPCS-Go share s <文件/目录1> <文件/目录2> ...
```

### 列出已分享文件/目录
```
BaiduPCS-Go share list
BaiduPCS-Go share l
```

### 取消分享文件/目录
```
BaiduPCS-Go share cancel <shareid_1> <shareid_2> ...
BaiduPCS-Go share c <shareid_1> <shareid_2> ...
```

目前只支持通过分享id (shareid) 来取消分享.

## 离线下载
```
BaiduPCS-Go offlinedl
BaiduPCS-Go clouddl
BaiduPCS-Go od
```

离线下载支持http/https/ftp/电驴/磁力链协议

离线下载同时进行的任务数量有限, 超出限制的部分将无法添加.

### 添加离线下载任务
```
BaiduPCS-Go offlinedl add -path=<离线下载文件保存的路径> 资源地址1 地址2 ...
```

添加任务成功之后, 返回离线下载的任务ID.

### 精确查询离线下载任务
```
BaiduPCS-Go offlinedl query 任务ID1 任务ID2 ...
```

### 查询离线下载任务列表
```
BaiduPCS-Go offlinedl list
```

### 取消离线下载任务
```
BaiduPCS-Go offlinedl cancel 任务ID1 任务ID2 ...
```

### 删除离线下载任务
```
BaiduPCS-Go offlinedl delete 任务ID1 任务ID2 ...

# 清空离线下载任务记录, 程序不会进行二次确认, 谨慎操作!!!
BaiduPCS-Go offlinedl delete -all
```

#### 例子
```
# 将百度和腾讯主页, 离线下载到根目录 /
BaiduPCS-Go offlinedl add -path=/ http://baidu.com http://qq.com

# 添加磁力链接任务
BaiduPCS-Go offlinedl add magnet:?xt=urn:btih:xxx

# 查询任务ID为 12345 的离线下载任务状态
BaiduPCS-Go offlinedl query 12345

# 取消任务ID为 12345 的离线下载任务
BaiduPCS-Go offlinedl cancel 12345
```

## 回收站
```
BaiduPCS-Go recycle
```

回收站操作.

### 列出回收站文件列表
```
BaiduPCS-Go recycle list
```

#### 可选参数
```
  --page value  回收站文件列表页数 (default: 1)
```

### 还原回收站文件或目录
```
BaiduPCS-Go recycle restore <fs_id 1> <fs_id 2> <fs_id 3> ...
```

根据文件/目录的 fs_id, 还原回收站指定的文件或目录.

### 删除回收站文件或目录/清空回收站
```
BaiduPCS-Go recycle delete [-all] <fs_id 1> <fs_id 2> <fs_id 3> ...
```

根据文件/目录的 fs_id 或 -all 参数, 删除回收站指定的文件或目录或清空回收站.

#### 例子
```
# 从回收站还原两个文件, 其中的两个文件的 fs_id 分别为 1013792297798440 和 643596340463870
BaiduPCS-Go recycle restore 1013792297798440 643596340463870

# 从回收站删除两个文件, 其中的两个文件的 fs_id 分别为 1013792297798440 和 643596340463870
BaiduPCS-Go recycle delete 1013792297798440 643596340463870

# 清空回收站, 程序不会进行二次确认, 谨慎操作!!!
BaiduPCS-Go recycle delete -all
```

## 显示程序环境变量
```
BaiduPCS-Go env
```

BAIDUPCS_GO_CONFIG_DIR: 配置文件路径,

BAIDUPCS_GO_VERBOSE: 是否启用调试.

## 显示和修改程序配置项
```
# 显示配置
BaiduPCS-Go config

# 设置配置
BaiduPCS-Go config set
```

注意: v3.5 以后, 程序对配置文件储存路径的寻找做了调整, 配置文件所在的目录可以是程序本身所在目录, 也可以是家目录.

配置文件所在的目录为家目录的情况:

Windows: `%APPDATA%\BaiduPCS-Go`

其他操作系统: `$HOME/.config/BaiduPCS-Go`

可通过设置环境变量 `BAIDUPCS_GO_CONFIG_DIR`, 指定配置文件存放的目录.

#### 例子
```
# 显示所有可以设置的值
BaiduPCS-Go config -h
BaiduPCS-Go config set -h

# 设置下载文件的储存目录
BaiduPCS-Go config set -savedir D:/Downloads

# 设置下载最大并发量为 150
BaiduPCS-Go config set -max_parallel 150

# 组合设置
BaiduPCS-Go config set -max_parallel 150 -savedir D:/Downloads
```

## 测试通配符
```
BaiduPCS-Go match <通配符表达式>
```

测试通配符匹配路径, 操作成功则输出所有匹配到的路径.

#### 例子
```
# 匹配 /我的资源 目录下所有mp4格式的文件
BaiduPCS-Go match /我的资源/*.mp4
```

## 工具箱
```
BaiduPCS-Go tool
```

目前工具箱支持加解密文件等.

# 初级使用教程

新手建议: **双击运行程序**, 进入仿 Linux shell 的 cli 交互模式;

cli交互模式下, 光标所在行的前缀应为 `BaiduPCS-Go >`, 如果登录了百度帐号则格式为 `BaiduPCS-Go:<工作目录> <百度ID>$ `

以下例子的命令, 均为 cli交互模式下的命令

运行命令的正确操作: **输入命令, 按一下回车键 (键盘上的 Enter 键)**, 程序会接收到命令并输出结果

## 1. 查看程序使用说明

cli交互模式下, 运行命令 `help`

## 2. 登录百度帐号 (必做)

cli交互模式下, 运行命令 `login -h` (注意空格) 查看帮助

cli交互模式下, 运行命令 `login` 程序将会提示你输入百度用户名(手机号/邮箱/用户名)和密码, 必要时还可以在线验证绑定的手机号或邮箱

## 3. 切换网盘工作目录

cli交互模式下, 运行命令 `cd /我的资源` 将工作目录切换为 `/我的资源` (前提: 该目录存在于网盘)

目录支持通配符匹配, 所以你也可以这样: 运行命令 `cd /我的*` 或 `cd /我的??` 将工作目录切换为 `/我的资源`, 简化输入.

将工作目录切换为 `/我的资源` 成功后, 运行命令 `cd ..` 切换上级目录, 即将工作目录切换为 `/`

为什么要这样设计呢, 举个例子,

假设 你要下载 `/我的资源` 内名为 `1.mp4` 和 `2.mp4` 两个文件, 而未切换工作目录, 你需要依次运行以下命令:

```
d /我的资源/1.mp4
d /我的资源/2.mp4
```

而切换网盘工作目录之后, 依次运行以下命令:

```
cd /我的资源
d 1.mp4
d 2.mp4
```

这样就达到了简化输入的目的

## 4. 网盘内列出文件和目录

cli交互模式下, 运行命令 `ls -h` (注意空格) 查看帮助

cli交互模式下, 运行命令 `ls` 来列出当前所在目录的文件和目录

cli交互模式下, 运行命令 `ls /我的资源` 来列出 `/我的资源` 内的文件和目录

cli交互模式下, 运行命令 `ls ..` 来列出当前所在目录的上级目录的文件和目录

## 5. 下载文件

说明: 下载的文件默认保存到 download/ 目录 (文件夹)

cli交互模式下, 运行命令 `d -h` (注意空格) 查看帮助

cli交互模式下, 运行命令 `d /我的资源/1.mp4` 来下载位于 `/我的资源/1.mp4` 的文件 `1.mp4` , 该操作等效于运行以下命令:

```
cd /我的资源
d 1.mp4
```

现在已经支持目录 (文件夹) 下载, 所以, 运行以下命令, 会下载 `/我的资源` 内的所有文件 (违规文件除外):

```
d /我的资源
```

参见 例6 设置下载最大并发量

## 6. 设置下载最大并发量

cli交互模式下, 运行命令 `config set -h` (注意空格) 查看设置帮助以及可供设置的值

cli交互模式下, 运行命令 `config set -max_parallel 250` 将下载最大并发量设置为 250

下载最大并发量建议值: 50~500, 太低下载速度提升不明显甚至速度会变为0, 太高可能会导致程序出错被操作系统结束掉.

## 7. 退出程序

运行命令 `quit` 或 `exit` 或 组合键 `Ctrl+C` 或 组合键 `Ctrl+D`

# 已知问题

* 分片上传文件时, 当文件分片数大于1, 网盘端最终计算所得的md5值和本地的不一致, 这可能是百度网盘的bug, 测试把上传的文件下载到本地后，对比md5值是匹配的. 可通过秒传的原理来修复md5值.

# 常见问题

参见 [常见问题](https://github.com/iikira/BaiduPCS-Go/wiki/%E5%B8%B8%E8%A7%81%E9%97%AE%E9%A2%98)

# TODO

1. 上传大文件;
2. 回收站操作, 例如查询回收站文件, 还原文件或目录等.

# 交流反馈

提交Issue: [Issues](https://github.com/iikira/BaiduPCS-Go/issues)

邮箱: i@mail.iikira.com

QQ群: 178324706

# 捐助

如果你愿意.

|支付宝|微信|
|:-----:|:-----:|
|![alipay](https://github.com/iikira/BaiduPCS-Go/raw/master/assets/donate/alipay.jpg)|![weixin](https://github.com/iikira/BaiduPCS-Go/raw/master/assets/donate/weixin.png)|
