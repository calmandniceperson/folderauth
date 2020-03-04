// Author(s): Michael Koeppl

package folderauth

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Client) addRole(data map[string]interface{}, route string) error {
	j, _ := json.Marshal(data)

	_, err := c.performRequest(http.MethodPost, route, bytes.NewReader(j), "application/json")
	if err != nil {
		return err
	}

	return nil
}

// AddGlobalRole uses the /folder-auth/addGlobalRole route to create a new
// global role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#addglobalrole
// for additional details.
func (c *Client) AddGlobalRole(name string, permissions []string) error {
	data := map[string]interface{}{
		"name":        name,
		"permissions": permissions,
	}

	return c.addRole(data, "addGlobalRole")
}

// AddFolderRole uses the /folder-auth/addFolderRole route to create a new
// folder role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#addfolderrole
// for additional details.
func (c *Client) AddFolderRole(name string, permissions []string, folderNames []string) error {
	data := map[string]interface{}{
		"name":        name,
		"permissions": permissions,
		"folderNames": folderNames,
	}

	return c.addRole(data, "addFolderRole")
}

// AddAgentRole uses the /folder-auth/addAgentRole route to create a new
// agent role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#addagentrole
// for additional details.
func (c *Client) AddAgentRole(name string, permissions []string, agentNames []string) error {
	data := map[string]interface{}{
		"name":        name,
		"permissions": permissions,
		"agentNames":  agentNames,
	}

	return c.addRole(data, "addAgentRole")
}
