# DChain Local Setup Guide

This guide mainly follows the one provided by [Forbole docs](https://docs.bigdipper.live/cosmos-based/parser/overview).
However, we have made some modification to easily setup for DChain.

## Set up postgres and hasura locally.

```sh

# this starts the database and hasaru
# database will start with the schema already
# i.e. do not need to do https://docs.bigdipper.live/cosmos-based/parser/database
docker compose up -d

```

## Set up callisto

Install callisto binary

```sh
make install

# check version is latest git hash
callisto version

callisto init
```

Instead of using the default value from `callisto init`,
you should have receieved pre-config files, you can save it here in `dchain-config/`.
Easiest way is copy the config folder to the default location.

```sh
cp -r dchain-config/ ~/.callisto
```

Now you can parse the genesis file and start callisto

```sh
callisto parse genesis-file
callisto start
```

Now you should see it starting to process blocks.
To check the local db
```sh
PGPASSWORD=<Postgres password in .env> psql -U dfoundation -h localhost -p 5432 -d bdjuno --set=sslmode=disable

SELECT height, hash, timestamp, num_txs, total_gas
FROM block
ORDER BY height DESC
LIMIT 1;
```

It shold be the latest block height which will increase as you repeat the command.

## Setup Hasura

The `docker-compose.yml` already has the hasura service running.

```sh
# apply the metadata and migrations
cd ./hasura/

hasura metadata apply --endpoint http://0.0.0.0:8080 --admin-secret <hasura secret in .env>

# quick test
curl http://localhost:3000/account_balance \                                                          130 ↵ belsy@MacBook-Pro-7
    --data '{ "input": { "address": "dchain1p45t0prxw0ylf6av03dt53z3hh4cznga3676e5"} }'
```
