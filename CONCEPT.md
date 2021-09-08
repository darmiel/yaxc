## CLI
* **gpg** signature check
* survey (`-m`)

**GLOBAL FLAGS:**
* `anywhere`
* `key` (whitelist key)

---

### 🚀 `push`
merge **paste** into **push**

**FLAGS:**
* `anywhere`*>
* `secret`>
* `ttl`
* `base64`
* `hash`
* `gpg-sign`

**OTHER:**
* from pipe
* from stdin
* from file
* from web (proxy?)

---

### 👈 `pull`
merge **get** into **pull**

**FLAGS:**
* `anywhere`*>
* `secret`>
* `base64`
* `gpg-sign`

**OTHER:**
* to stdout
* to file

---

### 👀 `watch`
* compare last updated
  * store DateTime

**FLAGS:**
* `anywhere`*>
* `secret`>
* `check-delay`
* `base64`
* `gpg-sign`
* `no-pull`
* `no-push`

## Server
* whitelist-key database

**PARAMETERS:**
* `base64`

**FLAGS:**
* `enable-whitelist-key`