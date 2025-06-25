package main

import (
	"fmt"
	"log"

	"github.com/DLiuHeSong/sdk/university"
)

func main() {
	sdk, err := university.New("../data/universities.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("所有学校:")
	for _, u := range sdk.ListAll() {
		fmt.Println(u)
	}

	fmt.Println("\n搜索‘北京’的学校:")
	for _, u := range sdk.SearchByProvince("北京") {
		fmt.Println(u)
	}

	fmt.Println("\n搜索名称包含‘工业’的学校:")
	for _, u := range sdk.SearchByName("工业") {
		fmt.Println(u)
	}
}
