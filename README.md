## SnapContractGo

Golang Unit Testing library to ensuring API/JSON based contracts using snapshots.

<h1 align="center">
<img src="./resources/snapcontract.png" with="300pt" height="300pt"></img>
</h1>

## Benefits

- Ignore changing values in snapshots 
- First class API JSON support
- Extends existing unit and integration tests

## How

1. Compare current and recorded responses from the API 
2. Autoupdate response snapshots for visibility.
3. Report errors if responses formats doesn't match
 
## Getting started

```golang
func TestAPIContract(t *testing.T) {
    var apiResponse interface{}
    // check that the result is the same as the last time the snapshot was updated
    snapcontract.SnapshotT(t, apiResponseast)
}
```

### Update a snapshot

```bash
UPDATE_SNAPSHOTS=true go test ./...
```
This will fail all tests where the snapshot was updated (to stop you accidentally updating snapshots in CI) but your snapshot files will now have been updated to reflect the current output of your code.

## Credits

Fork of https://github.com/bradleyjkemp/cupaloy
