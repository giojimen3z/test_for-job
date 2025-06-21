package challenges

import (
	"strconv"
	"strings"
)

type BrowserNavigator struct {
	backStack    []string
	forwardStack []string
	currentPage  string
}

func NewBrowserNavigator(start string) *BrowserNavigator {
	return &BrowserNavigator{
		currentPage:  start,
		backStack:    []string{},
		forwardStack: []string{},
	}
}

func (b *BrowserNavigator) hopTo(page string) string {
	b.backStack = append(b.backStack, b.currentPage)
	b.currentPage = page
	b.forwardStack = []string{}
	return b.currentPage
}

func (b *BrowserNavigator) backtrack(steps int) string {
	for steps > 0 && len(b.backStack) > 0 {
		b.forwardStack = append(b.forwardStack, b.currentPage)
		b.currentPage = b.backStack[len(b.backStack)-1]
		b.backStack = b.backStack[:len(b.backStack)-1]
		steps--
	}
	return b.currentPage
}

func (b *BrowserNavigator) leapForward(steps int) string {
	for steps > 0 && len(b.forwardStack) > 0 {
		b.backStack = append(b.backStack, b.currentPage)
		b.currentPage = b.forwardStack[len(b.forwardStack)-1]
		b.forwardStack = b.forwardStack[:len(b.forwardStack)-1]
		steps--
	}
	return b.currentPage
}

func SimulateBrowserNavigation(commands []string) []string {
	log := []string{}
	var browser *BrowserNavigator

	for _, cmd := range commands {
		parts := strings.Split(cmd, " ")
		switch parts[0] {
		case "BrowserNavigator":
			browser = NewBrowserNavigator(parts[1])
			log = append(log, "Starting page: "+browser.currentPage)
		case "hopTo":
			log = append(log, "Current page: "+browser.hopTo(parts[1]))
		case "backtrack":
			steps, _ := strconv.Atoi(parts[1])
			log = append(log, "Current page: "+browser.backtrack(steps))
		case "leapForward":
			steps, _ := strconv.Atoi(parts[1])
			log = append(log, "Current page: "+browser.leapForward(steps))
		}
	}
	return log
}
