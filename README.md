[![GoDoc](https://godoc.org/github.com/martinhej/go-dynamock?status.png)](https://godoc.org/github.com/martinhej/go-dynamock) [![Go Report Card](https://goreportcard.com/badge/github.com/martinhej/go-dynamock)](https://goreportcard.com/report/github.com/martinhej/go-dynamock) [![Build Status](https://travis-ci.com/martinhej/go-dynamock.svg?branch=master)](https://travis-ci.com/martinhej/go-dynamock)
# go-dynamock
Amazon Dynamo DB Mock Driver for Golang to Test Database Interactions

## Install
```
go get github.com/martinhej/go-dynamock
```

## Examples Usage
Visit [godoc](https://godoc.org/github.com/martinhej/go-dynamock) for general examples and public api reference.

### DynamoDB configuration
First of all, change the dynamodb configuration to use the ***dynamodb interface***. see code below:
``` go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MyDynamo struct {
    Db dynamodbiface.DynamoDBAPI
}

var Dyna *MyDynamo

func ConfigureDynamoDB() {
	Dyna = new(MyDynamo)
	awsSession, _ := session.NewSession(&aws.Config{Region: aws.String("ap-southeast-2")})
	svc := dynamodb.New(awsSession)
	Dyna.Db = dynamodbiface.DynamoDBAPI(svc)
}
```
the purpose of code above is to make your dynamoDB object can be mocked by ***dynamock*** through the dynamodbiface.

### Something you may wanna test
``` go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetName(id string) (*string, error) {
	parameter := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String("employee"),
	}

	response, err := Dyna.Db.GetItem(parameter)
	if err != nil {
		return nil, err
	}

	name := response.Item["name"].S
	return name, nil
}
```

### Test with DynaMock
``` go
package examples

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	dynamock "github.com/martinhej/go-dynamock"
)

var mock *dynamock.DynaMock

func init() {
	Dyna = new(MyDynamo)
	Dyna.Db, mock = dynamock.New()
}

func TestGetName(t *testing.T) {
	expectKey := map[string]*dynamodb.AttributeValue{
		"id": {
			N: aws.String("1"),
		},
	}

	expectedResult := aws.String("jaka")
	result := dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"name": {
				S: expectedResult,
			},
		},
	}

	//lets start dynamock in action
	mock.ExpectGetItem().ToTable("employee").WithKeys(expectKey).WillReturns(result, nil)

	actualResult, _ := GetName("1")
	if actualResult != expectedResult {
		t.Errorf("Test Fail")
	}
}
```
if you just wanna expect the table
``` go
mock.ExpectGetItem().ToTable("employee").WillReturns(result, nil)
```
or maybe you didn't care with any arguments, you just need to determine the result
``` go
mock.ExpectGetItem().WillReturns(result, nil)
```
and you can do multiple expectations at once, then the expectation will be executed sequentially.
``` go
mock.ExpectGetItem().WillReturns(resultOne, nil)
mock.ExpectUpdateItem().WillReturns(resultTwo, nil)
mock.ExpectGetItem().WillReturns(resultThree, nil)

/* Result
the first call of GetItem will return resultOne
the second call of GetItem will return resultThree
and the only call of UpdateItem will return resultTwo */
```
### Currently Supported Functions
``` go
CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
GetItemWithContext(aws.Context, *dynamodb.GetItemInput, ...request.Option) (*dynamodb.GetItemOutput, error)
PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
PutItemWithContext(aws.Context, *dynamodb.PutItemInput, ...request.Option) (*dynamodb.PutItemOutput, error)
UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
UpdateItemWithContext(aws.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error)
DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
DeleteItemWithContext(aws.Context, *dynamodb.DeleteItemInput, ...request.Option) (*dynamodb.DeleteItemOutput, error)
BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
BatchGetItemWithContext(aws.Context, *dynamodb.BatchGetItemInput, ...request.Option) (*dynamodb.BatchGetItemOutput, error)
BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
BatchWriteItemWithContext(aws.Context, *dynamodb.BatchWriteItemInput, ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
WaitUntilTableExists(*dynamodb.DescribeTableInput) error
Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
ScanPages(input *ScanInput, fn func(*ScanOutput, bool) bool) error
ScanPagesWithContext(ctx aws.Context, input *ScanInput, fn func(*ScanOutput, bool) bool, opts ...request.Option) error
ScanWithContext(aws.Context, *dynamodb.ScanInput, ...request.Option) (*dynamodb.ScanOutput, error)
Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
QueryWithContext(aws.Context, *dynamodb.QueryInput, request.Option) (*dynamodb.QueryOutput, error)
QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error
QueryPagesWithContext(aws.Context, *dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool, ...request.Option) error
```
## Contributions

Feel free to open a pull request. Note, if you wish to contribute an extension to public (exported methods or types) -
please open an issue before, to discuss whether these changes can be accepted. All backward incompatible changes are
and will be treated cautiously

## License

The [MIT License](https://github.com/martinhej/go-dynamock/blob/master/LICENSE)
