package app

import (

	"github.com/spf13/pflag"
)
type NamedFlagSet struct {
	Order [] string
	FlagSets map[string] *pflag.FlagSet
}
type CliOptions interface {
	Flags()(fs NamedFlagSet)
	Validate()[]error
}
func (nfs *NamedFlagSet)FlagSet(name string)*pflag.FlagSet{
	if nfs.FlagSets == nil{
		nfs.FlagSets = map[string]*pflag.FlagSet{}
	}
	if _,ok:= nfs.FlagSets[name];!ok{
		nfs.FlagSets[name] = pflag.NewFlagSet(name,pflag.ExitOnError)
		nfs.Order = append(nfs.Order,name)
	}
	return nfs.FlagSets[name]
}