/*
 * Merge HRIS API
 *
 * The unified API for building rich integrations with multiple HR Information System platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_hris_client

import (
	"encoding/json"
	"fmt"
)

// PayCurrencyEnum the model 'PayCurrencyEnum'
type PayCurrencyEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of PayCurrencyEnum
const (
    PAYCURRENCYENUM_MERGE_NONSTANDARD_VALUE PayCurrencyEnum = "MERGE_NONSTANDARD_VALUE"
    
	PAYCURRENCYENUM_XUA PayCurrencyEnum = "XUA"
	PAYCURRENCYENUM_AFN PayCurrencyEnum = "AFN"
	PAYCURRENCYENUM_AFA PayCurrencyEnum = "AFA"
	PAYCURRENCYENUM_ALL PayCurrencyEnum = "ALL"
	PAYCURRENCYENUM_ALK PayCurrencyEnum = "ALK"
	PAYCURRENCYENUM_DZD PayCurrencyEnum = "DZD"
	PAYCURRENCYENUM_ADP PayCurrencyEnum = "ADP"
	PAYCURRENCYENUM_AOA PayCurrencyEnum = "AOA"
	PAYCURRENCYENUM_AOK PayCurrencyEnum = "AOK"
	PAYCURRENCYENUM_AON PayCurrencyEnum = "AON"
	PAYCURRENCYENUM_AOR PayCurrencyEnum = "AOR"
	PAYCURRENCYENUM_ARA PayCurrencyEnum = "ARA"
	PAYCURRENCYENUM_ARS PayCurrencyEnum = "ARS"
	PAYCURRENCYENUM_ARM PayCurrencyEnum = "ARM"
	PAYCURRENCYENUM_ARP PayCurrencyEnum = "ARP"
	PAYCURRENCYENUM_ARL PayCurrencyEnum = "ARL"
	PAYCURRENCYENUM_AMD PayCurrencyEnum = "AMD"
	PAYCURRENCYENUM_AWG PayCurrencyEnum = "AWG"
	PAYCURRENCYENUM_AUD PayCurrencyEnum = "AUD"
	PAYCURRENCYENUM_ATS PayCurrencyEnum = "ATS"
	PAYCURRENCYENUM_AZN PayCurrencyEnum = "AZN"
	PAYCURRENCYENUM_AZM PayCurrencyEnum = "AZM"
	PAYCURRENCYENUM_BSD PayCurrencyEnum = "BSD"
	PAYCURRENCYENUM_BHD PayCurrencyEnum = "BHD"
	PAYCURRENCYENUM_BDT PayCurrencyEnum = "BDT"
	PAYCURRENCYENUM_BBD PayCurrencyEnum = "BBD"
	PAYCURRENCYENUM_BYN PayCurrencyEnum = "BYN"
	PAYCURRENCYENUM_BYB PayCurrencyEnum = "BYB"
	PAYCURRENCYENUM_BYR PayCurrencyEnum = "BYR"
	PAYCURRENCYENUM_BEF PayCurrencyEnum = "BEF"
	PAYCURRENCYENUM_BEC PayCurrencyEnum = "BEC"
	PAYCURRENCYENUM_BEL PayCurrencyEnum = "BEL"
	PAYCURRENCYENUM_BZD PayCurrencyEnum = "BZD"
	PAYCURRENCYENUM_BMD PayCurrencyEnum = "BMD"
	PAYCURRENCYENUM_BTN PayCurrencyEnum = "BTN"
	PAYCURRENCYENUM_BOB PayCurrencyEnum = "BOB"
	PAYCURRENCYENUM_BOL PayCurrencyEnum = "BOL"
	PAYCURRENCYENUM_BOV PayCurrencyEnum = "BOV"
	PAYCURRENCYENUM_BOP PayCurrencyEnum = "BOP"
	PAYCURRENCYENUM_BAM PayCurrencyEnum = "BAM"
	PAYCURRENCYENUM_BAD PayCurrencyEnum = "BAD"
	PAYCURRENCYENUM_BAN PayCurrencyEnum = "BAN"
	PAYCURRENCYENUM_BWP PayCurrencyEnum = "BWP"
	PAYCURRENCYENUM_BRC PayCurrencyEnum = "BRC"
	PAYCURRENCYENUM_BRZ PayCurrencyEnum = "BRZ"
	PAYCURRENCYENUM_BRE PayCurrencyEnum = "BRE"
	PAYCURRENCYENUM_BRR PayCurrencyEnum = "BRR"
	PAYCURRENCYENUM_BRN PayCurrencyEnum = "BRN"
	PAYCURRENCYENUM_BRB PayCurrencyEnum = "BRB"
	PAYCURRENCYENUM_BRL PayCurrencyEnum = "BRL"
	PAYCURRENCYENUM_GBP PayCurrencyEnum = "GBP"
	PAYCURRENCYENUM_BND PayCurrencyEnum = "BND"
	PAYCURRENCYENUM_BGL PayCurrencyEnum = "BGL"
	PAYCURRENCYENUM_BGN PayCurrencyEnum = "BGN"
	PAYCURRENCYENUM_BGO PayCurrencyEnum = "BGO"
	PAYCURRENCYENUM_BGM PayCurrencyEnum = "BGM"
	PAYCURRENCYENUM_BUK PayCurrencyEnum = "BUK"
	PAYCURRENCYENUM_BIF PayCurrencyEnum = "BIF"
	PAYCURRENCYENUM_XPF PayCurrencyEnum = "XPF"
	PAYCURRENCYENUM_KHR PayCurrencyEnum = "KHR"
	PAYCURRENCYENUM_CAD PayCurrencyEnum = "CAD"
	PAYCURRENCYENUM_CVE PayCurrencyEnum = "CVE"
	PAYCURRENCYENUM_KYD PayCurrencyEnum = "KYD"
	PAYCURRENCYENUM_XAF PayCurrencyEnum = "XAF"
	PAYCURRENCYENUM_CLE PayCurrencyEnum = "CLE"
	PAYCURRENCYENUM_CLP PayCurrencyEnum = "CLP"
	PAYCURRENCYENUM_CLF PayCurrencyEnum = "CLF"
	PAYCURRENCYENUM_CNX PayCurrencyEnum = "CNX"
	PAYCURRENCYENUM_CNY PayCurrencyEnum = "CNY"
	PAYCURRENCYENUM_CNH PayCurrencyEnum = "CNH"
	PAYCURRENCYENUM_COP PayCurrencyEnum = "COP"
	PAYCURRENCYENUM_COU PayCurrencyEnum = "COU"
	PAYCURRENCYENUM_KMF PayCurrencyEnum = "KMF"
	PAYCURRENCYENUM_CDF PayCurrencyEnum = "CDF"
	PAYCURRENCYENUM_CRC PayCurrencyEnum = "CRC"
	PAYCURRENCYENUM_HRD PayCurrencyEnum = "HRD"
	PAYCURRENCYENUM_HRK PayCurrencyEnum = "HRK"
	PAYCURRENCYENUM_CUC PayCurrencyEnum = "CUC"
	PAYCURRENCYENUM_CUP PayCurrencyEnum = "CUP"
	PAYCURRENCYENUM_CYP PayCurrencyEnum = "CYP"
	PAYCURRENCYENUM_CZK PayCurrencyEnum = "CZK"
	PAYCURRENCYENUM_CSK PayCurrencyEnum = "CSK"
	PAYCURRENCYENUM_DKK PayCurrencyEnum = "DKK"
	PAYCURRENCYENUM_DJF PayCurrencyEnum = "DJF"
	PAYCURRENCYENUM_DOP PayCurrencyEnum = "DOP"
	PAYCURRENCYENUM_NLG PayCurrencyEnum = "NLG"
	PAYCURRENCYENUM_XCD PayCurrencyEnum = "XCD"
	PAYCURRENCYENUM_DDM PayCurrencyEnum = "DDM"
	PAYCURRENCYENUM_ECS PayCurrencyEnum = "ECS"
	PAYCURRENCYENUM_ECV PayCurrencyEnum = "ECV"
	PAYCURRENCYENUM_EGP PayCurrencyEnum = "EGP"
	PAYCURRENCYENUM_GQE PayCurrencyEnum = "GQE"
	PAYCURRENCYENUM_ERN PayCurrencyEnum = "ERN"
	PAYCURRENCYENUM_EEK PayCurrencyEnum = "EEK"
	PAYCURRENCYENUM_ETB PayCurrencyEnum = "ETB"
	PAYCURRENCYENUM_EUR PayCurrencyEnum = "EUR"
	PAYCURRENCYENUM_XBA PayCurrencyEnum = "XBA"
	PAYCURRENCYENUM_XEU PayCurrencyEnum = "XEU"
	PAYCURRENCYENUM_XBB PayCurrencyEnum = "XBB"
	PAYCURRENCYENUM_XBC PayCurrencyEnum = "XBC"
	PAYCURRENCYENUM_XBD PayCurrencyEnum = "XBD"
	PAYCURRENCYENUM_FKP PayCurrencyEnum = "FKP"
	PAYCURRENCYENUM_FJD PayCurrencyEnum = "FJD"
	PAYCURRENCYENUM_FIM PayCurrencyEnum = "FIM"
	PAYCURRENCYENUM_FRF PayCurrencyEnum = "FRF"
	PAYCURRENCYENUM_XFO PayCurrencyEnum = "XFO"
	PAYCURRENCYENUM_XFU PayCurrencyEnum = "XFU"
	PAYCURRENCYENUM_GMD PayCurrencyEnum = "GMD"
	PAYCURRENCYENUM_GEK PayCurrencyEnum = "GEK"
	PAYCURRENCYENUM_GEL PayCurrencyEnum = "GEL"
	PAYCURRENCYENUM_DEM PayCurrencyEnum = "DEM"
	PAYCURRENCYENUM_GHS PayCurrencyEnum = "GHS"
	PAYCURRENCYENUM_GHC PayCurrencyEnum = "GHC"
	PAYCURRENCYENUM_GIP PayCurrencyEnum = "GIP"
	PAYCURRENCYENUM_XAU PayCurrencyEnum = "XAU"
	PAYCURRENCYENUM_GRD PayCurrencyEnum = "GRD"
	PAYCURRENCYENUM_GTQ PayCurrencyEnum = "GTQ"
	PAYCURRENCYENUM_GWP PayCurrencyEnum = "GWP"
	PAYCURRENCYENUM_GNF PayCurrencyEnum = "GNF"
	PAYCURRENCYENUM_GNS PayCurrencyEnum = "GNS"
	PAYCURRENCYENUM_GYD PayCurrencyEnum = "GYD"
	PAYCURRENCYENUM_HTG PayCurrencyEnum = "HTG"
	PAYCURRENCYENUM_HNL PayCurrencyEnum = "HNL"
	PAYCURRENCYENUM_HKD PayCurrencyEnum = "HKD"
	PAYCURRENCYENUM_HUF PayCurrencyEnum = "HUF"
	PAYCURRENCYENUM_IMP PayCurrencyEnum = "IMP"
	PAYCURRENCYENUM_ISK PayCurrencyEnum = "ISK"
	PAYCURRENCYENUM_ISJ PayCurrencyEnum = "ISJ"
	PAYCURRENCYENUM_INR PayCurrencyEnum = "INR"
	PAYCURRENCYENUM_IDR PayCurrencyEnum = "IDR"
	PAYCURRENCYENUM_IRR PayCurrencyEnum = "IRR"
	PAYCURRENCYENUM_IQD PayCurrencyEnum = "IQD"
	PAYCURRENCYENUM_IEP PayCurrencyEnum = "IEP"
	PAYCURRENCYENUM_ILS PayCurrencyEnum = "ILS"
	PAYCURRENCYENUM_ILP PayCurrencyEnum = "ILP"
	PAYCURRENCYENUM_ILR PayCurrencyEnum = "ILR"
	PAYCURRENCYENUM_ITL PayCurrencyEnum = "ITL"
	PAYCURRENCYENUM_JMD PayCurrencyEnum = "JMD"
	PAYCURRENCYENUM_JPY PayCurrencyEnum = "JPY"
	PAYCURRENCYENUM_JOD PayCurrencyEnum = "JOD"
	PAYCURRENCYENUM_KZT PayCurrencyEnum = "KZT"
	PAYCURRENCYENUM_KES PayCurrencyEnum = "KES"
	PAYCURRENCYENUM_KWD PayCurrencyEnum = "KWD"
	PAYCURRENCYENUM_KGS PayCurrencyEnum = "KGS"
	PAYCURRENCYENUM_LAK PayCurrencyEnum = "LAK"
	PAYCURRENCYENUM_LVL PayCurrencyEnum = "LVL"
	PAYCURRENCYENUM_LVR PayCurrencyEnum = "LVR"
	PAYCURRENCYENUM_LBP PayCurrencyEnum = "LBP"
	PAYCURRENCYENUM_LSL PayCurrencyEnum = "LSL"
	PAYCURRENCYENUM_LRD PayCurrencyEnum = "LRD"
	PAYCURRENCYENUM_LYD PayCurrencyEnum = "LYD"
	PAYCURRENCYENUM_LTL PayCurrencyEnum = "LTL"
	PAYCURRENCYENUM_LTT PayCurrencyEnum = "LTT"
	PAYCURRENCYENUM_LUL PayCurrencyEnum = "LUL"
	PAYCURRENCYENUM_LUC PayCurrencyEnum = "LUC"
	PAYCURRENCYENUM_LUF PayCurrencyEnum = "LUF"
	PAYCURRENCYENUM_MOP PayCurrencyEnum = "MOP"
	PAYCURRENCYENUM_MKD PayCurrencyEnum = "MKD"
	PAYCURRENCYENUM_MKN PayCurrencyEnum = "MKN"
	PAYCURRENCYENUM_MGA PayCurrencyEnum = "MGA"
	PAYCURRENCYENUM_MGF PayCurrencyEnum = "MGF"
	PAYCURRENCYENUM_MWK PayCurrencyEnum = "MWK"
	PAYCURRENCYENUM_MYR PayCurrencyEnum = "MYR"
	PAYCURRENCYENUM_MVR PayCurrencyEnum = "MVR"
	PAYCURRENCYENUM_MVP PayCurrencyEnum = "MVP"
	PAYCURRENCYENUM_MLF PayCurrencyEnum = "MLF"
	PAYCURRENCYENUM_MTL PayCurrencyEnum = "MTL"
	PAYCURRENCYENUM_MTP PayCurrencyEnum = "MTP"
	PAYCURRENCYENUM_MRU PayCurrencyEnum = "MRU"
	PAYCURRENCYENUM_MRO PayCurrencyEnum = "MRO"
	PAYCURRENCYENUM_MUR PayCurrencyEnum = "MUR"
	PAYCURRENCYENUM_MXV PayCurrencyEnum = "MXV"
	PAYCURRENCYENUM_MXN PayCurrencyEnum = "MXN"
	PAYCURRENCYENUM_MXP PayCurrencyEnum = "MXP"
	PAYCURRENCYENUM_MDC PayCurrencyEnum = "MDC"
	PAYCURRENCYENUM_MDL PayCurrencyEnum = "MDL"
	PAYCURRENCYENUM_MCF PayCurrencyEnum = "MCF"
	PAYCURRENCYENUM_MNT PayCurrencyEnum = "MNT"
	PAYCURRENCYENUM_MAD PayCurrencyEnum = "MAD"
	PAYCURRENCYENUM_MAF PayCurrencyEnum = "MAF"
	PAYCURRENCYENUM_MZE PayCurrencyEnum = "MZE"
	PAYCURRENCYENUM_MZN PayCurrencyEnum = "MZN"
	PAYCURRENCYENUM_MZM PayCurrencyEnum = "MZM"
	PAYCURRENCYENUM_MMK PayCurrencyEnum = "MMK"
	PAYCURRENCYENUM_NAD PayCurrencyEnum = "NAD"
	PAYCURRENCYENUM_NPR PayCurrencyEnum = "NPR"
	PAYCURRENCYENUM_ANG PayCurrencyEnum = "ANG"
	PAYCURRENCYENUM_TWD PayCurrencyEnum = "TWD"
	PAYCURRENCYENUM_NZD PayCurrencyEnum = "NZD"
	PAYCURRENCYENUM_NIO PayCurrencyEnum = "NIO"
	PAYCURRENCYENUM_NIC PayCurrencyEnum = "NIC"
	PAYCURRENCYENUM_NGN PayCurrencyEnum = "NGN"
	PAYCURRENCYENUM_KPW PayCurrencyEnum = "KPW"
	PAYCURRENCYENUM_NOK PayCurrencyEnum = "NOK"
	PAYCURRENCYENUM_OMR PayCurrencyEnum = "OMR"
	PAYCURRENCYENUM_PKR PayCurrencyEnum = "PKR"
	PAYCURRENCYENUM_XPD PayCurrencyEnum = "XPD"
	PAYCURRENCYENUM_PAB PayCurrencyEnum = "PAB"
	PAYCURRENCYENUM_PGK PayCurrencyEnum = "PGK"
	PAYCURRENCYENUM_PYG PayCurrencyEnum = "PYG"
	PAYCURRENCYENUM_PEI PayCurrencyEnum = "PEI"
	PAYCURRENCYENUM_PEN PayCurrencyEnum = "PEN"
	PAYCURRENCYENUM_PES PayCurrencyEnum = "PES"
	PAYCURRENCYENUM_PHP PayCurrencyEnum = "PHP"
	PAYCURRENCYENUM_XPT PayCurrencyEnum = "XPT"
	PAYCURRENCYENUM_PLN PayCurrencyEnum = "PLN"
	PAYCURRENCYENUM_PLZ PayCurrencyEnum = "PLZ"
	PAYCURRENCYENUM_PTE PayCurrencyEnum = "PTE"
	PAYCURRENCYENUM_GWE PayCurrencyEnum = "GWE"
	PAYCURRENCYENUM_QAR PayCurrencyEnum = "QAR"
	PAYCURRENCYENUM_XRE PayCurrencyEnum = "XRE"
	PAYCURRENCYENUM_RHD PayCurrencyEnum = "RHD"
	PAYCURRENCYENUM_RON PayCurrencyEnum = "RON"
	PAYCURRENCYENUM_ROL PayCurrencyEnum = "ROL"
	PAYCURRENCYENUM_RUB PayCurrencyEnum = "RUB"
	PAYCURRENCYENUM_RUR PayCurrencyEnum = "RUR"
	PAYCURRENCYENUM_RWF PayCurrencyEnum = "RWF"
	PAYCURRENCYENUM_SVC PayCurrencyEnum = "SVC"
	PAYCURRENCYENUM_WST PayCurrencyEnum = "WST"
	PAYCURRENCYENUM_SAR PayCurrencyEnum = "SAR"
	PAYCURRENCYENUM_RSD PayCurrencyEnum = "RSD"
	PAYCURRENCYENUM_CSD PayCurrencyEnum = "CSD"
	PAYCURRENCYENUM_SCR PayCurrencyEnum = "SCR"
	PAYCURRENCYENUM_SLL PayCurrencyEnum = "SLL"
	PAYCURRENCYENUM_XAG PayCurrencyEnum = "XAG"
	PAYCURRENCYENUM_SGD PayCurrencyEnum = "SGD"
	PAYCURRENCYENUM_SKK PayCurrencyEnum = "SKK"
	PAYCURRENCYENUM_SIT PayCurrencyEnum = "SIT"
	PAYCURRENCYENUM_SBD PayCurrencyEnum = "SBD"
	PAYCURRENCYENUM_SOS PayCurrencyEnum = "SOS"
	PAYCURRENCYENUM_ZAR PayCurrencyEnum = "ZAR"
	PAYCURRENCYENUM_ZAL PayCurrencyEnum = "ZAL"
	PAYCURRENCYENUM_KRH PayCurrencyEnum = "KRH"
	PAYCURRENCYENUM_KRW PayCurrencyEnum = "KRW"
	PAYCURRENCYENUM_KRO PayCurrencyEnum = "KRO"
	PAYCURRENCYENUM_SSP PayCurrencyEnum = "SSP"
	PAYCURRENCYENUM_SUR PayCurrencyEnum = "SUR"
	PAYCURRENCYENUM_ESP PayCurrencyEnum = "ESP"
	PAYCURRENCYENUM_ESA PayCurrencyEnum = "ESA"
	PAYCURRENCYENUM_ESB PayCurrencyEnum = "ESB"
	PAYCURRENCYENUM_XDR PayCurrencyEnum = "XDR"
	PAYCURRENCYENUM_LKR PayCurrencyEnum = "LKR"
	PAYCURRENCYENUM_SHP PayCurrencyEnum = "SHP"
	PAYCURRENCYENUM_XSU PayCurrencyEnum = "XSU"
	PAYCURRENCYENUM_SDD PayCurrencyEnum = "SDD"
	PAYCURRENCYENUM_SDG PayCurrencyEnum = "SDG"
	PAYCURRENCYENUM_SDP PayCurrencyEnum = "SDP"
	PAYCURRENCYENUM_SRD PayCurrencyEnum = "SRD"
	PAYCURRENCYENUM_SRG PayCurrencyEnum = "SRG"
	PAYCURRENCYENUM_SZL PayCurrencyEnum = "SZL"
	PAYCURRENCYENUM_SEK PayCurrencyEnum = "SEK"
	PAYCURRENCYENUM_CHF PayCurrencyEnum = "CHF"
	PAYCURRENCYENUM_SYP PayCurrencyEnum = "SYP"
	PAYCURRENCYENUM_STN PayCurrencyEnum = "STN"
	PAYCURRENCYENUM_STD PayCurrencyEnum = "STD"
	PAYCURRENCYENUM_TVD PayCurrencyEnum = "TVD"
	PAYCURRENCYENUM_TJR PayCurrencyEnum = "TJR"
	PAYCURRENCYENUM_TJS PayCurrencyEnum = "TJS"
	PAYCURRENCYENUM_TZS PayCurrencyEnum = "TZS"
	PAYCURRENCYENUM_XTS PayCurrencyEnum = "XTS"
	PAYCURRENCYENUM_THB PayCurrencyEnum = "THB"
	PAYCURRENCYENUM_XXX PayCurrencyEnum = "XXX"
	PAYCURRENCYENUM_TPE PayCurrencyEnum = "TPE"
	PAYCURRENCYENUM_TOP PayCurrencyEnum = "TOP"
	PAYCURRENCYENUM_TTD PayCurrencyEnum = "TTD"
	PAYCURRENCYENUM_TND PayCurrencyEnum = "TND"
	PAYCURRENCYENUM_TRY PayCurrencyEnum = "TRY"
	PAYCURRENCYENUM_TRL PayCurrencyEnum = "TRL"
	PAYCURRENCYENUM_TMT PayCurrencyEnum = "TMT"
	PAYCURRENCYENUM_TMM PayCurrencyEnum = "TMM"
	PAYCURRENCYENUM_USD PayCurrencyEnum = "USD"
	PAYCURRENCYENUM_USN PayCurrencyEnum = "USN"
	PAYCURRENCYENUM_USS PayCurrencyEnum = "USS"
	PAYCURRENCYENUM_UGX PayCurrencyEnum = "UGX"
	PAYCURRENCYENUM_UGS PayCurrencyEnum = "UGS"
	PAYCURRENCYENUM_UAH PayCurrencyEnum = "UAH"
	PAYCURRENCYENUM_UAK PayCurrencyEnum = "UAK"
	PAYCURRENCYENUM_AED PayCurrencyEnum = "AED"
	PAYCURRENCYENUM_UYW PayCurrencyEnum = "UYW"
	PAYCURRENCYENUM_UYU PayCurrencyEnum = "UYU"
	PAYCURRENCYENUM_UYP PayCurrencyEnum = "UYP"
	PAYCURRENCYENUM_UYI PayCurrencyEnum = "UYI"
	PAYCURRENCYENUM_UZS PayCurrencyEnum = "UZS"
	PAYCURRENCYENUM_VUV PayCurrencyEnum = "VUV"
	PAYCURRENCYENUM_VES PayCurrencyEnum = "VES"
	PAYCURRENCYENUM_VEB PayCurrencyEnum = "VEB"
	PAYCURRENCYENUM_VEF PayCurrencyEnum = "VEF"
	PAYCURRENCYENUM_VND PayCurrencyEnum = "VND"
	PAYCURRENCYENUM_VNN PayCurrencyEnum = "VNN"
	PAYCURRENCYENUM_CHE PayCurrencyEnum = "CHE"
	PAYCURRENCYENUM_CHW PayCurrencyEnum = "CHW"
	PAYCURRENCYENUM_XOF PayCurrencyEnum = "XOF"
	PAYCURRENCYENUM_YDD PayCurrencyEnum = "YDD"
	PAYCURRENCYENUM_YER PayCurrencyEnum = "YER"
	PAYCURRENCYENUM_YUN PayCurrencyEnum = "YUN"
	PAYCURRENCYENUM_YUD PayCurrencyEnum = "YUD"
	PAYCURRENCYENUM_YUM PayCurrencyEnum = "YUM"
	PAYCURRENCYENUM_YUR PayCurrencyEnum = "YUR"
	PAYCURRENCYENUM_ZWN PayCurrencyEnum = "ZWN"
	PAYCURRENCYENUM_ZRN PayCurrencyEnum = "ZRN"
	PAYCURRENCYENUM_ZRZ PayCurrencyEnum = "ZRZ"
	PAYCURRENCYENUM_ZMW PayCurrencyEnum = "ZMW"
	PAYCURRENCYENUM_ZMK PayCurrencyEnum = "ZMK"
	PAYCURRENCYENUM_ZWD PayCurrencyEnum = "ZWD"
	PAYCURRENCYENUM_ZWR PayCurrencyEnum = "ZWR"
	PAYCURRENCYENUM_ZWL PayCurrencyEnum = "ZWL"
)

var allowedPayCurrencyEnumEnumValues = []PayCurrencyEnum{
	"XUA",
	"AFN",
	"AFA",
	"ALL",
	"ALK",
	"DZD",
	"ADP",
	"AOA",
	"AOK",
	"AON",
	"AOR",
	"ARA",
	"ARS",
	"ARM",
	"ARP",
	"ARL",
	"AMD",
	"AWG",
	"AUD",
	"ATS",
	"AZN",
	"AZM",
	"BSD",
	"BHD",
	"BDT",
	"BBD",
	"BYN",
	"BYB",
	"BYR",
	"BEF",
	"BEC",
	"BEL",
	"BZD",
	"BMD",
	"BTN",
	"BOB",
	"BOL",
	"BOV",
	"BOP",
	"BAM",
	"BAD",
	"BAN",
	"BWP",
	"BRC",
	"BRZ",
	"BRE",
	"BRR",
	"BRN",
	"BRB",
	"BRL",
	"GBP",
	"BND",
	"BGL",
	"BGN",
	"BGO",
	"BGM",
	"BUK",
	"BIF",
	"XPF",
	"KHR",
	"CAD",
	"CVE",
	"KYD",
	"XAF",
	"CLE",
	"CLP",
	"CLF",
	"CNX",
	"CNY",
	"CNH",
	"COP",
	"COU",
	"KMF",
	"CDF",
	"CRC",
	"HRD",
	"HRK",
	"CUC",
	"CUP",
	"CYP",
	"CZK",
	"CSK",
	"DKK",
	"DJF",
	"DOP",
	"NLG",
	"XCD",
	"DDM",
	"ECS",
	"ECV",
	"EGP",
	"GQE",
	"ERN",
	"EEK",
	"ETB",
	"EUR",
	"XBA",
	"XEU",
	"XBB",
	"XBC",
	"XBD",
	"FKP",
	"FJD",
	"FIM",
	"FRF",
	"XFO",
	"XFU",
	"GMD",
	"GEK",
	"GEL",
	"DEM",
	"GHS",
	"GHC",
	"GIP",
	"XAU",
	"GRD",
	"GTQ",
	"GWP",
	"GNF",
	"GNS",
	"GYD",
	"HTG",
	"HNL",
	"HKD",
	"HUF",
	"IMP",
	"ISK",
	"ISJ",
	"INR",
	"IDR",
	"IRR",
	"IQD",
	"IEP",
	"ILS",
	"ILP",
	"ILR",
	"ITL",
	"JMD",
	"JPY",
	"JOD",
	"KZT",
	"KES",
	"KWD",
	"KGS",
	"LAK",
	"LVL",
	"LVR",
	"LBP",
	"LSL",
	"LRD",
	"LYD",
	"LTL",
	"LTT",
	"LUL",
	"LUC",
	"LUF",
	"MOP",
	"MKD",
	"MKN",
	"MGA",
	"MGF",
	"MWK",
	"MYR",
	"MVR",
	"MVP",
	"MLF",
	"MTL",
	"MTP",
	"MRU",
	"MRO",
	"MUR",
	"MXV",
	"MXN",
	"MXP",
	"MDC",
	"MDL",
	"MCF",
	"MNT",
	"MAD",
	"MAF",
	"MZE",
	"MZN",
	"MZM",
	"MMK",
	"NAD",
	"NPR",
	"ANG",
	"TWD",
	"NZD",
	"NIO",
	"NIC",
	"NGN",
	"KPW",
	"NOK",
	"OMR",
	"PKR",
	"XPD",
	"PAB",
	"PGK",
	"PYG",
	"PEI",
	"PEN",
	"PES",
	"PHP",
	"XPT",
	"PLN",
	"PLZ",
	"PTE",
	"GWE",
	"QAR",
	"XRE",
	"RHD",
	"RON",
	"ROL",
	"RUB",
	"RUR",
	"RWF",
	"SVC",
	"WST",
	"SAR",
	"RSD",
	"CSD",
	"SCR",
	"SLL",
	"XAG",
	"SGD",
	"SKK",
	"SIT",
	"SBD",
	"SOS",
	"ZAR",
	"ZAL",
	"KRH",
	"KRW",
	"KRO",
	"SSP",
	"SUR",
	"ESP",
	"ESA",
	"ESB",
	"XDR",
	"LKR",
	"SHP",
	"XSU",
	"SDD",
	"SDG",
	"SDP",
	"SRD",
	"SRG",
	"SZL",
	"SEK",
	"CHF",
	"SYP",
	"STN",
	"STD",
	"TVD",
	"TJR",
	"TJS",
	"TZS",
	"XTS",
	"THB",
	"XXX",
	"TPE",
	"TOP",
	"TTD",
	"TND",
	"TRY",
	"TRL",
	"TMT",
	"TMM",
	"USD",
	"USN",
	"USS",
	"UGX",
	"UGS",
	"UAH",
	"UAK",
	"AED",
	"UYW",
	"UYU",
	"UYP",
	"UYI",
	"UZS",
	"VUV",
	"VES",
	"VEB",
	"VEF",
	"VND",
	"VNN",
	"CHE",
	"CHW",
	"XOF",
	"YDD",
	"YER",
	"YUN",
	"YUD",
	"YUM",
	"YUR",
	"ZWN",
	"ZRN",
	"ZRZ",
	"ZMW",
	"ZMK",
	"ZWD",
	"ZWR",
	"ZWL",
}

func (v *PayCurrencyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PayCurrencyEnum(value)
	for _, existing := range allowedPayCurrencyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = PAYCURRENCYENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewPayCurrencyEnumFromValue returns a pointer to a valid PayCurrencyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPayCurrencyEnumFromValue(v string) (*PayCurrencyEnum, error) {
	ev := PayCurrencyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := PAYCURRENCYENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PayCurrencyEnum) IsValid() bool {
	for _, existing := range allowedPayCurrencyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PayCurrencyEnum value
func (v PayCurrencyEnum) Ptr() *PayCurrencyEnum {
	return &v
}

type NullablePayCurrencyEnum struct {
	value *PayCurrencyEnum
	isSet bool
}

func (v NullablePayCurrencyEnum) Get() *PayCurrencyEnum {
	return v.value
}

func (v *NullablePayCurrencyEnum) Set(val *PayCurrencyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePayCurrencyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePayCurrencyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayCurrencyEnum(val *PayCurrencyEnum) *NullablePayCurrencyEnum {
	return &NullablePayCurrencyEnum{value: val, isSet: true}
}

func (v NullablePayCurrencyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayCurrencyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

