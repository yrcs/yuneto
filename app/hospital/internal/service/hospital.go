package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/app/hospital/internal/pkg/do"
	"github.com/yrcs/yuneto/pkg/util"
	"github.com/yrcs/yuneto/third_party/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *HospitalService) ListHospitalSetting(ctx context.Context, req *pagination.PagingRequest) (*pagination.PagingReply, error) {
	reply, err := s.hsu.List(ctx, req)
	if err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	items := make([]*anypb.Any, 0)
	for _, hs := range reply {
		hsr := &v1.HospitalSettingReply{
			Id:                 hs.Id,
			CreatedAt:          timestamppb.New(hs.CreatedAt),
			UpdatedAt:          timestamppb.New(hs.UpdatedAt),
			Name:               hs.Name,
			RegistrationNumber: hs.RegistrationNumber,
			ContactPerson:      hs.ContactPerson,
			ContactMobile:      hs.ContactMobile,
			Locked:             uint32(hs.Locked),
			ApiUrl:             hs.ApiUrl,
			SignatureKey:       hs.SignatureKey,
		}
		hsAny, _ := anypb.New(hsr)
		items = append(items, hsAny)
	}
	return &pagination.PagingReply{
		Total: uint32(len(items)),
		Items: items,
	}, nil
}

func (s *HospitalService) AddHospitalSetting(ctx context.Context, req *v1.AddHospitalSettingRequest) (*v1.CommonAddReply, error) {
	var (
		err                                  error
		nameExists, registrationNumberExists bool
		reply                                *do.HospitalSetting
	)

	if nameExists, err = s.hsu.NameExists(ctx, req.GetName()); err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	if registrationNumberExists, err = s.hsu.RegistrationNumberExists(ctx, req.GetRegistrationNumber()); err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	if nameExists || registrationNumberExists {
		return nil, biz.ErrHospitalSettingSameDataExists
	}

	node, _ := snowflake.NewNode(1)
	reply, err = s.hsu.Add(ctx, &do.HospitalSetting{
		Id:                 uint64(node.Generate().Int64()),
		Name:               req.GetName(),
		RegistrationNumber: req.GetRegistrationNumber(),
		ContactPerson:      req.GetContactPerson(),
		ContactMobile:      req.GetContactMobile(),
		Locked:             uint8(req.GetLocked()),
		ApiUrl:             req.GetApiUrl(),
		SignatureKey:       req.GetSignatureKey(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CommonAddReply{
		Id:        reply.Id,
		CreatedAt: timestamppb.New(reply.CreatedAt),
		UpdatedAt: timestamppb.New(reply.UpdatedAt),
	}, nil
}

func (s *HospitalService) EditHospitalSetting(ctx context.Context, req *v1.EditHospitalSettingRequest) (*v1.CommonEditReply, error) {
	m := make(map[string]any, 2)
	m["Id"] = req.GetId()
	util.UpdateOptionalField(req, m)

	var (
		err                                  error
		nameUnique, registrationNumberUnique bool
		reply                                *do.HospitalSetting
	)

	if nameUnique, err = s.hsu.NameUnique(ctx, req.GetId(), req.GetName()); err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	if registrationNumberUnique, err = s.hsu.RegistrationNumberUnique(ctx, req.GetId(), req.GetRegistrationNumber()); err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	if !nameUnique || !registrationNumberUnique {
		return nil, biz.ErrHospitalSettingSameDataExists
	}

	reply, err = s.hsu.Edit(ctx, m)
	if err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	return &v1.CommonEditReply{
		Id:        reply.Id,
		UpdatedAt: timestamppb.New(reply.UpdatedAt),
	}, nil
}

func (s *HospitalService) DeleteHospitalSetting(ctx context.Context, req *v1.DeleteHospitalSettingRequest) (*emptypb.Empty, error) {
	err := s.hsu.Delete(ctx, req.GetId())
	if err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	return nil, nil
}

func (s *HospitalService) DeleteHospitalSettings(ctx context.Context, req *v1.DeleteHospitalSettingsRequest) (*emptypb.Empty, error) {
	if err := s.hsu.DeleteByIDs(ctx, req.GetIds()); err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	}
	return nil, nil
}

func (s *HospitalService) LockHospitalSetting(ctx context.Context, req *v1.LockHospitalSettingRequest) (*v1.CommonEditReply, error) {
	m := make(map[string]any, 2)
	m["Id"] = req.GetId()
	m["Locked"] = req.GetLocked()

	reply, err := s.hsu.Edit(ctx, m)
	if err != nil {
		return nil, biz.ErrHospitalSettingSystemError.WithCause(err)
	} else if reply == nil {
		return nil, biz.ErrHospitalSettingNotFound.WithCause(err)
	}
	return &v1.CommonEditReply{
		Id:        reply.Id,
		UpdatedAt: timestamppb.New(reply.UpdatedAt),
	}, nil
}
