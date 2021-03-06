## kev render

Generates application's deployment artefacts according to the specified output format for a given environment (ALL environments by default).

### Synopsis

(render) render Kubernetes manifests in selected format.

Examples:

  ### Render an app Kubernetes manifests (default) for all environments
  $ kev render

  ### Render an app Kubernetes manifests (default) for a specific environment(s)
  $ kev render -e staging [-e production ...]

```
kev render [flags]
```

### Options

```
  -f, --format string         Deployment files format. Default: Kubernetes manifests. (default "kubernetes")
  -s, --single                Controls whether to produce individual manifests or a single file output. Default: false
  -d, --dir string            Override default Kubernetes manifests output directory. Default: k8s/<env>
  -e, --environment strings   Target environment for which deployment files should be rendered
  -h, --help                  help for render
```

### SEE ALSO

* [kev](kev.md)	 - Develop Kubernetes apps iteratively using Docker-Compose.

###### Auto generated by spf13/cobra on 29-Oct-2020
