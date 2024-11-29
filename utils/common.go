package utils

const (
	DirectionNext = "next"
	DirectionPrev = "prev"

	ConstCursorHashSalt = "Curs0rH4shS4lt0rD342024"
	ConstHashLength     = 10

	DefaultPage  = 1
	DefaultLimit = 10
)

func GetPaginationCursor(dataCursor []int, isPrevCursor bool) (prevCursor, nextCursor int64) {
	prevCursor = int64(dataCursor[0])
	nextCursor = int64(dataCursor[len(dataCursor)-1])

	var dataCursorPrev []int
	if isPrevCursor {

		for dataCursorIdx := len(dataCursor) - 1; dataCursorIdx >= 0; dataCursorIdx-- {
			dataCursorPrev = append(dataCursorPrev, dataCursor[dataCursorIdx])
		}

		prevCursor = int64(dataCursorPrev[0])
		nextCursor = int64(dataCursorPrev[len(dataCursorPrev)-1])
	}

	return prevCursor, nextCursor
}
