#! /bin/bash
#
# formats the first partition in a disk image file
kpartx -a tmp.raw
mke2fs /dev/mapper/loop1p1
kpartx -d tmp.raw
