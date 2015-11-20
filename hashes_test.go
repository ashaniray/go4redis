package go4redis

import (
  "testing"
  "reflect"
)

func TestHDEL(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while HSet command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","77")
  if err != nil {
    t.Errorf("expected no error while HSet command, but got %s", err)
  }

  _, err = c.HSet("foo", "field3","78")
  if err != nil {
    t.Errorf("expected no error while HSet command, but got %s", err)
  }

  //Single field deletion
  ndel, err := c.HDel("foo","field")

  if err != nil {
    t.Errorf("expected no error while HDel command, but got %s", err)
  }

  if ndel != 1 {
    t.Errorf("expected del command to return 1, but got %d", ndel)
  }

  //Multiple field deletion

  ndel, err = c.HDel("foo","field2","field3","nofield")
  if err != nil {
    t.Errorf("expected no error while HDel command, but got %s", err)
  }

  if ndel != 2 {
    t.Errorf("expected del command to return 2, but got %d", ndel)
  }

}

func TestHEXISTS(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Field exists
  contains_field, err := c.HExists("foo","field")
  if err != nil {
    t.Errorf("expected no error while hexists command, but got %s", err)
  }
  if contains_field != 1 {
    t.Errorf("expected hexists command to return 1, but got %d", contains_field)
  }

  //Negative Test : Field does not exist
  contains_field, err = c.HExists("foo","nofield")
  if err != nil {
    t.Errorf("expected no error while hexists command, but got %s", err)
  }
  if contains_field != 0 {
    t.Errorf("expected hexists command to return 0, but got %d", contains_field)
  }

}

func TestHGETandHSET(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Field exists
  val, err := c.HGet("foo","field")
  if err != nil {
    t.Errorf("expected no error while hget command, but got %s", err)
  }
  if val != "76" {
    t.Errorf("expected hget command to return 76, but got %s", val)
  }

  //Negative Test : Field does not exist
  val, err = c.HGet("foo","nofield")
  if err != nil {
    t.Errorf("expected no error while hget command, but got %s", err)
  }
  if val != "-1" {
    t.Errorf("expected hget command to return -1, but got %s", val)
  }
}

func TestHGETALL(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Key exists
  array, err := c.HGetAll("foo")
  expected_array := []string{"field", "76", "field2", "abc"}
  if err != nil {
    t.Errorf("expected no error while hgetall command, but got %s", err)
  }
  if reflect.DeepEqual(array,expected_array) == false {
    t.Errorf("expected hgetall command to return 76, but got %v", array)
  }

  //Negative Test : Key does not exist
  array, err = c.HGetAll("key_not_present")
  var empty_array [] string
  if err != nil {
    t.Errorf("expected no error while hgetall command, but got %s", err)
  }
  if reflect.DeepEqual(array,empty_array) == false {
    t.Errorf("expected hgetall command to return empty_array, but got %v", array)
  }
}

func TestHINCRBY(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","1")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Integer value for the field exists
  val, err := c.HIncrBy("foo","field",5)
  if err != nil {
    t.Errorf("expected no error while hincrby command, but got %s", err)
  }
  if val != 6 {
    t.Errorf("expected hincrby command to return 6, but got %v", val)
  }

  val, err = c.HIncrBy("foo","field",-10)
  if err != nil {
    t.Errorf("expected no error while hincrby command, but got %s", err)
  }
  if val != -4 {
    t.Errorf("expected hincrby command to return -4, but got %v", val)
  }

  //Negative Tests

  //1.Integer value is not present for the field
  val, err = c.HIncrBy("foo","field2",-10)
  expected_error := "ERR hash value is not an integer"
  actual_error := err.Error()
  if err == nil {
    t.Errorf("expected error while integer value is no present while executing hincrby command")
  }  else if actual_error != expected_error {
    t.Errorf("expected : %s , but got : %s", expected_error, actual_error)
  }


  //2.key or Field is not present
  val, err = c.HIncrBy("key_not_present","nofield",-10)
  if err != nil {
    t.Errorf("expected no error while key or field not present while executing hincrby command, but got %s", err)
  }
  if val != -10 {
    t.Errorf("expected hincrby command to return -10, but got %d", val)
  }
}

