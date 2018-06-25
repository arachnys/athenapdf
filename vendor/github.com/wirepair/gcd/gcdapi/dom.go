// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOM functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Backend node with a friendly name.
type DOMBackendNode struct {
	NodeType      int    `json:"nodeType"`      // `Node`'s nodeType.
	NodeName      string `json:"nodeName"`      // `Node`'s nodeName.
	BackendNodeId int    `json:"backendNodeId"` //
}

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type DOMNode struct {
	NodeId           int               `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the `nodeId`. Backend will only push node with given `id` once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	ParentId         int               `json:"parentId,omitempty"`         // The id of the parent node if any.
	BackendNodeId    int               `json:"backendNodeId"`              // The BackendNodeId for this node.
	NodeType         int               `json:"nodeType"`                   // `Node`'s nodeType.
	NodeName         string            `json:"nodeName"`                   // `Node`'s nodeName.
	LocalName        string            `json:"localName"`                  // `Node`'s localName.
	NodeValue        string            `json:"nodeValue"`                  // `Node`'s nodeValue.
	ChildNodeCount   int               `json:"childNodeCount,omitempty"`   // Child count for `Container` nodes.
	Children         []*DOMNode        `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string          `json:"attributes,omitempty"`       // Attributes of the `Element` node in the form of flat array `[name1, value1, name2, value2]`.
	DocumentURL      string            `json:"documentURL,omitempty"`      // Document URL that `Document` or `FrameOwner` node points to.
	BaseURL          string            `json:"baseURL,omitempty"`          // Base URL that `Document` or `FrameOwner` node uses for URL completion.
	PublicId         string            `json:"publicId,omitempty"`         // `DocumentType`'s publicId.
	SystemId         string            `json:"systemId,omitempty"`         // `DocumentType`'s systemId.
	InternalSubset   string            `json:"internalSubset,omitempty"`   // `DocumentType`'s internalSubset.
	XmlVersion       string            `json:"xmlVersion,omitempty"`       // `Document`'s XML version in case of XML documents.
	Name             string            `json:"name,omitempty"`             // `Attr`'s name.
	Value            string            `json:"value,omitempty"`            // `Attr`'s value.
	PseudoType       string            `json:"pseudoType,omitempty"`       // Pseudo element type for this node. enum values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button
	ShadowRootType   string            `json:"shadowRootType,omitempty"`   // Shadow root type. enum values: user-agent, open, closed
	FrameId          string            `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *DOMNode          `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []*DOMNode        `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *DOMNode          `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []*DOMNode        `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	ImportedDocument *DOMNode          `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
	DistributedNodes []*DOMBackendNode `json:"distributedNodes,omitempty"` // Distributed nodes for given insertion point.
	IsSVG            bool              `json:"isSVG,omitempty"`            // Whether the node is SVG.
}

