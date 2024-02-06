package main

import (
	_ "fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	title string
	body  string
}

type API int

var database []Item

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for i, val := range database {
		if val.title == edit.title {
			database[i] = Item{edit.title, edit.body}
			changed = database[i]
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for i, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener Error", err)
	}

	log.Printf("Serving RPC on PORT %d", 4040)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error Serving: ", err)
	}

	// fmt.Println("Initial Database: ", database)
	// a := Item{"first", "first item"}
	// b := Item{"second", "second item"}
	// c := Item{"third", "third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("Second Database: ", database)

	// DeleteItem(b)
	// fmt.Println("Third Database: ", database)

	// EditItem("third", Item{"fourth", "new fourth item"})
	// fmt.Println("Fourth Database: ", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)
}
