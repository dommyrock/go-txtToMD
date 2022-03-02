package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/dommyrock/txtToMD/library/errorHandling"
	"github.com/dommyrock/txtToMD/library/files"
	"github.com/dommyrock/txtToMD/library/textUtil"
	"github.com/pkg/browser"

	types "github.com/dommyrock/txtToMD/types"
)

type Prefix = types.Prefix

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dict := map[string]Prefix{
		"#h1":    {Value: "#", Mode: "once"},
		"#h2":    {Value: "##", Mode: "once"},
		"#h3":    {Value: "###", Mode: "once"},
		"#h4":    {Value: "####", Mode: "once"},
		"#h5":    {Value: "#####", Mode: "once"},
		"-":      {Value: "---", Mode: "once"}, //line break
		"#code":  {Value: "```", Mode: "repeat"},
		"#b":     {Value: "**", Mode: "once"}, //bold
		"#img":   {Value: "#img", Mode: "once"},
		"#p":     {Value: ">", Mode: "multy"}, //paragraph
		"#link":  {Value: "link", Mode: "once"},
		"#bp":    {Value: "-", Mode: "multy"},     //bullet
		"#links": {Value: "links", Mode: "multy"}, //multiple links in a row
		"#table": {Value: "table", Mode: "multy"},
	}
	println(color.InCyan("Availiable mappings:"))
	fmt.Print(textUtil.MapKeys(dict))
	//TODO:
	//1 Only print above options if user entered options -h,
	//3 bugfix Md file extra ```` after code

	fileData, err := files.GetFileData()
	_, fileError := files.CheckIfFileIsValid(fileData.FilePath)

	//Validate inputs
	if err != nil {
		errorHandling.ExitGracefully(err)
	} else if fileError != nil {
		errorHandling.ExitGracefully(fileError)
	}

	writerChannel := make(chan string)
	filesCreated := make(chan bool)

	//Read file and and write to (MD,html)
	go files.ProcessFile(fileData, writerChannel, dict)
	go files.WriteToFiles(writerChannel, filesCreated)

	//Await channel to be closed, then unblock (Writing/Reading to channels is blocking)
	<-filesCreated

	//Print output file location
	pth, err := filepath.Abs(fileData.FilePath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	index := strings.LastIndex(pth, string(os.PathSeparator))
	fmt.Printf("Outputed markdown to dir: %s", pth[:index])

	//Open output html file
	browser.OpenFile(`D:\Downloads\output.html`)
}
