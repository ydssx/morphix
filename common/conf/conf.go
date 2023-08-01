package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func MustLoad(out *Bootstrap, path string) func() error {
	c := config.New(config.WithSource(file.NewSource(path)))

	if err := c.Load(); err != nil {
		panic(err)
	}

	if err := c.Scan(out); err != nil {
		panic(err)
	}
	return c.Close
}
