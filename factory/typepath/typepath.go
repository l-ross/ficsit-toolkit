package typepath

type TypePath string

const (
	//
	// Logistics
	//

	ConveyorBeltMk1 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C"
	ConveyorBeltMk2 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk2/Build_ConveyorBeltMk2.Build_ConveyorBeltMk2_C"
	ConveyorBeltMk3 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk3/Build_ConveyorBeltMk3.Build_ConveyorBeltMk3_C"
	ConveyorBeltMk4 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk4/Build_ConveyorBeltMk4.Build_ConveyorBeltMk4_C"
	ConveyorBeltMk5 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk5/Build_ConveyorBeltMk5.Build_ConveyorBeltMk5_C"

	ConveyorLiftMk1 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk1/Build_ConveyorLiftMk1.Build_ConveyorLiftMk1_C"
	ConveyorLiftMk2 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk2/Build_ConveyorLiftMk2.Build_ConveyorLiftMk2_C"
	ConveyorLiftMk3 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk3/Build_ConveyorLiftMk3.Build_ConveyorLiftMk3_C"
	ConveyorLiftMk4 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk4/Build_ConveyorLiftMk4.Build_ConveyorLiftMk4_C"
	ConveyorLiftMk5 TypePath = "/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk5/Build_ConveyorLiftMk5.Build_ConveyorLiftMk5_C"

	Merger               TypePath = "/Game/FactoryGame/Buildable/Factory/CA_Merger/Build_ConveyorAttachmentMerger.Build_ConveyorAttachmentMerger_C"
	Splitter             TypePath = "/Game/FactoryGame/Buildable/Factory/CA_Splitter/Build_ConveyorAttachmentSplitter.Build_ConveyorAttachmentSplitter_C"
	SmartSplitter        TypePath = "/Game/FactoryGame/Buildable/Factory/CA_SplitterSmart/Build_ConveyorAttachmentSplitterSmart.Build_ConveyorAttachmentSplitterSmart_C"
	ProgrammableSplitter TypePath = "/Game/FactoryGame/Buildable/Factory/CA_SplitterProgrammable/Build_ConveyorAttachmentSplitterProgrammable.Build_ConveyorAttachmentSplitterProgrammable_C"

	PipelineMk1           TypePath = "/Game/FactoryGame/Buildable/Factory/Pipeline/Build_Pipeline.Build_Pipeline_C"
	PipelineMk2           TypePath = "/Game/FactoryGame/Buildable/Factory/PipelineMk2/Build_PipelineMK2.Build_PipelineMK2_C"
	PipelineJunctionCross TypePath = "/Game/FactoryGame/Buildable/Factory/PipeJunction/Build_PipelineJunction_Cross.Build_PipelineJunction_Cross_C"
	PipelinePumpMk1       TypePath = "/Game/FactoryGame/Buildable/Factory/PipePump/Build_PipelinePump.Build_PipelinePump_C"
	PipelinePumpMk2       TypePath = "/Game/FactoryGame/Buildable/Factory/PipePumpMk2/Build_PipelinePumpMK2.Build_PipelinePumpMk2_C"
	Valve                 TypePath = "/Game/FactoryGame/Buildable/Factory/PipeValve/Build_Valve.Build_Valve_C"

	//
	// Storage
	//

	StorageContainerMk1 TypePath = "/Game/FactoryGame/Buildable/Factory/StorageContainerMk1/Build_StorageContainerMk1.Build_StorageContainerMk1_C"
	StorageContainerMk2 TypePath = "/Game/FactoryGame/Buildable/Factory/StorageContainerMk2/Build_StorageContainerMk2.Build_StorageContainerMk2_C"

	//
	// Production
	//

	Constructor         TypePath = "/Game/FactoryGame/Buildable/Factory/ConstructorMk1/Build_ConstructorMk1.Build_ConstructorMk1_C"
	Assembler           TypePath = "/Game/FactoryGame/Buildable/Factory/AssemblerMk1/Build_AssemblerMk1.Build_AssemblerMk1_C"
	Manufacturer        TypePath = "/Game/FactoryGame/Buildable/Factory/ManufacturerMk1/Build_ManufacturerMk1.Build_ManufacturerMk1_C"
	Packager            TypePath = "/Game/FactoryGame/Buildable/Factory/Packager/Build_Packager.Build_Packager_C"
	Refinery            TypePath = "/Game/FactoryGame/Buildable/Factory/OilRefinery/Build_OilRefinery.Build_OilRefinery_C"
	Blender             TypePath = "/Game/FactoryGame/Buildable/Factory/Blender/Build_Blender.Build_Blender_C"
	ParticleAccelerator TypePath = "/Game/FactoryGame/Buildable/Factory/HadronCollider/Build_HadronCollider.Build_HadronCollider_C"

	Smelter TypePath = "/Game/FactoryGame/Buildable/Factory/SmelterMk1/Build_SmelterMk1.Build_SmelterMk1_C"
	Foundry TypePath = "/Game/FactoryGame/Buildable/Factory/FoundryMk1/Build_FoundryMk1.Build_FoundryMk1_C"

	//
	// Extraction
	//

	MinerMk1 TypePath = "/Game/FactoryGame/Buildable/Factory/MinerMK1/Build_MinerMk1.Build_MinerMk1_C"
	MinerMk2 TypePath = "/Game/FactoryGame/Buildable/Factory/MinerMK2/Build_MinerMk2.Build_MinerMk2_C"
	MinerMk3 TypePath = "/Game/FactoryGame/Buildable/Factory/MinerMK3/Build_MinerMk3.Build_MinerMk3_C"

	WaterExtractor          TypePath = ""
	OilExtractor            TypePath = ""
	ResourceWellPressurizer TypePath = ""
	ResourceWellExtractor   TypePath = ""

	//
	// Power
	//

	BiomassBurner       TypePath = "/Game/FactoryGame/Buildable/Factory/GeneratorBiomass/Build_GeneratorBiomass.Build_GeneratorBiomass_C"
	CoalGenerator       TypePath = "/Game/FactoryGame/Buildable/Factory/GeneratorCoal/Build_GeneratorCoal.Build_GeneratorCoal_C"
	FuelGenerator       TypePath = "/Game/FactoryGame/Buildable/Factory/GeneratorFuel/Build_GeneratorFuel.Build_GeneratorFuel_C"
	GeothermalGenerator TypePath = ""
	NuclearPowerPlant   TypePath = "/Game/FactoryGame/Buildable/Factory/GeneratorNuclear/Build_GeneratorNuclear.Build_GeneratorNuclear_C"

	PowerPoleMk1 TypePath = "/Game/FactoryGame/Buildable/Factory/PowerPoleMk1/Build_PowerPoleMk1.Build_PowerPoleMk1_C"
	PowerPoleMk2 TypePath = "/Game/FactoryGame/Buildable/Factory/PowerPoleMk2/Build_PowerPoleMk2.Build_PowerPoleMk2_C"
	PowerPoleMk3 TypePath = "/Game/FactoryGame/Buildable/Factory/PowerPoleMk3/Build_PowerPoleMk3.Build_PowerPoleMk3_C"

	// TODO: Wall outlets

	PowerSwitch  TypePath = ""
	PowerStorage TypePath = "/Game/FactoryGame/Buildable/Factory/PowerStorage/Build_PowerStorageMk1.Build_PowerStorageMk1_C"

	//
	// Transportation
	//

	TruckStation TypePath = "/Game/FactoryGame/Buildable/Factory/TruckStation/Build_TruckStation.Build_TruckStation_C"
	DronePort    TypePath = "/Game/FactoryGame/Buildable/Factory/DroneStation/Build_DroneStation.Build_DroneStation_C"

	// TODO: Trains

	//
	// Other
	//

	WorkBench TypePath = "/Game/FactoryGame/Buildable/Factory/WorkBench/Build_WorkBench.Build_WorkBench_C"
	Workshop  TypePath = "/Game/FactoryGame/Buildable/Factory/Workshop/Build_Workshop.Build_Workshop_C"
)
