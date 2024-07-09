package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) printStory(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Print(node.text)
	fmt.Println()

	if node.yesPath != nil {
		node.yesPath.printStory(depth + 1)
	}
	if node.noPath != nil {
		node.noPath.printStory(depth + 1)
	}
}

func (node *storyNode) play() {
	fmt.Println(node.text)

	if node.yesPath == nil && node.noPath == nil {
		fmt.Println("Game Over!")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		answer := scanner.Text()

		if answer == "yes" {
			node.yesPath.play()
			break
		} else if answer == "no" {
			node.noPath.play()
			break
		} else {
			fmt.Println("Invalid input, please type yes or no")
		}
	}

}

func main() {
	root := storyNode{"You are at the entrace to a dark cave, do you want to go in?", nil, nil}
	yesAnswer := storyNode{"Do you want to take a few more steps?", nil, nil}
	winning := storyNode{"You have won!", nil, nil}
	losing := storyNode{"You have lost!", nil, nil}

	root.yesPath = &yesAnswer
	root.noPath = &losing
	yesAnswer.yesPath = &winning

	//root.play()
	root.printStory(0)
}
