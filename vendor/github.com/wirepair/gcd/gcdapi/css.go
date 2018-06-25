// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CSS functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// CSS rule collection for a single pseudo style.
type CSSPseudoElementMatches struct {
	PseudoType string          `json:"pseudoType"` // Pseudo element type. enum values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button
	Matches    []*CSSRuleMatch `json:"matches"`    // Matches of CSS rules applicable to the pseudo style.
}

// Inherited CSS rule collection from ancestor node.
type CSSInheritedStyleEntry struct {
	InlineStyle     *CSSCSSStyle    `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []*CSSRuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// Match data for a CSS rule.
type CSSRuleMatch struct {
	Rule              *CSSCSSRule `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int       `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Data for a simple selector (these are delimited by commas in a selector list).
type CSSValue struct {
	Text  string          `json:"text"`            // Value text.
	Range *CSSSourceRange `json:"range,omitempty"` // Value range in the underlying resource (if available).
}

// Selector list data.
type CSSSelectorList struct {
	Selectors []*CSSValue `json:"selectors"` // Selectors in the list.
	Text      string      `json:"text"`      // Rule selector text.
}

// CSS stylesheet metainformation.
type CSSCSSStyleSheetHeader struct {
	StyleSheetId string  `json:"styleSheetId"`           // The stylesheet identifier.
	FrameId      string  `json:"frameId"`                // Owner frame identifier.
	SourceURL    string  `json:"sourceURL"`              // Stylesheet resource URL.
	SourceMapURL string  `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin       string  `json:"origin"`                 // Stylesheet origin. enum values: injected, user-agent, inspector, regular
	Title        string  `json:"title"`                  // Stylesheet title.
	OwnerNode    int     `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled     bool    `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL bool    `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline     bool    `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine    float64 `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn  float64 `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
	Length       float64 `json:"length"`                 // Size of the content (in characters).
}

// CSS rule representation.
type CSSCSSRule struct {
	StyleSheetId string           `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList *CSSSelectorList `json:"selectorList"`           // Rule selector data.
	Origin       string           `json:"origin"`                 // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	Style        *CSSCSSStyle     `json:"style"`                  // Associated style declaration.
	Media        []*CSSCSSMedia   `json:"media,omitempty"`        // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
}

// CSS coverage information.
type CSSRuleUsage struct {
	StyleSheetId string  `json:"styleSheetId"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StartOffset  float64 `json:"startOffset"`  // Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	EndOffset    float64 `json:"endOffset"`    // Offset of the end of the rule body from the beginning of the stylesheet.
	Used         bool    `json:"used"`         // Indicates whether the rule was actually used by some element in the page.
}

// Text range within a resource. All numbers are zero-based.
type CSSSourceRange struct {
	StartLine   int `json:"startLine"`   // Start line of range.
	StartColumn int `json:"startColumn"` // Start column of range (inclusive).
	EndLine     int `json:"endLine"`     // End line of range
	EndColumn   int `json:"endColumn"`   // End column of range (exclusive).
}

// No Description.
type CSSShorthandEntry struct {
	Name      string `json:"name"`                // Shorthand name.
	Value     string `json:"value"`               // Shorthand value.
	Important bool   `json:"important,omitempty"` // Whether the property has "!important" annotation (implies `false` if absent).
}

