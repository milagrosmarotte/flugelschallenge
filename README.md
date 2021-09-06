# flugelschallenge
Hi! This is the project I´ve been developing for Flugel´s first stage process.
This project consists of writing Terraform code to create an AWS S3 bucket, testing it with Terratest, creating a GitHub actions pipeline of this Terraform code, and publishing everything to a GitHub directory. 

Inside of "s3.tf" file, you will find the terraform code I wrote in order to create the s3 bucket. 
Inside of "bucket-test.go" you will find the terratest code I wrote in order to test this bucket.

If you would like to run the automation, you must follow the following steps:
1. Get the go dependencies.
2. Run the test using "go test -v" command on the terminal.
3. If the test fails, you will get a "suspended" message. However, If the test passes you will get an "enabled" message on the console.

Thanks for checking my code! 
Any feedback is welcome :)
