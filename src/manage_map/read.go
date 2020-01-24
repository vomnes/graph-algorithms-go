package manageMap

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// StoreInputMap read the path file and store the data in an structure
func StoreInputMap(path string) *DataMap {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapData := DataMap{}
	m := mapData.New()
	y, x := 0, 0
	for scanner.Scan() {
		x = 0
		m.AllocNewY(y)
		strings.Map(func(r rune) rune {
			m.SetData(y, x, r)
			if r == StarKey {
				m.SetStart(y, x)
			}
			if r == EndKey {
				m.SetEnd(y, x)
			}
			x++
			return r
		}, scanner.Text())
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}