// A structure holding an RGBA color.
type DOMRGBA struct {
	R int     `json:"r"`           // The red component, in the [0-255] range.
	G int     `json:"g"`           // The green component, in the [0-255] range.
	B int     `json:"b"`           // The blue component, in the [0-255] range.
	A float64 `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}

// Box model.
type DOMBoxModel struct {
	Content      []float64            `json:"content"`                // Content box
	Padding      []float64            `json:"padding"`                // Padding box
	Border       []float64            `json:"border"`                 // Border box
	Margin       []float64            `json:"margin"`                 // Margin box
	Width        int                  `json:"width"`                  // Node width
	Height       int                  `json:"height"`                 // Node height
	ShapeOutside *DOMShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}

// CSS Shape Outside details.
type DOMShapeOutsideInfo struct {
	Bounds      []float64     `json:"bounds"`      // Shape bounds
	Shape       []interface{} `json:"shape"`       // Shape coordinate details
	MarginShape []interface{} `json:"marginShape"` // Margin shape bounds
}

// Rectangle.
type DOMRect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}

// Fired when `Element`'s attribute is modified.
type DOMAttributeModifiedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId int    `json:"nodeId"` // Id of the node that has changed.
		Name   string `json:"name"`   // Attribute name.
		Value  string `json:"value"`  // Attribute value.
	} `json:"Params,omitempty"`
}

// Fired when `Element`'s attribute is removed.
type DOMAttributeRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId int    `json:"nodeId"` // Id of the node that has changed.
		Name   string `json:"name"`   // A ttribute name.
	} `json:"Params,omitempty"`
}

// Mirrors `DOMCharacterDataModified` event.
type DOMCharacterDataModifiedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId        int    `json:"nodeId"`        // Id of the node that has changed.
		CharacterData string `json:"characterData"` // New text value.
	} `json:"Params,omitempty"`
}

// Fired when `Container`'s child node count has changed.
type DOMChildNodeCountUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeId         int `json:"nodeId"`         // Id of the node that has changed.
		ChildNodeCount int `json:"childNodeCount"` // New node count.
	} `json:"Params,omitempty"`
}

// Mirrors `DOMNodeInserted` event.
type DOMChildNodeInsertedEvent struct {
	Method string `json:"method"`
	Params struct {
		ParentNodeId   int      `json:"parentNodeId"`   // Id of the node that has changed.
		PreviousNodeId int      `json:"previousNodeId"` // If of the previous siblint.
		Node           *DOMNode `json:"node"`           // Inserted node data.
	} `json:"Params,omitempty"`
}

// Mirrors `DOMNodeRemoved` event.
type DOMChildNodeRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		ParentNodeId int `json:"parentNodeId"` // Parent id.
		NodeId       int `json:"nodeId"`       // Id of the node that has been removed.
	} `json:"Params,omitempty"`
}

// Called when distrubution is changed.
type DOMDistributedNodesUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		InsertionPointId int               `json:"insertionPointId"` // Insertion point where distrubuted nodes were updated.
		DistributedNodes []*DOMBackendNode `json:"distributedNodes"` // Distributed nodes for given insertion point.
	} `json:"Params,omitempty"`
}

// Fired when `Element`'s inline style is modified via a CSS property modification.
type DOMInlineStyleInvalidatedEvent struct {
	Method string `json:"method"`
	Params struct {
		NodeIds []int `json:"nodeIds"` // Ids of the nodes for which the inline styles have been invalidated.
	} `json:"Params,omitempty"`
}

// Called when a pseudo element is added to an element.
type DOMPseudoElementAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		ParentId      int      `json:"parentId"`      // Pseudo element's parent element id.
		PseudoElement *DOMNode `json:"pseudoElement"` // The added pseudo element.
	} `json:"Params,omitempty"`
}

// Called when a pseudo element is removed from an element.
type DOMPseudoElementRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		ParentId        int `json:"parentId"`        // Pseudo element's parent element id.
		PseudoElementId int `json:"pseudoElementId"` // The removed pseudo element id.
	} `json:"Params,omitempty"`
}

// Fired when backend wants to provide client with the missing DOM structure. This happens upon most of the calls requesting node ids.
type DOMSetChildNodesEvent struct {
	Method string `json:"method"`
	Params struct {
		ParentId int        `json:"parentId"` // Parent node id to populate with children.
		Nodes    []*DOMNode `json:"nodes"`    // Child nodes array.
	} `json:"Params,omitempty"`
}

// Called when shadow root is popped from the element.
type DOMShadowRootPoppedEvent struct {
	Method string `json:"method"`
	Params struct {
		HostId int `json:"hostId"` // Host element id.
		RootId int `json:"rootId"` // Shadow root id.
	} `json:"Params,omitempty"`
}

// Called when shadow root is pushed into the element.
type DOMShadowRootPushedEvent struct {
	Method string `json:"method"`
	Params struct {
		HostId int      `json:"hostId"` // Host element id.
		Root   *DOMNode `json:"root"`   // Shadow root.
	} `json:"Params,omitempty"`
}

type DOM struct {
	target gcdmessage.ChromeTargeter
}

func NewDOM(target gcdmessage.ChromeTargeter) *DOM {
	c := &DOM{target: target}
	return c
}

type DOMCollectClassNamesFromSubtreeParams struct {
	// Id of the node to collect class names.
	NodeId int `json:"nodeId"`
}

// CollectClassNamesFromSubtreeWithParams - Collects class names for the node with given id and all of it's child nodes.
// Returns -  classNames - Class name list.
func (c *DOM) CollectClassNamesFromSubtreeWithParams(v *DOMCollectClassNamesFromSubtreeParams) ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.collectClassNamesFromSubtree", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ClassNames []string
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

	return chromeData.Result.ClassNames, nil
}

// CollectClassNamesFromSubtree - Collects class names for the node with given id and all of it's child nodes.
// nodeId - Id of the node to collect class names.
// Returns -  classNames - Class name list.
func (c *DOM) CollectClassNamesFromSubtree(nodeId int) ([]string, error) {
	var v DOMCollectClassNamesFromSubtreeParams
	v.NodeId = nodeId
	return c.CollectClassNamesFromSubtreeWithParams(&v)
}

type DOMCopyToParams struct {
	// Id of the node to copy.
	NodeId int `json:"nodeId"`
	// Id of the element to drop the copy into.
	TargetNodeId int `json:"targetNodeId"`
	// Drop the copy before this node (if absent, the copy becomes the last child of `targetNodeId`).
	InsertBeforeNodeId int `json:"insertBeforeNodeId,omitempty"`
}

// CopyToWithParams - Creates a deep copy of the specified node and places it into the target container before the given anchor.
// Returns -  nodeId - Id of the node clone.
func (c *DOM) CopyToWithParams(v *DOMCopyToParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.copyTo", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// CopyTo - Creates a deep copy of the specified node and places it into the target container before the given anchor.
// nodeId - Id of the node to copy.
// targetNodeId - Id of the element to drop the copy into.
// insertBeforeNodeId - Drop the copy before this node (if absent, the copy becomes the last child of `targetNodeId`).
// Returns -  nodeId - Id of the node clone.
func (c *DOM) CopyTo(nodeId int, targetNodeId int, insertBeforeNodeId int) (int, error) {
	var v DOMCopyToParams
	v.NodeId = nodeId
	v.TargetNodeId = targetNodeId
	v.InsertBeforeNodeId = insertBeforeNodeId
	return c.CopyToWithParams(&v)
}

type DOMDescribeNodeParams struct {
	// Identifier of the node.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId string `json:"objectId,omitempty"`
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// DescribeNodeWithParams - Describes node given its id, does not require domain to be enabled. Does not start tracking any objects, can be used for automation.
// Returns -  node - Node description.
func (c *DOM) DescribeNodeWithParams(v *DOMDescribeNodeParams) (*DOMNode, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.describeNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Node *DOMNode
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

	return chromeData.Result.Node, nil
}

// DescribeNode - Describes node given its id, does not require domain to be enabled. Does not start tracking any objects, can be used for automation.
// nodeId - Identifier of the node.
// backendNodeId - Identifier of the backend node.
// objectId - JavaScript object id of the node wrapper.
// depth - The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
// pierce - Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
// Returns -  node - Node description.
func (c *DOM) DescribeNode(nodeId int, backendNodeId int, objectId string, depth int, pierce bool) (*DOMNode, error) {
	var v DOMDescribeNodeParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	v.Depth = depth
	v.Pierce = pierce
	return c.DescribeNodeWithParams(&v)
}

// Disables DOM agent for the given page.
func (c *DOM) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.disable"})
}

type DOMDiscardSearchResultsParams struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`
}

