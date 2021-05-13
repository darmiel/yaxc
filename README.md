# YAxC
Yet Another Cross Clipboard
> Allan, please add details!

## API
✨ API Docs moved to: [api.yaxc.d2a.io](https://api.yaxc.d2a.io)

## Demo
https://youtu.be/OVpH70byKRQ

## CLI
### Send Clipboard (one-time)
**HELP**
```
Usage:
  yaxc force push [flags]

Flags:
  -h, --help   help for push

Global Flags:
  -a, --anywhere string   Anywhere Path
  -b, --base64            Use Base64
      --config string     config file (default is $HOME/.yaxc.yaml)
  -S, --hide-secret       Hide Secret
  -U, --hide-url          Hide URL
  -s, --secret string     Encryption Key
      --server string     URL of API-Server (default "https://yaxc.d2a.io")
```

**EXAMPLE**
```bash
$ yaxc force push -a mypath [-s mypassword]
# INFO | Sent -> hello wor... -> /mypath
# INFO | 🔐 mypassword
# DBUG | URL: https://yaxc.d2a.io/mypath?secret=mypassword
```

### Receive Data And Paste To Clipboard (one-time)
**HELP**
```
Usage:
  yaxc force pull [flags]

Flags:
  -h, --help   help for pull

Global Flags:
  -a, --anywhere string   Anywhere Path
  -b, --base64            Use Base64
      --config string     config file (default is $HOME/.yaxc.yaml)
  -S, --hide-secret       Hide Secret
  -U, --hide-url          Hide URL
  -s, --secret string     Encryption Key
      --server string     URL of API-Server (default "https://yaxc.d2a.io")
```

**EXAMPLE**
```bash
$ yaxc force pull -a mypath [-s mypassword]
# INFO | Read <- hello wor...
```

### Receive Data And Output To Stdin (one-time)
**HELP**
```
Usage:
  yaxc get [flags]

Flags:
  -a, --anywhere string     Path (Anywhere)
  -h, --help                help for get
  -s, --passphrase string   Encryption Key

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server (default "https://yaxc.d2a.io")

```

**EXAMPLE**
```bash
$ yaxc get -a mypath [-s mypassword]
# Hello world!
```

### Watch Clipboard (Clipboard Sync)
**HELP**
```
Usage:
  yaxc watch [flags]

Flags:
  -a, --anywhere string     Path (Anywhere)
  -b, --base64              Use Base64?
  -h, --help                help for watch
      --ignore-client       Ignore Client Updates
      --ignore-server       Ignore Server Updates
  -s, --passphrase string   Encryption Key

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server (default "https://yaxc.d2a.io")
```

**EXAMPLE**
```bash
$ yaxc watch -a mypath [-b -s mypassword]
# INFO | Starting Watchers:
# INFO | * Server -> Client
# INFO | * Server <- Client
# INFO | Started clipboard-watcher. Press CTRL-C to stop.
# UPDT | Server <- Hello Wor...
```

### Host Your Own Server
**HELP**
```
Usage:
  yaxc serve [flags]

Flags:
  -b, --bind string                 Bind-Address (default ":1332")
  -t, --default-ttl duration        Default TTL (default 1m0s)
  -e, --enable-encryption           Enable Encryption (default true)
  -h, --help                        help for serve
  -x, --max-body-length int         Max Body Length (default 8192)
  -s, --max-ttl duration            Max TTL (default 1h0m0s)
  -l, --min-ttl duration            Min TTL (default 5s)
      --proxy-header string         Proxy Header
  -r, --redis-addr string           Redis Address
      --redis-db int                Redis Database
      --redis-pass string           Redis Password
      --redis-prefix-hash string    Redis Prefix (Hash) (default "yaxc::hash::")
      --redis-prefix-value string   Redis Prefix (Value) (default "yaxc::val::")

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server (default "https://yaxc.d2a.io")
```

**EXAMPLE**
```bash
$ yaxc serve -b :80
# INFO | Started clipboard-server. Press CTRL-C to stop.
```

---

```bash
alias yfll='yaxc force pull -a anywhere'
alias yfsh='yaxc force push -a anywhere'
```
