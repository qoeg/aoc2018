package day9

import (
	"container/list"
	"fmt"
)

func highScore(game Game) int {
	var hs int
	for _, p := range game.players {
		if p.score > hs {
			hs = p.score
		}
	}
	return hs
}

func getIndex(length, index, shift int) int {
	if (shift > 0) {
		for i := index; ; i++ {
			if i == length {
				i = 0
			}
			
			if shift == 0 {
				return i
			}
			shift--
		}
	} else if (shift < 0) {
		for i := index; ; i-- {
			if i == -1 {
				i = length-1
			}

			if shift == 0 {
				return i
			}
			shift++
		}
	}
	
	return index
}

func shiftTo(list *list.List, from *list.Element, to int) *list.Element {
	e := from
	if to > 0 {
		for shift := to; shift > 0; shift-- {
			e = e.Next()
			if e == nil {
				e = list.Front()
			}
		}
	} else if to < 0 {
		for shift := to; shift < 0; shift++ {
			e = e.Prev()
			if e == nil {
				e = list.Back()
			}
		}
	}
	return e
	
}

func printTurn(player, index int, circle []int) {
		fmt.Printf("[%d] ", player)
		for i := range circle {
			if i == index {
				fmt.Printf("(%d) ", circle[i])
				continue
			}
			fmt.Printf("%d ", circle[i])
		}
		fmt.Println()
}

func printList(player int, current *list.Element, circle *list.List) {
	fmt.Printf("[%d] ", player)
	for e := circle.Front(); e != nil; e = e.Next() {
		if e == current {
			fmt.Printf("(%d) ", e.Value)
			continue
		}
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}

// GetHighScore gets the high score for a game using Go slices (very slow)
func GetHighScore(game Game) int {
	circle := []int{0}
	player := 0
	index := 0

	for marble := 1; marble <= game.LastMarble; marble++ {
		if marble % 23 == 0 {
			index = getIndex(len(circle), index, -7)
			game.players[player].score += marble + circle[index]

			circle = append(circle[:index], circle[index+1:]...)
		} else {
			index = getIndex(len(circle), index, 2)
			if index == 0 {
				index = len(circle)
			}

			circle = append(circle, 0)
			copy(circle[index+1:], circle[index:])
			circle[index] = marble
		}

		player = marble % len(game.players)
	}

	return highScore(game)
}

// GetHighScoreFast get the high score for a game using a double-ended linked list
func GetHighScoreFast(game Game) int {
	player := 0
	circle := list.New()
	current := circle.PushFront(0)
	
	for marble := 1; marble <= game.LastMarble; marble++ {
		if marble % 23 == 0 {
			e := shiftTo(circle, current, -7)
			current = shiftTo(circle, e, 1)

			value := circle.Remove(e).(int)
			game.players[player].score += marble + value
		} else {
			e := shiftTo(circle, current, 1)
			current = circle.InsertAfter(marble, e)
		}

		player = marble % len(game.players)
	}

	return highScore(game)
}
