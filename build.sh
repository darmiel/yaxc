#!/bin/bash

# clear old builds
rm bin/*

NAME="server-only"
TAGS="server"

# ignored platforms
IGNORED=("aix" "android" "ios" "plan9" "js")

# get possible build configurations
POSSIBLE=$(go tool dist list)

for poss in $POSSIBLE; do

  IFS="/"
  read -r -a poss_array <<< "$poss"
  PLATFORM="${poss_array[0]}"
  ARCH="${poss_array[1]}"

  if [[ " ${IGNORED[@]} " =~ " ${PLATFORM} " ]]; then
    continue
  fi

  OUTF="./bin/yaxc-${NAME}-${PLATFORM}-${ARCH}"
  if [[ "${PLATFORM}" == "windows" ]]; then
    OUTF="${OUTF}.exe"
  fi

  echo ""
  echo "🔨 Building for ${PLATFORM}(${ARCH}) ..."
  echo "   👉 ${OUTF}"

  # build
  GOOS=${PLATFORM} GOARCH=${ARCH} go build \
    -ldflags="-s -w" \
    -o "${OUTF}" \
    ./main.go
  echo "    Done!"

done
