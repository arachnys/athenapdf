// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Overlay functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Configuration data for the highlighting of page elements.
type OverlayHighlightConfig struct {
	ShowInfo           bool     `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowRulers         bool     `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines bool     `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial  bool     `json:"displayAsMaterial,omitempty"`  //
	ContentColor       *DOMRGBA `json:"contentColor,omitempty"`       // The content box highlight fill color (default: transparent).
	PaddingColor       *DOMRGBA `json:"paddingColor,omitempty"`       // The padding highlight fill color (default: transparent).
	BorderColor        *DOMRGBA `json:"borderColor,omitempty"`        // The border highlight fill color (default: transparent).
	MarginColor        *DOMRGBA `json:"marginColor,omitempty"`        // The margin highlight fill color (default: transparent).
	EventTargetColor   *DOMRGBA `json:"eventTargetColor,omitempty"`   // The event target element highlight fill color (default: transparent).
	ShapeColor         *DOMRGBA `json:"shapeColor,omitempty"`         // The shape outside fill color (default: transparent).
	ShapeMarginColor   *DOMRGBA `json:"shapeMarginColor,omitempty"`   // The shape margin fill color (default: transparent).
	SelectorList       string   `json:"selectorList,omitempty"`       // Selectors to highlight relevant nodes.
}

// Fired when the node should be highlighted. This happens after call to <code>setInspectMode</code>.
type OverlayNodeHighlightRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId int `json:"nodeId"` //
	} `json:"Params,omitempty"`
}

// Fired when the node should be inspected. This happens after call to <code>setInspectMode</code> or when user manually inspects an element.
type OverlayInspectNodeRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		BackendNodeId int `json:"backendNodeId"` // Id of the node to inspect.
	} `json:"Params,omitempty"`
}

type Overlay struct {
	target gcdmessage.ChromeTargeter
}

func NewOverlay(target gcdmessage.ChromeTargeter) *Overlay {
	c := &Overlay{target: target}
	return c
}

// Enables domain notifications.
func (c *Overlay) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.enable"})
}

// Disables domain notifications.
func (c *Overlay) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.disable"})
}

type OverlaySetShowPaintRectsParams struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// SetShowPaintRectsWithParams - Requests that backend shows paint rectangles
func (c *Overlay) SetShowPaintRectsWithParams(v *OverlaySetShowPaintRectsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowPaintRects", Params: v})
}

// SetShowPaintRects - Requests that backend shows paint rectangles
// result - True for showing paint rectangles
func (c *Overlay) SetShowPaintRects(result bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowPaintRectsParams
	v.Result = result
	return c.SetShowPaintRectsWithParams(&v)
}

type OverlaySetShowDebugBordersParams struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// SetShowDebugBordersWithParams - Requests that backend shows debug borders on layers
func (c *Overlay) SetShowDebugBordersWithParams(v *OverlaySetShowDebugBordersParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowDebugBorders", Params: v})
}

// SetShowDebugBorders - Requests that backend shows debug borders on layers
// show - True for showing debug borders
func (c *Overlay) SetShowDebugBorders(show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowDebugBordersParams
	v.Show = show
	return c.SetShowDebugBordersWithParams(&v)
}

type OverlaySetShowFPSCounterParams struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// SetShowFPSCounterWithParams - Requests that backend shows the FPS counter
func (c *Overlay) SetShowFPSCounterWithParams(v *OverlaySetShowFPSCounterParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowFPSCounter", Params: v})
}

// SetShowFPSCounter - Requests that backend shows the FPS counter
// show - True for showing the FPS counter
func (c *Overlay) SetShowFPSCounter(show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowFPSCounterParams
	v.Show = show
	return c.SetShowFPSCounterWithParams(&v)
}

type OverlaySetShowScrollBottleneckRectsParams struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// SetShowScrollBottleneckRectsWithParams - Requests that backend shows scroll bottleneck rects
func (c *Overlay) SetShowScrollBottleneckRectsWithParams(v *OverlaySetShowScrollBottleneckRectsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowScrollBottleneckRects", Params: v})
}

// SetShowScrollBottleneckRects - Requests that backend shows scroll bottleneck rects
// show - True for showing scroll bottleneck rects
func (c *Overlay) SetShowScrollBottleneckRects(show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowScrollBottleneckRectsParams
	v.Show = show
	return c.SetShowScrollBottleneckRectsWithParams(&v)
}

type OverlaySetShowViewportSizeOnResizeParams struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// SetShowViewportSizeOnResizeWithParams - Paints viewport size upon main frame resize.
func (c *Overlay) SetShowViewportSizeOnResizeWithParams(v *OverlaySetShowViewportSizeOnResizeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setShowViewportSizeOnResize", Params: v})
}

// SetShowViewportSizeOnResize - Paints viewport size upon main frame resize.
// show - Whether to paint size or not.
func (c *Overlay) SetShowViewportSizeOnResize(show bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetShowViewportSizeOnResizeParams
	v.Show = show
	return c.SetShowViewportSizeOnResizeWithParams(&v)
}

type OverlaySetPausedInDebuggerMessageParams struct {
	// The message to display, also triggers resume and step over controls.
	Message string `json:"message,omitempty"`
}

// SetPausedInDebuggerMessageWithParams -
func (c *Overlay) SetPausedInDebuggerMessageWithParams(v *OverlaySetPausedInDebuggerMessageParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setPausedInDebuggerMessage", Params: v})
}

// SetPausedInDebuggerMessage -
// message - The message to display, also triggers resume and step over controls.
func (c *Overlay) SetPausedInDebuggerMessage(message string) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetPausedInDebuggerMessageParams
	v.Message = message
	return c.SetPausedInDebuggerMessageWithParams(&v)
}

type OverlaySetSuspendedParams struct {
	// Whether overlay should be suspended and not consume any resources until resumed.
	Suspended bool `json:"suspended"`
}

// SetSuspendedWithParams -
func (c *Overlay) SetSuspendedWithParams(v *OverlaySetSuspendedParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setSuspended", Params: v})
}

// SetSuspended -
// suspended - Whether overlay should be suspended and not consume any resources until resumed.
func (c *Overlay) SetSuspended(suspended bool) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetSuspendedParams
	v.Suspended = suspended
	return c.SetSuspendedWithParams(&v)
}

type OverlaySetInspectModeParams struct {
	// Set an inspection mode. enum values: searchForNode, searchForUAShadowDOM, none
	Mode string `json:"mode"`
	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
	HighlightConfig *OverlayHighlightConfig `json:"highlightConfig,omitempty"`
}

// SetInspectModeWithParams - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
func (c *Overlay) SetInspectModeWithParams(v *OverlaySetInspectModeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.setInspectMode", Params: v})
}

// SetInspectMode - Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
// mode - Set an inspection mode. enum values: searchForNode, searchForUAShadowDOM, none
// highlightConfig - A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
func (c *Overlay) SetInspectMode(mode string, highlightConfig *OverlayHighlightConfig) (*gcdmessage.ChromeResponse, error) {
	var v OverlaySetInspectModeParams
	v.Mode = mode
	v.HighlightConfig = highlightConfig
	return c.SetInspectModeWithParams(&v)
}

type OverlayHighlightRectParams struct {
	// X coordinate
	X int `json:"x"`
	// Y coordinate
	Y int `json:"y"`
	// Rectangle width
	Width int `json:"width"`
	// Rectangle height
	Height int `json:"height"`
	// The highlight fill color (default: transparent).
	Color *DOMRGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// HighlightRectWithParams - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
func (c *Overlay) HighlightRectWithParams(v *OverlayHighlightRectParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightRect", Params: v})
}

// HighlightRect - Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
// x - X coordinate
// y - Y coordinate
// width - Rectangle width
// height - Rectangle height
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *Overlay) HighlightRect(x int, y int, width int, height int, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightRectParams
	v.X = x
	v.Y = y
	v.Width = width
	v.Height = height
	v.Color = color
	v.OutlineColor = outlineColor
	return c.HighlightRectWithParams(&v)
}

type OverlayHighlightQuadParams struct {
	// Quad to highlight
	Quad []float64 `json:"quad"`
	// The highlight fill color (default: transparent).
	Color *DOMRGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *DOMRGBA `json:"outlineColor,omitempty"`
}

// HighlightQuadWithParams - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
func (c *Overlay) HighlightQuadWithParams(v *OverlayHighlightQuadParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightQuad", Params: v})
}

// HighlightQuad - Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
// quad - Quad to highlight
// color - The highlight fill color (default: transparent).
// outlineColor - The highlight outline color (default: transparent).
func (c *Overlay) HighlightQuad(quad []float64, color *DOMRGBA, outlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightQuadParams
	v.Quad = quad
	v.Color = color
	v.OutlineColor = outlineColor
	return c.HighlightQuadWithParams(&v)
}

type OverlayHighlightNodeParams struct {
	// A descriptor for the highlight appearance.
	HighlightConfig *OverlayHighlightConfig `json:"highlightConfig"`
	// Identifier of the node to highlight.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectId string `json:"objectId,omitempty"`
}

// HighlightNodeWithParams - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (c *Overlay) HighlightNodeWithParams(v *OverlayHighlightNodeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightNode", Params: v})
}

// HighlightNode - Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
// highlightConfig - A descriptor for the highlight appearance.
// nodeId - Identifier of the node to highlight.
// backendNodeId - Identifier of the backend node to highlight.
// objectId - JavaScript object id of the node to be highlighted.
func (c *Overlay) HighlightNode(highlightConfig *OverlayHighlightConfig, nodeId int, backendNodeId int, objectId string) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightNodeParams
	v.HighlightConfig = highlightConfig
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.HighlightNodeWithParams(&v)
}

type OverlayHighlightFrameParams struct {
	// Identifier of the frame to highlight.
	FrameId string `json:"frameId"`
	// The content box highlight fill color (default: transparent).
	ContentColor *DOMRGBA `json:"contentColor,omitempty"`
	// The content box highlight outline color (default: transparent).
	ContentOutlineColor *DOMRGBA `json:"contentOutlineColor,omitempty"`
}

// HighlightFrameWithParams - Highlights owner element of the frame with given id.
func (c *Overlay) HighlightFrameWithParams(v *OverlayHighlightFrameParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.highlightFrame", Params: v})
}

// HighlightFrame - Highlights owner element of the frame with given id.
// frameId - Identifier of the frame to highlight.
// contentColor - The content box highlight fill color (default: transparent).
// contentOutlineColor - The content box highlight outline color (default: transparent).
func (c *Overlay) HighlightFrame(frameId string, contentColor *DOMRGBA, contentOutlineColor *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v OverlayHighlightFrameParams
	v.FrameId = frameId
	v.ContentColor = contentColor
	v.ContentOutlineColor = contentOutlineColor
	return c.HighlightFrameWithParams(&v)
}

// Hides any highlight.
func (c *Overlay) HideHighlight() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.hideHighlight"})
}

type OverlayGetHighlightObjectForTestParams struct {
	// Id of the node to get highlight object for.
	NodeId int `json:"nodeId"`
}

// GetHighlightObjectForTestWithParams - For testing.
// Returns -  highlight - Highlight data for the node.
func (c *Overlay) GetHighlightObjectForTestWithParams(v *OverlayGetHighlightObjectForTestParams) (map[string]interface{}, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Overlay.getHighlightObjectForTest", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Highlight map[string]interface{}
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Highlight, nil
}

// GetHighlightObjectForTest - For testing.
// nodeId - Id of the node to get highlight object for.
// Returns -  highlight - Highlight data for the node.
func (c *Overlay) GetHighlightObjectForTest(nodeId int) (map[string]interface{}, error) {
	var v OverlayGetHighlightObjectForTestParams
	v.NodeId = nodeId
	return c.GetHighlightObjectForTestWithParams(&v)
}
