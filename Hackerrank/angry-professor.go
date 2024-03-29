package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the angryProfessor function below.
func angryProfessor(k int32, a []int32) string {
    var ontime int32 = 0
    var isCancel string = "YES"
    for _, i := range a {
        if (i <= 0){
            ontime++
        }
    }
    if (ontime >= k){
        isCancel = "NO"
    }
    return isCancel
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nk := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nk[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(nk[1], 10, 64)
        checkError(err)
        k := int32(kTemp)

        aTemp := strings.Split(readLine(reader), " ")

        var a []int32

        for i := 0; i < int(n); i++ {
            aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
            checkError(err)
            aItem := int32(aItemTemp)
            a = append(a, aItem)
        }

        result := angryProfessor(k, a)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
