package dynamock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type (
	// MockDynamoDB struct hold DynamoDBAPI implementation and mock object
	MockDynamoDB struct {
		dynamodbiface.DynamoDBAPI
		dynaMock *DynaMock
	}

	// DynaMock mock struct hold all expectation types
	DynaMock struct {
		GetItemExpect            []GetItemExpectation
		BatchGetItemExpect       []BatchGetItemExpectation
		UpdateItemExpect         []UpdateItemExpectation
		PutItemExpect            []PutItemExpectation
		DeleteItemExpect         []DeleteItemExpectation
		BatchWriteItemExpect     []BatchWriteItemExpectation
		CreateTableExpect        []CreateTableExpectation
		DescribeTableExpect      []DescribeTableExpectation
		WaitTableExistExpect     []WaitTableExistExpectation
		ScanExpect               []ScanExpectation
		QueryExpect              []QueryExpectation
		TransactWriteItemsExpect []TransactWriteItemsExpectation
	}

	// GetItemExpectation struct hold expectation field, err, and result
	GetItemExpectation struct {
		table  *string
		key    map[string]*dynamodb.AttributeValue
		output *dynamodb.GetItemOutput
		err    error
	}

	// BatchGetItemExpectation struct hold expectation field, err, and result
	BatchGetItemExpectation struct {
		input  map[string]*dynamodb.KeysAndAttributes
		output *dynamodb.BatchGetItemOutput
		err    error
	}

	// UpdateItemExpectation struct hold expectation field, err, and result
	UpdateItemExpectation struct {
		attributeUpdates map[string]*dynamodb.AttributeValueUpdate
		key              map[string]*dynamodb.AttributeValue
		table            *string
		output           *dynamodb.UpdateItemOutput
		err              error
	}

	// PutItemExpectation struct hold expectation field, err, and result
	PutItemExpectation struct {
		item   map[string]*dynamodb.AttributeValue
		table  *string
		output *dynamodb.PutItemOutput
		err    error
	}

	// DeleteItemExpectation struct hold expectation field, err, and result
	DeleteItemExpectation struct {
		key    map[string]*dynamodb.AttributeValue
		table  *string
		output *dynamodb.DeleteItemOutput
		err    error
	}

	// BatchWriteItemExpectation struct hold expectation field, err, and result
	BatchWriteItemExpectation struct {
		input  map[string][]*dynamodb.WriteRequest
		output *dynamodb.BatchWriteItemOutput
		err    error
	}

	// CreateTableExpectation struct hold expectation field, err, and result
	CreateTableExpectation struct {
		keySchema []*dynamodb.KeySchemaElement
		table     *string
		output    *dynamodb.CreateTableOutput
		err       error
	}

	// DescribeTableExpectation struct hold expectation field, err, and result
	DescribeTableExpectation struct {
		table  *string
		output *dynamodb.DescribeTableOutput
		err    error
	}

	// WaitTableExistExpectation struct hold expectation field, err, and result
	WaitTableExistExpectation struct {
		table *string
		err   error
	}

	// ScanExpectation struct hold expectation field, err, and result
	ScanExpectation struct {
		table  *string
		output *dynamodb.ScanOutput
		err    error
	}

	// QueryExpectation struct hold expectation field, err, and result
	QueryExpectation struct {
		table  *string
		output *dynamodb.QueryOutput
		err    error
	}

	// TransactWriteItemsExpectation struct holds field, err, and result
	TransactWriteItemsExpectation struct {
		table  *string
		items  []*dynamodb.TransactWriteItem
		output *dynamodb.TransactWriteItemsOutput
		err    error
	}
)
