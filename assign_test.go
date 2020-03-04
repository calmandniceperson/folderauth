// Author(s): Michael Koeppl

package folderauth_test

import (
	"os"
	"testing"

	"github.com/calmandniceperson/folderauth"
	"github.com/stretchr/testify/assert"
)

func TestAssignSIDToGlobalRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "assignTestGlobalRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.AddGlobalRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
	)
	if err != nil {
		t.Error("Failed to create global role for test")
		return
	}

	err = c.AssignSIDToGlobalRole(roleName, "admin")
	assert.NoError(t, err)

	err = c.AssignSIDToGlobalRole("lel", "admin")
	assert.Error(t, err)
}

func TestAssignSIDToFolderRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "assignTestFolderRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.AddFolderRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
		[]string{
			"testAssignFolder",
		},
	)
	if err != nil {
		t.Error("Failed to create folder role for test")
		return
	}

	err = c.AssignSIDToFolderRole(roleName, "admin")
	assert.NoError(t, err)

	err = c.AssignSIDToFolderRole("lel", "admin")
	assert.Error(t, err)
}

func TestAssignSIDToAgentRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "assignTestAgentRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.AddAgentRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
		[]string{
			"testAssignAgent",
		},
	)
	if err != nil {
		t.Error("Failed to create agent role for test")
		return
	}

	err = c.AssignSIDToAgentRole(roleName, "admin")
	assert.NoError(t, err)

	err = c.AssignSIDToAgentRole("lel", "admin")
	assert.Error(t, err)
}
