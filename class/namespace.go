package class

type NameSpace struct {
	Name     string                `json:"name"`
	Handlers map[string]SendMethod `json:"-"` // key: path, value: sendMethod
}
