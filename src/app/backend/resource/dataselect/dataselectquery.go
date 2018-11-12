package dataselect

type DataSelectQuery struct {
	PaginationQuery 	*PaginationQuery
	SortQuery					*SortQuery
	FilterQuery				*FilterQuery
}

//  NewDataSelectQuery creates DataSelectQuery object from data select queries
func NewDataSelectQuery(paginationQuery *PaginationQuery, sortQuery *SortQuery, filterQuery *FilterQuery) *DataSelectQuery {
	return &DataSelectQuery{
		PaginationQuery: paginationQuery,
		SortQuery: sortQuery,
		FilterQuery: filterQuery,
	}
}