func TestHINCRBYFLOAT(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","1")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Integer value for the field exists
  val, err := c.HIncrByFloat("foo","field","5.03")
  if err != nil {
    t.Errorf("expected no error while hincrbyfloat command, but got %s", err)
  }
  if val != "6.03" {
    t.Errorf("expected hincrbyfloat command to return 6.03, but got %v", val)
  }

  //Negative Tests

  //1.Float value is not present for the field
  val, err = c.HIncrByFloat("foo","field2","-10")
  expected_error := "ERR hash value is not a valid float"
  actual_error := err.Error()
  if err == nil {
    t.Errorf("expected error while float value is no present while executing hincrbyfloat command")
  }  else if actual_error != expected_error {
    t.Errorf("expected : %s , but got : %s", expected_error, actual_error)
  }

  //2.key or Field is not present
  val, err = c.HIncrByFloat("foo","nofield","-1e2")
  if err != nil {
    t.Errorf("expected no error while hincrbyfloat command, but got %s", err)
  }
  if val != "-100" {
    t.Errorf("expected hincrbyfloat command to return -3.97, but got %v", val)
  }
}

func TestHKEYS(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Key exists
  array, err := c.HKeys("foo")
  expected_array := []string{"field", "field2"}
  if err != nil {
    t.Errorf("expected no error while hgetall command, but got %s", err)
  }
  if reflect.DeepEqual(array,expected_array) == false {
    t.Errorf("expected hgetall command to return 76, but got %v", array)
  }

  //Negative Test : Key does not exist
  array, err = c.HKeys("key_not_present")
  var empty_array [] string
  if err != nil {
    t.Errorf("expected no error while hgetall command, but got %s", err)
  }
  if reflect.DeepEqual(array,empty_array) == false {
    t.Errorf("expected hgetall command to return empty_array, but got %v", array)
  }
}

func TestHLEN(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Key exists
  count, err := c.HLen("foo")
  if err != nil {
    t.Errorf("expected no error while hlen command, but got %s", err)
  }
  if count != 2 {
    t.Errorf("expected hlen command to return 2, but got %v", count)
  }

  //Negative Test : Key does not exist
  count, err = c.HLen("key_not_present")
  if err != nil {
    t.Errorf("expected no error while hlen command, but got %s", err)
  }
  if count != 0 {
    t.Errorf("expected hlen command to return 0, but got %v", count)
  }
}

func TestHMGET(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hmget command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hmget command, but got %s", err)
  }

  array, err := c.HMGet("foo","field","nofield","field2")
  expected_array := []string{"76","-1" ,"abc"}
  if err != nil {
    t.Errorf("expected no error while hmget command, but got %s", err)
  }
  if reflect.DeepEqual(array,expected_array) == false {
    t.Errorf("expected hmget command to return 76, but got %v", array)
  }
}

func TestHMSET(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HMSet("foo","field","76","field2","abc")
  if err != nil {
    t.Errorf("expected no error while hmset command, but got %s", err)
  }

  //Positive Test : Field exists
  val, err := c.HGet("foo","field")
  if err != nil {
    t.Errorf("expected no error while hget command, but got %s", err)
  }
  if val != "76" {
    t.Errorf("expected hget command to return 76, but got %s", val)
  }

  //Negative Test : Field does not exist
  val, err = c.HGet("foo","nofield")
  if err != nil {
    t.Errorf("expected no error while hget command, but got %s", err)
  }
  if val != "-1" {
    t.Errorf("expected hget command to return -1, but got %s", val)
  }
}

// HSETNX
func TestHSETNX(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  val, err := c.HSetNx("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }
  if val != 1 {
    t.Errorf("expected hget command to return 0, but got %d", val)
  }

  val, err = c.HSetNx("foo", "field","changed")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }
  if val != 0 {
    t.Errorf("expected hget command to return 0, but got %d", val)
  }

  str, err := c.HGet("foo","field")
  if err != nil {
    t.Errorf("expected no error while hget command, but got %s", err)
  }
  if str != "76" {
    t.Errorf("expected hget command to return 76, but got %s", val)
  }
}


func TestHVALS(t *testing.T) {
  c, err := DialAndFlush()
  if err != nil {
    t.Errorf("Expected no error while dialing and flushing but got %s", err)
  }

  _, err = c.HSet("foo", "field","76")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  _, err = c.HSet("foo", "field2","abc")
  if err != nil {
    t.Errorf("expected no error while hset command, but got %s", err)
  }

  //Positive Test : Key exists
  array, err := c.HVals("foo")
  expected_array := []string{"76", "abc"}
  if err != nil {
    t.Errorf("expected no error while hvals command, but got %s", err)
  }
  if reflect.DeepEqual(array,expected_array) == false {
    t.Errorf("expected hvals command to return 76, but got %v", array)
  }

  //Negative Test : Key does not exist
  array, err = c.HVals("key_not_present")
  var empty_array [] string
  if err != nil {
    t.Errorf("expected no error while hvals command, but got %s", err)
  }
  if reflect.DeepEqual(array,empty_array) == false {
    t.Errorf("expected hvals command to return empty_array, but got %v", array)
  }
}
