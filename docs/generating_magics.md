# Generating magic numbers

## To run

To run the magic number generator:
```
go run ./cmd/magics/main.go
```
This will spin up a number of workers depending on the machine you're using, that generate magic numbers for bishop and rooks. The numbers will get written out to `magic/magics.json`.

## Updating the magic numbers

If you are wanting to make a change to the magic numbers stored in this repository, submit a pull request, making sure to update the version string in `magic/version.txt` and the changelog in `magic/changelog.md`. The versioning works as follows:

### Major

- Changing the JSON structure
- Changing how the magic numbers are formatted (e.g., switching from hex strings to decimal)
- Adding required new fields
- Removing fields
- Changing the meaning of existing fields

### Minor

- Adding optional new fields (e.g., adding statistics about table sizes)

### Patch

- Finding better magic numbers that give smaller table sizes
- Fixing incorrect magic numbers
- Fixing typos (without changing the JSON structure)
- Updates to metadata like generation date

## What are magic numbers?

I'll get to this!
