# Playground in golang-based web frameworks

## Preparation

Install pre-commit hooks:

```console
$ pre-commit install
```

The first time you commit you may need to run:

```console
$ GOFLAGS=-buildvcs=false pre-commit
```

## Starting and stopping the application

If starting from scratch:

```console
$ docker compose up --wait
```

Stream logs:

```console
$ docker compose logs -f backend
```

Run with latest images and dependencies:

```console
$ docker compose build --pull
$ docker compose up --wait
```

Open a database shell:

```console
$ docker compose run psql
```

Run production backend:

```console
$ docker compose --profile production up --build prod-backend
```

Stopping:

```console
$ docker compose down
```

Stopping and removing local state:

```console
$ docker compose down --volumes --remove-orphans
```

## Updating database migrations

```console
$ docker compose run atlas migrate diff --env gorm
```

## Backend container

Configuration environment variables:

- `DATABASE_DSN` - (required) postgres connection string
- `HOST` - bind interface. Default: `0.0.0.0`
- `POST` - bind port. Default: `8000`
