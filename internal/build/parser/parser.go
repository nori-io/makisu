package parser

import "github.com/nori-io/nori-common/meta"

func ExtractMeta(file string) meta.Data {
	return meta.Data{
		ID:           meta.ID{},
		Author:       meta.Author{},
		Dependencies: nil,
		Description:  meta.Description{},
		Core:         meta.Core{},
		Interface:    "",
		License:      nil,
		Links:        nil,
		Repository:   meta.Repository{},
		Tags:         nil,
	}
}
