package entity

type Result struct {
	IP    string `json:"ip" firestore:"ip"`
	Host  string `json:"host" firestore:"host"`
	Kind  string `json:"kind" firestore:"kind"`
	Error string `json:"error" firestore:"error"`
	From  string `json:"from" firestore:"from"`
}

func (r *Result) IsEmpty() bool {
	result := Result{}
	if result == *r {
		return true
	}
	return false
}
