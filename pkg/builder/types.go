package builder

type DataType string

const (
	TypeString      DataType = "string"
	TypeInt         DataType = "int"
	TypeInt32       DataType = "int32"
	TypeInt64       DataType = "int64"
	TypeFloat       DataType = "float"
	TypeFloat32     DataType = "float32"
	TypeFloat64     DataType = "float64"
	TypeInterface   DataType = "interface{}"
	TypeEmptyStruct DataType = "type struct{} {\n %s \n}\n"
)

type AutoStruct struct {
	Check   bool `json:"check"`
	MyField struct {
		Test  []map[string]interface{} `json:"test"`
		Test3 struct {
			Some string `json:"some"`
		} `json:"test3"`
	} `json:"my_field"`
}
