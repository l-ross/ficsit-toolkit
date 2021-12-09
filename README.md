# FICSIT Toolkit

FICSIT Toolkit provides a collection of libraries for parsing a [Satisfactory](https://www.satisfactorygame.com/) save game file.

## Libraries

- save - A low level parser / serializer for Satisfactory save files.
  - You probably want to use the factory package.
- factory - Higher level parser for Satisfactory save files that understands and loads game state.
- resource - Auto generated structs based on the Satisfactory docs.json file

## Notes on API and version guarantees

- Satisfactory is still in early access as such my only interest is maintaining support for the latest
release on the early access branch.
- At this point I'm still experimenting with stuff so the API may change entirely and without notice.