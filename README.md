# procTop

为了 Linux 平台开发的性能优化分析工具。

<br>

**Download**

`wget https://raw.githubusercontent.com/matsuwin/proctop/main/dist/elf.x64-proctop.tar.gz`


kworker/*      内核工作线程
cpuhp/*        N=CPUs, CPU状态管理线程
migration/*    N=CPUs, 进程迁移管理线程
idle_inject/*  N=CPUs, CPU温度控制线程 for ubuntu
irq/*          系统中断处理线程
ksoftirqd/*    软中断处理线程
kdmflush/*     设备映射管理器
jbd2/*         ext4 日志管理器

mpt/*          RAID 控制器
kcryptd_io/*    磁盘加密(dm-crypt)
kcryptd/*       磁盘加密(dm-crypt)
dmcrypt_write/* 磁盘加密(dm-crypt)


scsi_tmf*         系统总线
scsi_eh*         系统总线
card0-*        存储
rcu_*          Read Copy Update