// No Description.
type CSSCSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// CSS style representation.
type CSSCSSStyle struct {
	StyleSheetId     string               `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CssProperties    []*CSSCSSProperty    `json:"cssProperties"`          // CSS properties in the style.
	ShorthandEntries []*CSSShorthandEntry `json:"shorthandEntries"`       // Computed values for all shorthands found in the style.
	CssText          string               `json:"cssText,omitempty"`      // Style declaration text (if available).
	Range            *CSSSourceRange      `json:"range,omitempty"`        // Style declaration range in the enclosing stylesheet (if available).
}

// CSS property declaration data.
type CSSCSSProperty struct {
	Name      string          `json:"name"`                // The property name.
	Value     string          `json:"value"`               // The property value.
	Important bool            `json:"important,omitempty"` // Whether the property has "!important" annotation (implies `false` if absent).
	Implicit  bool            `json:"implicit,omitempty"`  // Whether the property is implicit (implies `false` if absent).
	Text      string          `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  bool            `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies `true` if absent).
	Disabled  bool            `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *CSSSourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// CSS media rule descriptor.
type CSSCSSMedia struct {
	Text         string           `json:"text"`                   // Media query text.
	Source       string           `json:"source"`                 // Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL    string           `json:"sourceURL,omitempty"`    // URL of the document containing the media query description.
	Range        *CSSSourceRange  `json:"range,omitempty"`        // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	StyleSheetId string           `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList    []*CSSMediaQuery `json:"mediaList,omitempty"`    // Array of media queries.
}

// Media query descriptor.
type CSSMediaQuery struct {
	Expressions []*CSSMediaQueryExpression `json:"expressions"` // Array of media query expressions.
	Active      bool                       `json:"active"`      // Whether the media query condition is satisfied.
}

// Media query expression descriptor.
type CSSMediaQueryExpression struct {
	Value          float64         `json:"value"`                    // Media query expression value.
	Unit           string          `json:"unit"`                     // Media query expression units.
	Feature        string          `json:"feature"`                  // Media query expression feature.
	ValueRange     *CSSSourceRange `json:"valueRange,omitempty"`     // The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength float64         `json:"computedLength,omitempty"` // Computed length of media query expression (if applicable).
}

// Information about amount of glyphs that were rendered with given font.
type CSSPlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`   // Font's family name reported by platform.
	IsCustomFont bool    `json:"isCustomFont"` // Indicates if the font was downloaded or resolved locally.
	GlyphCount   float64 `json:"glyphCount"`   // Amount of glyphs that were rendered with this font.
}

// Properties of a web font: https://www.w3.org/TR/2008/REC-CSS2-20080411/fonts.html#font-descriptions
type CSSFontFace struct {
	FontFamily         string `json:"fontFamily"`         // The font-family.
	FontStyle          string `json:"fontStyle"`          // The font-style.
	FontVariant        string `json:"fontVariant"`        // The font-variant.
	FontWeight         string `json:"fontWeight"`         // The font-weight.
	FontStretch        string `json:"fontStretch"`        // The font-stretch.
	UnicodeRange       string `json:"unicodeRange"`       // The unicode-range.
	Src                string `json:"src"`                // The src.
	PlatformFontFamily string `json:"platformFontFamily"` // The resolved platform font family
}

// CSS keyframes rule representation.
type CSSCSSKeyframesRule struct {
	AnimationName *CSSValue             `json:"animationName"` // Animation name.
	Keyframes     []*CSSCSSKeyframeRule `json:"keyframes"`     // List of keyframes.
}

// CSS keyframe rule representation.
type CSSCSSKeyframeRule struct {
	StyleSheetId string       `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	Origin       string       `json:"origin"`                 // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	KeyText      *CSSValue    `json:"keyText"`                // Associated key text.
	Style        *CSSCSSStyle `json:"style"`                  // Associated style declaration.
}

// A descriptor of operation to mutate style declaration text.
type CSSStyleDeclarationEdit struct {
	StyleSheetId string          `json:"styleSheetId"` // The css style sheet identifier.
	Range        *CSSSourceRange `json:"range"`        // The range of the style text in the enclosing stylesheet.
	Text         string          `json:"text"`         // New style text.
}

// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded web font
type CSSFontsUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Font *CSSFontFace `json:"font,omitempty"` // The web font that has loaded.
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is added.
type CSSStyleSheetAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Header *CSSCSSStyleSheetHeader `json:"header"` // Added stylesheet metainfo.
	} `json:"Params,omitempty"`
}

// Fired whenever a stylesheet is changed as a result of the client operation.
type CSSStyleSheetChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` //
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is removed.
type CSSStyleSheetRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` // Identifier of the removed stylesheet.
	} `json:"Params,omitempty"`
}

type CSS struct {
	target gcdmessage.ChromeTargeter
}

func NewCSS(target gcdmessage.ChromeTargeter) *CSS {
	c := &CSS{target: target}
	return c
}

type CSSAddRuleParams struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetId string `json:"styleSheetId"`
	// The text of a new rule.
	RuleText string `json:"ruleText"`
	// Text position of a new rule in the target style sheet.
	Location *CSSSourceRange `json:"location"`
}

