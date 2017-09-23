// Copyright (c) 2017
// Author: Jeff Weisberg <jaw @ tcp4me.com>
// Created: 2017-Sep-06 00:21 (EDT)
// Function:

package group

import (
	"argus/api"
	"argus/configure"
	"argus/diag"
	"argus/monel"
)

type Group struct {
	mon   *monel.M
	GName string
}

var dl = diag.Logger("group")

// construction starts here:
func New(conf *configure.CF, parent *monel.M) (*monel.M, error) {

	g := &Group{}

	g.mon = monel.New(g, parent)

	err := g.mon.Config(conf)
	if err != nil {
		return nil, err
	}

	return g.mon, nil
}

func (g *Group) Config(conf *configure.CF) error {

	//conf.InitFromConfig(&g.cf, "group", "")

	g.GName = conf.Name
	g.mon.SetNames(g.GName, g.GName, g.GName)

	return nil
}

func (g *Group) Init() error {

	return nil
}

func (g *Group) DoneConfig() {

}

// destruction
func (g *Group) Recycle() {

}

func (g *Group) Persist(pm map[string]interface{}) {

}
func (g *Group) Restore(pm map[string]interface{}) {

}
func (g *Group) WebJson(md map[string]interface{}) {
}
func (g *Group) Children() []*monel.M {

	return g.mon.Children
}
func (g *Group) Dump(ctx *api.Context) {
}
func (g *Group) CheckNow() {
}
