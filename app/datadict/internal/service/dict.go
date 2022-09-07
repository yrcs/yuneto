package service

import (
	"context"

	v1 "github.com/yrcs/yuneto/api/datadict/v1"
	"github.com/yrcs/yuneto/app/datadict/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *DataDictService) ListChildren(ctx context.Context, req *v1.ListChildrenRequest) (*v1.ListChildrenReply, error) {
	reply, err := s.u.FindChildren(ctx, req.Id)
	if err != nil {
		return nil, biz.ErrDictSystemError.WithCause(err)
	}
	items := make([]*v1.ListChildrenReply_Dict, 0)
	for _, item := range reply {
		dict := &v1.ListChildrenReply_Dict{
			Id:          item.Id,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
			ParentId:    item.ParentId,
			Name:        item.Name,
			Value:       item.Value,
			DictCode:    item.DictCode,
			HasChildren: item.HasChildren,
		}
		items = append(items, dict)
	}
	return &v1.ListChildrenReply{Items: items}, nil
}
