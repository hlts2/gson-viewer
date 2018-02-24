package goson

// Result is search and path result struct
type Result struct {
	jsonObject interface{}
}

//Indent converts json object to json string
func (r *Result) Indent(prefix, indent string) (string, error) {
	return indentJSONString(r.jsonObject, prefix, indent)
}
