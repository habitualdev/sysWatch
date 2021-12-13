package main

import(
	g "github.com/AllenDang/giu"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var wnd = g.NewMasterWindow("sysWatch", 800, 600, 0)
var tabBarLayout g.Layout


var stringsTableEnabled string
var displayStringEnabled string

var displayStringDisabled string
var stringsTableDisabled string

var displayStringRunning string
var stringsTableRunning string




func rootCheck() {
	cmd := exec.Command("id", "-u")
	output, err := cmd.Output()

	rootInt, _ := strconv.Atoi(strings.ReplaceAll(string(output),"\n",""))

	if rootInt == 0 {
		tabBarLayout= g.Layout{
			g.TabBar().TabItems(g.TabItem("Enabled").Layout(g.Label(displayStringEnabled)),g.TabItem("Disabled").Layout(g.Label(displayStringDisabled)),g.TabItem("Running").Layout(g.Label(displayStringRunning))),
		}
	} else {
		tabBarLayout = g.Layout{
			g.Label("You must run this application as root"),
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}


func loop() {
	rootCheck()
	w1 := g.SingleWindow()
	w1.Layout(tabBarLayout)

}

func getSysEnabled() {
	for {
		displayStringEnabled = ""
		sysGet := exec.Command("systemctl", "list-unit-files", "--state=enabled")
		stringsTableEnabledB, _ := sysGet.Output()
		stringsTableEnabled = string(stringsTableEnabledB)

		lineSplit := strings.Split(stringsTableEnabled, "\n")
		for _, line := range lineSplit {
			if len(strings.Fields(line)) == 3 {
				displayStringEnabled = displayStringEnabled + strings.Fields(line)[0] + "\n"
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func getSysDisabled() {
	for {
		displayStringDisabled = ""
		sysGet := exec.Command("systemctl", "list-unit-files", "--state=disabled")
		stringsTableDisabledB, _ := sysGet.Output()
		stringsTableDisabled = string(stringsTableDisabledB)

		lineSplit := strings.Split(stringsTableDisabled, "\n")
		for _, line := range lineSplit {
			if len(strings.Fields(line)) == 3 {
				displayStringDisabled = displayStringDisabled + strings.Fields(line)[0] + "\n"
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func getSysRunning() {
	for {
		displayStringRunning = ""
	sysGet := exec.Command("systemctl", "list-units", "--type=service", "--state=active")
	stringsTableRunningB, _ := sysGet.Output()
	stringsTableRunning = string(stringsTableRunningB)

	lineSplit := strings.Split(stringsTableRunning, "\n")
	for _, line := range lineSplit {
		if len(strings.Fields(line)) >= 5 {
			if strings.Fields(line)[0] != strings.ToUpper(strings.Fields(line)[0]) {

				displayStringRunning = displayStringRunning + strings.Fields(line)[0] + "\n"
			}
		}
	}
	time.Sleep(time.Second * 10)
}
}


func main() {
	go getSysEnabled()
	go getSysDisabled()
	go getSysRunning()

wnd.Run(loop)
}