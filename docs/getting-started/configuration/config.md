---
title:  Config File
description:  Config File
summary: Adding an optional config file
author: tfec
tags: [configuration]
---

The tfsec config file can override various tfsec configurations.

The tfsec config file is a file in the `.tfsec` folder in the root check path named `config.json` or `config.yml` and is automatically loaded if it exists.

The config file can also be set with the `--config-file` option:

```
tfsec --config-file tfsec.yml
```

## Syntax and Overrides

### Severity Overrides

There are occasions where the default severity level for one of the built in checks is too severe or in some cases not strong enough.

The config file can be used to specify overrides for any check identifier to replace the result output.

```json
{
  "severity_overrides": {
    "CUS002": "ERROR",
    "aws-s3-enable-versioning": "LOW"
  }
}
```

or in yaml

```yaml
---
severity_overrides:
  CUS002: ERROR
  aws-s3-enable-versioning: HIGH
```

### Excluding checks

There are moments where the list of checks you'd want to exclude becomes larger and larger.
Rather than passing all the excluded checks via the command line, you can use the configuration
entry `exclude` to list them all out.

```json
{
  "exclude": ["CUS002", "aws-s3-enable-versioning"]
}
```

or in yaml

```yaml
---
exclude:
  - CUS002
  - aws-s3-enable-versioning
```

### Minimum required version

For your CI you might want to pull a config file into all of your build processes with a centrally managed config file. If this is the case, you might also want to require a minimum tfsec version to be used.

This can be achieved in the config file using the `min_required_version` setting. 


```json
{
  "min_required_version": "v1.1.2"
}

or in yaml

```yaml
---
min_required_version: v1.1.2
```