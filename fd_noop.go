//go:build !linux
// +build !linux

package spectator

func fdStats(s *sysStatsCollector) {
	// do nothing
}
