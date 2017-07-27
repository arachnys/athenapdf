// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Page functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Information about the Frame on the page.
type PageFrame struct {
	Id             string `json:"id"`                 // Frame unique identifier.
	ParentId       string `json:"parentId,omitempty"` // Parent frame identifier.
	LoaderId       string `json:"loaderId"`           // Identifier of the loader associated with this frame.
	Name           string `json:"name,omitempty"`     // Frame's name as specified in the tag.
	Url            string `json:"url"`                // Frame document's URL.
	SecurityOrigin string `json:"securityOrigin"`     // Frame document's security origin.
	MimeType       string `json:"mimeType"`           // Frame document's mimeType as determined by the browser.
}

// Information about the Resource on the page.
type PageFrameResource struct {
	Url          string  `json:"url"`                    // Resource URL.
	Type         string  `json:"type"`                   // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
	MimeType     string  `json:"mimeType"`               // Resource mimeType as determined by the browser.
	LastModified float64 `json:"lastModified,omitempty"` // last-modified timestamp as reported by server.
	ContentSize  float64 `json:"contentSize,omitempty"`  // Resource content size.
	Failed       bool    `json:"failed,omitempty"`       // True if the resource failed to load.
	Canceled     bool    `json:"canceled,omitempty"`     // True if the resource was canceled during loading.
}

// Information about the Frame hierarchy along with their cached resources.
type PageFrameResourceTree struct {
	Frame       *PageFrame               `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []*PageFrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []*PageFrameResource     `json:"resources"`             // Information about frame resources.
}

// Navigation history entry.
type PageNavigationEntry struct {
	Id             int    `json:"id"`             // Unique id of the navigation history entry.
	Url            string `json:"url"`            // URL of the navigation history entry.
	UserTypedURL   string `json:"userTypedURL"`   // URL that the user typed in the url bar.
	Title          string `json:"title"`          // Title of the navigation history entry.
	TransitionType string `json:"transitionType"` // Transition type. enum values: link, typed, auto_bookmark, auto_subframe, manual_subframe, generated, auto_toplevel, form_submit, reload, keyword, keyword_generated, other
}

// Screencast frame metadata.
type PageScreencastFrameMetadata struct {
	OffsetTop       float64 `json:"offsetTop"`           // Top offset in DIP.
	PageScaleFactor float64 `json:"pageScaleFactor"`     // Page scale factor.
	DeviceWidth     float64 `json:"deviceWidth"`         // Device screen width in DIP.
	DeviceHeight    float64 `json:"deviceHeight"`        // Device screen height in DIP.
	ScrollOffsetX   float64 `json:"scrollOffsetX"`       // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY   float64 `json:"scrollOffsetY"`       // Position of vertical scroll in CSS pixels.
	Timestamp       float64 `json:"timestamp,omitempty"` // Frame swap timestamp.
}

// Error while paring app manifest.
type PageAppManifestError struct {
	Message  string `json:"message"`  // Error message.
	Critical int    `json:"critical"` // If criticial, this is a non-recoverable parse error.
	Line     int    `json:"line"`     // Error line.
	Column   int    `json:"column"`   // Error column.
}

// Layout viewport position and dimensions.
type PageLayoutViewport struct {
	PageX        int `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        int `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  int `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
}

// Visual viewport position, dimensions, and scale.
type PageVisualViewport struct {
	OffsetX      float64 `json:"offsetX"`      // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`      // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`        // Scale relative to the ideal viewport (size at width=device-width).
}

//
type PageDomContentEventFiredEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

//
type PageLoadEventFiredEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp float64 `json:"timestamp"` //
	} `json:"Params,omitempty"`
}

// Fired when frame has been attached to its parent.
type PageFrameAttachedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId       string             `json:"frameId"`         // Id of the frame that has been attached.
		ParentFrameId string             `json:"parentFrameId"`   // Parent frame identifier.
		Stack         *RuntimeStackTrace `json:"stack,omitempty"` // JavaScript stack trace of when frame was attached, only set if frame initiated from script.
	} `json:"Params,omitempty"`
}

// Fired once navigation of the frame has completed. Frame is now associated with the new loader.
type PageFrameNavigatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Frame *PageFrame `json:"frame"` // Frame object.
	} `json:"Params,omitempty"`
}

