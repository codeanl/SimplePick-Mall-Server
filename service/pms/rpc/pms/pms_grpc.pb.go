// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pms.proto

package pms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Pms_CategoryAdd_FullMethodName     = "/pms.Pms/CategoryAdd"
	Pms_CategoryList_FullMethodName    = "/pms.Pms/CategoryList"
	Pms_CategoryUpdate_FullMethodName  = "/pms.Pms/CategoryUpdate"
	Pms_CategoryDelete_FullMethodName  = "/pms.Pms/CategoryDelete"
	Pms_AttributeAdd_FullMethodName    = "/pms.Pms/AttributeAdd"
	Pms_AttributeList_FullMethodName   = "/pms.Pms/AttributeList"
	Pms_AttributeUpdate_FullMethodName = "/pms.Pms/AttributeUpdate"
	Pms_AttributeDelete_FullMethodName = "/pms.Pms/AttributeDelete"
	Pms_ProductAdd_FullMethodName      = "/pms.Pms/ProductAdd"
	Pms_ProductList_FullMethodName     = "/pms.Pms/ProductList"
	Pms_ProductUpdate_FullMethodName   = "/pms.Pms/ProductUpdate"
	Pms_ProductDelete_FullMethodName   = "/pms.Pms/ProductDelete"
	Pms_ProductInfo_FullMethodName     = "/pms.Pms/ProductInfo"
	Pms_SkuAdd_FullMethodName          = "/pms.Pms/SkuAdd"
	Pms_SkuList_FullMethodName         = "/pms.Pms/SkuList"
	Pms_SkuUpdate_FullMethodName       = "/pms.Pms/SkuUpdate"
	Pms_SkuDelete_FullMethodName       = "/pms.Pms/SkuDelete"
	Pms_SkuInfo_FullMethodName         = "/pms.Pms/SkuInfo"
)

// PmsClient is the client API for Pms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PmsClient interface {
	// 添加分类
	CategoryAdd(ctx context.Context, in *CategoryAddReq, opts ...grpc.CallOption) (*CategoryAddResp, error)
	// 分类列表
	CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
	// 更新分类
	CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*CategoryUpdateResp, error)
	// 删除分类
	CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error)
	// 添加属性
	AttributeAdd(ctx context.Context, in *AttributeAddReq, opts ...grpc.CallOption) (*AttributeAddResp, error)
	// 属性列表
	AttributeList(ctx context.Context, in *AttributeListReq, opts ...grpc.CallOption) (*AttributeListResp, error)
	// 更新属性
	AttributeUpdate(ctx context.Context, in *AttributeUpdateReq, opts ...grpc.CallOption) (*AttributeUpdateResp, error)
	// 删除属性
	AttributeDelete(ctx context.Context, in *AttributeDeleteReq, opts ...grpc.CallOption) (*AttributeDeleteResp, error)
	// 添加商品
	ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error)
	// 商品列表
	ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error)
	// 更新商品
	ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error)
	// 删除商品
	ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error)
	// 查询商品详情
	ProductInfo(ctx context.Context, in *ProductInfoReq, opts ...grpc.CallOption) (*ProductInfoResp, error)
	// 添加Sku
	SkuAdd(ctx context.Context, in *SkuAddReq, opts ...grpc.CallOption) (*SkuAddResp, error)
	// Sku列表
	SkuList(ctx context.Context, in *SkuListReq, opts ...grpc.CallOption) (*SkuListResp, error)
	// 更新Sku
	SkuUpdate(ctx context.Context, in *SkuUpdateReq, opts ...grpc.CallOption) (*SkuUpdateResp, error)
	// 删除Sku
	SkuDelete(ctx context.Context, in *SkuDeleteReq, opts ...grpc.CallOption) (*SkuDeleteResp, error)
	// sku详情
	SkuInfo(ctx context.Context, in *SkuInfoReq, opts ...grpc.CallOption) (*SkuInfoResp, error)
}

type pmsClient struct {
	cc grpc.ClientConnInterface
}

func NewPmsClient(cc grpc.ClientConnInterface) PmsClient {
	return &pmsClient{cc}
}

