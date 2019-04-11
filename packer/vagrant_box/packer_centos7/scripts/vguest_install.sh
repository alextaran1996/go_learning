#!/bin/bash
mount -o loop VBoxGuestAdditions.iso /mnt
sh /mnt/VBoxLinuxAdditions.run
sync
sleep 10s
umount /mnt
rm VBoxGuestAdditions.iso
ln -sf /opt/VBoxGuestAdditions-6.0.4/other/mount.vboxsf /sbin/mount.vboxsf