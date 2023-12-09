### Steps

## Step 1: Making the Image
    docker build -t card_api:v1 .

## Step 2: Making the Container
    docker run --name card_api_container -p 8029:8029 card_api:v1

## Step 3: Deleting the Image and Container once it works 
    docker system prune 

## Step 4: Making the Docker compose file
    docker-compose up --detach

## Step 5: Stopping the Docker Compose container
    docker-compose down

