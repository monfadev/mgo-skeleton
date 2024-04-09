package helpers

type ResponseParams struct {
	StatusCode int
	Message    string
	Paginate   *Paginate
	Data       any
}
