package xmlrpc

import (
    "testing"
    "time"
)

func Test_Client_Call(t *testing.T) {
    client, err := NewClient("http://localhost:5001")

    assert_nil(t, err)

    defer client.Close()

    var result = new(time.Time)
    if err = client.Call("bugzilla.time", nil, result); err != nil {
        t.Fatal(err)
    }

    assert_not_nil(t, result)
}