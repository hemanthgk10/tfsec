package tfsec

import (
	"testing"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/scanner"

	"github.com/hemanthgk10/tfsec/pkg/app/tfsec/checks"
)

func Test_AWSOpenSecurityGroupRule(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.RuleID
		mustExcludeResultCode scanner.RuleID
	}{
		{
			name: "check aws_security_group_rule ingress on 0.0.0.0/0",
			source: `
resource "aws_security_group_rule" "my-rule" {
	type = "ingress"
	cidr_blocks = ["0.0.0.0/0"]
}`,
			mustIncludeResultCode: checks.AWSOpenIngressSecurityGroupRule,
		},
		{
			name: "check aws_security_group_rule egress on 0.0.0.0/0",
			source: `
resource "aws_security_group_rule" "my-rule" {
	type = "egress"
	cidr_blocks = ["0.0.0.0/0"]
}`,
			mustIncludeResultCode: checks.AWSOpenEgressSecurityGroupRule,
		},
		{
			name: "check aws_security_group_rule egress on 0.0.0.0/0 in list",
			source: `
resource "aws_security_group_rule" "my-rule" {
	type = "egress"
	cidr_blocks = ["10.0.0.0/16", "0.0.0.0/0"]
}`,
			mustIncludeResultCode: checks.AWSOpenEgressSecurityGroupRule,
		},
		{
			name: "check aws_security_group_rule egress on 10.0.0.0/16",
			source: `
resource "aws_security_group_rule" "my-rule" {
	type = "egress"
	cidr_blocks = ["10.0.0.0/16"]
}`,
			mustExcludeResultCode: checks.AWSOpenEgressSecurityGroupRule,
		},
		{
			name: "check variable containing 0.0.0.0/0",
			source: `
resource "aws_security_group_rule" "github" {
  description = "HTTPS (GitHub)"
  type        = "ingress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = var.blocks

  security_group_id = aws_security_group.sg.id
}

variable "blocks" {
	default=["0.0.0.0/0"]
}

`,
			mustIncludeResultCode: checks.AWSOpenIngressSecurityGroupRule,
		},
		{
			name: "check aws_security_group_rule ingress on ::/0",
			source: `
resource "aws_security_group_rule" "my-rule" {
	type = "ingress"
	ipv6_cidr_blocks = ["::/0"]
}`,
			mustIncludeResultCode: checks.AWSOpenIngressSecurityGroupRule,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
