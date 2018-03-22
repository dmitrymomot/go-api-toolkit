package api

// Responder is a response interface
type Responder interface {
	SetData(data interface{})
	SetMeta(meta Meta)
	AddError(err Error)
	AddMessage(msg string)
	AddLink(title, link string)
}

// Links type
type Links map[string]string

// Data type, because Data{...} is more short than map[string]interface{}{...}
type Data map[string]interface{}

// Meta type
type Meta map[string]interface{}

// Response structure
type Response struct {
	HTTPStatus int         `json:"-"`
	Links      Links       `json:"links,omitempty"`
	Errors     []Error     `json:"errors,omitempty"`
	Meta       Meta        `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Messages   []string    `json:"messages,omitempty"`
}

// SetData sets response data for BaseResponse
func (r *Response) SetData(data interface{}) {
	r.Data = data
}

// SetMeta sets response meta
func (r *Response) SetMeta(meta Meta) {
	r.Meta = meta
}

// AddError adds error into errors array
func (r *Response) AddError(err Error) {
	r.Errors = append(r.Errors, err)
}

// AddMessage adds message into messages array
func (r *Response) AddMessage(msg string) {
	r.Messages = append(r.Messages, msg)
}

// AddLink adds link into links array
func (r *Response) AddLink(title, link string) {
	if r.Links == nil {
		r.Links = Links{}
	}
	r.Links[title] = link
}
