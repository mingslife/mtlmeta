package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Config struct {
	Paths      []string
	Dictionary string
	Collect    bool
	Rename     bool
	Output     string
}

func parseConfig() *Config {
	c := &Config{}

	flag.StringVar(&c.Dictionary, "d", "dictionary.csv", "dictionary path")
	flag.BoolVar(&c.Collect, "c", false, "collect mode")
	flag.BoolVar(&c.Rename, "r", false, "rename mode")
	flag.StringVar(&c.Output, "o", "", "output file path")
	flag.Parse()
	c.Paths = flag.Args()

	return c
}

func getFilePaths(filePath, ext string) []string {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		panic(err)
	}
	filePaths := []string{}
	for _, file := range files {
		filePath := path.Join(filePath, file.Name())
		if file.IsDir() {
			filePaths = append(filePaths, getFilePaths(filePath, ext)...)
		} else if strings.ToLower(path.Ext(filePath)) == ext {
			filePaths = append(filePaths, filePath)
		}
	}
	return filePaths
}

func loadFilePaths(paths []string) []string {
	filePaths := []string{}

	for _, filePath := range paths {
		if Exists(filePath) {
			if IsDir(filePath) {
				filePaths = append(filePaths, getFilePaths(filePath, ".mtl")...)
			} else if strings.ToLower(path.Ext(filePath)) == ".mtl" {
				filePaths = append(filePaths, filePath)
			}
		}
	}

	return filePaths
}

func loadDictionary(dictionaryPath string) map[string]string {
	fs, err := os.Open(dictionaryPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer fs.Close()

	dictionary := map[string]string{}

	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		if err == io.EOF {
			break
		}
		original, translation := row[0], row[1]
		dictionary[original] = translation
	}

	return dictionary
}

func main() {
	c := parseConfig()

	dictionary := loadDictionary(c.Dictionary)
	tempData := []byte{}

	filePaths := loadFilePaths(c.Paths)

	for _, filePath := range filePaths {
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalln(err)
		}

		start, end, length, hasQuote := -1, -1, len(bytes), true
		for i, b := range bytes {
			if b == '#' {
				temp := []byte{}
				for j := i; j < length && bytes[j] != '\n'; j++ {
					temp = append(temp, bytes[j])
				}
				header := string(temp)
				if strings.HasPrefix(header, "#define ") {
					parts := strings.Split(strings.Trim(header[8:], " "), " ")
					if parts[0] == "material" {
						if strings.Index(header, "\"") != -1 {
							start = i + strings.Index(header, "\"") + 1
							end = i + strings.LastIndex(header, "\"")
						} else {
							hasQuote = false
							count := 0
							blank := false
							for j := 0; j < len(header); j++ {
								if header[j] == ' ' {
									if !blank {
										count++
										blank = true
									}
								} else {
									blank = false
								}
								if blank {
									switch count {
									case 2:
										start = i + j + 1
									case 3:
										end = i + j
										break
									}
								}
							}
						}
						break
					} else {
						continue
					}
				}
			}
		}

		materialName := string(bytes[start:end])
		fmt.Println("Material name: " + materialName)

		if c.Collect {
			fmt.Println("collect for: " + filePath)

			translation, original := "", ""
			for i, b := range materialName {
				if b == ' ' {
					translation, original = materialName[0:i], materialName[i+1:]
					break
				}
			}

			if original != "" && translation != "" {
				// handle Chinese in original
				words := strings.Split(original, " ")
				for len(words) > 1 {
					word := words[0]
					for _, c := range word {
						if !(c >= 0x21 && c <= 0x7e) {
							translation += " " + word
							words = words[1:]
							break
						}
					}
					break
				}
				original = strings.Join(words, " ")

				row := fmt.Sprintf("%s,%s,%s\n", original, translation, materialName)

				fmt.Print("collect sucessfully: " + row)
				tempData = append(tempData, []byte(row)...)
			} else if materialName != "" {
				row := fmt.Sprintf("%s,%s,%s\n", materialName, "", materialName)

				fmt.Print("collect empty record: " + row)
				tempData = append(tempData, []byte(row)...)
			} else {
				fmt.Println("collect failed for file: " + filePath)
			}
		} else if c.Rename {
			fmt.Println("rename for: " + filePath)

			// get output file path
			outputFilePath := c.Output
			if len(filePaths) == 1 {
				if outputFilePath == "" {
					_, fileName := path.Split(filePath)
					outputDir := "out"
					os.Mkdir(outputDir, os.ModePerm)
					outputFilePath = path.Join(outputDir, fileName)
				}
			} else {
				fileName := strings.ReplaceAll(path.Clean(filePath), "..", "")
				outputFilePath = path.Join(outputFilePath, fileName)
				outputDir := path.Dir(outputFilePath)
				os.MkdirAll(outputDir, os.ModePerm)
			}

			if translation, ok := dictionary[materialName]; ok {
				newMaterialName := materialName
				if translation != "" {
					newMaterialName = fmt.Sprintf("%s %s", translation, materialName)
				}

				newBytes := bytes[0:start:start]
				if hasQuote {
					newBytes = append(newBytes, []byte(newMaterialName)...)
				} else {
					newBytes = append(newBytes, []byte("\""+newMaterialName+"\"")...)
				}
				newBytes = append(newBytes, bytes[end:]...)

				ioutil.WriteFile(outputFilePath, newBytes, os.ModePerm)

				fmt.Printf("%s -> %s\n", materialName, newMaterialName)
			} else {
				ioutil.WriteFile(outputFilePath, bytes, os.ModePerm)

				fmt.Println("not found translation for: " + materialName)
			}
		}
	}

	if c.Collect {
		f, err := os.OpenFile(c.Dictionary, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}

		f.Write(tempData)
	}
}
