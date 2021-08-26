package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/inancgumus/screen"
)

func add(items []string, to string, h []card, b [][]card, ho []card) (bool, []card, [][]card, []card) {
	var cs []card
	num, err := strconv.Atoi(to)
	if err != nil {
		return false, h, b, ho
	}
	num--
	if num < 0 || num >= len(b) {
		return false, h, b, ho
	}
	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			return false, h, b, ho
		}
		cs = append(cs, c)
		b[num] = append(b[num], c)
	}
	b[num] = sortHand(b[num])
	if !isValid(b[num]) {
		return false, h, b, ho
	}
	for _, c := range cs {
		if isIn(c, h) != -1 {
			h = remove(h, c)
		} else if isIn(c, ho) != -1 {
			ho = remove(ho, c)
		} else {
			return false, h, b, ho
		}
	}
	return true, h, b, ho
}

func place(items []string, h []card, b [][]card, ho []card) (bool, []card, [][]card, []card, int) {
	var cs []card
	tot := 0
	for _, item := range items {
		c, fail := processItem(item)
		if fail == 1 {
			return false, h, b, ho, tot
		}
		cs = append(cs, c)
	}
	cs = sortHand(cs)
	if !isValid(cs) {
		return false, h, b, ho, tot
	}
	for _, c := range cs {
		if isIn(c, h) != -1 {
			h = remove(h, c)
			tot += c.number
		} else if isIn(c, ho) != -1 {
			ho = remove(ho, c)
		} else {
			return false, h, b, ho, tot
		}
	}
	b = append(b, cs)
	return true, h, b, ho, tot
}

func draw() {
	screen.Clear()
	randomIndex := rand.Intn(len(pool))
	pick := pool[randomIndex]
	hands[turn] = sortHand(append(hands[turn], pick))
	pool = removei(pool, randomIndex)
	fmt.Print("You drew: ")
	printCard(pick)
	fmt.Println("\nYour new hand is:")
	printCards(hands[turn])
	fmt.Scanln()
}

func exchange(item string, to string, h []card, b [][]card, ho []card) (bool, []card, [][]card, []card) {
	num, err := strconv.Atoi(to)
	if err != nil {
		return false, h, b, ho
	}
	num--
	if num < 0 || num >= len(b) {
		return false, h, b, ho
	}
	c, fail := processItem(item)
	if fail == 1 {
		return false, h, b, ho
	}
	cjok := c
	cjok.joker = 1
	i := isIn(cjok, b[num])
	if i == -1 {
		cjok.joker = 2
	}
	i = isIn(cjok, b[num])
	if i == -1 {
		return false, h, b, ho
	}
	h = remove(h, c)
	b[num][i].joker = 0
	ho = append(ho, cleanJoker(cjok))
	return true, h, b, ho
}
