package utils

import "fmt"

type byteUnit struct {
	Name  string
	Value uint64
}

var units = []byteUnit{
	{"TB", 1 << 40},
	{"GB", 1 << 30},
	{"MB", 1 << 20},
	{"KB", 1 << 10},
}

func FormatBytes(bytes uint64) string {
	for _, unit := range units {
		if bytes >= unit.Value {
			return fmt.Sprintf("%.2f %s", float64(bytes)/float64(unit.Value), unit.Name)
		}
	}
	return fmt.Sprintf("%d B", bytes)
}
