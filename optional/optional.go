package optional

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Optional[T any] struct {
	value   T
	isValid bool
}

func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{
		value:   v,
		isValid: true,
	}
}

func (o *Optional[T]) IsValid() bool {
	return o.isValid
}

func (o *Optional[T]) ShouldGet() T {
	return o.value
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.IsValid() {
		return json.Marshal(nil)
	}
	return json.Marshal(o.value)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	var v *T

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v == nil {
		o.isValid = false
		return nil
	}

	o.isValid = true
	o.value = *v
	return nil
}

func (o *Optional[T]) Scan(value interface{}) error {
	v, ok := value.(T)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	o.value = v
	o.isValid = true
	return nil
}

// Value return json value, implement driver.Valuer interface
func (o Optional[T]) Value() (driver.Value, error) {
	if o.isValid {
		return o.value, nil
	}
	return nil, nil
}
