# address-go
an address api based on golang

# Get Start

if you want to run this project, these software are required
1. [Docker](https://docs.docker.com/install/)
> if you use the Mac or Windows, please install `Docker Desktop`;
> if you use the Linux, please download [docker-compose](https://docs.docker.com/compose/install/) separately.
2. [Golang](https://golang.org/dl/)

## Install
```bash
git clone https://github.com/damingerdai/address-go.git
```

## Init

### Create Volume

```bash
docker volume create --name=address-go
```

### Create Network

```bash
docker network create address-networks
```

### Docker Compose

run for mysql
```bash
docker compose up mysql
```

### Init data

in `data` folder, exec these sql files to init data

# Run

```bash
make run
```

# Bazel

update repositories

```bash
bazel run //:gazelle
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

build

```
bazel build //cmd/address:address
```

# Data
i use this [data source](https://github.com/wecatch/china_regions)

# bazel

```
 bazel run //:gazelle
 bazel run //:gazelle -- update-repos -from_file=go.mod
 bazel build //cmd/address:address
```

# Roadmap

## 2019.12.30
1. use [golang-standards](https://github.com/golang-standards/project-layout) to refactor the project layout [x]

# Owner

[@大明二代](https://github.com/damingerdai)

# License

[MIT](LICENSE) © 大明二代