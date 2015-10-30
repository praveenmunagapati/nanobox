// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

//
package server

import (
	"github.com/nanobox-io/nanobox/config"
	"net/http"
	"time"
)

// Ping issues a ping to nanobox server
func Ping() (bool, error) {

	// a new client is used to allow for shortening the request timeout
	client := http.Client{Timeout: time.Duration(2 * time.Second)}

	//
	res, err := client.Get(config.ServerURL + "/ping")
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	//
	return res.StatusCode/100 == 2, nil
}