// DiscardSearchResultsWithParams - Discards search results from the session with the given id. `getSearchResults` should no longer be called for that search.
func (c *DOM) DiscardSearchResultsWithParams(v *DOMDiscardSearchResultsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.discardSearchResults", Params: v})
}

// DiscardSearchResults - Discards search results from the session with the given id. `getSearchResults` should no longer be called for that search.
// searchId - Unique search session identifier.
func (c *DOM) DiscardSearchResults(searchId string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDiscardSearchResultsParams
	v.SearchId = searchId
	return c.DiscardSearchResultsWithParams(&v)
}

// Enables DOM agent for the given page.
func (c *DOM) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.enable"})
}

type DOMFocusParams struct {
	// Identifier of the node.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId string `json:"objectId,omitempty"`
}

// FocusWithParams - Focuses the given element.
func (c *DOM) FocusWithParams(v *DOMFocusParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.focus", Params: v})
}

// Focus - Focuses the given element.
// nodeId - Identifier of the node.
// backendNodeId - Identifier of the backend node.
// objectId - JavaScript object id of the node wrapper.
func (c *DOM) Focus(nodeId int, backendNodeId int, objectId string) (*gcdmessage.ChromeResponse, error) {
	var v DOMFocusParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.FocusWithParams(&v)
}

type DOMGetAttributesParams struct {
	// Id of the node to retrieve attibutes for.
	NodeId int `json:"nodeId"`
}

// GetAttributesWithParams - Returns attributes for the specified node.
// Returns -  attributes - An interleaved array of node attribute names and values.
func (c *DOM) GetAttributesWithParams(v *DOMGetAttributesParams) ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getAttributes", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Attributes []string
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

	return chromeData.Result.Attributes, nil
}

// GetAttributes - Returns attributes for the specified node.
// nodeId - Id of the node to retrieve attibutes for.
// Returns -  attributes - An interleaved array of node attribute names and values.
func (c *DOM) GetAttributes(nodeId int) ([]string, error) {
	var v DOMGetAttributesParams
	v.NodeId = nodeId
	return c.GetAttributesWithParams(&v)
}

type DOMGetBoxModelParams struct {
	// Identifier of the node.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId string `json:"objectId,omitempty"`
}

// GetBoxModelWithParams - Returns boxes for the given node.
// Returns -  model - Box model for the node.
func (c *DOM) GetBoxModelWithParams(v *DOMGetBoxModelParams) (*DOMBoxModel, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getBoxModel", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Model *DOMBoxModel
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

	return chromeData.Result.Model, nil
}

