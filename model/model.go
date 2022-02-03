package model

type Subject struct {
	SubjectID           string `xml:"Subject_ID,attr" bson:"subjectId" yaml:"subjectId"`
	ParentRelationships struct {
		PreferredParent    []ParentRelationship `xml:"Preferred_Parent" bson:"preferredParents" yaml:"preferredParents"`
		NonPreferredParent []ParentRelationship `xml:"Non-Preferred_Parent" bson:"nonPreferredParents" yaml:"nonPreferredParents"`
	} `xml:"Parent_Relationships" bson:"parentRelationship" yaml:"parentRelationship"`
	NewNotes         []DescriptiveNote `yaml:"notes"` // new field for builder
	DescriptiveNotes DescriptiveNotes  `xml:"Descriptive_Notes" bson:"descriptiveNote" yaml:"descriptiveNotes"`
	RecordType       string            `xml:"Record_Type" bson:"recordType" yaml:"recordType"`
	MergedStatus     string            `xml:"Merged_Status" bson:"mergedStatus" yaml:"mergedStatus"`
	Hierarchy        string            `xml:"Hierarchy" bson:"hierarchy" yaml:"hierarchy"`
	SortOrder        string            `xml:"Sort_Order" bson:"sortOrder" yaml:"sortOrder"`
	Terms            struct {
		PreferredTerm    []Term `xml:"Preferred_Term" bson:"preferredTerms" yaml:"preferredTerms"`
		NonPreferredTerm []Term `xml:"Non-Preferred_Term" bson:"nonPreferredTerms" yaml:"nonPreferredTerms"`
	} `xml:"Terms" bson:"term" yaml:"term"`
	AssociativeRelationships struct {
		AssociativeRelationship []AssociativeRelationship `xml:"Associative_Relationship" bson:"associativeRelationships" yaml:"associativeRelationships"`
	} `xml:"Associative_Relationships" bson:"associativeRelationship" yaml:"associativeRelationship"`
	SubjectContributors struct {
		SubjectContributor []struct {
			ContributorID string `xml:"Contributor_id" bson:"contributorId" yaml:"contributorId"`
		} `xml:"Subject_Contributor" bson:"subjectContributors" yaml:"subjectContributors"`
	} `xml:"Subject_Contributors" bson:"subjectContributor" yaml:"subjectContributor"`
	SubjectSources struct {
		SubjectSource []struct {
			Source Source `xml:"Source" bson:"source" yaml:"source"`
		} `xml:"Subject_Source" bson:"subjectSources" yaml:"subjectSources"`
	} `xml:"Subject_Sources" bson:"subjectSource" yaml:"subjectSource"`
}

type ParentRelationship struct {
	ParentSubjectID  string `xml:"Parent_Subject_ID" bson:"parentSubjectId" yaml:"parentSubjectId"`
	RelationshipType string `xml:"Relationship_Type" bson:"relationshipType" yaml:"relationshipType"`
	HistoricFlag     string `xml:"Historic_Flag" bson:"historicFlag" yaml:"historicFlag"`
	ParentString     string `xml:"Parent_String" bson:"parentString" yaml:"text"` // use "text" for builder
	HierRelType      string `xml:"Hier_Rel_Type" bson:"hierRelType" yaml:"hierRelType"`
}

type DescriptiveNotes struct {
	DescriptiveNote []DescriptiveNote `xml:"Descriptive_Note" bson:"descriptiveNotes" yaml:"descriptiveNotes"`
}

type DescriptiveNote struct {
	NoteText         string `xml:"Note_Text" bson:"noteText" yaml:"text"` // use "text" for builder
	NoteLanguage     string `xml:"Note_Language" bson:"noteLanguage" yaml:"noteLanguage"`
	NoteContributors struct {
		NoteContributor []struct {
			ContributorID string `xml:"Contributor_id" bson:"contributorId" yaml:"contributorId"`
		} `xml:"Note_Contributor" bson:"noteContributors" yaml:"noteContributors"`
	} `xml:"Note_Contributors" bson:"noteContributor" yaml:"noteContributor"`
	NoteSources struct {
		NoteSource []struct {
			Source Source `xml:"Source" bson:"source" yaml:"source"`
		} `xml:"Note_Source" bson:"noteSources" yaml:"noteSources"`
	} `xml:"Note_Sources" bson:"noteSource" yaml:"noteSource"`
}

type AssociativeRelationship struct {
	RelationshipType string `xml:"Relationship_Type" bson:"relationshipType" yaml:"relationshipType"`
	RelatedSubjectID struct {
		VPSubjectID string `xml:"VP_Subject_ID" bson:"vpSubjectId" yaml:"vpSubjectId"`
	} `xml:"Related_Subject_ID" bson:"relatedSubjectId" yaml:"relatedSubjectId"`
	HistoricFlag string `xml:"Historic_Flag" bson:"historicFlag" yaml:"historicFlag"`
}

type Term struct {
	TermText      string `xml:"Term_Text" bson:"termText" yaml:"text"` // use "text" for builder
	DisplayName   string `xml:"Display_Name" bson:"displayName" yaml:"displayName"`
	HistoricFlag  string `xml:"Historic_Flag" bson:"historicFlag" yaml:"historicFlag"`
	Vernacular    string `xml:"Vernacular" bson:"vernacular" yaml:"vernacular"`
	TermID        string `xml:"Term_ID" bson:"termId" yaml:"termId"`
	TermLanguages struct {
		TermLanguage []struct {
			Language     string `xml:"Language" bson:"language" yaml:"language"`
			Preferred    string `xml:"Preferred" bson:"preferred" yaml:"preferred"`
			Qualifier    string `xml:"Qualifier" bson:"qualifier" yaml:"qualifier"`
			TermType     string `xml:"Term_Type" bson:"termType" yaml:"termType"`
			PartOfSpeech string `xml:"Part_of_Speech" bson:"partOfSpeech" yaml:"partOfSpeech"`
			LangStat     string `xml:"Lang_Stat" bson:"langStat" yaml:"langStat"`
		} `xml:"Term_Language" bson:"termLanguages" yaml:"termLanguages"`
	} `xml:"Term_Languages" bson:"termLanguage" yaml:"termLanguage"`
	TermContributors struct {
		TermContributor []struct {
			ContributorID string `xml:"Contributor_id" bson:"contributorId" yaml:"contributorId"`
			Preferred     string `xml:"Preferred" bson:"preferred" yaml:"preferred"`
		} `xml:"Term_Contributor" bson:"termContributors" yaml:"termContributors"`
	} `xml:"Term_Contributors" bson:"termContributor" yaml:"termContributor"`
	TermSources struct {
		TermSource []struct {
			Source    Source `xml:"Source" bson:"source" yaml:"source"`
			Page      string `xml:"Page" bson:"page" yaml:"page"`
			Preferred string `xml:"Preferred" bson:"preferred" yaml:"preferred"`
		} `xml:"Term_Source" bson:"termSources" yaml:"termSources"`
	} `xml:"Term_Sources" bson:"termSource" yaml:"termSource"`
}

type Source struct {
	SourceID string `xml:"Source_ID" bson:"sourceId" yaml:"sourceId"`
}
