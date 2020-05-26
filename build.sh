#! /bin/sh

NAME=kb
DEST=./dist
PLATFORMS="darwin/amd64 linux/amd64 windows/amd64"

for platform in $PLATFORMS; do
  export GOOS="${platform%/*}"
  export GOARCH="${platform#*/}"
  echo "Building for ${GOOS}/${GOARCH}..."
  out="$DEST/${NAME}_${GOOS}_${GOARCH}_$(git describe --tags --always --dirty)"
  post="gzip -v $out"
  if [ "$GOOS" = "windows" ]; then
    out="${out}.exe"
    post="zip --move ${out%.exe}.zip ${out}"
  fi
  go build -o "$out"
  $post
done
