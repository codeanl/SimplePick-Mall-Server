// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package pmsclient

import (
	"context"

	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Attribute                     = pms.Attribute
	AttributeAddReq               = pms.AttributeAddReq
	AttributeAddResp              = pms.AttributeAddResp
	AttributeCategoryAddReq       = pms.AttributeCategoryAddReq
	AttributeCategoryAddResp      = pms.AttributeCategoryAddResp
	AttributeCategoryDeleteAddReq = pms.AttributeCategoryDeleteAddReq
	AttributeCategoryDeleteResp   = pms.AttributeCategoryDeleteResp
	AttributeCategoryListData     = pms.AttributeCategoryListData
	AttributeCategoryListReq      = pms.AttributeCategoryListReq
	AttributeCategoryListResp     = pms.AttributeCategoryListResp
	AttributeCategoryUpdateReq    = pms.AttributeCategoryUpdateReq
	AttributeCategoryUpdateResp   = pms.AttributeCategoryUpdateResp
	AttributeDeleteReq            = pms.AttributeDeleteReq
	AttributeDeleteResp           = pms.AttributeDeleteResp
	AttributeListData             = pms.AttributeListData
	AttributeListReq              = pms.AttributeListReq
	AttributeListResp             = pms.AttributeListResp
	AttributeUpdateReq            = pms.AttributeUpdateReq
	AttributeUpdateResp           = pms.AttributeUpdateResp
	AttributeValueList            = pms.AttributeValueList
	CategoryAddReq                = pms.CategoryAddReq
	CategoryAddResp               = pms.CategoryAddResp
	CategoryDeleteReq             = pms.CategoryDeleteReq
	CategoryDeleteResp            = pms.CategoryDeleteResp
	CategoryListData              = pms.CategoryListData
	CategoryListReq               = pms.CategoryListReq
	CategoryListResp              = pms.CategoryListResp
	CategoryUpdateReq             = pms.CategoryUpdateReq
	CategoryUpdateResp            = pms.CategoryUpdateResp
	MerchantApplysListData        = pms.MerchantApplysListData
	MerchantsAddReq               = pms.MerchantsAddReq
	MerchantsAddResp              = pms.MerchantsAddResp
	MerchantsApplyAddReq          = pms.MerchantsApplyAddReq
	MerchantsApplyAddResp         = pms.MerchantsApplyAddResp
	MerchantsApplyDeleteReq       = pms.MerchantsApplyDeleteReq
	MerchantsApplyDeleteResp      = pms.MerchantsApplyDeleteResp
	MerchantsApplyListReq         = pms.MerchantsApplyListReq
	MerchantsApplyListResp        = pms.MerchantsApplyListResp
	MerchantsApplyUpdateReq       = pms.MerchantsApplyUpdateReq
	MerchantsApplyUpdateResp      = pms.MerchantsApplyUpdateResp
	MerchantsDeleteReq            = pms.MerchantsDeleteReq
	MerchantsDeleteResp           = pms.MerchantsDeleteResp
	MerchantsInfoReq              = pms.MerchantsInfoReq
	MerchantsInfoResp             = pms.MerchantsInfoResp
	MerchantsListData             = pms.MerchantsListData
	MerchantsListReq              = pms.MerchantsListReq
	MerchantsListResp             = pms.MerchantsListResp
	MerchantsUpdateReq            = pms.MerchantsUpdateReq
	MerchantsUpdateResp           = pms.MerchantsUpdateResp
	PlaceAddReq                   = pms.PlaceAddReq
	PlaceAddResp                  = pms.PlaceAddResp
	PlaceDeleteReq                = pms.PlaceDeleteReq
	PlaceDeleteResp               = pms.PlaceDeleteResp
	PlaceInfoReq                  = pms.PlaceInfoReq
	PlaceInfoResp                 = pms.PlaceInfoResp
	PlaceListData                 = pms.PlaceListData
	PlaceListReq                  = pms.PlaceListReq
	PlaceListResp                 = pms.PlaceListResp
	PlaceUpdateReq                = pms.PlaceUpdateReq
	PlaceUpdateResp               = pms.PlaceUpdateResp
	ProductAddReq                 = pms.ProductAddReq
	ProductAddResp                = pms.ProductAddResp
	ProductDeleteReq              = pms.ProductDeleteReq
	ProductDeleteResp             = pms.ProductDeleteResp
	ProductInfoReq                = pms.ProductInfoReq
	ProductInfoResp               = pms.ProductInfoResp
	ProductListData               = pms.ProductListData
	ProductListReq                = pms.ProductListReq
	ProductListResp               = pms.ProductListResp
	ProductUpdateReq              = pms.ProductUpdateReq
	ProductUpdateResp             = pms.ProductUpdateResp
	Size                          = pms.Size
	SkuAddReq                     = pms.SkuAddReq
	SkuAddResp                    = pms.SkuAddResp
	SkuDeleteReq                  = pms.SkuDeleteReq
	SkuDeleteResp                 = pms.SkuDeleteResp
	SkuInfoReq                    = pms.SkuInfoReq
	SkuInfoResp                   = pms.SkuInfoResp
	SkuListData                   = pms.SkuListData
	SkuListReq                    = pms.SkuListReq
	SkuListResp                   = pms.SkuListResp
	SkuUpdateReq                  = pms.SkuUpdateReq
	SkuUpdateResp                 = pms.SkuUpdateResp
	Values                        = pms.Values

	Pms interface {
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
		// 添加属性分类
		AttributeCategoryAdd(ctx context.Context, in *AttributeCategoryAddReq, opts ...grpc.CallOption) (*AttributeCategoryAddResp, error)
		// 更新属性分类
		AttributeCategoryUpdate(ctx context.Context, in *AttributeCategoryUpdateReq, opts ...grpc.CallOption) (*AttributeCategoryUpdateResp, error)
		// 删除属性分类
		AttributeCategoryDelete(ctx context.Context, in *AttributeCategoryDeleteAddReq, opts ...grpc.CallOption) (*AttributeCategoryDeleteResp, error)
		// 属性分类列表
		AttributeCategoryList(ctx context.Context, in *AttributeCategoryListReq, opts ...grpc.CallOption) (*AttributeCategoryListResp, error)
		// 添加商家
		MerchantsAdd(ctx context.Context, in *MerchantsAddReq, opts ...grpc.CallOption) (*MerchantsAddResp, error)
		// 商家列表
		MerchantsList(ctx context.Context, in *MerchantsListReq, opts ...grpc.CallOption) (*MerchantsListResp, error)
		// 更新商家
		MerchantsUpdate(ctx context.Context, in *MerchantsUpdateReq, opts ...grpc.CallOption) (*MerchantsUpdateResp, error)
		// 删除商家
		MerchantsDelete(ctx context.Context, in *MerchantsDeleteReq, opts ...grpc.CallOption) (*MerchantsDeleteResp, error)
		// 商家详情
		MerchantsInfo(ctx context.Context, in *MerchantsInfoReq, opts ...grpc.CallOption) (*MerchantsInfoResp, error)
		// 添加商家申请
		MerchantsApplyAdd(ctx context.Context, in *MerchantsApplyAddReq, opts ...grpc.CallOption) (*MerchantsApplyAddResp, error)
		// 商家申请列表
		MerchantsApplyList(ctx context.Context, in *MerchantsApplyListReq, opts ...grpc.CallOption) (*MerchantsApplyListResp, error)
		// 更新商家申请
		MerchantsApplyUpdate(ctx context.Context, in *MerchantsApplyUpdateReq, opts ...grpc.CallOption) (*MerchantsApplyUpdateResp, error)
		// 删除商家申请
		MerchantsApplyDelete(ctx context.Context, in *MerchantsApplyDeleteReq, opts ...grpc.CallOption) (*MerchantsApplyDeleteResp, error)
		// 添加自提点
		PlaceAdd(ctx context.Context, in *PlaceAddReq, opts ...grpc.CallOption) (*PlaceAddResp, error)
		// 自提点列表
		PlaceList(ctx context.Context, in *PlaceListReq, opts ...grpc.CallOption) (*PlaceListResp, error)
		// 更新自提点
		PlaceUpdate(ctx context.Context, in *PlaceUpdateReq, opts ...grpc.CallOption) (*PlaceUpdateResp, error)
		// 删除自提点
		PlaceDelete(ctx context.Context, in *PlaceDeleteReq, opts ...grpc.CallOption) (*PlaceDeleteResp, error)
		// 自提点详情
		PlaceInfo(ctx context.Context, in *PlaceInfoReq, opts ...grpc.CallOption) (*PlaceInfoResp, error)
	}

	defaultPms struct {
		cli zrpc.Client
	}
)

