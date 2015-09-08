package go4redis

import (
	"testing"
)

func DialAndFlush() (*Client, error) {
	c, err := Dial("localhost:6379")

	if err != nil {
		return nil, err
	}

	errf := c.FlushDB()

	if errf != nil {
		return nil, err
	}

	return c, nil
}

func TestLLEN(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("Expected no error while dialing and flushing but got %s", err)
	}

	l, err := c.llen("foo")

	if err != nil {
		t.Errorf("expected no error while llen command, but got %s", err)
	}

	if l != 0 {
		t.Errorf("expected llen to return zero for non existent keys but got %d", l)
	}

	total, err := c.lpush("foo", "1", "2", "3")

	if err != nil {
		t.Errorf("expected no error while lpush command, but got %s", err)
	}

	ll, err := c.llen("foo")

	if err != nil {
		t.Errorf("expected no error while lpush command, but got %s", err)
	}

	if ll != total {
		t.Errorf("expected llen to return %d, but got %d", total, ll)
	}
}

func TestLPUSH(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("Expected no error while dialing asnd flushing but got %s", err)
	}

	total, err := c.lpush("foo", "1", "2", "3")

	if err != nil {
		t.Error("expected no error while lpush but got %s", err)
	}

	if total != 3 {
		t.Error("expected lpush to return 3 but got %d", total)
	}
}

func TestLPUSHMultiple(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("expected no error while connecting and flushing db but got %s", err)
	}

	for expected := 1; expected < 4; expected++ {
		total, err := c.lpush("foo", "1")

		if err != nil {
			t.Error("expected no error while lpush but got %s", err)
		}

		if total != expected {
			t.Errorf("expected lpush to return %d but got %d", expected, total)
		}
	}

}

func TestLSET(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("expected no error while connecting and flushing db but got %s", err)
	}

	err = c.lset("foo", 0, "100") // lset on non-existent key

	if err == nil {
		t.Error("expected error while performing lset on non-existent key")
	}

	_, err = c.lpush("foo", "1")

	if err != nil {
		t.Errorf("expected no error while lpush but got %s", err)
	}

	err = c.lset("foo", 0, "100")

	if err != nil {
		t.Errorf("expected no error while lset but got %s", err)
	}

	value, err := c.lindex("foo", 0)

	if err != nil {
		t.Errorf("expected no error while lindex but got %s", err)
	}

	if value != "100" {
		t.Errorf("expected lindex to return 100  but got %s", value)
	}

}

func TestLPOP(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("expected no error while connecting and flushing db but got %s", err)
	}

	_, err = c.lpush("foo", "1", "2", "3")

	if err != nil {
		t.Errorf("expected no error while lpush but got %s", err)
	}

	value, err := c.lpop("foo")

	if err != nil {
		t.Errorf("expected no error while lpop but got %s", err)
	}

	if value != "3" {
		t.Errorf("expected lpop to return 3  but got %s", value)
	}
}

func TestRPOP(t *testing.T) {
	c, err := DialAndFlush()

	if err != nil {
		t.Errorf("expected no error while connecting and flushing db but got %s", err)
	}

	_, err = c.rpush("foo", "1", "2", "3")

	if err != nil {
		t.Errorf("expected no error while rpush but got %s", err)
	}

	value, err := c.rpop("foo")

	if err != nil {
		t.Errorf("expected no error while rpop but got %s", err)
	}

	if value != "3" {
		t.Errorf("expected rpop to return 3  but got %s", value)
	}
}

func TestLPUSHX(t *testing.T) {
	t.Skip()
}

func TestRPUSHX(t *testing.T) {
	t.Skip()
}

func TestBLPOP(t *testing.T) {
	t.Skip()
}

func TestBRPOP(t *testing.T) {
	t.Skip()
}

func TestBRPOPLPUSH(t *testing.T) {
	t.Skip()
}

func TestLINSERT(t *testing.T) {
	t.Skip()
}

func TestLRANGE(t *testing.T) {
	t.Skip()
}

func TestLREM(t *testing.T) {
	t.Skip()
}

func TestLTRIM(t *testing.T) {
	t.Skip()
}
