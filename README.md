# Playground in golang-based web frameworks

## Preparation

1. [Install Task](https://taskfile.dev/installation/)
2. [Install pre-commit](https://pre-commit.com/#install)
3. Install pre-commit hooks, pull docker images and build containers:

    ```sh
    task init
    ```

4. Start the application:

    ```sh
    task up
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
$ docker compose --profile=* down
```

Stopping and removing local state:

```console
$ docker compose --profile=* down --volumes --remove-orphans
```

## Testing

Running backend tests

```console
$ docker compose run backend-test
```

## Updating database migrations

```console
$ docker compose run atlas-make-migrations
```

## Backend container

Configuration environment variables:

- `DATABASE_DSN` - (required) postgres connection string
- `HOST` - bind interface. Default: `0.0.0.0`
- `POST` - bind port. Default: `8000`
