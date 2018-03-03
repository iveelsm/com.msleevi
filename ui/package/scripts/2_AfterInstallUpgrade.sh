#!/usr/bin/env bash

echo "==> Post-install processing %name-%_version-%release..."

if [ ! -e /usr/local/bin/ui ]; then
  sudo ln -s /opt/com/msleevi/ui/ui.sh /usr/local/bin/ui
fi
