package wsFramework

const (
	Green = "\033[32m"
	Red   = "\033[31m"
	Grey  = "\033[90m"

	Reset = "\033[0m"
)

func dyeText(text string, color string) string {
	return color + text + Reset
}

func outputError() string {
	return dyeText("[Error] ", Red)
}

func outputInfo() string {
	return dyeText("[Info] ", Green)
}

func outputDefault() string {
	return dyeText("[Default] ", Grey)
}
