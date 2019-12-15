# address-go
an address api based on golang

# Get Start

## Install
```bash
git clone https://github.com/damingerdai/address-go.git
```

## Init

### Create Volume

```bash
docker volume create --name=address-go
```

### Docker Compose

run for mysql
```bash
docker-compose up
```

### Init data

in `data` folder, exec these sql files to init data

# Run

```bash
make run
```

# Data
i use this [data source](https://github.com/wecatch/china_regions)

# Owner

[@大明二代](https://github.com/damingerdai)

# License

[MIT](LICENSE) © 大明二代