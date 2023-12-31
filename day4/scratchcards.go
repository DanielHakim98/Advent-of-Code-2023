package day4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DanielHakim98/aoc/utils"
)

type Day4 struct{}

func (d *Day4) PartOne(filename string, reader utils.AocReader) int {
	lines, err := reader(4, filename)
	if err != nil {
		log.Fatal(err)
	}

	var totalPoints int
	for _, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")

		// collect win numbers
		left := nums[0]
		wins := make(map[string]string)
		num, isDigit := "", false
		for _, char := range left {
			if char >= '0' && char <= '9' {
				if !isDigit {
					isDigit = true
				}
				num += string(char)
			} else {
				if isDigit {
					wins[num] = num
				}
				isDigit = false
				num = ""
			}
		}

		// check available numbers
		right := nums[1]
		var count, sum int
		num, isDigit = "", false
		for k, char := range right {
			if char >= '0' && char <= '9' {

				// if number then set to number
				if !isDigit {
					isDigit = true
				}
				// collect number
				num += string(char)

				// check if number locates at last index
				if k == len(right)-1 {
					_, ok := wins[num]
					if ok {
						if count == 0 {
							sum += 1
						} else {
							sum = sum * 2
						}
						count++
					}
				}
			} else {
				// If current 'char' is not not a number
				// then check isDigit is true (if true then previous is number)
				if isDigit {
					_, ok := wins[num]
					if ok {
						if count == 0 {
							sum += 1
						} else {
							sum = sum * 2
						}
						count++
					}
				}
				// reset isDigit and num
				isDigit = false
				num = ""
			}

		}

		fmt.Println(wins)
		fmt.Println("count: ", count)
		fmt.Println("sum: ", sum)
		fmt.Println()
		totalPoints += sum
	}
	return totalPoints
}

type ScratchCard struct {
	Id      int
	Matches int
}

type CardId int

func (d *Day4) PartTwo(filename string, reader utils.AocReader) int {
	lines, err := reader(4, filename)
	if err != nil {
		log.Fatal(err)
	}

	scratchCards := make(map[CardId]ScratchCard)
	for _, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")
		// collect total wins
		left := nums[0]
		wins := make(map[string]string)
		num, isDigit := "", false
		for _, char := range left {
			if char >= '0' && char <= '9' {
				if !isDigit {
					isDigit = true
				}
				num += string(char)
				// we dont' consider last index here because
				// it's always empty spaces
			} else {
				if isDigit {
					wins[num] = num
				}
				isDigit = false
				num = ""
			}
		}

		right := nums[1]
		var count int
		num, isDigit = "", false
		for k, char := range right {
			if char >= '0' && char <= '9' {
				// if number then set to number
				if !isDigit {
					isDigit = true
				}
				// collect number
				num += string(char)
				// check if a number locates at last index
				if k == len(right)-1 {
					_, ok := wins[num]
					if ok {
						count++
					}
				}
			} else {
				// If current 'char' is not not a number
				// then check isDigit is true (if true then previous is number)
				if isDigit {
					_, ok := wins[num]
					if ok {
						count++
					}
				}
				// reset isDigit and num
				isDigit = false
				num = ""
			}
		}

		id := d.GetCardNumber(card)
		scratchCards[CardId(id)] = ScratchCard{id, count}
	}

	var total int
	// Recursive Way
	for _, c := range scratchCards {
		// For every card, there must be a least:
		// 1 card + copies of cards won
		total += d.CountCards(c, scratchCards, 0)

	}

	return total
}

func (d *Day4) GetCardNumber(card []string) (id int) {
	// Previously, use strings.Split(card[0]," ")
	// but turns out spaces can vary between " " or "  " or "   "
	// so that's why we use strings.Field here to split by spaces
	// regardless of the space length
	id, _ = strconv.Atoi(strings.TrimSpace(strings.Fields(card[0])[1]))
	return
}

func (d *Day4) CountCards(c ScratchCard, ref map[CardId]ScratchCard, total int) int {
	// Everytime a card is passed into this method
	// then it's one card
	total++
	if c.Matches == 0 {
		return total
	}

	limit := c.Id + c.Matches
	cur := c.Id + 1

	for ; cur <= limit; cur++ {
		curCard, ok := ref[CardId(cur)]
		if ok {
			// For every copies of card won by a card,
			// pass those card into this method recursively
			total += d.CountCards(curCard, ref, 0)
		}
	}

	return total
}

/*
	var cards []ScratchCard
	instances := make([]int, len(cards))
	for idx, card := range cards {
		instances[idx] += 1

		cursor := card.Id + 1
		limit := card.Matches + card.Id
		fmt.Println("card: ", card)
		fmt.Println("start: ", cursor)
		fmt.Println("limit: ", limit)

		for ; cursor <= limit; cursor++ {
			// if cursor >= len(cards) {
			// 	continue
			// }
			cursorIdx := cursor - 1
			fmt.Println("child card: ", cards[cursorIdx])

			instances[cursorIdx] += 1
			fmt.Println(instances)

		}
		fmt.Println()
	}

	fmt.Println(instances)
*/
