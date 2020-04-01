package esquery

type Source struct {
	includes []string
}

func (source Source) Map() map[string]interface{} {
	return map[string]interface{}{
		"includes": source.includes,
	}
}

type Sort []map[string]interface{}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)
