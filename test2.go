package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
   var sorted []int
   for _,val :=range arr{
       sorted=append(sorted,int(val))

   }
sort.Ints(sorted)
var sum_min ,sum_max int64
for i:=0;i<4;i++{
    sum_min=sum_min+int64(sorted[i])
    sum_max=sum_max+int64(sorted[i+1])
}

fmt.Print(sum_min," ",sum_max)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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
