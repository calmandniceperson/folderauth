// Author(s): Michael Koeppl

package folderauth_test

import (
	"os"
	"testing"

	"github.com/calmandniceperson/folderauth"
	"github.com/stretchr/testify/assert"
)

func TestDeleteGlobalRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "deleteTestGlobalRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.DeleteGlobalRole(roleName)
	assert.Error(t, err)

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

	err = c.DeleteGlobalRole(roleName)
	assert.NoError(t, err)
}

func TestDeleteFolderRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "deleteTestFolderRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.DeleteFolderRole(roleName)
	assert.Error(t, err)

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

	err = c.DeleteFolderRole(roleName)
	assert.NoError(t, err)
}

func TestDeleteAgentRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "deleteTestAgentRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.DeleteAgentRole(roleName)
	assert.Error(t, err)

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

	err = c.DeleteAgentRole(roleName)
	assert.NoError(t, err)
}
