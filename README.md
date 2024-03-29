# Roba no hashi

From the German word "Eselsbrücke" meaning mnemonic.

This App allows users to lookup Radicals, Kanji and Vocabulary with mnemonics to aid learning.
The difference from something like WaniKani is that users can upload their own menmonics and share with a community.

The app currently runs on web here: [robanohashi.org](https://robanohashi.org)

API documentation: [api.robanohashi.org/docs/index.html](https://api.robanohashi.org/docs/index.html)

### Redis schema
![redis-schema](https://github.com/martingrzzler/robanohashi-backend/assets/54838006/376585fa-3f64-4d0d-a863-43280a278fe0)

### Authentication
Authentication is handled entirely by Firebase. This api only validates the JSON web token for endpoints that require it.

### Development

For development I suggest to start a redis docker container.
```bash
docker run -d \
  --name redis \
  -p 6379:6379 \
  -p 8001:8001 \
  -v "$(pwd)/data:/data" \
  redis/redis-stack:latest
```

#### Migrations
Migrations for the Redis instance are performed locally and the `dump.rdb` files are backed on `S3`.
Create a new package in `/cmd/<package>` for the migration code.

#### Docs
1. Download the `swag` cli.
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
2. Run `swag init` to generate the files

#### Indices
Run `go run ./cmd/indices` to recreate indices

#### Deployment
```bash
docker build -t martingrzzler/robanohashi-api:latest .
```
```bash
docker push martingrzzler/robanohashi-api:latest
```

On the production server update individual docker swarm services:
```bash
docker service update --image martingrzzler/robanohashi-api:latest robanohashi_api
```

```bash
docker service update robanohashi_redis --force
```

#### Test
Integration tests:

Make sure to start an empty redis instance.
```bash
go test ./persist
```

Unit tests:
```bash
go test ./internal/utils
```

