package textUtil

import "strings"

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
