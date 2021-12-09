# resource

This package provides the Docs.json data that ships with Satisfactory code gen'd in to a collection
of packages.

All data in these packages is the property of Coffee Stain Studios and can be removed from this
repo upon request

## Code Gen

To re-run the code gen grab the latest Docs.json and run the following from within the gen directory:
```shell
go run *.go --build BUILD_NUMBER --file PATH/TO/Docs.json
```