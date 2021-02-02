# ssm-bypass
This binary will export the path of SSM ParameterStore as an environment variables.
For example, it can be used in entrypoint.sh.

```
/your/pass/HOGE store hoge
/your/pass/FUGA store fuga
```
to
```
export HOGE=hoge
export FUGA=fuga
```

# usage
1. copy binary for your system from bin
```
$(ssm-bypass /your/path/)
```


# development
go get github.com/mitchellh/gox
gox --output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
