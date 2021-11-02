package main

import (
    "fmt"
    "os"
    "crypto/sha256"
    "encoding/json"
    "sort"

    "sigs.k8s.io/kustomize/kyaml/fn/framework"
    "sigs.k8s.io/kustomize/kyaml/fn/framework/command"
    "sigs.k8s.io/kustomize/kyaml/kio"
    "sigs.k8s.io/kustomize/kyaml/yaml"
)

type Spec struct {
    Enable bool `yaml:"enableSuffix" json:"enableSuffix"`
}

type SecretProviderClassHasher struct {
    Spec Spec `yaml:"spec" json:"spec"`
}

//func encodeSecretProviderClass(node *yaml.RNode) (string, err)