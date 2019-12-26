package main

import (
	"fmt"
	"sort"
	"strings"
)

func Meeting(s string) string {

	names := Names{}
	// seperate by lastNames
	for _, name := range strings.Split(s, ";") {
		firstName, lastName := splitName(name)
		newName := Name{firstName, lastName}
		names.append(newName)
	}

	names.sortByLastName()
	familyNames, orderOfLastNames := names.splitByFamilyName()
	fmt.Println(orderOfLastNames)

	sortedNames := Names{}

	for _, lastName := range orderOfLastNames {
		names := familyNames[lastName]
		sort.SliceStable(names, func(i, j int) bool {
			return names[i].firstName < names[j].firstName
		})
		sortedNames.appendNames(names)
	}

	return sortedNames.print()
}

type Names struct {
	names []Name
}

func (n *Names) append(name Name) {
	name.toUpperCase()
	n.names = append(n.names, name)
}

func (n *Names) appendNames(name []Name) {
	n.names = append(n.names, name...)
}

func (n *Names) sortByLastName() {
	sort.SliceStable(n.names, func(i, j int) bool {
		return n.names[i].lastName < n.names[j].lastName
	})
}

func (n Names) print() string {
	s := ""
	for _, name := range n.names {
		s += name.print()
	}
	return s
}

func (n Names) splitByFamilyName() (map[string][]Name, []string) {
	lastNames := make(map[string][]Name)
	orderOfLastNames := []string{}

	for _, name := range n.names {
		if _, ok := lastNames[name.lastName]; !ok {
			lastNames[name.lastName] = []Name{name}
			orderOfLastNames = append(orderOfLastNames, name.lastName)
		} else {
			lastNames[name.lastName] = append(lastNames[name.lastName], name)
		}
	}

	return lastNames, orderOfLastNames
}

type Name struct {
	firstName string
	lastName  string
}

func (n *Name) toUpperCase() {
	n.lastName = strings.ToUpper(n.lastName)
	n.firstName = strings.ToUpper(n.firstName)
}

func (n Name) print() string {
	return fmt.Sprintf("(%s, %s)", n.lastName, n.firstName)
}

func splitName(name string) (firstName, lastName string) {
	x := strings.Split(name, ":")
	firstName = x[0]
	lastName = x[1]
	return
}
