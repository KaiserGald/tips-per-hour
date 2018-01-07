// Package cli
// 6 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package cli

import (
	"fmt"
	"reflect"
	"regexp"
)

var (
	comList CommandList
)

// Command contains data about a particular command
type Command struct {
	name  string
	args  []string
	usage string
	flag  bool
	val   interface{}
}

// CommandList contains a list of all the commands
type CommandList struct {
	Commands []Command
}

func init() {
	comList = CommandList{}
}

// SetString associates a string variable to the specified command
func SetString(val *string, name string, initial string, usage string) {
	com := Command{}
	com.name = name
	com.usage = usage
	*val = initial
	com.val = val
	fmt.Println(reflect.ValueOf(com.val).Elem())
	if isFlag(name) {
		com.flag = true
	} else {
		com.flag = false
	}
	comList.Commands = append(comList.Commands, com)
}

// SetBool associates a bool variable to the specified command
func SetBool(val *bool, name string, initial bool, usage string) {

}

// SetValue sets the value of the command
func (c *Command) SetValue(i interface{}) {
	c.val = i
}

// SetInt associates an int variable to the specified command
func SetInt(val *int, name string, initial int, usage string) {

}

// Create Command
func CreateCommand(val interface{}, name string, initial interface{}, usage string) error {
	if reflect.TypeOf(val) != reflect.PtrTo(reflect.TypeOf(initial)) {
		fmt.Println("Error")
	} else {
		fmt.Println("Equal")
		com := Command{}
		com.name = name
		com.usage = usage
		if reflect.TypeOf(val).String() == "*string" {
			fmt.Println("It is string")
		} else if reflect.TypeOf(val).String() == "*bool" {
			fmt.Println("It is bool")
		} else if reflect.TypeOf(val).String() == "*int" {
			fmt.Println("It is int")
		}
	}

	return nil
}

func isFlag(s string) bool {
	matched, _ := regexp.MatchString("-(([a-zA-Z](\\d*))+)", s)
	return matched
}