// Fired when frame has been detached from its parent.
type PageFrameDetachedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has been detached.
	} `json:"Params,omitempty"`
}

// Fired when frame has started loading.
type PageFrameStartedLoadingEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has started loading.
	} `json:"Params,omitempty"`
}

// Fired when frame has stopped loading.
type PageFrameStoppedLoadingEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has stopped loading.
	} `json:"Params,omitempty"`
}

// Fired when frame schedules a potential navigation.
type PageFrameScheduledNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string  `json:"frameId"` // Id of the frame that has scheduled a navigation.
		Delay   float64 `json:"delay"`   // Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
	} `json:"Params,omitempty"`
}

// Fired when frame no longer has a scheduled navigation.
type PageFrameClearedScheduledNavigationEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId string `json:"frameId"` // Id of the frame that has cleared its scheduled navigation.
	} `json:"Params,omitempty"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to open.
type PageJavascriptDialogOpeningEvent struct {
	Method string `json:"method"`
	Params struct {
		Message string `json:"message"` // Message that will be displayed by the dialog.
		Type    string `json:"type"`    // Dialog type. enum values: alert, confirm, prompt, beforeunload
	} `json:"Params,omitempty"`
}

// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been closed.
type PageJavascriptDialogClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		Result bool `json:"result"` // Whether dialog was confirmed.
	} `json:"Params,omitempty"`
}

// Compressed image data requested by the <code>startScreencast</code>.
type PageScreencastFrameEvent struct {
	Method string `json:"method"`
	Params struct {
		Data      string                       `json:"data"`      // Base64-encoded compressed image.
		Metadata  *PageScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
		SessionId int                          `json:"sessionId"` // Frame number.
	} `json:"Params,omitempty"`
}

// Fired when the page with currently enabled screencast was shown or hidden </code>.
type PageScreencastVisibilityChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		Visible bool `json:"visible"` // True if the page is visible.
	} `json:"Params,omitempty"`
}

// Fired when a navigation is started if navigation throttles are enabled.  The navigation will be deferred until processNavigation is called.
type PageNavigationRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		IsInMainFrame bool   `json:"isInMainFrame"` // Whether the navigation is taking place in the main frame or in a subframe.
		IsRedirect    bool   `json:"isRedirect"`    // Whether the navigation has encountered a server redirect or not.
		NavigationId  int    `json:"navigationId"`  //
		Url           string `json:"url"`           // URL of requested navigation.
	} `json:"Params,omitempty"`
}

type Page struct {
	target gcdmessage.ChromeTargeter
}

func NewPage(target gcdmessage.ChromeTargeter) *Page {
	c := &Page{target: target}
	return c
}

// Enables page domain notifications.
func (c *Page) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.enable"})
}

// Disables page domain notifications.
func (c *Page) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.disable"})
}

type PageAddScriptToEvaluateOnLoadParams struct {
	//
	ScriptSource string `json:"scriptSource"`
}

// AddScriptToEvaluateOnLoadWithParams - Deprecated, please use addScriptToEvaluateOnNewDocument instead.
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnLoadWithParams(v *PageAddScriptToEvaluateOnLoadParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.addScriptToEvaluateOnLoad", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Identifier string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Identifier, nil
}

// AddScriptToEvaluateOnLoad - Deprecated, please use addScriptToEvaluateOnNewDocument instead.
// scriptSource -
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnLoad(scriptSource string) (string, error) {
	var v PageAddScriptToEvaluateOnLoadParams
	v.ScriptSource = scriptSource
	return c.AddScriptToEvaluateOnLoadWithParams(&v)
}

type PageRemoveScriptToEvaluateOnLoadParams struct {
	//
	Identifier string `json:"identifier"`
}

// RemoveScriptToEvaluateOnLoadWithParams - Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
func (c *Page) RemoveScriptToEvaluateOnLoadWithParams(v *PageRemoveScriptToEvaluateOnLoadParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.removeScriptToEvaluateOnLoad", Params: v})
}

// RemoveScriptToEvaluateOnLoad - Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
// identifier -
func (c *Page) RemoveScriptToEvaluateOnLoad(identifier string) (*gcdmessage.ChromeResponse, error) {
	var v PageRemoveScriptToEvaluateOnLoadParams
	v.Identifier = identifier
	return c.RemoveScriptToEvaluateOnLoadWithParams(&v)
}

type PageAddScriptToEvaluateOnNewDocumentParams struct {
	//
	Source string `json:"source"`
}

// AddScriptToEvaluateOnNewDocumentWithParams - Evaluates given script in every frame upon creation (before loading frame's scripts).
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnNewDocumentWithParams(v *PageAddScriptToEvaluateOnNewDocumentParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.addScriptToEvaluateOnNewDocument", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Identifier string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Identifier, nil
}

// AddScriptToEvaluateOnNewDocument - Evaluates given script in every frame upon creation (before loading frame's scripts).
// source -
// Returns -  identifier - Identifier of the added script.
func (c *Page) AddScriptToEvaluateOnNewDocument(source string) (string, error) {
	var v PageAddScriptToEvaluateOnNewDocumentParams
	v.Source = source
	return c.AddScriptToEvaluateOnNewDocumentWithParams(&v)
}

type PageRemoveScriptToEvaluateOnNewDocumentParams struct {
	//
	Identifier string `json:"identifier"`
}

// RemoveScriptToEvaluateOnNewDocumentWithParams - Removes given script from the list.
func (c *Page) RemoveScriptToEvaluateOnNewDocumentWithParams(v *PageRemoveScriptToEvaluateOnNewDocumentParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.removeScriptToEvaluateOnNewDocument", Params: v})
}

// RemoveScriptToEvaluateOnNewDocument - Removes given script from the list.
// identifier -
func (c *Page) RemoveScriptToEvaluateOnNewDocument(identifier string) (*gcdmessage.ChromeResponse, error) {
	var v PageRemoveScriptToEvaluateOnNewDocumentParams
	v.Identifier = identifier
	return c.RemoveScriptToEvaluateOnNewDocumentWithParams(&v)
}

type PageSetAutoAttachToCreatedPagesParams struct {
	// If true, browser will open a new inspector window for every page created from this one.
	AutoAttach bool `json:"autoAttach"`
}

// SetAutoAttachToCreatedPagesWithParams - Controls whether browser will open a new inspector window for connected pages.
func (c *Page) SetAutoAttachToCreatedPagesWithParams(v *PageSetAutoAttachToCreatedPagesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setAutoAttachToCreatedPages", Params: v})
}

// SetAutoAttachToCreatedPages - Controls whether browser will open a new inspector window for connected pages.
// autoAttach - If true, browser will open a new inspector window for every page created from this one.
func (c *Page) SetAutoAttachToCreatedPages(autoAttach bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetAutoAttachToCreatedPagesParams
	v.AutoAttach = autoAttach
	return c.SetAutoAttachToCreatedPagesWithParams(&v)
}

type PageReloadParams struct {
	// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache bool `json:"ignoreCache,omitempty"`
	// If set, the script will be injected into all frames of the inspected page after reload.
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

// ReloadWithParams - Reloads given page optionally ignoring the cache.
func (c *Page) ReloadWithParams(v *PageReloadParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.reload", Params: v})
}

// Reload - Reloads given page optionally ignoring the cache.
// ignoreCache - If true, browser cache is ignored (as if the user pressed Shift+refresh).
// scriptToEvaluateOnLoad - If set, the script will be injected into all frames of the inspected page after reload.
func (c *Page) Reload(ignoreCache bool, scriptToEvaluateOnLoad string) (*gcdmessage.ChromeResponse, error) {
	var v PageReloadParams
	v.IgnoreCache = ignoreCache
	v.ScriptToEvaluateOnLoad = scriptToEvaluateOnLoad
	return c.ReloadWithParams(&v)
}

type PageNavigateParams struct {
	// URL to navigate the page to.
	Url string `json:"url"`
	// Referrer URL.
	Referrer string `json:"referrer,omitempty"`
	// Intended transition type. enum values: link, typed, auto_bookmark, auto_subframe, manual_subframe, generated, auto_toplevel, form_submit, reload, keyword, keyword_generated, other
	TransitionType string `json:"transitionType,omitempty"`
}

// NavigateWithParams - Navigates current page to the given URL.
// Returns -  frameId - Frame id that will be navigated.
func (c *Page) NavigateWithParams(v *PageNavigateParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.navigate", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			FrameId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.FrameId, nil
}

// Navigate - Navigates current page to the given URL.
// url - URL to navigate the page to.
// referrer - Referrer URL.
// transitionType - Intended transition type. enum values: link, typed, auto_bookmark, auto_subframe, manual_subframe, generated, auto_toplevel, form_submit, reload, keyword, keyword_generated, other
// Returns -  frameId - Frame id that will be navigated.
func (c *Page) Navigate(url string, referrer string, transitionType string) (string, error) {
	var v PageNavigateParams
	v.Url = url
	v.Referrer = referrer
	v.TransitionType = transitionType
	return c.NavigateWithParams(&v)
}

// Force the page stop all navigations and pending resource fetches.
func (c *Page) StopLoading() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.stopLoading"})
}

// GetNavigationHistory - Returns navigation history for the current page.
// Returns -  currentIndex - Index of the current navigation history entry. entries - Array of navigation history entries.
func (c *Page) GetNavigationHistory() (int, []*PageNavigationEntry, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getNavigationHistory"})
	if err != nil {
		return 0, nil, err
	}

	var chromeData struct {
		Result struct {
			CurrentIndex int
			Entries      []*PageNavigationEntry
		}
	}

	if resp == nil {
		return 0, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, nil, err
	}

	return chromeData.Result.CurrentIndex, chromeData.Result.Entries, nil
}

type PageNavigateToHistoryEntryParams struct {
	// Unique id of the entry to navigate to.
	EntryId int `json:"entryId"`
}

// NavigateToHistoryEntryWithParams - Navigates current page to the given history entry.
func (c *Page) NavigateToHistoryEntryWithParams(v *PageNavigateToHistoryEntryParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.navigateToHistoryEntry", Params: v})
}

// NavigateToHistoryEntry - Navigates current page to the given history entry.
// entryId - Unique id of the entry to navigate to.
func (c *Page) NavigateToHistoryEntry(entryId int) (*gcdmessage.ChromeResponse, error) {
	var v PageNavigateToHistoryEntryParams
	v.EntryId = entryId
	return c.NavigateToHistoryEntryWithParams(&v)
}

// GetCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -  cookies - Array of cookie objects.
func (c *Page) GetCookies() ([]*NetworkCookie, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getCookies"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
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

	return chromeData.Result.Cookies, nil
}

type PageDeleteCookieParams struct {
	// Name of the cookie to remove.
	CookieName string `json:"cookieName"`
	// URL to match cooke domain and path.
	Url string `json:"url"`
}

// DeleteCookieWithParams - Deletes browser cookie with given name, domain and path.
func (c *Page) DeleteCookieWithParams(v *PageDeleteCookieParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.deleteCookie", Params: v})
}

// DeleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *Page) DeleteCookie(cookieName string, url string) (*gcdmessage.ChromeResponse, error) {
	var v PageDeleteCookieParams
	v.CookieName = cookieName
	v.Url = url
	return c.DeleteCookieWithParams(&v)
}

// GetResourceTree - Returns present frame / resource tree structure.
// Returns -  frameTree - Present frame / resource tree structure.
func (c *Page) GetResourceTree() (*PageFrameResourceTree, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getResourceTree"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			FrameTree *PageFrameResourceTree
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

	return chromeData.Result.FrameTree, nil
}

type PageGetResourceContentParams struct {
	// Frame id to get resource for.
	FrameId string `json:"frameId"`
	// URL of the resource to get content for.
	Url string `json:"url"`
}

// GetResourceContentWithParams - Returns content of the given resource.
// Returns -  content - Resource content. base64Encoded - True, if content was served as base64.
func (c *Page) GetResourceContentWithParams(v *PageGetResourceContentParams) (string, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getResourceContent", Params: v})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			Content       string
			Base64Encoded bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	return chromeData.Result.Content, chromeData.Result.Base64Encoded, nil
}

// GetResourceContent - Returns content of the given resource.
// frameId - Frame id to get resource for.
// url - URL of the resource to get content for.
// Returns -  content - Resource content. base64Encoded - True, if content was served as base64.
func (c *Page) GetResourceContent(frameId string, url string) (string, bool, error) {
	var v PageGetResourceContentParams
	v.FrameId = frameId
	v.Url = url
	return c.GetResourceContentWithParams(&v)
}

type PageSearchInResourceParams struct {
	// Frame id for resource to search in.
	FrameId string `json:"frameId"`
	// URL of the resource to search in.
	Url string `json:"url"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// SearchInResourceWithParams - Searches for given string in resource content.
