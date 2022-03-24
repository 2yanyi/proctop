# ProcTop

适用于 Linux 的性能分析工具，实时显示进程的资源占用状态，类似于 TOP。支持 Java 同名进程拆分。

**同样适用于 Raspberry Pi (树莓派)。*

<br>

## Features

- 单核 CPU 使用率 TOP 进程列表，自动刷新 2s/次。
- 分等级的彩色页面渲染：红 > 黄 > 青 > 蓝。
- 同名进程自动合并，资源利用累加。
- 主机信息和处理器型号抬头展示。
- 处理器温度实时预览。
- 一键安装。

*扩展功能*
  - 本地磁盘写入速率检测

<br>

<img src="demo.png">

<br>

## Installing

*linux/amd64*

```sh
sudo curl https://raw.githubusercontent.com/matsuwin/proctop/main/setup.sh | sh
```

*linux/arm for Raspberry Pi*

```sh
sudo curl https://raw.githubusercontent.com/matsuwin/proctop/main/setup-arm.sh | sh
```

<br>

## Quick Start

```sh
$ proctop --help
Usage of proctop:
  -diskw
        Disk write rate Test
  -l int
        limit (default 10)
  -version
        show version information
```

```sh
proctop -l 33
```

<br>
