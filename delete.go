// Author(s): Michael Koeppl

package folderauth

import (
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) deleteRole(roleName, route string) error {
	j := fmt.Sprintf("roleName=%s", roleName)

	_, err := c.performRequest(
		http.MethodPost,
		route,
		strings.NewReader(j),
		"application/x-www-form-urlencoded",
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteGlobalRole uses the /folder-auth/deleteGlobalRole route to delete a
// global role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#assignsidtoagentrole
// for additional details.
func (c *Client) DeleteGlobalRole(roleName string) error {
	return c.deleteRole(roleName, "deleteGlobalRole")
}

// DeleteFolderRole uses the /folder-auth/deleteFolderRole route to delete a
// folder role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#assignsidtoagentrole
// for additional details.
func (c *Client) DeleteFolderRole(roleName string) error {
	return c.deleteRole(roleName, "deleteFolderRole")
}

// DeleteAgentRole uses the /folder-auth/deleteAgentRole route to delete an
// agent role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#assignsidtoagentrole
// for additional details.
func (c *Client) DeleteAgentRole(roleName string) error {
	return c.deleteRole(roleName, "deleteAgentRole")
}
