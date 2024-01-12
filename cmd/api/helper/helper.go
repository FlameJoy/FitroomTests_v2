package helper

import "encoding/json"

func URLFormat(s string) string {
	if len(s) < 90 {
		for len(s) < 90 {
			s += " "
		}
	} else if len(s) > 90 {
		s = s[0:90-3] + "..."
	}
	return s
}

func PrepareReqBody(v any) []byte {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
