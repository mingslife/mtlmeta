package main

import (
	"fmt"
	"log"
	"path"

	"github.com/mingslife/mtlmeta/conf"
	"github.com/mingslife/mtlmeta/middleware"
	"github.com/mingslife/mtlmeta/mtl"
)

func main() {
	c := conf.ParseConfig()

	if len(c.Paths) == 0 {
		log.Fatalln("empty path")
	}

	// register middlewares
	middleware.RegisterMiddlewares(c)

	for _, mtlPath := range c.Paths {
		fmt.Printf("MTL file path: %s\n", mtlPath)

		mtlFile := mtl.New(mtlPath)

		if c.Debug {
			fmt.Println("==========CONTENT==========")
			fmt.Println(mtlFile.String())
			fmt.Println("==========CONTENT==========")
		}

		materialName, err := mtlFile.MaterialName()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Material name: %s\n", materialName)

		// handle
		middleware.Handle(mtlFile)

		if c.Name != "" {
			mtlFile.SetMaterialName(c.Name)
			fmt.Printf("New material name: %s\n", c.Name)
		}

		if c.Icon != "" {
			mtlFile.SaveIcon(c.Icon)
			fmt.Printf("Icon path: %s\n", c.Icon)
		}

		if c.Directory != "" {
			outputPath := path.Join(c.Directory, mtlPath)
			mtlFile.Save(outputPath)
			fmt.Printf("New MTL path: %s\n", outputPath)
		}
	}
}
