package tfsec

import (
	"testing"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/checks"
)

func Test_GoogleUnencryptedStorageBucket(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.RuleID
		mustExcludeResultCode scanner.RuleID
	}{
		{
			name: "check google_storage_bucket with no encryption block",
			source: `
resource "google_storage_bucket" "my-bucket" {
	
}`,
			mustIncludeResultCode: checks.GoogleUnencryptedStorageBucket,
		},
		{
			name: "check google_storage_bucket with no encryption kms key name",
			source: `
resource "google_storage_bucket" "my-bucket" {
	encryption {}	
}`,
			mustExcludeResultCode: checks.GoogleUnencryptedStorageBucket,
		},
		{
			name: "check google_storage_bucket with non-empty encryption kms key name",
			source: `
resource "google_storage_bucket" "my-bucket" {
	encryption {
		default_kms_key_name = "my-key"
	}	
}`,
			mustExcludeResultCode: checks.GoogleUnencryptedStorageBucket,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
