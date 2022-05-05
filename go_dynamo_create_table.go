package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main(){
	fmt.Println("Starting...")

	table 	:= flag.String("table","","")
	model 	:= flag.String("model","1","")

	flag.Parse()
	fmt.Printf("table: %s model: %s  \n", *table, *model)

	aws_region 		:= getEnvVar("AWS_REGION")
	aws_access_id 	:= getEnvVar("AWS_ACCESS_ID")
	aws_access_secret := getEnvVar("AWS_ACCESS_SECRET")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials( aws_access_id , aws_access_secret , ""),},
	)
	if err != nil {
		fmt.Println("Erro Create aws Session: ",err.Error())
		os.Exit(1)
	}

	svc := dynamodb.New(sess)
	createTable(svc, table, model)

	fmt.Println("Done...")
}

func getEnvVar(key string) string {
	fmt.Printf("Loading enviroment variable %s .. \n", key)
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file %s \n", err)
		os.Exit(1)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Printf("Invalid type \n")
		os.Exit(1)
	}
	return value
}

func createTable(svc *dynamodb.DynamoDB, tableName *string, model *string) {
	fmt.Printf("Creating table ( %s ) ... \n", *tableName)

	var input =  &dynamodb.CreateTableInput{}

	var pk_name = *tableName + "_id"
	var sk_name =  "issuer_id"
	var gsi_name =  "pk_tenant_id"

	if *model == "1" {
		input = &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String(pk_name),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String(sk_name),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String(pk_name),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String(sk_name),
					KeyType:       aws.String("RANGE"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			TableName: aws.String(*tableName),
		}
	} else if *model == "2" {
		input = &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String(pk_name),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String(sk_name),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String(gsi_name),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String(pk_name),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String(sk_name),
					KeyType:       aws.String("RANGE"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			TableName: aws.String(*tableName),
			GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
				{
					IndexName: aws.String("pk_gsi"),
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String(gsi_name),
							KeyType:       aws.String("HASH"),
						},
						{
							AttributeName: aws.String(sk_name),
							KeyType:       aws.String("RANGE"),
						},
					},
					Projection: &dynamodb.Projection{
						ProjectionType: aws.String("ALL"),
					},
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(1),
						WriteCapacityUnits: aws.Int64(1),
					},
				},
			},
		}
	} else if *model == "3" {
		input = &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String(pk_name),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String(pk_name),
					KeyType:       aws.String("HASH"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			TableName: aws.String(*tableName),
		}
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Error create table: ",err.Error())
		os.Exit(1)
	}
	
	fmt.Println("Success on create the table", *tableName)
}