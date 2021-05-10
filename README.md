# YAxC
Yet Another Cross Clipboard
> Allan, please add details!

## Demo
https://youtu.be/OVpH70byKRQ

```bash
alias yfll='yaxc force pull -a daniel-temp'
alias yfsh='yaxc force push -a daniel-temp'
```

## Server

### Set Data
Just make a POST request to any path:

**POST** `/hi`
```
Hello World!
```

#### TTL
By default, the data is kept for 5 minutes. This TTL can be changed via the `ttl`-parameter.

**POST** `/hi?ttl=1m30s`
```
Hello World!
```

#### Encryption
By default, the data is not encrypted. 
**It is not recommended to encrypt the data on server side. The data should always be encrypted on the client side.**

However, if this is not possible, the `secret`-parameter can be used to specify a password with which the data should be encrypted.

**POST** `/hi?secret=s3cr3tp455w0rd`
```
Hello World!
```
**Produces:**
```
gwttKS3Q2l0+YR+jQF/02u3fNVmMIcVOTNSGD5vWfrYTtH8adt8r
```

### Get Data
**GET** `/hi`
```
Hello World!
```

#### Encryption
If the data has been encrypted and should be decrypted on the server side (**which is not recommended**), the "password" can be passed via the `secret`-parameter.
**GET** `/hi`
```
gwttKS3Q2l0+YR+jQF/02u3fNVmMIcVOTNSGD5vWfrYTtH8adt8r
```

**GET** `/hi?secret=s3cr3tp455w0rd`
```
Hello World!
```


### CLI
```bash
Run the YAxC server

Usage:
  yaxc serve [flags]

Flags:
  -b, --bind string                 Bind-Address (default ":1332")
  -t, --default-ttl duration        Default TTL (default 1m0s)
  -e, --enable-encryption           Enable Encryption (default true)
  -h, --help                        help for serve
  -x, --max-body-length int         Max Body Length (default 1024)
  -s, --max-ttl duration            Max TTL (default 5m0s)
  -l, --min-ttl duration            Min TTL (default 5s)
      --proxy-header string         Proxy Header
  -r, --redis-addr string           Redis Address
      --redis-db int                Redis Database
      --redis-pass string           Redis Password
      --redis-prefix-hash string    Redis Prefix (Hash) (default "yaxc::hash::")
      --redis-prefix-value string   Redis Prefix (Value) (default "yaxc::val::")

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server

```

## Client
### Watch
```bash
Watch Clipboard

Usage:
  yaxc watch [flags]

Flags:
  -a, --anywhere string     Path (Anywhere)
  -h, --help                help for watch
      --ignore-client       Ignore Client Updates
      --ignore-server       Ignore Server Updates
  -s, --passphrase string   Encryption Key

Global Flags:
      --config string   config file (default is $HOME/.yaxc.yaml)
      --server string   URL of API-Server (default "https://yaxc.d2a.io")
```
