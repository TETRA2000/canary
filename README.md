# Canary

[![CircleCI](https://circleci.com/gh/TETRA2000/canary.svg?style=svg&circle-token=fb895ba698b6dd430d9b97a171a3184e9b02537e)](https://circleci.com/gh/TETRA2000/canary)

# Requirements

* Linux/Mac(Docker for Mac)

## Run

```bash
docker build -t canary .
docker run -d -v $PWD/data:/opt/canary/data -v /var/run/docker.sock:/var/run/docker.sock canary
```

## Test
```bash
./test.sh
```

## Directory structure

```
├── canary
│   ├── plugins/    -- Default plugins
│   ├── data/       -- App data
│   └── ...
```