// Returns -  result - List of search matches.
func (c *Page) SearchInResourceWithParams(v *PageSearchInResourceParams) ([]*DebuggerSearchMatch, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.searchInResource", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*DebuggerSearchMatch
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

	return chromeData.Result.Result, nil
}

// SearchInResource - Searches for given string in resource content.
// frameId - Frame id for resource to search in.
// url - URL of the resource to search in.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Page) SearchInResource(frameId string, url string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	var v PageSearchInResourceParams
	v.FrameId = frameId
	v.Url = url
	v.Query = query
	v.CaseSensitive = caseSensitive
	v.IsRegex = isRegex
	return c.SearchInResourceWithParams(&v)
}

type PageSetDocumentContentParams struct {
	// Frame id to set HTML for.
	FrameId string `json:"frameId"`
	// HTML content to set.
	Html string `json:"html"`
}

// SetDocumentContentWithParams - Sets given markup as the document's HTML.
func (c *Page) SetDocumentContentWithParams(v *PageSetDocumentContentParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDocumentContent", Params: v})
}

// SetDocumentContent - Sets given markup as the document's HTML.
// frameId - Frame id to set HTML for.
// html - HTML content to set.
func (c *Page) SetDocumentContent(frameId string, html string) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDocumentContentParams
	v.FrameId = frameId
	v.Html = html
	return c.SetDocumentContentWithParams(&v)
}

