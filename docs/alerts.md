# Cluster Network Addons Operator Alerts

### CnaoDown
**Summary:** CNAO pod is down..

**Description:** The total count of running CNAO operators is 0 (zero) for more than 5 minutes..

**Severity:** warning.

**Operator health impact:** warning.

**For:** 5m.

### KubeMacPoolDuplicateMacsFound
**Summary:** Duplicate macs found..

**Description:** There are {{ $value }} duplicate KubeMacPool MAC addresses..

**Severity:** warning.

**Operator health impact:** warning.

**For:** 5m.

### KubemacpoolDown
**Summary:** KubeMacpool is deployed by CNAO CR but KubeMacpool pod is down..

**Description:** The total count of running KubeMacPool manager pods is 0 (zero) for more than 5 minutes..

**Severity:** critical.

**Operator health impact:** critical.

**For:** 5m.

### NetworkAddonsConfigNotReady
**Summary:** CNAO CR NetworkAddonsConfig is not ready..

**Description:** There is no CNAO CR ready for more than 5 minutes..

**Severity:** warning.

**Operator health impact:** warning.

**For:** 5m.

## Developing new alerts

All alerts documented here are auto-generated and reflect exactly what is being
exposed. After developing new alerts or changing old ones please regenerate
this document.
