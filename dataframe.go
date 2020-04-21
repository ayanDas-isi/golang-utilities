package main

import (
	"bufio"
	"context"
	"dataframe-go/imports"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
)

//https://github.com/rocketlaunchr/dataframe-go
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
	//s := dataframe.NewSeriesInt64("random", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	//sf, _ := s.ToSeriesFloat64(ctx)
	//mean := stat.Mean(sf.Values, nil)
	//fmt.Println(mean)
	read_csv()
}

func read_csv() {

	ctx := context.Background()
	//df, _ := imports.LoadFromCSV(ctx, strings.NewReader(csvStr))
	df, _ := imports.LoadFromCSV(ctx, strings.NewReader(read_file("dr.csv")))
	fmt.Println(df.Table())
	write_csv(df, "output.csv")
}
func write_csv(df dataframe, fname string) {
	fo, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(fo)
	exports.ExportToCSV(ctx, w, df)
}

func read_file(fname string) string {
	b, err := ioutil.ReadFile(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
		return err.Error()
	}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'
	return str
}
