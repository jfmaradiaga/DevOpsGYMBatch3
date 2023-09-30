resource "aws_ssm_parameter" "foo" {
  name  = "foo"
  type  = "String"
  value = "barr"
}

resource "aws_instance" "public_instance" {
  ami           = "ami-053b0d53c279acc90"  # Replace with the desired AMI ID
  instance_type = "t2.micro"     # Replace with the desired instance type
  key_name      = aws_key_pair.key_pair.key_name  # Specify the key pair to use

  tags = {
    Name = "task19"
  }
}