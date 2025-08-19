# Gorgi - backend

First, clone the repository.
```sh
git clone --recurse-submodules https://github.com/gorgichat/backend.git
# or
git clone --recurse-submodules git@github.com:gorgichat/backend.git

cd "backend"
```

## Or clone first, then fetch submodules.

```sh
git clone https://github.com/gorgichat/backend.git
# or
git clone git@github.com:gorgichat/backend.git

cd "backend"
git submodule update --init --recursive
```


## Quick commands

- Update submodules and push changes:

```sh
git submodule foreach git pull
git add .
git commit -m "msg"
git push
```
