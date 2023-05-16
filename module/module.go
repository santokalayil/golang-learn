package sub_module

import "fmt"

func ModuleFunction(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
