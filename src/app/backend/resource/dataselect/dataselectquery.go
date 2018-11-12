package dataselect

type DataSelectQuery struct {
	PaginationQuery *PaginationQuery
}

//  NewDataSelectQuery creates DataSelectQuery object from data select queries
func NewDataSelectQuery(paginationQuery *PaginationQuery) *DataSelectQuery {
	return &DataSelectQuery{
		PaginationQuery: paginationQuery,
	}
}
