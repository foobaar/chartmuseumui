package controllers

import (
	"os"
	"os/exec"
	"quickstart/models"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

func getCharts() map[string][]models.Chart {

	l := logs.GetLogger()
	res, err := httplib.Get(os.Getenv("CHART_MUSESUM_URL") + api_get_charts).Debug(true).Bytes()
	if err != nil {
		l.Panic(err.Error)
	}

	charts, err := models.NewCharts(res)
	if err != nil {
		l.Panic(err)
	}
	return charts
}

func uploadChart(filePath string) {

	l := logs.GetLogger()

	cmd := exec.Command("curl", "-L", "--data-binary", "@"+filePath, os.Getenv("CHART_MUSESUM_URL")+api_get_charts)
	out, err := cmd.CombinedOutput()
	if err != nil {
		l.Fatalf("cmd.Run() failed with %s\n", err)
	}
	l.Printf("combined out:\n%s\n", string(out))
}

func deleteChart(name string, version string) {

	l := logs.GetLogger()
	l.Println("in deleteChart()")
	cmd := exec.Command("curl", "-X", "DELETE", os.Getenv("CHART_MUSESUM_URL")+api_get_charts+"/"+name+"/"+version)
	out, err := cmd.CombinedOutput()
	if err != nil {
		l.Fatalf("cmd.Run() failed with %s\n", err)
	}
	l.Printf("combined out:\n%s\n", string(out))
}
