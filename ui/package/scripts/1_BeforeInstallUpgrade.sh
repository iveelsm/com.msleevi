#!/usr/bin/env bash

echo "==> Pre-install processing %name-%_version-%release..."

/usr/bin/id ui > /dev/null 2>&1

if [ $? -ne 0 ]; then
  groupadd ui
  useradd -m -g ui ui
fi

sudo mkdir -p /var/log/com/msleevi/ui/
sudo chown -R ui:ui /var/log/com/msleevi/ui
