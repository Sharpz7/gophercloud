//go:build acceptance || identity || projects

package v3

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/v2/internal/acceptance/clients"
	"github.com/gophercloud/gophercloud/v2/internal/acceptance/tools"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
)

func TestProjectsListAvailable(t *testing.T) {
	clients.RequireNonAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	allPages, err := projects.ListAvailable(client).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	for _, project := range allProjects {
		tools.PrintResource(t, project)
	}
}

func TestProjectsList(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	var iTrue = true
	listOpts := projects.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == "admin" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	listOpts.Filters = map[string]string{
		"name__contains": "dmi",
	}

	allPages, err = projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err = projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	found = false
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == "admin" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	listOpts.Filters = map[string]string{
		"name__contains": "foo",
	}

	allPages, err = projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err = projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	found = false
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == "admin" {
			found = true
		}
	}

	th.AssertEquals(t, found, false)
}

func TestProjectsGet(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	allPages, err := projects.List(client, nil).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	project := allProjects[0]
	p, err := projects.Get(context.TODO(), client, project.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get project: %v", err)
	}

	tools.PrintResource(t, p)

	th.AssertEquals(t, project.Name, p.Name)
}

func TestProjectsCRUD(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	project, err := CreateProject(t, client, nil)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, project.ID)

	tools.PrintResource(t, project)

	description := ""
	iFalse := false
	updateOpts := projects.UpdateOpts{
		Description: &description,
		Enabled:     &iFalse,
	}

	updatedProject, err := projects.Update(context.TODO(), client, project.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedProject)
	th.AssertEquals(t, updatedProject.Description, description)
	th.AssertEquals(t, updatedProject.Enabled, iFalse)
}

func TestProjectsDomain(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	var iTrue = true
	createOpts := projects.CreateOpts{
		IsDomain: &iTrue,
	}

	projectDomain, err := CreateProject(t, client, &createOpts)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, projectDomain.ID)

	tools.PrintResource(t, projectDomain)

	createOpts = projects.CreateOpts{
		DomainID: projectDomain.ID,
	}

	project, err := CreateProject(t, client, &createOpts)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, project.ID)

	tools.PrintResource(t, project)

	th.AssertEquals(t, project.DomainID, projectDomain.ID)

	var iFalse = false
	updateOpts := projects.UpdateOpts{
		Enabled: &iFalse,
	}

	_, err = projects.Update(context.TODO(), client, projectDomain.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
}

func TestProjectsNested(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	projectMain, err := CreateProject(t, client, nil)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, projectMain.ID)

	tools.PrintResource(t, projectMain)

	createOpts := projects.CreateOpts{
		ParentID: projectMain.ID,
	}

	project, err := CreateProject(t, client, &createOpts)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, project.ID)

	tools.PrintResource(t, project)

	th.AssertEquals(t, project.ParentID, projectMain.ID)
}

func TestProjectsTags(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	createOpts := projects.CreateOpts{
		Tags: []string{"Tag1", "Tag2"},
	}

	projectMain, err := CreateProject(t, client, &createOpts)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, projectMain.ID)

	// Search using all tags
	listOpts := projects.ListOpts{
		Tags: "Tag1,Tag2",
	}

	allPages, err := projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	found := false
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == projectMain.Name {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	// Search using all tags, including a not existing one
	listOpts = projects.ListOpts{
		Tags: "Tag1,Tag2,Tag3",
	}

	allPages, err = projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err = projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(allProjects), 0)

	// Search matching at least one tag
	listOpts = projects.ListOpts{
		TagsAny: "Tag1,Tag2,Tag3",
	}

	allPages, err = projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err = projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	found = false
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == projectMain.Name {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	// Search not matching any single tag
	listOpts = projects.ListOpts{
		NotTagsAny: "Tag1",
	}

	allPages, err = projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err = projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	found = false
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == projectMain.Name {
			found = true
		}
	}

	th.AssertEquals(t, found, false)

	// Search matching not all tags
	listOpts = projects.ListOpts{
		NotTags: "Tag1,Tag2,Tag3",
	}

	allPages, err = projects.List(client, listOpts).AllPages(context.TODO())
	th.AssertNoErr(t, err)

	allProjects, err = projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	found = false
	for _, project := range allProjects {
		tools.PrintResource(t, project)

		if project.Name == "admin" {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	// Update the tags
	updateOpts := projects.UpdateOpts{
		Tags: &[]string{"Tag1"},
	}

	updatedProject, err := projects.Update(context.TODO(), client, projectMain.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedProject)
	th.AssertEquals(t, len(updatedProject.Tags), 1)
	th.AssertEquals(t, updatedProject.Tags[0], "Tag1")

	// Update the project, but not its tags
	description := "Test description"
	updateOpts = projects.UpdateOpts{
		Description: &description,
	}

	updatedProject, err = projects.Update(context.TODO(), client, projectMain.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedProject)
	th.AssertEquals(t, len(updatedProject.Tags), 1)
	th.AssertEquals(t, updatedProject.Tags[0], "Tag1")

	// Remove all Tags
	updateOpts = projects.UpdateOpts{
		Tags: &[]string{},
	}

	updatedProject, err = projects.Update(context.TODO(), client, projectMain.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, updatedProject)
	th.AssertEquals(t, len(updatedProject.Tags), 0)
}

func TestProjectsTagsCRUD(t *testing.T) {
	clients.RequireAdmin(t)

	client, err := clients.NewIdentityV3Client()
	th.AssertNoErr(t, err)

	createOpts := projects.CreateOpts{
		Tags: []string{"Tag1", "Tag2"},
	}

	projectMain, err := CreateProject(t, client, &createOpts)
	th.AssertNoErr(t, err)
	defer DeleteProject(t, client, projectMain.ID)

	projectTagsList, err := projects.ListTags(context.TODO(), client, projectMain.ID).Extract()
	tools.PrintResource(t, projectTagsList)
	th.AssertNoErr(t, err)

	modifyOpts := projects.ModifyTagsOpts{
		Tags: []string{"foo", "bar"},
	}
	projectTags, err := projects.ModifyTags(context.TODO(), client, projectMain.ID, modifyOpts).Extract()
	tools.PrintResource(t, projectTags)
	th.AssertNoErr(t, err)

	err = projects.DeleteTags(context.TODO(), client, projectMain.ID).ExtractErr()
	th.AssertNoErr(t, err)
}
