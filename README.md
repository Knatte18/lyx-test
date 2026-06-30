# lyx-test

Fixture host repo for the Loomyard **weft** sandbox. Paired with `lyx-test-weft`.

Carries a `services/api/` subpath used to exercise **weft path mirroring** — how the
weft worktree mirrors host subpaths. This is *not* an invitation to run `lyx` from that
subdirectory: `lyx` resolves against the current directory's own `_lyx/` (it does not
walk up to a parent), so it is expected to work only from the directory where it was
initialized — this repo's root. Running `lyx` from an uninitialized subdirectory
correctly reports `not initialized here; run "lyx init"`; that is expected behaviour,
not a defect.
