	"github.com/mutagen-io/mutagen/pkg/stream"
	// symbolicLinkMode is the symbolic link mode being used.
	symbolicLinkMode SymbolicLinkMode
	// problems are the problems encountered during transition operations.
		return errors.Wrap(err, "unable to read symbolic link target")
	// If we're in portable symbolic link mode, then we need to normalize the
	// target coming from disk, because some systems (e.g. Windows) won't
	// round-trip the target correctly.
	if t.symbolicLinkMode == SymbolicLinkMode_SymbolicLinkModePortable {
		target, err = normalizeSymbolicLinkAndEnsurePortable(path, target)
		return errors.New("symbolic link target does not match expected")
	// Ensure that this request is valid for the current symbolic link mode.
	if t.symbolicLinkMode == SymbolicLinkMode_SymbolicLinkModeIgnore {
	// Ensure that the existing symbolic link hasn't been modified from what
	// we're expecting.
	// Remove the symbolic link.
		} else if entry.Kind == EntryKind_SymbolicLink {
		entryCopy := entry.Copy(true)
	} else if entry.Kind == EntryKind_SymbolicLink {
			t.recordProblem(path, errors.Wrap(err, "unable to remove symbolic link"))
	preemptableTemporary := stream.NewPreemptableWriter(
		temporary,
		t.cancelled,
		transitionCopyPreemptionInterval,
	)
		if copyErr == stream.ErrWritePreempted {
	// Verify that the symbolic link agrees with our symbolic link mode.
	if t.symbolicLinkMode == SymbolicLinkMode_SymbolicLinkModeIgnore {
	} else if t.symbolicLinkMode == SymbolicLinkMode_SymbolicLinkModePortable {
		if normalized, err := normalizeSymbolicLinkAndEnsurePortable(path, target.Target); err != nil || normalized != target.Target {
	created := target.Copy(false)
		} else if entry.Kind == EntryKind_SymbolicLink {
	} else if target.Kind == EntryKind_SymbolicLink {
			t.recordProblem(path, errors.Wrap(err, "unable to create symbolic link"))
	symbolicLinkMode SymbolicLinkMode,
		symbolicLinkMode:               symbolicLinkMode,
		// fails, then record the reduced entry and continue to the next
		// transition.
		// entry is nil, then this is a no-op.