package csv

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

type CsvReader struct {
	rowVisitor RowVisitor
}

func NewReader(v RowVisitor) *CsvReader {
	return &CsvReader{
		rowVisitor: v,
	}
}

func (c *CsvReader) VisitRows(filename string) {
	readFile, err :=  os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
 	for fileScanner.Scan() {
		split := strings.Fields(fileScanner.Text())
 		c.rowVisitor.VisitRow(split)
	}
}

type RowVisitor interface {
	VisitRow(row []string)
}

type LowestValueTracker struct {
	minValue int
	label string
}

func (l *LowestValueTracker) UpdateMinVal(n int, nLabel string) {
	if n < l.minValue {
 		l.minValue = n
 		l.label = nLabel
	}
}

type FootballSpreadRowVisitor struct {
	*LowestValueTracker
}

func NewFootballSpreadRowVisitor() *FootballSpreadRowVisitor {
	lvt := &LowestValueTracker{
		minValue: math.MaxInt64,
	}
	return &FootballSpreadRowVisitor{
		 lvt,
	}
}

func (f *FootballSpreadRowVisitor) VisitRow(row []string) {
	if len(row) != 10 {
		return
	}
	forG, err := strconv.Atoi(row[6])
	if err != nil {
		return
	}
	against, err := strconv.Atoi(row[8])
	if err != nil {
		return
	}
	diff := int(math.Abs(float64(forG-against)))
	f.UpdateMinVal(diff, row[1])
}

func TeamWithMinForAgainstSpread(filename string) string {
	visitor := NewFootballSpreadRowVisitor()
	csvReader := &CsvReader{
		rowVisitor: visitor,
	}
	csvReader.VisitRows(filename)
	return visitor.label
}

type TempSpreadRowVisitor struct {
	*LowestValueTracker
}

func NewTempSpreadRowVisitor() *TempSpreadRowVisitor {
	lvt := &LowestValueTracker{
		minValue: math.MaxInt64,
	}
	return &TempSpreadRowVisitor{
		lvt,
	}
}

func (t *TempSpreadRowVisitor) VisitRow(row []string) {
	if len(row) == 0 {
		return
	}
	if row[0] == "mo" {
		return
	}
	minTemp, err := readIntFieldRemovingAsterisk(row[2])
	if err != nil {
		return
	}
	maxTemp, err := readIntFieldRemovingAsterisk(row[1])
	if err != nil {
		return
	}
	t.UpdateMinVal(maxTemp-minTemp, row[0])
}

func DayOfSmallestTempSpread(filename string) string {
	visitor := NewTempSpreadRowVisitor()
	csvReader := &CsvReader{
		rowVisitor: visitor,
	}
	csvReader.VisitRows(filename)
	return visitor.label
}

func readIntFieldRemovingAsterisk(i string) (int, error) {
	if i[len(i)-1] == '*' {
		i = i[:len(i)-1]
	}
	return strconv.Atoi(i)
}

func main() {
	weatherSpread := DayOfSmallestTempSpread("weather.dat")
	faSpread := TeamWithMinForAgainstSpread("football.dat")
	fmt.Println(faSpread, weatherSpread)
}

