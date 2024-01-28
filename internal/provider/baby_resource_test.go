// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	rsName := "aztfteam_baby.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccBabyResourceConfig("neo"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rsName, "name", "neo"),
					resource.TestCheckResourceAttrWith(rsName, "strength", validSpecial),
					resource.TestCheckResourceAttrSet(rsName, "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      rsName,
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{"name", "strength"},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccBabyResourceConfig(name string) string {
	return fmt.Sprintf(`
provider "aztfteam" {}
resource "aztfteam_baby" "test" {
  name = "%s"
}
`, name)
}

func validSpecial(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if v < 10 || v > 15 {
		return fmt.Errorf("invalid special")
	}
	return nil
}
