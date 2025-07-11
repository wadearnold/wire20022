package DrawdownRequest

func pathMapV1() map[string]any {
	return map[string]any{
		"CdtrPmtActvtnReq.GrpHdr.MsgId":                                                       "MessageId",
		"CdtrPmtActvtnReq.GrpHdr.CreDtTm":                                                     "CreatedDateTime",
		"CdtrPmtActvtnReq.GrpHdr.NbOfTxs":                                                     "NumberofTransaction",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm":                                                 "InitiatingParty.Name",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.StrtNm":                                     "InitiatingParty.Address.StreetName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.BldgNb":                                     "InitiatingParty.Address.BuildingNumber",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstCd":                                      "InitiatingParty.Address.PostalCode",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.TwnNm":                                      "InitiatingParty.Address.TownName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn":                                "InitiatingParty.Address.Subdivision",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Ctry":                                       "InitiatingParty.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].PmtInfId":                                                 "PaymentInfoId",
		"CdtrPmtActvtnReq.PmtInf[0].PmtMtd":                                                   "PaymentMethod",
		"CdtrPmtActvtnReq.PmtInf[0].ReqdExctnDt":                                              "RequestedExecutDate",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.Nm":                                                  "Debtor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.StrtNm":                                      "Debtor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.BldgNb":                                      "Debtor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.PstCd":                                       "Debtor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.TwnNm":                                       "Debtor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.CtrySubDvsn":                                 "Debtor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.Ctry":                                        "Debtor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":               "DebtorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":                     "DebtorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.InstrId":                                "CreditTransTransaction.PaymentInstructionId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.EndToEndId":                             "CreditTransTransaction.PaymentEndToEndId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.LclInstrm.Prtry":                     "CreditTransTransaction.PayRequestType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.CtgyPurp.Cd":                         "CreditTransTransaction.PayCategoryType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Value":                           "CreditTransTransaction.Amount.Amount",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Ccy":                             "CreditTransTransaction.Amount.Currency",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChrgBr":                                       "CreditTransTransaction.ChargeBearer",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "CreditTransTransaction.CreditorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "CreditTransTransaction.CreditorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Nm":                                      "CreditTransTransaction.Creditor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.StrtNm":                          "CreditTransTransaction.Creditor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.BldgNb":                          "CreditTransTransaction.Creditor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.PstCd":                           "CreditTransTransaction.Creditor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.TwnNm":                           "CreditTransTransaction.Creditor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.CtrySubDvsn":                     "CreditTransTransaction.Creditor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.Ctry":                            "CreditTransTransaction.Creditor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr.Id":                          "CreditTransTransaction.CreditorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Ustrd[0]":                              "CreditTransTransaction.RemittanceInformation",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd": "CreditTransTransaction.Document.CodeOrProprietary",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Nb":              "CreditTransTransaction.Document.Number",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt":          "CreditTransTransaction.Document.RelatedDate",
	}
}
func pathMapV2() map[string]any {
	return pathMapV5()
}
func pathMapV3() map[string]any {
	return pathMapV5()
}
func pathMapV4() map[string]any {
	return pathMapV5()
}
func pathMapV5() map[string]any {
	return map[string]any{
		"CdtrPmtActvtnReq.GrpHdr.MsgId":                                                       "MessageId",
		"CdtrPmtActvtnReq.GrpHdr.CreDtTm":                                                     "CreatedDateTime",
		"CdtrPmtActvtnReq.GrpHdr.NbOfTxs":                                                     "NumberofTransaction",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm":                                                 "InitiatingParty.Name",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.StrtNm":                                     "InitiatingParty.Address.StreetName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.BldgNb":                                     "InitiatingParty.Address.BuildingNumber",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstCd":                                      "InitiatingParty.Address.PostalCode",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.TwnNm":                                      "InitiatingParty.Address.TownName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn":                                "InitiatingParty.Address.Subdivision",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Ctry":                                       "InitiatingParty.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].PmtInfId":                                                 "PaymentInfoId",
		"CdtrPmtActvtnReq.PmtInf[0].PmtMtd":                                                   "PaymentMethod",
		"CdtrPmtActvtnReq.PmtInf[0].ReqdExctnDt":                                              "RequestedExecutDate",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.Nm":                                                  "Debtor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.StrtNm":                                      "Debtor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.BldgNb":                                      "Debtor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.PstCd":                                       "Debtor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.TwnNm":                                       "Debtor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.CtrySubDvsn":                                 "Debtor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.Ctry":                                        "Debtor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr.Id":                                      "AccountEnhancement.DebtorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":               "DebtorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":                     "DebtorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.InstrId":                                "CreditTransTransaction.PaymentInstructionId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.EndToEndId":                             "CreditTransTransaction.PaymentEndToEndId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.LclInstrm.Prtry":                     "CreditTransTransaction.PayRequestType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.CtgyPurp.Cd":                         "CreditTransTransaction.PayCategoryType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Value":                           "CreditTransTransaction.Amount.Amount",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Ccy":                             "CreditTransTransaction.Amount.Currency",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChrgBr":                                       "CreditTransTransaction.ChargeBearer",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "CreditTransTransaction.CreditorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "CreditTransTransaction.CreditorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Nm":                                      "CreditTransTransaction.Creditor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.StrtNm":                          "CreditTransTransaction.Creditor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.BldgNb":                          "CreditTransTransaction.Creditor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.PstCd":                           "CreditTransTransaction.Creditor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.TwnNm":                           "CreditTransTransaction.Creditor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.CtrySubDvsn":                     "CreditTransTransaction.Creditor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.Ctry":                            "CreditTransTransaction.Creditor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr.Id":                          "CreditTransTransaction.CreditorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Ustrd[0]":                              "CreditTransTransaction.RemittanceInformation",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd": "CreditTransTransaction.Document.CodeOrProprietary",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Nb":              "CreditTransTransaction.Document.Number",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt":          "CreditTransTransaction.Document.RelatedDate",
	}
}
func pathMapV6() map[string]any {
	return map[string]any{
		"CdtrPmtActvtnReq.GrpHdr.MsgId":                                                       "MessageId",
		"CdtrPmtActvtnReq.GrpHdr.CreDtTm":                                                     "CreatedDateTime",
		"CdtrPmtActvtnReq.GrpHdr.NbOfTxs":                                                     "NumberofTransaction",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm":                                                 "InitiatingParty.Name",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.StrtNm":                                     "InitiatingParty.Address.StreetName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.BldgNb":                                     "InitiatingParty.Address.BuildingNumber",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstCd":                                      "InitiatingParty.Address.PostalCode",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.TwnNm":                                      "InitiatingParty.Address.TownName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn":                                "InitiatingParty.Address.Subdivision",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Ctry":                                       "InitiatingParty.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].PmtInfId":                                                 "PaymentInfoId",
		"CdtrPmtActvtnReq.PmtInf[0].PmtMtd":                                                   "PaymentMethod",
		"CdtrPmtActvtnReq.PmtInf[0].ReqdExctnDt.Dt":                                           "RequestedExecutDate",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.Nm":                                                  "Debtor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.StrtNm":                                      "Debtor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.BldgNb":                                      "Debtor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.PstCd":                                       "Debtor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.TwnNm":                                       "Debtor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.CtrySubDvsn":                                 "Debtor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.Ctry":                                        "Debtor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr.Id":                                      "AccountEnhancement.DebtorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":               "DebtorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":                     "DebtorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.InstrId":                                "CreditTransTransaction.PaymentInstructionId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.EndToEndId":                             "CreditTransTransaction.PaymentEndToEndId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.LclInstrm.Prtry":                     "CreditTransTransaction.PayRequestType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.CtgyPurp.Cd":                         "CreditTransTransaction.PayCategoryType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Value":                           "CreditTransTransaction.Amount.Amount",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Ccy":                             "CreditTransTransaction.Amount.Currency",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChrgBr":                                       "CreditTransTransaction.ChargeBearer",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "CreditTransTransaction.CreditorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "CreditTransTransaction.CreditorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Nm":                                      "CreditTransTransaction.Creditor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.StrtNm":                          "CreditTransTransaction.Creditor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.BldgNb":                          "CreditTransTransaction.Creditor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.PstCd":                           "CreditTransTransaction.Creditor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.TwnNm":                           "CreditTransTransaction.Creditor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.CtrySubDvsn":                     "CreditTransTransaction.Creditor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.Ctry":                            "CreditTransTransaction.Creditor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr.Id":                          "CreditTransTransaction.CreditorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Ustrd[0]":                              "CreditTransTransaction.RemittanceInformation",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd": "CreditTransTransaction.Document.CodeOrProprietary",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Nb":              "CreditTransTransaction.Document.Number",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt":          "CreditTransTransaction.Document.RelatedDate",
	}
}
func pathMapV7() map[string]any {
	return map[string]any{
		"CdtrPmtActvtnReq.GrpHdr.MsgId":                                                       "MessageId",
		"CdtrPmtActvtnReq.GrpHdr.CreDtTm":                                                     "CreatedDateTime",
		"CdtrPmtActvtnReq.GrpHdr.NbOfTxs":                                                     "NumberofTransaction",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm":                                                 "InitiatingParty.Name",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.StrtNm":                                     "InitiatingParty.Address.StreetName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.BldgNb":                                     "InitiatingParty.Address.BuildingNumber",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Room":                                       "InitiatingParty.Address.RoomNumber",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstCd":                                      "InitiatingParty.Address.PostalCode",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.TwnNm":                                      "InitiatingParty.Address.TownName",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.CtrySubDvsn":                                "InitiatingParty.Address.Subdivision",
		"CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.Ctry":                                       "InitiatingParty.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].PmtInfId":                                                 "PaymentInfoId",
		"CdtrPmtActvtnReq.PmtInf[0].PmtMtd":                                                   "PaymentMethod",
		"CdtrPmtActvtnReq.PmtInf[0].ReqdExctnDt.Dt":                                           "RequestedExecutDate",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.Nm":                                                  "Debtor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.StrtNm":                                      "Debtor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.BldgNb":                                      "Debtor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.Room":                                        "Debtor.Address.RoomNumber",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.PstCd":                                       "Debtor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.TwnNm":                                       "Debtor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.CtrySubDvsn":                                 "Debtor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].Dbtr.PstlAdr.Ctry":                                        "Debtor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAcct.Id.Othr.Id":                                      "AccountEnhancement.DebtorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":               "DebtorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].DbtrAgt.FinInstnId.ClrSysMmbId.MmbId":                     "DebtorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.InstrId":                                "CreditTransTransaction.PaymentInstructionId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.EndToEndId":                             "CreditTransTransaction.PaymentEndToEndId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtId.UETR":                                   "CreditTransTransaction.PaymentUniqueId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.LclInstrm.Prtry":                     "CreditTransTransaction.PayRequestType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].PmtTpInf.CtgyPurp.Cd":                         "CreditTransTransaction.PayCategoryType",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Value":                           "CreditTransTransaction.Amount.Amount",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Amt.InstdAmt.Ccy":                             "CreditTransTransaction.Amount.Currency",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].ChrgBr":                                       "CreditTransTransaction.ChargeBearer",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":   "CreditTransTransaction.CreditorAgent.PaymentSysCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAgt.FinInstnId.ClrSysMmbId.MmbId":         "CreditTransTransaction.CreditorAgent.PaymentSysMemberId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.Nm":                                      "CreditTransTransaction.Creditor.Name",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.StrtNm":                          "CreditTransTransaction.Creditor.Address.StreetName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.BldgNb":                          "CreditTransTransaction.Creditor.Address.BuildingNumber",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.Room":                            "CreditTransTransaction.Creditor.Address.RoomNumber",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.PstCd":                           "CreditTransTransaction.Creditor.Address.PostalCode",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.TwnNm":                           "CreditTransTransaction.Creditor.Address.TownName",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.CtrySubDvsn":                     "CreditTransTransaction.Creditor.Address.Subdivision",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].Cdtr.PstlAdr.Ctry":                            "CreditTransTransaction.Creditor.Address.Country",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].CdtrAcct.Id.Othr.Id":                          "CreditTransTransaction.CreditorAccountOtherId",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Ustrd[0]":                              "CreditTransTransaction.RemittanceInformation",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Tp.CdOrPrtry.Cd": "CreditTransTransaction.Document.CodeOrProprietary",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].Nb":              "CreditTransTransaction.Document.Number",
		"CdtrPmtActvtnReq.PmtInf[0].CdtTrfTx[0].RmtInf.Strd[0].RfrdDocInf[0].RltdDt":          "CreditTransTransaction.Document.RelatedDate",
	}
}
func pathMapV8() map[string]any {
	return pathMapV7()
}
func pathMapV9() map[string]any {
	return pathMapV7()
}
func pathMapV10() map[string]any {
	return pathMapV7()
}
