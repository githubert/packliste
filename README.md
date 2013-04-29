packliste
=========

A simplistic tool for managing pack lists. Look at example.json to get a rough
idea on how to use it.

Options for items:

* Name	id/name
* Note	note to be shown in the summary ("charge batteries")
* Weight	weight in grams
* Scale	after how many days a new set of 'Number' items will be added, 0 if unaffected by days. For example: You might want to add a pair of socks for each day you are away. If 0: Your population of cameras will be constant.
* Number	usually just 1, but if you want to add three snacks per day, set this to three. 
* Spare	one set of additional 'Number' items will be added
* Limit	maximum number of items (e.g. you might have only 8 t-shirts)


Sample output using example.json

    $ go run packliste.go example.json 7 general clothing tent "first aid" stuff
     
     === general ===
     1 x phone
     1 x phone charger
     1 x keys
     1 x train ticket
     --- 0.267kg
     === clothing ===
     3 x t-shirt
     2 x pullover
     2 x trousers
     --- 3.010kg
     === tent ===
     1 x tent
     1 x sleeping bag
     1 x sleeping pad
     --- 3.250kg
     === first aid ===
     1 x first aid set
     --- 0.660kg
     === stuff ===
     1 x multitool
     1 x knife
     1 x camera
     --- 1.020kg
     
     Total weight: 8.207kg
     
     
     Notes: 
     * phone: charge batteries, check prepaid account
     * camera: charge batteries
 