// GetBoxModel - Returns boxes for the given node.
// nodeId - Identifier of the node.
// backendNodeId - Identifier of the backend node.
// objectId - JavaScript object id of the node wrapper.
// Returns -  model - Box model for the node.
func (c *DOM) GetBoxModel(nodeId int, backendNodeId int, objectId string) (*DOMBoxModel, error) {
	var v DOMGetBoxModelParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.GetBoxModelWithParams(&v)
}

type DOMGetDocumentParams struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// GetDocumentWithParams - Returns the root DOM node (and optionally the subtree) to the caller.
// Returns -  root - Resulting node.
func (c *DOM) GetDocumentWithParams(v *DOMGetDocumentParams) (*DOMNode, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getDocument", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Root *DOMNode
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

	return chromeData.Result.Root, nil
}

// GetDocument - Returns the root DOM node (and optionally the subtree) to the caller.
// depth - The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
// pierce - Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
// Returns -  root - Resulting node.
func (c *DOM) GetDocument(depth int, pierce bool) (*DOMNode, error) {
	var v DOMGetDocumentParams
	v.Depth = depth
	v.Pierce = pierce
	return c.GetDocumentWithParams(&v)
}

type DOMGetFlattenedDocumentParams struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// GetFlattenedDocumentWithParams - Returns the root DOM node (and optionally the subtree) to the caller.
// Returns -  nodes - Resulting node.
func (c *DOM) GetFlattenedDocumentWithParams(v *DOMGetFlattenedDocumentParams) ([]*DOMNode, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getFlattenedDocument", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Nodes []*DOMNode
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

	return chromeData.Result.Nodes, nil
}

// GetFlattenedDocument - Returns the root DOM node (and optionally the subtree) to the caller.
// depth - The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
// pierce - Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
// Returns -  nodes - Resulting node.
func (c *DOM) GetFlattenedDocument(depth int, pierce bool) ([]*DOMNode, error) {
	var v DOMGetFlattenedDocumentParams
	v.Depth = depth
	v.Pierce = pierce
	return c.GetFlattenedDocumentWithParams(&v)
}

