// Package api :: This is auto generated file, do not edit manually
package api

// Conf receives custom request config definition, e.g. custom headers, custom OData mod
func (files *Files) Conf(config *RequestConfig) *Files {
	files.config = config
	return files
}

// Select adds $select OData modifier
func (files *Files) Select(oDataSelect string) *Files {
	files.modifiers.AddSelect(oDataSelect)
	return files
}

// Expand adds $expand OData modifier
func (files *Files) Expand(oDataExpand string) *Files {
	files.modifiers.AddExpand(oDataExpand)
	return files
}

// Filter adds $filter OData modifier
func (files *Files) Filter(oDataFilter string) *Files {
	files.modifiers.AddFilter(oDataFilter)
	return files
}

// Top adds $top OData modifier
func (files *Files) Top(oDataTop int) *Files {
	files.modifiers.AddTop(oDataTop)
	return files
}

// OrderBy adds $orderby OData modifier
func (files *Files) OrderBy(oDataOrderBy string, ascending bool) *Files {
	files.modifiers.AddOrderBy(oDataOrderBy, ascending)
	return files
}