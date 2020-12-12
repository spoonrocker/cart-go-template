# Cart API

This is a REST API for creating carts and products using Golang.

## Development deployment

First, it is needed to create a .env file, there is a .env.example file in order to know the right environment variables.  
Once the environment variables are set at the .env file, set the ports on the docker-compose.yml file.

```bash
docker compose up -d
```

## Production deployment

First, it is needed to build the Docker image
```bash
docker build . -t cart-api
```

Then, just create the Docker container
```bash
docker run --publish="3000:3000" --rm -it --env PORT=3000 --env DATABASE=db.db cart-api
```
