package main

const (
	queryMetadataNamespace   = "query_metadata_namespace"
	queryMetadataTaskType    = "query_metadata_task_type"
	queryMetadataOdinService = "odin_service"
)

type QueryMeta struct {
	namespace string
	taskType  string
	service   string
}

type QueryMetadataMatcher struct {
	namespaceMatcher *KeyValueMatcher
	taskTypeMatcher  *KeyValueMatcher
	serviceMatcher   *KeyValueMatcher
}

func NewQueryMetadataMatcher() *QueryMetadataMatcher {
	namespaceMatcher, _ := NewKeyValueMatcher(queryMetadataNamespace)
	taskTypeMatcher, _ := NewKeyValueMatcher(queryMetadataTaskType)
	serviceMatcher, _ := NewKeyValueMatcher(queryMetadataOdinService)
	return &QueryMetadataMatcher{namespaceMatcher: namespaceMatcher, taskTypeMatcher: taskTypeMatcher, serviceMatcher: serviceMatcher}
}

func (qm *QueryMetadataMatcher) MatchQuery(query string) *QueryMeta {
	return &QueryMeta{
		namespace: qm.namespaceMatcher.Match(query),
		taskType:  qm.taskTypeMatcher.Match(query),
		service:   qm.serviceMatcher.Match(query),
	}
}
