package main

import (
	"log"

	"github.com/dommyrock/txtToMD/library/errorHandling"
	"github.com/dommyrock/txtToMD/library/files"
	"github.com/dommyrock/txtToMD/static"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/browser"

	types "github.com/dommyrock/txtToMD/types"
)

type Prefix = types.Prefix

func main() {
	mappings := map[string]Prefix{
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

	if errorHandling.ShouldExit(mappings) {
		return
	}
	//Log config
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fileData, err := files.GetFileData()
	_, fileError := files.CheckIfFileIsValid(fileData.FilePath)

	if err != nil {
		errorHandling.ExitGracefully(err)
	} else if fileError != nil {
		errorHandling.ExitGracefully(fileError)
	}

	writerChannel := make(chan string)
	filesCreated := make(chan bool)

	//Read /Write to [MD,html]
	go files.ProcessFile(fileData, writerChannel, mappings)
	go files.WriteToFiles(writerChannel, filesCreated)

	//Await channel to be closed, then unblock
	<-filesCreated

	homeDir, err := homedir.Dir()
	static.PrintOutputDirLocation(homeDir, err)
	//Open browser
	browser.OpenFile(homeDir + "\\Downloads\\generated.html")

	//Hosting static file in cloud
	static.StaticFileHostingNote()
}
