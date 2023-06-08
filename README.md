# DevOpsGYMBatch3
DevOpsGYMBatch3 Repository for Julio Maradiaga


#Task 10
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"

this will generate private and public key, the public key needs to be added to ~/.ssh/authorized_keys in the EC2 instance and the private key to the SSH_PRIVATE_KEY secret