package main

import (
	"fmt"
	"hash/fnv"
)

/*

key:value
hash funcion
hash table

*/

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	fmt.Println(hash("saurabh") % 20)
	fmt.Println(hash("saurabh") % 20)
	fmt.Println(byte("2" ))


}
