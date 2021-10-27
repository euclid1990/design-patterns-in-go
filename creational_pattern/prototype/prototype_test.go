package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) SetupTest() {
}

func (suite *TestSuite) TestFilePrint() {
	type parameters struct {
		f *file
	}
	testCases := []struct {
		name   string
		args   parameters
		expect string
	}{
		{
			name:   "Abnormal",
			args:   parameters{f: &file{name: ""}},
			expect: "-\n",
		},
		{
			name:   "Normal 1",
			args:   parameters{f: &file{name: "file1"}},
			expect: "-file1\n",
		},
		{
			name:   "Normal 2",
			args:   parameters{f: &file{name: "file2"}},
			expect: "-file2\n",
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			out := captureOutput(func() {
				tc.args.f.print("-")
			})
			assert.Equal(suite.T(), tc.expect, out)
		})
	}
}

func (suite *TestSuite) TestFileClone() {
	type parameters struct {
		f *file
	}
	testCases := []struct {
		name   string
		args   parameters
		expect *file
	}{
		{
			name:   "Abnormal",
			args:   parameters{f: &file{name: ""}},
			expect: &file{name: "_clone"},
		},
		{
			name:   "Normal 1",
			args:   parameters{f: &file{name: "file1"}},
			expect: &file{name: "file1_clone"},
		},
		{
			name:   "Normal 2",
			args:   parameters{f: &file{name: "file2"}},
			expect: &file{name: "file2_clone"},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			assert.Equal(suite.T(), tc.expect, tc.args.f.clone())
		})
	}
}

func (suite *TestSuite) TestFolderPrint() {
	type parameters struct {
		f *folder
	}
	testCases := []struct {
		name   string
		args   parameters
		expect string
	}{
		{
			name:   "Abnormal",
			args:   parameters{f: &folder{name: ""}},
			expect: "-\n",
		},
		{
			name:   "Normal 1",
			args:   parameters{f: &folder{name: "folder1"}},
			expect: "-folder1\n",
		},
		{
			name:   "Normal 2",
			args:   parameters{f: &folder{name: "folder2", children: []iNode{&file{name: "file1"}}}},
			expect: "-folder2\n--file1\n",
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			out := captureOutput(func() {
				tc.args.f.print("-")
			})
			assert.Equal(suite.T(), tc.expect, out)
		})
	}
}

func (suite *TestSuite) TestFolderClone() {
	type parameters struct {
		f *folder
	}
	testCases := []struct {
		name   string
		args   parameters
		expect *folder
	}{
		{
			name:   "Abnormal",
			args:   parameters{f: &folder{name: ""}},
			expect: &folder{name: "_clone"},
		},
		{
			name:   "Normal 1",
			args:   parameters{f: &folder{name: "folder1"}},
			expect: &folder{name: "folder1_clone"},
		},
		{
			name:   "Normal 2",
			args:   parameters{f: &folder{name: "folder2", children: []iNode{&file{name: "file1"}}}},
			expect: &folder{name: "folder2_clone", children: []iNode{&file{name: "file1_clone"}}},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			assert.Equal(suite.T(), tc.expect, tc.args.f.clone())
		})
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
