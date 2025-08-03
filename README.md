# Docker Projects 🐳
Practical Docker projects demonstrating containerization, networking, and multi-stage builds.

```bash
git clone https://github.com/Bharathi-vbr/Docker-Projects.git
cd Docker-Projects/project-name
docker-compose up -d
```

# 🐳 Docker Essentials
Docker packages applications into portable containers that run consistently anywhere.

- Image - Application blueprint
- Container - Running application instance
- Dockerfile - Build instructions
- Compose - Multi-container orchestration
  
```bash
# 🛠️ Docker Commands
Images & Containers  
# Build image
docker build -t myapp .

# Run container (detached, port mapping)
docker run -d -p 8080:80 myapp

# List running containers
docker ps

# Stop container
docker stop <container>

# Remove container
docker rm <container>

# List images
docker images

# Remove image
docker rmi <image>
```
# Container Operations
```bash
# View container logs
docker logs <container>

# Access container shell
docker exec -it <container> bash

# Copy files to container
docker cp file.txt <container>:/app

# View resource usage
docker stats
```
# Docker Compose
```bash
# Start all services (detached)
docker-compose up -d

# Follow logs
docker-compose logs -f

# Stop and remove containers
docker-compose down

# Stop and remove volumes
docker-compose down -v

# Rebuild images
docker-compose build
```
# 🌐 Docker Networks
Enable container communication and isolation.
```bash
# List networks
docker network ls

# Create custom network
docker network create mynet

# Connect container to network
docker network connect mynet web
```
# Network Communication:
```bash
# Containers on same network communicate by name
docker run --network mynet --name api node-app
docker run --network mynet --name db postgres
# 'api' container can reach 'db' at hostname 'db'
```
# 🏗️ Multi-Stage Builds
Create smaller, secure production images.
```bash
# Build stage
FROM node:16 AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install && npm run build

# Production stage  
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
```
