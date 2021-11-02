// Code generated by ../../gen/docs_json. DO NOT EDIT.

package EquipmentStunSpear

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/factory/resource"
)

type FGEquipmentStunSpear struct {
	Name                      string
	ClassName                 string
	MArmAnimation             string
	MAttachSocket             string
	MAttackDistance           float64
	MBackAnimation            string
	MCostToUse                string
	MCurrentMontageSection    string
	MDamage                   int
	MEquipmentSlot            resource.EquipmentSlot
	MFirstAttackTimer         float64
	MHasPersistentOwner       bool
	MPlayingSound             bool
	MRandomAttackAnim         int
	MRandomEquipAnim          int
	MRandomStingerAnim        int
	MSecondAttackTimer        float64
	MSecondSwingCooldDownTime float64
	MSecondSwingMaxTime       float64
	MUseDefaultPrimaryFire    bool
}

var (
	CandyCaneBasher = FGEquipmentStunSpear{
		Name:                      "CandyCaneBasher",
		ClassName:                 "Equip_CandyCaneBasher_C",
		MArmAnimation:             `AE_StunSpear`,
		MAttachSocket:             `hand_rSocket`,
		MAttackDistance:           400.000000,
		MBackAnimation:            `BE_None`,
		MCostToUse:                ``,
		MCurrentMontageSection:    `None`,
		MDamage:                   7,
		MEquipmentSlot:            resource.Arms,
		MFirstAttackTimer:         0.000000,
		MHasPersistentOwner:       false,
		MRandomAttackAnim:         0,
		MRandomEquipAnim:          0,
		MSecondAttackTimer:        0.000000,
		MSecondSwingCooldDownTime: 0.700000,
		MSecondSwingMaxTime:       0.600000,
		MUseDefaultPrimaryFire:    false,
	}

	ShockShank = FGEquipmentStunSpear{
		Name:                      "ShockShank",
		ClassName:                 "Equip_ShockShank_C",
		MArmAnimation:             `AE_ShockShank`,
		MAttachSocket:             `hand_rSocket`,
		MAttackDistance:           350.000000,
		MBackAnimation:            `BE_None`,
		MCostToUse:                ``,
		MDamage:                   5,
		MEquipmentSlot:            resource.Arms,
		MHasPersistentOwner:       false,
		MPlayingSound:             false,
		MRandomAttackAnim:         0,
		MRandomStingerAnim:        0,
		MSecondSwingCooldDownTime: 1.000000,
		MSecondSwingMaxTime:       1.100000,
		MUseDefaultPrimaryFire:    false,
	}

	StunSpear = FGEquipmentStunSpear{
		Name:                      "StunSpear",
		ClassName:                 "Equip_StunSpear_C",
		MArmAnimation:             `AE_StunSpear`,
		MAttachSocket:             `hand_rSocket`,
		MAttackDistance:           400.000000,
		MBackAnimation:            `BE_None`,
		MCostToUse:                ``,
		MCurrentMontageSection:    `None`,
		MDamage:                   7,
		MEquipmentSlot:            resource.Arms,
		MFirstAttackTimer:         0.000000,
		MHasPersistentOwner:       false,
		MRandomAttackAnim:         0,
		MRandomEquipAnim:          0,
		MSecondAttackTimer:        0.000000,
		MSecondSwingCooldDownTime: 0.700000,
		MSecondSwingMaxTime:       0.600000,
		MUseDefaultPrimaryFire:    false,
	}
)

func GetByClassName(className string) (FGEquipmentStunSpear, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return FGEquipmentStunSpear{}, fmt.Errorf("failed to find FGEquipmentStunSpear with class name %s", className)
}

var classNameToVar = map[string]FGEquipmentStunSpear{
	"Equip_CandyCaneBasher_C": CandyCaneBasher,
	"Equip_ShockShank_C":      ShockShank,
	"Equip_StunSpear_C":       StunSpear,
}
