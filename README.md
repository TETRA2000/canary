# Canary

[![CircleCI](https://circleci.com/gh/TETRA2000/canary.svg?style=svg&circle-token=fb895ba698b6dd430d9b97a171a3184e9b02537e)](https://circleci.com/gh/TETRA2000/canary)

# Requirements

* Linux/Mac(Docker for Mac)

## Run

```bash
docker build -t canary .
docker run -d -v $PWD/data:/opt/canary/data -v /var/run/docker.sock:/var/run/docker.sock canary
```

If you building for development, adding `--build-arg USE_HOST_VENDOR=1` will keep files in `vendor/` and skip `dep ensure`.


## Test
```bash
./test.sh
```

## Directory structure

```
├── canary
│   ├── data/       -- App data
│   ├── plugins/    -- Default plugins
│   ├── scripts/    -- Development tools
│   └── ...
```
