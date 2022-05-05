# snipet-aws-go-dynamo-create-table

How to use

1. Create a go module
go mod init github.com/snipet-aws-go-dynamo-create-table/main

2. Create a config.yaml with your secrets in the root
AWS_REGION: "AWS_REGION"
AWS_ACCESS_ID: "AWS_ACCESS_ID"
AWS_ACCESS_SECRET: "AWS_ACCESS_SECRET"

3. Run passing the table name and the model choice
3.1 Model 1 - NO GSI
3.2 Model 2 - WITH GSI
3.2 Model 3 - only PK

Ex:
go run . --table account --model 1

or

go run . --table InvoiceT --model 2

or

go run . --table InvoiceT --model 3