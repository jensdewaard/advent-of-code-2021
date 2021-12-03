package day2

import (
	"log"
	"strings"

	shared_strings "github.com/jensdewaard/advent-of-code-2021/shared/strings"
)

func Move(start Location, command Command) Location {
	switch command.Direction {
	case Forward:
		return Location{
			Depth:      start.Depth,
			Horizontal: start.Horizontal + command.Length,
		}
	case Up:
		return Location{
			Depth:      start.Depth - command.Length,
			Horizontal: start.Horizontal,
		}
	case Down:
		return Location{
			Depth:      start.Depth + command.Length,
			Horizontal: start.Horizontal,
		}
	}
	log.Fatalf("unknown direction: %v", command.Direction)
	return start
}

func MoveWithAim(start LocationWithAim, command Command) LocationWithAim {
	switch command.Direction {
	case Forward:
		return LocationWithAim{
			Depth:      start.Depth + (start.Aim * command.Length),
			Horizontal: start.Horizontal + command.Length,
			Aim:        start.Aim,
		}
	case Up:
		return LocationWithAim{
			Depth:      start.Depth,
			Horizontal: start.Horizontal,
			Aim:        start.Aim - command.Length,
		}
	case Down:
		return LocationWithAim{
			Depth:      start.Depth,
			Horizontal: start.Horizontal,
			Aim:        start.Aim + command.Length,
		}
	}
	log.Fatalf("unknown direction: %v", command.Direction)
	return start
}

func FoldCommands(acc Location, cs []Command) Location {
	for _, c := range cs {
		acc = Move(acc, c)
	}
	return acc
}

func FoldCommandsWithAim(acc LocationWithAim, cs []Command) LocationWithAim {
	for _, c := range cs {
		acc = MoveWithAim(acc, c)
	}
	return acc
}

func ParseCommand(s string) Command {
	ss := strings.Split(s, " ")
	switch ss[0] {
	case "forward":
		return Command{
			Forward,
			shared_strings.ParseInt(ss[1]),
		}
	case "up":
		return Command{
			Up,
			shared_strings.ParseInt(ss[1]),
		}
	case "down":
		return Command{
			Down,
			shared_strings.ParseInt(ss[1]),
		}
	default:
		log.Fatalf("could not parse command: %s", ss[0])
	}
	return Command{}
}

func MapStringsToCommands(ss []string) []Command {
	cs := make([]Command, len(ss))
	for i, s := range ss {
		cs[i] = ParseCommand(s)
	}
	return cs
}
