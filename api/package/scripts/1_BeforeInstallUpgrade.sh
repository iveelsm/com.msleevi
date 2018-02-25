#!/usr/bin/env bash

echo "==> Pre-install processing %name-%_version-%release..."

/usr/bin/id api > /dev/null 2>&1

if [ $? -ne 0 ]; then
  groupadd api
  useradd -m -g api api
fi

sudo mkdir -p /var/log/com/msleevi/api/
sudo chown -R api:api /var/log/com/msleevi/api