type DOMGetNodeForLocationParams struct {
	// X coordinate.
	X int `json:"x"`
	// Y coordinate.
	Y int `json:"y"`
	// False to skip to the nearest non-UA shadow root ancestor (default: false).
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

// GetNodeForLocationWithParams - Returns node id at given location.
// Returns -  nodeId - Id of the node at given coordinates.
func (c *DOM) GetNodeForLocationWithParams(v *DOMGetNodeForLocationParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getNodeForLocation", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// GetNodeForLocation - Returns node id at given location.
// x - X coordinate.
// y - Y coordinate.
// includeUserAgentShadowDOM - False to skip to the nearest non-UA shadow root ancestor (default: false).
// Returns -  nodeId - Id of the node at given coordinates.
func (c *DOM) GetNodeForLocation(x int, y int, includeUserAgentShadowDOM bool) (int, error) {
	var v DOMGetNodeForLocationParams
	v.X = x
	v.Y = y
	v.IncludeUserAgentShadowDOM = includeUserAgentShadowDOM
	return c.GetNodeForLocationWithParams(&v)
}

type DOMGetOuterHTMLParams struct {
	// Identifier of the node.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId string `json:"objectId,omitempty"`
}

// GetOuterHTMLWithParams - Returns node's HTML markup.
// Returns -  outerHTML - Outer HTML markup.
func (c *DOM) GetOuterHTMLWithParams(v *DOMGetOuterHTMLParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getOuterHTML", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			OuterHTML string
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

	return chromeData.Result.OuterHTML, nil
}

// GetOuterHTML - Returns node's HTML markup.
// nodeId - Identifier of the node.
// backendNodeId - Identifier of the backend node.
// objectId - JavaScript object id of the node wrapper.
// Returns -  outerHTML - Outer HTML markup.
func (c *DOM) GetOuterHTML(nodeId int, backendNodeId int, objectId string) (string, error) {
	var v DOMGetOuterHTMLParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.GetOuterHTMLWithParams(&v)
}

type DOMGetRelayoutBoundaryParams struct {
	// Id of the node.
	NodeId int `json:"nodeId"`
}

// GetRelayoutBoundaryWithParams - Returns the id of the nearest ancestor that is a relayout boundary.
// Returns -  nodeId - Relayout boundary node id for the given node.
func (c *DOM) GetRelayoutBoundaryWithParams(v *DOMGetRelayoutBoundaryParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getRelayoutBoundary", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// GetRelayoutBoundary - Returns the id of the nearest ancestor that is a relayout boundary.
// nodeId - Id of the node.
// Returns -  nodeId - Relayout boundary node id for the given node.
func (c *DOM) GetRelayoutBoundary(nodeId int) (int, error) {
	var v DOMGetRelayoutBoundaryParams
	v.NodeId = nodeId
	return c.GetRelayoutBoundaryWithParams(&v)
}

type DOMGetSearchResultsParams struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`
	// Start index of the search result to be returned.
	FromIndex int `json:"fromIndex"`
	// End index of the search result to be returned.
	ToIndex int `json:"toIndex"`
}

// GetSearchResultsWithParams - Returns search results from given `fromIndex` to given `toIndex` from the search with the given identifier.
// Returns -  nodeIds - Ids of the search result nodes.
func (c *DOM) GetSearchResultsWithParams(v *DOMGetSearchResultsParams) ([]int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getSearchResults", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

// GetSearchResults - Returns search results from given `fromIndex` to given `toIndex` from the search with the given identifier.
// searchId - Unique search session identifier.
// fromIndex - Start index of the search result to be returned.
// toIndex - End index of the search result to be returned.
// Returns -  nodeIds - Ids of the search result nodes.
func (c *DOM) GetSearchResults(searchId string, fromIndex int, toIndex int) ([]int, error) {
	var v DOMGetSearchResultsParams
	v.SearchId = searchId
	v.FromIndex = fromIndex
	v.ToIndex = toIndex
	return c.GetSearchResultsWithParams(&v)
}

// Hides any highlight.
func (c *DOM) HideHighlight() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.hideHighlight"})
}

// Highlights DOM node.
func (c *DOM) HighlightNode() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.highlightNode"})
}

// Highlights given rectangle.
func (c *DOM) HighlightRect() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.highlightRect"})
}

// Marks last undoable state.
func (c *DOM) MarkUndoableState() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.markUndoableState"})
}

type DOMMoveToParams struct {
	// Id of the node to move.
	NodeId int `json:"nodeId"`
	// Id of the element to drop the moved node into.
	TargetNodeId int `json:"targetNodeId"`
	// Drop node before this one (if absent, the moved node becomes the last child of `targetNodeId`).
	InsertBeforeNodeId int `json:"insertBeforeNodeId,omitempty"`
}

// MoveToWithParams - Moves node into the new container, places it before the given anchor.
// Returns -  nodeId - New id of the moved node.
func (c *DOM) MoveToWithParams(v *DOMMoveToParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.moveTo", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// MoveTo - Moves node into the new container, places it before the given anchor.
// nodeId - Id of the node to move.
// targetNodeId - Id of the element to drop the moved node into.
// insertBeforeNodeId - Drop node before this one (if absent, the moved node becomes the last child of `targetNodeId`).
// Returns -  nodeId - New id of the moved node.
func (c *DOM) MoveTo(nodeId int, targetNodeId int, insertBeforeNodeId int) (int, error) {
	var v DOMMoveToParams
	v.NodeId = nodeId
	v.TargetNodeId = targetNodeId
	v.InsertBeforeNodeId = insertBeforeNodeId
	return c.MoveToWithParams(&v)
}

type DOMPerformSearchParams struct {
	// Plain text or query selector or XPath search query.
	Query string `json:"query"`
	// True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}

// PerformSearchWithParams - Searches for a given string in the DOM tree. Use `getSearchResults` to access search results or `cancelSearch` to end this search session.
// Returns -  searchId - Unique search session identifier. resultCount - Number of search results.
func (c *DOM) PerformSearchWithParams(v *DOMPerformSearchParams) (string, int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.performSearch", Params: v})
	if err != nil {
		return "", 0, err
	}

	var chromeData struct {
		Result struct {
			SearchId    string
			ResultCount int
		}
	}

	if resp == nil {
		return "", 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", 0, err
	}

	return chromeData.Result.SearchId, chromeData.Result.ResultCount, nil
}

// PerformSearch - Searches for a given string in the DOM tree. Use `getSearchResults` to access search results or `cancelSearch` to end this search session.
// query - Plain text or query selector or XPath search query.
// includeUserAgentShadowDOM - True to search in user agent shadow DOM.
// Returns -  searchId - Unique search session identifier. resultCount - Number of search results.
func (c *DOM) PerformSearch(query string, includeUserAgentShadowDOM bool) (string, int, error) {
	var v DOMPerformSearchParams
	v.Query = query
	v.IncludeUserAgentShadowDOM = includeUserAgentShadowDOM
	return c.PerformSearchWithParams(&v)
}

type DOMPushNodeByPathToFrontendParams struct {
	// Path to node in the proprietary format.
	Path string `json:"path"`
}

// PushNodeByPathToFrontendWithParams - Requests that the node is sent to the caller given its path. // FIXME, use XPath
// Returns -  nodeId - Id of the node for given path.
func (c *DOM) PushNodeByPathToFrontendWithParams(v *DOMPushNodeByPathToFrontendParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.pushNodeByPathToFrontend", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// PushNodeByPathToFrontend - Requests that the node is sent to the caller given its path. // FIXME, use XPath
// path - Path to node in the proprietary format.
// Returns -  nodeId - Id of the node for given path.
func (c *DOM) PushNodeByPathToFrontend(path string) (int, error) {
	var v DOMPushNodeByPathToFrontendParams
	v.Path = path
	return c.PushNodeByPathToFrontendWithParams(&v)
}

type DOMPushNodesByBackendIdsToFrontendParams struct {
	// The array of backend node ids.
	BackendNodeIds []int `json:"backendNodeIds"`
}

// PushNodesByBackendIdsToFrontendWithParams - Requests that a batch of nodes is sent to the caller given their backend node ids.
// Returns -  nodeIds - The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
func (c *DOM) PushNodesByBackendIdsToFrontendWithParams(v *DOMPushNodesByBackendIdsToFrontendParams) ([]int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.pushNodesByBackendIdsToFrontend", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

// PushNodesByBackendIdsToFrontend - Requests that a batch of nodes is sent to the caller given their backend node ids.
// backendNodeIds - The array of backend node ids.
// Returns -  nodeIds - The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
func (c *DOM) PushNodesByBackendIdsToFrontend(backendNodeIds []int) ([]int, error) {
	var v DOMPushNodesByBackendIdsToFrontendParams
	v.BackendNodeIds = backendNodeIds
	return c.PushNodesByBackendIdsToFrontendWithParams(&v)
}

type DOMQuerySelectorParams struct {
	// Id of the node to query upon.
	NodeId int `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

// QuerySelectorWithParams - Executes `querySelector` on a given node.
// Returns -  nodeId - Query selector result.
func (c *DOM) QuerySelectorWithParams(v *DOMQuerySelectorParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.querySelector", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// QuerySelector - Executes `querySelector` on a given node.
// nodeId - Id of the node to query upon.
// selector - Selector string.
// Returns -  nodeId - Query selector result.
func (c *DOM) QuerySelector(nodeId int, selector string) (int, error) {
	var v DOMQuerySelectorParams
	v.NodeId = nodeId
	v.Selector = selector
	return c.QuerySelectorWithParams(&v)
}

type DOMQuerySelectorAllParams struct {
	// Id of the node to query upon.
	NodeId int `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

// QuerySelectorAllWithParams - Executes `querySelectorAll` on a given node.
// Returns -  nodeIds - Query selector result.
func (c *DOM) QuerySelectorAllWithParams(v *DOMQuerySelectorAllParams) ([]int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.querySelectorAll", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			NodeIds []int
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

	return chromeData.Result.NodeIds, nil
}

// QuerySelectorAll - Executes `querySelectorAll` on a given node.
// nodeId - Id of the node to query upon.
// selector - Selector string.
// Returns -  nodeIds - Query selector result.
func (c *DOM) QuerySelectorAll(nodeId int, selector string) ([]int, error) {
	var v DOMQuerySelectorAllParams
	v.NodeId = nodeId
	v.Selector = selector
	return c.QuerySelectorAllWithParams(&v)
}

// Re-does the last undone action.
func (c *DOM) Redo() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.redo"})
}

type DOMRemoveAttributeParams struct {
	// Id of the element to remove attribute from.
	NodeId int `json:"nodeId"`
	// Name of the attribute to remove.
	Name string `json:"name"`
}

// RemoveAttributeWithParams - Removes attribute with given name from an element with given id.
func (c *DOM) RemoveAttributeWithParams(v *DOMRemoveAttributeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.removeAttribute", Params: v})
}

// RemoveAttribute - Removes attribute with given name from an element with given id.
// nodeId - Id of the element to remove attribute from.
// name - Name of the attribute to remove.
func (c *DOM) RemoveAttribute(nodeId int, name string) (*gcdmessage.ChromeResponse, error) {
	var v DOMRemoveAttributeParams
	v.NodeId = nodeId
	v.Name = name
	return c.RemoveAttributeWithParams(&v)
}

type DOMRemoveNodeParams struct {
	// Id of the node to remove.
	NodeId int `json:"nodeId"`
}

// RemoveNodeWithParams - Removes node with given id.
func (c *DOM) RemoveNodeWithParams(v *DOMRemoveNodeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.removeNode", Params: v})
}

// RemoveNode - Removes node with given id.
// nodeId - Id of the node to remove.
func (c *DOM) RemoveNode(nodeId int) (*gcdmessage.ChromeResponse, error) {
	var v DOMRemoveNodeParams
	v.NodeId = nodeId
	return c.RemoveNodeWithParams(&v)
}

type DOMRequestChildNodesParams struct {
	// Id of the node to get children for.
	NodeId int `json:"nodeId"`
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
	Pierce bool `json:"pierce,omitempty"`
}

// RequestChildNodesWithParams - Requests that children of the node with given id are returned to the caller in form of `setChildNodes` events where not only immediate children are retrieved, but all children down to the specified depth.
func (c *DOM) RequestChildNodesWithParams(v *DOMRequestChildNodesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.requestChildNodes", Params: v})
}

// RequestChildNodes - Requests that children of the node with given id are returned to the caller in form of `setChildNodes` events where not only immediate children are retrieved, but all children down to the specified depth.
// nodeId - Id of the node to get children for.
// depth - The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
// pierce - Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
func (c *DOM) RequestChildNodes(nodeId int, depth int, pierce bool) (*gcdmessage.ChromeResponse, error) {
	var v DOMRequestChildNodesParams
	v.NodeId = nodeId
	v.Depth = depth
	v.Pierce = pierce
	return c.RequestChildNodesWithParams(&v)
}

type DOMRequestNodeParams struct {
	// JavaScript object id to convert into node.
	ObjectId string `json:"objectId"`
}

// RequestNodeWithParams - Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of `setChildNodes` notifications.
// Returns -  nodeId - Node id for given object.
func (c *DOM) RequestNodeWithParams(v *DOMRequestNodeParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.requestNode", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// RequestNode - Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of `setChildNodes` notifications.
// objectId - JavaScript object id to convert into node.
// Returns -  nodeId - Node id for given object.
func (c *DOM) RequestNode(objectId string) (int, error) {
	var v DOMRequestNodeParams
	v.ObjectId = objectId
	return c.RequestNodeWithParams(&v)
}

type DOMResolveNodeParams struct {
	// Id of the node to resolve.
	NodeId int `json:"nodeId,omitempty"`
	// Backend identifier of the node to resolve.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// ResolveNodeWithParams - Resolves the JavaScript node object for a given NodeId or BackendNodeId.
// Returns -  object - JavaScript object wrapper for given node.
func (c *DOM) ResolveNodeWithParams(v *DOMResolveNodeParams) (*RuntimeRemoteObject, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.resolveNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Object *RuntimeRemoteObject
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

	return chromeData.Result.Object, nil
}

// ResolveNode - Resolves the JavaScript node object for a given NodeId or BackendNodeId.
// nodeId - Id of the node to resolve.
// backendNodeId - Backend identifier of the node to resolve.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  object - JavaScript object wrapper for given node.
func (c *DOM) ResolveNode(nodeId int, backendNodeId int, objectGroup string) (*RuntimeRemoteObject, error) {
	var v DOMResolveNodeParams
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectGroup = objectGroup
	return c.ResolveNodeWithParams(&v)
}

type DOMSetAttributeValueParams struct {
	// Id of the element to set attribute for.
	NodeId int `json:"nodeId"`
	// Attribute name.
	Name string `json:"name"`
	// Attribute value.
	Value string `json:"value"`
}

// SetAttributeValueWithParams - Sets attribute for an element with given id.
func (c *DOM) SetAttributeValueWithParams(v *DOMSetAttributeValueParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setAttributeValue", Params: v})
}

// SetAttributeValue - Sets attribute for an element with given id.
// nodeId - Id of the element to set attribute for.
// name - Attribute name.
// value - Attribute value.
func (c *DOM) SetAttributeValue(nodeId int, name string, value string) (*gcdmessage.ChromeResponse, error) {
	var v DOMSetAttributeValueParams
	v.NodeId = nodeId
	v.Name = name
	v.Value = value
	return c.SetAttributeValueWithParams(&v)
}

type DOMSetAttributesAsTextParams struct {
	// Id of the element to set attributes for.
	NodeId int `json:"nodeId"`
	// Text with a number of attributes. Will parse this text using HTML parser.
	Text string `json:"text"`
	// Attribute name to replace with new attributes derived from text in case text parsed successfully.
	Name string `json:"name,omitempty"`
}

// SetAttributesAsTextWithParams - Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
func (c *DOM) SetAttributesAsTextWithParams(v *DOMSetAttributesAsTextParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setAttributesAsText", Params: v})
}

// SetAttributesAsText - Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
// nodeId - Id of the element to set attributes for.
// text - Text with a number of attributes. Will parse this text using HTML parser.
// name - Attribute name to replace with new attributes derived from text in case text parsed successfully.
func (c *DOM) SetAttributesAsText(nodeId int, text string, name string) (*gcdmessage.ChromeResponse, error) {
	var v DOMSetAttributesAsTextParams
	v.NodeId = nodeId
	v.Text = text
	v.Name = name
	return c.SetAttributesAsTextWithParams(&v)
}

type DOMSetFileInputFilesParams struct {
	// Array of file paths to set.
	Files []string `json:"files"`
	// Identifier of the node.
	NodeId int `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId int `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId string `json:"objectId,omitempty"`
}

// SetFileInputFilesWithParams - Sets files for the given file input element.
func (c *DOM) SetFileInputFilesWithParams(v *DOMSetFileInputFilesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setFileInputFiles", Params: v})
}

// SetFileInputFiles - Sets files for the given file input element.
// files - Array of file paths to set.
// nodeId - Identifier of the node.
// backendNodeId - Identifier of the backend node.
// objectId - JavaScript object id of the node wrapper.
func (c *DOM) SetFileInputFiles(files []string, nodeId int, backendNodeId int, objectId string) (*gcdmessage.ChromeResponse, error) {
	var v DOMSetFileInputFilesParams
	v.Files = files
	v.NodeId = nodeId
	v.BackendNodeId = backendNodeId
	v.ObjectId = objectId
	return c.SetFileInputFilesWithParams(&v)
}

type DOMSetInspectedNodeParams struct {
	// DOM node id to be accessible by means of $x command line API.
	NodeId int `json:"nodeId"`
}

// SetInspectedNodeWithParams - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
func (c *DOM) SetInspectedNodeWithParams(v *DOMSetInspectedNodeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setInspectedNode", Params: v})
}

// SetInspectedNode - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// nodeId - DOM node id to be accessible by means of $x command line API.
func (c *DOM) SetInspectedNode(nodeId int) (*gcdmessage.ChromeResponse, error) {
	var v DOMSetInspectedNodeParams
	v.NodeId = nodeId
	return c.SetInspectedNodeWithParams(&v)
}

type DOMSetNodeNameParams struct {
	// Id of the node to set name for.
	NodeId int `json:"nodeId"`
	// New node's name.
	Name string `json:"name"`
}

// SetNodeNameWithParams - Sets node name for a node with given id.
// Returns -  nodeId - New node's id.
func (c *DOM) SetNodeNameWithParams(v *DOMSetNodeNameParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setNodeName", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// SetNodeName - Sets node name for a node with given id.
// nodeId - Id of the node to set name for.
// name - New node's name.
// Returns -  nodeId - New node's id.
func (c *DOM) SetNodeName(nodeId int, name string) (int, error) {
	var v DOMSetNodeNameParams
	v.NodeId = nodeId
	v.Name = name
	return c.SetNodeNameWithParams(&v)
}

type DOMSetNodeValueParams struct {
	// Id of the node to set value for.
	NodeId int `json:"nodeId"`
	// New node's value.
	Value string `json:"value"`
}

// SetNodeValueWithParams - Sets node value for a node with given id.
func (c *DOM) SetNodeValueWithParams(v *DOMSetNodeValueParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setNodeValue", Params: v})
}

// SetNodeValue - Sets node value for a node with given id.
// nodeId - Id of the node to set value for.
// value - New node's value.
func (c *DOM) SetNodeValue(nodeId int, value string) (*gcdmessage.ChromeResponse, error) {
	var v DOMSetNodeValueParams
	v.NodeId = nodeId
	v.Value = value
	return c.SetNodeValueWithParams(&v)
}

type DOMSetOuterHTMLParams struct {
	// Id of the node to set markup for.
	NodeId int `json:"nodeId"`
	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

// SetOuterHTMLWithParams - Sets node HTML markup, returns new node id.
func (c *DOM) SetOuterHTMLWithParams(v *DOMSetOuterHTMLParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.setOuterHTML", Params: v})
}

// SetOuterHTML - Sets node HTML markup, returns new node id.
// nodeId - Id of the node to set markup for.
// outerHTML - Outer HTML markup to set.
func (c *DOM) SetOuterHTML(nodeId int, outerHTML string) (*gcdmessage.ChromeResponse, error) {
	var v DOMSetOuterHTMLParams
	v.NodeId = nodeId
	v.OuterHTML = outerHTML
	return c.SetOuterHTMLWithParams(&v)
}

// Undoes the last performed action.
func (c *DOM) Undo() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.undo"})
}

type DOMGetFrameOwnerParams struct {
	//
	FrameId string `json:"frameId"`
}

// GetFrameOwnerWithParams - Returns iframe node that owns iframe with the given domain.
// Returns -  nodeId -
func (c *DOM) GetFrameOwnerWithParams(v *DOMGetFrameOwnerParams) (int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOM.getFrameOwner", Params: v})
	if err != nil {
		return 0, err
	}

	var chromeData struct {
		Result struct {
			NodeId int
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

	return chromeData.Result.NodeId, nil
}

// GetFrameOwner - Returns iframe node that owns iframe with the given domain.
// frameId -
// Returns -  nodeId -
func (c *DOM) GetFrameOwner(frameId string) (int, error) {
	var v DOMGetFrameOwnerParams
	v.FrameId = frameId
	return c.GetFrameOwnerWithParams(&v)
}
