package main

import (
  "os"
  "strings"
  "sigs.k8s.io/kustomize/kyaml/fn/framework"
  "sigs.k8s.io/kustomize/kyaml/fn/framework/command"
  "sigs.k8s.io/kustomize/kyaml/kio"
  "sigs.k8s.io/kustomize/kyaml/yaml"
)

type Spec struct {
  Names []string `yaml:"resourceNames" json:"resourceNames"`
  Kinds []string `yaml:"resourceKinds,omitempty" json:"kinds,omitempty"`
}

// The yaml will look like the following
// ---
// kind: ResourcePicker
// spec:
//   resourceNames:
//   - list_goes_on
//   resourceKinds:
//   - Namespace
//   - Deployment //etc.
// ---
type ResourcePicker struct {
  Spec Spec `yaml:"spec,omitempty" json:"spec,omitempty"`
}

// Function stringInArray will return true if,
// a given string 'a' is a substring of
// any element in the provided string list

func stringInArray(a string, list []string) bool {
  for _, b := range list {
    if (strings.Contains(a, b)) {
      return true
    }
  }
  return false
}

// Function stringInArrayExact will return true if,
// a given string 'a' is an exact match of
// any element in the provided string list

func stringInArrayExact(a string, list []string) bool {
  for _, b := range list {
    if a == b {
      return true
    }
  }
  return false
}

// This function will filter out the resources by
// its name and/or kind
func main() {
  config := new(ResourcePicker)
  fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
    var outNodes []*yaml.RNode
    for i := range items {
      kind := items[i].GetKind()
      meta, err := items[i].GetMeta()
      if err != nil {
        return nil, err
      }
      if stringInArrayExact(kind, config.Spec.Kinds) || stringInArray(meta.Name, config.Spec.Names) {
        outNodes = append(outNodes, items[i])
      }
    }
    return outNodes, nil
  }
  p := framework.SimpleProcessor{Config: config, Filter: kio.FilterFunc(fn)}
  cmd := command.Build(p, command.StandaloneDisabled, false)
  command.AddGenerateDockerfile(cmd)
  if err := cmd.Execute(); err != nil {
    os.Exit(1)
  }
}