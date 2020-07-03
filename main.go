package main

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/KeThichDua/ex1go/src"
)

type serve struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func list(result []serve) {
	for i := range result {
		log.Println("name: " + result[i].Name)
		log.Println("class: " + result[i].Class)
	}
}

func listaddress(result []serve) {
	for i := range result {
		fmt.Printf("địa chỉ: %p \n", &result[i])
	}
}

func main() {
	//mở jsonFile
	jsonFile, err := os.Open("serve.json")

	//nếu os.Open trả về lỗi
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n	Mở thành công serve.json")

	//
	defer jsonFile.Close()

	//đọc jsonFile dưới 1 mảng byte
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []serve
	json.Unmarshal(byteValue, &result)

	fmt.Println("\n	2. Dữ liệu file json")
	list(result)

	fmt.Println("\n	3. Lọc class chứa \"admin\" ")
	for i := range result {
		class := strings.ToUpper(result[i].Class)
		admin := strings.ToUpper("admin")
		if strings.Contains(class, admin) == true {
			fmt.Println("name: " + result[i].Name)
			fmt.Println("class: " + result[i].Class)
		} else {
			continue
		}

	}

	fmt.Println("\n	4. Thêm 1 phần tử")
	var x serve
	x.Name = "fileCustome"
	x.Class = "org.cofax.cds.FileServlet.Custome"
	result = append(result, x)
	fmt.Println("	Sau khi thêm phần tử: ")
	list(result)

	fmt.Println("\n	5. In ra địa chỉ item trong slice")
	listaddress(result)

	src.Run6()
	src.Run7()

}
