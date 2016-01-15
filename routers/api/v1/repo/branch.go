// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	//TODO change for merge
	api "github.com/gogits/go-gogs-client"

	"github.com/gogits/gogs/modules/middleware"
	"github.com/gogits/gogs/routers/api/v1/convert"
)

//TODO add to github.com/gogits/go-gogs-client
// https://github.com/gogits/go-gogs-client/wiki/Repositories#get-branch
func GetBranch(ctx *middleware.Context) {
	branch, err := ctx.Repo.Repository.GetBranch(ctx.Params(":id"))
	if err != nil {
		//TODO handle error
		return
	}
	c, err := branch.GetCommit()
	if err != nil {
		//TODO handle error
		return
	}
	ctx.JSON(200, convert.ToApiBranch(branch,c))
}

//TODO add to github.com/gogits/go-gogs-client
// https://github.com/gogits/go-gogs-client/wiki/Repositories#list-branches
func ListBranches(ctx *middleware.Context) {


	Branches, err := ctx.Repo.Repository.GetBranches()
	if err != nil {
		//TODO handle error
		return
	}
	apiBranches := make([]*api.Branch, len(Branches))
	for i := range Branches {
		c, err := Branches[i].GetCommit()
		if err != nil {
			//TODO handle error
			continue
		}
		apiBranches[i] = convert.ToApiBranch(Branches[i],c)
	}

	ctx.JSON(200, &apiBranches)
}
