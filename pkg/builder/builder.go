package builder

import (
	"fmt"
	"sync"
)

type StructBuilder struct {
	mtx  *sync.Mutex
	keys map[string]DataType

	// could be useful for determining the indentation
	maxKeyLen      int
	maxDataTypeLen int
}

func NewStruct() *StructBuilder {
	return &StructBuilder{
		mtx:  &sync.Mutex{},
		keys: make(map[string]DataType),
	}
}

func (sb *StructBuilder) GetKey() map[string]DataType {
	return sb.keys
}

func (sb *StructBuilder) MaxKeyLen() int {
	return sb.maxKeyLen
}

func (sb *StructBuilder) MaxDataTypeLen() int {
	return sb.maxDataTypeLen
}

func (sb *StructBuilder) AddKey(key string, dataType DataType, nested bool) error {
	if key == "" || dataType == "" {
		return fmt.Errorf("'key or data type is empty")
	}
	val, ok := sb.CheckKey(key)
	if ok && val != dataType {
		dataType = TypeInterface
	}

	sb.mtx.Lock()
	defer sb.mtx.Unlock()
	if !nested {
		if len(key) > sb.maxKeyLen {
			sb.maxKeyLen = len(key)
		}
		if len(dataType) > sb.maxDataTypeLen {
			sb.maxDataTypeLen = len(dataType)
		}
	}

	sb.keys[key] = dataType
	return nil
}

func (sb *StructBuilder) CheckKey(key string) (DataType, bool) {
	if k, ok := sb.keys[key]; ok {
		return k, true
	}
	return "", false
}
