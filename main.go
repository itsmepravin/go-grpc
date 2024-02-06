package main

import "fmt"

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	return getItem
}

func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changed Item

	for i, val := range database {
		if val.title == edit.title {
			database[i] = edit
			changed = edit
		}
	}

	return changed
}

func DeleteItem(item Item) Item {
	var del Item

	for i, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	return del
}

func main() {
	fmt.Println("Initial Database: ", database)
	a := Item{"first", "first item"}
	b := Item{"second", "second item"}
	c := Item{"third", "third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("Second Database: ", database)

	DeleteItem(b)
	fmt.Println("Third Database: ", database)

	EditItem("third", Item{"fourth", "new fourth item"})
	fmt.Println("Fourth Database: ", database)

	x := GetByName("fourth")
	y := GetByName("first")
	fmt.Println(x, y)

}
