package methods

import (
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "github.com/xuri/excelize/v2"
)

func WriteXMLX(original []string, translation []string, outputPath string) {
    f := excelize.NewFile()
    f.SetColWidth("Sheet1", "A", "B", 75)

    var orI, trI int

    orI = 1
    trI = 1

    for i := 0; i < len(original); i++ {
       orCell := "A" + strconv.Itoa(orI)
       trCell := "B" + strconv.Itoa(trI)

       orI += 1
       trI += 1

       f.SetCellValue("Sheet1", orCell, original[i])
       f.SetCellValue("Sheet1", trCell, translation[i])
    }

    if err := f.SaveAs(outputPath); err != nil {
        fmt.Println(err)
    }
}

func MakeDir(filePath string) {
	Dir := filepath.Dir(strings.TrimSuffix(filePath, filepath.Base(filePath)))
	_, err := os.Stat(Dir)

	if os.IsNotExist(err) {
		os.MkdirAll(Dir, os.ModePerm)
	}
}
