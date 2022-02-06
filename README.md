# snipet-aws-go-dynamo-query

Dynamo queries in GO

How to use

1. Create a go module
go mod init github.com/snipet-aws-go-dynamo-create-query/main

2. Create a config.yaml with your secrets in the root
AWS_REGION: "AWS_REGION"
AWS_ACCESS_ID: "AWS_ACCESS_ID"
AWS_ACCESS_SECRET: "AWS_ACCESS_SECRET"

3. Load data
go run . --option load_invoice --table Invoice_Tenant

4. Query data only pk
aws dynamodb query --table-name Invoice_Tenant \
	--key-condition-expression "pk= :v1" \
	--expression-attribute-values '{ ":v1": {"S":"invoice-44"}}' \
	--return-consumed-capacity TOTAL

go run . --option query_invoice --table Invoice --pk invoice-47

4. Query data with pk and sk
aws dynamodb query --table-name Invoice_Tenant  \
	--key-condition-expression "pk = :v1 AND sk = :v2" \
	--expression-attribute-values '{ ":v1": {"S":"invoice-44"} , ":v2": {"S":"invoice-44"} }' \
	--return-consumed-capacity TOTAL

go run . --option query_invoice --table Invoice_Tenant --pk invoice-44 --sk invoice-44


aws dynamodb query --table-name Invoice_Tenant  \
	--key-condition-expression "pk = :v1 AND begins_with(sk, :v2)" \
	--expression-attribute-values '{ ":v1": {"S":"invoice-44"} , ":v2": {"S":"order"} }' \
	--return-consumed-capacity TOTAL


5. Query data through GSI

aws dynamodb query --table-name Invoice_Tenant  \
	--index-name pk_gsi \
	--key-condition-expression "pk_gsi = :v1" \
	--expression-attribute-values '{ ":v1": {"S":"tenant-0"} }' \
	--return-consumed-capacity TOTAL

go run . --option query_invoice_gsi --table InvoiceT --pk tenant-0

6. Query data through GSI and sk
aws dynamodb query --table-name InvoiceT  \
	--index-name pk_gsi \
	--key-condition-expression "pk_gsi = :v1 AND begins_with(sk, :v2)" \
	--expression-attribute-values '{ ":v1": {"S":"tenant-0"} , ":v2": {"S":"order"} }' \
	--return-consumed-capacity TOTAL
go run . --option query_invoice_gsi --table InvoiceT --pk tenant-0 --sk order

7. Query data through GSI
go run . --option update_transaction --table InvoiceT --pk invoice-47 --sk invoice-47 --order_id order-178 --amount 3.4