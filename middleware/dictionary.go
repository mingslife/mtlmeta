package middleware

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mingslife/mtlmeta/conf"
	"github.com/mingslife/mtlmeta/mtl"
)

const (
	dictionaryPath = "dictionary.csv"
	defaultLayout  = "%T %O"
)

type DictionaryMiddleware struct {
	dictionary map[string]string
	layout     string
}

func NewDictionaryMiddleware(c *conf.Config) (m *DictionaryMiddleware) {
	m = &DictionaryMiddleware{}

	mark := c.Name
	if strings.HasPrefix(mark, "<dictionary") && strings.HasSuffix(mark, ">") {
		c.Name = ""
		fmt.Println("Enable dictionary middleware")

		m.layout = defaultLayout
		if parts := strings.Split(mark[1:len(mark)-1], ":"); len(parts) >= 2 {
			m.layout = parts[1]
		}

		fs, err := os.Open(dictionaryPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fs.Close()

		m.dictionary = map[string]string{}

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
			m.dictionary[original] = translation
		}
	}

	return
}

func (m *DictionaryMiddleware) Handle(mtlFile *mtl.MTLFile) bool {
	if m.dictionary == nil {
		return false
	}

	materialName, err := mtlFile.MaterialName()
	if err != nil {
		log.Fatalln(err)
	}
	if translation, ok := m.dictionary[materialName]; ok {
		newMaterialName := strings.ReplaceAll(m.layout, "%T", translation)
		newMaterialName = strings.ReplaceAll(newMaterialName, "%O", materialName)
		mtlFile.SetMaterialName(newMaterialName)
		fmt.Printf("New material name: %s\n", newMaterialName)
	} else {
		fmt.Printf("Not found translation for: %s\n", materialName)
	}

	return false
}

var _ MTLMiddleware = (*DictionaryMiddleware)(nil)
