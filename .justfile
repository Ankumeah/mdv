default: run

run *opts:
  CGO_ENABLED=0 \
  go build \
    -ldflags='-s -w' \
    -o mdv \
    .

  ./mdv {{ opts }}

ci:
  ./.ci.sh
