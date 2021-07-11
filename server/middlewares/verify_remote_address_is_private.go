package middlewares

import (
	"net"
	"net/http"
)

func PrivateAddressPool() []net.IPNet {
	pool := []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	}

	var privateIPBlocks []net.IPNet
	for _, cidr := range pool {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, *block)
	}

	return privateIPBlocks
}

func VerifyRemoteAddressIsPrivate(privateAddressPool []net.IPNet) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			if !isPrivateIP(net.ParseIP(ip), privateAddressPool) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func isPrivateIP(ip net.IP, blocks []net.IPNet) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range blocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
