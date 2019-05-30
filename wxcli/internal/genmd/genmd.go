package genmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jpillora/md-tmpl/mdtmpl"
	"github.com/wxio/wxcli"
)

type genmd struct {
	Filename   string `wxcli:"mode=arg"`
	WorkingDir string
	Preview    bool
}

func Register() wxcli.SubWXCli {
	gen := genmd{
		WorkingDir: ".",
	}
	return wxcli.NewSub(&gen).SubName("gen-markdown")
}

func (gen *genmd) Run() error {
	fp := filepath.Join(gen.WorkingDir, "README.md")
	if b, err := ioutil.ReadFile(gen.Filename); err != nil {
		return err
	} else {
		if gen.Preview {
			for i, c := range mdtmpl.Commands(b) {
				fmt.Printf("%18s#%d %s\n", gen.Filename, i+1, c)
			}
			return nil
		}
		b = mdtmpl.ExecuteIn(b, gen.WorkingDir)
		if err := ioutil.WriteFile(fp, b, 0655); err != nil {
			return err
		}
		log.Printf("executed templates and rewrote '%s'", gen.Filename)
		return nil
	}
}
