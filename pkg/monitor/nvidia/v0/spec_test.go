package v0

import "testing"

func TestDeviceMemoryTotalUsesMonitorUsedWhenLedgerEmpty(t *testing.T) {
	var sr sharedRegionT
	sr.procs[0].monitorused[0] = 998244352
	sr.procs[1].monitorused[0] = 1024
	sr.procs[0].monitorused[1] = 4096

	got := (Spec{sr: &sr}).DeviceMemoryTotal(0)
	if got != 998244352+1024 {
		t.Fatalf("device 0 = %d, want monitorused sum", got)
	}
	if got := (Spec{sr: &sr}).DeviceMemoryTotal(1); got != 4096 {
		t.Fatalf("device 1 = %d, want 4096", got)
	}
}

func TestDeviceMemoryTotalPrefersHookLedgerPerProcess(t *testing.T) {
	var sr sharedRegionT
	sr.procs[0].used[0].total = 100
	sr.procs[0].monitorused[0] = 999
	sr.procs[1].monitorused[0] = 50

	got := (Spec{sr: &sr}).DeviceMemoryTotal(0)
	if got != 150 {
		t.Fatalf("got %d, want hook total of proc0 plus monitorused of proc1", got)
	}
}
