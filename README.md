# articles-server

### Local development

1. Build an image and run docker container:  
`$ docker compose up --build`

2. Apply database migrations if necessary:  
`$ goose -dir ./migrations postgres "postgres://postgres:secret_password@localhost:5432/articles?sslmode=disable" up`