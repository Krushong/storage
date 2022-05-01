package main

import (
	"fmt"
	"github.com/Krushong/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredfile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it is restored", restoredfile)
}
