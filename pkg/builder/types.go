/*
Copyright © 2022 Samit <samitghimire@gmail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
