package examples

import (
	"errors"
	dynamock "github.com/gusaul/go-dynamock"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

	mock.ExpectGetItem().ToTable("employee").WithKeys(expectKey).WillReturns(dynamodb.GetItemOutput{}, errors.New("error"))

	_, err := GetName("1")
	if err.Error() != "error" {
		t.Error("Error was expected")
	}
}

func TestGetTransactGetItems(t *testing.T) {
	databaseOutput := dynamodb.TransactWriteItemsOutput{}

	mock.ExpectTransactWriteItems().Table("wrongTable").WillReturns(databaseOutput, nil)

	err := GetTransactGetItems("")

	if err == nil {
		t.Errorf("Test failed")
	}
}
