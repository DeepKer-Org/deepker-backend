package enums

// RoleEnum represent the roles that a user can have
type RoleEnum string

const (
	Admin  RoleEnum = "admin"
	Tester RoleEnum = "tester"
	Doctor RoleEnum = "doctor"
	Nurse  RoleEnum = "nurse"
)

// ToStringArray converts a list of RoleEnum to a list of strings
func ToStringArray(roles ...RoleEnum) []string {
	var result []string
	for _, role := range roles {
		result = append(result, string(role))
	}
	return result
}

func MapStringsToRoleEnums(roles []string) []*RoleEnum {
	var result []*RoleEnum
	for _, role := range roles {
		r := RoleEnum(role)
		result = append(result, &r)
	}
	return result
}

func MapRoleEnumsToStrings(roles []*RoleEnum) []string {
	var result []string
	for _, role := range roles {
		result = append(result, string(*role))
	}
	return result
}
