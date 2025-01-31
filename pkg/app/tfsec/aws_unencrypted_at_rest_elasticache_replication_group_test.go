package tfsec

import (
	"testing"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/checks"
)

func Test_AWSUnencryptedAtRestElasticacheReplicationGroup(t *testing.T) {
	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.RuleID
		mustExcludeResultCode scanner.RuleID
	}{
		{
			name: "check aws_elasticache_replication_group missing at_rest_encryption_enabled",
			source: `
resource "aws_elasticache_replication_group" "my-resource" {
        replication_group_id = "foo"
        replication_group_description = "my foo cluster"
}`,
			mustIncludeResultCode: checks.AWSUnencryptedAtRestElasticacheReplicationGroup,
		},
		{
			name: "check aws_elasticache_replication_group with at_rest_encryption_enabled",
			source: `
resource "aws_elasticache_replication_group" "my-resource" {
        replication_group_id = "foo"
        replication_group_description = "my foo cluster"

        at_rest_encryption_enabled = true
}`,
			mustExcludeResultCode: checks.AWSUnencryptedAtRestElasticacheReplicationGroup,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
