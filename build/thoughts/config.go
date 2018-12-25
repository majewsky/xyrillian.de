/*******************************************************************************
*
* Copyright 2016-2018 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under the
* terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

import (
	"io/ioutil"
)

const (
	//SourceURL is a constant.
	SourceURL = "https://github.com/majewsky/xyrillian.de"
	//TargetURL is a constant.
	TargetURL = "https://xyrillian.de"
	//PageName is a constant.
	PageName = "Xyrillian Thoughts"
	//PageDescription is a constant.
	PageDescription = "Personal blog of Stefan Majewsky"
)

//Configuration contains the configuration of the program.
type Configuration struct {
	TemplateHTML string
}

//Config contains the configuration of the program.
var Config Configuration

func init() {
	//read template.html
	bytes, err := ioutil.ReadFile("build/thoughts/template.html")
	FailOnErr(err)
	Config.TemplateHTML = string(bytes)
}
