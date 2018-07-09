#!/usr/bin/env bash

PACKAGE_NAME=com-msleevi
DIR_NAME=com/msleevi/ui

build_rpm() {
  opt=build/root/opt/$DIR_NAME
  mkdir -p build/dist

  cp -rp root/ build/root

  fpm -s dir \
    -t rpm \
    -a x86_64 \
	-p build/dist \
	-n $PACKAGE_NAME \
    -v $VERSION.$BUILD_NUMBER \
    --iteration $ITERATION --prefix / \
    --before-install scripts/1_BeforeInstallUpgrade.sh \
    --before-upgrade scripts/1_BeforeInstallUpgrade.sh \
    --after-install  scripts/2_AfterInstallUpgrade.sh  \
    --after-upgrade  scripts/2_AfterInstallUpgrade.sh  \
    --before-remove  scripts/3_BeforeRemove.sh         \
    --after-remove   scripts/4_AfterRemove.sh          \
    --rpm-rpmbuild-define "_version $VERSION.$BUILD_NUMBER" \
    --rpm-rpmbuild-define "_release $ITERATION" \
    -C build/root \
    --log info \
    .
}

set -e

if [ -z "${VERSION}" ]; then
    VERSION=1.0
fi

if [ -z "${BUILD_NUMBER}" ]; then
    # If the package has already been installed, make sure to use the same build number.
    # This is used for developers. The CI machine will specify the build number
    BUILD_NUMBER=$(rpm -qa ${PACKAGE_NAME} | cut -d- -f4 | cut -d. -f3)
    if [ -z "${BUILD_NUMBER}" ]; then
        BUILD_NUMBER=0
    fi
fi
if [ -z "${ITERATION}" ]; then
    # If the package has already been installed, make sure to increment the iteration.
    # This is used for developers. The CI machine will always use 1
    ITERATION=$(rpm -qa ${PACKAGE_NAME} | cut -d- -f5 | cut -d. -f1)
    if [ -z "${ITERATION}" ]; then
        ITERATION=1
    else
        ((ITERATION += 1))
    fi
fi

echo Building $VERSION.$BUILD_NUMBER

build_rpm
echo "$VERSION.$BUILD_NUMBER-$ITERATION" > build/dist/$PACKAGE_NAME-latest-version.txt
