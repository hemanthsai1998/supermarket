package supermarket

import (
	"errors"
	s "strings"
)

var item = map[string]interface{}{
	"flour":  90.10,
	"cheese": 35.60,
	"meat":   349.99,
	"onions": 30.50,
}

// item = make(map[string]interface{})
// item["a"] = 475.60
// item["b"] = 67.5

func Get(name string) (interface{}, error) {
	s.ToLower(name)
	if item[name] != nil {
		//fmt.Println("Price of", name, "is", item[name])
		return item[name], nil
	} else {
		return nil, errors.New("Item not found")
	}
}

func Post(name string, value interface{}) error {
	s.ToLower(name)
	if item[name] != nil {
		//fmt.Println("Item already exists")
		return errors.New("Item already exists")
	} else {
		if item == nil {
			item = make(map[string]interface{})
			item[name] = value
		} else {
			item[name] = value
		}
		return nil
	}

}

func Put(name string, value interface{}) error {
	s.ToLower(name)
	if item[name] == nil {
		return errors.New("Item does not exist")
	} else {
		item[name] = value
		return nil
	}
}

func Delete(name string) error {
	s.ToLower(name)
	if item[name] != nil {
		delete(item, name)
		return nil
	} else {
		return errors.New("Item not found")
		//fmt.Println("Item doesn't exist or already deleted")
	}
}
func Print() map[string]interface{} {
	// fmt.Println("Items in supermarket:")
	// fmt.Println(item)
	return item
}
