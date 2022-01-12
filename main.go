package main

import (
	"fmt"
	"log"

	"github.com/alanadiprastyo/learn-by-code/kube"
)

func main() {

	// key value [key] : [value]
	// myMap := make(map[string]string)

	// myMap["name"] = "Alan"
	// myMap["kota"] = "Depok"

	//Map tanpa make
	// myMap := map[string]string{
	// 	"nama": "Alan",
	// 	"kota": "Depok",
	// }

	// map dengan interface kosong
	// makanan := map[string]interface{}{
	// 	"nama":  "nasi goreng",
	// 	"harga": 15000,
	// 	"komposisi": map[string]string{
	// 		"gurih": "micin",
	// 		"pedas": "cabe",
	// 		"manis": "kecap",
	// 	},
	// }

	// fmt.Println(fmt.Sprintf("menu makan siang hari ini adalah %s dengan harga %d dimana komposisinya adalah %s, %s, %s",
	// 	makanan["nama"],
	// 	makanan["harga"],
	// 	makanan["komposisi"].(map[string]string)["gurih"],
	// 	makanan["komposisi"].(map[string]string)["pedas"],
	// 	makanan["komposisi"].(map[string]string)["manis"],
	// ))

	name, harga, err := kube.Makanan()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("makanan hari ini adalah %s dengan harga %d", name, harga))
	}

	age, err := kube.Umur(1995)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("umurnya adalah %d tahun", age))
	}
}