func NewPms(cli zrpc.Client) Pms {
	return &defaultPms{
		cli: cli,
	}
}

// 添加分类
func (m *defaultPms) CategoryAdd(ctx context.Context, in *CategoryAddReq, opts ...grpc.CallOption) (*CategoryAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryAdd(ctx, in, opts...)
}

// 分类列表
func (m *defaultPms) CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryList(ctx, in, opts...)
}

// 更新分类
func (m *defaultPms) CategoryUpdate(ctx context.Context, in *CategoryUpdateReq, opts ...grpc.CallOption) (*CategoryUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryUpdate(ctx, in, opts...)
}

// 删除分类
func (m *defaultPms) CategoryDelete(ctx context.Context, in *CategoryDeleteReq, opts ...grpc.CallOption) (*CategoryDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.CategoryDelete(ctx, in, opts...)
}

// 添加属性
func (m *defaultPms) AttributeAdd(ctx context.Context, in *AttributeAddReq, opts ...grpc.CallOption) (*AttributeAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeAdd(ctx, in, opts...)
}

// 属性列表
func (m *defaultPms) AttributeList(ctx context.Context, in *AttributeListReq, opts ...grpc.CallOption) (*AttributeListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeList(ctx, in, opts...)
}

// 更新属性
func (m *defaultPms) AttributeUpdate(ctx context.Context, in *AttributeUpdateReq, opts ...grpc.CallOption) (*AttributeUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeUpdate(ctx, in, opts...)
}

