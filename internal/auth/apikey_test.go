package auth

import (
	"strings"
	"testing"
)

// hostVariants spans the shapes a real machine reports: short, typical, an
// over-long name, an FQDN, and the fallback when os.Hostname fails.
var hostVariants = []string{
	"laptop",
	"sapman-mbp",
	"sapman-ThinkPad-X1-Carbon-Gen-13",
	"buildbox.internal.example.com",
	"",
	strings.Repeat("y", 300),
}

var userVariants = []string{
	"sapirm@nimbleway.com",
	"oauth-user@example.com",
	"omerm",
	"",
	"averyveryverylongusername.with.dots@corp.example.com",
	strings.Repeat("x", 300),
	"@leading-at-sign.com",
}

// TestCLIKeyNameLength guards the server's 50-character limit on key_name.
// Exceeding it makes key creation fail with a 400.
func TestCLIKeyNameLength(t *testing.T) {
	for _, host := range hostVariants {
		for _, user := range userVariants {
			name := cliKeyNameFor(user, host)
			if len(name) > maxKeyNameLen {
				t.Errorf("cliKeyNameFor(%q, %q) = %q (%d chars), exceeds limit of %d",
					user, host, name, len(name), maxKeyNameLen)
			}
			if !strings.HasPrefix(name, cliKeyPrefix) {
				t.Errorf("cliKeyNameFor(%q, %q) = %q, want prefix %q", user, host, name, cliKeyPrefix)
			}
		}
	}
}

// TestCLIKeyNameDoesNotOverTruncate guards a bug an earlier version had: once
// the name needed trimming it applied a fixed budget split, discarding
// characters it was still allowed to use. Anything that must be truncated
// should spend the whole budget.
func TestCLIKeyNameDoesNotOverTruncate(t *testing.T) {
	for _, host := range hostVariants {
		name := cliKeyNameFor(strings.Repeat("x", 300), host)
		if len(name) != maxKeyNameLen {
			t.Errorf("cliKeyNameFor(<300 chars>, %q) = %q (%d chars), want the full %d",
				host, name, len(name), maxKeyNameLen)
		}
	}
}

// TestCLIKeyNameDistinguishesUsers is the whole point of scoping the name:
// two people sharing an account must not match each other's keys.
func TestCLIKeyNameDistinguishesUsers(t *testing.T) {
	for _, host := range hostVariants {
		a := cliKeyNameFor("alice@nimbleway.com", host)
		b := cliKeyNameFor("bob@nimbleway.com", host)
		if a == b {
			t.Errorf("names collided for different users on host %q: %q", host, a)
		}
	}
}

// TestCLIKeyNameDistinguishesHosts covers the same person on two machines,
// who would otherwise revoke their own key on every login.
func TestCLIKeyNameDistinguishesHosts(t *testing.T) {
	a := cliKeyNameFor("alice@nimbleway.com", "laptop")
	b := cliKeyNameFor("alice@nimbleway.com", "desktop")
	if a == b {
		t.Errorf("names collided for different hosts: %q", a)
	}
}

// TestCLIKeyNameStable ensures the same user on the same machine reuses one
// name, so stale-key cleanup keeps working across logins.
func TestCLIKeyNameStable(t *testing.T) {
	first := cliKeyNameFor("alice@nimbleway.com", "laptop")
	second := cliKeyNameFor("alice@nimbleway.com", "laptop")
	if first != second {
		t.Errorf("name not stable: %q then %q", first, second)
	}
}

// TestCLIKeyNameDropsSharedSuffixes spends the budget on the parts that
// actually identify someone, not the domain everyone on the account shares.
func TestCLIKeyNameDropsSharedSuffixes(t *testing.T) {
	name := cliKeyNameFor("sapirm@nimbleway.com", "buildbox.internal.example.com")
	for _, unwanted := range []string{"nimbleway.com", "internal.example.com"} {
		if strings.Contains(name, unwanted) {
			t.Errorf("name = %q, want %q dropped", name, unwanted)
		}
	}
	if !strings.Contains(name, "sapirm") || !strings.Contains(name, "buildbox") {
		t.Errorf("name = %q, want it to keep both identifying parts", name)
	}
}

// TestCLIKeyNameKeepsNormalUsernameIntact ensures an ordinary username is
// never trimmed: only the hostname gives way when the two together overflow.
func TestCLIKeyNameKeepsNormalUsernameIntact(t *testing.T) {
	for _, host := range hostVariants {
		name := cliKeyNameFor("sapirm@nimbleway.com", host)
		if !strings.Contains(name, "(sapirm @ ") {
			t.Errorf("cliKeyNameFor(sapirm, %q) = %q, want the username kept whole", host, name)
		}
	}
}

// TestCLIKeyNameUsesRealHostname checks the exported wrapper actually consults
// os.Hostname rather than silently falling back.
func TestCLIKeyNameUsesRealHostname(t *testing.T) {
	if got, want := CLIKeyName("alice@nimbleway.com"), "CLI (alice @ "; !strings.HasPrefix(got, want) {
		t.Errorf("CLIKeyName = %q, want prefix %q", got, want)
	}
	if strings.Contains(CLIKeyName("alice@nimbleway.com"), "unknown-host") {
		t.Error("CLIKeyName fell back to unknown-host on a machine with a hostname")
	}
}
