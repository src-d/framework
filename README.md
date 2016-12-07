# Framework

## Description

Base package for src-d codebase.

## Modeling your structs

Our intention is to use Go structs as the source of truth for our data models.
We plan to generate
[proto3](https://developers.google.com/protocol-buffers/docs/proto3) files from
Go structs. Thus, we need to follow some recommendation to simplify the process
and ensure the output is right.

Things to consider:

* Try to use only basic types and enum-ish constructions in Go.
* Try to avoid maps: though not a limitation of the tool nor the spec, it can be hard to
  work with in other parts of the projects like a GraphQL API. In case you use
  one, try to include the key as part of the value.
  not work.

Things to consider while we try to fix them (any of these can be upgraded to
a plan _thing to consider_ after further research):

* Try not to use nested structs: currently the generated output for this does
  not work as expected and the fix might be difficult or unusable.
* Beware of embeds: `go2idl` generates the un-shortcut version and the fix
  might require redoing the whole tool (feature still in research).
* Try not to use pointers, the part of the tool that processes `gogo/protobuf`
  back does not know which pointers to remove and which pointers to leave, and
  it assumes that all need to be removed.

## Using protobuffers

Current plan is using protobuffers for communication only, please, take into
account that if we plan to use them for storage of any kind, some problems may
arise. Especially, the problem in case the generated proto file changes the
index of a field.

We need to come up with some solution for this.
