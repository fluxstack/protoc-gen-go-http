// Code generated by protoc-gen-gohttp. DO NOT EDIT.
// source: routeguide/route_guide.proto

package routeguidepb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RouteGuideHTTPService is the server API for RouteGuide service.
type RouteGuideHTTPService interface {
	GetFeature(context.Context, *Point) (*Feature, error)
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	RecordRoute(RouteGuide_RecordRouteServer) error
	RouteChat(RouteGuide_RouteChatServer) error
}

// RouteGuideHTTPConverter has a function to convert RouteGuideHTTPService interface to http.HandlerFunc.
type RouteGuideHTTPConverter struct {
	srv RouteGuideHTTPService
}

// NewRouteGuideHTTPConverter returns RouteGuideHTTPConverter.
func NewRouteGuideHTTPConverter(srv RouteGuideHTTPService) *RouteGuideHTTPConverter {
	return &RouteGuideHTTPConverter{
		srv: srv,
	}
}

// GetFeature returns RouteGuideHTTPService interface's GetFeature converted to http.HandlerFunc.
func (h *RouteGuideHTTPConverter) GetFeature(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type")); contentType {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					if err := json.NewEncoder(w).Encode(p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		arg := &Point{}
		contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if r.Method != http.MethodGet {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}

			switch contentType {
			case "application/protobuf", "application/x-protobuf":
				if err := proto.Unmarshal(body, arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			case "application/json":
				if err := jsonpb.Unmarshal(bytes.NewBuffer(body), arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
				cb(ctx, w, r, nil, nil, err)
				return
			}
		}

		n := len(interceptors)
		chained := func(ctx context.Context, arg interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return currentInter(currentCtx, currentReq, info, currentHandler)
				}
			}

			chainedHandler := handler
			for i := n - 1; i >= 0; i-- {
				chainedHandler = chainer(interceptors[i], chainedHandler)
			}
			return chainedHandler(ctx, arg)
		}

		info := &grpc.UnaryServerInfo{
			Server:     h.srv,
			FullMethod: "/routeguide.RouteGuide/GetFeature",
		}

		handler := func(c context.Context, req interface{}) (interface{}, error) {
			return h.srv.GetFeature(c, req.(*Point))
		}

		iret, err := chained(ctx, arg, info, handler)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		ret, ok := iret.(*Feature)
		if !ok {
			cb(ctx, w, r, arg, nil, fmt.Errorf("/routeguide.RouteGuide/GetFeature: interceptors have not return Feature"))
			return
		}

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			m := jsonpb.Marshaler{
				EnumsAsInts:  true,
				EmitDefaults: true,
			}
			if err := m.Marshal(w, ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

// GetFeatureWithName returns Service name, Method name and RouteGuideHTTPService interface's GetFeature converted to http.HandlerFunc.
func (h *RouteGuideHTTPConverter) GetFeatureWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) (string, string, http.HandlerFunc) {
	return "RouteGuide", "GetFeature", h.GetFeature(cb, interceptors...)
}
