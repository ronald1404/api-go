package request

type UpdateTagsRequest struct {
	Id   int    `validadte:"required"`
	Name string `validade:"required, max=200, min=1" json="name"`
}
