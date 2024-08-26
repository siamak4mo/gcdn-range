# cdn-range
Golang cdn-search implementation
It's a tool to download IP ranges used by CDNs
currently supported: AWS, Cloudflare, Akamai, Incapsula, Fastly, ArvanCloud

---

# Compilation

#### go
``` shell
go build -o cdn-range main.go
```

#### Makefile
``` shell
# to compile with gccgo (smaller binary size)
make

# without gccgo
make no_gcc

# to run tests
make test
```


# Usage
``` shell
# verbose
./cdn-range -v

# write to file
./cdn-range -o /tmp/out.txt

# output formats
./cdn-range -json  [also: -csv, -tsv]

# specific provider
./cdn-range -p aws,cloudflare
```
