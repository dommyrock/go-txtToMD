package types

type Prefix struct {
	Value string
	Mode  string
}
type InputFile struct {
	FilePath string
}
type MapKeys map[string]Prefix
