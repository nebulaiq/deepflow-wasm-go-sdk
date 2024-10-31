/*
 * Copyright (c) 2022 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/deepflowio/deepflow-wasm-go-sdk/sdk"
	"slices"
)

type parser struct {
}

func (p parser) HookIn() []sdk.HookBitmap {
	return []sdk.HookBitmap{
		//sdk.HOOK_POINT_HTTP_REQ,
		sdk.HOOK_POINT_PAYLOAD_PARSE,
		//sdk.HOOK_POINT_HTTP_RESP,
	}
}

func (p parser) OnHttpReq(ctx *sdk.HttpReqCtx) sdk.Action {
	return sdk.ActionNext()
}

//func (p parser) OnHttpReq(ctx *sdk.HttpReqCtx) sdk.Action {
//	baseCtx := &ctx.BaseCtx
//
//	payload, err := baseCtx.GetPayload()
//	if err != nil {
//		return sdk.ActionAbortWithErr(err)
//	}
//
//	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(payload)))
//	if err != nil {
//		return sdk.ActionAbortWithErr(err)
//	}
//
//	excludeNames := []string{
//		"Accept",
//		"Connection",
//		"User-Agent",
//		"Accept-Encoding",
//		"Metadata",
//		"X-Ms-Agent-Name",
//		"X-Ms-Version",
//		"X-Ms-Containerid",
//		"X-Ms-Host-Config-Name",
//		"X-Prometheus-Scrape-Timeout-Seconds",
//		"X-Prometheus-Scrape-Timeout-Seconds",
//		"X-Forwarded-For",
//		"X-Forwarded-Uri",
//		"Content-Length",
//		"Content-Type",
//		"Referer",
//		"Accept-Language",
//		"Upgrade-Insecure-Requests",
//		"Cookie",
//		"X-User-Id",
//		"X-User-Type",
//		"Origin",
//		"Cache-Control",
//		"Sec-Websocket-Extensions",
//		"Pragma",
//		"Sec-Websocket-Version",
//		"Sec-Websocket-Key",
//	}
//	// Loop over header names
//	for name, values := range req.Header {
//		if slices.Contains(excludeNames, name) {
//			continue
//		}
//		// Loop over all values for the name.
//		for _, value := range values {
//			sdk.Warn("got header", name, value, req.URL.Path)
//		}
//	}
//
//	//var (
//	//	traceID      string
//	//	spanID       string
//	//	parentSpanID string
//	//	trace        *sdk.Trace
//	//)
//	//
//	//key := "Custom-Trace-Info"
//	//
//	//traceInfo, ok := req.Header[key]
//	//if ok && len(traceInfo) != 0 {
//	//	sdk.Warn("got trace info from header", traceInfo)
//	//	s := strings.Split(traceInfo[0], ":")
//	//	if len(s) == 3 {
//	//		traceID = strings.TrimSpace(s[0])
//	//		spanID = strings.TrimSpace(s[1])
//	//		parentSpanID = strings.TrimSpace(s[2])
//	//	}
//	//}
//	//
//	//if traceID != "" && spanID != "" {
//	//	trace = &sdk.Trace{
//	//		TraceID:      traceID,
//	//		SpanID:       spanID,
//	//		ParentSpanID: parentSpanID,
//	//	}
//	//} else {
//	//	s := strconv.FormatUint(baseCtx.FlowID, 10)
//	//	trace = &sdk.Trace{
//	//		TraceID: s,
//	//		SpanID:  s,
//	//	}
//	//}
//	//
//	//val := trace.TraceID + ":" + trace.SpanID + ":" + trace.ParentSpanID
//	//
//	//req.Header.Set(key, val)
//	//attr := []sdk.KeyVal{
//	//	{
//	//		Key: key,
//	//		Val: val,
//	//	},
//	//}
//
//	//return sdk.HttpReqActionAbortWithResult(nil, trace, attr)
//	return sdk.ActionAbort()
//}

/*
assume resp as follow:

	HTTP/1.1 200 OK

	{"code": 0, "data": {"user_id": 12345, "register_time": 1682050409}}
*/

func (p parser) OnHttpResp(ctx *sdk.HttpRespCtx) sdk.Action {
	return sdk.ActionNext()
}

