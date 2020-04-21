package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/big"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int32) {
 f := new(big.Int)
f.SetString("1", 10) 
tmp:=new(big.Int)
for i:=1;i<=int(n);i++{
    tmp.SetInt64(int64(i))
   
    f.Mul(f,tmp)
}

fmt.Println(f)

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    extraLongFactorials(n)
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
