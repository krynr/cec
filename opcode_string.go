// Code generated by "stringer -type=OpCode"; DO NOT EDIT.

package cec

import "strconv"

const _OpCode_name = "OpFeatureAbortOpImageViewOnOpTunerStepIncrementOpTunerStepDecrementOpTunerDeviceStatusOpGiveTunerDeviceStatusOpRecordOnOpRecordStatusOpRecordOffOpTextViewOnOpRecordTVScreenOpGiveDeckStatusOpDeckStatusOpSetMenuLanguageOpClearAnalogTimerOpSetAnalogTimerOpTimerStatusOpStandbyOpPlayOpDeckControlOpTimerClearedStatusOpUserControlPressedOpUserControlReleasedOpGiveOSDNameOpSetOSDNameOpSetOSDStringOpSetTimerProgramTitleOpSystemAudioModeRequestOpGiveAudioStatusOpSetSystemAudioModeOpReportAudioStatusOpGiveSystemAudioModeStatusOpSystemAudioModeStatusOpRoutingChangeOpRoutingInformationOpActiveSourceOpGivePhysicalAddressOpReportPhysicalAddressOpRequestActiveSourceOpSetStreamPathOpDeviceVendorIDOpVendorCommandOpVendorRemoteButtonDownOpVendorRemoteButtonUpOpGiveDeviceVendorIDOpMenuRequestOpMenuStatusOpGiveDevicePowerStatusOpReportPowerStatusOpGetMenuLanguageOpSelectAnalogServiceOpSelectDigitalServiceOpSetDigitalTimerOpClearDigitalTimerOpSetAudioRateOpInactiveSourceOpCECVersionOpGetCECVersionOpVendorCommandWithIDOpClearExternalTimerOpSetExternalTimerOpAbort"

var _OpCode_map = map[OpCode]string{
	0:   _OpCode_name[0:14],
	4:   _OpCode_name[14:27],
	5:   _OpCode_name[27:47],
	6:   _OpCode_name[47:67],
	7:   _OpCode_name[67:86],
	8:   _OpCode_name[86:109],
	9:   _OpCode_name[109:119],
	10:  _OpCode_name[119:133],
	11:  _OpCode_name[133:144],
	13:  _OpCode_name[144:156],
	15:  _OpCode_name[156:172],
	26:  _OpCode_name[172:188],
	27:  _OpCode_name[188:200],
	50:  _OpCode_name[200:217],
	51:  _OpCode_name[217:235],
	52:  _OpCode_name[235:251],
	53:  _OpCode_name[251:264],
	54:  _OpCode_name[264:273],
	65:  _OpCode_name[273:279],
	66:  _OpCode_name[279:292],
	67:  _OpCode_name[292:312],
	68:  _OpCode_name[312:332],
	69:  _OpCode_name[332:353],
	70:  _OpCode_name[353:366],
	71:  _OpCode_name[366:378],
	100: _OpCode_name[378:392],
	103: _OpCode_name[392:414],
	112: _OpCode_name[414:438],
	113: _OpCode_name[438:455],
	114: _OpCode_name[455:475],
	122: _OpCode_name[475:494],
	125: _OpCode_name[494:521],
	126: _OpCode_name[521:544],
	128: _OpCode_name[544:559],
	129: _OpCode_name[559:579],
	130: _OpCode_name[579:593],
	131: _OpCode_name[593:614],
	132: _OpCode_name[614:637],
	133: _OpCode_name[637:658],
	134: _OpCode_name[658:673],
	135: _OpCode_name[673:689],
	137: _OpCode_name[689:704],
	138: _OpCode_name[704:728],
	139: _OpCode_name[728:750],
	140: _OpCode_name[750:770],
	141: _OpCode_name[770:783],
	142: _OpCode_name[783:795],
	143: _OpCode_name[795:818],
	144: _OpCode_name[818:837],
	145: _OpCode_name[837:854],
	146: _OpCode_name[854:875],
	147: _OpCode_name[875:897],
	151: _OpCode_name[897:914],
	153: _OpCode_name[914:933],
	154: _OpCode_name[933:947],
	157: _OpCode_name[947:963],
	158: _OpCode_name[963:975],
	159: _OpCode_name[975:990],
	160: _OpCode_name[990:1011],
	161: _OpCode_name[1011:1031],
	162: _OpCode_name[1031:1049],
	255: _OpCode_name[1049:1056],
}

func (i OpCode) String() string {
	if str, ok := _OpCode_map[i]; ok {
		return str
	}
	return "OpCode(" + strconv.FormatInt(int64(i), 10) + ")"
}
