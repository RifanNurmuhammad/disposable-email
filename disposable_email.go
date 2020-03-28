package disposableemail

import (
	"strings"
	"sync"
)

var domains = new(collection)

type collection struct {
	items map[string]struct{}
	err   error
	once  sync.Once
}

// IsDisabledEmail for split and validate email domain
func IsDisabledEmail(email string) bool {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return false
	}
	return IsDisabledDomain(parts[1])
}

// IsDisabledDomain for validate domain
func IsDisabledDomain(domain string) bool {
	domains.once.Do(func() { domains.loadDomainList() })
	if domains.err != nil {
		return false
	}
	domain = strings.TrimSpace(domain)
	return domains.hasValidDomain(strings.ToLower(domain))
}

func (c *collection) hasValidDomain(item string) bool {
	_, ok := c.items[item]
	return ok
}

func (c *collection) loadDomainList() {
	c.items = make(map[string]struct{})
	for _, value := range DisposableDomains {
		c.items[value] = struct{}{}
	}
}
