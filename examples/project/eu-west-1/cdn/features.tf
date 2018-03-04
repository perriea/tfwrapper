resource "aws_s3_bucket" "state-bucket" {
  bucket = "tfwrapper-${terraform.workspace}"
  acl    = "private"

  tags {
    Name        = "tfwrapper-${terraform.workspace}-state"
    Environment = "${terraform.workspace}"
  }
}