// 删除属性
func (m *defaultPms) AttributeDelete(ctx context.Context, in *AttributeDeleteReq, opts ...grpc.CallOption) (*AttributeDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeDelete(ctx, in, opts...)
}

// 添加商品
func (m *defaultPms) ProductAdd(ctx context.Context, in *ProductAddReq, opts ...grpc.CallOption) (*ProductAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductAdd(ctx, in, opts...)
}

// 商品列表
func (m *defaultPms) ProductList(ctx context.Context, in *ProductListReq, opts ...grpc.CallOption) (*ProductListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductList(ctx, in, opts...)
}

// 更新商品
func (m *defaultPms) ProductUpdate(ctx context.Context, in *ProductUpdateReq, opts ...grpc.CallOption) (*ProductUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductUpdate(ctx, in, opts...)
}

// 删除商品
func (m *defaultPms) ProductDelete(ctx context.Context, in *ProductDeleteReq, opts ...grpc.CallOption) (*ProductDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductDelete(ctx, in, opts...)
}

// 查询商品详情
func (m *defaultPms) ProductInfo(ctx context.Context, in *ProductInfoReq, opts ...grpc.CallOption) (*ProductInfoResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.ProductInfo(ctx, in, opts...)
}

// 添加Sku
func (m *defaultPms) SkuAdd(ctx context.Context, in *SkuAddReq, opts ...grpc.CallOption) (*SkuAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuAdd(ctx, in, opts...)
}

// Sku列表
func (m *defaultPms) SkuList(ctx context.Context, in *SkuListReq, opts ...grpc.CallOption) (*SkuListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuList(ctx, in, opts...)
}

// 更新Sku
func (m *defaultPms) SkuUpdate(ctx context.Context, in *SkuUpdateReq, opts ...grpc.CallOption) (*SkuUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuUpdate(ctx, in, opts...)
}

// 删除Sku
func (m *defaultPms) SkuDelete(ctx context.Context, in *SkuDeleteReq, opts ...grpc.CallOption) (*SkuDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuDelete(ctx, in, opts...)
}

// sku详情
func (m *defaultPms) SkuInfo(ctx context.Context, in *SkuInfoReq, opts ...grpc.CallOption) (*SkuInfoResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.SkuInfo(ctx, in, opts...)
}

