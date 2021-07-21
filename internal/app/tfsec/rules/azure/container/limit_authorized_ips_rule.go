package container

import (
	"fmt"

	"github.com/aquasecurity/tfsec/pkg/result"
	"github.com/aquasecurity/tfsec/pkg/severity"

	"github.com/aquasecurity/tfsec/pkg/provider"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/hclcontext"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"

	"github.com/aquasecurity/tfsec/pkg/rule"

	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
)


func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID:   "AZU008",
		Service:   "container",
		ShortCode: "limit-authorized-ips",
		Documentation: rule.RuleDocumentation{
			Summary:      "Ensure AKS has an API Server Authorized IP Ranges enabled",
			Impact:       "Any IP can interact with the API server",
			Resolution:   "Limit the access to the API server to a limited IP range",
			Explanation:  `
The API server is the central way to interact with and manage a cluster. To improve cluster security and minimize attacks, the API server should only be accessible from a limited set of IP address ranges.
`,
			BadExample:   `
resource "azurerm_kubernetes_cluster" "bad_example" {

}
`,
			GoodExample:  `
resource "azurerm_kubernetes_cluster" "good_example" {
    api_server_authorized_ip_ranges = [
		"1.2.3.4/32"
	]
}
`,
			Links: []string{
				"https://docs.microsoft.com/en-us/azure/aks/api-server-authorized-ip-ranges",
				"https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster.html#api_server_authorized_ip_ranges",
			},
		},
		Provider:        provider.AzureProvider,
		RequiredTypes:   []string{"resource"},
		RequiredLabels:  []string{"azurerm_kubernetes_cluster"},
		DefaultSeverity: severity.Critical,
		CheckFunc: func(set result.Set, resourceBlock block.Block, _ *hclcontext.Context) {

			if (resourceBlock.MissingChild("api_server_authorized_ip_ranges") ||
				resourceBlock.GetAttribute("api_server_authorized_ip_ranges").Value().LengthInt() < 1) &&
				(resourceBlock.MissingChild("private_cluster_enabled") ||
					resourceBlock.GetAttribute("private_cluster_enabled").IsFalse()) {
				{
					set.Add(
						result.New(resourceBlock).
							WithDescription(fmt.Sprintf("Resource '%s' defined without limited set of IP address ranges to the API server.", resourceBlock.FullName())).
							WithRange(resourceBlock.Range()),
					)
				}
			}
		},
	})
}