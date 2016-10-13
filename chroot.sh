# run these commands as part of chroot
echo "args ${#}"
if test ${#} -eq 0 ; then
  echo "Usage:  `basename $0` <mount point>"
  exit 1
fi

MNT=${1}
if test ! -d ${MNT} ; then
  echo "`basename $0` -- ${MNT} is not a directory"
  exit 2
fi

cd ${MNT}
mount -t proc proc proc/
mount -t sysfs sys sys/
mount -o bind /dev dev/
cp /etc/resolv.conf etc/
chroot ${MNT}
umount ${MNT}/proc
umount ${MNT}/sys
umount ${MNT}/dev
