package parserlib

import (
	"regexp"
	"sort"
	"strconv"
)

func GetLargestSrc(value string) (string, error) {
	rx, err := regexp.Compile(`(https://.+?\.jpg) (\d+)w`)
	if err != nil {
		return "", err
	}

	type srcset struct {
		URL  string
		Size int
	}

	list := []srcset{}

	for _, item := range rx.FindAllStringSubmatch(value, -1) {
		size, err := strconv.Atoi(item[2])
		if err != nil {
			continue
		}

		list = append(list, srcset{
			URL:  item[1],
			Size: size,
		})
	}

	if len(list) == 0 {
		return "", ErrNotFound
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Size > list[j].Size
	})

	return list[0].URL, nil
}
