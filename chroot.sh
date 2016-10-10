# run these commands as part of chroot

cd "to mount point"
mount -t proc proc proc/
mount -t sysfs sys sys/
mount -o bind /dev dev/
cp /etc/resolv.conf etc/
chroot "mount point"