type PageSetDeviceMetricsOverrideParams struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`
	// Whether a view that exceeds the available browser window area should be scaled down to fit.
	FitWindow bool `json:"fitWindow"`
	// Scale to apply to resulting view image. Ignored in |fitWindow| mode.
	Scale float64 `json:"scale,omitempty"`
	// X offset to shift resulting view image by. Ignored in |fitWindow| mode.
	OffsetX float64 `json:"offsetX,omitempty"`
	// Y offset to shift resulting view image by. Ignored in |fitWindow| mode.
	OffsetY float64 `json:"offsetY,omitempty"`
	// Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenWidth int `json:"screenWidth,omitempty"`
	// Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenHeight int `json:"screenHeight,omitempty"`
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionX int `json:"positionX,omitempty"`
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionY int `json:"positionY,omitempty"`
	// Screen orientation override.
	ScreenOrientation *EmulationScreenOrientation `json:"screenOrientation,omitempty"`
}

// SetDeviceMetricsOverrideWithParams - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (c *Page) SetDeviceMetricsOverrideWithParams(v *PageSetDeviceMetricsOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDeviceMetricsOverride", Params: v})
}

// SetDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
// width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
// mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
// fitWindow - Whether a view that exceeds the available browser window area should be scaled down to fit.
// scale - Scale to apply to resulting view image. Ignored in |fitWindow| mode.
// offsetX - X offset to shift resulting view image by. Ignored in |fitWindow| mode.
// offsetY - Y offset to shift resulting view image by. Ignored in |fitWindow| mode.
// screenWidth - Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenHeight - Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionX - Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionY - Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenOrientation - Screen orientation override.
func (c *Page) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64, screenWidth int, screenHeight int, positionX int, positionY int, screenOrientation *EmulationScreenOrientation) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDeviceMetricsOverrideParams
	v.Width = width
	v.Height = height
	v.DeviceScaleFactor = deviceScaleFactor
	v.Mobile = mobile
	v.FitWindow = fitWindow
	v.Scale = scale
	v.OffsetX = offsetX
	v.OffsetY = offsetY
	v.ScreenWidth = screenWidth
	v.ScreenHeight = screenHeight
	v.PositionX = positionX
	v.PositionY = positionY
	v.ScreenOrientation = screenOrientation
	return c.SetDeviceMetricsOverrideWithParams(&v)
}

// Clears the overriden device metrics.
func (c *Page) ClearDeviceMetricsOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearDeviceMetricsOverride"})
}

type PageSetGeolocationOverrideParams struct {
	// Mock latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy float64 `json:"accuracy,omitempty"`
}

// SetGeolocationOverrideWithParams - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (c *Page) SetGeolocationOverrideWithParams(v *PageSetGeolocationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setGeolocationOverride", Params: v})
}

// SetGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *Page) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*gcdmessage.ChromeResponse, error) {
	var v PageSetGeolocationOverrideParams
	v.Latitude = latitude
	v.Longitude = longitude
	v.Accuracy = accuracy
	return c.SetGeolocationOverrideWithParams(&v)
}

// Clears the overriden Geolocation Position and Error.
func (c *Page) ClearGeolocationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearGeolocationOverride"})
}

type PageSetDeviceOrientationOverrideParams struct {
	// Mock alpha
	Alpha float64 `json:"alpha"`
	// Mock beta
	Beta float64 `json:"beta"`
	// Mock gamma
	Gamma float64 `json:"gamma"`
}

// SetDeviceOrientationOverrideWithParams - Overrides the Device Orientation.
func (c *Page) SetDeviceOrientationOverrideWithParams(v *PageSetDeviceOrientationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setDeviceOrientationOverride", Params: v})
}

// SetDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *Page) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*gcdmessage.ChromeResponse, error) {
	var v PageSetDeviceOrientationOverrideParams
	v.Alpha = alpha
	v.Beta = beta
	v.Gamma = gamma
	return c.SetDeviceOrientationOverrideWithParams(&v)
}

// Clears the overridden Device Orientation.
func (c *Page) ClearDeviceOrientationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.clearDeviceOrientationOverride"})
}

type PageSetTouchEmulationEnabledParams struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration string `json:"configuration,omitempty"`
}

// SetTouchEmulationEnabledWithParams - Toggles mouse event-based touch event emulation.
func (c *Page) SetTouchEmulationEnabledWithParams(v *PageSetTouchEmulationEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setTouchEmulationEnabled", Params: v})
}

// SetTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *Page) SetTouchEmulationEnabled(enabled bool, configuration string) (*gcdmessage.ChromeResponse, error) {
	var v PageSetTouchEmulationEnabledParams
	v.Enabled = enabled
	v.Configuration = configuration
	return c.SetTouchEmulationEnabledWithParams(&v)
}

type PageCaptureScreenshotParams struct {
	// Image compression format (defaults to png).
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100] (jpeg only).
	Quality int `json:"quality,omitempty"`
	// Capture the screenshot from the surface, rather than the view. Defaults to true.
	FromSurface bool `json:"fromSurface,omitempty"`
}

// CaptureScreenshotWithParams - Capture page screenshot.
// Returns -  data - Base64-encoded image data.
func (c *Page) CaptureScreenshotWithParams(v *PageCaptureScreenshotParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.captureScreenshot", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Data string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Data, nil
}

// CaptureScreenshot - Capture page screenshot.
// format - Image compression format (defaults to png).
// quality - Compression quality from range [0..100] (jpeg only).
// fromSurface - Capture the screenshot from the surface, rather than the view. Defaults to true.
// Returns -  data - Base64-encoded image data.
func (c *Page) CaptureScreenshot(format string, quality int, fromSurface bool) (string, error) {
	var v PageCaptureScreenshotParams
	v.Format = format
	v.Quality = quality
	v.FromSurface = fromSurface
	return c.CaptureScreenshotWithParams(&v)
}

type PagePrintToPDFParams struct {
	// Paper orientation. Defaults to false.
	Landscape bool `json:"landscape,omitempty"`
	// Display header and footer. Defaults to false.
	DisplayHeaderFooter bool `json:"displayHeaderFooter,omitempty"`
	// Print background graphics. Defaults to false.
	PrintBackground bool `json:"printBackground,omitempty"`
	// Scale of the webpage rendering. Defaults to 1.
	Scale float64 `json:"scale,omitempty"`
	// Paper width in inches. Defaults to 8.5 inches.
	PaperWidth float64 `json:"paperWidth,omitempty"`
	// Paper height in inches. Defaults to 11 inches.
	PaperHeight float64 `json:"paperHeight,omitempty"`
	// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop float64 `json:"marginTop,omitempty"`
	// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom float64 `json:"marginBottom,omitempty"`
	// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft float64 `json:"marginLeft,omitempty"`
	// Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight float64 `json:"marginRight,omitempty"`
	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	PageRanges string `json:"pageRanges,omitempty"`
}

// PrintToPDFWithParams - Print page as PDF.
// Returns -  data - Base64-encoded pdf data.
func (c *Page) PrintToPDFWithParams(v *PagePrintToPDFParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.printToPDF", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Data string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Data, nil
}

// PrintToPDF - Print page as PDF.
// landscape - Paper orientation. Defaults to false.
// displayHeaderFooter - Display header and footer. Defaults to false.
// printBackground - Print background graphics. Defaults to false.
// scale - Scale of the webpage rendering. Defaults to 1.
// paperWidth - Paper width in inches. Defaults to 8.5 inches.
// paperHeight - Paper height in inches. Defaults to 11 inches.
// marginTop - Top margin in inches. Defaults to 1cm (~0.4 inches).
// marginBottom - Bottom margin in inches. Defaults to 1cm (~0.4 inches).
// marginLeft - Left margin in inches. Defaults to 1cm (~0.4 inches).
// marginRight - Right margin in inches. Defaults to 1cm (~0.4 inches).
// pageRanges - Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
// Returns -  data - Base64-encoded pdf data.
func (c *Page) PrintToPDF(landscape bool, displayHeaderFooter bool, printBackground bool, scale float64, paperWidth float64, paperHeight float64, marginTop float64, marginBottom float64, marginLeft float64, marginRight float64, pageRanges string) (string, error) {
	var v PagePrintToPDFParams
	v.Landscape = landscape
	v.DisplayHeaderFooter = displayHeaderFooter
	v.PrintBackground = printBackground
	v.Scale = scale
	v.PaperWidth = paperWidth
	v.PaperHeight = paperHeight
	v.MarginTop = marginTop
	v.MarginBottom = marginBottom
	v.MarginLeft = marginLeft
	v.MarginRight = marginRight
	v.PageRanges = pageRanges
	return c.PrintToPDFWithParams(&v)
}

type PageStartScreencastParams struct {
	// Image compression format.
	Format string `json:"format,omitempty"`
	// Compression quality from range [0..100].
	Quality int `json:"quality,omitempty"`
	// Maximum screenshot width.
	MaxWidth int `json:"maxWidth,omitempty"`
	// Maximum screenshot height.
	MaxHeight int `json:"maxHeight,omitempty"`
	// Send every n-th frame.
	EveryNthFrame int `json:"everyNthFrame,omitempty"`
}

// StartScreencastWithParams - Starts sending each frame using the <code>screencastFrame</code> event.
func (c *Page) StartScreencastWithParams(v *PageStartScreencastParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.startScreencast", Params: v})
}

// StartScreencast - Starts sending each frame using the <code>screencastFrame</code> event.
// format - Image compression format.
// quality - Compression quality from range [0..100].
// maxWidth - Maximum screenshot width.
// maxHeight - Maximum screenshot height.
// everyNthFrame - Send every n-th frame.
func (c *Page) StartScreencast(format string, quality int, maxWidth int, maxHeight int, everyNthFrame int) (*gcdmessage.ChromeResponse, error) {
	var v PageStartScreencastParams
	v.Format = format
	v.Quality = quality
	v.MaxWidth = maxWidth
	v.MaxHeight = maxHeight
	v.EveryNthFrame = everyNthFrame
	return c.StartScreencastWithParams(&v)
}

// Stops sending each frame in the <code>screencastFrame</code>.
func (c *Page) StopScreencast() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.stopScreencast"})
}

type PageScreencastFrameAckParams struct {
	// Frame number.
	SessionId int `json:"sessionId"`
}

// ScreencastFrameAckWithParams - Acknowledges that a screencast frame has been received by the frontend.
func (c *Page) ScreencastFrameAckWithParams(v *PageScreencastFrameAckParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.screencastFrameAck", Params: v})
}

// ScreencastFrameAck - Acknowledges that a screencast frame has been received by the frontend.
// sessionId - Frame number.
func (c *Page) ScreencastFrameAck(sessionId int) (*gcdmessage.ChromeResponse, error) {
	var v PageScreencastFrameAckParams
	v.SessionId = sessionId
	return c.ScreencastFrameAckWithParams(&v)
}

type PageHandleJavaScriptDialogParams struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`
	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
	PromptText string `json:"promptText,omitempty"`
}

