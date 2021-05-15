<!-- Logo -->
<p align="center">
  <img src="./assets/YAxC-Filled@128px.png" alt="yaxc">
</p>

<!-- Header -->
<h1 align="center">YAxC</h1>
<p align="center">
  <strong>Y</strong>et 
  <strong>A</strong>nother 
  <i>Cross</i>
  <strong>C</strong>lipboard
</p>

<!-- Links -->
<p align="center">
  [
  <a href="https://github.com/darmiel/yaxc/releases">📦 Download</a> |
  <a href="https://api.yaxc.d2a.io">📚 API-Docs</a> |
  <a href="https://youtu.be/OVpH70byKRQ" target="_blank">🎥 Demo</a> |
  <a href="#client">⌨️ Usage</a>
  ]
</p>

## Introduction
YAxC is my attempt to develop a cross-platform clipboard that is as simple as possible. YAxC consists of two components:

### Server
The server was kept very minimalistic and simple: 

The server accepts a `POST` request (`text/plain`) to any path `/{anywhere}` (except `/` and `/hash`) and stores the sent data there for 5 minutes by default, but this can be changed with the `ttl` query-parameter. An MD5 hash is then generated and can be retrieved at `/hash/{anywhere}`. 

This hash is used to see if the data has changed on the server. This hash can also be specified during the upload, e.g. if you want to use a different hash method.

For more information, see the 📚 API docs [here](https://api.yaxc.d2a.io)

#### 🖥 Host Your Own Server
**EXAMPLE**
```bash
$ yaxc serve -b :80
# INFO | Started clipboard-server. Press CTRL-C to stop.
```

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
```

### Client
The client is a command line interface that provides the following functions:

#### 👉 Send Clipboard (one-time)
**EXAMPLE**
```bash
$ yaxc force push -a mypath [-s mypassword]
# INFO | Sent -> hello wor... -> /mypath
# INFO | 🔐 mypassword
# DBUG | URL: https://yaxc.d2a.io/mypath?secret=mypassword
```

**HELP**
```
Usage:
  yaxc force push [flags]
  
Global Flags:
  -a, --anywhere string   Anywhere Path
  -b, --base64            Use Base64
      --config string     config file (default is $HOME/.yaxc.yaml)
  -S, --hide-secret       Hide Secret
  -U, --hide-url          Hide URL
  -s, --secret string     Encryption Key
      --server string     URL of API-Server (default "https://yaxc.d2a.io")
```

#### 👈 Receive Data And Paste To Clipboard (one-time)
**EXAMPLE**
```bash
$ yaxc force pull -a mypath [-s mypassword]
# INFO | Read <- hello wor...
```

**HELP**
```
Usage:
  yaxc force pull [flags]

Global Flags:
  -a, --anywhere string   Anywhere Path
  -b, --base64            Use Base64
      --config string     config file (default is $HOME/.yaxc.yaml)
  -S, --hide-secret       Hide Secret
  -U, --hide-url          Hide URL
  -s, --secret string     Encryption Key
      --server string     URL of API-Server (default "https://yaxc.d2a.io")
```

#### 👈 Receive data and output them (one-time)
**EXAMPLE**
```bash
$ yaxc get -a mypath [-s mypassword]
# Hello world!
```

**HELP**
```
Usage:
  yaxc get [flags]

Flags:
  -a, --anywhere string     Path (Anywhere)
  -h, --help                help for get
  -s, --passphrase string   Encryption Key

```

#### ♻️ Watch Clipboard (Clipboard Sync)
**EXAMPLE**
```bash
$ yaxc watch -a mypath [-b -s mypassword]
INFO | Started clipboard-watcher. Press CTRL-C to stop.
UPDT | Server <- Hello Wor...
```

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
```

---

## Useful aliases
```bash
# PUSH
# push clipboard contents (one-time) to /anywhere
alias yfsh='yaxc force push -a anywhere'

# push clipboard contents (one-time) encrypted to /anywhere
alias yfshs='yaxc force push -a anywhere -s secret'

# PULL
# pull clipboard contents (one-time) to /anywhere
alias yfll='yaxc force pull -a anywhere'

# pull encrypted clipboard contents (one-time) to /anywhere
alias yflls='yaxc force pull -a anywhere -s secret'
```
