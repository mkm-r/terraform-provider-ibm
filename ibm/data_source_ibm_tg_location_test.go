package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMTransitGatewaysLocationDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_location.test_tg_location"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewaysLocationDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "billing_location"),
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "type"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewaysLocationDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_tg_location" "test_tg_location" {
		name = "us-south"
		}   `)
}