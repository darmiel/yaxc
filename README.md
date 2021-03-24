# YAxC
Yet Another Cross Clipboard
> Allan, please add details!

## Server
```bash
Run the YAxC server

Usage:
  yaxc serve [flags]

Flags:
  -b, --bind string            Bind-Address (default ":1332")
  -t, --default-ttl duration   Default TTL (default 1m0s)
  -h, --help                   help for serve
  -x, --max-body-length int    Max Body Length (default 1024)
  -s, --max-ttl duration       Max TTL (default 5m0s)
  -l, --min-ttl duration       Min TTL (default 5s)
  -r, --redis-addr string      Redis Address
      --redis-db int           Redis Database
      --redis-pass string      Redis Password
      --redis-prefix string    Redis Prefix (default "yaxc::")

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server
```