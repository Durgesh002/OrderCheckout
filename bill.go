package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newbill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"cake": 8, "pizza": 6.5, "french fries": 2},
		tip:   0,
	}

	return b
}

func (b *bill) format() string {
	f := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		f += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	f += fmt.Sprintf("%-25v ...$%v\n", "tip: ", b.tip)
	f += fmt.Sprintf("%-25v ...$%0.2f\n", "total: ", total)

	return f
}

//update tip
func (b *bill) updatetip(tip float64) {
	b.tip = tip
}

//add items
func (b *bill) additems(name string, price float64) {
	b.items[name] = price
}

//delete items
func (b *bill) deleteitems(name string) {
	delete(b.items, name)
}

//save blls
func (b bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill saved to the folder")
}

//creating function to take inputs
func getinput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

//creating bills
func createbill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getinput("Create new bill name: ", reader)

	b := newbill(name)
	fmt.Println("Created bill name - ", b.name)

	return b
}

//various prompts
func prompts(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getinput("Choose an option (a - add item, s - save bill, t - add tip, d - delete item): ", reader)

	switch opt {
	case "a":
		name, _ := getinput("Item name: ", reader)
		price, _ := getinput("Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error while parsing: Please enter a valid number...")
			prompts(b)
		}
		b.additems(name, p)
		fmt.Println("Items added: ", name, p)
		prompts(b)

	case "t":
		tip, _ := getinput("Enter tip value($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Error while parsing: Please enter valid number...")
		}
		b.updatetip(t)
		fmt.Println("Tip added - ", t)
		prompts(b)

	case "s":
		ans, _ := getinput("Do you want to save your bill? [y/n]: ", reader)
		if ans == "y" {
			b.save()
			fmt.Println("Bill saved - ", b.name)
		} else if ans == "n" {
			prompts(b)
		}

	case "d":
		rem, _ := getinput("Imput the item to delete: ", reader)
		b.deleteitems(rem)
		fmt.Println("Item deleted...")
		prompts(b)

	default:
		fmt.Println("oops, choose a valid option...")
		prompts(b)
	}
}
