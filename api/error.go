package api

// Error structure
type Error struct {
	ID        string      `json:"id,omitempty"`
	Status    int         `json:"-"`
	Code      int         `json:"code,omitempty"`
	Title     string      `json:"title,omitempty"`
	Detail    interface{} `json:"detail,omitempty"`
	Links     Links       `json:"links,omitempty"`
	PrevError error       `json:"-"`
}

// Error text
func (e *Error) Error() string {
	return e.Title
}