// HandleJavaScriptDialogWithParams - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
func (c *Page) HandleJavaScriptDialogWithParams(v *PageHandleJavaScriptDialogParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.handleJavaScriptDialog", Params: v})
}

// HandleJavaScriptDialog - Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
// accept - Whether to accept or dismiss the dialog.
// promptText - The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
func (c *Page) HandleJavaScriptDialog(accept bool, promptText string) (*gcdmessage.ChromeResponse, error) {
	var v PageHandleJavaScriptDialogParams
	v.Accept = accept
	v.PromptText = promptText
	return c.HandleJavaScriptDialogWithParams(&v)
}

// GetAppManifest -
// Returns -  url - Manifest location. errors -  data - Manifest content.
func (c *Page) GetAppManifest() (string, []*PageAppManifestError, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getAppManifest"})
	if err != nil {
		return "", nil, "", err
	}

	var chromeData struct {
		Result struct {
			Url    string
			Errors []*PageAppManifestError
			Data   string
		}
	}

	if resp == nil {
		return "", nil, "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, "", err
	}

	return chromeData.Result.Url, chromeData.Result.Errors, chromeData.Result.Data, nil
}

//
func (c *Page) RequestAppBanner() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.requestAppBanner"})
}

type PageSetControlNavigationsParams struct {
	//
	Enabled bool `json:"enabled"`
}

