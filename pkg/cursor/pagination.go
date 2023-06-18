package cursor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const (
	ERRORED_CURSOR             = "ERRORED"
	DEFAULT_LIMIT        int64 = 20
	DEFAULT_START_OFFSET int64 = 0
)

type PaginationInfo struct {
	Next string
	Prev string
}

type PaginationCursor struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  *int64 `json:"total,omitempty"`
}

func CreatePaginationCursor(limit, offset, total int64) PaginationCursor {
	return PaginationCursor{
		Limit: limit, Offset: offset,
	}
}

func (c PaginationCursor) String() string {
	if *c.Total < c.Limit+c.Offset {
		return ""
	}

	serialized, err := json.Marshal(c)
	if err != nil {
		return ERRORED_CURSOR
	}

	return base64.StdEncoding.EncodeToString(serialized)
}

func DecodePaginationCursor(cursor *string) (PaginationCursor, error) {
	fmt.Println(cursor)
	if cursor == nil {
		return startingCursor(), nil
	}

	decodedCursor, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return PaginationCursor{}, err
	}
	fmt.Println(decodedCursor)

	var cur PaginationCursor
	if err := json.Unmarshal(decodedCursor, &cur); err != nil {
		return PaginationCursor{}, err
	}
	fmt.Println(cur)

	return cur, nil
}

func startingCursor() PaginationCursor {
	return CreatePaginationCursor(DEFAULT_LIMIT, DEFAULT_START_OFFSET, 0)
}

func (c *PaginationCursor) Next(total int64) PaginationCursor {
	c.Total = &total

	if *c.Total <= c.Limit + c.Offset {
		c.Offset = *c.Total
	} else {
		c.Offset += c.Limit
	}

	return *c
}

func MakePaginationInfo(next PaginationCursor, prev PaginationCursor) PaginationInfo {
	prev.Total = next.Total
	return PaginationInfo{
		Next: next.String(),
		Prev: prev.String(),
	}
}
