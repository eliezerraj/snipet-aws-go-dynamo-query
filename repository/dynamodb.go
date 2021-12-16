package dynamodb

import (
	"fmt"
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

)

type Repository struct {
	client dynamodbiface.DynamoDBAPI
	table  *string
}

func NewRepository(client dynamodbiface.DynamoDBAPI, table string) *Repository {
	return &Repository{client: client, table: aws.String(table)}
}

func (r *Repository) AddInvoice(ctx context.Context, inter interface{}) error {
	fmt.Printf("AddInvoice ->  %s  \n", inter)

	item, err := dynamodbattribute.MarshalMap(inter)
	if err != nil {
		fmt.Println("Erro marshalling: ",err.Error())
		return err
	}

	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: r.table,
		Item:      item,
	}})

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		return err
	}

	_, err = r.client.TransactWriteItemsWithContext(ctx, transaction)
	return err
}

func (r *Repository) QueryInvoice(ctx context.Context, pk string, sk string) (interface{}, error) {
	fmt.Printf("QueryInvoice \n")

	var keyCond expression.KeyConditionBuilder
	if len(sk) == 0 {
		keyCond = expression.Key("pk").Equal(expression.Value(pk))
	}else {
		keyCond = expression.KeyAnd(
			expression.Key("pk").Equal(expression.Value(pk)),
			expression.Key("sk").BeginsWith(sk),
		)
	}

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		return nil, err
	}

	key := &dynamodb.QueryInput{
		TableName:                 r.table,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	fmt.Println(key)

	result, err := r.client.QueryWithContext(ctx, key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) QueryInvoiceGsi(ctx context.Context, pk string, sk string) (interface{}, error) {
	fmt.Printf("QueryInvoiceGsi \n")

	var keyCond expression.KeyConditionBuilder
	if len(sk) > 0 {
		keyCond = expression.KeyAnd(
			expression.Key("pk_gsi").Equal(expression.Value(pk)),
			expression.Key("sk").BeginsWith(sk),
		)
	} else {
		keyCond = expression.Key("pk_gsi").Equal(expression.Value(pk))
	}

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		return nil, err
	}

	key := &dynamodb.QueryInput{
					TableName:                 r.table,
					IndexName: 				   aws.String("pk_gsi"),
					ExpressionAttributeNames:  expr.Names(),
					ExpressionAttributeValues: expr.Values(),
					FilterExpression:          expr.Filter(),
					KeyConditionExpression:    expr.KeyCondition(),
					}

	fmt.Println(key)

	result, err := r.client.QueryWithContext(ctx, key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) UpdateInvoiceTransaction(ctx context.Context, pk string, sk string, order_id string, amount float32 ) (interface{}, error) {
	fmt.Printf("UpdateInvoiceTransaction \n")

	primaryKey := map[string]string{
		"pk": pk,
		"sk": sk,
	}
	pri_k, err := dynamodbattribute.MarshalMap(primaryKey)
	if err != nil {
		return nil, err
	}

	upd := expression.Set(
				expression.Name("update_at"),
				expression.Value(time.Now()),
			).Set(
				expression.Name("amount"),
				expression.Value(amount),
			)

	filter := expression.Equal(expression.Name("sk"), expression.Value(sk))
	filter2 := expression.AttributeNotExists(expression.Name("deletedAt"))

	expr, err := expression.NewBuilder().
					WithCondition(filter.And(filter2)).
					WithUpdate(upd).
					Build()
	if err != nil {
		return nil, err
	}

	key := &dynamodb.Update{
		TableName:                 r.table,
		Key:                       pri_k,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ConditionExpression: 	   expr.Condition(),
	}

	fmt.Println(key)

	primaryKey_02 := map[string]string{
		"pk": pk,
		"sk": order_id,
	}
	pri_k_02, err := dynamodbattribute.MarshalMap(primaryKey_02)
	if err != nil {
		return nil, err
	}
	upd_02 := expression.Set(
								expression.Name("update_at"),
								expression.Value(time.Now()),
							).Set(
								expression.Name("price"),
								expression.Value(amount),
							)

	filter = expression.Equal(expression.Name("sk"), expression.Value(order_id))
	filter2 = expression.AttributeNotExists(expression.Name("deletedAt"))

	expr_02, err := expression.NewBuilder().
					WithCondition(filter.And(filter2)).
					WithUpdate(upd_02).
					Build()
	if err != nil {
		return nil, err
	}

	key_02 := &dynamodb.Update{
		TableName:                 r.table,
		Key:                       pri_k_02,
		ExpressionAttributeNames:  expr_02.Names(),
		ExpressionAttributeValues: expr_02.Values(),
		UpdateExpression:          expr_02.Update(),
		ConditionExpression: 	   expr_02.Condition(),
	}

	fmt.Println("-----------------------------")
	fmt.Println(key_02)

	transactItems := make([]*dynamodb.TransactWriteItem, 2)
	transactItems[0] = &dynamodb.TransactWriteItem{Update: key}
	transactItems[1] = &dynamodb.TransactWriteItem{Update: key_02}

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		return nil, err
	}
	_, err = r.client.TransactWriteItemsWithContext(ctx, transaction)
	
	return nil, err

}