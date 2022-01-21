#!/bin/bash

pkg='elf.raspberry-proctop.tar.gz'
wget https://github.com/matsuwin/proctop/releases/download/v0.1.11/$pkg
sudo tar -C /bin -xf $pkg
rm -f $pkg
