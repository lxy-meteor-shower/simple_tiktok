// Code generated by Kitex v0.4.4. DO NOT EDIT.

package interactservice

import (
	"context"
	interactdemo "demo/tiktok/kitex_gen/interactdemo"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return interactServiceServiceInfo
}

var interactServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InteractService"
	handlerType := (*interactdemo.InteractService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateFavorite":  kitex.NewMethodInfo(createFavoriteHandler, newCreateFavoriteArgs, newCreateFavoriteResult, false),
		"DeleteFavorite":  kitex.NewMethodInfo(deleteFavoriteHandler, newDeleteFavoriteArgs, newDeleteFavoriteResult, false),
		"GetUserFavorite": kitex.NewMethodInfo(getUserFavoriteHandler, newGetUserFavoriteArgs, newGetUserFavoriteResult, false),
		"CheckFavorite":   kitex.NewMethodInfo(checkFavoriteHandler, newCheckFavoriteArgs, newCheckFavoriteResult, false),
		"CreateComment":   kitex.NewMethodInfo(createCommentHandler, newCreateCommentArgs, newCreateCommentResult, false),
		"DeleteComment":   kitex.NewMethodInfo(deleteCommentHandler, newDeleteCommentArgs, newDeleteCommentResult, false),
		"GetVideoComment": kitex.NewMethodInfo(getVideoCommentHandler, newGetVideoCommentArgs, newGetVideoCommentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "interact",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.FavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).CreateFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateFavoriteArgs:
		success, err := handler.(interactdemo.InteractService).CreateFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newCreateFavoriteArgs() interface{} {
	return &CreateFavoriteArgs{}
}

func newCreateFavoriteResult() interface{} {
	return &CreateFavoriteResult{}
}

type CreateFavoriteArgs struct {
	Req *interactdemo.FavoriteRequest
}

func (p *CreateFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.FavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.FavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateFavoriteArgs_Req_DEFAULT *interactdemo.FavoriteRequest

func (p *CreateFavoriteArgs) GetReq() *interactdemo.FavoriteRequest {
	if !p.IsSetReq() {
		return CreateFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateFavoriteResult struct {
	Success *interactdemo.FavoriteResponse
}

var CreateFavoriteResult_Success_DEFAULT *interactdemo.FavoriteResponse

func (p *CreateFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.FavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateFavoriteResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.FavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateFavoriteResult) GetSuccess() *interactdemo.FavoriteResponse {
	if !p.IsSetSuccess() {
		return CreateFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.FavoriteResponse)
}

func (p *CreateFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.FavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).DeleteFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteFavoriteArgs:
		success, err := handler.(interactdemo.InteractService).DeleteFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteFavoriteArgs() interface{} {
	return &DeleteFavoriteArgs{}
}

func newDeleteFavoriteResult() interface{} {
	return &DeleteFavoriteResult{}
}

type DeleteFavoriteArgs struct {
	Req *interactdemo.FavoriteRequest
}

func (p *DeleteFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.FavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.FavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteFavoriteArgs_Req_DEFAULT *interactdemo.FavoriteRequest

func (p *DeleteFavoriteArgs) GetReq() *interactdemo.FavoriteRequest {
	if !p.IsSetReq() {
		return DeleteFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteFavoriteResult struct {
	Success *interactdemo.FavoriteResponse
}

var DeleteFavoriteResult_Success_DEFAULT *interactdemo.FavoriteResponse

func (p *DeleteFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.FavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteFavoriteResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.FavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteFavoriteResult) GetSuccess() *interactdemo.FavoriteResponse {
	if !p.IsSetSuccess() {
		return DeleteFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.FavoriteResponse)
}

func (p *DeleteFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.GetUserFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).GetUserFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserFavoriteArgs:
		success, err := handler.(interactdemo.InteractService).GetUserFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserFavoriteArgs() interface{} {
	return &GetUserFavoriteArgs{}
}

func newGetUserFavoriteResult() interface{} {
	return &GetUserFavoriteResult{}
}

type GetUserFavoriteArgs struct {
	Req *interactdemo.GetUserFavoriteRequest
}

func (p *GetUserFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.GetUserFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.GetUserFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserFavoriteArgs_Req_DEFAULT *interactdemo.GetUserFavoriteRequest

func (p *GetUserFavoriteArgs) GetReq() *interactdemo.GetUserFavoriteRequest {
	if !p.IsSetReq() {
		return GetUserFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserFavoriteResult struct {
	Success *interactdemo.GetUserFavoriteResponse
}

var GetUserFavoriteResult_Success_DEFAULT *interactdemo.GetUserFavoriteResponse

func (p *GetUserFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.GetUserFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserFavoriteResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.GetUserFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserFavoriteResult) GetSuccess() *interactdemo.GetUserFavoriteResponse {
	if !p.IsSetSuccess() {
		return GetUserFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.GetUserFavoriteResponse)
}

func (p *GetUserFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.CheckFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).CheckFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckFavoriteArgs:
		success, err := handler.(interactdemo.InteractService).CheckFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newCheckFavoriteArgs() interface{} {
	return &CheckFavoriteArgs{}
}

func newCheckFavoriteResult() interface{} {
	return &CheckFavoriteResult{}
}

type CheckFavoriteArgs struct {
	Req *interactdemo.CheckFavoriteRequest
}

func (p *CheckFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.CheckFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.CheckFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckFavoriteArgs_Req_DEFAULT *interactdemo.CheckFavoriteRequest

func (p *CheckFavoriteArgs) GetReq() *interactdemo.CheckFavoriteRequest {
	if !p.IsSetReq() {
		return CheckFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckFavoriteResult struct {
	Success *interactdemo.CheckFavoriteResponse
}

var CheckFavoriteResult_Success_DEFAULT *interactdemo.CheckFavoriteResponse

func (p *CheckFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.CheckFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckFavoriteResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.CheckFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckFavoriteResult) GetSuccess() *interactdemo.CheckFavoriteResponse {
	if !p.IsSetSuccess() {
		return CheckFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.CheckFavoriteResponse)
}

func (p *CheckFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.CreateCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).CreateComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateCommentArgs:
		success, err := handler.(interactdemo.InteractService).CreateComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateCommentResult)
		realResult.Success = success
	}
	return nil
}
func newCreateCommentArgs() interface{} {
	return &CreateCommentArgs{}
}

func newCreateCommentResult() interface{} {
	return &CreateCommentResult{}
}

type CreateCommentArgs struct {
	Req *interactdemo.CreateCommentRequest
}

func (p *CreateCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.CreateCommentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateCommentArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.CreateCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateCommentArgs_Req_DEFAULT *interactdemo.CreateCommentRequest

func (p *CreateCommentArgs) GetReq() *interactdemo.CreateCommentRequest {
	if !p.IsSetReq() {
		return CreateCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateCommentResult struct {
	Success *interactdemo.CommentResponse
}

var CreateCommentResult_Success_DEFAULT *interactdemo.CommentResponse

func (p *CreateCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.CommentResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateCommentResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.CommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateCommentResult) GetSuccess() *interactdemo.CommentResponse {
	if !p.IsSetSuccess() {
		return CreateCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.CommentResponse)
}

func (p *CreateCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.DeleteCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).DeleteComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteCommentArgs:
		success, err := handler.(interactdemo.InteractService).DeleteComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteCommentResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteCommentArgs() interface{} {
	return &DeleteCommentArgs{}
}

func newDeleteCommentResult() interface{} {
	return &DeleteCommentResult{}
}

type DeleteCommentArgs struct {
	Req *interactdemo.DeleteCommentRequest
}

func (p *DeleteCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.DeleteCommentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteCommentArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.DeleteCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteCommentArgs_Req_DEFAULT *interactdemo.DeleteCommentRequest

func (p *DeleteCommentArgs) GetReq() *interactdemo.DeleteCommentRequest {
	if !p.IsSetReq() {
		return DeleteCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteCommentResult struct {
	Success *interactdemo.CommentResponse
}

var DeleteCommentResult_Success_DEFAULT *interactdemo.CommentResponse

func (p *DeleteCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.CommentResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteCommentResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.CommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteCommentResult) GetSuccess() *interactdemo.CommentResponse {
	if !p.IsSetSuccess() {
		return DeleteCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.CommentResponse)
}

func (p *DeleteCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getVideoCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(interactdemo.GetVideoCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(interactdemo.InteractService).GetVideoComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideoCommentArgs:
		success, err := handler.(interactdemo.InteractService).GetVideoComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideoCommentResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideoCommentArgs() interface{} {
	return &GetVideoCommentArgs{}
}

func newGetVideoCommentResult() interface{} {
	return &GetVideoCommentResult{}
}

type GetVideoCommentArgs struct {
	Req *interactdemo.GetVideoCommentRequest
}

func (p *GetVideoCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(interactdemo.GetVideoCommentRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetVideoCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetVideoCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetVideoCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideoCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideoCommentArgs) Unmarshal(in []byte) error {
	msg := new(interactdemo.GetVideoCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideoCommentArgs_Req_DEFAULT *interactdemo.GetVideoCommentRequest

func (p *GetVideoCommentArgs) GetReq() *interactdemo.GetVideoCommentRequest {
	if !p.IsSetReq() {
		return GetVideoCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideoCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetVideoCommentResult struct {
	Success *interactdemo.GetVideoCommentRespense
}

var GetVideoCommentResult_Success_DEFAULT *interactdemo.GetVideoCommentRespense

func (p *GetVideoCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(interactdemo.GetVideoCommentRespense)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetVideoCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetVideoCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetVideoCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideoCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideoCommentResult) Unmarshal(in []byte) error {
	msg := new(interactdemo.GetVideoCommentRespense)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoCommentResult) GetSuccess() *interactdemo.GetVideoCommentRespense {
	if !p.IsSetSuccess() {
		return GetVideoCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideoCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*interactdemo.GetVideoCommentRespense)
}

func (p *GetVideoCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateFavorite(ctx context.Context, Req *interactdemo.FavoriteRequest) (r *interactdemo.FavoriteResponse, err error) {
	var _args CreateFavoriteArgs
	_args.Req = Req
	var _result CreateFavoriteResult
	if err = p.c.Call(ctx, "CreateFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteFavorite(ctx context.Context, Req *interactdemo.FavoriteRequest) (r *interactdemo.FavoriteResponse, err error) {
	var _args DeleteFavoriteArgs
	_args.Req = Req
	var _result DeleteFavoriteResult
	if err = p.c.Call(ctx, "DeleteFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserFavorite(ctx context.Context, Req *interactdemo.GetUserFavoriteRequest) (r *interactdemo.GetUserFavoriteResponse, err error) {
	var _args GetUserFavoriteArgs
	_args.Req = Req
	var _result GetUserFavoriteResult
	if err = p.c.Call(ctx, "GetUserFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckFavorite(ctx context.Context, Req *interactdemo.CheckFavoriteRequest) (r *interactdemo.CheckFavoriteResponse, err error) {
	var _args CheckFavoriteArgs
	_args.Req = Req
	var _result CheckFavoriteResult
	if err = p.c.Call(ctx, "CheckFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateComment(ctx context.Context, Req *interactdemo.CreateCommentRequest) (r *interactdemo.CommentResponse, err error) {
	var _args CreateCommentArgs
	_args.Req = Req
	var _result CreateCommentResult
	if err = p.c.Call(ctx, "CreateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, Req *interactdemo.DeleteCommentRequest) (r *interactdemo.CommentResponse, err error) {
	var _args DeleteCommentArgs
	_args.Req = Req
	var _result DeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoComment(ctx context.Context, Req *interactdemo.GetVideoCommentRequest) (r *interactdemo.GetVideoCommentRespense, err error) {
	var _args GetVideoCommentArgs
	_args.Req = Req
	var _result GetVideoCommentResult
	if err = p.c.Call(ctx, "GetVideoComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
