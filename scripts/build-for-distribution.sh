#!/usr/bin/env bash
set -o errexit

ARCHS=(amd64 arm arm64 386)
VERSION=`may -V`
BASEDIR=`pwd`

echo "==============================================="
echo "Building may version $VERSION for distribution."
echo "==============================================="
echo "Step 1: Build"
echo "==============================================="

echo "Cleaning up past builds"
rm -rf $BASEDIR/dist

mkdir $BASEDIR/dist
touch $BASEDIR/dist/.gitkeep
mkdir $BASEDIR/dist/linux

for ARCH in "${ARCHS[@]}"
do
	echo "Building architecture linux/$ARCH "
    mkdir $BASEDIR/dist/linux/$ARCH
    cp $BASEDIR/README.md $BASEDIR/dist/linux/$ARCH/README.md
    cp $BASEDIR/LICENSE $BASEDIR/dist/linux/$ARCH/LICENSE
    cp -r $BASEDIR/doc $BASEDIR/dist/linux/$ARCH/doc

    GOOS=linux GOARCH=$ARCH go build -ldflags="-s -w" -o="$BASEDIR/dist/linux/$ARCH/may" $BASEDIR/cmd/may
done

echo "========================"
echo "Step 2: Archive Creation"
echo "========================"

cd $BASEDIR/dist
for ARCH in "${ARCHS[@]}"
do
    echo "Archiving for linux/$ARCH"
    cd linux/$ARCH
    tar czf ../../may-$VERSION-linux-$ARCH.tar.gz .
    cd ../../

    echo "Removing tmp files of linux/$ARCH"
    rm -rf $BASEDIR/dist/linux/$ARCH
done

echo "Calculating sha256sums"
cd $BASEDIR/dist
for TAR in *.tar.gz
do
    sha256sum $TAR > $TAR.sha256sum
done
cd ..

rm -rf $BASEDIR/dist/linux

echo "======="
echo "Done."
echo "======="

tree -n $BASEDIR/dist
