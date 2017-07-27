// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMSnapshot functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// A Node in the DOM tree.
type DOMSnapshotDOMNode struct {
	NodeType              int                     `json:"nodeType"`                        // <code>Node</code>'s nodeType.
	NodeName              string                  `json:"nodeName"`                        // <code>Node</code>'s nodeName.
	NodeValue             string                  `json:"nodeValue"`                       // <code>Node</code>'s nodeValue.
	BackendNodeId         int                     `json:"backendNodeId"`                   // <code>Node</code>'s id, corresponds to DOM.Node.backendNodeId.
	ChildNodeIndexes      []int                   `json:"childNodeIndexes,omitempty"`      // The indexes of the node's child nodes in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	Attributes            []*DOMSnapshotNameValue `json:"attributes,omitempty"`            // Attributes of an <code>Element</code> node.
	PseudoElementIndexes  []int                   `json:"pseudoElementIndexes,omitempty"`  // Indexes of pseudo elements associated with this node in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	LayoutNodeIndex       int                     `json:"layoutNodeIndex,omitempty"`       // The index of the node's related layout tree node in the <code>layoutTreeNodes</code> array returned by <code>getSnapshot</code>, if any.
	DocumentURL           string                  `json:"documentURL,omitempty"`           // Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	BaseURL               string                  `json:"baseURL,omitempty"`               // Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	ContentLanguage       string                  `json:"contentLanguage,omitempty"`       // Only set for documents, contains the document's content language.
	PublicId              string                  `json:"publicId,omitempty"`              // <code>DocumentType</code> node's publicId.
	SystemId              string                  `json:"systemId,omitempty"`              // <code>DocumentType</code> node's systemId.
	FrameId               string                  `json:"frameId,omitempty"`               // Frame ID for frame owner elements.
	ContentDocumentIndex  int                     `json:"contentDocumentIndex,omitempty"`  // The index of a frame owner element's content document in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	ImportedDocumentIndex int                     `json:"importedDocumentIndex,omitempty"` // Index of the imported document's node of a link element in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	TemplateContentIndex  int                     `json:"templateContentIndex,omitempty"`  // Index of the content node of a template element in the <code>domNodes</code> array returned by <code>getSnapshot</code>.
	PseudoType            string                  `json:"pseudoType,omitempty"`            // Type of a pseudo element node. enum values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button
	IsClickable           bool                    `json:"isClickable,omitempty"`           // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
}

// Details of an element in the DOM tree with a LayoutObject.
type DOMSnapshotLayoutTreeNode struct {
	DomNodeIndex    int                 `json:"domNodeIndex"`              // The index of the related DOM node in the <code>domNodes</code> array returned by <code>getSnapshot</code>.
	BoundingBox     *DOMRect            `json:"boundingBox"`               // The absolute position bounding box.
	LayoutText      string              `json:"layoutText,omitempty"`      // Contents of the LayoutText, if any.
	InlineTextNodes []*CSSInlineTextBox `json:"inlineTextNodes,omitempty"` // The post-layout inline text nodes, if any.
	StyleIndex      int                 `json:"styleIndex,omitempty"`      // Index into the <code>computedStyles</code> array returned by <code>getSnapshot</code>.
}

// A subset of the full ComputedStyle as defined by the request whitelist.
type DOMSnapshotComputedStyle struct {
	Properties []*DOMSnapshotNameValue `json:"properties"` // Name/value pairs of computed style properties.
}

// A name/value pair.
type DOMSnapshotNameValue struct {
	Name  string `json:"name"`  // Attribute/property name.
	Value string `json:"value"` // Attribute/property value.
}

type DOMSnapshot struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMSnapshot(target gcdmessage.ChromeTargeter) *DOMSnapshot {
	c := &DOMSnapshot{target: target}
	return c
}

type DOMSnapshotGetSnapshotParams struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}

// GetSnapshotWithParams - Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
// Returns -  domNodes - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. layoutTreeNodes - The nodes in the layout tree. computedStyles - Whitelisted ComputedStyle properties for each node in the layout tree.
func (c *DOMSnapshot) GetSnapshotWithParams(v *DOMSnapshotGetSnapshotParams) ([]*DOMSnapshotDOMNode, []*DOMSnapshotLayoutTreeNode, []*DOMSnapshotComputedStyle, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMSnapshot.getSnapshot", Params: v})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			DomNodes        []*DOMSnapshotDOMNode
			LayoutTreeNodes []*DOMSnapshotLayoutTreeNode
			ComputedStyles  []*DOMSnapshotComputedStyle
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

	return chromeData.Result.DomNodes, chromeData.Result.LayoutTreeNodes, chromeData.Result.ComputedStyles, nil
}

// GetSnapshot - Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
// computedStyleWhitelist - Whitelist of computed styles to return.
// Returns -  domNodes - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. layoutTreeNodes - The nodes in the layout tree. computedStyles - Whitelisted ComputedStyle properties for each node in the layout tree.
func (c *DOMSnapshot) GetSnapshot(computedStyleWhitelist []string) ([]*DOMSnapshotDOMNode, []*DOMSnapshotLayoutTreeNode, []*DOMSnapshotComputedStyle, error) {
	var v DOMSnapshotGetSnapshotParams
	v.ComputedStyleWhitelist = computedStyleWhitelist
	return c.GetSnapshotWithParams(&v)
}
