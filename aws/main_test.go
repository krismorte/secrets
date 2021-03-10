package aws

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func Test_getSecret(t *testing.T) {
	validSecretID := os.Getenv("validSecretID")
	if GetSecret(validSecretID) == nil {
		t.Error("Failed to get aws secret with aws credentials")
	}
}

func Test_getSecretPanic(t *testing.T) {
	invalidSecretID := "wrong-name"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	GetSecret(invalidSecretID)
}

func Test_getSecretLocally(t *testing.T) {
	validSecretID := "local-secret"
	localURL := os.Getenv("localURL")
	awsRegion := os.Getenv("region")
	if GetSecretLocally(validSecretID, awsRegion, localURL) == nil {
		t.Error("Failed to get aws secret locally with aws credentials")
	}
}

func Test_getSecretWithCredential(t *testing.T) {
	validSecretID := os.Getenv("validSecretID")
	testCredential := Credential{
		Region:          os.Getenv("region"),
		AccessKeyID:     os.Getenv("aws_access_key_id"),
		SecretAccessKey: os.Getenv("aws_secret_access_key"),
	}
	if GetSecretWithCredential(testCredential, validSecretID) == nil {
		t.Error("Failed to get aws secret with aws credentials")
	}
}

func Test_getSecretWithCredentialPanic(t *testing.T) {
	invalidSecretID := "wrong-name"
	testCredential := Credential{
		Region:          os.Getenv("region"),
		AccessKeyID:     os.Getenv("aws_access_key_id"),
		SecretAccessKey: os.Getenv("aws_secret_access_key"),
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	GetSecretWithCredential(testCredential, invalidSecretID)
}
