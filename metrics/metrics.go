// package metrics contains expvar & Prometheus types and code used by
// Tailscale for monitoring.

package metrics

import "expvar"

type Set struct {
	expvar.Map
}
