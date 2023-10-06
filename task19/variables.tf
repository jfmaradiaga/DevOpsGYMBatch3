variable "region" {
  description = "The AWS region where resources will be created."
  default     = "us-east-1"
}

variable "public_key" {
  description = "The public SSH key to use for EC2 instances."
}

variable "private_key" {
  description = "The private SSH key to use for EC2 instances."
}

variable "key_name" {
  description = "The name of the SSH key pair for EC2 instances."
}
