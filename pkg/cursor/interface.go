package cursor

type cursor map[string]interface{}

type Cursor interface {
  CreateCursor(...any) (cursor, error)
  Encode() string
  startingCursor() cursor
}
