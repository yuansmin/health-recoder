package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuansmin/health-recoder/pkg/models"
)

func TestCreateGetDeleteUser(t *testing.T) {
	url := genURL("/api/users")
	user := models.User{Name: "function-test-user"}
	d, _ := json.Marshal(&user)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(d))
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, 201, resp.StatusCode, string(data))

	createdUser := models.User{}
	err = json.Unmarshal(data, &createdUser)
	assert.Nil(t, err, "unmarshal response err", string(data))
	assert.Equal(t, user.Name, createdUser.Name)

	// check if create user success, return if created failed
	if createdUser.ID == 0 {
		return
	}

	// test get user
	url = genURL(fmt.Sprintf("/api/users/%d", createdUser.ID))
	req, _ = http.NewRequest("GET", url, nil)
	resp, err = client.Do(req)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode, string(data))

	res := models.User{}
	err = json.Unmarshal(data, &res)
	assert.Nil(t, err, "unmarshal response err", string(data))
	assert.Equal(t, user.Name, res.Name)

	// test delete user
	url = genURL(fmt.Sprintf("/api/users/%d", createdUser.ID))
	req, _ = http.NewRequest("DELETE", url, nil)
	resp, err = client.Do(req)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode, string(data))

	// res = models.User{}
	// err = json.Unmarshal(data, &res)
	// assert.Nil(t, err, "unmarshal response err", string(data))
	// assert.Equal(t, user.Name, res.Name)

	// check if delete op success
	url = genURL(fmt.Sprintf("/api/users/%d", createdUser.ID))
	req, _ = http.NewRequest("get", url, nil)
	resp, err = client.Do(req)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)
	assert.NotEqual(t, 200, resp.StatusCode, string(data))

}
