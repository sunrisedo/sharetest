PROGRAM="ShareTest"
VERSION="v0.0.1"
BUILD=`date +%FT%T%z`
COMMIT_SHA1=`git rev-parse HEAD`
LDFLAGS="-X main.PROGRAM=${PROGRAM} -X main.VERSION=${VERSION} -X main.BUILD=${BUILD} -X main.COMMIT_SHA1=${COMMIT_SHA1}"
# echo "VERSION:${VERSION}"
# echo "BUILD:${BUILD}"
# echo "COMMIT_SHA1:${COMMIT_SHA1}"
# echo "LDFLAGS:${LDFLAGS}"
go build -ldflags "${LDFLAGS}" -o sharetest