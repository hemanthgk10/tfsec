package tfsec

import (
	"testing"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/checks"
)

func Test_AWSOutdatedSSLPolicy(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.RuleID
		mustExcludeResultCode scanner.RuleID
	}{
		{
			name: "check aws_alb_listener with outdated policy",
			source: `
resource "aws_alb_listener" "my-resource" {
	ssl_policy = "ELBSecurityPolicy-TLS-1-1-2017-01"
	protocol = "HTTPS"
}`,
			mustIncludeResultCode: checks.AWSOutdatedSSLPolicy,
		},
		{
			name: "check aws_lb_listener with outdated policy",
			source: `
resource "aws_lb_listener" "my-resource" {
	ssl_policy = "ELBSecurityPolicy-TLS-1-1-2017-01"
	protocol = "HTTPS"
}`,
			mustIncludeResultCode: checks.AWSOutdatedSSLPolicy,
		},
		{
			name: "check aws_alb_listener with ok policy",
			source: `
resource "aws_alb_listener" "my-resource" {
	ssl_policy = "ELBSecurityPolicy-TLS-1-2-2017-01"
	protocol = "HTTPS"
}`,
			mustExcludeResultCode: checks.AWSOutdatedSSLPolicy,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
