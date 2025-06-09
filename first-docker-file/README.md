# Simple Python Docker App

This is a basic Flask app containerized with Docker.

## Build the Docker Image

```bash
docker build -t simple-python-docker-app .
docker run -p 5000:5000 simple-python-docker-app

Visit http://localhost:5000 in your browser.
