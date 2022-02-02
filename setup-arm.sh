#!/bin/bash

pkg='elf.raspberry-proctop.tar.gz'
wget https://github.com/matsuwin/proctop/releases/download/v0.1.12/$pkg
sudo tar -C /bin -xf $pkg
sudo chmod 0777 /bin/proctop
rm -f $pkg
