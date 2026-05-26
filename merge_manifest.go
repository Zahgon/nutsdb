package nutsdb

type mergeManifestStatus string

const (
	mergeManifestFileName       = "merge_manifest.json"
	mergeManifestTempFileSuffix = ".tmp"

	manifestStatusWriting   mergeManifestStatus = "writing"
	manifestStatusCommitted mergeManifestStatus = "committed"
)

type mergeManifest struct {
	Status            mergeManifestStatus `json:"status"`
	MergeSeqMax       int                 `json:"mergeSeqMax"`
	PendingOldFileIDs []int64             `json:"pendingOldFileIDs"`
}

func loadMergeManifest(dir string) (*mergeManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeMergeManifest(dir string, manifest *mergeManifest) error {
	_ = "STUB: not implemented"
	return nil
}

func removeMergeManifest(dir string) error { _ = "STUB: not implemented"; return nil }
