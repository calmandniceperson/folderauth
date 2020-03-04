// Author(s): Michael Koeppl

package folderauth

import (
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) assignRole(roleName, sid, route string) error {
	j := fmt.Sprintf("roleName=%s&sid=%s", roleName, sid)

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

// AssignSIDToGlobalRole uses the /folder-auth/assignSidToGlobalRole route
// to assign the user with the given SID to a global role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#assignsidtoglobalrole
// for additional details.
func (c *Client) AssignSIDToGlobalRole(roleName, sid string) error {
	return c.assignRole(roleName, sid, "assignSidToGlobalRole")
}

// AssignSIDToFolderRole uses the /folder-auth/assignSidToGlobalRole route
// to assign the user with the given SID to a folder role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#assignsidtofolderrole
// for additional details.
func (c *Client) AssignSIDToFolderRole(roleName, sid string) error {
	return c.assignRole(roleName, sid, "assignSidToFolderRole")
}

// AssignSIDToAgentRole uses the /folder-auth/assignSidToAgentRole route
// to assign the user with the given SID to an agent role.
// See https://github.com/jenkinsci/folder-auth-plugin/blob/master/docs/rest-api.adoc#assignsidtoagentrole
// for additional details.
func (c *Client) AssignSIDToAgentRole(roleName, sid string) error {
	return c.assignRole(roleName, sid, "assignSidToAgentRole")
}