// AddRuleWithParams - Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the position specified by `location`.
// Returns -  rule - The newly created rule.
func (c *CSS) AddRuleWithParams(v *CSSAddRuleParams) (*CSSCSSRule, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.addRule", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Rule *CSSCSSRule
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

	return chromeData.Result.Rule, nil
}

// AddRule - Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the position specified by `location`.
// styleSheetId - The css style sheet identifier where a new rule should be inserted.
// ruleText - The text of a new rule.
// location - Text position of a new rule in the target style sheet.
// Returns -  rule - The newly created rule.
func (c *CSS) AddRule(styleSheetId string, ruleText string, location *CSSSourceRange) (*CSSCSSRule, error) {
	var v CSSAddRuleParams
	v.StyleSheetId = styleSheetId
	v.RuleText = ruleText
	v.Location = location
	return c.AddRuleWithParams(&v)
}

type CSSCollectClassNamesParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
}

// CollectClassNamesWithParams - Returns all class names from specified stylesheet.
// Returns -  classNames - Class name list.
func (c *CSS) CollectClassNamesWithParams(v *CSSCollectClassNamesParams) ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.collectClassNames", Params: v})
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

// CollectClassNames - Returns all class names from specified stylesheet.
// styleSheetId -
// Returns -  classNames - Class name list.
func (c *CSS) CollectClassNames(styleSheetId string) ([]string, error) {
	var v CSSCollectClassNamesParams
	v.StyleSheetId = styleSheetId
	return c.CollectClassNamesWithParams(&v)
}

type CSSCreateStyleSheetParams struct {
	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameId string `json:"frameId"`
}

// CreateStyleSheetWithParams - Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
// Returns -  styleSheetId - Identifier of the created "via-inspector" stylesheet.
func (c *CSS) CreateStyleSheetWithParams(v *CSSCreateStyleSheetParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.createStyleSheet", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			StyleSheetId string
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

	return chromeData.Result.StyleSheetId, nil
}

// CreateStyleSheet - Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
// frameId - Identifier of the frame where "via-inspector" stylesheet should be created.
// Returns -  styleSheetId - Identifier of the created "via-inspector" stylesheet.
func (c *CSS) CreateStyleSheet(frameId string) (string, error) {
	var v CSSCreateStyleSheetParams
	v.FrameId = frameId
	return c.CreateStyleSheetWithParams(&v)
}

// Disables the CSS agent for the given page.
func (c *CSS) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.disable"})
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (c *CSS) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.enable"})
}

type CSSForcePseudoStateParams struct {
	// The element id for which to force the pseudo state.
	NodeId int `json:"nodeId"`
	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

// ForcePseudoStateWithParams - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
func (c *CSS) ForcePseudoStateWithParams(v *CSSForcePseudoStateParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.forcePseudoState", Params: v})
}

// ForcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *CSS) ForcePseudoState(nodeId int, forcedPseudoClasses []string) (*gcdmessage.ChromeResponse, error) {
	var v CSSForcePseudoStateParams
	v.NodeId = nodeId
	v.ForcedPseudoClasses = forcedPseudoClasses
	return c.ForcePseudoStateWithParams(&v)
}

type CSSGetBackgroundColorsParams struct {
	// Id of the node to get background colors for.
	NodeId int `json:"nodeId"`
}

// GetBackgroundColorsWithParams -
// Returns -  backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load). computedFontSize - The computed font size for this node, as a CSS computed value string (e.g. '12px'). computedFontWeight - The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100'). computedBodyFontSize - The computed font size for the document body, as a computed CSS value string (e.g. '16px').
func (c *CSS) GetBackgroundColorsWithParams(v *CSSGetBackgroundColorsParams) ([]string, string, string, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getBackgroundColors", Params: v})
	if err != nil {
		return nil, "", "", "", err
	}

	var chromeData struct {
		Result struct {
			BackgroundColors     []string
			ComputedFontSize     string
			ComputedFontWeight   string
			ComputedBodyFontSize string
		}
	}

	if resp == nil {
		return nil, "", "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, "", "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, "", "", "", err
	}

	return chromeData.Result.BackgroundColors, chromeData.Result.ComputedFontSize, chromeData.Result.ComputedFontWeight, chromeData.Result.ComputedBodyFontSize, nil
}

