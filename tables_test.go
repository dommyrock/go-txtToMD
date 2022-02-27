package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type testData struct {
	md   []byte
	html []byte
}

func testDataToStrArray(tests []*testData) []string {
	res := []string{}
	for _, td := range tests {
		res = append(res, string(td.md))
		res = append(res, string(td.html))
	}
	return res
}

func readTestFile(t *testing.T, fileName string) []*testData {
	path := filepath.Join("testdata", fileName)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("ioutil.ReadFile('%s') failed with %s", path, err)
	}
	parts := bytes.Split(d, []byte("+++\n"))
	if len(parts)%2 != 0 {
		t.Fatalf("odd test tuples in file %s: %d", path, len(parts))
	}
	res := []*testData{}
	n := len(parts) / 2
	for i := 0; i < n; i++ {
		j := i * 2
		td := &testData{
			md:   parts[j],
			html: parts[j+1],
		}
		res = append(res, td)
	}
	return res
}

func TestTable(t *testing.T) {
	doTestsBlock(t, "Table.tests", parser.Tables)
}

type TestParams struct {
	extensions        parser.Extensions
	referenceOverride parser.ReferenceOverrideFunc
	html.Flags
	html.RendererOptions
}

func doTestsBlock(t *testing.T, path string, extensions parser.Extensions) {
	tests := readTestFile2(t, path)
	doTestsParam(t, tests, TestParams{
		extensions: extensions,
		Flags:      html.UseXHTML,
	})
}

func readTestFile2(t *testing.T, fileName string) []string {
	tests := readTestFile(t, fileName)
	return testDataToStrArray(tests)
}

func doTestsParam(t *testing.T, tests []string, params TestParams) {
	for i := 0; i+1 < len(tests); i += 2 {
		input := tests[i]
		expected := tests[i+1]
		got := runMarkdown(input, params)
		if got != expected {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nGot     [%#v]\nInput:\n%s\nExpected:\n%s\nGot:\n%s\n",
				input, expected, got, input, expected, got)
		}
	}
}

func runMarkdown(input string, params TestParams) string {
	params.RendererOptions.Flags = params.Flags
	parser := parser.NewWithExtensions(params.extensions)
	parser.ReferenceOverride = params.referenceOverride
	renderer := html.NewRenderer(params.RendererOptions)

	d := markdown.ToHTML([]byte(input), parser, renderer)
	return string(d)
}
