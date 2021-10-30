# Resource Picker

Picks Only what You Want

---

## How to Use

`cd` to the [src](src) directory and execute,

`go run main.go gen .`

This will generate the Dockerfile within the src directory. Once generated please make sure that the golang version in
the Dockerfile is 1.17. Also you can make use of the pre-built Dockerfile under [gen](gen).

Then, build and tag the docker image giving any tag you like.

```
ex: docker build . -t example.com/resourcepicker:1.0.0
```

Use this docker image in your function specification.

### ResourcePicker Definition

The definition file needs to look like something shown below. The important fields are `annotations` and `kind`.

```
apiVersion: tarnsformers.api/v1              # this is arbitary
kind: ResourcePicker
metadata:
  name: resourcepicker                       # this is arbitary
  annotations:
    config.kubernetes.io/function: |
      container:
        image: <Docker image name with tag>  # Fill with your image name
spec:
  resourceNames:                             # Specify names of the resources
    - deployment1                            # you want in your output
  resourceKinds:                             # Specify kinds of the resources
    - Namespace                              # you want in your output
```

The above specification will output the resources whose names consists of the names listed under `resourceNames`
or whose kind is equal to what is listed under `resourceKinds`.

### Build the sample

Simply execute

```
kustomize build --enable-alpha-plugins > output.yaml
```

within the [sample](sample) directory after replacing the docker image name [here](sample/resourcepicker.yaml#L8). Now
you will see that the output only contains the two namespaces and the deployment1.