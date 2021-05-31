# mschartgen

Generate an org chart from the microsoft graph api

This project uses <https://github.com/dabeng/OrgChart> for generating the org chart.

## Usage

### Export Bearer Access Token

You can get this without having to create an app in Azure.

Browse to https://developer.microsoft.com/en-us/graph/graph-explorer and login.

You can then copy the Access Token for usage with this program, export in your shell env:

```
export MSGRAPH_BEARER_TOKEN=thisisquitelong
```

### Get UID

Get the UID of your CEO or the person you'd like to generate the chart down from.

You can get this in delve by navigating to the person and getting the id from the URL, e.g.

https://eur.delve.office.com/?u=2a6e0631-9431-4aed-b250-87b5d03a9fde (the value of `u`)

### Run

`go run cmd/mschartgen/mschartgen.go 2a6e0621-9441-4aed-b750-87c5d03a9fde`

## License

- Author: Chris Fordham (<chris@fordham.id.au>)

```text
Copyright 2021, Chris Fordham

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
