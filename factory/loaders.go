package factory

import "github.com/l-ross/ficsit-toolkit/save"

type loader func(f *Factory, b *building, s *save.Save) (Building, error)

var prioritisedLoading = []map[string]loader{
	{
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C": (*Factory).loadConveyorBelt,
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk2/Build_ConveyorBeltMk2.Build_ConveyorBeltMk2_C": (*Factory).loadConveyorBelt,
	},
	{
		"/Game/FactoryGame/Buildable/Factory/CA_Merger/Build_ConveyorAttachmentMerger.Build_ConveyorAttachmentMerger_C":       (*Factory).LoadSplitter,
		"/Game/FactoryGame/Buildable/Factory/CA_Splitter/Build_ConveyorAttachmentSplitter.Build_ConveyorAttachmentSplitter_C": (*Factory).LoadMerger,
		"/Game/FactoryGame/Buildable/Factory/StorageContainerMk2/Build_StorageContainerMk2.Build_StorageContainerMk2_C":       (*Factory).LoadStorageContainer,
		"/Game/FactoryGame/Buildable/Factory/ConstructorMk1/Build_ConstructorMk1.Build_ConstructorMk1_C":                      (*Factory).loadConstructor,
	},
}
