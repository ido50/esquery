package esquery

type TermsAggregation struct {
	name        string
	field       string
	size        *uint64
	shardSize   *float64
	showTermDoc *bool
	aggs        []Aggregation
}

func TermsAgg(name, field string) *TermsAggregation {
	return &TermsAggregation{
		name:  name,
		field: field,
	}
}

func (agg *TermsAggregation) Name() string {
	return agg.name
}

func (agg *TermsAggregation) Size(size uint64) *TermsAggregation {
	agg.size = &size
	return agg
}

func (agg *TermsAggregation) ShardSize(size float64) *TermsAggregation {
	agg.shardSize = &size
	return agg
}

func (agg *TermsAggregation) ShowTermDocCountError(b bool) *TermsAggregation {
	agg.showTermDoc = &b
	return agg
}

func (agg *TermsAggregation) Aggs(aggs ...Aggregation) *TermsAggregation {
	agg.aggs = aggs
	return agg
}

func (agg *TermsAggregation) Map() map[string]interface{} {
	innerMap := map[string]interface{}{
		"field": agg.field,
	}

	if agg.size != nil {
		innerMap["size"] = *agg.size
	}
	if agg.shardSize != nil {
		innerMap["shard_size"] = *agg.shardSize
	}
	if agg.showTermDoc != nil {
		innerMap["show_term_doc_count_error"] = *agg.showTermDoc
	}

	outerMap := map[string]interface{}{
		"terms": innerMap,
	}
	if len(agg.aggs) > 0 {
		subAggs := make(map[string]map[string]interface{})
		for _, sub := range agg.aggs {
			subAggs[sub.Name()] = sub.Map()
		}
		outerMap["aggs"] = subAggs
	}

	return outerMap
}
