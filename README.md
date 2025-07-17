# movies-api
JSON API for retrieveing and managing info about movies.
practice from the book "Let's Go Further"

## how to run

current version can be ran with Makefile:
```bash
make stest
```
it runs server and `curl`s the first endpoint

## implemented

| N | Method | URL Pattern | Action |
| -- | ------ | ----------- | ------ |
| 1 | GET |  /v1/healthcheck |  show application health and version information |
| 2 | POST | /v1/movies | create a new movie |
| 3 | GET | /v1/movies/:id | show the details of a specific movie |

> *note*: for now  there's no movies data

## still to be done

| Method | URL Pattern | Action |
| ------ | ----------- | ------ |
| GET | /v1/movies | show the details of all movies |

