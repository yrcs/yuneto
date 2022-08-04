package service

import (
	"context"

	v1 "github.com/yrcs/yuneto/api/hospital/v1"
	"github.com/yrcs/yuneto/app/hospital/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
