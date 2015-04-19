package main

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"log"
	"net"
	"testing"
	"time"
)

type ExampleTestSuite struct {
	suite.Suite
	conn     net.Conn
	listener net.Listener
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	suite.listener = Listener(":7435")
	conn, err := net.Dial("tcp", ":7435")
	if err != nil {
		log.Fatal(err)
	}
	suite.conn = conn
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	go Handler(suite.listener)
	// suite.NoError(err)
	suite.conn.Write([]byte(`<policy-file-request/>\0`))
	var a []byte = make([]byte, len(policy)/2)
	var b []byte = make([]byte, len(policy)/2)
	suite.conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	suite.conn.Read(a)
	suite.conn.Read(b)
	var buff bytes.Buffer
	buff.Write(a)
	buff.Write(b)
	suite.Equal(policy, buff.Bytes())
	suite.listener.Close()
	// suite.Equal(suite.VariableThatShouldStartAtFive, 5)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
