// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favorservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	favor "toktik-rpc/kitex_gen/favor"
)

func serviceInfo() *kitex.ServiceInfo {
	return favorServiceServiceInfo
}

var favorServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavorService"
	handlerType := (*favor.FavorService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteList":     kitex.NewMethodInfo(favoriteListHandler, newFavoriteListArgs, newFavoriteListResult, false),
		"FavoriteAction":   kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"IsFavoriteVideos": kitex.NewMethodInfo(isFavoriteVideosHandler, newIsFavoriteVideosArgs, newIsFavoriteVideosResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favor",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favor.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favor.FavorService).FavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteListArgs:
		success, err := handler.(favor.FavorService).FavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteListArgs() interface{} {
	return &FavoriteListArgs{}
}

func newFavoriteListResult() interface{} {
	return &FavoriteListResult{}
}

type FavoriteListArgs struct {
	Req *favor.FavoriteListRequest
}

func (p *FavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favor.FavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(favor.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteListArgs_Req_DEFAULT *favor.FavoriteListRequest

func (p *FavoriteListArgs) GetReq() *favor.FavoriteListRequest {
	if !p.IsSetReq() {
		return FavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteListResult struct {
	Success *favor.FavoriteListResponse
}

var FavoriteListResult_Success_DEFAULT *favor.FavoriteListResponse

func (p *FavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favor.FavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteListResult) Unmarshal(in []byte) error {
	msg := new(favor.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteListResult) GetSuccess() *favor.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return FavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*favor.FavoriteListResponse)
}

func (p *FavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteListResult) GetResult() interface{} {
	return p.Success
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favor.FavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favor.FavorService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(favor.FavorService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *favor.FavoriteActionRequest
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favor.FavoriteActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(favor.FavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *favor.FavoriteActionRequest

func (p *FavoriteActionArgs) GetReq() *favor.FavoriteActionRequest {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteActionResult struct {
	Success *favor.FavoriteActionResponse
}

var FavoriteActionResult_Success_DEFAULT *favor.FavoriteActionResponse

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favor.FavoriteActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(favor.FavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *favor.FavoriteActionResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*favor.FavoriteActionResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
	return p.Success
}

func isFavoriteVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favor.IsFavoriteVideosRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favor.FavorService).IsFavoriteVideos(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IsFavoriteVideosArgs:
		success, err := handler.(favor.FavorService).IsFavoriteVideos(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsFavoriteVideosResult)
		realResult.Success = success
	}
	return nil
}
func newIsFavoriteVideosArgs() interface{} {
	return &IsFavoriteVideosArgs{}
}

func newIsFavoriteVideosResult() interface{} {
	return &IsFavoriteVideosResult{}
}

type IsFavoriteVideosArgs struct {
	Req *favor.IsFavoriteVideosRequest
}

func (p *IsFavoriteVideosArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favor.IsFavoriteVideosRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *IsFavoriteVideosArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *IsFavoriteVideosArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *IsFavoriteVideosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IsFavoriteVideosArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *IsFavoriteVideosArgs) Unmarshal(in []byte) error {
	msg := new(favor.IsFavoriteVideosRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsFavoriteVideosArgs_Req_DEFAULT *favor.IsFavoriteVideosRequest

func (p *IsFavoriteVideosArgs) GetReq() *favor.IsFavoriteVideosRequest {
	if !p.IsSetReq() {
		return IsFavoriteVideosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsFavoriteVideosArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *IsFavoriteVideosArgs) GetFirstArgument() interface{} {
	return p.Req
}

type IsFavoriteVideosResult struct {
	Success *favor.IsFavoriteVideosResponse
}

var IsFavoriteVideosResult_Success_DEFAULT *favor.IsFavoriteVideosResponse

func (p *IsFavoriteVideosResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favor.IsFavoriteVideosResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *IsFavoriteVideosResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *IsFavoriteVideosResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *IsFavoriteVideosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IsFavoriteVideosResult")
	}
	return proto.Marshal(p.Success)
}

func (p *IsFavoriteVideosResult) Unmarshal(in []byte) error {
	msg := new(favor.IsFavoriteVideosResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsFavoriteVideosResult) GetSuccess() *favor.IsFavoriteVideosResponse {
	if !p.IsSetSuccess() {
		return IsFavoriteVideosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsFavoriteVideosResult) SetSuccess(x interface{}) {
	p.Success = x.(*favor.IsFavoriteVideosResponse)
}

func (p *IsFavoriteVideosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IsFavoriteVideosResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteList(ctx context.Context, Req *favor.FavoriteListRequest) (r *favor.FavoriteListResponse, err error) {
	var _args FavoriteListArgs
	_args.Req = Req
	var _result FavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *favor.FavoriteActionRequest) (r *favor.FavoriteActionResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavoriteVideos(ctx context.Context, Req *favor.IsFavoriteVideosRequest) (r *favor.IsFavoriteVideosResponse, err error) {
	var _args IsFavoriteVideosArgs
	_args.Req = Req
	var _result IsFavoriteVideosResult
	if err = p.c.Call(ctx, "IsFavoriteVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}