func (c *pmsClient) CategoryAdd(ctx context.Context, in *CategoryAddReq, opts ...grpc.CallOption) (*CategoryAddResp, error) {
	out := new(CategoryAddResp)
	err := c.cc.Invoke(ctx, Pms_CategoryAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	out := new(CategoryListResp)
	err := c.cc.Invoke(ctx, Pms_CategoryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*CategoryUpdateResp, error) {
	out := new(CategoryUpdateResp)
	err := c.cc.Invoke(ctx, Pms_CategoryUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error) {
	out := new(CategoryDeleteResp)
	err := c.cc.Invoke(ctx, Pms_CategoryDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) AttributeAdd(ctx context.Context, in *AttributeAddReq, opts ...grpc.CallOption) (*AttributeAddResp, error) {
	out := new(AttributeAddResp)
	err := c.cc.Invoke(ctx, Pms_AttributeAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) AttributeList(ctx context.Context, in *AttributeListReq, opts ...grpc.CallOption) (*AttributeListResp, error) {
	out := new(AttributeListResp)
	err := c.cc.Invoke(ctx, Pms_AttributeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) AttributeUpdate(ctx context.Context, in *AttributeUpdateReq, opts ...grpc.CallOption) (*AttributeUpdateResp, error) {
	out := new(AttributeUpdateResp)
	err := c.cc.Invoke(ctx, Pms_AttributeUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) AttributeDelete(ctx context.Context, in *AttributeDeleteReq, opts ...grpc.CallOption) (*AttributeDeleteResp, error) {
	out := new(AttributeDeleteResp)
	err := c.cc.Invoke(ctx, Pms_AttributeDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error) {
	out := new(ProductAddResp)
	err := c.cc.Invoke(ctx, Pms_ProductAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error) {
	out := new(ProductListResp)
	err := c.cc.Invoke(ctx, Pms_ProductList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error) {
	out := new(ProductUpdateResp)
	err := c.cc.Invoke(ctx, Pms_ProductUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error) {
	out := new(ProductDeleteResp)
	err := c.cc.Invoke(ctx, Pms_ProductDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ProductInfo(ctx context.Context, in *ProductInfoReq, opts ...grpc.CallOption) (*ProductInfoResp, error) {
	out := new(ProductInfoResp)
	err := c.cc.Invoke(ctx, Pms_ProductInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) SkuAdd(ctx context.Context, in *SkuAddReq, opts ...grpc.CallOption) (*SkuAddResp, error) {
	out := new(SkuAddResp)
	err := c.cc.Invoke(ctx, Pms_SkuAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) SkuList(ctx context.Context, in *SkuListReq, opts ...grpc.CallOption) (*SkuListResp, error) {
	out := new(SkuListResp)
	err := c.cc.Invoke(ctx, Pms_SkuList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) SkuUpdate(ctx context.Context, in *SkuUpdateReq, opts ...grpc.CallOption) (*SkuUpdateResp, error) {
	out := new(SkuUpdateResp)
	err := c.cc.Invoke(ctx, Pms_SkuUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) SkuDelete(ctx context.Context, in *SkuDeleteReq, opts ...grpc.CallOption) (*SkuDeleteResp, error) {
	out := new(SkuDeleteResp)
	err := c.cc.Invoke(ctx, Pms_SkuDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) SkuInfo(ctx context.Context, in *SkuInfoReq, opts ...grpc.CallOption) (*SkuInfoResp, error) {
	out := new(SkuInfoResp)
	err := c.cc.Invoke(ctx, Pms_SkuInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PmsServer is the server API for Pms service.
// All implementations must embed UnimplementedPmsServer
// for forward compatibility
type PmsServer interface {
	// 添加分类
	CategoryAdd(context.Context, *CategoryAddReq) (*CategoryAddResp, error)
	// 分类列表
	CategoryList(context.Context, *CategoryListReq) (*CategoryListResp, error)
	// 更新分类
	CategoryUpdate(context.Context, *CategoryUpdateReq) (*CategoryUpdateResp, error)
	// 删除分类
	CategoryDelete(context.Context, *CategoryDeleteReq) (*CategoryDeleteResp, error)
	// 添加属性
	AttributeAdd(context.Context, *AttributeAddReq) (*AttributeAddResp, error)
	// 属性列表
	AttributeList(context.Context, *AttributeListReq) (*AttributeListResp, error)
	// 更新属性
	AttributeUpdate(context.Context, *AttributeUpdateReq) (*AttributeUpdateResp, error)
	// 删除属性
	AttributeDelete(context.Context, *AttributeDeleteReq) (*AttributeDeleteResp, error)
	// 添加商品
	ProductAdd(context.Context, *ProductAddReq) (*ProductAddResp, error)
	// 商品列表
	ProductList(context.Context, *ProductListReq) (*ProductListResp, error)
	// 更新商品
	ProductUpdate(context.Context, *ProductUpdateReq) (*ProductUpdateResp, error)
	// 删除商品
	ProductDelete(context.Context, *ProductDeleteReq) (*ProductDeleteResp, error)
	// 查询商品详情
	ProductInfo(context.Context, *ProductInfoReq) (*ProductInfoResp, error)
	// 添加Sku
	SkuAdd(context.Context, *SkuAddReq) (*SkuAddResp, error)
	// Sku列表
	SkuList(context.Context, *SkuListReq) (*SkuListResp, error)
	// 更新Sku
	SkuUpdate(context.Context, *SkuUpdateReq) (*SkuUpdateResp, error)
	// 删除Sku
	SkuDelete(context.Context, *SkuDeleteReq) (*SkuDeleteResp, error)
	// sku详情
	SkuInfo(context.Context, *SkuInfoReq) (*SkuInfoResp, error)
	mustEmbedUnimplementedPmsServer()
}

// UnimplementedPmsServer must be embedded to have forward compatible implementations.
type UnimplementedPmsServer struct {
}

func (UnimplementedPmsServer) CategoryAdd(context.Context, *CategoryAddReq) (*CategoryAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryAdd not implemented")
}
func (UnimplementedPmsServer) CategoryList(context.Context, *CategoryListReq) (*CategoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryList not implemented")
}
func (UnimplementedPmsServer) CategoryUpdate(context.Context, *CategoryUpdateReq) (*CategoryUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryUpdate not implemented")
}
func (UnimplementedPmsServer) CategoryDelete(context.Context, *CategoryDeleteReq) (*CategoryDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryDelete not implemented")
}
func (UnimplementedPmsServer) AttributeAdd(context.Context, *AttributeAddReq) (*AttributeAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeAdd not implemented")
}
func (UnimplementedPmsServer) AttributeList(context.Context, *AttributeListReq) (*AttributeListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeList not implemented")
}
func (UnimplementedPmsServer) AttributeUpdate(context.Context, *AttributeUpdateReq) (*AttributeUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeUpdate not implemented")
}
func (UnimplementedPmsServer) AttributeDelete(context.Context, *AttributeDeleteReq) (*AttributeDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttributeDelete not implemented")
}
func (UnimplementedPmsServer) ProductAdd(context.Context, *ProductAddReq) (*ProductAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductAdd not implemented")
}
func (UnimplementedPmsServer) ProductList(context.Context, *ProductListReq) (*ProductListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductList not implemented")
}
func (UnimplementedPmsServer) ProductUpdate(context.Context, *ProductUpdateReq) (*ProductUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductUpdate not implemented")
}
func (UnimplementedPmsServer) ProductDelete(context.Context, *ProductDeleteReq) (*ProductDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDelete not implemented")
}
func (UnimplementedPmsServer) ProductInfo(context.Context, *ProductInfoReq) (*ProductInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfo not implemented")
}
func (UnimplementedPmsServer) SkuAdd(context.Context, *SkuAddReq) (*SkuAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SkuAdd not implemented")
}
func (UnimplementedPmsServer) SkuList(context.Context, *SkuListReq) (*SkuListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SkuList not implemented")
}
func (UnimplementedPmsServer) SkuUpdate(context.Context, *SkuUpdateReq) (*SkuUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SkuUpdate not implemented")
}
func (UnimplementedPmsServer) SkuDelete(context.Context, *SkuDeleteReq) (*SkuDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SkuDelete not implemented")
}
func (UnimplementedPmsServer) SkuInfo(context.Context, *SkuInfoReq) (*SkuInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SkuInfo not implemented")
}
func (UnimplementedPmsServer) mustEmbedUnimplementedPmsServer() {}

// UnsafePmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PmsServer will
// result in compilation errors.
type UnsafePmsServer interface {
	mustEmbedUnimplementedPmsServer()
}

func RegisterPmsServer(s grpc.ServiceRegistrar, srv PmsServer) {
	s.RegisterService(&Pms_ServiceDesc, srv)
}

func _Pms_CategoryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).CategoryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_CategoryAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).CategoryAdd(ctx, req.(*CategoryAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_CategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).CategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_CategoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).CategoryList(ctx, req.(*CategoryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_CategoryUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).CategoryUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_CategoryUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).CategoryUpdate(ctx, req.(*CategoryUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_CategoryDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).CategoryDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_CategoryDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).CategoryDelete(ctx, req.(*CategoryDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_AttributeAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).AttributeAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_AttributeAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).AttributeAdd(ctx, req.(*AttributeAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_AttributeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).AttributeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_AttributeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).AttributeList(ctx, req.(*AttributeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_AttributeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).AttributeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_AttributeUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).AttributeUpdate(ctx, req.(*AttributeUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_AttributeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).AttributeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_AttributeDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).AttributeDelete(ctx, req.(*AttributeDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ProductAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ProductAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_ProductAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ProductAdd(ctx, req.(*ProductAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_ProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ProductList(ctx, req.(*ProductListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ProductUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ProductUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_ProductUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ProductUpdate(ctx, req.(*ProductUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ProductDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ProductDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_ProductDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ProductDelete(ctx, req.(*ProductDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_ProductInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ProductInfo(ctx, req.(*ProductInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_SkuAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).SkuAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_SkuAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).SkuAdd(ctx, req.(*SkuAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_SkuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).SkuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_SkuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).SkuList(ctx, req.(*SkuListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_SkuUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).SkuUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_SkuUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).SkuUpdate(ctx, req.(*SkuUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_SkuDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).SkuDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_SkuDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).SkuDelete(ctx, req.(*SkuDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_SkuInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SkuInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).SkuInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pms_SkuInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).SkuInfo(ctx, req.(*SkuInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pms_ServiceDesc is the grpc.ServiceDesc for Pms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pms.Pms",
	HandlerType: (*PmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CategoryAdd",
			Handler:    _Pms_CategoryAdd_Handler,
		},
		{
			MethodName: "CategoryList",
			Handler:    _Pms_CategoryList_Handler,
		},
		{
			MethodName: "CategoryUpdate",
			Handler:    _Pms_CategoryUpdate_Handler,
		},
		{
			MethodName: "CategoryDelete",
			Handler:    _Pms_CategoryDelete_Handler,
		},
		{
			MethodName: "AttributeAdd",
			Handler:    _Pms_AttributeAdd_Handler,
		},
		{
			MethodName: "AttributeList",
			Handler:    _Pms_AttributeList_Handler,
		},
		{
			MethodName: "AttributeUpdate",
			Handler:    _Pms_AttributeUpdate_Handler,
		},
		{
			MethodName: "AttributeDelete",
			Handler:    _Pms_AttributeDelete_Handler,
		},
		{
			MethodName: "ProductAdd",
			Handler:    _Pms_ProductAdd_Handler,
		},
		{
			MethodName: "ProductList",
			Handler:    _Pms_ProductList_Handler,
		},
		{
			MethodName: "ProductUpdate",
			Handler:    _Pms_ProductUpdate_Handler,
		},
		{
			MethodName: "ProductDelete",
			Handler:    _Pms_ProductDelete_Handler,
		},
		{
			MethodName: "ProductInfo",
			Handler:    _Pms_ProductInfo_Handler,
		},
		{
			MethodName: "SkuAdd",
			Handler:    _Pms_SkuAdd_Handler,
		},
		{
			MethodName: "SkuList",
			Handler:    _Pms_SkuList_Handler,
		},
		{
			MethodName: "SkuUpdate",
			Handler:    _Pms_SkuUpdate_Handler,
		},
		{
			MethodName: "SkuDelete",
			Handler:    _Pms_SkuDelete_Handler,
		},
		{
			MethodName: "SkuInfo",
			Handler:    _Pms_SkuInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pms.proto",
}
