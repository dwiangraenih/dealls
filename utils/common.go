package utils

import "errors"

const (
	DirectionNext = "next"
	DirectionPrev = "prev"

	ConstCursorHashSalt = "Curs0rH4shS4lt0rD342024"
	ConstHashLength     = 10

	DefaultPage  = 1
	DefaultLimit = 10
)

var (
	ErrBadRequest       = errors.New("bad request")
	ErrInternal         = errors.New("error internal")
	ErrInvalidParameter = errors.New("invalid parameters, please check your input")
	ErrDuplicateData    = errors.New("duplicate data")
	ErrDataNotFound     = errors.New("data not found")
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

func IsIntInSlice(list []int, a int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
