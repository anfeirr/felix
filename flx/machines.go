package flx

import (
	"fmt"
	"github.com/dejavuzhou/felix/models"
	"github.com/mattn/go-isatty"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"runtime"
	"strconv"
)

const caption = "ssh login by ID: felix ssh 2"

func AllMachines(search string) {
	data := fetchMachineToRows(search)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Addr", "User", "IP", "TYPE", "PORT"})
	table.SetBorder(true) // Set Border to false
	table.SetCaption(true, caption)
	//table.SetAutoMergeCells(true)
	//table.SetRowLine(true)

	setListTableColor(table)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func fetchMachineToRows(search string) [][]string {
	mcs, err := models.MachineAll(search)
	if err != nil {
		log.Fatal(err)
	}
	var rows [][]string
	for _, mc := range mcs {
		id := fmt.Sprintf("%d", mc.ID)
		one := []string{id, mc.Name, mc.Host, mc.User, mc.Ip, mc.Type, strconv.Itoa(int(mc.Port))}
		rows = append(rows, one)
	}
	return rows
}
func setListTableColor(table *tablewriter.Table) {
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) || (runtime.GOOS != "windows") {
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold})
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.FgRedColor},
			tablewriter.Colors{tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.FgCyanColor},
			tablewriter.Colors{tablewriter.FgCyanColor})
	}
}
