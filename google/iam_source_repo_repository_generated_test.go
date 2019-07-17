// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSourceRepoRepositoryIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/editor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSourceRepoRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceRepoRepositoryIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/editor", getTestProjectFromEnv(), fmt.Sprintf("my-repository-%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSourceRepoRepositoryIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/editor", getTestProjectFromEnv(), fmt.Sprintf("my-repository-%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSourceRepoRepositoryIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/editor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSourceRepoRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSourceRepoRepositoryIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/editor user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("my-repository-%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSourceRepoRepositoryIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/editor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSourceRepoRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceRepoRepositoryIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s", getTestProjectFromEnv(), fmt.Sprintf("my-repository-%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSourceRepoRepositoryIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my-repository-%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_member" "foo" {
	repository = "${google_sourcerepo_repository.my-repo.id}"
	role          = "%{role}"
	member        = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSourceRepoRepositoryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my-repository-%{random_suffix}"
}

data "google_iam_policy" "foo" {
	binding {
		role    = "%{role}"
		members = ["user:admin@hashicorptest.com"]
	}
}

resource "google_sourcerepo_repository_iam_policy" "foo" {
	repository = "${google_sourcerepo_repository.my-repo.id}"
	policy_data   = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccSourceRepoRepositoryIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my-repository-%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_binding" "foo" {
	repository = "${google_sourcerepo_repository.my-repo.id}"
	role          = "%{role}"
	members       = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSourceRepoRepositoryIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my-repository-%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_binding" "foo" {
	repository = "${google_sourcerepo_repository.my-repo.id}"
	role          = "%{role}"
	members       = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}