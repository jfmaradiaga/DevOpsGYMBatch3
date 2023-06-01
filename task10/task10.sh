#!/bin/bash
#set -e

# go to the git repo path
cd /home/jmaradiaga/DevOps/DevOpsGYMBatch3/

# Pull the lastest changes
echo "Pulling the latest changes from the remote repository..."
ssh-agent bash -c 'ssh-add /home/jmaradiaga/.ssh/id_rsa; git pull'

# Install Docker
sudo apt-get update
sudo apt-get install -y docker.io

# Install Ansible
sudo apt-get install -y ansible

# Install Go
sudo apt-get install -y golang

# Compile the Go application
cd /home/jmaradiaga/DevOps/DevOpsGYMBatch3/
go build -o /home/jmaradiaga/DevOps/DevOpsGYMBatch3/hello /home/jmaradiaga/DevOps/DevOpsGYMBatch3/hello.go

# Create/build the Docker image
sudo docker build -t jfmaradiaga/helloworld:latest /home/jmaradiaga/DevOps/DevOpsGYMBatch3/

# Log in to Docker Hub or your desired Docker registry
echo "Logging in to the Docker registry..."
docker login -u jfmaradiaga -p dckr_pat_dp-A2vXuxznjjdyUWqiBnLlV2yU

# Push the Docker image to the Docker repository
sudo docker push jfmaradiaga/helloworld:latest

# Add changes to the local repository
echo "Adding changes to the local repository..."
git add .

# Commit the changes
echo "Committing changes..."
git commit -m "Update Docker image"

# Push the changes to the remote repository
echo "Pushing changes to the remote repository..."
ssh-agent bash -c 'ssh-add /home/jmaradiaga/.ssh/id_rsa; git push'

echo "Script execution completed."
