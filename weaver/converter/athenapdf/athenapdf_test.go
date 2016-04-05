package athenapdf

import (
	"github.com/arachnys/athenapdf/weaver/testutil"
	"reflect"
	"testing"
)

func TestConstructCMD(t *testing.T) {
	got := constructCMD("athenapdf -S -T 120", "test_file.html", false)
	want := []string{"athenapdf", "-S", "-T", "120", "test_file.html"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected constructed athenapdf command to be %+v, got %+v", want, got)
	}
}

func TestConstructCMD_aggressive(t *testing.T) {
	cmd := constructCMD("athenapdf -S -T 60", "test_file.html", true)
	if got, want := cmd[len(cmd)-1], "-A"; got != want {
		t.Errorf("expected last argument of constructed athenapdf command to be %s, got %+v", want, got)
	}
}

func mockConversion(path, cmd string) ([]byte, error) {
	c := AthenaPDF{}
	c.Path = path
	c.CMD = cmd
	t := make(chan struct{}, 1)
	return c.Convert(t)
}

func TestConvert(t *testing.T) {
	ts := testutil.MockHTTPServer("", "test AthenaPDF convert")
	defer ts.Close()
	got, err := mockConversion(ts.URL, "echo")
	if err != nil {
		t.Fatalf("convert returned an unexpected error: %+v", err)
	}
	if want := []byte(ts.URL + "\n"); !reflect.DeepEqual(got, want) {
		t.Errorf("expected output of athenapdf conversion to be %+v, got %+v", want, got)
	}
}

func TestConvert_badURL(t *testing.T) {
	got, err := mockConversion("bad URL", "echo")
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
	if got != nil {
		t.Errorf("expected output of athenapdf conversion to be nil, got %+v", got)
	}
}

func TestConvert_badCMD(t *testing.T) {
	ts := testutil.MockHTTPServer("", "test Athena convert")
	defer ts.Close()
	got, err := mockConversion(ts.URL, "echo-broken")
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
	if got != nil {
		t.Errorf("expected output of athenapdf conversion to be nil, got %+v", got)
	}
}

func TestConvert_octet(t *testing.T) {
	ts := testutil.MockHTTPServer("application/octet-stream", "test application/octet-stream bytes")
	defer ts.Close()
	got, err := mockConversion(ts.URL, "echo")
	if err != nil {
		t.Fatalf("convert returned an unexpected error: %+v", err)
	}
	if want := []byte(ts.URL + "\n"); reflect.DeepEqual(got, want) {
		t.Errorf("expected value should not be equal (new conversion path should be a local tmp file)")
	}
}
