package mtl

import (
	"errors"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type MTLFile struct {
	document *MTLDocument
}

func New(path string) *MTLFile {
	mtlFile := &MTLFile{
		document: NewMTLDocument(),
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	for pos, length := 0, len(bytes); pos < length; {
		b := bytes[pos]
		switch b {
		case '/':
			temp := []byte{}
			for ; pos < length && bytes[pos] != '\n'; pos++ {
				temp = append(temp, bytes[pos])
			}

			if string(temp[0:2]) == "//" {
				mtlFile.document.AppendChild(NewMTLComment(string(temp[2:])))
			} else {
				log.Fatalln("parse error")
			}
		case 'i':
			header := []byte{}
			for end := pos + 5; pos < end; pos++ {
				header = append(header, bytes[pos])
			}

			if string(header) != "icon " {
				log.Fatalln("parse error")
			}

			mark := bytes[pos : pos+4]
			pos += 4
			switch {
			case reflect.DeepEqual(mark, MTLIconMarkBinary):
				// binary
				size := int(bytes[pos+1])<<8 + int(bytes[pos])
				meta := bytes[pos : pos+5]
				pos += 5
				data := bytes[pos : pos+size]
				pos += size
				mtlFile.document.AppendChild(NewMTLIcon(MTLIconKindBinary, meta, data))
			case reflect.DeepEqual(mark, MTLIconMarkString):
				// string
				data := []byte{}
				for ; pos < length && bytes[pos] != '\n'; pos++ {
					data = append(data, bytes[pos])
				}
				mtlFile.document.AppendChild(NewMTLIcon(MTLIconKindString, nil, data))
			default:
				log.Fatalln("parse error")
			}
		case '#':
			temp := []byte{}
			for ; pos < length && bytes[pos] != '{'; pos++ {
				temp = append(temp, bytes[pos])
			}
			header := string(temp)

			temp = []byte{}
			for ; pos < length && bytes[pos-1] != '}'; pos++ {
				temp = append(temp, bytes[pos])
			}
			body := string(temp)

			words := strings.Split(string(header), " ")
			if words[0] == "#define" {
				switch words[1] {
				case "shader":
					index, _ := strconv.Atoi(words[2])
					name := words[3]
					mtlFile.document.AppendChild(NewMTLShader(index, name, body))
				case "material":
					name := header[strings.Index(header, "\"")+1 : strings.LastIndex(header, "\"")]
					mtlFile.document.AppendChild(NewMTLMaterial(name, body))
				}
			} else {
				log.Fatalln("parse error")
			}
		case '\n':
			mtlFile.document.AppendChild(NewMTLDelimiter(false))
			pos++
		case '\r':
			if pos+1 < length && bytes[pos+1] == '\n' {
				mtlFile.document.AppendChild(NewMTLDelimiter(true))
				pos += 2
			} else {
				log.Fatalln("parse error")
			}
		default:
			log.Fatalf("parse error in position: %d\n", pos)
		}
	}

	return mtlFile
}

func (f *MTLFile) String() string {
	return f.document.String()
}

func (f *MTLFile) Bytes() []byte {
	return f.document.Bytes()
}

func (f *MTLFile) Save(path string) {
	ioutil.WriteFile(path, f.document.Bytes(), 0644)
}

func (f *MTLFile) GetMaterial() *MTLMaterial {
	children := f.document.Children()
	for _, child := range children {
		if material, ok := child.(*MTLMaterial); ok {
			return material
		}
	}
	return nil
}

func (f *MTLFile) MaterialName() (string, error) {
	material := f.GetMaterial()
	if material != nil {
		return material.Name(), nil
	}
	return "", errors.New("material not found")
}

func (f *MTLFile) SetMaterialName(materialName string) error {
	material := f.GetMaterial()
	if material != nil {
		material.SetName(materialName)
		return nil
	}
	return errors.New("material not found")
}

func (f *MTLFile) GetIcon() *MTLIcon {
	children := f.document.Children()
	for _, child := range children {
		if icon, ok := child.(*MTLIcon); ok {
			return icon
		}
	}
	return nil
}

func (f *MTLFile) SaveIcon(path string) error {
	icon := f.GetIcon()
	if icon != nil {
		icon.Save(path)
	}
	return errors.New("icon not found")
}
