package main

import (
	"strings"
	"gopkg.in/guregu/null.v3"
	"errors"
)

type RuntimeOptions struct {
	Pagesize   null.String
	Margins    null.String
	Zoom       null.Int
	Aggressive null.Bool
}

var RuntimeOptionsPageSizeError error = errors.New("Invalid Pagesize! Use: A3, A4, A5, Legal, Letter or Tabloid")
var RuntimeOptionsMarginsError error = errors.New("Invalid Margins! Use: standard, none or minimal")

func (r *RuntimeOptions) Validate() error {

	if r.Pagesize.Valid {
		switch strings.ToLower(strings.TrimSpace(r.Pagesize.String)) {
		case "a3":
		case "a4":
		case "a5":
		case "legal":
		case "letter":
		case "tabloid":
		default:
			return RuntimeOptionsPageSizeError
		}
	}

	if r.Margins.Valid {
		switch strings.ToLower(strings.TrimSpace(r.Margins.String)) {
		case "standard":
		case "none":
		case "minimal":
		default:
			return RuntimeOptionsMarginsError
		}
	}

	return nil
}

func (r *RuntimeOptions) BuildCommand() []string {
	args := []string{}

	if r.Pagesize.Valid {
		args = append(args, "--pagesize")
		args = append(args, strings.Title(strings.TrimSpace(r.Pagesize.String)))
	}

	if r.Margins.Valid {
		args = append(args, "--margins")
		args = append(args, strings.ToLower(strings.TrimSpace(r.Margins.String)))
	}

	return args
}

