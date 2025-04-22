package checkStruct

func NewService() *Service {
	return &Service{}
}

// Verify interface implementation
var _ DirectoryLister = (*Service)(nil)
var _ StructureComparer = (*Service)(nil)
