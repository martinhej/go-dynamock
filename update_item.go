package dynamock

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ToTable - method for set Table expectation
func (e *UpdateItemExpectation) ToTable(table string) *UpdateItemExpectation {
	e.table = &table
	return e
}

// WithKeys - method for set Keys expectation
func (e *UpdateItemExpectation) WithKeys(keys map[string]*dynamodb.AttributeValue) *UpdateItemExpectation {
	e.key = keys
	return e
}

// Updates - method for set Updates expectation
func (e *UpdateItemExpectation) Updates(attrs map[string]*dynamodb.AttributeValueUpdate) *UpdateItemExpectation {
	e.attributeUpdates = attrs
	return e
}

// WillReturns - method for set desired result
func (e *UpdateItemExpectation) WillReturns(res dynamodb.UpdateItemOutput) *UpdateItemExpectation {
	e.output = &res
	return e
}

// WillError - mocks the call to return the provided error.
func (e *UpdateItemExpectation) WillError(err error) *UpdateItemExpectation {
	e.err = err
	return e
}

// UpdateItem - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, x.err
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item Expectation Not Found")
}

// UpdateItemWithContext - this func will be invoked when test running matching expectation with actual input
func (e *MockDynamoDB) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opt ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	if len(e.dynaMock.UpdateItemExpect) > 0 {
		x := e.dynaMock.UpdateItemExpect[0] //get first element of expectation

		if x.table != nil {
			if *x.table != *input.TableName {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect table %s but found table %s", *x.table, *input.TableName)
			}
		}

		if x.key != nil {
			if !reflect.DeepEqual(x.key, input.Key) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.key, input.Key)
			}
		}

		if x.attributeUpdates != nil {
			if !reflect.DeepEqual(x.attributeUpdates, input.AttributeUpdates) {
				return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Expect key %+v but found key %+v", x.attributeUpdates, input.AttributeUpdates)
			}
		}

		// delete first element of expectation
		e.dynaMock.UpdateItemExpect = append(e.dynaMock.UpdateItemExpect[:0], e.dynaMock.UpdateItemExpect[1:]...)

		return x.output, nil
	}

	return &dynamodb.UpdateItemOutput{}, fmt.Errorf("Update Item With Context Expectation Not Found")
}
