package models

// Generated with gopkg.in/reform.v1. Do not edit by hand.

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type personTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *personTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("people").
func (v *personTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *personTableType) Columns() []string {
	return []string{"id", "group_id", "name", "email", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *personTableType) NewStruct() reform.Struct {
	return new(Person)
}

// NewRecord makes a new record for that table.
func (v *personTableType) NewRecord() reform.Record {
	return new(Person)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *personTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// PersonTable represents people view or table in SQL database.
var PersonTable = &personTableType{
	s: parse.StructInfo{Type: "Person", SQLSchema: "", SQLName: "people", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "GroupID", Type: "", Column: "group_id"}, {Name: "Name", Type: "", Column: "name"}, {Name: "Email", Type: "", Column: "email"}, {Name: "CreatedAt", Type: "", Column: "created_at"}, {Name: "UpdatedAt", Type: "", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(Person).Values(),
}

// String returns a string representation of this struct or record.
func (s Person) String() string {
	res := make([]string, 6)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "GroupID: " + reform.Inspect(s.GroupID, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "Email: " + reform.Inspect(s.Email, true)
	res[4] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[5] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Person) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.GroupID,
		s.Name,
		s.Email,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Person) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.GroupID,
		&s.Name,
		&s.Email,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Person) View() reform.View {
	return PersonTable
}

// Table returns Table object for that record.
func (s *Person) Table() reform.Table {
	return PersonTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Person) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Person) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Person) HasPK() bool {
	return s.ID != PersonTable.z[PersonTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Person) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = PersonTable
	_ reform.Struct = (*Person)(nil)
	_ reform.Table  = PersonTable
	_ reform.Record = (*Person)(nil)
	_ fmt.Stringer  = (*Person)(nil)
)

type projectTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *projectTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("projects").
func (v *projectTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *projectTableType) Columns() []string {
	return []string{"name", "id", "start", "end"}
}

// NewStruct makes a new struct for that view or table.
func (v *projectTableType) NewStruct() reform.Struct {
	return new(Project)
}

// NewRecord makes a new record for that table.
func (v *projectTableType) NewRecord() reform.Record {
	return new(Project)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *projectTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ProjectTable represents projects view or table in SQL database.
var ProjectTable = &projectTableType{
	s: parse.StructInfo{Type: "Project", SQLSchema: "", SQLName: "projects", Fields: []parse.FieldInfo{{Name: "Name", Type: "", Column: "name"}, {Name: "ID", Type: "string", Column: "id"}, {Name: "Start", Type: "", Column: "start"}, {Name: "End", Type: "", Column: "end"}}, PKFieldIndex: 1},
	z: new(Project).Values(),
}

// String returns a string representation of this struct or record.
func (s Project) String() string {
	res := make([]string, 4)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "ID: " + reform.Inspect(s.ID, true)
	res[2] = "Start: " + reform.Inspect(s.Start, true)
	res[3] = "End: " + reform.Inspect(s.End, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Project) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.ID,
		s.Start,
		s.End,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Project) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.ID,
		&s.Start,
		&s.End,
	}
}

// View returns View object for that struct.
func (s *Project) View() reform.View {
	return ProjectTable
}

// Table returns Table object for that record.
func (s *Project) Table() reform.Table {
	return ProjectTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Project) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Project) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Project) HasPK() bool {
	return s.ID != ProjectTable.z[ProjectTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Project) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = string(i64)
	} else {
		s.ID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = ProjectTable
	_ reform.Struct = (*Project)(nil)
	_ reform.Table  = ProjectTable
	_ reform.Record = (*Project)(nil)
	_ fmt.Stringer  = (*Project)(nil)
)

type personProjectViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *personProjectViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("person_project").
func (v *personProjectViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *personProjectViewType) Columns() []string {
	return []string{"person_id", "project_id"}
}

// NewStruct makes a new struct for that view or table.
func (v *personProjectViewType) NewStruct() reform.Struct {
	return new(PersonProject)
}

// PersonProjectView represents person_project view or table in SQL database.
var PersonProjectView = &personProjectViewType{
	s: parse.StructInfo{Type: "PersonProject", SQLSchema: "", SQLName: "person_project", Fields: []parse.FieldInfo{{Name: "PersonID", Type: "", Column: "person_id"}, {Name: "ProjectID", Type: "", Column: "project_id"}}, PKFieldIndex: -1},
	z: new(PersonProject).Values(),
}

