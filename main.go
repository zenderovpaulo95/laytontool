package main

import(
     "fmt"
     "os"
     "strings"
     "laytontool.com/m/v2/methods"
     "golang.org/x/text/encoding/charmap"
)

func main() {
 //fmt.Println("Hello World from go!")
 args := os.Args

 if len(args) == 3 && args[1] == "merge" {
   files, err := methods.GetFiles(args[1])

   if err != nil {
     panic(err)
   }

   var originalStrings []string
   var translateStrings []string

   for _, file := range files {
      //fmt.Println(file.Name())
       strs := methods.ReadTxt(args[1] + "/" + file.Name())

       if err != nil {
         panic(err)
       }

     if strings.Contains(file.Name(), "_EN") {
       originalStrings = methods.AppendStrings(originalStrings, strs)
     } else {
       translateStrings = methods.AppendStrings(translateStrings, strs)
     }
   }

   methods.WriteXMLX(originalStrings, translateStrings, "test.xlsx")

   fmt.Println("Done.")
 } else if len(args) >= 4 {
	 var needConvert bool = false
	 var needReplace bool = false
	 var charReplace []string

	 for c := 0; c < len(args); c++ {
		if args[c] == "convert" {
			needConvert = true
		}
		if (args[c] == "replace") && (c + 1 < len(args)) {
			needReplace = true
			var err error

			charReplace, err = methods.ReadReplaceChars(args[c + 1])

			if err != nil {
				panic(err)
			}
		}
		if args[c] == "maketxt" {
		strs := methods.ReadXlsx(args[c + 1], 1)

		var checkStrs = []string{"×", "£"}
		var replaceStrs = []string{"<multiSign>", "<poundSign>"}

		ch := []rune(charReplace[0])
		newCh := []rune(charReplace[1])

		for i := 0; i < len(strs); i = i + 2 {
			if !strings.Contains(strs[i], "tobj/charm.txt") {
				filePath := args[c + 2] + "/" + strs[i]
				methods.MakeDir(filePath)

				file, _ := os.Create(filePath)

				if strings.Contains(strs[i + 1], "\\n") {
					strs[i + 1] = strings.Replace(strs[i + 1], "\\n", "\n", -1)
				}

				if needReplace == true {
					for cr := 0; cr < len(ch); cr++ {
						strs[i + 1] = strings.Replace(strs[i + 1], string(ch[cr]), string(newCh[cr]), -1)
					}
				}

				if needConvert == true {
					for ci, ch := range checkStrs {
						if strings.Contains(strs[i + 1], ch) {
							strs[i + 1] = strings.Replace(strs[i + 1], ch, replaceStrs[ci], -1)
						}
					}

					b := []byte(strs[i + 1])
					encoder := charmap.Windows1251.NewEncoder()
					newB, e := encoder.Bytes(b)
					if e != nil {
						panic(e)
					}

					decoder := charmap.Windows1252.NewDecoder()
					b, e = decoder.Bytes(newB)
					if e != nil {
						panic(e)
					}

					strs[i + 1] = string(b)

					for ci, c := range checkStrs {
						if strings.Contains(strs[i + 1], replaceStrs[ci]) {
							strs[i + 1] = strings.Replace(strs[i + 1], replaceStrs[ci], c, -1)
						}
					}
				}

				_, _ = file.WriteString(strs[i + 1])
				file.Close()

				fmt.Println("Создан файл " + filePath)
			}
		}
		}
	 }
 } else if len(args) == 3 && args[1] == "checkls" {
	 or_strs := methods.ReadXlsx(args[2], 0)
	 tr_strs := methods.ReadXlsx(args[2], 1)

	 maxLen := 0
	 word := ""

	 for i := 0; i < len(or_strs); i++ {
		if strings.Contains(or_strs[i], "jiten") {
			if maxLen < len(or_strs[i + 1]) {
				maxLen = len(or_strs[i + 1])
				word = or_strs[i + 1]
			}
			if len(or_strs[i + 1]) < len(tr_strs[i + 1]) {
				fmt.Println(tr_strs[i] + "\t" + or_strs[i + 1] + "\t" + tr_strs[i + 1])
			}
		}
	 }

	 fmt.Println(maxLen)
	 fmt.Println(word)
 } else {
    fmt.Println("Как пользоваться этой программой:")
    fmt.Println(args[0] + " merge <Путь к папке с файлами из Нотабенойд> - Объединяет текстовые файлы в xlsx файл")
    fmt.Println(args[0] + " maketxt <Путь к файлу xlsx> <Путь к папке для будущих файлов> - Создаёт файлы по xlsx файлу")
    fmt.Println(args[0] + " checkls <Путь к файлу xlsx> - Проверить на наличие длинных фраз в названии локаций")
 }
}
