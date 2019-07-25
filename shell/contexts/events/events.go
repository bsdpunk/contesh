package events

import "time"
import "fmt"

type Event struct {
	// Event Id
	Id int64 `json:"id"`
	// Event Description
	Desc string `json:"description"`
	// Latitude
	Lat float64 `json:"lat"`
	// Longitude
	Lon float64 `json:"lon"`
	//Action func()
	// Theme of event
	Theme string `json:"theme"`
	// When the event occoured
	When time.Time `json:"when"`
	// Unified Name
	Name string `json:"event"`
	// Tags
	Tags []string `json:"tags"`
	// MetaData Id
	Mid int64 `json:"mId"`
}

type Events []Event

func Import() {
	fmt.Println("Not yet Implimented")
	return
}

//func (c CommandsByName) Len() int {
//	return len(c)
//}
//
//func (c CommandsByName) Swap(i, j int) {
//	c[i], c[j] = c[j], c[i]
//}
//
//// FullName returns the full name of the command.
//// For subcommands this ensures that parent commands are part of the command path
//
//func (c Command) Names() []string {
//	names := []string{c.Name}
//
//	if c.ShortName != "" {
//		names = append(names, c.ShortName)
//	}
//
//	return names
//	//return append(names, c.Aliases...)
//}
//func (cs Commands) HasCommand(name string) bool {
//	for _, i := range cs {
//		for _, n := range i.Names() {
//			if n == name {
//				return true
//			}
//		}
//	}
//	return false
//}
//
////func (c Command) HasName(name string) bool {
////	for _, n := range c.Names() {
////		if n == name {
////			return true
////		}
////	}
////	return false
////}
//func (cs Commands) NameIs(name string) Command {
//	for _, c := range cs {
//		if c.ShortName == name || c.Name == name {
//			return c
//		}
//	}
//	c := Command{}
//	return c
//}
