package service

import (
	"context"

	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"github.com/yrcs/yuneto/third_party/pagination"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *HospitalService) ListHospitalSetting(ctx context.Context, req *pagination.PagingRequest) (*pagination.PagingReply, error) {
	reply, err := s.hsu.List(ctx, req)
	if err != nil {
		return nil, err
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
		Total: uint32(len(reply)),
		Items: items,
	}, nil
}

func (s *HospitalService) AddHospitalSetting(ctx context.Context, req *v1.AddHospitalSettingRequest) (*v1.CommonAddReply, error) {
	reply, err := s.hsu.Add(ctx, &biz.HospitalSetting{
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
	reply, err := s.hsu.Edit(ctx, &biz.HospitalSetting{
		Id:                 req.GetId(),
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
	return &v1.CommonEditReply{
		Id:        reply.Id,
		UpdatedAt: timestamppb.New(reply.UpdatedAt),
	}, nil
}

func (s *HospitalService) DeleteHospitalSetting(ctx context.Context, req *v1.DeleteHospitalSettingRequest) (*emptypb.Empty, error) {
	err := s.hsu.Delete(ctx, &biz.HospitalSetting{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *HospitalService) DeleteHospitalSettings(ctx context.Context, req *v1.DeleteHospitalSettingsRequest) (*emptypb.Empty, error) {
	err := s.hsu.DeleteInBatches(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
