# FICSIT Toolkit

FICSIT Toolkit provides a collection of Go libraries for parsing a [Satisfactory](https://www.satisfactorygame.com/) save game file.

🚧 This repo is still a massive work in progress. 🚧

## Libraries

- save - A low level parser / serializer for Satisfactory save files.
- factory - Higher level parser for Satisfactory save files that understands and loads game state.
  - Huge WIP can currently only load some building types and understand conveyor connections
- resource - Auto generated structs based on the Satisfactory docs.json file

## Notes on API and version guarantees

- Satisfactory is still in early access as such my only interest is maintaining support for the latest
release on the early access branch.
- At this point I'm still experimenting with stuff so the API may change entirely and without notice.

## Other Packages

These packages were immensely helpful in understanding the Satisfactory save file structures:
- [SatisfactorySaveEditor](https://github.com/Goz3rr/SatisfactorySaveEditor)
- [satisfactory-json](https://github.com/ficsit-felix/satisfactory-json)