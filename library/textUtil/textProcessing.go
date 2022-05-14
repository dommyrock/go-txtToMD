package textUtil

import (
	"fmt"
	"strings"

	types "github.com/dommyrock/txtToMD/types"
)

type MapKeys map[string]types.Prefix

//Prints all available txt file mappings
func (s MapKeys) String() string {
	var str string
	for k := range s {
		str += fmt.Sprintf("%v >> %v\n", k, decodePrefix(k))
	}
	return str
}

//Translates prefix to description
func decodePrefix(r string) string {
	switch r {
	case "#h1":
		return "Heading 1"
	case "#h2":
		return "Heading 2"
	case "#h3":
		return "Heading 3"
	case "#h4":
		return "Heading 4"
	case "#h5":
		return "Heading 5"
	case "-":
		return "Line break"
	case "#code":
		return "Code block"
	case "#b":
		return "Bold"
	case "#bp":
		return "Bullet point"
	case "#p":
		return "Paragraph"
	case "#link":
		return "Single link"
	case "#links":
		return "Multiple links"
	case "#table":
		return "Table"
	case "#img":
		return "Image"
	default:
		return r
	}
}

func InsertHeaderLine(text string) string {
	pipeRepeated := strings.Count(text, "|")
	result := ""
	i := 0
	for i < pipeRepeated-1 {
		result += "| ------ "
		i++
	}
	return result + "|"
}

func HandlePrefix(prefix string, txt string) string {
	switch prefix {
	case "```":
		return "```\n" + txt
	case "**":
		return "**" + txt + "**"
	case "table":
		{
			slices := strings.Split(txt, ",")
			builtString := "| "
			i := 0
			for i < len(slices) {
				builtString += slices[i] + " | "
				i++
			}
			return builtString
		}
	case "-":
		return "- " + txt
	case "#img":
		return "![image](" + txt + ")"
	case "link", "links":
		title, link, description := SplitLink(txt, ",")
		return "[" + title + "]" + "(" + link + ")" + " -" + description
	default:
		return prefix + " " + txt
	}
}

func SplitLink(s, sep string) (string, string, string) {
	x := strings.Split(s, sep)
	switch len(x) {
	case 3:
		return x[0], x[1], x[2] //title, link, description
	case 2:
		return x[0], x[1], ""
	case 1:
		return x[0], "", ""
	default:
		return "", "", ""
	}
}

func TrimCodeStart(text string) string {
	if strings.HasPrefix(text, "`") {
		return TrimCodeStart(strings.TrimPrefix(text, "`"))
	}
	return text
}
