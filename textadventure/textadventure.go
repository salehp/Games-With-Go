package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type storyNode struct {
	text    string
	choices []*choice
}

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry I didn't understand that command")
	return node
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, ":", choice.description)
		}
	}
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in a large chamber, deep underground.
	You see three passages leading out. A north passage leads to darkness.
	To the south, a package appears to be heading upwards. 
	The eastern passage appears flatand well travelled.
	`}

	darkRoom := storyNode{text: "It is pitch black, you cannot see a thing"}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your lantern, you can head north or head back south"}

	ghoul := storyNode{text: "While stumbling around in the dark, you are eaten by a ghoul"}

	trap := storyNode{text: "You head down the well travelled path when suddenly a trap door opens and you fall into a pit"}

	treasure := storyNode{text: "You arrive at a small chamber filled with treasures!"}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try To Go Back", &ghoul)
	darkRoom.addChoice("O", "Turn on your lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()
	fmt.Println("The End!")

}
