package resources

func (src *Source) FetchStock(category, key string) (interface{}, error) {
	return src.get("stock", category, key)
}

func (src *Source) FetchAllStock(category, pagesize string) (map[string]interface{}, error) {
	res, err := src.get("stock", category, pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