// GetBackgroundColors -
// nodeId - Id of the node to get background colors for.
// Returns -  backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load). computedFontSize - The computed font size for this node, as a CSS computed value string (e.g. '12px'). computedFontWeight - The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100'). computedBodyFontSize - The computed font size for the document body, as a computed CSS value string (e.g. '16px').
func (c *CSS) GetBackgroundColors(nodeId int) ([]string, string, string, string, error) {
	var v CSSGetBackgroundColorsParams
	v.NodeId = nodeId
	return c.GetBackgroundColorsWithParams(&v)
}

type CSSGetComputedStyleForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetComputedStyleForNodeWithParams - Returns the computed style for a DOM node identified by `nodeId`.
// Returns -  computedStyle - Computed style for the specified DOM node.
func (c *CSS) GetComputedStyleForNodeWithParams(v *CSSGetComputedStyleForNodeParams) ([]*CSSCSSComputedStyleProperty, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getComputedStyleForNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ComputedStyle []*CSSCSSComputedStyleProperty
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

	return chromeData.Result.ComputedStyle, nil
}

// GetComputedStyleForNode - Returns the computed style for a DOM node identified by `nodeId`.
// nodeId -
// Returns -  computedStyle - Computed style for the specified DOM node.
func (c *CSS) GetComputedStyleForNode(nodeId int) ([]*CSSCSSComputedStyleProperty, error) {
	var v CSSGetComputedStyleForNodeParams
	v.NodeId = nodeId
	return c.GetComputedStyleForNodeWithParams(&v)
}

type CSSGetInlineStylesForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetInlineStylesForNodeWithParams - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by `nodeId`.
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *CSS) GetInlineStylesForNodeWithParams(v *CSSGetInlineStylesForNodeParams) (*CSSCSSStyle, *CSSCSSStyle, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getInlineStylesForNode", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			InlineStyle     *CSSCSSStyle
			AttributesStyle *CSSCSSStyle
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, nil
}

// GetInlineStylesForNode - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by `nodeId`.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *CSS) GetInlineStylesForNode(nodeId int) (*CSSCSSStyle, *CSSCSSStyle, error) {
	var v CSSGetInlineStylesForNodeParams
	v.NodeId = nodeId
	return c.GetInlineStylesForNodeWithParams(&v)
}

type CSSGetMatchedStylesForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetMatchedStylesForNodeWithParams - Returns requested styles for a DOM node identified by `nodeId`.
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%"). matchedCSSRules - CSS rules matching this node, from all applicable stylesheets. pseudoElements - Pseudo style matches for this node. inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root). cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (c *CSS) GetMatchedStylesForNodeWithParams(v *CSSGetMatchedStylesForNodeParams) (*CSSCSSStyle, *CSSCSSStyle, []*CSSRuleMatch, []*CSSPseudoElementMatches, []*CSSInheritedStyleEntry, []*CSSCSSKeyframesRule, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMatchedStylesForNode", Params: v})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			InlineStyle       *CSSCSSStyle
			AttributesStyle   *CSSCSSStyle
			MatchedCSSRules   []*CSSRuleMatch
			PseudoElements    []*CSSPseudoElementMatches
			Inherited         []*CSSInheritedStyleEntry
			CssKeyframesRules []*CSSCSSKeyframesRule
		}
	}

	if resp == nil {
		return nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, chromeData.Result.MatchedCSSRules, chromeData.Result.PseudoElements, chromeData.Result.Inherited, chromeData.Result.CssKeyframesRules, nil
}

// GetMatchedStylesForNode - Returns requested styles for a DOM node identified by `nodeId`.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%"). matchedCSSRules - CSS rules matching this node, from all applicable stylesheets. pseudoElements - Pseudo style matches for this node. inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root). cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (c *CSS) GetMatchedStylesForNode(nodeId int) (*CSSCSSStyle, *CSSCSSStyle, []*CSSRuleMatch, []*CSSPseudoElementMatches, []*CSSInheritedStyleEntry, []*CSSCSSKeyframesRule, error) {
	var v CSSGetMatchedStylesForNodeParams
	v.NodeId = nodeId
	return c.GetMatchedStylesForNodeWithParams(&v)
}