// SetControlNavigationsWithParams - Toggles navigation throttling which allows programatic control over navigation and redirect response.
func (c *Page) SetControlNavigationsWithParams(v *PageSetControlNavigationsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.setControlNavigations", Params: v})
}

// SetControlNavigations - Toggles navigation throttling which allows programatic control over navigation and redirect response.
// enabled -
func (c *Page) SetControlNavigations(enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v PageSetControlNavigationsParams
	v.Enabled = enabled
	return c.SetControlNavigationsWithParams(&v)
}

type PageProcessNavigationParams struct {
	//  enum values: Proceed, Cancel, CancelAndIgnore
	Response string `json:"response"`
	//
	NavigationId int `json:"navigationId"`
}

// ProcessNavigationWithParams - Should be sent in response to a navigationRequested or a redirectRequested event, telling the browser how to handle the navigation.
func (c *Page) ProcessNavigationWithParams(v *PageProcessNavigationParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.processNavigation", Params: v})
}

// ProcessNavigation - Should be sent in response to a navigationRequested or a redirectRequested event, telling the browser how to handle the navigation.
// response -  enum values: Proceed, Cancel, CancelAndIgnore
// navigationId -
func (c *Page) ProcessNavigation(response string, navigationId int) (*gcdmessage.ChromeResponse, error) {
	var v PageProcessNavigationParams
	v.Response = response
	v.NavigationId = navigationId
	return c.ProcessNavigationWithParams(&v)
}

