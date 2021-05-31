package mschartgen

import (
	//	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func convert(org Organisation) {
	log.Debug("organisation pre-convert:", org)

	// start with the head of the org
	orgChart := OrgChart{}
	child := Child{}
	child.Name = org.Head.Name
	child.Title = org.Head.Title
	orgChart.Data = append(orgChart.Data, child)

	// then move on to the root's members and beyond
	//log.Debugf("CHILDREN OF ROOT:", orgChart.Data[0].Children)
	//log.Debugf("CHILDREN OF ROOT:", org.Head.DirectReports)
	for k, v := range org.Head.DirectReports {
		log.Debug(k, v)
		//orgChart.Data[0].Children[0].Name = "foo"

		// for some reason the head might report to themselves, sometimes
		if v.Name != org.Head.Name {
			orgChart.Data[0].Children = append(orgChart.Data[0].Children, Child{Name:v.Name, Title: v.Title})
		}
	}
	//orgChart.Data[0].Children[0].Name = "foo"

	// iterate over members
	// for k, v := range org.Members {
	// 	//fmt.Printf("H -> %s\n", k, v)
	// 	fmt.Printf("KEY: %s VALUE: %s", k, v)
	// }

	// log the org chat before conversion to json for saving
	log.Debugf("OrgChart:", orgChart)

	// render it to a json file for orgchartjs
	json, err := json.MarshalIndent(orgChart.Data[0], "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./srv/data.json", json, 0644)
	if err != nil {
		panic(err)
	}
}
