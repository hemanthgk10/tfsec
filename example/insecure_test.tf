# Test fixture for COD-7216: intentionally insecure IaC used to trigger new
# FortiCNAPP IaC (Infrastructure as Code) violations on a pull request so the
# status-check deep-link can be validated on DEV5. Do NOT deploy.

# SSH open to the entire internet (0.0.0.0/0 on port 22) — critical IaC violation.
resource "aws_security_group" "cod7216_open_ssh" {
  name        = "cod7216-open-ssh"
  description = "Test SG that allows SSH from anywhere"



  ingress {
    description = "SSH from anywhere"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Public, unencrypted S3 bucket — should flag public-access + encryption violations.
resource "aws_s3_bucket" "cod7216_public_bucket" {
  bucket = "cod7216-public-test-bucket"
  acl    = "public-read"
}

# COD-7216 rescan trigger
