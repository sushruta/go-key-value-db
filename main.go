package main

import "fmt"
import "github.com/sushruta/go-key-value-db/memtable"

func main() {
        bst := memtable.BstConstructor()
        bst.Insert("handbag", 8786)
        bst.Insert("handlebars", 3869)
        bst.Insert("handicap", 70836)
        bst.Insert("handkerchief", 16433)

        fmt.Println(bst.Inorder())
}
