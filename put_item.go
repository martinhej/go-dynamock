package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *PutItemExpectation) ToTable(table string) *PutItemExpectation {
	e.table = &table
	return e
}

// WithItems - method for set Items expectation
func (e *PutItemExpectation) WithItems(item map[string]*dynamodb.AttributeValue) *PutItemExpectation {
	e.item = item
	return e
}

// WillReturns - method for set desired result
func (e *PutItemExpectation) WillReturns(res dynamodb.PutItemOutput, err error) *PutItemExpectation {
	e.output = &res
	return e
}

// WillError - mocks the call to return the provided error.
func (e *PutItemExpectation) WillError(err error) *PutItemExpectation {
	e.err = err
	return e
}

// PutItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if len(e.dynaMock.PutItemExpect) > 0 {
		x := e.dynaMock.PutItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.PutItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.item != nil {
			if !reflect.DeepEqual(x.item, input.Item) {
				return &dynamodb.PutItemOutput{}, fmt.Errorf("Expect item %+v but found item %+v", x.item, input.Item)
			}
		}

		// delete first element of expectation
		e.dynaMock.PutItemExpect = append(e.dynaMock.PutItemExpect[:0], e.dynaMock.PutItemExpect[1:]...)

		return x.output, x.err
	}

	return &dynamodb.PutItemOutput{}, fmt.Errorf("Put Item Expectation Not Found")
}

// PutItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opt ...request.Option) (*dynamodb.PutItemOutput, error) {
	if len(e.dynaMock.PutItemExpect) > 0 {
		x := e.dynaMock.PutItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.PutItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.item != nil {
			if !reflect.DeepEqual(x.item, input.Item) {
				return &dynamodb.PutItemOutput{}, fmt.Errorf("Expect item %+v but found item %+v", x.item, input.Item)
			}
		}

		// delete first element of expectation
		e.dynaMock.PutItemExpect = append(e.dynaMock.PutItemExpect[:0], e.dynaMock.PutItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.PutItemOutput{}, fmt.Errorf("Put Item With Context Expectation Not Found")
}
