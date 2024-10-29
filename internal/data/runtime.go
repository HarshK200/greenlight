package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// NOTE: wrapping the string in quotes for it to be a valid json string (could've used \ for escape chars)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(JSONData []byte) error {
	unquotedJSON, err := strconv.Unquote(string(JSONData)) // NOTE: unwarpping the quotes 'cause json string
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	splits := strings.Split(unquotedJSON, " ")
	if len(splits) != 2 || splits[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	runtimeInt, err := strconv.ParseInt(splits[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

    // NOTE: can't do int32() typecaste since Runtime is of custom Runtime type even though it's kinda an int32 alias
    *r = Runtime(runtimeInt) // NOTE: destructure the pointer here

	return nil
}
