default: run

run file = "README.md":
  go build \
    -ldflags='-s -w' \
    -o mdv \
    .

  ./mdv {{ file }}

ci:
  ./.ci.sh
