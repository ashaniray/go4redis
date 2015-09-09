package go4redis

import (
	"testing"
)

func TestDEL(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("Expected no error while dialing and flushing but got %s", err)
	}

	_, err = c.lpush("foo", "1")

	if err != nil {
		t.Errorf("expected no error while lpush command, but got %s", err)
	}

	ndel, err := c.Del("foo")

	if err != nil {
		t.Errorf("expected no error while del command, but got %s", err)
	}

	if ndel != 1 {
		t.Errorf("expected del command to return 1, but got %d", ndel)
	}

}

func TestEXISTS(t *testing.T) {
  c, err := DialAndFlush()

  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.lpush("foo", "1")

  if err != nil {
    t.Errorf("expected no error while lpush command, but got %s", err)
  }

  nexists, err := c.Exists("foo")

  if err != nil {
		t.Errorf("expected no error while exists command, but got %s", err)
	}

  if nexists != 1 {
		t.Errorf("expected exists command to return 1, but got %d", nexists)
	}

}

/*
func TestDUMP(t *testing.T)      { t.Fail() }

func TestEXPIRE(t *testing.T)    { t.Fail() }
func TestEXPIREAT(t *testing.T)  { t.Fail() }
func TestKEYS(t *testing.T)      { t.Fail() }
func TestMIGRATE(t *testing.T)   { t.Fail() }
func TestMOVE(t *testing.T)      { t.Fail() }
func TestOBJECT(t *testing.T)    { t.Fail() }
func TestPERSIST(t *testing.T)   { t.Fail() }
func TestPEXPIRE(t *testing.T)   { t.Fail() }
func TestPEXPIREAT(t *testing.T) { t.Fail() }
func TestPTTL(t *testing.T)      { t.Fail() }
func TestRANDOMKEY(t *testing.T) { t.Fail() }
func TestRENAME(t *testing.T)    { t.Fail() }
func TestRENAMENX(t *testing.T)  { t.Fail() }
func TestRESTORE(t *testing.T)   { t.Fail() }
func TestSORT(t *testing.T)      { t.Fail() }
func TestTTL(t *testing.T)       { t.Fail() }
func TestTYPE(t *testing.T)      { t.Fail() }
func TestWAIT(t *testing.T)      { t.Fail() }
func TestSCAN(t *testing.T)      { t.Fail() }
*/
