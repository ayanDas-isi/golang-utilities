package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/rocketlaunchr/dataframe-go"
	"gonum.org/v1/gonum/stat"
)

//go get github.com/rocketlaunchr/dataframe-go

func main() {
	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)

	fmt.Print(df.Table())

	//append

	df.Append(nil, 9, 123.6)

	df.Append(nil, map[string]interface{}{
		"day":   10,
		"sales": nil,
	})

	df.Remove(0)
	fmt.Print(df.Table())

	df.UpdateRow(0, nil, map[string]interface{}{
		"day":   3,
		"sales": 45,
	})
	sks := []dataframe.SortKey{
		{Key: "sales", Desc: true},
		{Key: "day", Desc: true},
	}
	ctx := context.Background()
	df.Sort(ctx, sks)
	fmt.Println(df.Table())

	//iterate
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 3, true}) // Don't apply read lock because we are write locking from outside.

	df.Lock()
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		fmt.Println(*row, vals)
	}
	df.Unlock()
	//stat
	s := dataframe.NewSeriesInt64("random", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	sf, _ := s.ToSeriesFloat64(ctx)
	mean := stat.Mean(sf.Values, nil)
	fmt.Println(mean)
}

func insert_position(score []int32, val int32, pos int32) []int32 {
	fmt.Println(val, pos)
	score = append(score, 0)         // Step 1
	copy(score[pos+1:], score[pos:]) // Step 2
	score[pos] = val                 // Step 3
	return score
}

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(score []int32, alice []int32) []int32 {
	var stand []int32
	var pos int32
	for _, val := range alice {
		pos = 1
		//fmt.Println(score)
		if score[0] <= val {
			stand = append(stand, pos)
			continue
		}
		for i := 0; i < len(score)-1; i++ {

			if score[i] > val && score[i+1] <= val {
				stand = append(stand, pos+1)
				break
			}
			if i > 0 && score[i] < score[i-1] {
				pos = pos + 1
				fmt.Println(pos)
			}

		}
		if score[len(score)-1] > val {
			stand = append(stand, pos+1)
		}

		fmt.Println("---------------")
	}
	fmt.Println(score)

	return stand
}

func main2() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
