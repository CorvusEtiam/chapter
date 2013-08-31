package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const PATH = "./game.xml"

type Game struct {
	Place []Place `xml:"Place"`
}

type Place struct {
	Name    string `xml:"name,attr"`
  Actions []Action `xml:"Action"`
}

type Action struct {
	Type    string `xml:"type,attr"`
	Keyword string `xml:"keyword,attr"`
  Has     string `xml:"has,attr"`
  Goto    string `xml:"go-to, attr"`
  Text    string `xml:",innerxml"`
  Item    []Items `xml:"Item"`
}

type Items struct {
  Type string `xml:"type,attr"`
  Num string `xml:"num,attr"`
  Item string `xml:",innerxml"`
}

func main() {
	var g Game
  g = getData(PATH)
  for _, act := range g.Place {
		for _, v := range act.Actions {
			fmt.Println("->", v.Text)
		}
	}
}

func getData(path string) Game  {
  var x Game
  xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("*DEBUG_ME*",err)
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(b, &x)
  return x
}

