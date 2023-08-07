package repos

import (
	"io"

	"github.com/Valentin-Kaiser/go-dbase/dbase"
	"github.com/Volkov-D-A/gku-data-analisys/pkg/models"
	"golang.org/x/text/encoding/charmap"
)

func ImportDbfData(filepath string, f io.Writer) (map[string][]models.GkuData, error) {
	dbase.Debug(true, io.Writer(f))
	conv := dbase.NewDefaultConverter(charmap.CodePage866)

	table, err := dbase.OpenTable(&dbase.Config{
		Filename:   filepath,
		Converter:  conv,
		Untested:   true,
		TrimSpaces: true,
	})

	if err != nil {
		return nil, dbase.GetErrorTrace(err)
	}

	defer table.Close()

	data := make(map[string][]models.GkuData, 0)

	for !table.EOF() {
		row, err := table.Next()
		if err != nil {
			return nil, dbase.GetErrorTrace(err)
		}

		gk := &models.GkuData{}
		err = row.ToStruct(gk)
		if err != nil {
			return nil, dbase.GetErrorTrace(err)
		}
		data[gk.Lschet] = append(data[gk.Lschet], *gk)
	}

	return data, nil

}
