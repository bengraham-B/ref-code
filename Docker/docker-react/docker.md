# Docker steps

## Step 1: Create Docker file in client directory
dockerfile \

    FROM node:17-alpine

    WORKDIR /app

    COPY package.json .

    RUN npm install

    COPY . . 

    EXPOSE 3000

    CMD ["npm", "start"]

## Step 2: Create Image
    docker build -t react_10dec:v1.0.0 .

## Step 3: Show Docker Images
    docker images

## Step 4: Create Container
    docker run --name react_10dec_container -p 3000:3000 react_10dec:v1.0.0

## Step 5: Show Docker Containers
    docker ps -a

## Step 6: Start the Docker Container
    docker start react_10dec_container

## Step 7: DELETE IMAGE (Force)
    docker image rm react_10dec:v1.0.0 -f

## Step 8: Delete Container 
    docker container rm react_10dec_container

## Step 9: Stop Container 
    docker stop react_10dec_container

## Step 10: Start Container 
    docker start react_10dec_container

---

# Docker Compose 

## Step 11: Build Container with Docker Compose
    docker-compose up --build

## Step 12: Run the docker container mad with docker compose (Detached mode)
    docker-compose up -d

## Step 13: Show which docker containers are running
    docker ps

## Step 14: Stop the docker container mad with docker compose (Detached mode)
    docker-compose down

---

# Push to Dockerhub
##  Step 15: Create new image with username in the image
    docker build -t goosecpt/react_10dec:v1.0.0 .

##  Step 16: Show images
    docker images

##  Step 17: Login to Dockerhub locally
    docker login

##  Step 18: Push image to docker
    docker push goosecpt/react_10dec:v1.0.0

## Step 19: Pull from Dockerhub
    docker pull goosecpt/react_10dec:v1.0.0


