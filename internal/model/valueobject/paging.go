package valueobject

import "net/url"

// Cursors ...
type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// Paging ...
type Paging struct {
	u           *url.URL `json:"-"`
	AfterCount  int      `json:"after_count,omitempty"`
	BeforeCount int      `json:"before_count,omitempty"`
	Count       int      `json:"count,omitempty"`
	Cursors     Cursors  `json:"cursors,omitempty"`
	Next        string   `json:"next"`
	Prev        string   `json:"prev"`
}

// CopyUrlValues ...
func CopyUrlValues(original url.Values) url.Values {
	copied := url.Values{}
	for k, v := range original {
		copied[k] = v
	}
	return copied
}

// BuildLinks ...
func (o *Paging) BuildLinks(params url.Values) {
	if o.Cursors.After != "" {
		nextParams := CopyUrlValues(params)
		nextParams.Set("after", o.Cursors.After)
		nextParams.Del("before")
		o.Next = nextParams.Encode()
		if o.u != nil {
			o.u.RawQuery = o.Next
			o.Next = o.u.String()
		}
	}
	if o.Cursors.Before != "" {
		prevParams := CopyUrlValues(params)
		prevParams.Del("after")
		prevParams.Set("before", o.Cursors.Before)
		o.Prev = prevParams.Encode()
		if o.u != nil {
			o.u.RawQuery = o.Prev
			o.Prev = o.u.String()
		}
	}
}

// BuildLinks ...
func (o *Paging) WithURL(u *url.URL) {
	o.u = u
}
