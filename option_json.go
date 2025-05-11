package shine

import "encoding/json"

type JSONOption[T any] [1]Option[T]

func (o *JSONOption[T]) UnmarshalJSON(data []byte) error {
	var v *T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var opt Option[T]
	if v == nil {
		opt = NewNone[T]()
	} else {
		opt = NewSome(*v)
	}

	*o = [1]Option[T]{opt}
	return nil
}

func (o JSONOption[T]) MarshalJSON() ([]byte, error) {
	if o, ok := o[0].(Some[T]); ok {
		return json.Marshal(o.Value())
	}

	return []byte(`null`), nil
}

func (o JSONOption[T]) Option() Option[T] {
	return o[0]
}
