package output

const (
	Green = "\033[32m"
	Red   = "\033[31m"
	Grey  = "\033[90m"

	Reset = "\033[0m"
)

func DyeText(text string, color string) string {
	return color + text + Reset
}

func OutputError() string {
	return DyeText("[Error] ", Red)
}

func OutputInfo() string {
	return DyeText("[Info] ", Green)
}

func OutputDefault() string {
	return DyeText("[Default] ", Grey)
}
