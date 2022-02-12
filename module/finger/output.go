package finger

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func outjson(filename string, data []byte) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		defer f.Close()
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func outxlsx(filename string, msg []Outrestul) {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "url")
	xlsx.SetCellValue("Sheet1", "B1", "cms")
	xlsx.SetCellValue("Sheet1", "C1", "server")
	xlsx.SetCellValue("Sheet1", "D1", "statuscode")
	xlsx.SetCellValue("Sheet1", "E1", "length")
	xlsx.SetCellValue("Sheet1", "F1", "title")
	for k, v := range msg {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(k+2), v.Url)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(k+2), v.Cms)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(k+2), v.Server)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(k+2), v.Statuscode)
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(k+2), v.Length)
		xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(k+2), v.Title)
	}
	err := xlsx.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}

func outtxt(filename string, allresult []Outrestul) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		defer f.Close()
		for _,aas := range allresult {
			data := fmt.Sprintf(fmt.Sprintf("%s, %s, %s, %d, %d, %s\n", aas.Url, aas.Cms, aas.Server, aas.Statuscode, aas.Length, aas.Title))
			if _, err = f.Write([]byte(data)); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func outfile(filename string, allresult []Outrestul) {
	file := strings.Split(filename, ".")
	fileSuffix := file[len(file)-1]
	switch fileSuffix {
	case "txt":
		outtxt(filename, allresult)
	case "json":
		buf, err := json.MarshalIndent(allresult, "", " ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		outjson(filename, buf)
	case "xlsx":
		outxlsx(filename, allresult)
	default:
		color.RGBStyleFromString("237,64,35").Printf("[error] Output file suffix error: %s\n", filename)
	}

}
