#!/usr/bin/env bash

echo "==> Post-install processing %name-%_version-%release..."

if [ ! -e /usr/local/bin/api ]; then
  sudo ln -s /opt/com/msleevi/api/api /usr/local/bin/api
fi
