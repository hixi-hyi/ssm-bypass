# ssm-bypass
This binary will export the path of SSM ParameterStore as an environment variables.
For example, it can be used in entrypoint.sh.

The following are the registered contents of ssm-parameter and the values output by ssm-bypass.
* ssm-parameter
```
/your/pass/HOGE → hoge
/your/pass/FUGA → fuga
```

* `ssm-bypass /your/pass`
```
export HOGE=hoge
export FUGA=fuga
```

# usage
1. copy binary for your system from `bin/`
2. write the bellow code into your entrypoint.sh
```
$(ssm-bypass /your/path/)
```

# causion
* If the value of ssm parameter contains shell's special characters such as single quotes, back quotes and dollar sign, please verify the operation thoroughly before using it. It has not been fully verified.
* Multi-line values are not supported.


# development
```
rm -rf bin/
go get github.com/mitchellh/gox
gox --output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
```
