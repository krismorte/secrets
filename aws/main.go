package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

//Credential is a basic AWS credential
type Credential struct {
	// AWS Region
	Region string
	// AWS Access key ID
	AccessKeyID string

	// AWS Secret Access Key
	SecretAccessKey string
}

func GetSecret(secretID string) interface{} {
	var secret interface{}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	svc := getSecretClient()

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(*result.SecretString), &secret)

	return secret
}

func GetSecretWithRegion(secretID string, region string) interface{} {
	var secret interface{}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	svc := getSecretClientWithRegion(region)

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(*result.SecretString), &secret)

	return secret
}

func GetSecretLocally(secretID string, region string, localURL string) interface{} {
	var secret interface{}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	svc := getSecretClientLocally(localURL, region)

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(*result.SecretString), &secret)

	return secret
}

func GetSecretWithCredential(credential Credential, secretID string) interface{} {
	var secret interface{}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}

	svc := getSecretClientWithCredential(credential)

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(*result.SecretString), &secret)

	return secret
}

func getSecretClient() *secretsmanager.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	if &cfg != nil {
		return secretsmanager.NewFromConfig(cfg)
	}
	return nil
}

func getSecretClientWithRegion(region string) *secretsmanager.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	if &cfg != nil {
		return secretsmanager.NewFromConfig(cfg)
	}

	return nil
}

func getSecretClientWithCredential(credential Credential) *secretsmanager.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(credential.AccessKeyID, credential.SecretAccessKey, "")),
		config.WithRegion(credential.Region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	if &cfg != nil {
		return secretsmanager.NewFromConfig(cfg)
	}

	return nil
}

func getSecretClientLocally(localURL, region string) *secretsmanager.Client {

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if service == secretsmanager.ServiceID && region == region {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           localURL,
				SigningRegion: region,
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	if &cfg != nil {
		return secretsmanager.NewFromConfig(cfg)
	}

	return nil
}
