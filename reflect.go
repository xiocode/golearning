package golearning

import (
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	client := new(http.Client)
	response := reflect.ValueOf(client).MethodByName("Get").Call([]reflect.Value{reflect.ValueOf("http://www.baidu.com")})[0].Interface()
	resp := response.(*Response)
	fmt.Println(resp)
}
