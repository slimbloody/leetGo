package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Info struct {
	actId int64
	staCode string
	staType int
	staName string
	staNum int
	createTime string
}
func main()  {
	file, err := os.Open("./log.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "sta_code=,", "sta_code=\"\",", -1)
		split1 := strings.Split(line, "source: ")[1]
		split2 := strings.Split(split1, string("\",\"__source__\""))[0]
		split3 := strings.Split(split2, ", target: ")
		source, target := split3[0], split3[1]

		var sourceInfo Info = extractInfo(source)
		var targetInfo Info = extractInfo(target)

		//{act_id=2105010333324564, sta_code=Dongguan, sta_type=6, sta_name=城市, sta_num=2, create_time=2021-05-04 08:48:54.0
		if sourceInfo.actId % 10 == 0 && sourceInfo.staNum != targetInfo.staNum && sourceInfo.actId >= 2105000000000000 && sourceInfo.actId < 2105020000000000{
			log.Println("sourceInfo: ", sourceInfo)
			log.Println("targetInfo: ", targetInfo)
			log.Println("  ")
		}

	}
	file.Close()
}

func MultiSplit(r rune) bool {
	return r == '=' || r == ','
}

func extractInfo(structStr string) Info {
	fieldsSplits := strings.FieldsFunc(structStr, MultiSplit)
	sourceActId, _ := strconv.ParseInt(fieldsSplits[1], 10, 64)
	staType, _ := strconv.Atoi(fieldsSplits[5])
	staNum, _ := strconv.Atoi(fieldsSplits[9])
	var createTime string = fieldsSplits[11][0 : len(fieldsSplits[11])-1]

	info := Info{sourceActId, fieldsSplits[3], staType, fieldsSplits[7], staNum, createTime}
	return info
}
