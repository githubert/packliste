package main

// TODO: Warn if undefined items are referenced
// TODO: Compact items defined multiple times in a kit

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
	"math"
	"strconv"
)

type Item struct {
	Name	string
	Note	string
	Weight	float64
	Scale	int
	Number	int
	Spare	bool
	Limit	int
}

type Kit struct {
	Name	string
	Items	[]string
}

type List struct {
	Kits	[]Kit
	Items	[]Item
}

var kits map[string]Kit
var items map[string]Item
var notes []string

func main() {

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s filename days kit [kit] ...\n", os.Args[0])
		os.Exit(1)
	}
	
	// arguments
	
	filename := os.Args[1]
	
	days, err := strconv.Atoi(os.Args[2])
	
	if err != nil {
		log.Fatal(err)
	}

	// open JSON data

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	
	// parse data
	
	dec := json.NewDecoder(file)
	
	var l List
	err = dec.Decode(&l)
		
	if err != nil {
		log.Fatal(err)
	}
	
	// make maps of kits and items
	
	kits = make(map[string]Kit, 16)
	
	for _, kit := range l.Kits {
		kits[kit.Name] = kit
	}
	
	items = make(map[string]Item, 64)
	
	for _, item := range l.Items {
		items[item.Name] = item
	}

	// iterate over kit names provided on command line

	var weight float64

	for _, kitname := range os.Args[3:] {
		weight += useKit(kitname, days)	
	}
	
	fmt.Printf("\nTotal weight: %.3fkg\n", weight / 1000)
	
	fmt.Println("\n\nNotes: ")
	
	for _, note := range notes {
		fmt.Printf("* %s\n", note)
	}

}

func useKit(kitname string, days int) float64 {

	var weight float64
	var qty int
	
	fmt.Printf("=== %s ===\n", kitname)
	
	for _, itemname := range kits[kitname].Items {
	
		item, present := items[itemname]
	
		if !present {
			continue
		}

		// Scale == 0 means there will be item.Number items,
		// independent of the number of days
		if item.Scale == 0 {
			qty = item.Number
		} else {
			qty = int( math.Ceil(float64(item.Number) / float64(item.Scale) * float64(days)))
		}
		
		// Spare == true means there will one spare set added
		if item.Spare {
			qty += item.Number
		}
		
		// TODO: Should this be (item.Limit * item.Number), or is the
		// current solution fine? Unsure. Need to use brain.
		if item.Limit > 0 && qty > (item.Limit) {
			qty = item.Limit
		}

		weight += float64(qty) * item.Weight
		
		if item.Note != "" {
			notes = append(notes, item.Name + ": " + item.Note)
		}
		
		fmt.Printf("%d x %s\n", qty, item.Name)
	}

	fmt.Printf("--- %.3fkg\n", weight/1000)
	
	return weight
}
