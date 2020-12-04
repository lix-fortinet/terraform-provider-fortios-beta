// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemSitTunnel_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemSitTunnel_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemSitTunnelConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemSitTunnelExists("fortios_system_sittunnel.trname"),
					resource.TestCheckResourceAttr("fortios_system_sittunnel.trname", "destination", "1.1.1.1"),
					resource.TestCheckResourceAttr("fortios_system_sittunnel.trname", "interface", "port2"),
					resource.TestCheckResourceAttr("fortios_system_sittunnel.trname", "ip6", "::/0"),
					resource.TestCheckResourceAttr("fortios_system_sittunnel.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_sittunnel.trname", "source", "2.2.2.2"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemSitTunnelExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemSitTunnel: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemSitTunnel is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemSitTunnel(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemSitTunnel: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemSitTunnel: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemSitTunnelDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_sittunnel" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemSitTunnel(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemSitTunnel %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemSitTunnelConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_sittunnel" "trname" {
  destination = "1.1.1.1"
  interface   = "port2"
  ip6         = "::/0"
  name        = "%[1]s"
  source      = "2.2.2.2"
}
`, name)
}