// GetLayoutMetrics - Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
// Returns -  layoutViewport - Metrics relating to the layout viewport. visualViewport - Metrics relating to the visual viewport. contentSize - Size of scrollable area.
func (c *Page) GetLayoutMetrics() (*PageLayoutViewport, *PageVisualViewport, *DOMRect, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.getLayoutMetrics"})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			LayoutViewport *PageLayoutViewport
			VisualViewport *PageVisualViewport
			ContentSize    *DOMRect
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.LayoutViewport, chromeData.Result.VisualViewport, chromeData.Result.ContentSize, nil
}

type PageCreateIsolatedWorldParams struct {
	// Id of the frame in which the isolated world should be created.
	FrameId string `json:"frameId"`
	// An optional name which is reported in the Execution Context.
	WorldName string `json:"worldName,omitempty"`
	// Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
}

// CreateIsolatedWorldWithParams - Creates an isolated world for the given frame.
// Returns -  executionContextId - Execution context of the isolated world.
func (c *Page) CreateIsolatedWorldWithParams(v *PageCreateIsolatedWorldParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Page.createIsolatedWorld", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			ExecutionContextId int
		}
	}

	if resp == nil {
		return 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, err
	}

	return chromeData.Result.ExecutionContextId, nil
}

// CreateIsolatedWorld - Creates an isolated world for the given frame.
// frameId - Id of the frame in which the isolated world should be created.
// worldName - An optional name which is reported in the Execution Context.
// grantUniveralAccess - Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
// Returns -  executionContextId - Execution context of the isolated world.
func (c *Page) CreateIsolatedWorld(frameId string, worldName string, grantUniveralAccess bool) (int, error) {
	var v PageCreateIsolatedWorldParams
	v.FrameId = frameId
	v.WorldName = worldName
	v.GrantUniveralAccess = grantUniveralAccess
	return c.CreateIsolatedWorldWithParams(&v)
}