// 添加属性分类
func (m *defaultPms) AttributeCategoryAdd(ctx context.Context, in *AttributeCategoryAddReq, opts ...grpc.CallOption) (*AttributeCategoryAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryAdd(ctx, in, opts...)
}

// 更新属性分类
func (m *defaultPms) AttributeCategoryUpdate(ctx context.Context, in *AttributeCategoryUpdateReq, opts ...grpc.CallOption) (*AttributeCategoryUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryUpdate(ctx, in, opts...)
}

// 删除属性分类
func (m *defaultPms) AttributeCategoryDelete(ctx context.Context, in *AttributeCategoryDeleteAddReq, opts ...grpc.CallOption) (*AttributeCategoryDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryDelete(ctx, in, opts...)
}

// 属性分类列表
func (m *defaultPms) AttributeCategoryList(ctx context.Context, in *AttributeCategoryListReq, opts ...grpc.CallOption) (*AttributeCategoryListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.AttributeCategoryList(ctx, in, opts...)
}

// 添加商家
func (m *defaultPms) MerchantsAdd(ctx context.Context, in *MerchantsAddReq, opts ...grpc.CallOption) (*MerchantsAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsAdd(ctx, in, opts...)
}

// 商家列表
func (m *defaultPms) MerchantsList(ctx context.Context, in *MerchantsListReq, opts ...grpc.CallOption) (*MerchantsListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsList(ctx, in, opts...)
}

// 更新商家
func (m *defaultPms) MerchantsUpdate(ctx context.Context, in *MerchantsUpdateReq, opts ...grpc.CallOption) (*MerchantsUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsUpdate(ctx, in, opts...)
}

// 删除商家
func (m *defaultPms) MerchantsDelete(ctx context.Context, in *MerchantsDeleteReq, opts ...grpc.CallOption) (*MerchantsDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsDelete(ctx, in, opts...)
}

// 商家详情
func (m *defaultPms) MerchantsInfo(ctx context.Context, in *MerchantsInfoReq, opts ...grpc.CallOption) (*MerchantsInfoResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsInfo(ctx, in, opts...)
}

// 添加商家申请
func (m *defaultPms) MerchantsApplyAdd(ctx context.Context, in *MerchantsApplyAddReq, opts ...grpc.CallOption) (*MerchantsApplyAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyAdd(ctx, in, opts...)
}

// 商家申请列表
func (m *defaultPms) MerchantsApplyList(ctx context.Context, in *MerchantsApplyListReq, opts ...grpc.CallOption) (*MerchantsApplyListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyList(ctx, in, opts...)
}

// 更新商家申请
func (m *defaultPms) MerchantsApplyUpdate(ctx context.Context, in *MerchantsApplyUpdateReq, opts ...grpc.CallOption) (*MerchantsApplyUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyUpdate(ctx, in, opts...)
}

// 删除商家申请
func (m *defaultPms) MerchantsApplyDelete(ctx context.Context, in *MerchantsApplyDeleteReq, opts ...grpc.CallOption) (*MerchantsApplyDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.MerchantsApplyDelete(ctx, in, opts...)
}

// 添加自提点
func (m *defaultPms) PlaceAdd(ctx context.Context, in *PlaceAddReq, opts ...grpc.CallOption) (*PlaceAddResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.PlaceAdd(ctx, in, opts...)
}

// 自提点列表
func (m *defaultPms) PlaceList(ctx context.Context, in *PlaceListReq, opts ...grpc.CallOption) (*PlaceListResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.PlaceList(ctx, in, opts...)
}

// 更新自提点
func (m *defaultPms) PlaceUpdate(ctx context.Context, in *PlaceUpdateReq, opts ...grpc.CallOption) (*PlaceUpdateResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.PlaceUpdate(ctx, in, opts...)
}

// 删除自提点
func (m *defaultPms) PlaceDelete(ctx context.Context, in *PlaceDeleteReq, opts ...grpc.CallOption) (*PlaceDeleteResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.PlaceDelete(ctx, in, opts...)
}

// 自提点详情
func (m *defaultPms) PlaceInfo(ctx context.Context, in *PlaceInfoReq, opts ...grpc.CallOption) (*PlaceInfoResp, error) {
	client := pms.NewPmsClient(m.cli.Conn())
	return client.PlaceInfo(ctx, in, opts...)
}
