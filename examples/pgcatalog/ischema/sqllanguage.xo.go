// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/go-goracle/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// SQLLanguage represents a row from 'information_schema.sql_languages'.
type SQLLanguage struct {
	Tableoid                       pgtypes.Oid           `json:"tableoid"`                          // tableoid
	Cmax                           pgtypes.Cid           `json:"cmax"`                              // cmax
	Xmax                           pgtypes.Xid           `json:"xmax"`                              // xmax
	Cmin                           pgtypes.Cid           `json:"cmin"`                              // cmin
	Xmin                           pgtypes.Xid           `json:"xmin"`                              // xmin
	Ctid                           pgtypes.Tid           `json:"ctid"`                              // ctid
	SQLLanguageSource              pgtypes.CharacterData `json:"sql_language_source"`               // sql_language_source
	SQLLanguageYear                pgtypes.CharacterData `json:"sql_language_year"`                 // sql_language_year
	SQLLanguageConformance         pgtypes.CharacterData `json:"sql_language_conformance"`          // sql_language_conformance
	SQLLanguageIntegrity           pgtypes.CharacterData `json:"sql_language_integrity"`            // sql_language_integrity
	SQLLanguageImplementation      pgtypes.CharacterData `json:"sql_language_implementation"`       // sql_language_implementation
	SQLLanguageBindingStyle        pgtypes.CharacterData `json:"sql_language_binding_style"`        // sql_language_binding_style
	SQLLanguageProgrammingLanguage pgtypes.CharacterData `json:"sql_language_programming_language"` // sql_language_programming_language
}
