#!/bin/bash

pkg='elf.arm-proctop.tar.gz'
wget https://github.com/matsuwin/proctop/releases/download/v0.1.13/$pkg
sudo tar -C /bin -xf $pkg
sudo chmod 0777 /bin/proctop
rm -f $pkg
