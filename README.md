# checkout

Get checkouted branches from the `reflog` and create a fuzzy-finding interface

## Install

### manually

Download the pre-compiled binaries from the [OSS releases page](github.com/wesleimp/checkout/releases) and copy them to the desired location.

### Compiling from source

If you just want to build from source for whatever reason, follow these steps:

**clone**

```sh
git clone git@github.com/wesleimp/checkout.git
cd checkout
```

**get the dependencies**

```sh
make deps # or go mod install
```

**install**

```sh
make install
```

**verify it works**

```sh
checkout -v
```
