package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Part1() int {
	var sum int
	for _, name := range Data {
		if sectorId, _, ok := validateRoomName(name); ok {
			sum += sectorId
		}
	}
	return sum
}

func Part2() int {
	roomNames := map[string]int{}
	for _, name := range Data {
		if sectorId, letters, ok := validateRoomName(name); ok {
			roomName := shiftCipher(letters, sectorId)
			roomNames[roomName] = sectorId
		}
	}
	return roomNames["northpole-object-storage"]
}

func shiftCipher(name string, shift int) string {
	start := int('a')
	end := int('z')

	value := []rune{}
	for _, letter := range name {
		if letter == '-' {
			value = append(value, '-')
			continue
		}

		l := int(letter)
		value = append(value, rune(((l-start+shift)%(end-start+1))+start))
	}

	return string(value)
}

func validateRoomName(name string) (int, string, bool) {
	re := regexp.MustCompile(`^(.+)-(\d+)\[(.+)\]$`)
	match := re.FindStringSubmatch(name)

	letters := strings.Replace(match[1], "-", "", -1)
	sectorId, _ := strconv.Atoi(match[2])
	checksumLetters := match[3]

	pairList := map[string]int{}
	for _, letter := range letters {
		pairList[string(letter)] += 1
	}

	checksum := strings.Join(rankMapStringInt(pairList)[:5], "")

	return sectorId, match[1], checksum == checksumLetters
}

func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}

		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
