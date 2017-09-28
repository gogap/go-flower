package main

import (
	"github.com/emirpasic/gods/maps/treemap"
	"strings"
)

type Args struct {
	ctx *treemap.Map
}

func NewArgs(strArgs ...string) *Args {

	args := &Args{
		ctx: treemap.NewWithStringComparator(),
	}

	for i := 0; i < len(strArgs); i++ {
		arg := strings.Split(strArgs[i], ":")

		if len(arg) == 1 {
			args.ctx.Put(arg[0], true)
		} else if len(arg) == 2 {
			args.ctx.Put(arg[0], arg[1])
		} else if len(arg) > 2 {
			args.ctx.Put(arg[0], strings.Join(arg[1:], ":"))
		}
	}

	return args
}

func (p *Args) Get(key string) (value interface{}, exist bool) {
	return p.ctx.Get(key)
}

func (p *Args) Exist(key string) bool {
	_, exist := p.ctx.Get(key)
	return exist
}

func (p Args) Keys() []string {
	var retval []string

	for _, key := range p.ctx.Keys() {
		retval = append(retval, key.(string))
	}

	return retval
}

func (p Args) GetAll() map[string]interface{} {
	var retval = make(map[string]interface{})
	p.ctx.All(func(k, v interface{}) bool {
		retval[k.(string)] = v
		return true
	})

	return retval
}