// String returns a string representation of this struct or record.
func (s PersonProject) String() string {
	res := make([]string, 2)
	res[0] = "PersonID: " + reform.Inspect(s.PersonID, true)
	res[1] = "ProjectID: " + reform.Inspect(s.ProjectID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *PersonProject) Values() []interface{} {
	return []interface{}{
		s.PersonID,
		s.ProjectID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *PersonProject) Pointers() []interface{} {
	return []interface{}{
		&s.PersonID,
		&s.ProjectID,
	}
}

// View returns View object for that struct.
func (s *PersonProject) View() reform.View {
	return PersonProjectView
}

// check interfaces
var (
	_ reform.View   = PersonProjectView
	_ reform.Struct = (*PersonProject)(nil)
	_ fmt.Stringer  = (*PersonProject)(nil)
)

type legacyPersonTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("legacy").
func (v *legacyPersonTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("people").
func (v *legacyPersonTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *legacyPersonTableType) Columns() []string {
	return []string{"id", "name"}
}

// NewStruct makes a new struct for that view or table.
func (v *legacyPersonTableType) NewStruct() reform.Struct {
	return new(LegacyPerson)
}

// NewRecord makes a new record for that table.
func (v *legacyPersonTableType) NewRecord() reform.Record {
	return new(LegacyPerson)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *legacyPersonTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// LegacyPersonTable represents people view or table in SQL database.
var LegacyPersonTable = &legacyPersonTableType{
	s: parse.StructInfo{Type: "LegacyPerson", SQLSchema: "legacy", SQLName: "people", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Name", Type: "", Column: "name"}}, PKFieldIndex: 0},
	z: new(LegacyPerson).Values(),
}

// String returns a string representation of this struct or record.
func (s LegacyPerson) String() string {
	res := make([]string, 2)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *LegacyPerson) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *LegacyPerson) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *LegacyPerson) View() reform.View {
	return LegacyPersonTable
}

// Table returns Table object for that record.
func (s *LegacyPerson) Table() reform.Table {
	return LegacyPersonTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *LegacyPerson) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *LegacyPerson) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *LegacyPerson) HasPK() bool {
	return s.ID != LegacyPersonTable.z[LegacyPersonTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *LegacyPerson) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = LegacyPersonTable
	_ reform.Struct = (*LegacyPerson)(nil)
	_ reform.Table  = LegacyPersonTable
	_ reform.Record = (*LegacyPerson)(nil)
	_ fmt.Stringer  = (*LegacyPerson)(nil)
)

type iDOnlyTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *iDOnlyTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("id_only").
func (v *iDOnlyTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *iDOnlyTableType) Columns() []string {
	return []string{"id"}
}

// NewStruct makes a new struct for that view or table.
func (v *iDOnlyTableType) NewStruct() reform.Struct {
	return new(IDOnly)
}

// NewRecord makes a new record for that table.
func (v *iDOnlyTableType) NewRecord() reform.Record {
	return new(IDOnly)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *iDOnlyTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// IDOnlyTable represents id_only view or table in SQL database.
var IDOnlyTable = &iDOnlyTableType{
	s: parse.StructInfo{Type: "IDOnly", SQLSchema: "", SQLName: "id_only", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}}, PKFieldIndex: 0},
	z: new(IDOnly).Values(),
}

// String returns a string representation of this struct or record.
func (s IDOnly) String() string {
	res := make([]string, 1)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *IDOnly) Values() []interface{} {
	return []interface{}{
		s.ID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *IDOnly) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
	}
}

// View returns View object for that struct.
func (s *IDOnly) View() reform.View {
	return IDOnlyTable
}

// Table returns Table object for that record.
func (s *IDOnly) Table() reform.Table {
	return IDOnlyTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *IDOnly) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *IDOnly) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *IDOnly) HasPK() bool {
	return s.ID != IDOnlyTable.z[IDOnlyTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *IDOnly) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = IDOnlyTable
	_ reform.Struct = (*IDOnly)(nil)
	_ reform.Table  = IDOnlyTable
	_ reform.Record = (*IDOnly)(nil)
	_ fmt.Stringer  = (*IDOnly)(nil)
)

func init() {
	parse.AssertUpToDate(&PersonTable.s, new(Person))
	parse.AssertUpToDate(&ProjectTable.s, new(Project))
	parse.AssertUpToDate(&PersonProjectView.s, new(PersonProject))
	parse.AssertUpToDate(&LegacyPersonTable.s, new(LegacyPerson))
	parse.AssertUpToDate(&IDOnlyTable.s, new(IDOnly))
}
