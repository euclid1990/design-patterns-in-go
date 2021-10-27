package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {
	instance = nil
}

func (suite *TestSuite) TestDbGetConnection() {
	connectionString := "myuser@example.com:3306/main-schema"
	db := &db{Connection: connectionString}
	assert.Equal(suite.T(), connectionString, db.GetConnection())
}

func (suite *TestSuite) TestSetInstance() {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()
	setInstance("myuser@example.com:3306/main-schema")
	assert.Equal(suite.T(), "Singleton instance has been created.\n", out.(*bytes.Buffer).String())
}

func (suite *TestSuite) TestGetInstanceByDoOnce() {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()
	connectionStringFirst := "myuser@example.com:3306/main-schema-do-one-first"
	connectionStringSecond := "myuser@example.com:3306/main-schema-do-one-second"
	instanceFirst := GetInstanceByDoOnce(connectionStringFirst)
	instanceSecond := GetInstanceByDoOnce(connectionStringSecond)
	assert.Equal(suite.T(), connectionStringFirst, instanceFirst.GetConnection())
	assert.Equal(suite.T(), connectionStringFirst, instanceSecond.GetConnection())
}

func (suite *TestSuite) TestGetInstanceByDoLock() {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()
	connectionStringFirst := "myuser@example.com:3306/main-schema-do-lock-first"
	connectionStringSecond := "myuser@example.com:3306/main-schema-do-lock-second"
	instanceFirst := GetInstanceByDoLock(connectionStringFirst)
	instanceSecond := GetInstanceByDoLock(connectionStringSecond)
	assert.Equal(suite.T(), connectionStringFirst, instanceFirst.GetConnection())
	assert.Equal(suite.T(), connectionStringFirst, instanceSecond.GetConnection())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
