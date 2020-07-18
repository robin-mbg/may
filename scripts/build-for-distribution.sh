
ARCHS=(amd64 arm arm64)
VERSION=`may -V`
BASEDIR=`pwd`

echo "==============================================="
echo "Building may version $VERSION for distribution."
echo ""
echo "Step 1: Build"
echo "==============================================="

echo "Cleaning up past builds"
rm -rf $BASEDIR/dist/linux
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


for ARCH in "${ARCHS[@]}"
do
    echo "Archiving for linux/$ARCH"
    tar czf $BASEDIR/dist/may-$VERSION-linux-$ARCH.tar.gz $BASEDIR/dist/linux/$ARCH

    echo "Removing tmp files of linux/$ARCH"
    rm -rf $BASEDIR/dist/linux/$ARCH
done

echo "Calculating sha256sums"
for TAR in $BASEDIR/dist/*.tar.gz
do
    sha256sum $TAR | awk '{ print $1 }' > $TAR.sha256sum
done

rm -rf $BASEDIR/dist/linux

echo "The following has been created:"
tree $BASEDIR/dist
