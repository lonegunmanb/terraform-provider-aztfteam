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
					resource.TestCheckResourceAttrWith(rsName, "agility", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "charisma", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "endurance", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "luck", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "strength", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "perception", validSpecial),
					resource.TestCheckResourceAttrSet(rsName, "id"),
					resource.TestCheckResourceAttrSet(rsName, "intelligence"),
					resource.TestCheckResourceAttrSet(rsName, "birthday"),
					resource.TestCheckResourceAttrSet(rsName, "age"),
					resource.TestCheckResourceAttrSet(rsName, "id"),
					resource.TestCheckResourceAttrSet(rsName, "biological_gender"),
					resource.TestCheckResourceAttr(rsName, "tags.blessed_by", "terraform engineering China team"),
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
				ImportStateVerifyIgnore: []string{
					"age",
					"agility",
					"birthday",
					"charisma",
					"endurance",
					"intelligence",
					"luck",
					"name",
					"strength",
					"birthday",
					"age",
					"perception",
					"biological_gender",
					"tags",
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccExampleResource_withBirthday(t *testing.T) {
	rsName := "aztfteam_baby.test"
	birthday := "2024-02-23T20:09:00+08:00"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccBabyResourceConfig_withBirthday("neo", birthday),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rsName, "name", "neo"),
					resource.TestCheckResourceAttrWith(rsName, "agility", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "endurance", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "luck", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "strength", validSpecial),
					resource.TestCheckResourceAttrWith(rsName, "perception", validSpecial),
					resource.TestCheckResourceAttr(rsName, "birthday", birthday),
					resource.TestCheckResourceAttrSet(rsName, "id"),
					resource.TestCheckResourceAttrSet(rsName, "age"),
					resource.TestCheckResourceAttr(rsName, "tags.blessed_by", "terraform engineering China team"),
				),
			},
		},
	})
}

func testAccBabyResourceConfig_withBirthday(name string, birthday string) string {
	return fmt.Sprintf(`
provider "aztfteam" {}
resource "aztfteam_baby" "test" {
  name 	   = "%s"
  birthday = "%s"
}
`, name, birthday)
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
