resource "aws_s3_bucket" "tf_challenge" {
  bucket = "milagroschallenge"
  acl="private"
}


resource "aws_s3_bucket_object" "objects" {
for_each = fileset("C:\\Users\\Dra. Marotte Romina\\Desktop\\Terraform\\files", "*")
bucket = aws_s3_bucket.tf_challenge.id
key = each.value
source = "C:\\Users\\Dra. Marotte Romina\\Desktop\\Terraform\\files\\${each.value}"
etag = filemd5("C:\\Users\\Dra. Marotte Romina\\Desktop\\Terraform\\files\\${each.value}")
}
