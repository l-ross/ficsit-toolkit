package ConsumableEquipment

import (
	"fmt"
)

type FGConsumableEquipment struct {
	ClassName              string
	MAnimData              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCanPress              bool
	MCostToUse             string
	MCurrentAnimData       string
	MEquipmentSlot         string
	MHasPersistentOwner    bool
	MRandomAnim            int
	MUseDefaultPrimaryFire bool
}

var (
	Beacon = FGConsumableEquipment{
		ClassName:              "Equip_Beacon_C",
		MArmAnimation:          `AE_Generic1Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             ``,
		MEquipmentSlot:         `ES_ARMS`,
		MHasPersistentOwner:    false,
		MUseDefaultPrimaryFire: false,
	}

	ConsumeableEquipment = FGConsumableEquipment{
		ClassName:              "BP_ConsumeableEquipment_C",
		MArmAnimation:          `AE_Generic1Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MCanPress:              false,
		MCostToUse:             ``,
		MEquipmentSlot:         `ES_ARMS`,
		MHasPersistentOwner:    false,
		MRandomAnim:            0,
		MUseDefaultPrimaryFire: false,
	}

	MedKit = FGConsumableEquipment{
		ClassName:              "Equip_MedKit_C",
		MAnimData:              `((Montage_7_2E66F6A948A8606E71185682EA2AC4EC=AnimMontage'"/Game/FactoryGame/Character/Player/Animation/FirstPerson/MedkitUse_01_Montage.MedkitUse_01_Montage"',CameraAnim_8_AA01C2B248FF438D6C2816B2FA94F1BD=CameraAnim'"/Game/FactoryGame/Character/Player/CameraShake/MedkitUse_01_CameraAnim.MedkitUse_01_CameraAnim"'),(Montage_7_2E66F6A948A8606E71185682EA2AC4EC=AnimMontage'"/Game/FactoryGame/Character/Player/Animation/FirstPerson/MedkitUse_02_Montage.MedkitUse_02_Montage"',CameraAnim_8_AA01C2B248FF438D6C2816B2FA94F1BD=CameraAnim'"/Game/FactoryGame/Character/Player/CameraShake/MedkitUse_02_CameraAnim.MedkitUse_02_CameraAnim"'),(Montage_7_2E66F6A948A8606E71185682EA2AC4EC=AnimMontage'"/Game/FactoryGame/Character/Player/Animation/FirstPerson/MedkitUse_03_Montage.MedkitUse_03_Montage"',CameraAnim_8_AA01C2B248FF438D6C2816B2FA94F1BD=CameraAnim'"/Game/FactoryGame/Character/Player/CameraShake/MedkitUse_03_CameraAnim.MedkitUse_03_CameraAnim"'))`,
		MArmAnimation:          `AE_Generic1Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MCanPress:              false,
		MCostToUse:             ``,
		MCurrentAnimData:       `()`,
		MEquipmentSlot:         `ES_ARMS`,
		MHasPersistentOwner:    false,
		MRandomAnim:            0,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGConsumableEquipment, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGConsumableEquipment with name %s", className)
}

var classNameToVar = map[string]*FGConsumableEquipment{
	"Equip_Beacon_C":            &Beacon,
	"BP_ConsumeableEquipment_C": &ConsumeableEquipment,
	"Equip_MedKit_C":            &MedKit,
}