/*func (p parser) OnHttpResp(ctx *sdk.HttpRespCtx) sdk.Action {
	baseCtx := &ctx.BaseCtx
	if baseCtx.SrcPort != 8080 {
		return sdk.ActionNext()
	}
	payload, err := baseCtx.GetPayload()
	if err != nil {
		return sdk.ActionAbortWithErr(err)
	}

	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(payload)), nil)
	if err != nil {
		return sdk.ActionAbortWithErr(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return sdk.ActionAbortWithErr(err)
	}
	if fastjson.Exists(body, "code") && fastjson.Exists(body, "data") {
		code := fastjson.GetInt(body, "code")
		if code == 0 {
			userID := fastjson.GetInt(body, "data", "user_id")
			t := fastjson.GetInt(body, "data", "register_time")

			return sdk.HttpRespActionAbortWithResult(nil, nil, []sdk.KeyVal{
				{
					Key: "user_id",
					Val: strconv.Itoa(userID),
				},

				{
					Key: "register_time",
					Val: time.Unix(int64(t), 0).String(),
				},
			})
		}

	}
	return sdk.ActionAbort()

}*/

func (p parser) OnCheckPayload(baseCtx *sdk.ParseCtx) (uint8, string) {
	return 0, ""
}

//func (p parser) OnParsePayload(baseCtx *sdk.ParseCtx) sdk.Action {
//	return sdk.ActionNext()
//}

const (
	HTTP_PROTOCOL          = 0
	GO_HTTP2_EBPF_PROTOCOL = 1
)

func (p parser) OnParsePayload(ctx *sdk.ParseCtx) sdk.Action {
	payload, err := ctx.GetPayload()
	if err != nil {
		return sdk.ActionAbortWithErr(err)
	}

	excludeNames := []string{
		"Accept",
		"Connection",
		"User-Agent",
		"Accept-Encoding",
		"Metadata",
		"X-Ms-Agent-Name",
		"X-Ms-Version",
		"X-Ms-Containerid",
		"X-Ms-Host-Config-Name",
		"X-Prometheus-Scrape-Timeout-Seconds",
		"X-Prometheus-Scrape-Timeout-Seconds",
		"X-Forwarded-For",
		"X-Forwarded-Uri",
		"Content-Length",
		"Content-Type",
		"Referer",
		"Accept-Language",
		"Upgrade-Insecure-Requests",
		"Cookie",
		"X-User-Id",
		"X-User-Type",
		"Origin",
		"Cache-Control",
		"Sec-Websocket-Extensions",
		"Pragma",
		"Sec-Websocket-Version",
		"Sec-Websocket-Key",
	}

	switch ctx.EbpfType {
	case sdk.EbpfTypeGoHttp2Uprobe:
		streamID, key, val, err := parseHeader(payload)
		if err != nil {
			return sdk.ActionAbortWithErr(err)
		}

		if slices.Contains(excludeNames, key) {
			return sdk.ActionAbort()
		}

		sdk.Warn("EbpfTypeGoHttp2Uprobe -- StreamID/Key/Val", streamID, key, val)
		return sdk.ActionAbort()
	default:
		//req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(payload)))
		//if err != nil {
		//	return sdk.ActionAbortWithErr(err)
		//}
		//
		//// Loop over header names
		//for name, values := range req.Header {
		//	if slices.Contains(excludeNames, name) {
		//		continue
		//	}
		//	// Loop over all values for the name.
		//	for _, value := range values {
		//		sdk.Warn("got header", name, value, req.URL.Path)
		//	}
		//}

		return sdk.ActionAbort()
	}
}

/*
fd(4 bytes)
stream id (4 bytes)
header key len (4 bytes)
header value len (4 bytes)
header key value (xxx bytes)
header value value (xxx bytes)
*/
func parseHeader(payload []byte) (uint32, string, string, error) {
	if len(payload) < 16 {
		return 0, "", "", errors.New("header payload too short")
	}

	streamID := binary.LittleEndian.Uint32(payload[4:8])
	keyLen := int(binary.LittleEndian.Uint32(payload[8:12]))
	valLen := int(binary.LittleEndian.Uint32(payload[12:16]))
	if keyLen < 0 || keyLen < 0 || keyLen+valLen+16 > len(payload) {
		return 0, "", "", fmt.Errorf("header kv length too short, key len: %d, val len: %d, payload len: %d", keyLen, valLen, len(payload))
	}

	return streamID, string(payload[16 : 16+keyLen]), string(payload[16+keyLen : 16+keyLen+valLen]), nil
}

func main() {
	sdk.Warn("wasm register http hook")
	sdk.SetParser(parser{})
}