// GetMediaQueries - Returns all media queries parsed by the rendering engine.
// Returns -  medias -
func (c *CSS) GetMediaQueries() ([]*CSSCSSMedia, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMediaQueries"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Medias []*CSSCSSMedia
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

	return chromeData.Result.Medias, nil
}

type CSSGetPlatformFontsForNodeParams struct {
	//
	NodeId int `json:"nodeId"`
}

// GetPlatformFontsForNodeWithParams - Requests information about platform fonts which we used to render child TextNodes in the given node.
// Returns -  fonts - Usage statistics for every employed platform font.
func (c *CSS) GetPlatformFontsForNodeWithParams(v *CSSGetPlatformFontsForNodeParams) ([]*CSSPlatformFontUsage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getPlatformFontsForNode", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Fonts []*CSSPlatformFontUsage
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

	return chromeData.Result.Fonts, nil
}

// GetPlatformFontsForNode - Requests information about platform fonts which we used to render child TextNodes in the given node.
// nodeId -
// Returns -  fonts - Usage statistics for every employed platform font.
func (c *CSS) GetPlatformFontsForNode(nodeId int) ([]*CSSPlatformFontUsage, error) {
	var v CSSGetPlatformFontsForNodeParams
	v.NodeId = nodeId
	return c.GetPlatformFontsForNodeWithParams(&v)
}

type CSSGetStyleSheetTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
}

// GetStyleSheetTextWithParams - Returns the current textual content for a stylesheet.
// Returns -  text - The stylesheet text.
func (c *CSS) GetStyleSheetTextWithParams(v *CSSGetStyleSheetTextParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getStyleSheetText", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Text string
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

	return chromeData.Result.Text, nil
}

// GetStyleSheetText - Returns the current textual content for a stylesheet.
// styleSheetId -
// Returns -  text - The stylesheet text.
func (c *CSS) GetStyleSheetText(styleSheetId string) (string, error) {
	var v CSSGetStyleSheetTextParams
	v.StyleSheetId = styleSheetId
	return c.GetStyleSheetTextWithParams(&v)
}

type CSSSetEffectivePropertyValueForNodeParams struct {
	// The element id for which to set property.
	NodeId int `json:"nodeId"`
	//
	PropertyName string `json:"propertyName"`
	//
	Value string `json:"value"`
}

// SetEffectivePropertyValueForNodeWithParams - Find a rule with the given active property for the given node and set the new value for this property
func (c *CSS) SetEffectivePropertyValueForNodeWithParams(v *CSSSetEffectivePropertyValueForNodeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setEffectivePropertyValueForNode", Params: v})
}

// SetEffectivePropertyValueForNode - Find a rule with the given active property for the given node and set the new value for this property
// nodeId - The element id for which to set property.
// propertyName -
// value -
func (c *CSS) SetEffectivePropertyValueForNode(nodeId int, propertyName string, value string) (*gcdmessage.ChromeResponse, error) {
	var v CSSSetEffectivePropertyValueForNodeParams
	v.NodeId = nodeId
	v.PropertyName = propertyName
	v.Value = value
	return c.SetEffectivePropertyValueForNodeWithParams(&v)
}

type CSSSetKeyframeKeyParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	KeyText string `json:"keyText"`
}

// SetKeyframeKeyWithParams - Modifies the keyframe rule key text.
// Returns -  keyText - The resulting key text after modification.
func (c *CSS) SetKeyframeKeyWithParams(v *CSSSetKeyframeKeyParams) (*CSSValue, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setKeyframeKey", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			KeyText *CSSValue
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

	return chromeData.Result.KeyText, nil
}

// SetKeyframeKey - Modifies the keyframe rule key text.
// styleSheetId -
// range -
// keyText -
// Returns -  keyText - The resulting key text after modification.
func (c *CSS) SetKeyframeKey(styleSheetId string, theRange *CSSSourceRange, keyText string) (*CSSValue, error) {
	var v CSSSetKeyframeKeyParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.KeyText = keyText
	return c.SetKeyframeKeyWithParams(&v)
}

type CSSSetMediaTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	Text string `json:"text"`
}

