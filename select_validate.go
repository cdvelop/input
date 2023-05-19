package input

// validación con datos de entrada
func (s selecTag) ValidateField(data_in string, skip_validation bool) bool {
	if !skip_validation {
		if data_in != "" {
			if _, exists := s.Data.SourceData()[data_in]; exists {
				return true
			}
		}
	} else {
		return true
	}
	return false
}
