package impl

import (
	"bufio"
	slaveTemplate "github.com/dirkarnez/my-slave/template"
	"github.com/dirkarnez/my-slave/util"
	"log"
	"os"
	"regexp"
	"text/template"
)

type Field struct {
	Name, Value string
}

func ApiParamToCsField(inPath, outPath string) error {
	re := regexp.MustCompile("([^=&?]+)=([^&]*)&?")
	content, err := util.ReadString(inPath)

	if err != nil {
		return err
	}

	arr := re.FindAllStringSubmatch(*content, -1)

	t := template.Must(template.New("csharpfield").Parse(slaveTemplate.CSharpField))

	return util.CreateFile(outPath, func(file *os.File) error {
		w := bufio.NewWriter(file)

		for _, r := range arr {
			field := Field{Name: r[1], Value: r[2]}
			err := t.Execute(w, field)
			if err != nil {
				log.Println("executing template:", err)
			}
		}

		return w.Flush()
	})
}