// SetMediaTextWithParams - Modifies the rule selector.
// Returns -  media - The resulting CSS media rule after modification.
func (c *CSS) SetMediaTextWithParams(v *CSSSetMediaTextParams) (*CSSCSSMedia, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setMediaText", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Media *CSSCSSMedia
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

	return chromeData.Result.Media, nil
}

// SetMediaText - Modifies the rule selector.
// styleSheetId -
// range -
// text -
// Returns -  media - The resulting CSS media rule after modification.
func (c *CSS) SetMediaText(styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSMedia, error) {
	var v CSSSetMediaTextParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.Text = text
	return c.SetMediaTextWithParams(&v)
}

type CSSSetRuleSelectorParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	TheRange *CSSSourceRange `json:"range"`
	//
	Selector string `json:"selector"`
}

// SetRuleSelectorWithParams - Modifies the rule selector.
// Returns -  selectorList - The resulting selector list after modification.
func (c *CSS) SetRuleSelectorWithParams(v *CSSSetRuleSelectorParams) (*CSSSelectorList, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setRuleSelector", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			SelectorList *CSSSelectorList
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

	return chromeData.Result.SelectorList, nil
}

// SetRuleSelector - Modifies the rule selector.
// styleSheetId -
// range -
// selector -
// Returns -  selectorList - The resulting selector list after modification.
func (c *CSS) SetRuleSelector(styleSheetId string, theRange *CSSSourceRange, selector string) (*CSSSelectorList, error) {
	var v CSSSetRuleSelectorParams
	v.StyleSheetId = styleSheetId
	v.TheRange = theRange
	v.Selector = selector
	return c.SetRuleSelectorWithParams(&v)
}

type CSSSetStyleSheetTextParams struct {
	//
	StyleSheetId string `json:"styleSheetId"`
	//
	Text string `json:"text"`
}

// SetStyleSheetTextWithParams - Sets the new stylesheet text.
// Returns -  sourceMapURL - URL of source map associated with script (if any).
func (c *CSS) SetStyleSheetTextWithParams(v *CSSSetStyleSheetTextParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleSheetText", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SourceMapURL string
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

	return chromeData.Result.SourceMapURL, nil
}

// SetStyleSheetText - Sets the new stylesheet text.
// styleSheetId -
// text -
// Returns -  sourceMapURL - URL of source map associated with script (if any).
func (c *CSS) SetStyleSheetText(styleSheetId string, text string) (string, error) {
	var v CSSSetStyleSheetTextParams
	v.StyleSheetId = styleSheetId
	v.Text = text
	return c.SetStyleSheetTextWithParams(&v)
}

type CSSSetStyleTextsParams struct {
	//
	Edits []*CSSStyleDeclarationEdit `json:"edits"`
}

// SetStyleTextsWithParams - Applies specified style edits one after another in the given order.
// Returns -  styles - The resulting styles after modification.
func (c *CSS) SetStyleTextsWithParams(v *CSSSetStyleTextsParams) ([]*CSSCSSStyle, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleTexts", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Styles []*CSSCSSStyle
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

	return chromeData.Result.Styles, nil
}

// SetStyleTexts - Applies specified style edits one after another in the given order.
// edits -
// Returns -  styles - The resulting styles after modification.
func (c *CSS) SetStyleTexts(edits []*CSSStyleDeclarationEdit) ([]*CSSCSSStyle, error) {
	var v CSSSetStyleTextsParams
	v.Edits = edits
	return c.SetStyleTextsWithParams(&v)
}

// Enables the selector recording.
func (c *CSS) StartRuleUsageTracking() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.startRuleUsageTracking"})
}

// StopRuleUsageTracking - Stop tracking rule usage and return the list of rules that were used since last call to `takeCoverageDelta` (or since start of coverage instrumentation)
// Returns -  ruleUsage -
func (c *CSS) StopRuleUsageTracking() ([]*CSSRuleUsage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.stopRuleUsageTracking"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			RuleUsage []*CSSRuleUsage
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

	return chromeData.Result.RuleUsage, nil
}

// TakeCoverageDelta - Obtain list of rules that became used since last call to this method (or since start of coverage instrumentation)
// Returns -  coverage -
func (c *CSS) TakeCoverageDelta() ([]*CSSRuleUsage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.takeCoverageDelta"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Coverage []*CSSRuleUsage
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

	return chromeData.Result.Coverage, nil
}
