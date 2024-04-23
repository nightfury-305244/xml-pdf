package main

import "encoding/xml"

type ClinicalDocument struct {
	XMLName        xml.Name `xml:"ClinicalDocument" json:"clinicaldocument,omitempty"`
	Text           string   `xml:",chardata" json:"text,omitempty"`
	Xsi            string   `xml:"xsi,attr" json:"xsi,omitempty"`
	Xmlns          string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	SchemaLocation string   `xml:"schemaLocation,attr" json:"schemalocation,omitempty"`
	RealmCode      struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Code string `xml:"code,attr" json:"code,omitempty"`
	} `xml:"realmCode" json:"realmcode,omitempty"`
	TypeId struct {
		Text      string `xml:",chardata" json:"text,omitempty"`
		Root      string `xml:"root,attr" json:"root,omitempty"`
		Extension string `xml:"extension,attr" json:"extension,omitempty"`
	} `xml:"typeId" json:"typeid,omitempty"`
	TemplateId []struct {
		Text      string `xml:",chardata" json:"text,omitempty"`
		Root      string `xml:"root,attr" json:"root,omitempty"`
		Extension string `xml:"extension,attr" json:"extension,omitempty"`
	} `xml:"templateId" json:"templateid,omitempty"`
	ID struct {
		Text                   string `xml:",chardata" json:"text,omitempty"`
		Root                   string `xml:"root,attr" json:"root,omitempty"`
		Extension              string `xml:"extension,attr" json:"extension,omitempty"`
		AssigningAuthorityName string `xml:"assigningAuthorityName,attr" json:"assigningauthorityname,omitempty"`
	} `xml:"id" json:"id,omitempty"`
	Code struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		Code           string `xml:"code,attr" json:"code,omitempty"`
		CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
		CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
		DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
	} `xml:"code" json:"code,omitempty"`
	Title         string `xml:"title"`
	EffectiveTime struct {
		Text  string `xml:",chardata" json:"text,omitempty"`
		Value string `xml:"value,attr" json:"value,omitempty"`
	} `xml:"effectiveTime" json:"effectivetime,omitempty"`
	ConfidentialityCode struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		Code           string `xml:"code,attr" json:"code,omitempty"`
		CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
		CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
	} `xml:"confidentialityCode" json:"confidentialitycode,omitempty"`
	LanguageCode struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Code string `xml:"code,attr" json:"code,omitempty"`
	} `xml:"languageCode" json:"languagecode,omitempty"`
	RecordTarget struct {
		Text        string `xml:",chardata" json:"text,omitempty"`
		PatientRole struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			ID   []struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				Root      string `xml:"root,attr" json:"root,omitempty"`
				Extension string `xml:"extension,attr" json:"extension,omitempty"`
			} `xml:"id" json:"id,omitempty"`
			Addr struct {
				Text              string `xml:",chardata" json:"text,omitempty"`
				Use               string `xml:"use,attr" json:"use,omitempty"`
				State             string `xml:"state"`
				City              string `xml:"city"`
				PostalCode        string `xml:"postalCode"`
				StreetAddressLine string `xml:"streetAddressLine"`
				UseablePeriod     struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					Type string `xml:"type,attr" json:"type,omitempty"`
					Low  struct {
						Text  string `xml:",chardata" json:"text,omitempty"`
						Value string `xml:"value,attr" json:"value,omitempty"`
					} `xml:"low" json:"low,omitempty"`
					High struct {
						Text       string `xml:",chardata" json:"text,omitempty"`
						NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
					} `xml:"high" json:"high,omitempty"`
				} `xml:"useablePeriod" json:"useableperiod,omitempty"`
			} `xml:"addr" json:"addr,omitempty"`
			Telecom struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Value string `xml:"value,attr" json:"value,omitempty"`
				Use   string `xml:"use,attr" json:"use,omitempty"`
			} `xml:"telecom" json:"telecom,omitempty"`
			Patient struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Name struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					Use    string `xml:"use,attr" json:"use,omitempty"`
					Given  string `xml:"given"`
					Family string `xml:"family"`
				} `xml:"name" json:"name,omitempty"`
				AdministrativeGenderCode struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					Code        string `xml:"code,attr" json:"code,omitempty"`
					CodeSystem  string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
					DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
				} `xml:"administrativeGenderCode" json:"administrativegendercode,omitempty"`
				BirthTime struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Value string `xml:"value,attr" json:"value,omitempty"`
				} `xml:"birthTime" json:"birthtime,omitempty"`
				MaritalStatusCode struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
				} `xml:"maritalStatusCode" json:"maritalstatuscode,omitempty"`
				RaceCode struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
				} `xml:"raceCode" json:"racecode,omitempty"`
				EthnicGroupCode struct {
					Text        string `xml:",chardata" json:"text,omitempty"`
					NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
					DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
				} `xml:"ethnicGroupCode" json:"ethnicgroupcode,omitempty"`
				LanguageCommunication struct {
					Text         string `xml:",chardata" json:"text,omitempty"`
					LanguageCode struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Code string `xml:"code,attr" json:"code,omitempty"`
					} `xml:"languageCode" json:"languagecode,omitempty"`
					PreferenceInd struct {
						Text  string `xml:",chardata" json:"text,omitempty"`
						Value string `xml:"value,attr" json:"value,omitempty"`
					} `xml:"preferenceInd" json:"preferenceind,omitempty"`
				} `xml:"languageCommunication" json:"languagecommunication,omitempty"`
			} `xml:"patient" json:"patient,omitempty"`
		} `xml:"patientRole" json:"patientrole,omitempty"`
	} `xml:"recordTarget" json:"recordtarget,omitempty"`
	Author struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Time struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			Value string `xml:"value,attr" json:"value,omitempty"`
		} `xml:"time" json:"time,omitempty"`
		AssignedAuthor struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			ID   struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Root string `xml:"root,attr" json:"root,omitempty"`
			} `xml:"id" json:"id,omitempty"`
			Addr struct {
				Text              string `xml:",chardata" json:"text,omitempty"`
				Use               string `xml:"use,attr" json:"use,omitempty"`
				State             string `xml:"state"`
				City              string `xml:"city"`
				PostalCode        string `xml:"postalCode"`
				StreetAddressLine string `xml:"streetAddressLine"`
			} `xml:"addr" json:"addr,omitempty"`
			Telecom struct {
				Text  string `xml:",chardata" json:"text,omitempty"`
				Value string `xml:"value,attr" json:"value,omitempty"`
				Use   string `xml:"use,attr" json:"use,omitempty"`
			} `xml:"telecom" json:"telecom,omitempty"`
			AssignedPerson struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Name struct {
					Text   string `xml:",chardata" json:"text,omitempty"`
					Use    string `xml:"use,attr" json:"use,omitempty"`
					Given  string `xml:"given"`
					Family string `xml:"family"`
					Suffix string `xml:"suffix"`
				} `xml:"name" json:"name,omitempty"`
			} `xml:"assignedPerson" json:"assignedperson,omitempty"`
			RepresentedOrganization struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				ID   struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					Root      string `xml:"root,attr" json:"root,omitempty"`
					Extension string `xml:"extension,attr" json:"extension,omitempty"`
				} `xml:"id" json:"id,omitempty"`
				Name    string `xml:"name"`
				Telecom struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Value string `xml:"value,attr" json:"value,omitempty"`
					Use   string `xml:"use,attr" json:"use,omitempty"`
				} `xml:"telecom" json:"telecom,omitempty"`
				Addr struct {
					Text              string `xml:",chardata" json:"text,omitempty"`
					Use               string `xml:"use,attr" json:"use,omitempty"`
					State             string `xml:"state"`
					City              string `xml:"city"`
					PostalCode        string `xml:"postalCode"`
					StreetAddressLine string `xml:"streetAddressLine"`
				} `xml:"addr" json:"addr,omitempty"`
			} `xml:"representedOrganization" json:"representedorganization,omitempty"`
		} `xml:"assignedAuthor" json:"assignedauthor,omitempty"`
	} `xml:"author" json:"author,omitempty"`
	Custodian struct {
		Text              string `xml:",chardata" json:"text,omitempty"`
		AssignedCustodian struct {
			Text                             string `xml:",chardata" json:"text,omitempty"`
			RepresentedCustodianOrganization struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				ID   struct {
					Text      string `xml:",chardata" json:"text,omitempty"`
					Root      string `xml:"root,attr" json:"root,omitempty"`
					Extension string `xml:"extension,attr" json:"extension,omitempty"`
				} `xml:"id" json:"id,omitempty"`
				Name    string `xml:"name"`
				Telecom struct {
					Text  string `xml:",chardata" json:"text,omitempty"`
					Value string `xml:"value,attr" json:"value,omitempty"`
					Use   string `xml:"use,attr" json:"use,omitempty"`
				} `xml:"telecom" json:"telecom,omitempty"`
				Addr struct {
					Text              string `xml:",chardata" json:"text,omitempty"`
					Use               string `xml:"use,attr" json:"use,omitempty"`
					State             string `xml:"state"`
					City              string `xml:"city"`
					PostalCode        string `xml:"postalCode"`
					StreetAddressLine string `xml:"streetAddressLine"`
				} `xml:"addr" json:"addr,omitempty"`
			} `xml:"representedCustodianOrganization" json:"representedcustodianorganization,omitempty"`
		} `xml:"assignedCustodian" json:"assignedcustodian,omitempty"`
	} `xml:"custodian" json:"custodian,omitempty"`
	DocumentationOf struct {
		Text         string `xml:",chardata" json:"text,omitempty"`
		ServiceEvent struct {
			Text          string `xml:",chardata" json:"text,omitempty"`
			ClassCode     string `xml:"classCode,attr" json:"classcode,omitempty"`
			EffectiveTime struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				Low  struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
				} `xml:"low" json:"low,omitempty"`
				High struct {
					Text       string `xml:",chardata" json:"text,omitempty"`
					NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
				} `xml:"high" json:"high,omitempty"`
			} `xml:"effectiveTime" json:"effectivetime,omitempty"`
			Performer struct {
				Text         string `xml:",chardata" json:"text,omitempty"`
				TypeCode     string `xml:"typeCode,attr" json:"typecode,omitempty"`
				FunctionCode struct {
					Text           string `xml:",chardata" json:"text,omitempty"`
					Code           string `xml:"code,attr" json:"code,omitempty"`
					CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
					CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
					OriginalText   string `xml:"originalText"`
				} `xml:"functionCode" json:"functioncode,omitempty"`
				AssignedEntity struct {
					Text string `xml:",chardata" json:"text,omitempty"`
					ID   struct {
						Text      string `xml:",chardata" json:"text,omitempty"`
						Root      string `xml:"root,attr" json:"root,omitempty"`
						Extension string `xml:"extension,attr" json:"extension,omitempty"`
					} `xml:"id" json:"id,omitempty"`
					Code struct {
						Text           string `xml:",chardata" json:"text,omitempty"`
						Code           string `xml:"code,attr" json:"code,omitempty"`
						CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
						CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
						DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
					} `xml:"code" json:"code,omitempty"`
					Addr struct {
						Text              string `xml:",chardata" json:"text,omitempty"`
						Use               string `xml:"use,attr" json:"use,omitempty"`
						State             string `xml:"state"`
						City              string `xml:"city"`
						PostalCode        string `xml:"postalCode"`
						StreetAddressLine string `xml:"streetAddressLine"`
					} `xml:"addr" json:"addr,omitempty"`
					Telecom struct {
						Text  string `xml:",chardata" json:"text,omitempty"`
						Value string `xml:"value,attr" json:"value,omitempty"`
						Use   string `xml:"use,attr" json:"use,omitempty"`
					} `xml:"telecom" json:"telecom,omitempty"`
					AssignedPerson struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						Name struct {
							Text   string `xml:",chardata" json:"text,omitempty"`
							Use    string `xml:"use,attr" json:"use,omitempty"`
							Given  string `xml:"given"`
							Family string `xml:"family"`
							Suffix string `xml:"suffix"`
						} `xml:"name" json:"name,omitempty"`
					} `xml:"assignedPerson" json:"assignedperson,omitempty"`
					RepresentedOrganization struct {
						Text string `xml:",chardata" json:"text,omitempty"`
						ID   struct {
							Text      string `xml:",chardata" json:"text,omitempty"`
							Root      string `xml:"root,attr" json:"root,omitempty"`
							Extension string `xml:"extension,attr" json:"extension,omitempty"`
						} `xml:"id" json:"id,omitempty"`
						Name    string `xml:"name"`
						Telecom struct {
							Text  string `xml:",chardata" json:"text,omitempty"`
							Value string `xml:"value,attr" json:"value,omitempty"`
							Use   string `xml:"use,attr" json:"use,omitempty"`
						} `xml:"telecom" json:"telecom,omitempty"`
						Addr struct {
							Text              string `xml:",chardata" json:"text,omitempty"`
							Use               string `xml:"use,attr" json:"use,omitempty"`
							State             string `xml:"state"`
							City              string `xml:"city"`
							PostalCode        string `xml:"postalCode"`
							StreetAddressLine string `xml:"streetAddressLine"`
						} `xml:"addr" json:"addr,omitempty"`
					} `xml:"representedOrganization" json:"representedorganization,omitempty"`
				} `xml:"assignedEntity" json:"assignedentity,omitempty"`
			} `xml:"performer" json:"performer,omitempty"`
		} `xml:"serviceEvent" json:"serviceevent,omitempty"`
	} `xml:"documentationOf" json:"documentationof,omitempty"`
	Component struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		StructuredBody struct {
			Text      string `xml:",chardata" json:"text,omitempty"`
			Component []struct {
				Text    string `xml:",chardata" json:"text,omitempty"`
				Section struct {
					Chardata   string `xml:",chardata" json:"chardata,omitempty"`
					NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
					TemplateId []struct {
						Text      string `xml:",chardata" json:"text,omitempty"`
						Root      string `xml:"root,attr" json:"root,omitempty"`
						Extension string `xml:"extension,attr" json:"extension,omitempty"`
					} `xml:"templateId" json:"templateid,omitempty"`
					Code struct {
						Text           string `xml:",chardata" json:"text,omitempty"`
						Code           string `xml:"code,attr" json:"code,omitempty"`
						CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
						CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
						DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
					} `xml:"code" json:"code,omitempty"`
					Title string `xml:"title"`
					Text  struct {
						Text  string `xml:",chardata" json:"text,omitempty"`
						Table []struct {
							Text        string `xml:",chardata" json:"text,omitempty"`
							Border      string `xml:"border,attr" json:"border,omitempty"`
							Width       string `xml:"width,attr" json:"width,omitempty"`
							Cellpadding string `xml:"cellpadding,attr" json:"cellpadding,omitempty"`
							Cellspacing string `xml:"cellspacing,attr" json:"cellspacing,omitempty"`
							Thead       struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Tr   struct {
									Text string `xml:",chardata" json:"text,omitempty"`
									Th   []struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Sub  string `xml:"sub"`
									} `xml:"th" json:"th,omitempty"`
								} `xml:"tr" json:"tr,omitempty"`
							} `xml:"thead" json:"thead,omitempty"`
							Tbody struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Tr   []struct {
									Text string `xml:",chardata" json:"text,omitempty"`
									Td   []struct {
										Text    string `xml:",chardata" json:"text,omitempty"`
										ID      string `xml:"ID,attr" json:"id,omitempty"`
										Align   string `xml:"align,attr" json:"align,omitempty"`
										Colspan string `xml:"colspan,attr" json:"colspan,omitempty"`
										Content struct {
											Text      string   `xml:",chardata" json:"text,omitempty"`
											StyleCode string   `xml:"styleCode,attr" json:"stylecode,omitempty"`
											Br        []string `xml:"br"`
										} `xml:"content" json:"content,omitempty"`
										Paragraph string `xml:"paragraph"`
										List      struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											Item []struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												ID   string `xml:"ID,attr" json:"id,omitempty"`
											} `xml:"item" json:"item,omitempty"`
										} `xml:"list" json:"list,omitempty"`
									} `xml:"td" json:"td,omitempty"`
								} `xml:"tr" json:"tr,omitempty"`
							} `xml:"tbody" json:"tbody,omitempty"`
						} `xml:"table" json:"table,omitempty"`
						Content struct {
							Text string `xml:",chardata" json:"text,omitempty"`
							ID   string `xml:"ID,attr" json:"id,omitempty"`
						} `xml:"content" json:"content,omitempty"`
					} `xml:"text" json:"text,omitempty"`
					Entry []struct {
						Text     string `xml:",chardata" json:"text,omitempty"`
						TypeCode string `xml:"typeCode,attr" json:"typecode,omitempty"`
						Act      struct {
							Text       string `xml:",chardata" json:"text,omitempty"`
							ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
							MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
							TemplateId []struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"templateId" json:"templateid,omitempty"`
							ID struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"id" json:"id,omitempty"`
							Code struct {
								Text           string `xml:",chardata" json:"text,omitempty"`
								Code           string `xml:"code,attr" json:"code,omitempty"`
								CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
								CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
								DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
							} `xml:"code" json:"code,omitempty"`
							StatusCode struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Code string `xml:"code,attr" json:"code,omitempty"`
							} `xml:"statusCode" json:"statuscode,omitempty"`
							EffectiveTime struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Low  struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
									Value      string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"low" json:"low,omitempty"`
								High struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
									Value      string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"high" json:"high,omitempty"`
							} `xml:"effectiveTime" json:"effectivetime,omitempty"`
							EntryRelationship []struct {
								Text         string `xml:",chardata" json:"text,omitempty"`
								TypeCode     string `xml:"typeCode,attr" json:"typecode,omitempty"`
								InversionInd string `xml:"inversionInd,attr" json:"inversionind,omitempty"`
								Observation  struct {
									Chardata   string `xml:",chardata" json:"chardata,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
									TemplateId []struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									ID struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									Code struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										Code           string `xml:"code,attr" json:"code,omitempty"`
										CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
										DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
										Translation    struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"translation" json:"translation,omitempty"`
									} `xml:"code" json:"code,omitempty"`
									Text struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Reference struct {
											Text  string `xml:",chardata" json:"text,omitempty"`
											Value string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"reference" json:"reference,omitempty"`
									} `xml:"text" json:"text,omitempty"`
									StatusCode struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Code string `xml:"code,attr" json:"code,omitempty"`
									} `xml:"statusCode" json:"statuscode,omitempty"`
									EffectiveTime struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Low  struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											Value      string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"low" json:"low,omitempty"`
										High struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											Value      string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"high" json:"high,omitempty"`
									} `xml:"effectiveTime" json:"effectivetime,omitempty"`
									Value struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										Type           string `xml:"type,attr" json:"type,omitempty"`
										Code           string `xml:"code,attr" json:"code,omitempty"`
										CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
										DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
										NullFlavor     string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										Translation    struct {
											Text           string `xml:",chardata" json:"text,omitempty"`
											Code           string `xml:"code,attr" json:"code,omitempty"`
											CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
											CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
											DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
										} `xml:"translation" json:"translation,omitempty"`
									} `xml:"value" json:"value,omitempty"`
									Author struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										TemplateId struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											Root string `xml:"root,attr" json:"root,omitempty"`
										} `xml:"templateId" json:"templateid,omitempty"`
										Time struct {
											Text  string `xml:",chardata" json:"text,omitempty"`
											Value string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"time" json:"time,omitempty"`
										AssignedAuthor struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											ID   struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											AssignedPerson struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Name struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													Use        string `xml:"use,attr" json:"use,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
													Given      string `xml:"given"`
													Family     string `xml:"family"`
													Suffix     string `xml:"suffix"`
												} `xml:"name" json:"name,omitempty"`
											} `xml:"assignedPerson" json:"assignedperson,omitempty"`
											RepresentedOrganization struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												ID   struct {
													Text      string `xml:",chardata" json:"text,omitempty"`
													Root      string `xml:"root,attr" json:"root,omitempty"`
													Extension string `xml:"extension,attr" json:"extension,omitempty"`
												} `xml:"id" json:"id,omitempty"`
												Name    string `xml:"name"`
												Telecom struct {
													Text  string `xml:",chardata" json:"text,omitempty"`
													Value string `xml:"value,attr" json:"value,omitempty"`
													Use   string `xml:"use,attr" json:"use,omitempty"`
												} `xml:"telecom" json:"telecom,omitempty"`
												Addr struct {
													Text              string `xml:",chardata" json:"text,omitempty"`
													Use               string `xml:"use,attr" json:"use,omitempty"`
													State             string `xml:"state"`
													City              string `xml:"city"`
													PostalCode        string `xml:"postalCode"`
													StreetAddressLine string `xml:"streetAddressLine"`
												} `xml:"addr" json:"addr,omitempty"`
											} `xml:"representedOrganization" json:"representedorganization,omitempty"`
										} `xml:"assignedAuthor" json:"assignedauthor,omitempty"`
									} `xml:"author" json:"author,omitempty"`
									EntryRelationship struct {
										Text         string `xml:",chardata" json:"text,omitempty"`
										TypeCode     string `xml:"typeCode,attr" json:"typecode,omitempty"`
										InversionInd string `xml:"inversionInd,attr" json:"inversionind,omitempty"`
										Observation  struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
											MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
											TemplateId struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Root string `xml:"root,attr" json:"root,omitempty"`
											} `xml:"templateId" json:"templateid,omitempty"`
											Code struct {
												Text           string `xml:",chardata" json:"text,omitempty"`
												Code           string `xml:"code,attr" json:"code,omitempty"`
												CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
												CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
												DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
												Translation    struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												} `xml:"translation" json:"translation,omitempty"`
											} `xml:"code" json:"code,omitempty"`
											StatusCode struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Code string `xml:"code,attr" json:"code,omitempty"`
											} `xml:"statusCode" json:"statuscode,omitempty"`
											Value struct {
												Text           string `xml:",chardata" json:"text,omitempty"`
												Type           string `xml:"type,attr" json:"type,omitempty"`
												Code           string `xml:"code,attr" json:"code,omitempty"`
												CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
												CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
												DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
												Translation    struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												} `xml:"translation" json:"translation,omitempty"`
											} `xml:"value" json:"value,omitempty"`
										} `xml:"observation" json:"observation,omitempty"`
									} `xml:"entryRelationship" json:"entryrelationship,omitempty"`
									Participant struct {
										Text            string `xml:",chardata" json:"text,omitempty"`
										TypeCode        string `xml:"typeCode,attr" json:"typecode,omitempty"`
										ParticipantRole struct {
											Text          string `xml:",chardata" json:"text,omitempty"`
											ClassCode     string `xml:"classCode,attr" json:"classcode,omitempty"`
											PlayingEntity struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												ClassCode string `xml:"classCode,attr" json:"classcode,omitempty"`
												Code      struct {
													Text        string `xml:",chardata" json:"text,omitempty"`
													NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
													DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
												} `xml:"code" json:"code,omitempty"`
											} `xml:"playingEntity" json:"playingentity,omitempty"`
										} `xml:"participantRole" json:"participantrole,omitempty"`
									} `xml:"participant" json:"participant,omitempty"`
								} `xml:"observation" json:"observation,omitempty"`
								Act struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
									TemplateId []struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									ID struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									Code struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										Type           string `xml:"type,attr" json:"type,omitempty"`
										Code           string `xml:"code,attr" json:"code,omitempty"`
										CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
										DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
										Translation    struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"translation" json:"translation,omitempty"`
									} `xml:"code" json:"code,omitempty"`
									StatusCode struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Code string `xml:"code,attr" json:"code,omitempty"`
									} `xml:"statusCode" json:"statuscode,omitempty"`
									Performer struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										TypeCode   string `xml:"typeCode,attr" json:"typecode,omitempty"`
										TemplateId struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											Root string `xml:"root,attr" json:"root,omitempty"`
										} `xml:"templateId" json:"templateid,omitempty"`
										AssignedEntity struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											ID   struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											Code struct {
												Text           string `xml:",chardata" json:"text,omitempty"`
												Code           string `xml:"code,attr" json:"code,omitempty"`
												CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
												CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
												DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
											} `xml:"code" json:"code,omitempty"`
											RepresentedOrganization struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												ID   struct {
													Text      string `xml:",chardata" json:"text,omitempty"`
													Root      string `xml:"root,attr" json:"root,omitempty"`
													Extension string `xml:"extension,attr" json:"extension,omitempty"`
												} `xml:"id" json:"id,omitempty"`
												Name    string `xml:"name"`
												Telecom struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												} `xml:"telecom" json:"telecom,omitempty"`
											} `xml:"representedOrganization" json:"representedorganization,omitempty"`
											Addr struct {
												Text              string `xml:",chardata" json:"text,omitempty"`
												NullFlavor        string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												StreetAddressLine string `xml:"streetAddressLine"`
												City              string `xml:"city"`
												State             string `xml:"state"`
												PostalCode        string `xml:"postalCode"`
											} `xml:"addr" json:"addr,omitempty"`
											Telecom struct {
												Text  string `xml:",chardata" json:"text,omitempty"`
												Value string `xml:"value,attr" json:"value,omitempty"`
												Use   string `xml:"use,attr" json:"use,omitempty"`
											} `xml:"telecom" json:"telecom,omitempty"`
											AssignedPerson struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Name struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													Use        string `xml:"use,attr" json:"use,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
													Given      string `xml:"given"`
													Family     string `xml:"family"`
												} `xml:"name" json:"name,omitempty"`
											} `xml:"assignedPerson" json:"assignedperson,omitempty"`
										} `xml:"assignedEntity" json:"assignedentity,omitempty"`
									} `xml:"performer" json:"performer,omitempty"`
									Participant struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										TypeCode   string `xml:"typeCode,attr" json:"typecode,omitempty"`
										TemplateId struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											Root string `xml:"root,attr" json:"root,omitempty"`
										} `xml:"templateId" json:"templateid,omitempty"`
										ParticipantRole struct {
											Text      string `xml:",chardata" json:"text,omitempty"`
											ClassCode string `xml:"classCode,attr" json:"classcode,omitempty"`
											ID        struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											Code struct {
												Text           string `xml:",chardata" json:"text,omitempty"`
												Code           string `xml:"code,attr" json:"code,omitempty"`
												CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
												CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
												DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
											} `xml:"code" json:"code,omitempty"`
										} `xml:"participantRole" json:"participantrole,omitempty"`
									} `xml:"participant" json:"participant,omitempty"`
									EntryRelationship struct {
										Text     string `xml:",chardata" json:"text,omitempty"`
										TypeCode string `xml:"typeCode,attr" json:"typecode,omitempty"`
										Act      struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
											MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
											TemplateId struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Root string `xml:"root,attr" json:"root,omitempty"`
											} `xml:"templateId" json:"templateid,omitempty"`
											ID struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											Code struct {
												Text       string `xml:",chardata" json:"text,omitempty"`
												NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											} `xml:"code" json:"code,omitempty"`
											EntryRelationship struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												TypeCode  string `xml:"typeCode,attr" json:"typecode,omitempty"`
												Procedure struct {
													Text      string `xml:",chardata" json:"text,omitempty"`
													ClassCode string `xml:"classCode,attr" json:"classcode,omitempty"`
													MoodCode  string `xml:"moodCode,attr" json:"moodcode,omitempty"`
													Code      struct {
														Text       string `xml:",chardata" json:"text,omitempty"`
														NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
													} `xml:"code" json:"code,omitempty"`
												} `xml:"procedure" json:"procedure,omitempty"`
											} `xml:"entryRelationship" json:"entryrelationship,omitempty"`
										} `xml:"act" json:"act,omitempty"`
									} `xml:"entryRelationship" json:"entryrelationship,omitempty"`
								} `xml:"act" json:"act,omitempty"`
							} `xml:"entryRelationship" json:"entryrelationship,omitempty"`
						} `xml:"act" json:"act,omitempty"`
						SubstanceAdministration struct {
							Chardata    string `xml:",chardata" json:"chardata,omitempty"`
							ClassCode   string `xml:"classCode,attr" json:"classcode,omitempty"`
							MoodCode    string `xml:"moodCode,attr" json:"moodcode,omitempty"`
							NegationInd string `xml:"negationInd,attr" json:"negationind,omitempty"`
							NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
							TemplateId  []struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"templateId" json:"templateid,omitempty"`
							ID struct {
								Text       string `xml:",chardata" json:"text,omitempty"`
								Root       string `xml:"root,attr" json:"root,omitempty"`
								Extension  string `xml:"extension,attr" json:"extension,omitempty"`
								NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
							} `xml:"id" json:"id,omitempty"`
							Text struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Reference struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"reference" json:"reference,omitempty"`
							} `xml:"text" json:"text,omitempty"`
							StatusCode struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Code string `xml:"code,attr" json:"code,omitempty"`
							} `xml:"statusCode" json:"statuscode,omitempty"`
							EffectiveTime []struct {
								Text                 string `xml:",chardata" json:"text,omitempty"`
								Type                 string `xml:"type,attr" json:"type,omitempty"`
								Operator             string `xml:"operator,attr" json:"operator,omitempty"`
								InstitutionSpecified string `xml:"institutionSpecified,attr" json:"institutionspecified,omitempty"`
								NullFlavor           string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								Value                string `xml:"value,attr" json:"value,omitempty"`
								Low                  struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									Value      string `xml:"value,attr" json:"value,omitempty"`
									NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								} `xml:"low" json:"low,omitempty"`
								High struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									Value      string `xml:"value,attr" json:"value,omitempty"`
									NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								} `xml:"high" json:"high,omitempty"`
								Period struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
									Unit  string `xml:"unit,attr" json:"unit,omitempty"`
								} `xml:"period" json:"period,omitempty"`
							} `xml:"effectiveTime" json:"effectivetime,omitempty"`
							RouteCode struct {
								Text        string `xml:",chardata" json:"text,omitempty"`
								NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
							} `xml:"routeCode" json:"routecode,omitempty"`
							DoseQuantity struct {
								Text        string `xml:",chardata" json:"text,omitempty"`
								NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								Translation struct {
									Text        string `xml:",chardata" json:"text,omitempty"`
									NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
									DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
								} `xml:"translation" json:"translation,omitempty"`
							} `xml:"doseQuantity" json:"dosequantity,omitempty"`
							Consumable struct {
								Text                string `xml:",chardata" json:"text,omitempty"`
								ManufacturedProduct struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									TemplateId []struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									ManufacturedMaterial struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Code struct {
											Text           string `xml:",chardata" json:"text,omitempty"`
											Code           string `xml:"code,attr" json:"code,omitempty"`
											CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
											CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
											DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
											NullFlavor     string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											OriginalText   struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Reference struct {
													Text  string `xml:",chardata" json:"text,omitempty"`
													Value string `xml:"value,attr" json:"value,omitempty"`
												} `xml:"reference" json:"reference,omitempty"`
											} `xml:"originalText" json:"originaltext,omitempty"`
											Translation struct {
												Text           string `xml:",chardata" json:"text,omitempty"`
												Code           string `xml:"code,attr" json:"code,omitempty"`
												CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
												CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
												DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
											} `xml:"translation" json:"translation,omitempty"`
										} `xml:"code" json:"code,omitempty"`
										LotNumberText string `xml:"lotNumberText"`
									} `xml:"manufacturedMaterial" json:"manufacturedmaterial,omitempty"`
									ManufacturerOrganization struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Name struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"name" json:"name,omitempty"`
									} `xml:"manufacturerOrganization" json:"manufacturerorganization,omitempty"`
								} `xml:"manufacturedProduct" json:"manufacturedproduct,omitempty"`
							} `xml:"consumable" json:"consumable,omitempty"`
							Author struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Time struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"time" json:"time,omitempty"`
								AssignedAuthor struct {
									Text string `xml:",chardata" json:"text,omitempty"`
									ID   struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									AssignedPerson struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Name struct {
											Text   string `xml:",chardata" json:"text,omitempty"`
											Use    string `xml:"use,attr" json:"use,omitempty"`
											Given  string `xml:"given"`
											Family string `xml:"family"`
											Suffix string `xml:"suffix"`
										} `xml:"name" json:"name,omitempty"`
									} `xml:"assignedPerson" json:"assignedperson,omitempty"`
									RepresentedOrganization struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										ID   []struct {
											Text      string `xml:",chardata" json:"text,omitempty"`
											Root      string `xml:"root,attr" json:"root,omitempty"`
											Extension string `xml:"extension,attr" json:"extension,omitempty"`
										} `xml:"id" json:"id,omitempty"`
										Name struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"name" json:"name,omitempty"`
										Telecom struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"telecom" json:"telecom,omitempty"`
										Addr struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											Use        string `xml:"use,attr" json:"use,omitempty"`
										} `xml:"addr" json:"addr,omitempty"`
									} `xml:"representedOrganization" json:"representedorganization,omitempty"`
								} `xml:"assignedAuthor" json:"assignedauthor,omitempty"`
							} `xml:"author" json:"author,omitempty"`
							EntryRelationship []struct {
								Text        string `xml:",chardata" json:"text,omitempty"`
								TypeCode    string `xml:"typeCode,attr" json:"typecode,omitempty"`
								Observation struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
									TemplateId struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Root string `xml:"root,attr" json:"root,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									Code struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										Code           string `xml:"code,attr" json:"code,omitempty"`
										CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
										DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
										Translation    struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"translation" json:"translation,omitempty"`
									} `xml:"code" json:"code,omitempty"`
									Value struct {
										Text        string `xml:",chardata" json:"text,omitempty"`
										Type        string `xml:"type,attr" json:"type,omitempty"`
										Code        string `xml:"code,attr" json:"code,omitempty"`
										CodeSystem  string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
									} `xml:"value" json:"value,omitempty"`
								} `xml:"observation" json:"observation,omitempty"`
								Supply struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
									TemplateId []struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									ID struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									StatusCode struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Code string `xml:"code,attr" json:"code,omitempty"`
									} `xml:"statusCode" json:"statuscode,omitempty"`
									EffectiveTime struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Type string `xml:"type,attr" json:"type,omitempty"`
										High struct {
											Text  string `xml:",chardata" json:"text,omitempty"`
											Value string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"high" json:"high,omitempty"`
									} `xml:"effectiveTime" json:"effectivetime,omitempty"`
									RepeatNumber struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										Value      string `xml:"value,attr" json:"value,omitempty"`
										NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
									} `xml:"repeatNumber" json:"repeatnumber,omitempty"`
									Quantity struct {
										Text  string `xml:",chardata" json:"text,omitempty"`
										Value string `xml:"value,attr" json:"value,omitempty"`
									} `xml:"quantity" json:"quantity,omitempty"`
									Product struct {
										Text                string `xml:",chardata" json:"text,omitempty"`
										ManufacturedProduct struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
											TemplateId []struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"templateId" json:"templateid,omitempty"`
											ManufacturedMaterial struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Code struct {
													Text           string `xml:",chardata" json:"text,omitempty"`
													Code           string `xml:"code,attr" json:"code,omitempty"`
													CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
													CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
													DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
													OriginalText   struct {
														Text      string `xml:",chardata" json:"text,omitempty"`
														Reference struct {
															Text  string `xml:",chardata" json:"text,omitempty"`
															Value string `xml:"value,attr" json:"value,omitempty"`
														} `xml:"reference" json:"reference,omitempty"`
													} `xml:"originalText" json:"originaltext,omitempty"`
												} `xml:"code" json:"code,omitempty"`
											} `xml:"manufacturedMaterial" json:"manufacturedmaterial,omitempty"`
											ManufacturerOrganization struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Name string `xml:"name"`
											} `xml:"manufacturerOrganization" json:"manufacturerorganization,omitempty"`
										} `xml:"manufacturedProduct" json:"manufacturedproduct,omitempty"`
									} `xml:"product" json:"product,omitempty"`
									Author struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Time struct {
											Text  string `xml:",chardata" json:"text,omitempty"`
											Value string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"time" json:"time,omitempty"`
										AssignedAuthor struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											ID   []struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											AssignedPerson struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Name struct {
													Text   string `xml:",chardata" json:"text,omitempty"`
													Use    string `xml:"use,attr" json:"use,omitempty"`
													Given  string `xml:"given"`
													Family string `xml:"family"`
													Suffix string `xml:"suffix"`
												} `xml:"name" json:"name,omitempty"`
											} `xml:"assignedPerson" json:"assignedperson,omitempty"`
										} `xml:"assignedAuthor" json:"assignedauthor,omitempty"`
									} `xml:"author" json:"author,omitempty"`
									Performer struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										AssignedEntity struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											ID   struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											Addr struct {
												Text       string `xml:",chardata" json:"text,omitempty"`
												NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											} `xml:"addr" json:"addr,omitempty"`
											AssignedPerson struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Name struct {
													Text   string `xml:",chardata" json:"text,omitempty"`
													Use    string `xml:"use,attr" json:"use,omitempty"`
													Given  string `xml:"given"`
													Family string `xml:"family"`
													Suffix string `xml:"suffix"`
												} `xml:"name" json:"name,omitempty"`
											} `xml:"assignedPerson" json:"assignedperson,omitempty"`
											RepresentedOrganization struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												ID   []struct {
													Text      string `xml:",chardata" json:"text,omitempty"`
													Root      string `xml:"root,attr" json:"root,omitempty"`
													Extension string `xml:"extension,attr" json:"extension,omitempty"`
												} `xml:"id" json:"id,omitempty"`
												Name struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												} `xml:"name" json:"name,omitempty"`
												Telecom struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												} `xml:"telecom" json:"telecom,omitempty"`
												Addr struct {
													Text       string `xml:",chardata" json:"text,omitempty"`
													NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
													Use        string `xml:"use,attr" json:"use,omitempty"`
												} `xml:"addr" json:"addr,omitempty"`
											} `xml:"representedOrganization" json:"representedorganization,omitempty"`
										} `xml:"assignedEntity" json:"assignedentity,omitempty"`
									} `xml:"performer" json:"performer,omitempty"`
								} `xml:"supply" json:"supply,omitempty"`
							} `xml:"entryRelationship" json:"entryrelationship,omitempty"`
							ApproachSiteCode struct {
								Text        string `xml:",chardata" json:"text,omitempty"`
								NullFlavor  string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								DisplayName string `xml:"displayName,attr" json:"displayname,omitempty"`
							} `xml:"approachSiteCode" json:"approachsitecode,omitempty"`
							Performer struct {
								Text           string `xml:",chardata" json:"text,omitempty"`
								AssignedEntity struct {
									Text string `xml:",chardata" json:"text,omitempty"`
									ID   struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									AssignedPerson struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Name struct {
											Text   string `xml:",chardata" json:"text,omitempty"`
											Use    string `xml:"use,attr" json:"use,omitempty"`
											Given  string `xml:"given"`
											Family string `xml:"family"`
										} `xml:"name" json:"name,omitempty"`
									} `xml:"assignedPerson" json:"assignedperson,omitempty"`
									RepresentedOrganization struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										ID   []struct {
											Text      string `xml:",chardata" json:"text,omitempty"`
											Root      string `xml:"root,attr" json:"root,omitempty"`
											Extension string `xml:"extension,attr" json:"extension,omitempty"`
										} `xml:"id" json:"id,omitempty"`
										Name struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"name" json:"name,omitempty"`
										Telecom struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"telecom" json:"telecom,omitempty"`
										Addr struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											Use        string `xml:"use,attr" json:"use,omitempty"`
										} `xml:"addr" json:"addr,omitempty"`
									} `xml:"representedOrganization" json:"representedorganization,omitempty"`
								} `xml:"assignedEntity" json:"assignedentity,omitempty"`
							} `xml:"performer" json:"performer,omitempty"`
						} `xml:"substanceAdministration" json:"substanceadministration,omitempty"`
						Organizer struct {
							Text       string `xml:",chardata" json:"text,omitempty"`
							ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
							MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
							TemplateId []struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"templateId" json:"templateid,omitempty"`
							ID struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"id" json:"id,omitempty"`
							Code struct {
								Text           string `xml:",chardata" json:"text,omitempty"`
								NullFlavor     string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
								Code           string `xml:"code,attr" json:"code,omitempty"`
								CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
								CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
								Translation    struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									Code       string `xml:"code,attr" json:"code,omitempty"`
									CodeSystem string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
								} `xml:"translation" json:"translation,omitempty"`
							} `xml:"code" json:"code,omitempty"`
							StatusCode struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Code string `xml:"code,attr" json:"code,omitempty"`
							} `xml:"statusCode" json:"statuscode,omitempty"`
							EffectiveTime struct {
								Text  string `xml:",chardata" json:"text,omitempty"`
								Value string `xml:"value,attr" json:"value,omitempty"`
								Low   struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"low" json:"low,omitempty"`
								High struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"high" json:"high,omitempty"`
							} `xml:"effectiveTime" json:"effectivetime,omitempty"`
							Component []struct {
								Text        string `xml:",chardata" json:"text,omitempty"`
								Observation struct {
									Chardata   string `xml:",chardata" json:"chardata,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
									TemplateId []struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									ID struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Root      string `xml:"root,attr" json:"root,omitempty"`
										Extension string `xml:"extension,attr" json:"extension,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									Code struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										NullFlavor     string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										Code           string `xml:"code,attr" json:"code,omitempty"`
										CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
										DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
										Translation    struct {
											Text       string `xml:",chardata" json:"text,omitempty"`
											NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										} `xml:"translation" json:"translation,omitempty"`
									} `xml:"code" json:"code,omitempty"`
									Text struct {
										Text      string `xml:",chardata" json:"text,omitempty"`
										Reference struct {
											Text  string `xml:",chardata" json:"text,omitempty"`
											Value string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"reference" json:"reference,omitempty"`
									} `xml:"text" json:"text,omitempty"`
									StatusCode struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Code string `xml:"code,attr" json:"code,omitempty"`
									} `xml:"statusCode" json:"statuscode,omitempty"`
									EffectiveTime struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										Value      string `xml:"value,attr" json:"value,omitempty"`
									} `xml:"effectiveTime" json:"effectivetime,omitempty"`
									Value struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										Type       string `xml:"type,attr" json:"type,omitempty"`
										NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										Value      string `xml:"value,attr" json:"value,omitempty"`
										Unit       string `xml:"unit,attr" json:"unit,omitempty"`
									} `xml:"value" json:"value,omitempty"`
									InterpretationCode struct {
										Text           string `xml:",chardata" json:"text,omitempty"`
										NullFlavor     string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
										CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
										Code           string `xml:"code,attr" json:"code,omitempty"`
									} `xml:"interpretationCode" json:"interpretationcode,omitempty"`
									ReferenceRange struct {
										Text             string `xml:",chardata" json:"text,omitempty"`
										ObservationRange struct {
											Chardata string `xml:",chardata" json:"chardata,omitempty"`
											Text     struct {
												Text       string `xml:",chardata" json:"text,omitempty"`
												NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
												Reference  struct {
													Text  string `xml:",chardata" json:"text,omitempty"`
													Value string `xml:"value,attr" json:"value,omitempty"`
												} `xml:"reference" json:"reference,omitempty"`
											} `xml:"text" json:"text,omitempty"`
											Value struct {
												Text       string `xml:",chardata" json:"text,omitempty"`
												Type       string `xml:"type,attr" json:"type,omitempty"`
												NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
											} `xml:"value" json:"value,omitempty"`
										} `xml:"observationRange" json:"observationrange,omitempty"`
									} `xml:"referenceRange" json:"referencerange,omitempty"`
									EntryRelationship struct {
										Text     string `xml:",chardata" json:"text,omitempty"`
										TypeCode string `xml:"typeCode,attr" json:"typecode,omitempty"`
										Act      struct {
											Chardata   string `xml:",chardata" json:"chardata,omitempty"`
											ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
											MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
											TemplateId struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Root      string `xml:"root,attr" json:"root,omitempty"`
												Extension string `xml:"extension,attr" json:"extension,omitempty"`
											} `xml:"templateId" json:"templateid,omitempty"`
											ID struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Root string `xml:"root,attr" json:"root,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											Code struct {
												Text           string `xml:",chardata" json:"text,omitempty"`
												Code           string `xml:"code,attr" json:"code,omitempty"`
												CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
												CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
												DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
												Translation    struct {
													Text           string `xml:",chardata" json:"text,omitempty"`
													Code           string `xml:"code,attr" json:"code,omitempty"`
													CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
													CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
													DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
												} `xml:"translation" json:"translation,omitempty"`
											} `xml:"code" json:"code,omitempty"`
											Text struct {
												Text      string `xml:",chardata" json:"text,omitempty"`
												Reference struct {
													Text  string `xml:",chardata" json:"text,omitempty"`
													Value string `xml:"value,attr" json:"value,omitempty"`
												} `xml:"reference" json:"reference,omitempty"`
											} `xml:"text" json:"text,omitempty"`
											StatusCode struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Code string `xml:"code,attr" json:"code,omitempty"`
											} `xml:"statusCode" json:"statuscode,omitempty"`
											EffectiveTime struct {
												Text  string `xml:",chardata" json:"text,omitempty"`
												Value string `xml:"value,attr" json:"value,omitempty"`
											} `xml:"effectiveTime" json:"effectivetime,omitempty"`
											Author struct {
												Text       string `xml:",chardata" json:"text,omitempty"`
												TemplateId struct {
													Text string `xml:",chardata" json:"text,omitempty"`
													Root string `xml:"root,attr" json:"root,omitempty"`
												} `xml:"templateId" json:"templateid,omitempty"`
												Time struct {
													Text  string `xml:",chardata" json:"text,omitempty"`
													Value string `xml:"value,attr" json:"value,omitempty"`
												} `xml:"time" json:"time,omitempty"`
												AssignedAuthor struct {
													Text string `xml:",chardata" json:"text,omitempty"`
													ID   struct {
														Text string `xml:",chardata" json:"text,omitempty"`
														Root string `xml:"root,attr" json:"root,omitempty"`
													} `xml:"id" json:"id,omitempty"`
													RepresentedOrganization struct {
														Text string `xml:",chardata" json:"text,omitempty"`
														ID   struct {
															Text      string `xml:",chardata" json:"text,omitempty"`
															Root      string `xml:"root,attr" json:"root,omitempty"`
															Extension string `xml:"extension,attr" json:"extension,omitempty"`
														} `xml:"id" json:"id,omitempty"`
														Name    string `xml:"name"`
														Telecom struct {
															Text  string `xml:",chardata" json:"text,omitempty"`
															Value string `xml:"value,attr" json:"value,omitempty"`
															Use   string `xml:"use,attr" json:"use,omitempty"`
														} `xml:"telecom" json:"telecom,omitempty"`
														Addr struct {
															Text              string `xml:",chardata" json:"text,omitempty"`
															Use               string `xml:"use,attr" json:"use,omitempty"`
															State             string `xml:"state"`
															City              string `xml:"city"`
															PostalCode        string `xml:"postalCode"`
															StreetAddressLine string `xml:"streetAddressLine"`
														} `xml:"addr" json:"addr,omitempty"`
													} `xml:"representedOrganization" json:"representedorganization,omitempty"`
												} `xml:"assignedAuthor" json:"assignedauthor,omitempty"`
											} `xml:"author" json:"author,omitempty"`
										} `xml:"act" json:"act,omitempty"`
									} `xml:"entryRelationship" json:"entryrelationship,omitempty"`
									Author struct {
										Text       string `xml:",chardata" json:"text,omitempty"`
										TemplateId struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											Root string `xml:"root,attr" json:"root,omitempty"`
										} `xml:"templateId" json:"templateid,omitempty"`
										Time struct {
											Text  string `xml:",chardata" json:"text,omitempty"`
											Value string `xml:"value,attr" json:"value,omitempty"`
										} `xml:"time" json:"time,omitempty"`
										AssignedAuthor struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											ID   struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Root string `xml:"root,attr" json:"root,omitempty"`
											} `xml:"id" json:"id,omitempty"`
											AssignedPerson struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												Name struct {
													Text   string `xml:",chardata" json:"text,omitempty"`
													Use    string `xml:"use,attr" json:"use,omitempty"`
													Given  string `xml:"given"`
													Family string `xml:"family"`
													Suffix string `xml:"suffix"`
													Prefix struct {
														Text      string `xml:",chardata" json:"text,omitempty"`
														Qualifier string `xml:"qualifier,attr" json:"qualifier,omitempty"`
													} `xml:"prefix" json:"prefix,omitempty"`
												} `xml:"name" json:"name,omitempty"`
											} `xml:"assignedPerson" json:"assignedperson,omitempty"`
											RepresentedOrganization struct {
												Text string `xml:",chardata" json:"text,omitempty"`
												ID   struct {
													Text      string `xml:",chardata" json:"text,omitempty"`
													Root      string `xml:"root,attr" json:"root,omitempty"`
													Extension string `xml:"extension,attr" json:"extension,omitempty"`
												} `xml:"id" json:"id,omitempty"`
												Name    string `xml:"name"`
												Telecom struct {
													Text  string `xml:",chardata" json:"text,omitempty"`
													Value string `xml:"value,attr" json:"value,omitempty"`
													Use   string `xml:"use,attr" json:"use,omitempty"`
												} `xml:"telecom" json:"telecom,omitempty"`
												Addr struct {
													Text              string `xml:",chardata" json:"text,omitempty"`
													Use               string `xml:"use,attr" json:"use,omitempty"`
													State             string `xml:"state"`
													City              string `xml:"city"`
													PostalCode        string `xml:"postalCode"`
													StreetAddressLine string `xml:"streetAddressLine"`
												} `xml:"addr" json:"addr,omitempty"`
											} `xml:"representedOrganization" json:"representedorganization,omitempty"`
										} `xml:"assignedAuthor" json:"assignedauthor,omitempty"`
									} `xml:"author" json:"author,omitempty"`
									MethodCode struct {
										Text         string `xml:",chardata" json:"text,omitempty"`
										NullFlavor   string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										DisplayName  string `xml:"displayName,attr" json:"displayname,omitempty"`
										OriginalText struct {
											Text      string `xml:",chardata" json:"text,omitempty"`
											Reference struct {
												Text  string `xml:",chardata" json:"text,omitempty"`
												Value string `xml:"value,attr" json:"value,omitempty"`
											} `xml:"reference" json:"reference,omitempty"`
										} `xml:"originalText" json:"originaltext,omitempty"`
									} `xml:"methodCode" json:"methodcode,omitempty"`
									TargetSiteCode struct {
										Text         string `xml:",chardata" json:"text,omitempty"`
										NullFlavor   string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
										DisplayName  string `xml:"displayName,attr" json:"displayname,omitempty"`
										OriginalText struct {
											Text      string `xml:",chardata" json:"text,omitempty"`
											Reference struct {
												Text  string `xml:",chardata" json:"text,omitempty"`
												Value string `xml:"value,attr" json:"value,omitempty"`
											} `xml:"reference" json:"reference,omitempty"`
										} `xml:"originalText" json:"originaltext,omitempty"`
									} `xml:"targetSiteCode" json:"targetsitecode,omitempty"`
								} `xml:"observation" json:"observation,omitempty"`
							} `xml:"component" json:"component,omitempty"`
						} `xml:"organizer" json:"organizer,omitempty"`
						Encounter struct {
							Text       string `xml:",chardata" json:"text,omitempty"`
							ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
							MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
							TemplateId []struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"templateId" json:"templateid,omitempty"`
							ID struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"id" json:"id,omitempty"`
							Code struct {
								Text           string `xml:",chardata" json:"text,omitempty"`
								Code           string `xml:"code,attr" json:"code,omitempty"`
								CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
								CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
								DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
								Translation    struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									NullFlavor string `xml:"nullFlavor,attr" json:"nullflavor,omitempty"`
								} `xml:"translation" json:"translation,omitempty"`
							} `xml:"code" json:"code,omitempty"`
							EffectiveTime struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Low  struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"low" json:"low,omitempty"`
							} `xml:"effectiveTime" json:"effectivetime,omitempty"`
						} `xml:"encounter" json:"encounter,omitempty"`
						Observation struct {
							Chardata   string `xml:",chardata" json:"chardata,omitempty"`
							ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
							MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
							TemplateId []struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"templateId" json:"templateid,omitempty"`
							ID struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"id" json:"id,omitempty"`
							Code struct {
								Text           string `xml:",chardata" json:"text,omitempty"`
								Code           string `xml:"code,attr" json:"code,omitempty"`
								CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
								CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
								DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
							} `xml:"code" json:"code,omitempty"`
							StatusCode struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Code string `xml:"code,attr" json:"code,omitempty"`
							} `xml:"statusCode" json:"statuscode,omitempty"`
							EffectiveTime struct {
								Text  string `xml:",chardata" json:"text,omitempty"`
								Value string `xml:"value,attr" json:"value,omitempty"`
							} `xml:"effectiveTime" json:"effectivetime,omitempty"`
							Value struct {
								Text           string `xml:",chardata" json:"text,omitempty"`
								Type           string `xml:"type,attr" json:"type,omitempty"`
								Code           string `xml:"code,attr" json:"code,omitempty"`
								CodeSystem     string `xml:"codeSystem,attr" json:"codesystem,omitempty"`
								CodeSystemName string `xml:"codeSystemName,attr" json:"codesystemname,omitempty"`
								DisplayName    string `xml:"displayName,attr" json:"displayname,omitempty"`
								OriginalText   struct {
									Text      string `xml:",chardata" json:"text,omitempty"`
									Reference struct {
										Text  string `xml:",chardata" json:"text,omitempty"`
										Value string `xml:"value,attr" json:"value,omitempty"`
									} `xml:"reference" json:"reference,omitempty"`
								} `xml:"originalText" json:"originaltext,omitempty"`
							} `xml:"value" json:"value,omitempty"`
							Text struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Reference struct {
									Text  string `xml:",chardata" json:"text,omitempty"`
									Value string `xml:"value,attr" json:"value,omitempty"`
								} `xml:"reference" json:"reference,omitempty"`
							} `xml:"text" json:"text,omitempty"`
						} `xml:"observation" json:"observation,omitempty"`
						Supply struct {
							Text       string `xml:",chardata" json:"text,omitempty"`
							ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
							MoodCode   string `xml:"moodCode,attr" json:"moodcode,omitempty"`
							TemplateId []struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"templateId" json:"templateid,omitempty"`
							ID struct {
								Text      string `xml:",chardata" json:"text,omitempty"`
								Root      string `xml:"root,attr" json:"root,omitempty"`
								Extension string `xml:"extension,attr" json:"extension,omitempty"`
							} `xml:"id" json:"id,omitempty"`
							StatusCode struct {
								Text string `xml:",chardata" json:"text,omitempty"`
								Code string `xml:"code,attr" json:"code,omitempty"`
							} `xml:"statusCode" json:"statuscode,omitempty"`
							Participant struct {
								Text            string `xml:",chardata" json:"text,omitempty"`
								TypeCode        string `xml:"typeCode,attr" json:"typecode,omitempty"`
								ParticipantRole struct {
									Text       string `xml:",chardata" json:"text,omitempty"`
									ClassCode  string `xml:"classCode,attr" json:"classcode,omitempty"`
									TemplateId struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										Root string `xml:"root,attr" json:"root,omitempty"`
									} `xml:"templateId" json:"templateid,omitempty"`
									ID struct {
										Text                   string `xml:",chardata" json:"text,omitempty"`
										Root                   string `xml:"root,attr" json:"root,omitempty"`
										Extension              string `xml:"extension,attr" json:"extension,omitempty"`
										AssigningAuthorityName string `xml:"assigningAuthorityName,attr" json:"assigningauthorityname,omitempty"`
									} `xml:"id" json:"id,omitempty"`
									PlayingDevice string `xml:"playingDevice"`
									ScopingEntity struct {
										Text string `xml:",chardata" json:"text,omitempty"`
										ID   struct {
											Text string `xml:",chardata" json:"text,omitempty"`
											Root string `xml:"root,attr" json:"root,omitempty"`
										} `xml:"id" json:"id,omitempty"`
									} `xml:"scopingEntity" json:"scopingentity,omitempty"`
								} `xml:"participantRole" json:"participantrole,omitempty"`
							} `xml:"participant" json:"participant,omitempty"`
						} `xml:"supply" json:"supply,omitempty"`
					} `xml:"entry" json:"entry,omitempty"`
				} `xml:"section" json:"section,omitempty"`
			} `xml:"component" json:"component,omitempty"`
		} `xml:"structuredBody" json:"structuredbody,omitempty"`
	} `xml:"component" json:"component,omitempty"`
}
