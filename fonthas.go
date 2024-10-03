// ttfフォントが対応していない文字をCSVの文字列から検出する
// 1行目のデータを各カラム名として扱う

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: fonthas <ttf path> <csv path>")
		os.Exit(1)
	}
	ttfPath := flag.Arg(0)
	ttfData, err := loadTtf(ttfPath)
	if err != nil {
		panic(err)
	}

	csvPath := flag.Arg(1)
	csvData, err := loadCsv(csvPath)
	if err != nil {
		panic(err)
	}

	var columnName []string
	for i := 0; i < len(csvData); i++ {
		if i == 0 {
			columnName = make([]string, len(csvData[i]))
			copy(columnName, csvData[i])
			continue
		}
		if len(csvData[i]) <= 1 {
			continue
		}

		for j := 0; j < len(csvData[i]) && j < len(columnName); j++ {
			str := csvData[i][j]
			for _, rv := range str {
				if ttfData.Index(rv) == 0 {
					fmt.Printf("%d[%s]: %c in %s\n", i, columnName[j], rv, str)
				}
			}
		}
	}
}

func loadTtf(path string) (*truetype.Font, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ttfBytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	ft, err := truetype.Parse(ttfBytes)
	if err != nil {
		return nil, err
	}

	return ft, nil
}

const BUFSIZE = 1024

func loadCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("csv file open: %w", err)
	}
	defer file.Close()

	var csvStr string
	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("csv file read: %w", err)
		}

		csvStr += string(buf[:n])
	}

	r := csv.NewReader(strings.NewReader(csvStr))
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv string read: %w", err)
	}

	return records, nil
}
