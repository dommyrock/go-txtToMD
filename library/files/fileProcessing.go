package files

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/dommyrock/txtToMD/library/errorHandling"
	"github.com/dommyrock/txtToMD/library/textUtil"
	"github.com/dommyrock/txtToMD/static"
	types "github.com/dommyrock/txtToMD/types"
	"github.com/mitchellh/go-homedir"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type InputFile = types.InputFile

func GetFileData() (InputFile, error) {
	if len(os.Args) < 2 {
		return InputFile{}, errors.New("filepath must be specified")
	}
	//Parse all cli args
	flag.Parse()
	return InputFile{flag.Arg(0)}, nil
}

//Reads from input File and writes parsed lines to writerChannel
func ProcessFile(fileData InputFile, writerChannel chan<- string, dict map[string]types.Prefix) {
	file, err := os.Open(fileData.FilePath)

	if err != nil {
		errorHandling.ExitGracefully(err)
	}
	defer file.Close()

	var currentPrefix string = ""
	var mode string = ""
	var blockStarted bool = false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		line := strings.TrimSpace(scanner.Text())
		val, found := dict[line]
		//Check how many times to repeat cmd, set prefix
		if found {
			mode = val.Mode
			currentPrefix = val.Value
			//edge case
			if currentPrefix == "---" {
				writerChannel <- currentPrefix
			}
			continue
		}

		if line != "" {
			if mode == "multy" {
				text := textUtil.HandlePrefix(currentPrefix, line)
				writerChannel <- text
				//also write header split
				if !blockStarted && currentPrefix == "table" {
					blockStarted = true
					writerChannel <- textUtil.InsertHeaderLine(text)
				}
			} else if mode == "once" && currentPrefix != "" {
				writerChannel <- textUtil.HandlePrefix(currentPrefix, line)
				currentPrefix = ""
			} else if mode == "repeat" && !blockStarted {
				writerChannel <- textUtil.HandlePrefix(currentPrefix, line)
				blockStarted = true
			} else {
				writerChannel <- line
			}
		} else { // newline, empty line
			if mode == "repeat" {
				writerChannel <- line + currentPrefix + "\n"
			} else {
				writerChannel <- line
			}
			//RESET after writing to channel
			if !found {
				mode = ""
				currentPrefix = ""
				blockStarted = false

			}
		}
	}
	close(writerChannel)
}

func CheckIfFileIsValid(filename string) (bool, error) {
	// Check if file is txt
	if fileExtension := filepath.Ext(filename); fileExtension != ".txt" {
		return false, fmt.Errorf("file %s is not txt, it has %s extension", filename, fileExtension)
	}

	// Check if file exists
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", filename)
	}
	return true, nil
}

func CreateFileWriter(path string) func(string, bool) {
	file, err := os.Create(path)
	if err != nil {
		errorHandling.ExitGracefully(err)
	}

	return func(data string, close bool) {
		_, err := file.WriteString(data + "\n")
		if err != nil {
			errorHandling.ExitGracefully(err)
		}

		if close {
			file.Close()
		}
	}
}

//1. Awaits writes from writerChannel
//2.Writes to file
//3.Signals done <-
func WriteToFiles(writerChannel <-chan string, done chan<- bool) {
	fmt.Printf("Reading from writerChannel and writing to files...\n")

	//Get user home dir regardless of the OS
	homeDir, err := homedir.Dir()
	if err != nil {
		errorHandling.ExitGracefully(err)
	}
	htmlLocation := fmt.Sprintf("%s/%s", homeDir+"\\Downloads", "generated.html")
	writeToHTML := CreateFileWriter(htmlLocation)
	mdLocation := fmt.Sprintf("%s/%s", homeDir+"\\Downloads", "generated.md")
	writeToMD := CreateFileWriter(mdLocation)

	//Tables & code blocks need to be converted from MD to html all at once
	builder := strings.Builder{}

	//HTML Writer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)

	beginningOfFile, wroteEnd, detectedTable, detectedCode := true, false, false, false
	for {
		text, more := <-writerChannel
		var buf bytes.Buffer

		if more {
			writeToMD(text, false)

			if beginningOfFile {
				//construct html template/styles
				writeToHTML(htmlStleBuilder(), false)
				beginningOfFile = false
			}
			if text == "```\n" && builder.Len() == 0 {
				continue
			}

			if strings.HasPrefix(text, "|") || detectedTable { //table
				if text == "" { //end of table
					mdToHTML(md, builder.String(), buf, writeToHTML)
					detectedTable = false
					builder.Reset()
				} else {
					detectedTable = true
					builder.WriteString(text + "\n")
				}
			} else if strings.HasPrefix(text, "```") || detectedCode { //code block
				//end of code block
				if text == "```\n" && builder.Len() > 0 {
					err := quick.Highlight(&builder, builder.String(), "go", "html", "dracula") //options: monokai,dracula,rainbow_dash
					if err != nil {
						log.Fatalf("error highlighting code bock. %s", err)
					}

					innerBody := strings.Split(strings.Split(builder.String(), `<body class="bg">`)[1], `</body>`)[0]
					//Raw HTML is passed (no need to re encode it)
					writeToHTML(innerBody, false)
					detectedCode = false
					builder.Reset()

				} else {
					detectedCode = true
					builder.WriteString(textUtil.TrimCodeStart(text) + "\n")
				}

			} else {
				mdToHTML(md, text, buf, writeToHTML)
			}

		} else { //Writer channel is closed
			//Write remaining text from StringBuiilder to file
			if builder.Len() > 0 && detectedTable {
				mdToHTML(md, builder.String(), buf, writeToHTML)
				detectedTable = false
			}
			if !wroteEnd {
				writeToHTML(static.HtmlEnd, false)
				wroteEnd = true
			}
			done <- true
		}
	}
}

//Converts md syntax to html and writes to file
func mdToHTML(md goldmark.Markdown, text string, buf bytes.Buffer, writeToHTML func(string, bool)) {
	if errr := md.Convert([]byte(text), &buf); errr != nil {
		log.Fatalf("error converting txt to")
	}
	writeToHTML(buf.String(), false)
}

//Builds <style> based on theme arg
func htmlStleBuilder() string {
	args := os.Args
	switch len(args) {
	case 3:
		if args[2] == "mid" {
			return static.HtmlOpen + static.Root_midTheme + static.HtmlMid
		}
		return static.HtmlOpen + static.Root_darkTheme + static.HtmlMid
	default:
		return static.HtmlOpen + static.Root_defaultTheme + static.HtmlMid
	}
}
