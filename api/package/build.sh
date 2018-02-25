#!/usr/bin/env bash

PACKAGE_NAME=com-msleevi-api
DIR_NAME=com/msleevi/api

build_rpm() {
  local opt="build/root/opt/${DIR_NAME}"
  mkdir -p build/dist
  mkdir -p $opt

  cp ../src/bin/api $opt/

  fpm -s dir                                            \
    -t rpm                                              \
    -a x86_64                                           \
    -p build/dist                                       \
    -n ${PACKAGE_NAME}                                  \
    -v ${API_VERSION}.${BUILD_NUMBER}                   \
    --iteration ${ITERATION}                            \
    --prefix /                                          \
    --before-install scripts/1_BeforeInstallUpgrade.sh  \
    --before-upgrade scripts/1_BeforeInstallUpgrade.sh  \
    --after-install scripts/2_AfterInstallUpgrade.sh    \
    --after-upgrade scripts/2_AfterInstallUpgrade.sh    \
    --before-remove scripts/3_BeforeRemove.sh           \
    --after-remove scripts/4_AfterRemove.sh             \
    --rpm-rpmbuild-define "_version ${API_VERSION}.${BUILD_NUMBER}" \
    --rpm-rpmbuild-define "_release ${ITERATION}"               \
    -C build/root   \
    --log info      \
    .
}

set -e

if [ -z ${API_VERSION} ]; then
  API_VERSION="1.0"
fi

if [ -z ${BUILD_NUMBER} ]; then
  BUILD_NUMBER=$(rpm -qa ${PACKAGE_NAME} | cut -d- -f3 | cut -d. -f3)
  if [ -z ${BUILD_NUMBER} ]; then
    BUILD_NUMBER=0
  fi
fi

if [ -z ${ITERATION} ]; then
  ITERATION=$(rpm -qa ${PACKAGE_NAME} | cut -d- -f4 | cut -d. -f1)
  if [ -z ${ITERATION} ]; then
    ITERATION=1
  else
    ((ITERATION += 1))
  fi
fi

echo "Building ${API_VERSION}.${BUILD_NUMBER}-${ITERATION}"
build_rpm
echo "${API_VERSION}.${BUILD_NUMBER}-${ITERATION}" > build/dist/${PACKAGE_NAME}-latest-version.txt
