package test

import (
  "fmt"
  "strings"
  "testing"

  "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/gruntwork-io/terratest/modules/random"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
)

func TestS3Bucket(t *testing.T) {
	awsRegion := "us-east-1"

terraformOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
	TerraformDir: "C:\\Users\\Dra. Marotte Romina\\Desktop\\Terraform\\Milagros",

Vars: map[string]interface{}{
	"bucket_name": fmt.Sprintf("-%v", strings.ToLower(random.UniqueId())),
	  },

EnvVars: map[string]string{
	"AWS_DEFAULT_REGION": "us-east-1",
  },
})

defer terraform.Destroy(t, terraformOpts)
terraform.InitAndApply(t, terraformOpts)
bucketID := terraform.Output(t, terraformOpts, "bucket_id")
actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
assert.Equal(t, "Enabled", actualStatus)
}