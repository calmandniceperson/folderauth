// Author(s): Michael Koeppl

package folderauth_test

import (
	"os"
	"testing"

	"github.com/calmandniceperson/folderauth"
	"github.com/stretchr/testify/assert"
)

func TestAddGlobalRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "addTestGlobalRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.AddGlobalRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
	)
	assert.NoError(t, err)

	// Perform the same request again. This time, the plugin's API should
	// return an error.
	err = c.AddGlobalRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
	)
	assert.Error(t, err)
}

func TestAddFolderRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "addTestFolderRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.AddFolderRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
		[]string{
			"foo",
			"foo/bar",
		},
	)
	assert.NoError(t, err)

	// Perform the same request again. This time, the plugin's API should
	// return an error.
	err = c.AddFolderRole(
		roleName,
		[]string{
			"hudson.model.Item.Delete",
			"hudson.model.Item.Configure",
		},
		[]string{
			"foo",
			"foo/bar",
		},
	)
	assert.Error(t, err)
}

func TestAddAgentRole(t *testing.T) {
	hostname := os.Getenv("JENKINS_HOSTNAME")
	username := os.Getenv("JENKINS_USERNAME")
	token := os.Getenv("JENKINS_TOKEN")
	roleName := "addTestAgentRole"

	c, err := folderauth.NewClient(hostname, username, token)
	assert.NoError(t, err)

	err = c.AddAgentRole(
		roleName,
		[]string{
			"hudson.model.Computer.Configure",
		},
		[]string{
			"fooAgent",
			"barAgent",
		},
	)
	assert.NoError(t, err)

	// Perform the same request again. This time, the plugin's API should
	// return an error.
	err = c.AddAgentRole(
		roleName,
		[]string{
			"hudson.model.Computer.Configure",
		},
		[]string{
			"fooAgent",
			"barAgent",
		},
	)
	assert.Error(t, err)
}
