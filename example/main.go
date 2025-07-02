// main.go
package main

import (
	"context"
	"fmt"
	"io"
	"mytemplate"
	"os"

	"github.com/a-h/templ"
)

func main() {
	err := mytemplate.GenerateJS(withoutParameters(), withParameters("", "", 0))

	if err != nil {
		fmt.Println("panic 1")
		panic(err)
	}

	var MyHome mytemplate.Component

	str, err := MyHome.RenderHtmlToString(nil, Home())
	if err != nil {
		fmt.Println("panic 2")
		panic(err)
	}

	str = MyHome.RemoveScriptTags(str)

	MyHome.ComponentFunc = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, str)
		return err
	})

	err = MyHome.ComponentFunc.Render(context.Background(), os.Stdout)
	if err != nil {
		fmt.Println("panic 3")
		panic(err)
	}

}
