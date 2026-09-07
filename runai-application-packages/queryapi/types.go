package queryapi


type APIFieldsMapping map[string]FieldMapping

// MappingFromColumns mirrors dao-utils CreateFieldConfigMap by applying UserInputValidator to every field.
func MappingFromColumns(fieldToColumn map[string]string) APIFieldsMapping {
	mapping := make(APIFieldsMapping, len(fieldToColumn))
	for field, column := range fieldToColumn {
		mapping[field] = AlwaysValidField(column)
	}
	return mapping
}

func AlwaysValidField(column string) FieldMapping {
	return FieldMapping{
		Column:    column,
		Validator: UserInputValidator{},
	}
}

type FieldMapping struct {
	Column    string
	Validator Validator
}

type Filter struct {
	Field       string
	Operator    string
	ValueFormat string
	Values      []string
}

type Sort struct {
	Field string
	Direction string
}

type operatorInfo struct {
	Operator    string
	ValueFormat string
}