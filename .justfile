default: run

run file = "example.md":
  go build \
    -ldflags='-s -w' \
    -o mdv \
    .

  ./mdv {{ file }}

ci:
  ./.ci.sh
