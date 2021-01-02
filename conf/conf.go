package conf

import (
	"flag"
)

type Config struct {
	Paths     []string
	Name      string
	Icon      string
	Directory string
	Debug     bool
}

func ParseConfig() *Config {
	c := &Config{}

	// flag.StringVar(&c.Path, "p", "", "mtl file path")
	flag.StringVar(&c.Name, "n", "", "new material name")
	// flag.StringVar(&c.Icon, "i", "", "output icon path")
	flag.StringVar(&c.Directory, "d", "", "output directory")
	flag.BoolVar(&c.Debug, "D", false, "debug level command output")
	flag.Parse()

	c.Paths = flag.Args()

	return c
}
