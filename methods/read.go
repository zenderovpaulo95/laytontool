package methods

import(
    "os"
    "fmt"
    "bufio"
    "io/fs"
    "io/ioutil"
    "errors"
    "github.com/xuri/excelize/v2"
)

func GetFiles(dirPath string) (files []fs.FileInfo, err error) {
    files, err = ioutil.ReadDir(dirPath)
    if err != nil {
     return
    }

    return files, nil
}

func AppendStrings(strings []string, appendStrings []string) []string {
    for i := 0; i < len(appendStrings); i++ {
      strings = append(strings, appendStrings[i])
    }

  return strings
}

func RemoveString(strings []string, index int) ([]string, error) {
  if index >= len(strings) {
     return nil, errors.New("Out of Range error")
  }

  return append(strings[:index], strings[index + 1:]...), nil
}

func ReadReplaceChars(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadTxt(filePath string) []string {
    file, err := os.Open(filePath)

    if err != nil {
        return nil
    }
    defer file.Close()

    var lines []string
    buf := bufio.NewScanner(file)

    for buf.Scan() {
        lines = append(lines, buf.Text())
    }

    if lines[len(lines) - 1] == "" {
        lines, _ = RemoveString(lines, len(lines) - 1)
    }

    return lines
}

func ReadXlsx(filePath string, index int) []string {
    f, err := excelize.OpenFile(filePath)

    if err != nil {
	    return nil
    }

    defer func(){
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
    }()

    cols, _ := f.GetCols("Sheet1")


    if len(cols) == 2 && index < 2 {
	    rows, _ := f.GetRows("Sheet1")
	    var strs []string

	    for _, row := range rows {
		strs = append(strs, row[index])
	    }

	    return strs
    } else {
	fmt.Println("В файле должно быть только 2 столбца: для оригинала и для перевода!")
    }

    return nil
}
